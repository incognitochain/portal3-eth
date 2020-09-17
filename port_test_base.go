package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/incognitochain/bridge-eth/blockchain"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
	"github.com/incognitochain/bridge-eth/jsonresult"
	"github.com/incognitochain/portal3-eth/portal/portalv3"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strings"
	"time"

	// "github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/suite"
	// "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/portal3-eth/portal/delegator"
	"github.com/incognitochain/portal3-eth/portal/erc20"
	"github.com/incognitochain/portal3-eth/portal/erc20/bnb"
	"github.com/incognitochain/portal3-eth/portal/erc20/dai"
	"github.com/incognitochain/portal3-eth/portal/erc20/dless"
	"github.com/incognitochain/portal3-eth/portal/erc20/fail"
	"github.com/incognitochain/portal3-eth/portal/erc20/usdt"
	"github.com/incognitochain/portal3-eth/portal/incognitoproxy"
	"github.com/pkg/errors"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
var auth *bind.TransactOpts
var genesisAcc *account

const (
	IncPaymentAddr    = "12S5YAYzDr75NJWa4YoUDB9d2c2GKQMREtkRL19JrnZngnkwahP6aBDz7qnKHBGLmDpp2FUA7yqa58fuMqoaUHtE97FRvReBMssWMWm"
	DepositERC20Topic = "0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e"
)

type account struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

type TokenerInfo struct {
	addr common.Address
	c    Tokener
}

type Tokener interface {
	BalanceOf(*bind.CallOpts, common.Address) (*big.Int, error)
	Approve(*bind.TransactOpts, common.Address, *big.Int) (*types.Transaction, error)
}

type contracts struct {
	delegatorAddr common.Address
	portalV3Ins   *portalv3.Portalv3
	portalv3      common.Address
	inc           *incognitoproxy.Incognitoproxy
	incAddr       common.Address
	token         *erc20.Erc20
	tokenAddr     common.Address

	tokens       map[int]tokenInfo       // mapping from dcommonimal => token
	customErc20s map[string]*TokenerInfo // mapping from name => token
}

type tokenInfo struct {
	c    *erc20.Erc20
	addr common.Address
}

type Platform struct {
	*contracts
	sim *backends.SimulatedBackend
}

type committees struct {
	beacons     []common.Address
	bridges     []common.Address
	beaconPrivs [][]byte
	bridgePrivs [][]byte
}

type portalV3Base struct {
	sc *contracts
	p  *Platform
	c  *committees
}

type getProofResult struct {
	Result jsonresult.GetInstructionProof
	Error  struct {
		Code       int
		Message    string
		StackTrace string
	}
}

type decodedProof struct {
	Instruction []byte
	Heights     [2]*big.Int

	InstPaths       [2][][32]byte
	InstPathIsLefts [2][]bool
	InstRoots       [2][32]byte
	BlkData         [2][32]byte
	SigIdxs         [2][]*big.Int
	SigVs           [2][]uint8
	SigRs           [2][][32]byte
	SigSs           [2][][32]byte
}

type instProof struct {
	isBeacon       bool
	instHash       [32]byte
	blkHeight      *big.Int
	instPath       [][32]byte
	instPathIsLeft []bool
	instRoot       [32]byte
	blkData        [32]byte
	sigIdx         []*big.Int
	sigV           []uint8
	sigR           [][32]byte
	sigS           [][32]byte
}

type merklePath struct {
	merkles [][]byte
	leaf    [32]byte
	root    [32]byte
	path    [][32]byte
	left    []bool
}

func init() {
	fmt.Println("Initializing genesis account...")
	genesisAcc = loadAccount()
	auth = bind.NewKeyedTransactor(genesisAcc.PrivateKey)
}

func loadAccount() *account {
	key, err := crypto.LoadECDSA("genesisKey.hex")
	if err != nil {
		return newAccount()
	}
	return &account{
		PrivateKey: key,
		Address:    crypto.PubkeyToAddress(key.PublicKey),
	}
}

func setupFixedCommittee(accs ...common.Address) (*Platform, *committees, error) {
	c := getFixedCommittee()
	p, err := setup(c.beacons, c.bridges, accs...)
	return p, c, err
}

// getFixedCommittee is for unittest
func getFixedCommittee() *committees {
	beaconCommPrivs := []string{
		"aad53b70ad9ed01b75238533dd6b395f4d300427da0165aafbd42ea7a606601f",
		"ca71365ceddfa8e0813cf184463bd48f0b62c9d7d5825cf95263847628816e82",
		"1e4d2244506211200640567630e3951abadbc2154cf772e4f0d2ff0770290c7c",
		"c7146b500240ed7aac9445e2532ae8bf6fc7108f6ea89fde5eebdf2fb6cefa5a",
	}
	beaconComm := []string{
		"0xD7d93b7fa42b60b6076f3017fCA99b69257A912D",
		"0xf25ee30cfed2d2768C51A6Eb6787890C1c364cA4",
		"0x0D8c517557f3edE116988DD7EC0bAF83b96fe0Cb",
		"0xc225fcd5CE8Ad42863182Ab71acb6abD9C4ddCbE",
	}
	beaconPrivs := make([][]byte, len(beaconCommPrivs))
	for i, p := range beaconCommPrivs {
		priv, _ := hex.DecodeString(p)
		beaconPrivs[i] = priv
	}

	bridgeComm := []string{
		"0x3c78124783E8e39D1E084FdDD0E097334ba2D945",
		"0x76E34d8a527961286E55532620Af5b84F3C6538F",
		"0x68686dB6874588D2404155D00A73F82a50FDd190",
		"0x1533ac4d2922C150551f2F5dc2b0c1eDE382b890",
	}
	bridgeCommPrivs := []string{
		"3560e649ce326a2eb9fbb59fba4b29e10fb064627f61487acommonc8b92afbb127dd",
		"b71af1a7e2ca74400187cbf2333ab1f20e9b39517347fb655ffa309d1b51b2b0",
		"07f91f98513c203103f8d44683ce47920d1aea0eaf1cb86a373be835374d1490",
		"7412e24d4ac1796866c44a0d5b966f8db1c3022bba8afd370a09dc49a14efeb4",
	}

	bridgePrivs := make([][]byte, len(bridgeCommPrivs))
	for i, p := range bridgeCommPrivs {
		priv, _ := hex.DecodeString(p)
		bridgePrivs[i] = priv
	}

	beacons, bridges := toAddresses(beaconComm, bridgeComm)
	return &committees{
		beacons:     beacons,
		beaconPrivs: beaconPrivs,
		bridges:     bridges,
		bridgePrivs: bridgePrivs,
	}
}

func toAddresses(beaconComm, bridgeComm []string) ([]common.Address, []common.Address) {
	beacons := make([]common.Address, len(beaconComm))
	for i, p := range beaconComm {
		beacons[i] = common.HexToAddress(p)
	}

	bridges := make([]common.Address, len(bridgeComm))
	for i, p := range bridgeComm {
		bridges[i] = common.HexToAddress(p)
	}
	return beacons, bridges
}

func setup(
	beaconComm []common.Address,
	bridgeComm []common.Address,
	accs ...common.Address,
) (*Platform, error) {
	alloc := make(core.GenesisAlloc)
	balance, _ := big.NewInt(1).SetString("1000000000000000000000000000000", 10) // 1E30 wei
	alloc[auth.From] = core.GenesisAccount{Balance: balance}
	for _, acc := range accs {
		alloc[acc] = core.GenesisAccount{Balance: balance}
	}
	sim := backends.NewSimulatedBackend(alloc, 8000000)
	p := &Platform{
		sim: sim,
		contracts: &contracts{
			tokens:       map[int]tokenInfo{},
			customErc20s: map[string]*TokenerInfo{},
		},
	}

	var err error
	var tx *types.Transaction
	_ = tx

	// ERC20: always deploy first so its address is fixed
	p.tokenAddr, tx, p.token, err = erc20.DeployErc20(auth, sim, "MyErc20", "ERC", big.NewInt(0), big.NewInt(int64(1e18)))
	if err != nil {
		return nil, fmt.Errorf("failed to deploy ERC20 contract: %v", err)
	}
	// fmt.Printf("token addr: %s\n", p.tokenAddr.Hex())
	sim.Commit()

	// Custom tokens
	err = setupCustomTokens(p)
	if err != nil {
		return nil, err
	}

	// IncognitoProxy
	admin := auth.From
	p.incAddr, tx, p.inc, err = incognitoproxy.DeployIncognitoproxy(auth, sim, admin, beaconComm)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy IncognitoProxy contract: %v", err)
	}
	sim.Commit()

	p.portalv3, tx, _, err = portalv3.DeployPortalv3(auth, sim)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy Portal contract: %v", err)
	}
	sim.Commit()

	// Portal
	p.delegatorAddr, _, _, err = delegator.DeployDelegator(auth, sim, auth.From, p.portalv3, p.incAddr)
	if err != nil {
		return nil, err
	}
	sim.Commit()

	p.portalV3Ins, err = portalv3.NewPortalv3(p.delegatorAddr, sim)
	if err != nil {
		return nil, fmt.Errorf("failed to assgin portal contract to delegator address: %v", err)
	}

	return p, nil
}

func setupCustomTokens(p *Platform) error {
	// Deploy BNB
	bal, _ := big.NewInt(1).SetString("200000000000000000000000000", 10)
	addr, _, bnb, err := bnb.DeployBnb(auth, p.sim, bal, "BNB", uint8(18), "BNB")
	if err != nil {
		return errors.Errorf("failed to deploy BNB contract: %v", err)
	}
	p.sim.Commit()
	p.contracts.customErc20s["BNB"] = &TokenerInfo{addr: addr, c: bnb}

	// Deploy USDT
	bal, _ = big.NewInt(1).SetString("100000000000", 10)
	addr, _, usdt, err := usdt.DeployUsdt(auth, p.sim, bal, "Tether USD", "USDT", big.NewInt(6))
	if err != nil {
		return errors.Errorf("failed to deploy USDT contract: %v", err)
	}
	p.sim.Commit()
	p.contracts.customErc20s["USDT"] = &TokenerInfo{addr: addr, c: usdt}

	// Deploy DAI
	symbol := [32]byte{'D', 'A', 'I'}
	addr, _, d, err := dai.DeployDai(auth, p.sim, symbol)
	if err != nil {
		return errors.Errorf("failed to deploy DAI contract: %v", err)
	}
	p.sim.Commit()
	p.contracts.customErc20s["DAI"] = &TokenerInfo{addr: addr, c: d}

	// Mint DAI
	bal, _ = big.NewInt(1).SetString("1000000000000000000000000000", 10)
	_, err = d.Mint(auth, bal)
	if err != nil {
		return errors.Errorf("failed to mint DAI: %v", err)
	}
	p.sim.Commit()

	// // Deploy USDC
	// // symbol := [32]byte{'D', 'A', 'I'}
	// addr, _, dc, err := usdc.DeployUsdc(auth, p.sim)
	// if err != nil {
	// 	fmt.Println("ASDASD", err)
	// 	return errors.Errorf("failed to deploy USDC contract: %v", err)
	// }
	// p.sim.Commit()
	// p.contracts.customErc20s["USDC"] = &TokenerInfo{c: dc}

	// // Deploy USDC wrapper
	// addr, _, _, err = usdc_wrap.DeployUsdcWrap(auth, p.sim, addr)
	// if err != nil {
	// 	fmt.Println("!@(*#&!@*(#&", err)
	// 	return errors.Errorf("failed to deploy USDCWrap contract: %v", err)
	// }
	// p.sim.Commit()
	// p.contracts.customErc20s["USDC"].addr = addr

	// Deploy FAIL token
	bal, _ = big.NewInt(1).SetString("1000000000000000000", 10)
	addr, _, fail, err := fail.DeployFAIL(auth, p.sim, bal, "FAIL", 6, "FAIL")
	if err != nil {
		return errors.Errorf("failed to deploy FAIL contract: %v", err)
	}
	p.sim.Commit()
	p.contracts.customErc20s["FAIL"] = &TokenerInfo{addr: addr, c: fail}

	// Deploy DLESS token
	bal, _ = big.NewInt(1).SetString("1000000000000000000", 10)
	addr, _, dless, err := dless.DeployDLESS(auth, p.sim, bal, "DLESS", "DLESS")
	if err != nil {
		return errors.Errorf("failed to deploy DLESS contract: %v", err)
	}
	p.sim.Commit()
	p.contracts.customErc20s["DLESS"] = &TokenerInfo{addr: addr, c: dless}

	return nil
}

func deployERC20(decimals []int, p *Platform) error {
	// Deploy erc20s with different decimals to test
	ercBal := big.NewInt(20)
	ercBal = ercBal.Mul(ercBal, big.NewInt(int64(1e18)))
	ercBal = ercBal.Mul(ercBal, big.NewInt(int64(1e18)))
	for _, d := range decimals {
		tokenAddr, _, token, err := erc20.DeployErc20(auth, p.sim, "MyErc20", "ERC", big.NewInt(int64(d)), ercBal)
		if err != nil {
			return fmt.Errorf("failed to deploy ERC20 contract: %v", err)
		}
		p.tokens[d] = tokenInfo{c: token, addr: tokenAddr}
	}
	p.sim.Commit()
	return nil
}

func newAccount() *account {
	key, _ := crypto.GenerateKey()
	// crypto.SaveECDSA("genesisKey.hex", key)
	return &account{
		PrivateKey: key,
		Address:    crypto.PubkeyToAddress(key.PublicKey),
	}
}

//func setupFixedERC20s(decimals []int) (*Platform, *committees, error) {
//	c := getFixedCommittee()
//	p, err := setup(c.beacons, c.bridges, decimals)
//	return p, c, err
//}

func getBalanceERC20(token Tokener, addr common.Address) *big.Int {
	bal, err := token.BalanceOf(nil, addr)
	if err != nil {
		return big.NewInt(-1)
	}
	return bal
}

func lockSimERC20WithTxs(
	p *Platform,
	token Tokener,
	tokenAddr common.Address,
	amount *big.Int,
) (*types.Transaction, *types.Transaction, error) {
	txApprove, err := approveERC20(genesisAcc.PrivateKey, p.delegatorAddr, token, amount)
	if err != nil {
		return nil, nil, err
	}
	p.sim.Commit()

	txDeposit, err := depositERC20(genesisAcc.PrivateKey, p.portalV3Ins, tokenAddr, amount)
	if err != nil {
		return txApprove, nil, err
	}
	p.sim.Commit()
	return txApprove, txDeposit, nil
}

func approveERC20(privKey *ecdsa.PrivateKey, spender common.Address, token Tokener, amount *big.Int) (*types.Transaction, error) {
	return approveERC20Detail(
		privKey,
		spender,
		token,
		amount,
		0,
		0,
		nil,
	)
}

func approveERC20Detail(
	privKey *ecdsa.PrivateKey,
	spender common.Address,
	token Tokener,
	amount *big.Int,
	nonce uint64,
	gasLimit uint64,
	gasPrice *big.Int,
) (*types.Transaction, error) {
	// Check balance
	// userAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	// bal, _ := token.BalanceOf(nil, userAddr)
	// fmt.Printf("erc20 balance: %d\n", bal)

	// Approve
	auth := bind.NewKeyedTransactor(privKey)
	if gasLimit > 0 {
		auth.GasLimit = gasLimit
	}
	if gasPrice != nil {
		auth.GasPrice = gasPrice
	}
	if nonce > 0 {
		auth.Nonce = big.NewInt(int64(nonce))
	}

	tx, err := token.Approve(auth, spender, amount)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// txHash := tx.Hash()
	// fmt.Printf("erc20 approved, txHash: %x\n", txHash[:])
	return tx, nil
}

func depositERC20(
	privKey *ecdsa.PrivateKey,
	v *portalv3.Portalv3,
	tokenAddr common.Address,
	amount *big.Int,
) (*types.Transaction, error) {
	return depositERC20Detail(
		privKey,
		v,
		tokenAddr,
		amount,
		IncPaymentAddr,
		0,
		0,
		nil,
	)
}

func depositERC20Detail(
	privKey *ecdsa.PrivateKey,
	v *portalv3.Portalv3,
	tokenAddr common.Address,
	amount *big.Int,
	incPaymentAddr string,
	nonce uint64,
	gasLimit uint64,
	gasPrice *big.Int,
) (*types.Transaction, error) {
	auth := bind.NewKeyedTransactor(privKey)
	if gasLimit > 0 {
		auth.GasLimit = gasLimit
	}
	if gasPrice != nil {
		auth.GasPrice = gasPrice
	}
	if nonce > 0 {
		auth.Nonce = big.NewInt(int64(nonce))
	}

	tx, err := v.DepositERC20(auth, tokenAddr, amount, incPaymentAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// txHash := tx.Hash()
	// fmt.Printf("erc20 deposited, txHash: %x\n", txHash[:])
	return tx, nil
}

func extractAmountInDepositERC20Event(sim *backends.SimulatedBackend, tx *types.Transaction) (*big.Int, error) {
	_, events, err := retrieveEvents(sim, tx)
	if err != nil {
		return nil, err
	}
	data, ok := events[DepositERC20Topic]
	if !ok {
		return nil, errors.Errorf("no erc20 deposit event found in tx %v", tx.Hash().Hex())
	}
	cAbi, err := abi.JSON(strings.NewReader(vault.VaultABI))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	e := struct {
		Token            common.Address
		IncognitoAddress string
		Amount           *big.Int
	}{}

	err = cAbi.Unpack(&e, "Deposit", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return e.Amount, nil
}

func retrieveEvents(sim *backends.SimulatedBackend, tx *types.Transaction) (*types.Receipt, map[string][]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	receipt, err := sim.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	if len(receipt.Logs) == 0 {
		fmt.Println("empty log")
		return receipt, nil, nil
	}

	events := map[string][]byte{}
	for _, log := range receipt.Logs {
		events[log.Topics[0].Hex()] = log.Data
	}
	return receipt, events, nil
}

func deposit(p *Platform, amount *big.Int) (*big.Int, *big.Int, error) {
	initBalance := p.getBalance(p.delegatorAddr)
	auth := bind.NewKeyedTransactor(genesisAcc.PrivateKey)
	auth.GasLimit = 0
	auth.Value = amount

	_, err := p.portalV3Ins.Deposit(auth, "")
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	p.sim.Commit()
	newBalance := p.getBalance(p.delegatorAddr)
	return initBalance, newBalance, nil
}

func (p *Platform) getBalance(addr common.Address) *big.Int {
	bal, _ := p.sim.BalanceAt(context.Background(), addr, nil)
	return bal
}

func buildWithdrawTestcase(c *committees, meta, shard int, tokenID common.Address, amount *big.Int, withdrawer common.Address) (*decodedProof, [32]byte) {
	inst, mp, blkData, blkHash := buildWithdrawData(meta, shard, tokenID, amount, withdrawer)
	ipBeacon := signAndReturnInstProof(c.beaconPrivs, true, mp, blkData, blkHash[:])
	return &decodedProof{
		Instruction: inst,
		Heights:     [2]*big.Int{big.NewInt(1), big.NewInt(1)},

		InstPaths:       [2][][32]byte{ipBeacon.instPath},
		InstPathIsLefts: [2][]bool{ipBeacon.instPathIsLeft},
		InstRoots:       [2][32]byte{ipBeacon.instRoot},
		BlkData:         [2][32]byte{ipBeacon.blkData},
		SigIdxs:         [2][]*big.Int{ipBeacon.sigIdx},
		SigVs:           [2][]uint8{ipBeacon.sigV},
		SigRs:           [2][][32]byte{ipBeacon.sigR},
		SigSs:           [2][][32]byte{ipBeacon.sigS},
	}, ipBeacon.instHash
}

func buildWithdrawData(meta, shard int, tokenID common.Address, amount *big.Int, withdrawer common.Address) ([]byte, *merklePath, []byte, []byte) {
	// Build instruction merkle tree
	numInst := 10
	startNodeID := 7
	height := big.NewInt(1)
	inst := buildDecodedWithdrawInst(meta, shard, tokenID, withdrawer, amount)
	instWithHeight := append(inst, toBytes32BigEndian(height.Bytes())...)
	data := randomMerkleHashes(numInst)
	data[startNodeID] = instWithHeight
	mp := buildInstructionMerklePath(data, numInst, startNodeID)

	// Generate random blkHash
	h := randomMerkleHashes(1)
	blkData := h[0]
	blkHash := rawsha3(append(blkData, mp.root[:]...))
	return inst, mp, blkData, blkHash[:]
}

func buildDecodedWithdrawInst(meta, shard int, tokenID, withdrawer common.Address, amount *big.Int) []byte {
	decoded := []byte{byte(meta)}
	decoded = append(decoded, byte(shard))
	decoded = append(decoded, toBytes32BigEndian(tokenID[:])...)
	decoded = append(decoded, toBytes32BigEndian(withdrawer[:])...)
	decoded = append(decoded, toBytes32BigEndian(amount.Bytes())...)
	txId := make([]byte, 32)
	rand.Read(txId)
	decoded = append(decoded, toBytes32BigEndian(txId)...) // txID
	decoded = append(decoded, make([]byte, 16)...)         // incTokenID, variable length
	return decoded
}

func signAndReturnInstProof(
	privs [][]byte,
	isBeacon bool,
	mp *merklePath,
	blkData []byte,
	blkHash []byte,
) *instProof {
	sigV := make([]uint8, len(privs))
	sigR := make([][32]byte, len(privs))
	sigS := make([][32]byte, len(privs))
	sigIdx := make([]*big.Int, len(privs))
	for i, p := range privs {
		sig, _ := bridgesig.Sign(p, blkHash)
		sigV[i] = uint8(sig[64] + 27)
		sigR[i] = toByte32(sig[:32])
		sigS[i] = toByte32(sig[32:64])
		sigIdx[i] = big.NewInt(int64(i))
	}

	return &instProof{
		isBeacon:       isBeacon,
		instHash:       mp.leaf,
		blkHeight:      big.NewInt(0),
		instPath:       mp.path,
		instPathIsLeft: mp.left,
		instRoot:       mp.root,
		blkData:        toByte32(blkData),
		sigIdx:         sigIdx,
		sigV:           sigV,
		sigR:           sigR,
		sigS:           sigS,
	}
}

func buildInstructionMerklePath(data [][]byte, numInst, startNodeID int) *merklePath {
	merkles := blockchain.BuildKeccak256MerkleTree(data)
	p, l := blockchain.GetKeccak256MerkleProofFromTree(merkles, startNodeID)
	path := [][32]byte{}
	left := []bool{}
	for i, x := range p {
		path = append(path, toByte32(x))
		left = append(left, l[i])
	}

	return &merklePath{
		merkles: merkles,
		leaf:    toByte32(merkles[startNodeID]),
		root:    toByte32(merkles[len(merkles)-1]),
		path:    path,
		left:    left,
	}
}

func Withdraw(v *portalv3.Portalv3, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	// auth.GasPrice = big.NewInt(20000000000)
	tx, err := v.WithdrawLockedTokens(
		auth,
		proof.Instruction,
		proof.Heights[0],

		proof.InstPaths[0],
		proof.InstPathIsLefts[0],
		proof.InstRoots[0],
		proof.BlkData[0],
		proof.SigIdxs[0],
		proof.SigVs[0],
		proof.SigRs[0],
		proof.SigSs[0],
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func toBytes32BigEndian(b []byte) []byte {
	a := copyToBytes32(b)
	return a[:]
}

func copyToBytes32(b []byte) [32]byte {
	a := [32]byte{}
	copy(a[32-len(b):], b)
	return a
}

func toByte32(s []byte) [32]byte {
	a := [32]byte{}
	copy(a[:], s)
	return a
}
func randomMerkleHashes(n int) [][]byte {
	h := [][]byte{}
	for i := 0; i < n; i++ {
		b := make([]byte, 32)
		rand.Read(b)
		h = append(h, b)
	}
	return h
}

func rawsha3(b []byte) []byte {
	hashF := sha3.NewLegacyKeccak256()
	hashF.Write(b)
	buf := hashF.Sum(nil)
	return buf
}

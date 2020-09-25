package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/incognitochain/bridge-eth/common/base58"
	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
	"github.com/incognitochain/portal3-eth/portal/portalv3"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/light"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/incognitochain/bridge-eth/rpccaller"
	"github.com/stretchr/testify/require"
)

type IssuingETHRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type BurningForDepositToSCRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type Receipt struct {
	Result *types.Receipt `json:"result"`
}

type NormalResult struct {
	Result interface{} `json:"result"`
}


// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type PortalV3BaseTestSuite struct {
	suite.Suite
	IncBurningAddrStr string
	IncPrivKeyStr     string
	IncPaymentAddrStr string

	GeneratedPrivKeyForSC ecdsa.PrivateKey
	GeneratedPubKeyForSC  ecdsa.PublicKey

	IncBridgeHost string
	IncRPCHost    string
	IncEtherTokenIDStr string

	EtherAddressStr string
	ETHPrivKeyStr   string
	ETHOwnerAddrStr string

	ETHHost    string
	ETHPrivKey *ecdsa.PrivateKey
	ETHClient  *ethclient.Client

	Portalv3            common.Address
	KBNTradeDeployedAddr common.Address

	KyberContractAddr common.Address
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (portalV3Suite *PortalV3BaseTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")

	// 0x kovan env
	portalV3Suite.IncBurningAddrStr = "15pABFiJVeh9D5uiQEhQX4SVibGGbdAVipQxBdxkmDqAJaoG1EdFKHBrNfs"
	portalV3Suite.IncPrivKeyStr = "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"
	portalV3Suite.IncPaymentAddrStr = "12S5Lrs1XeQLbqN4ySyKtjAjd2d7sBP2tjFijzmp6avrrkQCNFMpkXm3FPzj2Wcu2ZNqJEmh9JriVuRErVwhuQnLmWSaggobEWsBEci"
	// portalV3Suite.GeneratedPubKeyForSCStr = "8224890Cd5A537792d1B8B56c95FAb8a1A5E98B1"

	portalV3Suite.EtherAddressStr = "0x0000000000000000000000000000000000000000"
	portalV3Suite.IncEtherTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000099"
	portalV3Suite.ETHPrivKeyStr = "B8DB29A7A43FB88AD520F762C5FDF6F1B0155637FA1E5CB2C796AFE9E5C04E31"
	portalV3Suite.ETHOwnerAddrStr = "FD94c46ab8dCF0928d5113a6fEaa925793504e16"

	portalV3Suite.ETHHost = "https://kovan.infura.io/v3/93fe721349134964aa71071a713c5cef"
	portalV3Suite.IncBridgeHost = "http://127.0.0.1:9338"
	portalV3Suite.IncRPCHost = "http://127.0.0.1:9334"

	portalV3Suite.Portalv3 = common.HexToAddress("0x88D9531eCCDee7fDd2061D2053F92B1E00596804")

	// generate a new keys pair for SC
	portalV3Suite.genKeysPairForSC()

	// connect to ethereum network
	portalV3Suite.connectToETH()
}

func (portalV3Suite *PortalV3BaseTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	portalV3Suite.ETHClient.Close()
}

func (portalV3Suite *PortalV3BaseTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (portalV3Suite *PortalV3BaseTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

func (portalV3Suite *PortalV3BaseTestSuite) TestPortalV3BaseTestSuite() {
	fmt.Println("This is generic test suite")
}

func (portalV3Suite *PortalV3BaseTestSuite) getBalanceOnETHNet(
	tokenAddr common.Address,
	ownerAddr common.Address,
) *big.Int {
	if tokenAddr.Hex() == portalV3Suite.EtherAddressStr {
		balance, err := portalV3Suite.ETHClient.BalanceAt(context.Background(), ownerAddr, nil)
		require.Equal(portalV3Suite.T(), nil, err)
		return balance
	}
	// erc20 token
	instance, err := erc20.NewErc20(tokenAddr, portalV3Suite.ETHClient)
	require.Equal(portalV3Suite.T(), nil, err)

	balance, err := instance.BalanceOf(&bind.CallOpts{}, ownerAddr)
	require.Equal(portalV3Suite.T(), nil, err)
	return balance
}

func (portalV3Suite *PortalV3BaseTestSuite) connectToETH() {
	privKeyHex := portalV3Suite.ETHPrivKeyStr
	privKey, err := crypto.HexToECDSA(privKeyHex)
	require.Equal(portalV3Suite.T(), nil, err)

	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "development"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial(portalV3Suite.ETHHost)
	require.Equal(portalV3Suite.T(), nil, err)

	portalV3Suite.ETHClient = client
	portalV3Suite.ETHPrivKey = privKey
}

func (portalV3Suite *PortalV3BaseTestSuite) depositETH(
	amt float64,
	incPaymentAddrStr string,
) common.Hash {
	c, err := portalv3.NewPortalv3(portalV3Suite.Portalv3, portalV3Suite.ETHClient)
	require.Equal(portalV3Suite.T(), nil, err)

	auth := bind.NewKeyedTransactor(portalV3Suite.ETHPrivKey)
	auth.Value = big.NewInt(int64(amt * params.Ether))
	tx, err := c.Deposit(auth, incPaymentAddrStr)
	require.Equal(portalV3Suite.T(), nil, err)
	txHash := tx.Hash()

	if err := wait(portalV3Suite.ETHClient, txHash); err != nil {
		require.Equal(portalV3Suite.T(), nil, err)
	}
	fmt.Printf("deposited, txHash: %x\n", txHash[:])
	return txHash
}

func (portalV3Suite *PortalV3BaseTestSuite) depositERC20ToBridge(
	amt *big.Int,
	tokenAddr common.Address,
	incPaymentAddrStr string,
) common.Hash {
	auth := bind.NewKeyedTransactor(portalV3Suite.ETHPrivKey)
	c, err := portalv3.NewPortalv3(portalV3Suite.Portalv3, portalV3Suite.ETHClient)
	require.Equal(portalV3Suite.T(), nil, err)

	erc20Token, _ := erc20.NewErc20(tokenAddr, portalV3Suite.ETHClient)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 1000000
	tx2, apprErr := erc20Token.Approve(auth, portalV3Suite.Portalv3, amt)
	tx2Hash := tx2.Hash()
	fmt.Printf("Approve tx, txHash: %x\n", tx2Hash[:])
	require.Equal(portalV3Suite.T(), nil, apprErr)
	time.Sleep(15 * time.Second)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 1000000

	fmt.Println("Starting deposit erc20 to portalv3 contract")
	tx, err := c.DepositERC20(auth, tokenAddr, amt, incPaymentAddrStr)
	require.Equal(portalV3Suite.T(), nil, err)
	fmt.Println("Finished deposit erc20 to portlv3 contract")
	txHash := tx.Hash()

	if err := wait(portalV3Suite.ETHClient, txHash); err != nil {
		require.Equal(portalV3Suite.T(), nil, err)
	}
	fmt.Printf("deposited erc20 token to bridge, txHash: %x\n", txHash[:])
	return txHash
}

func (portalV3Suite *PortalV3BaseTestSuite) callCustodianDeposit(
	incTokenIDStr string,
	ethDepositProof []string,
	ethBlockHash string,
	ethTxIdx uint,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	remoteAddresses :=  map[string]interface{}{
		"6abd698ea7ddd1f98b1ecaaddab5db0453b8363ff092f0d8d7d4c6b1155fb693" : "tbnb1fau9kq605jwkyfea2knw495we8cpa47r9r6uxv",
	} 
	meta := map[string]interface{}{
		"RemoteAddresses": remoteAddresses,
		"BlockHash":  ethBlockHash,
		"ProofStrs":  ethDepositProof,
		"TxIndex":    ethTxIdx,
	}
	params := []interface{}{
		portalV3Suite.IncPrivKeyStr,
		nil,
		5,
		-1,
		meta,
	}
	var res IssuingETHRes
	err := rpcClient.RPCCall(
		"",
		portalV3Suite.IncRPCHost,
		"",
		"createandsendtxwithcustodiandepositv3",
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}

	response, _ := json.Marshal(res)
	fmt.Println("get response", string(response))

	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}

func (portalV3Suite *PortalV3BaseTestSuite) callUnlockCollateralToken(
	incTokenIDStr string,
	amount *big.Int,
	remoteAddrStr string,
	burningMethod string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"TokenID":     incTokenIDStr,
		"TokenTxType": 1,
		"TokenName":   "",
		"TokenSymbol": "",
		"TokenAmount": amount.Uint64(),
		"TokenReceivers": map[string]uint64{
			portalV3Suite.IncBurningAddrStr: amount.Uint64(),
		},
		"RemoteAddress": remoteAddrStr,
		"Privacy":       true,
		"TokenFee":      0,
	}
	params := []interface{}{
		portalV3Suite.IncPrivKeyStr,
		nil,
		5,
		-1,
		meta,
		"",
		0,
	}
	var res BurningForDepositToSCRes
	err := rpcClient.RPCCall(
		"",
		portalV3Suite.IncRPCHost,
		"",
		burningMethod,
		params,
		&res,
	)
	if err != nil {
		fmt.Println("calling burning ptokens err: ", err)
		return nil, err
	}
	bb, _ := json.Marshal(res)
	fmt.Println("calling burning ptokens res: ", string(bb))
	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}

func (portalV3Suite *PortalV3BaseTestSuite) genKeysPairForSC() {
	incPriKeyBytes, _, err := base58.Base58Check{}.Decode(portalV3Suite.IncPrivKeyStr)
	require.Equal(portalV3Suite.T(), nil, err)

	portalV3Suite.GeneratedPrivKeyForSC, portalV3Suite.GeneratedPubKeyForSC = bridgesig.KeyGen(incPriKeyBytes)
}

func randomizeTimestamp() string {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)
	return randomNow.String()
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func wait(client *ethclient.Client, tx common.Hash) error {
	ctx := context.Background()
	for range time.Tick(10 * time.Second) {
		_, err := client.TransactionReceipt(ctx, tx)
		if err == nil {
			break
		} else if err == ethereum.NotFound {
			continue
		} else {
			return err
		}
	}
	return nil
}

func getETHDepositProof(
	url string,
	txHash common.Hash,
) (*big.Int, string, uint, []string, error) {
	// Get tx content
	txContent, err := getETHTransactionByHash(url, txHash)
	if err != nil {
		fmt.Println("Cannot get transaction by hash : ", err)
		return nil, "", 0, nil, err
	}
	blockHash := txContent["blockHash"].(string)
	if err != nil {
		return nil, "", 0, nil, err
	}
	txIndexStr, success := txContent["transactionIndex"].(string)
	if !success {
		return nil, "", 0, nil, errors.New("Cannot find transactionIndex field")
	}
	txIndex, err := strconv.ParseUint(txIndexStr[2:], 16, 64)
	if err != nil {
		return nil, "", 0, nil, err
	}

	// Get tx's block for constructing receipt trie
	blockNumString, success := txContent["blockNumber"].(string)
	if !success {
		return nil, "", 0, nil, errors.New("Cannot find blockNumber field")
	}
	blockNumber := new(big.Int)
	_, success = blockNumber.SetString(blockNumString[2:], 16)
	if !success {
		return nil, "", 0, nil, errors.New("Cannot convert blockNumber into integer")
	}
	blockHeader, err := getETHBlockByHash(url, blockHash)
	if err != nil {
		return nil, "", 0, nil, err
	}

	// Get all sibling Txs
	siblingTxs, success := blockHeader["transactions"].([]interface{})
	if !success {
		return nil, "", 0, nil, errors.New("Cannot find transactions field")
	}

	// Constructing the receipt trie (source: go-ethereum/core/types/derive_sha.go)
	keybuf := new(bytes.Buffer)
	receiptTrie := new(trie.Trie)
	for i, tx := range siblingTxs {
		siblingReceipt, err := getETHTransactionReceipt(url, common.HexToHash(tx.(string)))
		if err != nil {
			return nil, "", 0, nil, err
		}
		keybuf.Reset()
		rlp.Encode(keybuf, uint(i))
		encodedReceipt, err := rlp.EncodeToBytes(siblingReceipt)
		if err != nil {
			return nil, "", 0, nil, err
		}
		receiptTrie.Update(keybuf.Bytes(), encodedReceipt)
	}

	// Constructing the proof for the current receipt (source: go-ethereum/trie/proof.go)
	proof := light.NewNodeSet()
	keybuf.Reset()
	rlp.Encode(keybuf, uint(txIndex))
	err = receiptTrie.Prove(keybuf.Bytes(), 0, proof)
	if err != nil {
		return nil, "", 0, nil, err
	}

	nodeList := proof.NodeList()
	encNodeList := make([]string, 0)
	for _, node := range nodeList {
		str := base64.StdEncoding.EncodeToString(node)
		encNodeList = append(encNodeList, str)
	}

	return blockNumber, blockHash, uint(txIndex), encNodeList, nil
}



// getTransactionByHashToInterface returns the transaction as a map[string]interface{} type
func getETHTransactionByHash(
	url string,
	tx common.Hash,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{tx.String()}
	var res NormalResult
	err := rpcClient.RPCCall(
		"",
		url,
		"",
		"eth_getTransactionByHash",
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}
	if res.Result == nil {
		return nil, errors.New("eth tx by hash not found")
	}
	return res.Result.(map[string]interface{}), nil
}

func getETHBlockByHash(
	url string,
	blockHash string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{blockHash, false}
	var res NormalResult
	err := rpcClient.RPCCall(
		"",
		url,
		"",
		"eth_getBlockByHash",
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}
	return res.Result.(map[string]interface{}), nil
}

func getETHTransactionReceipt(url string, txHash common.Hash) (*types.Receipt, error) {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{txHash.String()}
	var res Receipt
	err := rpcClient.RPCCall(
		"",
		url,
		"",
		"eth_getTransactionReceipt",
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

package main

import (
	"fmt"
	"math/big"
	"testing"
	"time"
	"crypto/ecdsa"
	"os/exec"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/require"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	
	"github.com/incognitochain/portal3-eth/portal/delegator"
	"github.com/incognitochain/portal3-eth/portal/incognitoproxy"
	"github.com/incognitochain/portal3-eth/portal/portalv3"

)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type PortalIntegrationTestSuite struct {
	*PortalV3BaseTestSuite

	WETHAddr                    common.Address

	IncKBNTokenIDStr  string
	IncSALTTokenIDStr string
	IncOMGTokenIDStr  string
	IncSNTTokenIDStr  string

	EtherAddressStrKyber string
	KBNAddressStr        string
	SALTAddressStr       string
	OMGAddressStr        string
	SNTAddressStr        string

	// token amounts for tests
	DepositingEther       float64
	KBNBalanceAfterStep1  *big.Int
	SALTBalanceAfterStep2 *big.Int
	auth              *bind.TransactOpts
}

func NewPortalIntegrationTestSuite(tradingTestSuite *PortalV3BaseTestSuite) *PortalIntegrationTestSuite {
	return &PortalIntegrationTestSuite{
		PortalV3BaseTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (pg *PortalIntegrationTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env

	pg.IncKBNTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000082"
	pg.IncSALTTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000081"
	pg.IncOMGTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000072"
	pg.IncSNTTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000071"
	pg.KBNAddressStr = "0xad67cB4d63C9da94AcA37fDF2761AaDF780ff4a2"                                    // kovan
	pg.SALTAddressStr = "0x6fEE5727EE4CdCBD91f3A873ef2966dF31713A04"                                   // kovan
	pg.OMGAddressStr = "0xdB7ec4E4784118D9733710e46F7C83fE7889596a"                                    // kovan
	pg.SNTAddressStr = "0x4c99B04682fbF9020Fcb31677F8D8d66832d3322"                                    // kovan
	pg.DepositingEther = float64(0.05)
	pg.ETHPrivKeyStr = "1ABA488300A9D7297A315D127837BE4219107C62C61966ECDF7A75431D75CC61"
	pg.ETHHost = "http://localhost:8545"

	var err error
	fmt.Println("Pulling image if not exist, please wait...")
	// remove container if already running
	exec.Command("/bin/sh", "-c", "docker rm -f portalv3").Output()
	exec.Command("/bin/sh", "-c", "docker rm -f incognito").Output()
	_, err = exec.Command("/bin/sh", "-c", "docker run -d -p 8545:8545 --name portalv3 bomtb/kybertrade").Output()
	require.Equal(pg.T(), nil, err)
	time.Sleep(10 * time.Second)

	ETHPrivKey, ETHClient, err := ethInstance(pg.ETHPrivKeyStr, pg.ETHHost)
	require.Equal(pg.T(), nil, err)
	pg.ETHClient = ETHClient
	pg.ETHPrivKey = ETHPrivKey
	c := getFixedCommittee()
	pg.auth = bind.NewKeyedTransactor(ETHPrivKey)

	incAddr, _, _, err := incognitoproxy.DeployIncognitoproxy(pg.auth, pg.ETHClient, pg.auth.From, c.beacons)
	require.Equal(pg.T(), nil, err)
	fmt.Printf("Proxy address: %s\n", incAddr.Hex())
	portalv3Logic, _, _, err := portalv3.DeployPortalv3(pg.auth, pg.ETHClient)
	require.Equal(pg.T(), nil, err)
	fmt.Printf("portalv3 address: %s\n", portalv3Logic.Hex())
	// PortalV3
	pg.Portalv3, _, _, err = delegator.DeployDelegator(pg.auth, pg.ETHClient, pg.auth.From, portalv3Logic, incAddr)
	require.Equal(pg.T(), nil, err)
	fmt.Printf("delegator address: %s\n", pg.Portalv3.Hex())

	// run incognito chain
	incogitoWithArgument := fmt.Sprintf("docker run -d -p 9334:9334 -p 9338:9338 --name incognito incognito %v", pg.Portalv3.Hex())
	_, err = exec.Command("/bin/sh", "-c", incogitoWithArgument).Output()
	require.Equal(pg.T(), nil, err)
	time.Sleep(15 * time.Second)
}

func (pg *PortalIntegrationTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	// _, err := exec.Command("/bin/sh", "-c", "docker rm -f portalv3").Output()
	// require.Equal(pg.T(), nil, err)
	// pg.ETHClient.Close()

	// _, err = exec.Command("/bin/sh", "-c", "docker rm -f incognito").Output()
	// require.Equal(pg.T(), nil, err)
}

func (pg *PortalIntegrationTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (pg *PortalIntegrationTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPortalIntegrationTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Kyber test suite...")

	pg := new(PortalV3BaseTestSuite)
	suite.Run(t, pg)

	portalv3Suite := NewPortalIntegrationTestSuite(pg)
	suite.Run(t, portalv3Suite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (pg *PortalIntegrationTestSuite) Test1CustodianDeposit() {
	fmt.Println("============ TEST CUSTODIAN DEPOSIT ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	// depositAmount := big.NewInt(int64(pg.DepositingEther * params.Ether))

	// _ := crypto.PubkeyToAddress(pg.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: Custodian deposit ETH --------------")
	txHash := pg.depositETH(
		pg.DepositingEther,
		pg.IncPaymentAddrStr,
	)
	// time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(pg.ETHHost, txHash)
	require.Equal(pg.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 15s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)
	_, err = pg.callCustodianDeposit(
		pg.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(pg.T(), nil, err)
	time.Sleep(120 * time.Second)
}

func ethInstance(ethPrivate string, ethEnpoint string) (*ecdsa.PrivateKey, *ethclient.Client, error) {
	privKey, err := crypto.HexToECDSA(ethPrivate)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "development"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial(ethEnpoint)
	if err != nil {
		return nil, nil, err
	}
	return privKey, client, nil
}
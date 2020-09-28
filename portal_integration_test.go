package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/portal3-eth/portal/delegator"
	"github.com/incognitochain/portal3-eth/portal/incognitoproxy"
	"github.com/incognitochain/portal3-eth/portal/portalv3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/big"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type PortalIntegrationTestSuite struct {
	*PortalV3BaseTestSuite

	WETHAddr common.Address

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
	auth                  *bind.TransactOpts
}

func NewPortalIntegrationTestSuite(tradingTestSuite *PortalV3BaseTestSuite) *PortalIntegrationTestSuite {
	return &PortalIntegrationTestSuite{
		PortalV3BaseTestSuite: tradingTestSuite,
	}
}

func (pg *PortalIntegrationTestSuite) genBlock() {
	for i := 0; i < 16; i++ {
		nonce, err := pg.ETHClient.PendingNonceAt(context.Background(), pg.auth.From)
		require.Equal(pg.T(), nil, err)
		transfer(pg.ETHClient, pg.ETHPrivKey, pg.auth.From.Hex(), nonce, big.NewInt(0), uint64(21000), big.NewInt(5000000000))
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
	pg.KBNAddressStr = "0xad67cB4d63C9da94AcA37fDF2761AaDF780ff4a2"  // kovan
	pg.SALTAddressStr = "0x6fEE5727EE4CdCBD91f3A873ef2966dF31713A04" // kovan
	pg.OMGAddressStr = "0xdB7ec4E4784118D9733710e46F7C83fE7889596a"  // kovan
	pg.SNTAddressStr = "0x4c99B04682fbF9020Fcb31677F8D8d66832d3322"  // kovan
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
	pg.auth = bind.NewKeyedTransactor(ETHPrivKey)
	c := getFixedCommittee()

	incAddr, _, _, err := incognitoproxy.DeployIncognitoproxy(pg.auth, pg.ETHClient, pg.auth.From, c.beacons)
	require.Equal(pg.T(), nil, err)
	fmt.Printf("Proxy address: %s\n", incAddr.Hex())
	portalv3Logic, _, _, err := portalv3.DeployPortalv3(pg.auth, pg.ETHClient)
	require.Equal(pg.T(), nil, err)
	fmt.Printf("portalv3 address: %s\n", portalv3Logic.Hex())
	//PortalV3
	pg.Portalv3, _, _, err = delegator.DeployDelegator(pg.auth, pg.ETHClient, pg.auth.From, portalv3Logic, incAddr)
	require.Equal(pg.T(), nil, err)
	fmt.Printf("delegator address: %s\n", pg.Portalv3.Hex())
	//pg.Portalv3 = common.HexToAddress("0x1B4c8873Fd83aB7E2eaB66cDB238F884827a61d4")

	//get portalv3 ip
	ipAddress, err := exec.Command("/bin/sh", "-c", "docker inspect -f \"{{ .NetworkSettings.IPAddress }}\" portalv3").Output()
	require.Equal(pg.T(), nil, err)

	// run incognito chaind
	incogitoWithArgument := fmt.Sprintf("docker run -d -p 9334:9334 -p 9338:9338 --name incognito -e GETH_NAME=%v -e PORTAL_CONTRACT=%v incognito", string(ipAddress), pg.Portalv3.Hex())
	incogitoWithArgument = strings.Replace(incogitoWithArgument, "\n", "", -1)
	_, err = exec.Command("/bin/sh", "-c", incogitoWithArgument).Output()
	require.Equal(pg.T(), nil, err)

	for {
		time.Sleep(15 * time.Second)
		if checkRepsonse(pg.IncBridgeHost) {
			break
		}
	}
	time.Sleep(60 * time.Second)
}

func (pg *PortalIntegrationTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	_, err := exec.Command("/bin/sh", "-c", "docker rm -f portalv3").Output()
	require.Equal(pg.T(), nil, err)
	pg.ETHClient.Close()

	_, err = exec.Command("/bin/sh", "-c", "docker rm -f incognito").Output()
	require.Equal(pg.T(), nil, err)
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

	fmt.Println("Generate blocks to pass 15 confirmations ")
	pg.genBlock()

	// Submit proof first time must pass
	depositRes, err := pg.callCustodianDeposit(
		pg.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(pg.T(), nil, err)
	time.Sleep(30 * time.Second)
	require.NotEqual(pg.T(), nil, depositRes)
	TxId := depositRes["TxID"]
	TxDepositStatus, err := getPortalCustodianDepositStatusv3(pg.IncRPCHost, TxId.(string))
	require.Equal(pg.T(), float64(1), TxDepositStatus["Status"].(float64))
	require.Equal(pg.T(), pg.DepositingEther*1e9, TxDepositStatus["DepositAmount"].(float64))
	require.Equal(pg.T(), pg.EtherAddressStr, TxDepositStatus["ExternalTokenID"].(string))

	// Submit the same as above proof must failed
	depositRes, err = pg.callCustodianDeposit(
		pg.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(pg.T(), nil, err)
	time.Sleep(30 * time.Second)
	require.NotEqual(pg.T(), nil, depositRes)
	TxId = depositRes["TxID"]
	TxDepositStatus, err = getPortalCustodianDepositStatusv3(pg.IncRPCHost, TxId.(string))
	require.Equal(pg.T(), float64(2), TxDepositStatus["Status"].(float64))

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

func checkRepsonse(url string) bool {
	resp, err := http.Get(url)
	if err != nil || resp == nil {
		fmt.Println("Incognito chain is running please wait...")
		return false
	}
	return true
}

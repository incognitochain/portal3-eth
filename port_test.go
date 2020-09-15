package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ec "github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// // Define the suite, and absorb the built-in basic suite
// // functionality from testify - including assertion methods.
type PortalV3TestSuite struct {
	suite.Suite
	portalV3Base
	withdrawer   ec.Address
	auth         *bind.TransactOpts
	EtherAddress ec.Address
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (v3 *PortalV3TestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	v3.withdrawer = ec.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
	v3.EtherAddress = ec.HexToAddress("0x0000000000000000000000000000000000000000")
}

func (v3 *PortalV3TestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
}

func (v3 *PortalV3TestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
	p, c, err := setupFixedCommittee()
	require.Equal(v3.T(), nil, err)
	v3.p = p
	v3.c = c
}

func (v3 *PortalV3TestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPortalV3(t *testing.T) {
	fmt.Println("Starting entry point for vault v3 test suite...")
	suite.Run(t, new(PortalV3TestSuite))

	fmt.Println("Finishing entry point for vault v3 test suite...")
}

func (v3 *PortalV3TestSuite) TestPortalV3LockCustodianTokens() {
	fmt.Println("==== PORTAL 3 TEST CUSTODIAN DEPOSIT ====")

	// Deposit ETH
	oldBalance, newBalance, err := deposit(v3.p, big.NewInt(int64(5e18)))

	require.Equal(v3.T(), nil, err)
	require.Equal(v3.T(), oldBalance.Add(oldBalance, big.NewInt(int64(5e18))), newBalance)

	// Deposit ERC20
	b2e27, _ := big.NewInt(1).SetString("2000000000000000000000000000", 10)
	testCases := []struct {
		desc    string
		decimal int
		value   *big.Int
		emit    *big.Int
		err     bool
	}{
		{
			desc:    "DAI (d=18)",
			decimal: 18,
			value:   big.NewInt(int64(5e18)),
			emit:    big.NewInt(int64(5e9)),
		},
		{
			desc:    "ZIL (d=12)",
			decimal: 12,
			value:   big.NewInt(int64(3e12)),
			emit:    big.NewInt(int64(3e9)),
		},
		{
			desc:    "ABC (d=27)",
			decimal: 27,
			value:   b2e27,
			emit:    big.NewInt(int64(2e9)),
		},
		{
			desc:    "XYZ (d=9)",
			decimal: 9,
			value:   big.NewInt(int64(4e9)),
			emit:    big.NewInt(int64(4e9)),
		},
		{
			desc:    "USDT (d=6)",
			decimal: 6,
			value:   big.NewInt(int64(8e6)),
			emit:    big.NewInt(int64(8e6)),
		},
		{
			desc:    "IJK (d=0)",
			decimal: 0,
			value:   big.NewInt(9),
			emit:    big.NewInt(9),
		},
	}

	for _, tc := range testCases {
		v3.T().Run(tc.desc, func(t *testing.T) {
			decimals := []int{tc.decimal}
			err := deployERC20(decimals, v3.p)
			assert.Nil(t, err)

			tinfo := v3.p.tokens[tc.decimal]
			oldBalance := getBalanceERC20(tinfo.c, v3.p.delegatorAddr)
			_, tx, err := lockSimERC20WithTxs(v3.p, tinfo.c, tinfo.addr, tc.value)
			if assert.Nil(t, err) {
				newBalance := getBalanceERC20(tinfo.c, v3.p.delegatorAddr)
				require.Equal(t, oldBalance.Add(oldBalance, tc.value), newBalance)

				emitted, err := extractAmountInDepositERC20Event(v3.p.sim, tx)
				if assert.Nil(t, err) {
					require.Equal(t, tc.emit, emitted)
				}
			}
		})
	}
}

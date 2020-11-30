package main

import (
	"fmt"
	"github.com/incognitochain/portal3-eth/portal/delegator"
	"github.com/incognitochain/portal3-eth/portal/incognitoproxy"
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
func TestUnitPortalV3(t *testing.T) {
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

func (v3 *PortalV3TestSuite) TestPortalV3LockCustodianTokensWhenPortalPaused() {
	fmt.Println("==== PORTAL 3 TEST CUSTODIAN DEPOSIT WHILE PORTAL PAUSED ====")
	portalStorage, err := delegator.NewDelegator(v3.p.delegatorAddr, v3.p.sim)
	require.Equal(v3.T(), nil, err)
	_, err = portalStorage.Pause(auth2)
	require.Equal(v3.T(), nil, err)
	v3.p.sim.Commit()

	// Deposit ETH
	_, _, err = deposit(v3.p, big.NewInt(int64(5e18)))
	require.NotEqual(v3.T(), nil, err)

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
			_, _, err = lockSimERC20WithTxs(v3.p, tinfo.c, tinfo.addr, tc.value)
			require.NotEqual(v3.T(), nil, err)
		})
	}
}

func (v3 *PortalV3TestSuite) TestPortalV3UnLockCustodianTokens() {
	fmt.Println("==== PORTAL 3 TEST CUSTODIAN UNLOCK TOKEN FROM PORTAL ====")
	b2e27, _ := big.NewInt(1).SetString("2000000000000000000000000000", 10)
	//isPaused, _ := v3.p.portalV3Ins.Paused(nil)
	//fmt.Println(isPaused)

	testCases := []struct {
		desc     string
		decimal  int
		deposit  *big.Int
		withdraw *big.Int
		remain   *big.Int
		err      bool
	}{
		{
			desc:     "DAI (d=18)",
			decimal:  18,
			deposit:  big.NewInt(int64(5e18)),
			withdraw: big.NewInt(int64(4e9)),
			remain:   big.NewInt(int64(1e18)),
		},
		{
			desc:     "ZIL (d=12)",
			decimal:  12,
			deposit:  big.NewInt(int64(3e12)),
			withdraw: big.NewInt(int64(3e8)),
			remain:   big.NewInt(int64(2.7e12)),
		},
		{
			desc:     "ABC (d=27)",
			decimal:  27,
			deposit:  b2e27,
			withdraw: big.NewInt(int64(2e9)),
			remain:   big.NewInt(int64(0)),
		},
		{
			desc:     "XYZ (d=9)",
			decimal:  9,
			deposit:  big.NewInt(int64(4e9)),
			withdraw: big.NewInt(int64(1)),
			remain:   big.NewInt(int64(4e9 - 1)),
		},
		{
			desc:     "USDT (d=6)",
			decimal:  6,
			deposit:  big.NewInt(int64(8e6)),
			withdraw: big.NewInt(int64(7e6)),
			remain:   big.NewInt(int64(1e6)),
		},
		{
			desc:     "IJK (d=0)",
			decimal:  0,
			deposit:  big.NewInt(9),
			withdraw: big.NewInt(2),
			remain:   big.NewInt(7),
		},
	}

	for _, tc := range testCases {
		v3.T().Run(tc.desc, func(t *testing.T) {
			decimals := []int{tc.decimal}
			err := deployERC20(decimals, v3.p)
			assert.Nil(t, err)
			tinfo := v3.p.tokens[tc.decimal]

			// Deposit, must success
			_, _, err = lockSimERC20WithTxs(v3.p, tinfo.c, tinfo.addr, tc.deposit)
			assert.Nil(t, err)
			v3.p.sim.Commit()
			meta := 170
			shardID := 0
			proof, instHash := buildWithdrawTestcase(v3.c, meta, shardID, []ec.Address{tinfo.addr}, []*big.Int{tc.withdraw}, auth.From)
			incognitoProxy, _ := incognitoproxy.NewIncognitoproxy(v3.p.incAddr, v3.p.sim)
			isApproved, err := incognitoProxy.InstructionApproved(
				nil,
				true,
				instHash,
				proof.Heights,
				proof.InstPaths,
				proof.InstPathIsLefts,
				proof.InstRoots,
				proof.BlkData,
				proof.SigIdxs,
				proof.SigVs,
				proof.SigRs,
				proof.SigSs,
			)
			require.Equal(v3.T(), nil, err)
			require.Equal(v3.T(), true, isApproved)

			_, err = Withdraw(v3.p.portalV3Ins, auth, proof)
			if assert.Nil(t, err) {
				v3.p.sim.Commit()

				// Check balance
				bal := getBalanceERC20(tinfo.c, v3.p.delegatorAddr)
				assert.Zero(t, tc.remain.Cmp(bal))
			}
		})
	}
}

func (v3 *PortalV3TestSuite) TestPortalV3UnLockCustodianTokensWhilePaused() {
	fmt.Println("==== PORTAL 3 TEST CUSTODIAN UNLOCK TOKEN FROM PORTAL WHILE PORTAL PAUSED ====")
	b2e27, _ := big.NewInt(1).SetString("2000000000000000000000000000", 10)

	testCases := []struct {
		desc     string
		decimal  int
		deposit  *big.Int
		withdraw *big.Int
		remain   *big.Int
		err      bool
	}{
		{
			desc:     "DAI (d=18)",
			decimal:  18,
			deposit:  big.NewInt(int64(5e18)),
			withdraw: big.NewInt(int64(4e9)),
			remain:   big.NewInt(int64(1e18)),
		},
		{
			desc:     "ZIL (d=12)",
			decimal:  12,
			deposit:  big.NewInt(int64(3e12)),
			withdraw: big.NewInt(int64(3e8)),
			remain:   big.NewInt(int64(2.7e12)),
		},
		{
			desc:     "ABC (d=27)",
			decimal:  27,
			deposit:  b2e27,
			withdraw: big.NewInt(int64(2e9)),
			remain:   big.NewInt(int64(0)),
		},
		{
			desc:     "XYZ (d=9)",
			decimal:  9,
			deposit:  big.NewInt(int64(4e9)),
			withdraw: big.NewInt(int64(1)),
			remain:   big.NewInt(int64(4e9 - 1)),
		},
		{
			desc:     "USDT (d=6)",
			decimal:  6,
			deposit:  big.NewInt(int64(8e6)),
			withdraw: big.NewInt(int64(7e6)),
			remain:   big.NewInt(int64(1e6)),
		},
		{
			desc:     "IJK (d=0)",
			decimal:  0,
			deposit:  big.NewInt(9),
			withdraw: big.NewInt(2),
			remain:   big.NewInt(7),
		},
	}

	for _, tc := range testCases {
		v3.T().Run(tc.desc, func(t *testing.T) {
			decimals := []int{tc.decimal}
			err := deployERC20(decimals, v3.p)
			assert.Nil(t, err)
			tinfo := v3.p.tokens[tc.decimal]

			// Deposit, must success
			_, _, err = lockSimERC20WithTxs(v3.p, tinfo.c, tinfo.addr, tc.deposit)
			assert.Nil(t, err)
		})
	}

	portalStorage, err := delegator.NewDelegator(v3.p.delegatorAddr, v3.p.sim)
	require.Equal(v3.T(), nil, err)
	_, err = portalStorage.Pause(auth2)
	v3.p.sim.Commit()

	for _, tc := range testCases {
		v3.T().Run(tc.desc, func(t *testing.T) {
			tinfo := v3.p.tokens[tc.decimal]
			meta := 170
			shardID := 0
			proof, _ := buildWithdrawTestcase(v3.c, meta, shardID, []ec.Address{tinfo.addr}, []*big.Int{tc.withdraw}, auth.From)

			_, err = Withdraw(v3.p.portalV3Ins, auth, proof)
			// Withdraw, must fail
			require.NotEqual(v3.T(), nil, err)
		})
	}
}

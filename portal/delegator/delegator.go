// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package delegator

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DelegatorABI is the input ABI used to generate the binding from.
const DelegatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_incognito\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousIncognito\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncognito\",\"type\":\"address\"}],\"name\":\"IncognitoChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousSuccessor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSuccessor\",\"type\":\"address\"}],\"name\":\"SuccessorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSuccessor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newIncognito\",\"type\":\"address\"}],\"name\":\"upgradeIncognito\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// DelegatorBin is the compiled bytecode used for deploying new contracts.
var DelegatorBin = "0x60806040526040516200164f3803806200164f833981810160405260808110156200002957600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805160405193929190846401000000008211156200006857600080fd5b838201915060208201858111156200007f57600080fd5b82518660018202830111640100000000821117156200009d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000d3578082015181840152602081019050620000b6565b50505050905090810190601f168015620001015780820380516001836020036101000a031916815260200191505b506040525050508381600160405180807f656970313936372e70726f78792e696d706c656d656e746174696f6e00000000815250601c019050604051809103902060001c0360001b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b146200017557fe5b62000186826200043260201b60201c565b600081511115620002585760008273ffffffffffffffffffffffffffffffffffffffff16826040518082805190602001908083835b60208310620001e05780518252602082019150602081019050602083039250620001bb565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811462000242576040519150601f19603f3d011682016040523d82523d6000602084013e62000247565b606091505b50509050806200025657600080fd5b505b5050600160405180807f656970313936372e70726f78792e61646d696e000000000000000000000000008152506013019050604051809103902060001c0360001b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b14620002c557fe5b600160405180807f656970313936372e70726f78792e737563636573736f720000000000000000008152506017019050604051809103902060001c0360001b7f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f860001b146200033057fe5b600160405180807f656970313936372e70726f78792e7061757365640000000000000000000000008152506014019050604051809103902060001c0360001b7f8dea8703c3cf94703383ce38a9c894669dccd4ca8e65ddb43267aa024871145060001b146200039b57fe5b600160405180807f656970313936372e70726f78792e696e636f676e69746f2e00000000000000008152506018019050604051809103902060001c0360001b7f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd260001b146200040657fe5b6200041783620004c960201b60201c565b6200042882620004f860201b60201c565b505050506200053a565b62000443816200052760201b60201c565b6200049a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180620016196036913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b60007f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd260001b90508181555050565b600080823b905060008111915050919050565b6110cf806200054a6000396000f3fe6080604052600436106100955760003560e01c80635c60da1b116100595780635c60da1b146102175780636ff968c31461026e5780638456cb59146102c55780639e6371ba146102dc578063f851a4401461032d576100a4565b80631c587771146100ae5780633659cfe6146100ff5780633f4ba83a146101505780634e71d92d146101675780634f1ef2861461017e576100a4565b366100a4576100a2610384565b005b6100ac610384565b005b3480156100ba57600080fd5b506100fd600480360360208110156100d157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061039e565b005b34801561010b57600080fd5b5061014e6004803603602081101561012257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610517565b005b34801561015c57600080fd5b5061016561056c565b005b34801561017357600080fd5b5061017c61061e565b005b6102156004803603604081101561019457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156101d157600080fd5b8201836020820111156101e357600080fd5b8035906020019184600183028401116401000000008311171561020557600080fd5b9091929391929390505050610725565b005b34801561022357600080fd5b5061022c6107fb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561027a57600080fd5b50610283610853565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102d157600080fd5b506102da6108ab565b005b3480156102e857600080fd5b5061032b600480360360208110156102ff57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061095e565b005b34801561033957600080fd5b50610342610ad7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61038c610b2f565b61039c610397610c23565b610c54565b565b6103a6610c7a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561050b57600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561045f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526040815260200180610f856040913960400191505060405180910390fd5b7f86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c610488610cab565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161050681610cdc565b610514565b610513610384565b5b50565b61051f610c7a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156105605761055b81610d0b565b610569565b610568610384565b5b50565b610574610c7a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610613576105af610d5a565b610604576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806110286030913960400191505060405180910390fd5b61060e6000610d8b565b61061c565b61061b610384565b5b565b610626610dba565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106a9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180610eec6029913960400191505060405180910390fd5b7f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc6106d2610dba565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a161072361071e610dba565b610deb565b565b61072d610c7a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156107ed5761076983610d0b565b60008373ffffffffffffffffffffffffffffffffffffffff168383604051808383808284378083019250505092505050600060405180830381855af49150503d80600081146107d4576040519150601f19603f3d011682016040523d82523d6000602084013e6107d9565b606091505b50509050806107e757600080fd5b506107f6565b6107f5610384565b5b505050565b6000610805610c7a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561084757610840610c23565b9050610850565b61084f610384565b5b90565b600061085d610c7a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561089f57610898610dba565b90506108a8565b6108a7610384565b5b90565b6108b3610c7a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610953576108ee610d5a565b15610944576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526034815260200180610fc56034913960400191505060405180910390fd5b61094e6001610d8b565b61095c565b61095b610384565b5b565b610966610c7a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610acb57600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a1f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180610f15603a913960400191505060405180910390fd5b7ff966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974610a48610dba565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1610ac681610e1a565b610ad4565b610ad3610384565b5b50565b6000610ae1610c7a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610b2357610b1c610c7a565b9050610b2c565b610b2b610384565b5b90565b610b37610c7a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610bbb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260428152602001806110586042913960600191505060405180910390fd5b610bc3610d5a565b15610c19576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180610ff9602f913960400191505060405180910390fd5b610c21610e49565b565b6000807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b9050805491505090565b3660008037600080366000845af43d6000803e8060008114610c75573d6000f35b3d6000fd5b6000807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b9050805491505090565b6000807f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd260001b9050805491505090565b60007f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd260001b90508181555050565b610d1481610e4b565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b6000807f8dea8703c3cf94703383ce38a9c894669dccd4ca8e65ddb43267aa024871145060001b9050805491505090565b60007f8dea8703c3cf94703383ce38a9c894669dccd4ca8e65ddb43267aa024871145060001b90508181555050565b6000807f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f860001b9050805491505090565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b60007f7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f860001b90508181555050565b565b610e5481610ed8565b610ea9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180610f4f6036913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b600080823b90506000811191505091905056fe5472616e73706172656e745570677261646561626c6550726f78793a20756e617574686f72697a65645472616e73706172656e745570677261646561626c6550726f78793a20737563636573736f7220697320746865207a65726f20616464726573735570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a20696e636f676e69746f2070726f787920697320746865207a65726f20616464726573735472616e73706172656e745570677261646561626c6550726f78793a20636f6e74726163742070617573656420616c72656164795472616e73706172656e745570677261646561626c6550726f78793a20636f6e7472616374206973207061757365645472616e73706172656e745570677261646561626c6550726f78793a20636f6e7472616374206e6f74207061757365645472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a2646970667358221220ee7fec885a2bc1364166200039c9913c39d9f074fea65789cd01b8a37097ebbc64736f6c634300060600335570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374"

// DeployDelegator deploys a new Ethereum contract, binding an instance of Delegator to it.
func DeployDelegator(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _admin common.Address, _incognito common.Address, _data []byte) (common.Address, *types.Transaction, *Delegator, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DelegatorBin), backend, _logic, _admin, _incognito, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Delegator{DelegatorCaller: DelegatorCaller{contract: contract}, DelegatorTransactor: DelegatorTransactor{contract: contract}, DelegatorFilterer: DelegatorFilterer{contract: contract}}, nil
}

// Delegator is an auto generated Go binding around an Ethereum contract.
type Delegator struct {
	DelegatorCaller     // Read-only binding to the contract
	DelegatorTransactor // Write-only binding to the contract
	DelegatorFilterer   // Log filterer for contract events
}

// DelegatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegatorSession struct {
	Contract     *Delegator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegatorCallerSession struct {
	Contract *DelegatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DelegatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegatorTransactorSession struct {
	Contract     *DelegatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DelegatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegatorRaw struct {
	Contract *Delegator // Generic contract binding to access the raw methods on
}

// DelegatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegatorCallerRaw struct {
	Contract *DelegatorCaller // Generic read-only contract binding to access the raw methods on
}

// DelegatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegatorTransactorRaw struct {
	Contract *DelegatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegator creates a new instance of Delegator, bound to a specific deployed contract.
func NewDelegator(address common.Address, backend bind.ContractBackend) (*Delegator, error) {
	contract, err := bindDelegator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Delegator{DelegatorCaller: DelegatorCaller{contract: contract}, DelegatorTransactor: DelegatorTransactor{contract: contract}, DelegatorFilterer: DelegatorFilterer{contract: contract}}, nil
}

// NewDelegatorCaller creates a new read-only instance of Delegator, bound to a specific deployed contract.
func NewDelegatorCaller(address common.Address, caller bind.ContractCaller) (*DelegatorCaller, error) {
	contract, err := bindDelegator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatorCaller{contract: contract}, nil
}

// NewDelegatorTransactor creates a new write-only instance of Delegator, bound to a specific deployed contract.
func NewDelegatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegatorTransactor, error) {
	contract, err := bindDelegator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatorTransactor{contract: contract}, nil
}

// NewDelegatorFilterer creates a new log filterer instance of Delegator, bound to a specific deployed contract.
func NewDelegatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegatorFilterer, error) {
	contract, err := bindDelegator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegatorFilterer{contract: contract}, nil
}

// bindDelegator binds a generic wrapper to an already deployed contract.
func bindDelegator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegator *DelegatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Delegator.Contract.DelegatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegator *DelegatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.Contract.DelegatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegator *DelegatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegator.Contract.DelegatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegator *DelegatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Delegator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegator *DelegatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegator *DelegatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegator.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Delegator *DelegatorTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Delegator *DelegatorSession) Admin() (*types.Transaction, error) {
	return _Delegator.Contract.Admin(&_Delegator.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Delegator *DelegatorTransactorSession) Admin() (*types.Transaction, error) {
	return _Delegator.Contract.Admin(&_Delegator.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Delegator *DelegatorTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Delegator *DelegatorSession) Claim() (*types.Transaction, error) {
	return _Delegator.Contract.Claim(&_Delegator.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Delegator *DelegatorTransactorSession) Claim() (*types.Transaction, error) {
	return _Delegator.Contract.Claim(&_Delegator.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_Delegator *DelegatorTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_Delegator *DelegatorSession) Implementation() (*types.Transaction, error) {
	return _Delegator.Contract.Implementation(&_Delegator.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_Delegator *DelegatorTransactorSession) Implementation() (*types.Transaction, error) {
	return _Delegator.Contract.Implementation(&_Delegator.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Delegator *DelegatorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Delegator *DelegatorSession) Pause() (*types.Transaction, error) {
	return _Delegator.Contract.Pause(&_Delegator.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Delegator *DelegatorTransactorSession) Pause() (*types.Transaction, error) {
	return _Delegator.Contract.Pause(&_Delegator.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address newSuccessor) returns()
func (_Delegator *DelegatorTransactor) Retire(opts *bind.TransactOpts, newSuccessor common.Address) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "retire", newSuccessor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address newSuccessor) returns()
func (_Delegator *DelegatorSession) Retire(newSuccessor common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.Retire(&_Delegator.TransactOpts, newSuccessor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address newSuccessor) returns()
func (_Delegator *DelegatorTransactorSession) Retire(newSuccessor common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.Retire(&_Delegator.TransactOpts, newSuccessor)
}

// Successor is a paid mutator transaction binding the contract method 0x6ff968c3.
//
// Solidity: function successor() returns(address)
func (_Delegator *DelegatorTransactor) Successor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "successor")
}

// Successor is a paid mutator transaction binding the contract method 0x6ff968c3.
//
// Solidity: function successor() returns(address)
func (_Delegator *DelegatorSession) Successor() (*types.Transaction, error) {
	return _Delegator.Contract.Successor(&_Delegator.TransactOpts)
}

// Successor is a paid mutator transaction binding the contract method 0x6ff968c3.
//
// Solidity: function successor() returns(address)
func (_Delegator *DelegatorTransactorSession) Successor() (*types.Transaction, error) {
	return _Delegator.Contract.Successor(&_Delegator.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Delegator *DelegatorTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Delegator *DelegatorSession) Unpause() (*types.Transaction, error) {
	return _Delegator.Contract.Unpause(&_Delegator.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Delegator *DelegatorTransactorSession) Unpause() (*types.Transaction, error) {
	return _Delegator.Contract.Unpause(&_Delegator.TransactOpts)
}

// UpgradeIncognito is a paid mutator transaction binding the contract method 0x1c587771.
//
// Solidity: function upgradeIncognito(address newIncognito) returns()
func (_Delegator *DelegatorTransactor) UpgradeIncognito(opts *bind.TransactOpts, newIncognito common.Address) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "upgradeIncognito", newIncognito)
}

// UpgradeIncognito is a paid mutator transaction binding the contract method 0x1c587771.
//
// Solidity: function upgradeIncognito(address newIncognito) returns()
func (_Delegator *DelegatorSession) UpgradeIncognito(newIncognito common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.UpgradeIncognito(&_Delegator.TransactOpts, newIncognito)
}

// UpgradeIncognito is a paid mutator transaction binding the contract method 0x1c587771.
//
// Solidity: function upgradeIncognito(address newIncognito) returns()
func (_Delegator *DelegatorTransactorSession) UpgradeIncognito(newIncognito common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.UpgradeIncognito(&_Delegator.TransactOpts, newIncognito)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Delegator *DelegatorTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Delegator *DelegatorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.UpgradeTo(&_Delegator.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Delegator *DelegatorTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.UpgradeTo(&_Delegator.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Delegator *DelegatorTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Delegator *DelegatorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Delegator.Contract.UpgradeToAndCall(&_Delegator.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Delegator *DelegatorTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Delegator.Contract.UpgradeToAndCall(&_Delegator.TransactOpts, newImplementation, data)
}

// DelegatorAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Delegator contract.
type DelegatorAdminChangedIterator struct {
	Event *DelegatorAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegatorAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegatorAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegatorAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorAdminChanged represents a AdminChanged event raised by the Delegator contract.
type DelegatorAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Delegator *DelegatorFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*DelegatorAdminChangedIterator, error) {

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &DelegatorAdminChangedIterator{contract: _Delegator.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Delegator *DelegatorFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *DelegatorAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorAdminChanged)
				if err := _Delegator.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Delegator *DelegatorFilterer) ParseAdminChanged(log types.Log) (*DelegatorAdminChanged, error) {
	event := new(DelegatorAdminChanged)
	if err := _Delegator.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DelegatorClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Delegator contract.
type DelegatorClaimIterator struct {
	Event *DelegatorClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegatorClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegatorClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegatorClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorClaim represents a Claim event raised by the Delegator contract.
type DelegatorClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Delegator *DelegatorFilterer) FilterClaim(opts *bind.FilterOpts) (*DelegatorClaimIterator, error) {

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &DelegatorClaimIterator{contract: _Delegator.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Delegator *DelegatorFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *DelegatorClaim) (event.Subscription, error) {

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorClaim)
				if err := _Delegator.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Delegator *DelegatorFilterer) ParseClaim(log types.Log) (*DelegatorClaim, error) {
	event := new(DelegatorClaim)
	if err := _Delegator.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DelegatorIncognitoChangedIterator is returned from FilterIncognitoChanged and is used to iterate over the raw logs and unpacked data for IncognitoChanged events raised by the Delegator contract.
type DelegatorIncognitoChangedIterator struct {
	Event *DelegatorIncognitoChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegatorIncognitoChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorIncognitoChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegatorIncognitoChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegatorIncognitoChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorIncognitoChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorIncognitoChanged represents a IncognitoChanged event raised by the Delegator contract.
type DelegatorIncognitoChanged struct {
	PreviousIncognito common.Address
	NewIncognito      common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterIncognitoChanged is a free log retrieval operation binding the contract event 0x86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c.
//
// Solidity: event IncognitoChanged(address previousIncognito, address newIncognito)
func (_Delegator *DelegatorFilterer) FilterIncognitoChanged(opts *bind.FilterOpts) (*DelegatorIncognitoChangedIterator, error) {

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "IncognitoChanged")
	if err != nil {
		return nil, err
	}
	return &DelegatorIncognitoChangedIterator{contract: _Delegator.contract, event: "IncognitoChanged", logs: logs, sub: sub}, nil
}

// WatchIncognitoChanged is a free log subscription operation binding the contract event 0x86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c.
//
// Solidity: event IncognitoChanged(address previousIncognito, address newIncognito)
func (_Delegator *DelegatorFilterer) WatchIncognitoChanged(opts *bind.WatchOpts, sink chan<- *DelegatorIncognitoChanged) (event.Subscription, error) {

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "IncognitoChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorIncognitoChanged)
				if err := _Delegator.contract.UnpackLog(event, "IncognitoChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncognitoChanged is a log parse operation binding the contract event 0x86d392a76e88298144124db3dd7265135d76810f52d747dc329a0f7722135e5c.
//
// Solidity: event IncognitoChanged(address previousIncognito, address newIncognito)
func (_Delegator *DelegatorFilterer) ParseIncognitoChanged(log types.Log) (*DelegatorIncognitoChanged, error) {
	event := new(DelegatorIncognitoChanged)
	if err := _Delegator.contract.UnpackLog(event, "IncognitoChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DelegatorSuccessorChangedIterator is returned from FilterSuccessorChanged and is used to iterate over the raw logs and unpacked data for SuccessorChanged events raised by the Delegator contract.
type DelegatorSuccessorChangedIterator struct {
	Event *DelegatorSuccessorChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegatorSuccessorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorSuccessorChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegatorSuccessorChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegatorSuccessorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorSuccessorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorSuccessorChanged represents a SuccessorChanged event raised by the Delegator contract.
type DelegatorSuccessorChanged struct {
	PreviousSuccessor common.Address
	NewSuccessor      common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSuccessorChanged is a free log retrieval operation binding the contract event 0xf966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974.
//
// Solidity: event SuccessorChanged(address previousSuccessor, address newSuccessor)
func (_Delegator *DelegatorFilterer) FilterSuccessorChanged(opts *bind.FilterOpts) (*DelegatorSuccessorChangedIterator, error) {

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "SuccessorChanged")
	if err != nil {
		return nil, err
	}
	return &DelegatorSuccessorChangedIterator{contract: _Delegator.contract, event: "SuccessorChanged", logs: logs, sub: sub}, nil
}

// WatchSuccessorChanged is a free log subscription operation binding the contract event 0xf966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974.
//
// Solidity: event SuccessorChanged(address previousSuccessor, address newSuccessor)
func (_Delegator *DelegatorFilterer) WatchSuccessorChanged(opts *bind.WatchOpts, sink chan<- *DelegatorSuccessorChanged) (event.Subscription, error) {

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "SuccessorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorSuccessorChanged)
				if err := _Delegator.contract.UnpackLog(event, "SuccessorChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSuccessorChanged is a log parse operation binding the contract event 0xf966f857c3c376f2e1df873bbe2596a18675dc056dc3465dfbbe8fe9ac02c974.
//
// Solidity: event SuccessorChanged(address previousSuccessor, address newSuccessor)
func (_Delegator *DelegatorFilterer) ParseSuccessorChanged(log types.Log) (*DelegatorSuccessorChanged, error) {
	event := new(DelegatorSuccessorChanged)
	if err := _Delegator.contract.UnpackLog(event, "SuccessorChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DelegatorUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Delegator contract.
type DelegatorUpgradedIterator struct {
	Event *DelegatorUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegatorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegatorUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegatorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorUpgraded represents a Upgraded event raised by the Delegator contract.
type DelegatorUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Delegator *DelegatorFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*DelegatorUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorUpgradedIterator{contract: _Delegator.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Delegator *DelegatorFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *DelegatorUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorUpgraded)
				if err := _Delegator.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Delegator *DelegatorFilterer) ParseUpgraded(log types.Log) (*DelegatorUpgraded, error) {
	event := new(DelegatorUpgraded)
	if err := _Delegator.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package incognitoproxy

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

// IncognitoProxyCommittee is an auto generated low-level Go binding around an user-defined struct.
type IncognitoProxyCommittee struct {
	Pubkeys    []common.Address
	StartBlock *big.Int
}

// IncognitoproxyABI is the input ABI used to generate the binding from.
const IncognitoproxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"}],\"name\":\"BeaconCommitteeSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numVals\",\"type\":\"uint256\"}],\"name\":\"extractCommitteeFromInstruction\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractMetaFromInstruction\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"}],\"name\":\"findBeaconCommitteeFromHeight\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getBeaconCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"pubkeys\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIncognitoProxy.Committee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"}],\"name\":\"instructionInMerkleTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"swapBeaconCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"verifySig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// IncognitoproxyBin is the compiled bytecode used for deploying new contracts.
var IncognitoproxyBin = "0x60806040523480156200001157600080fd5b5060405162002a3438038062002a3483398181016040528101906200003791906200028c565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160146101000a81548160ff0219169083151502179055506301e13380420160028190555060036040518060400160405280838152602001600081525090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000019080519060200190620000fe92919062000113565b5060208201518160010155505050506200038b565b8280548282559060005260206000209081019282156200018f579160200282015b828111156200018e5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000134565b5b5090506200019e9190620001a2565b5090565b620001e591905b80821115620001e157600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101620001a9565b5090565b90565b600081519050620001f98162000371565b92915050565b600082601f8301126200021157600080fd5b815162000228620002228262000314565b620002e6565b915081818352602084019350602081019050838560208402820111156200024e57600080fd5b60005b83811015620002825781620002678882620001e8565b84526020840193506020830192505060018101905062000251565b5050505092915050565b60008060408385031215620002a057600080fd5b6000620002b085828601620001e8565b925050602083015167ffffffffffffffff811115620002ce57600080fd5b620002dc85828601620001ff565b9150509250929050565b6000604051905081810181811067ffffffffffffffff821117156200030a57600080fd5b8060405250919050565b600067ffffffffffffffff8211156200032c57600080fd5b602082029050602081019050919050565b60006200034a8262000351565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200037c816200033d565b81146200038857600080fd5b50565b612699806200039b6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80638eb60066116100a2578063b600ffdb11610071578063b600ffdb146102be578063e41be775146102ef578063f203a5ed1461030b578063f851a4401461033b578063faea31671461035957610116565b80638eb600661461022357806390500bae146102535780639714378c146102865780639e6371ba146102a257610116565b80635c975abb116100e95780635c975abb1461018f5780636ff968c3146101ad57806379599f96146101cb5780638456cb59146101e957806384dc44cc146101f357610116565b80633aacfdad1461011b5780633f4ba83a1461014b57806347c4b328146101555780634e71d92d14610185575b600080fd5b610135600480360381019061013091906118ab565b610389565b60405161014291906121db565b60405180910390f35b6101536104a9565b005b61016f600480360381019061016a9190611982565b6105db565b60405161017c91906121db565b60405180910390f35b61018d610705565b005b610197610895565b6040516101a491906121db565b60405180910390f35b6101b56108a8565b6040516101c29190612153565b60405180910390f35b6101d36108ce565b6040516101e0919061233d565b60405180910390f35b6101f16108d4565b005b61020d60048036038101906102089190611a15565b610a4a565b60405161021a91906121db565b60405180910390f35b61023d60048036038101906102389190611d2f565b610c34565b60405161024a9190612189565b60405180910390f35b61026d60048036038101906102689190611b80565b610d10565b60405161027d9493929190612381565b60405180910390f35b6102a0600480360381019061029b9190611d83565b610d8b565b005b6102bc60048036038101906102b79190611882565b610eec565b005b6102d860048036038101906102d39190611d83565b611003565b6040516102e69291906121ab565b60405180910390f35b61030960048036038101906103049190611bc1565b611131565b005b61032560048036038101906103209190611d83565b611332565b604051610332919061233d565b60405180910390f35b61034361135d565b6040516103509190612153565b60405180910390f35b610373600480360381019061036e9190611d83565b611382565b604051610380919061231b565b60405180910390f35b6000825184511461039957600080fd5b81518451146103a757600080fd5b60008090505b845181101561049a578681815181106103c257fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166001878784815181106103ef57fe5b602002602001015187858151811061040357fe5b602002602001015187868151811061041757fe5b60200260200101516040516000815260200160405260405161043c94939291906121f6565b6020604051602081039080840390855afa15801561045e573d6000803e3d6000fd5b5050506020604051035173ffffffffffffffffffffffffffffffffffffffff161461048d5760009150506104a0565b80806001019150506103ad565b50600190505b95945050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610538576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052f906122fb565b60405180910390fd5b600160149054906101000a900460ff16610587576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057e9061223b565b60405180910390fd5b6000600160146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336040516105d1919061216e565b60405180910390a1565b600082518251146105eb57600080fd5b600085905060008090505b84518110156106f65783818151811061060b57fe5b60200260200101511561065b5784818151811061062457fe5b60200260200101518260405160200161063e929190612127565b6040516020818303038152906040528051906020012091506106e9565b6000801b85828151811061066b57fe5b602002602001015114156106a957818260405160200161068c929190612127565b6040516020818303038152906040528051906020012091506106e8565b818582815181106106b657fe5b60200260200101516040516020016106cf929190612127565b6040516020818303038152906040528051906020012091505b5b80806001019150506105f6565b50848114915050949350505050565b6002544210610749576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107409061225b565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d09061229b565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405161088b9190612153565b60405180910390a1565b600160149054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610963576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095a906122fb565b60405180910390fd5b600160149054906101000a900460ff16156109b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109aa906122db565b60405180910390fd5b60025442106109f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ee9061225b565b60405180910390fd5b60018060146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833604051610a40919061216e565b60405180910390a1565b600060606000610a598c611003565b80925081935050508651865114610a6f57600080fd5b8451865114610a7d57600080fd5b8351865114610a8b57600080fd5b60008090505b8751811015610b7d57600081118015610ad35750876001820381518110610ab457fe5b6020026020010151888281518110610ac857fe5b602002602001015111155b80610af257508251888281518110610ae757fe5b602002602001015110155b15610b035760009350505050610c26565b82888281518110610b1057fe5b602002602001015181518110610b2257fe5b6020026020010151838281518110610b3657fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610a91565b506000888a604051602001610b93929190612127565b60405160208183030381529060405280519060200120604051602001610bb9919061210c565b6040516020818303038152906040528051906020012090506003600284510281610bdf57fe5b04885111610bf35760009350505050610c26565b610c008382898989610389565b610c0957600080fd5b610c158e8b8e8e6105db565b610c1e57600080fd5b600193505050505b9a9950505050505050505050565b606060208202604201835114610c4957600080fd5b60608267ffffffffffffffff81118015610c6257600080fd5b50604051908082528060200260200182016040528015610c915781602001602082028036833780820191505090505b509050600080600090505b84811015610d045760208102606287010151915081838281518110610cbd57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610c9c565b50819250505092915050565b600080600080604285511015610d2557600080fd5b600085600081518110610d3457fe5b602001015160f81c60f81b60f81c9050600086600181518110610d5357fe5b602001015160f81c60f81b60f81c90506000806022890151915060428901519050838383839750975097509750505050509193509193565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e11906122fb565b60405180910390fd5b6002544210610e5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e559061225b565b60405180910390fd5b61016e8110610ea2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e999061227b565b60405180910390fd5b620151808102600254016002819055507f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e881604051610ee1919061233d565b60405180910390a150565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f7b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f72906122fb565b60405180910390fd5b6002544210610fbf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb69061225b565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b606060008060009050600060038054905090506000811161102357600080fd5b6001810390505b80821461107f57600060026001838501018161104257fe5b049050856003828154811061105357fe5b9060005260206000209060020201600101541161107257809250611079565b6001810391505b5061102a565b6003828154811061108c57fe5b9060005260206000209060020201600001828180548060200260200160405190810160405280929190818152602001828054801561111f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116110d5575b50505050509150935093505050915091565b600160149054906101000a900460ff1615611181576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611178906122db565b60405180910390fd5b6000898051906020012090506111c5816003600160038054905003815481106111a657fe5b9060005260206000209060020201600101548b8b8b8b8b8b8b8b610a4a565b6111ce57600080fd5b6000806000806111dd8e610d10565b935093509350935060468460ff161480156111fb575060018360ff16145b61120457600080fd5b60036001600380549050038154811061121957fe5b906000526020600020906002020160010154821161126c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611263906122bb565b60405180910390fd5b60606112788f83610c34565b90506003604051806040016040528083815260200185815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190805190602001906112d692919061144f565b506020820151816001015550507fe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d160038054905084604051611319929190612358565b60405180910390a1505050505050505050505050505050565b6003818154811061133f57fe5b90600052602060002090600202016000915090508060010154905081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61138a6114d9565b6003828154811061139757fe5b90600052602060002090600202016040518060400160405290816000820180548060200260200160405190810160405280929190818152602001828054801561143557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116113eb575b505050505081526020016001820154815250509050919050565b8280548282559060005260206000209081019282156114c8579160200282015b828111156114c75782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019061146f565b5b5090506114d591906114f3565b5090565b604051806040016040528060608152602001600081525090565b61153391905b8082111561152f57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016114f9565b5090565b90565b600081359050611545816125f0565b92915050565b600082601f83011261155c57600080fd5b813561156f61156a826123f3565b6123c6565b9150818183526020840193506020810190508385602084028201111561159457600080fd5b60005b838110156115c457816115aa8882611536565b845260208401935060208301925050600181019050611597565b5050505092915050565b600082601f8301126115df57600080fd5b81356115f26115ed8261241b565b6123c6565b9150818183526020840193506020810190508385602084028201111561161757600080fd5b60005b83811015611647578161162d88826117da565b84526020840193506020830192505060018101905061161a565b5050505092915050565b600082601f83011261166257600080fd5b813561167561167082612443565b6123c6565b9150818183526020840193506020810190508385602084028201111561169a57600080fd5b60005b838110156116ca57816116b088826117ef565b84526020840193506020830192505060018101905061169d565b5050505092915050565b600082601f8301126116e557600080fd5b81356116f86116f38261246b565b6123c6565b9150818183526020840193506020810190508385602084028201111561171d57600080fd5b60005b8381101561174d57816117338882611858565b845260208401935060208301925050600181019050611720565b5050505092915050565b600082601f83011261176857600080fd5b813561177b61177682612493565b6123c6565b915081818352602084019350602081019050838560208402820111156117a057600080fd5b60005b838110156117d057816117b6888261186d565b8452602084019350602083019250506001810190506117a3565b5050505092915050565b6000813590506117e981612607565b92915050565b6000813590506117fe8161261e565b92915050565b600082601f83011261181557600080fd5b8135611828611823826124bb565b6123c6565b9150808252602083016020830185838301111561184457600080fd5b61184f8382846125d7565b50505092915050565b60008135905061186781612635565b92915050565b60008135905061187c8161264c565b92915050565b60006020828403121561189457600080fd5b60006118a284828501611536565b91505092915050565b600080600080600060a086880312156118c357600080fd5b600086013567ffffffffffffffff8111156118dd57600080fd5b6118e98882890161154b565b95505060206118fa888289016117ef565b945050604086013567ffffffffffffffff81111561191757600080fd5b61192388828901611757565b935050606086013567ffffffffffffffff81111561194057600080fd5b61194c88828901611651565b925050608086013567ffffffffffffffff81111561196957600080fd5b61197588828901611651565b9150509295509295909350565b6000806000806080858703121561199857600080fd5b60006119a6878288016117ef565b94505060206119b7878288016117ef565b935050604085013567ffffffffffffffff8111156119d457600080fd5b6119e087828801611651565b925050606085013567ffffffffffffffff8111156119fd57600080fd5b611a09878288016115ce565b91505092959194509250565b6000806000806000806000806000806101408b8d031215611a3557600080fd5b6000611a438d828e016117ef565b9a50506020611a548d828e01611858565b99505060408b013567ffffffffffffffff811115611a7157600080fd5b611a7d8d828e01611651565b98505060608b013567ffffffffffffffff811115611a9a57600080fd5b611aa68d828e016115ce565b9750506080611ab78d828e016117ef565b96505060a0611ac88d828e016117ef565b95505060c08b013567ffffffffffffffff811115611ae557600080fd5b611af18d828e016116d4565b94505060e08b013567ffffffffffffffff811115611b0e57600080fd5b611b1a8d828e01611757565b9350506101008b013567ffffffffffffffff811115611b3857600080fd5b611b448d828e01611651565b9250506101208b013567ffffffffffffffff811115611b6257600080fd5b611b6e8d828e01611651565b9150509295989b9194979a5092959850565b600060208284031215611b9257600080fd5b600082013567ffffffffffffffff811115611bac57600080fd5b611bb884828501611804565b91505092915050565b60008060008060008060008060006101208a8c031215611be057600080fd5b60008a013567ffffffffffffffff811115611bfa57600080fd5b611c068c828d01611804565b99505060208a013567ffffffffffffffff811115611c2357600080fd5b611c2f8c828d01611651565b98505060408a013567ffffffffffffffff811115611c4c57600080fd5b611c588c828d016115ce565b9750506060611c698c828d016117ef565b9650506080611c7a8c828d016117ef565b95505060a08a013567ffffffffffffffff811115611c9757600080fd5b611ca38c828d016116d4565b94505060c08a013567ffffffffffffffff811115611cc057600080fd5b611ccc8c828d01611757565b93505060e08a013567ffffffffffffffff811115611ce957600080fd5b611cf58c828d01611651565b9250506101008a013567ffffffffffffffff811115611d1357600080fd5b611d1f8c828d01611651565b9150509295985092959850929598565b60008060408385031215611d4257600080fd5b600083013567ffffffffffffffff811115611d5c57600080fd5b611d6885828601611804565b9250506020611d7985828601611858565b9150509250929050565b600060208284031215611d9557600080fd5b6000611da384828501611858565b91505092915050565b6000611db88383611dd3565b60208301905092915050565b611dcd816125a1565b82525050565b611ddc81612542565b82525050565b611deb81612542565b82525050565b6000611dfc826124f7565b611e06818561250f565b9350611e11836124e7565b8060005b83811015611e42578151611e298882611dac565b9750611e3483612502565b925050600181019050611e15565b5085935050505092915050565b6000611e5a826124f7565b611e648185612520565b9350611e6f836124e7565b8060005b83811015611ea0578151611e878882611dac565b9750611e9283612502565b925050600181019050611e73565b5085935050505092915050565b611eb681612554565b82525050565b611ec581612560565b82525050565b611edc611ed782612560565b6125e6565b82525050565b6000611eef601483612531565b91507f6e6f7420706175736564207269676874206e6f770000000000000000000000006000830152602082019050919050565b6000611f2f600783612531565b91507f65787069726564000000000000000000000000000000000000000000000000006000830152602082019050919050565b6000611f6f601a83612531565b91507f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e670000000000006000830152602082019050919050565b6000611faf600c83612531565b91507f756e617574686f72697a656400000000000000000000000000000000000000006000830152602082019050919050565b6000611fef601b83612531565b91507f63616e6e6f74206368616e6765206f6c6420636f6d6d697474656500000000006000830152602082019050919050565b600061202f601083612531565b91507f706175736564207269676874206e6f77000000000000000000000000000000006000830152602082019050919050565b600061206f600983612531565b91507f6e6f742061646d696e00000000000000000000000000000000000000000000006000830152602082019050919050565b600060408301600083015184820360008601526120bf8282611df1565b91505060208301516120d460208601826120df565b508091505092915050565b6120e88161258a565b82525050565b6120f78161258a565b82525050565b61210681612594565b82525050565b60006121188284611ecb565b60208201915081905092915050565b60006121338285611ecb565b6020820191506121438284611ecb565b6020820191508190509392505050565b60006020820190506121686000830184611de2565b92915050565b60006020820190506121836000830184611dc4565b92915050565b600060208201905081810360008301526121a38184611e4f565b905092915050565b600060408201905081810360008301526121c58185611e4f565b90506121d460208301846120ee565b9392505050565b60006020820190506121f06000830184611ead565b92915050565b600060808201905061220b6000830187611ebc565b61221860208301866120fd565b6122256040830185611ebc565b6122326060830184611ebc565b95945050505050565b6000602082019050818103600083015261225481611ee2565b9050919050565b6000602082019050818103600083015261227481611f22565b9050919050565b6000602082019050818103600083015261229481611f62565b9050919050565b600060208201905081810360008301526122b481611fa2565b9050919050565b600060208201905081810360008301526122d481611fe2565b9050919050565b600060208201905081810360008301526122f481612022565b9050919050565b6000602082019050818103600083015261231481612062565b9050919050565b6000602082019050818103600083015261233581846120a2565b905092915050565b600060208201905061235260008301846120ee565b92915050565b600060408201905061236d60008301856120ee565b61237a60208301846120ee565b9392505050565b600060808201905061239660008301876120fd565b6123a360208301866120fd565b6123b060408301856120ee565b6123bd60608301846120ee565b95945050505050565b6000604051905081810181811067ffffffffffffffff821117156123e957600080fd5b8060405250919050565b600067ffffffffffffffff82111561240a57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561243257600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561245a57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561248257600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156124aa57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156124d257600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061254d8261256a565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006125ac826125b3565b9050919050565b60006125be826125c5565b9050919050565b60006125d08261256a565b9050919050565b82818337600083830152505050565b6000819050919050565b6125f981612542565b811461260457600080fd5b50565b61261081612554565b811461261b57600080fd5b50565b61262781612560565b811461263257600080fd5b50565b61263e8161258a565b811461264957600080fd5b50565b61265581612594565b811461266057600080fd5b5056fea2646970667358221220eeb776d7d62977ef969b498a55e9657eb8329302d1ed8960513a8bccbee9742e64736f6c63430006060033"

// DeployIncognitoproxy deploys a new Ethereum contract, binding an instance of Incognitoproxy to it.
func DeployIncognitoproxy(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, beaconCommittee []common.Address) (common.Address, *types.Transaction, *Incognitoproxy, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoproxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IncognitoproxyBin), backend, _admin, beaconCommittee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Incognitoproxy{IncognitoproxyCaller: IncognitoproxyCaller{contract: contract}, IncognitoproxyTransactor: IncognitoproxyTransactor{contract: contract}, IncognitoproxyFilterer: IncognitoproxyFilterer{contract: contract}}, nil
}

// Incognitoproxy is an auto generated Go binding around an Ethereum contract.
type Incognitoproxy struct {
	IncognitoproxyCaller     // Read-only binding to the contract
	IncognitoproxyTransactor // Write-only binding to the contract
	IncognitoproxyFilterer   // Log filterer for contract events
}

// IncognitoproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncognitoproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncognitoproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncognitoproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncognitoproxySession struct {
	Contract     *Incognitoproxy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncognitoproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncognitoproxyCallerSession struct {
	Contract *IncognitoproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IncognitoproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncognitoproxyTransactorSession struct {
	Contract     *IncognitoproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IncognitoproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncognitoproxyRaw struct {
	Contract *Incognitoproxy // Generic contract binding to access the raw methods on
}

// IncognitoproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncognitoproxyCallerRaw struct {
	Contract *IncognitoproxyCaller // Generic read-only contract binding to access the raw methods on
}

// IncognitoproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncognitoproxyTransactorRaw struct {
	Contract *IncognitoproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncognitoproxy creates a new instance of Incognitoproxy, bound to a specific deployed contract.
func NewIncognitoproxy(address common.Address, backend bind.ContractBackend) (*Incognitoproxy, error) {
	contract, err := bindIncognitoproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Incognitoproxy{IncognitoproxyCaller: IncognitoproxyCaller{contract: contract}, IncognitoproxyTransactor: IncognitoproxyTransactor{contract: contract}, IncognitoproxyFilterer: IncognitoproxyFilterer{contract: contract}}, nil
}

// NewIncognitoproxyCaller creates a new read-only instance of Incognitoproxy, bound to a specific deployed contract.
func NewIncognitoproxyCaller(address common.Address, caller bind.ContractCaller) (*IncognitoproxyCaller, error) {
	contract, err := bindIncognitoproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoproxyCaller{contract: contract}, nil
}

// NewIncognitoproxyTransactor creates a new write-only instance of Incognitoproxy, bound to a specific deployed contract.
func NewIncognitoproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IncognitoproxyTransactor, error) {
	contract, err := bindIncognitoproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoproxyTransactor{contract: contract}, nil
}

// NewIncognitoproxyFilterer creates a new log filterer instance of Incognitoproxy, bound to a specific deployed contract.
func NewIncognitoproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IncognitoproxyFilterer, error) {
	contract, err := bindIncognitoproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncognitoproxyFilterer{contract: contract}, nil
}

// bindIncognitoproxy binds a generic wrapper to an already deployed contract.
func bindIncognitoproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoproxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incognitoproxy *IncognitoproxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Incognitoproxy.Contract.IncognitoproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incognitoproxy *IncognitoproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognitoproxy.Contract.IncognitoproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incognitoproxy *IncognitoproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incognitoproxy.Contract.IncognitoproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incognitoproxy *IncognitoproxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Incognitoproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incognitoproxy *IncognitoproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognitoproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incognitoproxy *IncognitoproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incognitoproxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Incognitoproxy *IncognitoproxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Incognitoproxy.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Incognitoproxy *IncognitoproxySession) Admin() (common.Address, error) {
	return _Incognitoproxy.Contract.Admin(&_Incognitoproxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Incognitoproxy *IncognitoproxyCallerSession) Admin() (common.Address, error) {
	return _Incognitoproxy.Contract.Admin(&_Incognitoproxy.CallOpts)
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock)
func (_Incognitoproxy *IncognitoproxyCaller) BeaconCommittees(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Incognitoproxy.contract.Call(opts, out, "beaconCommittees", arg0)
	return *ret0, err
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock)
func (_Incognitoproxy *IncognitoproxySession) BeaconCommittees(arg0 *big.Int) (*big.Int, error) {
	return _Incognitoproxy.Contract.BeaconCommittees(&_Incognitoproxy.CallOpts, arg0)
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock)
func (_Incognitoproxy *IncognitoproxyCallerSession) BeaconCommittees(arg0 *big.Int) (*big.Int, error) {
	return _Incognitoproxy.Contract.BeaconCommittees(&_Incognitoproxy.CallOpts, arg0)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_Incognitoproxy *IncognitoproxyCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Incognitoproxy.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_Incognitoproxy *IncognitoproxySession) Expire() (*big.Int, error) {
	return _Incognitoproxy.Contract.Expire(&_Incognitoproxy.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_Incognitoproxy *IncognitoproxyCallerSession) Expire() (*big.Int, error) {
	return _Incognitoproxy.Contract.Expire(&_Incognitoproxy.CallOpts)
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_Incognitoproxy *IncognitoproxyCaller) ExtractCommitteeFromInstruction(opts *bind.CallOpts, inst []byte, numVals *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Incognitoproxy.contract.Call(opts, out, "extractCommitteeFromInstruction", inst, numVals)
	return *ret0, err
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_Incognitoproxy *IncognitoproxySession) ExtractCommitteeFromInstruction(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _Incognitoproxy.Contract.ExtractCommitteeFromInstruction(&_Incognitoproxy.CallOpts, inst, numVals)
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_Incognitoproxy *IncognitoproxyCallerSession) ExtractCommitteeFromInstruction(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _Incognitoproxy.Contract.ExtractCommitteeFromInstruction(&_Incognitoproxy.CallOpts, inst, numVals)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256)
func (_Incognitoproxy *IncognitoproxyCaller) ExtractMetaFromInstruction(opts *bind.CallOpts, inst []byte) (uint8, uint8, *big.Int, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(uint8)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Incognitoproxy.contract.Call(opts, out, "extractMetaFromInstruction", inst)
	return *ret0, *ret1, *ret2, *ret3, err
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256)
func (_Incognitoproxy *IncognitoproxySession) ExtractMetaFromInstruction(inst []byte) (uint8, uint8, *big.Int, *big.Int, error) {
	return _Incognitoproxy.Contract.ExtractMetaFromInstruction(&_Incognitoproxy.CallOpts, inst)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256)
func (_Incognitoproxy *IncognitoproxyCallerSession) ExtractMetaFromInstruction(inst []byte) (uint8, uint8, *big.Int, *big.Int, error) {
	return _Incognitoproxy.Contract.ExtractMetaFromInstruction(&_Incognitoproxy.CallOpts, inst)
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_Incognitoproxy *IncognitoproxyCaller) FindBeaconCommitteeFromHeight(opts *bind.CallOpts, blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Incognitoproxy.contract.Call(opts, out, "findBeaconCommitteeFromHeight", blkHeight)
	return *ret0, *ret1, err
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_Incognitoproxy *IncognitoproxySession) FindBeaconCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	return _Incognitoproxy.Contract.FindBeaconCommitteeFromHeight(&_Incognitoproxy.CallOpts, blkHeight)
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) view returns(address[], uint256)
func (_Incognitoproxy *IncognitoproxyCallerSession) FindBeaconCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, *big.Int, error) {
	return _Incognitoproxy.Contract.FindBeaconCommitteeFromHeight(&_Incognitoproxy.CallOpts, blkHeight)
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256))
func (_Incognitoproxy *IncognitoproxyCaller) GetBeaconCommittee(opts *bind.CallOpts, i *big.Int) (IncognitoProxyCommittee, error) {
	var (
		ret0 = new(IncognitoProxyCommittee)
	)
	out := ret0
	err := _Incognitoproxy.contract.Call(opts, out, "getBeaconCommittee", i)
	return *ret0, err
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256))
func (_Incognitoproxy *IncognitoproxySession) GetBeaconCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _Incognitoproxy.Contract.GetBeaconCommittee(&_Incognitoproxy.CallOpts, i)
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256))
func (_Incognitoproxy *IncognitoproxyCallerSession) GetBeaconCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _Incognitoproxy.Contract.GetBeaconCommittee(&_Incognitoproxy.CallOpts, i)
}

// InstructionApproved is a free data retrieval call binding the contract method 0x84dc44cc.
//
// Solidity: function instructionApproved(bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_Incognitoproxy *IncognitoproxyCaller) InstructionApproved(opts *bind.CallOpts, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incognitoproxy.contract.Call(opts, out, "instructionApproved", instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
	return *ret0, err
}

// InstructionApproved is a free data retrieval call binding the contract method 0x84dc44cc.
//
// Solidity: function instructionApproved(bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_Incognitoproxy *IncognitoproxySession) InstructionApproved(instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _Incognitoproxy.Contract.InstructionApproved(&_Incognitoproxy.CallOpts, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionApproved is a free data retrieval call binding the contract method 0x84dc44cc.
//
// Solidity: function instructionApproved(bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) view returns(bool)
func (_Incognitoproxy *IncognitoproxyCallerSession) InstructionApproved(instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _Incognitoproxy.Contract.InstructionApproved(&_Incognitoproxy.CallOpts, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionInMerkleTree is a free data retrieval call binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) pure returns(bool)
func (_Incognitoproxy *IncognitoproxyCaller) InstructionInMerkleTree(opts *bind.CallOpts, leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incognitoproxy.contract.Call(opts, out, "instructionInMerkleTree", leaf, root, path, left)
	return *ret0, err
}

// InstructionInMerkleTree is a free data retrieval call binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) pure returns(bool)
func (_Incognitoproxy *IncognitoproxySession) InstructionInMerkleTree(leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	return _Incognitoproxy.Contract.InstructionInMerkleTree(&_Incognitoproxy.CallOpts, leaf, root, path, left)
}

// InstructionInMerkleTree is a free data retrieval call binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) pure returns(bool)
func (_Incognitoproxy *IncognitoproxyCallerSession) InstructionInMerkleTree(leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	return _Incognitoproxy.Contract.InstructionInMerkleTree(&_Incognitoproxy.CallOpts, leaf, root, path, left)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Incognitoproxy *IncognitoproxyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incognitoproxy.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Incognitoproxy *IncognitoproxySession) Paused() (bool, error) {
	return _Incognitoproxy.Contract.Paused(&_Incognitoproxy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Incognitoproxy *IncognitoproxyCallerSession) Paused() (bool, error) {
	return _Incognitoproxy.Contract.Paused(&_Incognitoproxy.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_Incognitoproxy *IncognitoproxyCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Incognitoproxy.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_Incognitoproxy *IncognitoproxySession) Successor() (common.Address, error) {
	return _Incognitoproxy.Contract.Successor(&_Incognitoproxy.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_Incognitoproxy *IncognitoproxyCallerSession) Successor() (common.Address, error) {
	return _Incognitoproxy.Contract.Successor(&_Incognitoproxy.CallOpts)
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_Incognitoproxy *IncognitoproxyCaller) VerifySig(opts *bind.CallOpts, committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incognitoproxy.contract.Call(opts, out, "verifySig", committee, msgHash, v, r, s)
	return *ret0, err
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_Incognitoproxy *IncognitoproxySession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	return _Incognitoproxy.Contract.VerifySig(&_Incognitoproxy.CallOpts, committee, msgHash, v, r, s)
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_Incognitoproxy *IncognitoproxyCallerSession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	return _Incognitoproxy.Contract.VerifySig(&_Incognitoproxy.CallOpts, committee, msgHash, v, r, s)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Incognitoproxy *IncognitoproxyTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognitoproxy.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Incognitoproxy *IncognitoproxySession) Claim() (*types.Transaction, error) {
	return _Incognitoproxy.Contract.Claim(&_Incognitoproxy.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Incognitoproxy *IncognitoproxyTransactorSession) Claim() (*types.Transaction, error) {
	return _Incognitoproxy.Contract.Claim(&_Incognitoproxy.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Incognitoproxy *IncognitoproxyTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Incognitoproxy.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Incognitoproxy *IncognitoproxySession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Incognitoproxy.Contract.Extend(&_Incognitoproxy.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Incognitoproxy *IncognitoproxyTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Incognitoproxy.Contract.Extend(&_Incognitoproxy.TransactOpts, n)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Incognitoproxy *IncognitoproxyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognitoproxy.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Incognitoproxy *IncognitoproxySession) Pause() (*types.Transaction, error) {
	return _Incognitoproxy.Contract.Pause(&_Incognitoproxy.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Incognitoproxy *IncognitoproxyTransactorSession) Pause() (*types.Transaction, error) {
	return _Incognitoproxy.Contract.Pause(&_Incognitoproxy.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Incognitoproxy *IncognitoproxyTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _Incognitoproxy.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Incognitoproxy *IncognitoproxySession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Incognitoproxy.Contract.Retire(&_Incognitoproxy.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Incognitoproxy *IncognitoproxyTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Incognitoproxy.Contract.Retire(&_Incognitoproxy.TransactOpts, _successor)
}

// SwapBeaconCommittee is a paid mutator transaction binding the contract method 0xe41be775.
//
// Solidity: function swapBeaconCommittee(bytes inst, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) returns()
func (_Incognitoproxy *IncognitoproxyTransactor) SwapBeaconCommittee(opts *bind.TransactOpts, inst []byte, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (*types.Transaction, error) {
	return _Incognitoproxy.contract.Transact(opts, "swapBeaconCommittee", inst, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// SwapBeaconCommittee is a paid mutator transaction binding the contract method 0xe41be775.
//
// Solidity: function swapBeaconCommittee(bytes inst, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) returns()
func (_Incognitoproxy *IncognitoproxySession) SwapBeaconCommittee(inst []byte, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (*types.Transaction, error) {
	return _Incognitoproxy.Contract.SwapBeaconCommittee(&_Incognitoproxy.TransactOpts, inst, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// SwapBeaconCommittee is a paid mutator transaction binding the contract method 0xe41be775.
//
// Solidity: function swapBeaconCommittee(bytes inst, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) returns()
func (_Incognitoproxy *IncognitoproxyTransactorSession) SwapBeaconCommittee(inst []byte, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (*types.Transaction, error) {
	return _Incognitoproxy.Contract.SwapBeaconCommittee(&_Incognitoproxy.TransactOpts, inst, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Incognitoproxy *IncognitoproxyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognitoproxy.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Incognitoproxy *IncognitoproxySession) Unpause() (*types.Transaction, error) {
	return _Incognitoproxy.Contract.Unpause(&_Incognitoproxy.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Incognitoproxy *IncognitoproxyTransactorSession) Unpause() (*types.Transaction, error) {
	return _Incognitoproxy.Contract.Unpause(&_Incognitoproxy.TransactOpts)
}

// IncognitoproxyBeaconCommitteeSwappedIterator is returned from FilterBeaconCommitteeSwapped and is used to iterate over the raw logs and unpacked data for BeaconCommitteeSwapped events raised by the Incognitoproxy contract.
type IncognitoproxyBeaconCommitteeSwappedIterator struct {
	Event *IncognitoproxyBeaconCommitteeSwapped // Event containing the contract specifics and raw log

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
func (it *IncognitoproxyBeaconCommitteeSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoproxyBeaconCommitteeSwapped)
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
		it.Event = new(IncognitoproxyBeaconCommitteeSwapped)
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
func (it *IncognitoproxyBeaconCommitteeSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoproxyBeaconCommitteeSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoproxyBeaconCommitteeSwapped represents a BeaconCommitteeSwapped event raised by the Incognitoproxy contract.
type IncognitoproxyBeaconCommitteeSwapped struct {
	Id          *big.Int
	StartHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeaconCommitteeSwapped is a free log retrieval operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_Incognitoproxy *IncognitoproxyFilterer) FilterBeaconCommitteeSwapped(opts *bind.FilterOpts) (*IncognitoproxyBeaconCommitteeSwappedIterator, error) {

	logs, sub, err := _Incognitoproxy.contract.FilterLogs(opts, "BeaconCommitteeSwapped")
	if err != nil {
		return nil, err
	}
	return &IncognitoproxyBeaconCommitteeSwappedIterator{contract: _Incognitoproxy.contract, event: "BeaconCommitteeSwapped", logs: logs, sub: sub}, nil
}

// WatchBeaconCommitteeSwapped is a free log subscription operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_Incognitoproxy *IncognitoproxyFilterer) WatchBeaconCommitteeSwapped(opts *bind.WatchOpts, sink chan<- *IncognitoproxyBeaconCommitteeSwapped) (event.Subscription, error) {

	logs, sub, err := _Incognitoproxy.contract.WatchLogs(opts, "BeaconCommitteeSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoproxyBeaconCommitteeSwapped)
				if err := _Incognitoproxy.contract.UnpackLog(event, "BeaconCommitteeSwapped", log); err != nil {
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

// ParseBeaconCommitteeSwapped is a log parse operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_Incognitoproxy *IncognitoproxyFilterer) ParseBeaconCommitteeSwapped(log types.Log) (*IncognitoproxyBeaconCommitteeSwapped, error) {
	event := new(IncognitoproxyBeaconCommitteeSwapped)
	if err := _Incognitoproxy.contract.UnpackLog(event, "BeaconCommitteeSwapped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoproxyClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Incognitoproxy contract.
type IncognitoproxyClaimIterator struct {
	Event *IncognitoproxyClaim // Event containing the contract specifics and raw log

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
func (it *IncognitoproxyClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoproxyClaim)
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
		it.Event = new(IncognitoproxyClaim)
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
func (it *IncognitoproxyClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoproxyClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoproxyClaim represents a Claim event raised by the Incognitoproxy contract.
type IncognitoproxyClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Incognitoproxy *IncognitoproxyFilterer) FilterClaim(opts *bind.FilterOpts) (*IncognitoproxyClaimIterator, error) {

	logs, sub, err := _Incognitoproxy.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &IncognitoproxyClaimIterator{contract: _Incognitoproxy.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Incognitoproxy *IncognitoproxyFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *IncognitoproxyClaim) (event.Subscription, error) {

	logs, sub, err := _Incognitoproxy.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoproxyClaim)
				if err := _Incognitoproxy.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_Incognitoproxy *IncognitoproxyFilterer) ParseClaim(log types.Log) (*IncognitoproxyClaim, error) {
	event := new(IncognitoproxyClaim)
	if err := _Incognitoproxy.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoproxyExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the Incognitoproxy contract.
type IncognitoproxyExtendIterator struct {
	Event *IncognitoproxyExtend // Event containing the contract specifics and raw log

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
func (it *IncognitoproxyExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoproxyExtend)
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
		it.Event = new(IncognitoproxyExtend)
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
func (it *IncognitoproxyExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoproxyExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoproxyExtend represents a Extend event raised by the Incognitoproxy contract.
type IncognitoproxyExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Incognitoproxy *IncognitoproxyFilterer) FilterExtend(opts *bind.FilterOpts) (*IncognitoproxyExtendIterator, error) {

	logs, sub, err := _Incognitoproxy.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &IncognitoproxyExtendIterator{contract: _Incognitoproxy.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Incognitoproxy *IncognitoproxyFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *IncognitoproxyExtend) (event.Subscription, error) {

	logs, sub, err := _Incognitoproxy.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoproxyExtend)
				if err := _Incognitoproxy.contract.UnpackLog(event, "Extend", log); err != nil {
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

// ParseExtend is a log parse operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Incognitoproxy *IncognitoproxyFilterer) ParseExtend(log types.Log) (*IncognitoproxyExtend, error) {
	event := new(IncognitoproxyExtend)
	if err := _Incognitoproxy.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoproxyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Incognitoproxy contract.
type IncognitoproxyPausedIterator struct {
	Event *IncognitoproxyPaused // Event containing the contract specifics and raw log

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
func (it *IncognitoproxyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoproxyPaused)
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
		it.Event = new(IncognitoproxyPaused)
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
func (it *IncognitoproxyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoproxyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoproxyPaused represents a Paused event raised by the Incognitoproxy contract.
type IncognitoproxyPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Incognitoproxy *IncognitoproxyFilterer) FilterPaused(opts *bind.FilterOpts) (*IncognitoproxyPausedIterator, error) {

	logs, sub, err := _Incognitoproxy.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IncognitoproxyPausedIterator{contract: _Incognitoproxy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Incognitoproxy *IncognitoproxyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IncognitoproxyPaused) (event.Subscription, error) {

	logs, sub, err := _Incognitoproxy.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoproxyPaused)
				if err := _Incognitoproxy.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Incognitoproxy *IncognitoproxyFilterer) ParsePaused(log types.Log) (*IncognitoproxyPaused, error) {
	event := new(IncognitoproxyPaused)
	if err := _Incognitoproxy.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoproxyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Incognitoproxy contract.
type IncognitoproxyUnpausedIterator struct {
	Event *IncognitoproxyUnpaused // Event containing the contract specifics and raw log

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
func (it *IncognitoproxyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoproxyUnpaused)
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
		it.Event = new(IncognitoproxyUnpaused)
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
func (it *IncognitoproxyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoproxyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoproxyUnpaused represents a Unpaused event raised by the Incognitoproxy contract.
type IncognitoproxyUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Incognitoproxy *IncognitoproxyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IncognitoproxyUnpausedIterator, error) {

	logs, sub, err := _Incognitoproxy.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IncognitoproxyUnpausedIterator{contract: _Incognitoproxy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Incognitoproxy *IncognitoproxyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IncognitoproxyUnpaused) (event.Subscription, error) {

	logs, sub, err := _Incognitoproxy.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoproxyUnpaused)
				if err := _Incognitoproxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Incognitoproxy *IncognitoproxyFilterer) ParseUnpaused(log types.Log) (*IncognitoproxyUnpaused, error) {
	event := new(IncognitoproxyUnpaused)
	if err := _Incognitoproxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

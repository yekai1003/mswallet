// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bcos

import (
	"math/big"
	"strings"

	"github.com/yekai1003/gobcos/accounts/abi"
	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/common"
	"github.com/yekai1003/gobcos/core/types"
	"github.com/yekai1003/gobcos/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = common.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenadminABI is the input ABI used to generate the binding from.
const TokenadminABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenaddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TokenadminBin is the compiled bytecode used for deploying new contracts.
var TokenadminBin = "0x608060405234801561001057600080fd5b50604051611e5e380380611e5e83398101806040528101908080518201929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508061007c61014e565b8080602001828103825283818151815260200191508051906020019080838360005b838110156100b957808201518184015260208101905061009e565b50505050905090810190601f1680156100e65780820380516001836020036101000a031916815260200191505b5092505050604051809103906000f080158015610107573d6000803e3d6000fd5b50600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061015e565b604051611567806108f783390190565b61078a8061016d6000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630900f0101461007d57806340c10f19146100c05780639dc29fac1461010d578063a9059cbb1461015a578063bce934a9146101a7578063f851a440146101fe575b600080fd5b34801561008957600080fd5b506100be600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610255565b005b3480156100cc57600080fd5b5061010b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506102f4565b005b34801561011957600080fd5b50610158600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610430565b005b34801561016657600080fd5b506101a5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061056c565b005b3480156101b357600080fd5b506101bc610671565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561020a57600080fd5b50610213610739565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156102b057600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561034f57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561041457600080fd5b505af1158015610428573d6000803e3d6000fd5b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561048b57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561055057600080fd5b505af1158015610564573d6000803e3d6000fd5b505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561063157600080fd5b505af1158015610645573d6000803e3d6000fd5b505050506040513d602081101561065b57600080fd5b8101908080519060200190929190505050505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663767800de6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156106f957600080fd5b505af115801561070d573d6000803e3d6000fd5b505050506040513d602081101561072357600080fd5b8101908080519060200190929190505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820b62b8463cc4350e2d12f0e8c90d5730d9f94967917350103913d8116a6ac5d60002960806040523480156200001157600080fd5b5060405162001567380380620015678339810180604052810190808051820192919050505060006004819055508060009080519060200190620000569291906200009f565b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200014e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000e257805160ff191683800117855562000113565b8280016001018555821562000113579182015b8281111562000112578251825591602001919060010190620000f5565b5b50905062000122919062000126565b5090565b6200014b91905b80821115620001475760008160009055506001016200012d565b5090565b90565b611409806200015e6000396000f3006080604052600436106100c5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063095ea7b3146100ca57806318160ddd1461012f57806323b872dd1461015a57806339509351146101df57806340c10f191461024457806370a0823114610291578063767800de146102e85780639dc29fac1461033f578063a457c2d71461038c578063a75b54c4146103f1578063a9059cbb14610448578063b09f1266146104ad578063dd62ed3e1461053d575b600080fd5b3480156100d657600080fd5b50610115600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105b4565b604051808215151515815260200191505060405180910390f35b34801561013b57600080fd5b506101446105cb565b6040518082815260200191505060405180910390f35b34801561016657600080fd5b506101c5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105d5565b604051808215151515815260200191505060405180910390f35b3480156101eb57600080fd5b5061022a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610686565b604051808215151515815260200191505060405180910390f35b34801561025057600080fd5b5061028f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061072b565b005b34801561029d57600080fd5b506102d2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610795565b6040518082815260200191505060405180910390f35b3480156102f457600080fd5b506102fd6107de565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561034b57600080fd5b5061038a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506107e6565b005b34801561039857600080fd5b506103d7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610850565b604051808215151515815260200191505060405180910390f35b3480156103fd57600080fd5b506104066108f5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561045457600080fd5b50610493600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061091b565b604051808215151515815260200191505060405180910390f35b3480156104b957600080fd5b506104c2610932565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105025780820151818401526020810190506104e7565b50505050905090810190601f16801561052f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561054957600080fd5b5061059e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109d0565b6040518082815260200191505060405180910390f35b60006105c1338484610a57565b6001905092915050565b6000600454905090565b60006105e2848484610cd8565b61067b843361067685600360008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461100290919063ffffffff16565b610a57565b600190509392505050565b6000610721338461071c85600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461101b90919063ffffffff16565b610a57565b6001905092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561078757600080fd5b6107918282611039565b5050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600030905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561084257600080fd5b61084c82826111f8565b5050565b60006108eb33846108e685600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461100290919063ffffffff16565b610a57565b6001905092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000610928338484610cd8565b6001905092915050565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109c85780601f1061099d576101008083540402835291602001916109c8565b820191906000526020600020905b8154815290600101906020018083116109ab57829003601f168201915b505050505081565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610b22576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001807f45524332303a20617070726f76652066726f6d20746865207a65726f2061646481526020017f726573730000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515610bed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001807f45524332303a20617070726f766520746f20746865207a65726f20616464726581526020017f737300000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b80600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610da3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001807f45524332303a207472616e736665722066726f6d20746865207a65726f20616481526020017f647265737300000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515610e6e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001807f45524332303a207472616e7366657220746f20746865207a65726f206164647281526020017f657373000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b610ec081600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461100290919063ffffffff16565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610f5581600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461101b90919063ffffffff16565b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b600082821115151561101057fe5b818303905092915050565b600080828401905083811015151561102f57fe5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156110de576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b6110f38160045461101b90919063ffffffff16565b60048190555061114b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461101b90919063ffffffff16565b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156112c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f45524332303a206275726e2066726f6d20746865207a65726f2061646472657381526020017f730000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b6112d88160045461100290919063ffffffff16565b60048190555061133081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461100290919063ffffffff16565b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a350505600a165627a7a72305820a2500ee85a5e4f09d0378835c7562887bbc381daad90e4d9ca5817b43df05c820029"

// DeployTokenadmin deploys a new Ethereum contract, binding an instance of Tokenadmin to it.
func DeployTokenadmin(auth *bind.TransactOpts, backend bind.ContractBackend, symbol string) (common.Address, *types.RawTransaction, *Tokenadmin, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenadminABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenadminBin), backend, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokenadmin{TokenadminCaller: TokenadminCaller{contract: contract}, TokenadminTransactor: TokenadminTransactor{contract: contract}, TokenadminFilterer: TokenadminFilterer{contract: contract}}, nil
}

// Tokenadmin is an auto generated Go binding around an Ethereum contract.
type Tokenadmin struct {
	TokenadminCaller     // Read-only binding to the contract
	TokenadminTransactor // Write-only binding to the contract
	TokenadminFilterer   // Log filterer for contract events
}

// TokenadminCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenadminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenadminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenadminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenadminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenadminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenadminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenadminSession struct {
	Contract     *Tokenadmin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenadminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenadminCallerSession struct {
	Contract *TokenadminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenadminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenadminTransactorSession struct {
	Contract     *TokenadminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenadminRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenadminRaw struct {
	Contract *Tokenadmin // Generic contract binding to access the raw methods on
}

// TokenadminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenadminCallerRaw struct {
	Contract *TokenadminCaller // Generic read-only contract binding to access the raw methods on
}

// TokenadminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenadminTransactorRaw struct {
	Contract *TokenadminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenadmin creates a new instance of Tokenadmin, bound to a specific deployed contract.
func NewTokenadmin(address common.Address, backend bind.ContractBackend) (*Tokenadmin, error) {
	contract, err := bindTokenadmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenadmin{TokenadminCaller: TokenadminCaller{contract: contract}, TokenadminTransactor: TokenadminTransactor{contract: contract}, TokenadminFilterer: TokenadminFilterer{contract: contract}}, nil
}

// NewTokenadminCaller creates a new read-only instance of Tokenadmin, bound to a specific deployed contract.
func NewTokenadminCaller(address common.Address, caller bind.ContractCaller) (*TokenadminCaller, error) {
	contract, err := bindTokenadmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenadminCaller{contract: contract}, nil
}

// NewTokenadminTransactor creates a new write-only instance of Tokenadmin, bound to a specific deployed contract.
func NewTokenadminTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenadminTransactor, error) {
	contract, err := bindTokenadmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenadminTransactor{contract: contract}, nil
}

// NewTokenadminFilterer creates a new log filterer instance of Tokenadmin, bound to a specific deployed contract.
func NewTokenadminFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenadminFilterer, error) {
	contract, err := bindTokenadmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenadminFilterer{contract: contract}, nil
}

// bindTokenadmin binds a generic wrapper to an already deployed contract.
func bindTokenadmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenadminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenadmin *TokenadminRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tokenadmin.Contract.TokenadminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenadmin *TokenadminRaw) Transfer(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.TokenadminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenadmin *TokenadminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.TokenadminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenadmin *TokenadminCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tokenadmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenadmin *TokenadminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenadmin *TokenadminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Tokenadmin *TokenadminCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tokenadmin.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Tokenadmin *TokenadminSession) Admin() (common.Address, error) {
	return _Tokenadmin.Contract.Admin(&_Tokenadmin.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Tokenadmin *TokenadminCallerSession) Admin() (common.Address, error) {
	return _Tokenadmin.Contract.Admin(&_Tokenadmin.CallOpts)
}

// Tokenaddr is a free data retrieval call binding the contract method 0xbce934a9.
//
// Solidity: function tokenaddr() constant returns(address)
func (_Tokenadmin *TokenadminCaller) Tokenaddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tokenadmin.contract.Call(opts, out, "tokenaddr")
	return *ret0, err
}

// Tokenaddr is a free data retrieval call binding the contract method 0xbce934a9.
//
// Solidity: function tokenaddr() constant returns(address)
func (_Tokenadmin *TokenadminSession) Tokenaddr() (common.Address, error) {
	return _Tokenadmin.Contract.Tokenaddr(&_Tokenadmin.CallOpts)
}

// Tokenaddr is a free data retrieval call binding the contract method 0xbce934a9.
//
// Solidity: function tokenaddr() constant returns(address)
func (_Tokenadmin *TokenadminCallerSession) Tokenaddr() (common.Address, error) {
	return _Tokenadmin.Contract.Tokenaddr(&_Tokenadmin.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address owner, uint256 value) returns()
func (_Tokenadmin *TokenadminTransactor) Burn(opts *bind.TransactOpts, owner common.Address, value *big.Int) (*types.RawTransaction, error) {
	return _Tokenadmin.contract.Transact(opts, "burn", owner, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address owner, uint256 value) returns()
func (_Tokenadmin *TokenadminSession) Burn(owner common.Address, value *big.Int) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.Burn(&_Tokenadmin.TransactOpts, owner, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address owner, uint256 value) returns()
func (_Tokenadmin *TokenadminTransactorSession) Burn(owner common.Address, value *big.Int) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.Burn(&_Tokenadmin.TransactOpts, owner, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Tokenadmin *TokenadminTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.RawTransaction, error) {
	return _Tokenadmin.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Tokenadmin *TokenadminSession) Mint(to common.Address, value *big.Int) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.Mint(&_Tokenadmin.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Tokenadmin *TokenadminTransactorSession) Mint(to common.Address, value *big.Int) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.Mint(&_Tokenadmin.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns()
func (_Tokenadmin *TokenadminTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.RawTransaction, error) {
	return _Tokenadmin.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns()
func (_Tokenadmin *TokenadminSession) Transfer(to common.Address, value *big.Int) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.Transfer(&_Tokenadmin.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns()
func (_Tokenadmin *TokenadminTransactorSession) Transfer(to common.Address, value *big.Int) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.Transfer(&_Tokenadmin.TransactOpts, to, value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address addr) returns()
func (_Tokenadmin *TokenadminTransactor) Upgrade(opts *bind.TransactOpts, addr common.Address) (*types.RawTransaction, error) {
	return _Tokenadmin.contract.Transact(opts, "upgrade", addr)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address addr) returns()
func (_Tokenadmin *TokenadminSession) Upgrade(addr common.Address) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.Upgrade(&_Tokenadmin.TransactOpts, addr)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address addr) returns()
func (_Tokenadmin *TokenadminTransactorSession) Upgrade(addr common.Address) (*types.RawTransaction, error) {
	return _Tokenadmin.Contract.Upgrade(&_Tokenadmin.TransactOpts, addr)
}

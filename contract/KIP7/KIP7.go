// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kip7

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Kip7MetaData contains all meta data concerning the Kip7 contract.
var Kip7MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
	Bin: "0x608060405234801561001057600080fd5b506100276301ffc9a760e01b61004260201b60201c565b61003d636578737160e01b61004260201b60201c565b61014a565b63ffffffff60e01b817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614156100de576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4b495031333a20696e76616c696420696e74657266616365206964000000000081525060200191505060405180910390fd5b6001600080837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b611244806101596000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806342842e0e1161007157806342842e0e1461026b57806370a08231146102d9578063a9059cbb14610331578063b88d4fde14610397578063dd62ed3e1461049c578063eb79554914610514576100a9565b806301ffc9a7146100ae578063095ea7b31461011357806318160ddd1461017957806323b872dd14610197578063423f6cef1461021d575b600080fd5b6100f9600480360360208110156100c457600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690602001909291905050506105f9565b604051808215151515815260200191505060405180910390f35b61015f6004803603604081101561012957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610660565b604051808215151515815260200191505060405180910390f35b610181610677565b6040518082815260200191505060405180910390f35b610203600480360360608110156101ad57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610681565b604051808215151515815260200191505060405180910390f35b6102696004803603604081101561023357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610732565b005b6102d76004803603606081101561028157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610750565b005b61031b600480360360208110156102ef57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610770565b6040518082815260200191505060405180910390f35b61037d6004803603604081101561034757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506107b9565b604051808215151515815260200191505060405180910390f35b61049a600480360360808110156103ad57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561041457600080fd5b82018360208201111561042657600080fd5b8035906020019184600183028401116401000000008311171561044857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506107d0565b005b6104fe600480360360408110156104b257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610843565b6040518082815260200191505060405180910390f35b6105f76004803603606081101561052a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561057157600080fd5b82018360208201111561058357600080fd5b803590602001918460018302840111640100000000831117156105a557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506108ca565b005b6000806000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff169050919050565b600061066d33848461093b565b6001905092915050565b6000600354905090565b600061068e848484610b32565b610727843361072285600260008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610dd290919063ffffffff16565b61093b565b600190509392505050565b61074c8282604051806020016040528060008152506108ca565b5050565b61076b838383604051806020016040528060008152506107d0565b505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60006107c6338484610b32565b6001905092915050565b6107db848484610681565b506107e884848484610e1c565b61083d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611182602e913960400191505060405180910390fd5b50505050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6108d483836107b9565b506108e133848484610e1c565b610936576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611182602e913960400191505060405180910390fd5b505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156109c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806111f66023913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a47576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806111616021913960400191505060405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610bb8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806111d26024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610c3e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806111b06022913960400191505060405180910390fd5b610c9081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610dd290919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610d2581600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461100590919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000610e1483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061108d565b905092915050565b6000610e3d8473ffffffffffffffffffffffffffffffffffffffff1661114d565b610e4a5760019050610ffd565b60008473ffffffffffffffffffffffffffffffffffffffff16639d188c22338887876040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610f25578082015181840152602081019050610f0a565b50505050905090810190601f168015610f525780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610f7457600080fd5b505af1158015610f88573d6000803e3d6000fd5b505050506040513d6020811015610f9e57600080fd5b81019080805190602001909291905050509050639d188c2260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149150505b949350505050565b600080828401905083811015611083576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600083831115829061113a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156110ff5780820151818401526020810190506110e4565b50505050905090810190601f16801561112c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600080823b90506000811191505091905056fe4b4950373a20617070726f766520746f20746865207a65726f20616464726573734b4950373a207472616e7366657220746f206e6f6e204b495037526563656976657220696d706c656d656e7465724b4950373a207472616e7366657220746f20746865207a65726f20616464726573734b4950373a207472616e736665722066726f6d20746865207a65726f20616464726573734b4950373a20617070726f76652066726f6d20746865207a65726f2061646472657373a165627a7a723058205d233e51511049470bafe5a395741c47585054c9caa8e343ec38d666aef9528d0029",
}

// Kip7ABI is the input ABI used to generate the binding from.
// Deprecated: Use Kip7MetaData.ABI instead.
var Kip7ABI = Kip7MetaData.ABI

// Kip7Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Kip7MetaData.Bin instead.
var Kip7Bin = Kip7MetaData.Bin

// DeployKip7 deploys a new Ethereum contract, binding an instance of Kip7 to it.
func DeployKip7(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Kip7, error) {
	parsed, err := Kip7MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Kip7Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kip7{Kip7Caller: Kip7Caller{contract: contract}, Kip7Transactor: Kip7Transactor{contract: contract}, Kip7Filterer: Kip7Filterer{contract: contract}}, nil
}

// Kip7 is an auto generated Go binding around an Ethereum contract.
type Kip7 struct {
	Kip7Caller     // Read-only binding to the contract
	Kip7Transactor // Write-only binding to the contract
	Kip7Filterer   // Log filterer for contract events
}

// Kip7Caller is an auto generated read-only Go binding around an Ethereum contract.
type Kip7Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Kip7Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Kip7Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Kip7Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Kip7Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Kip7Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Kip7Session struct {
	Contract     *Kip7             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Kip7CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Kip7CallerSession struct {
	Contract *Kip7Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Kip7TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Kip7TransactorSession struct {
	Contract     *Kip7Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Kip7Raw is an auto generated low-level Go binding around an Ethereum contract.
type Kip7Raw struct {
	Contract *Kip7 // Generic contract binding to access the raw methods on
}

// Kip7CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Kip7CallerRaw struct {
	Contract *Kip7Caller // Generic read-only contract binding to access the raw methods on
}

// Kip7TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Kip7TransactorRaw struct {
	Contract *Kip7Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKip7 creates a new instance of Kip7, bound to a specific deployed contract.
func NewKip7(address common.Address, backend bind.ContractBackend) (*Kip7, error) {
	contract, err := bindKip7(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kip7{Kip7Caller: Kip7Caller{contract: contract}, Kip7Transactor: Kip7Transactor{contract: contract}, Kip7Filterer: Kip7Filterer{contract: contract}}, nil
}

// NewKip7Caller creates a new read-only instance of Kip7, bound to a specific deployed contract.
func NewKip7Caller(address common.Address, caller bind.ContractCaller) (*Kip7Caller, error) {
	contract, err := bindKip7(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Kip7Caller{contract: contract}, nil
}

// NewKip7Transactor creates a new write-only instance of Kip7, bound to a specific deployed contract.
func NewKip7Transactor(address common.Address, transactor bind.ContractTransactor) (*Kip7Transactor, error) {
	contract, err := bindKip7(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Kip7Transactor{contract: contract}, nil
}

// NewKip7Filterer creates a new log filterer instance of Kip7, bound to a specific deployed contract.
func NewKip7Filterer(address common.Address, filterer bind.ContractFilterer) (*Kip7Filterer, error) {
	contract, err := bindKip7(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Kip7Filterer{contract: contract}, nil
}

// bindKip7 binds a generic wrapper to an already deployed contract.
func bindKip7(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Kip7ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kip7 *Kip7Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kip7.Contract.Kip7Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kip7 *Kip7Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kip7.Contract.Kip7Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kip7 *Kip7Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kip7.Contract.Kip7Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kip7 *Kip7CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kip7.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kip7 *Kip7TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kip7.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kip7 *Kip7TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kip7.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Kip7 *Kip7Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kip7.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Kip7 *Kip7Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Kip7.Contract.Allowance(&_Kip7.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Kip7 *Kip7CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Kip7.Contract.Allowance(&_Kip7.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Kip7 *Kip7Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kip7.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Kip7 *Kip7Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Kip7.Contract.BalanceOf(&_Kip7.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Kip7 *Kip7CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Kip7.Contract.BalanceOf(&_Kip7.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kip7 *Kip7Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Kip7.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kip7 *Kip7Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Kip7.Contract.SupportsInterface(&_Kip7.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kip7 *Kip7CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Kip7.Contract.SupportsInterface(&_Kip7.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kip7 *Kip7Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kip7.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kip7 *Kip7Session) TotalSupply() (*big.Int, error) {
	return _Kip7.Contract.TotalSupply(&_Kip7.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kip7 *Kip7CallerSession) TotalSupply() (*big.Int, error) {
	return _Kip7.Contract.TotalSupply(&_Kip7.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Kip7 *Kip7Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Kip7.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Kip7 *Kip7Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Kip7.Contract.Approve(&_Kip7.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Kip7 *Kip7TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Kip7.Contract.Approve(&_Kip7.TransactOpts, spender, value)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_Kip7 *Kip7Transactor) SafeTransfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.contract.Transact(opts, "safeTransfer", recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_Kip7 *Kip7Session) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.Contract.SafeTransfer(&_Kip7.TransactOpts, recipient, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x423f6cef.
//
// Solidity: function safeTransfer(address recipient, uint256 amount) returns()
func (_Kip7 *Kip7TransactorSession) SafeTransfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.Contract.SafeTransfer(&_Kip7.TransactOpts, recipient, amount)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_Kip7 *Kip7Transactor) SafeTransfer0(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Kip7.contract.Transact(opts, "safeTransfer0", recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_Kip7 *Kip7Session) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Kip7.Contract.SafeTransfer0(&_Kip7.TransactOpts, recipient, amount, data)
}

// SafeTransfer0 is a paid mutator transaction binding the contract method 0xeb795549.
//
// Solidity: function safeTransfer(address recipient, uint256 amount, bytes data) returns()
func (_Kip7 *Kip7TransactorSession) SafeTransfer0(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Kip7.Contract.SafeTransfer0(&_Kip7.TransactOpts, recipient, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_Kip7 *Kip7Transactor) SafeTransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.contract.Transact(opts, "safeTransferFrom", sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_Kip7 *Kip7Session) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.Contract.SafeTransferFrom(&_Kip7.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount) returns()
func (_Kip7 *Kip7TransactorSession) SafeTransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.Contract.SafeTransferFrom(&_Kip7.TransactOpts, sender, recipient, amount)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_Kip7 *Kip7Transactor) SafeTransferFrom0(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Kip7.contract.Transact(opts, "safeTransferFrom0", sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_Kip7 *Kip7Session) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Kip7.Contract.SafeTransferFrom0(&_Kip7.TransactOpts, sender, recipient, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address sender, address recipient, uint256 amount, bytes data) returns()
func (_Kip7 *Kip7TransactorSession) SafeTransferFrom0(sender common.Address, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Kip7.Contract.SafeTransferFrom0(&_Kip7.TransactOpts, sender, recipient, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Kip7 *Kip7Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Kip7 *Kip7Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.Contract.Transfer(&_Kip7.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Kip7 *Kip7TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.Contract.Transfer(&_Kip7.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Kip7 *Kip7Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Kip7 *Kip7Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.Contract.TransferFrom(&_Kip7.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Kip7 *Kip7TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kip7.Contract.TransferFrom(&_Kip7.TransactOpts, sender, recipient, amount)
}

// Kip7ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Kip7 contract.
type Kip7ApprovalIterator struct {
	Event *Kip7Approval // Event containing the contract specifics and raw log

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
func (it *Kip7ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Kip7Approval)
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
		it.Event = new(Kip7Approval)
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
func (it *Kip7ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Kip7ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Kip7Approval represents a Approval event raised by the Kip7 contract.
type Kip7Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Kip7 *Kip7Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Kip7ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Kip7.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Kip7ApprovalIterator{contract: _Kip7.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Kip7 *Kip7Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Kip7Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Kip7.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Kip7Approval)
				if err := _Kip7.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Kip7 *Kip7Filterer) ParseApproval(log types.Log) (*Kip7Approval, error) {
	event := new(Kip7Approval)
	if err := _Kip7.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Kip7TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Kip7 contract.
type Kip7TransferIterator struct {
	Event *Kip7Transfer // Event containing the contract specifics and raw log

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
func (it *Kip7TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Kip7Transfer)
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
		it.Event = new(Kip7Transfer)
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
func (it *Kip7TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Kip7TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Kip7Transfer represents a Transfer event raised by the Kip7 contract.
type Kip7Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Kip7 *Kip7Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Kip7TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Kip7.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Kip7TransferIterator{contract: _Kip7.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Kip7 *Kip7Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Kip7Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Kip7.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Kip7Transfer)
				if err := _Kip7.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Kip7 *Kip7Filterer) ParseTransfer(log types.Log) (*Kip7Transfer, error) {
	event := new(Kip7Transfer)
	if err := _Kip7.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

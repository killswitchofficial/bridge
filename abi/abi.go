// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ba23902febd09b9aec9c7bf92ff40e4246ff1b18f75e543afcea9fea08683e5c64736f6c63430008040033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// BridgeBaseABI is the input ABI used to generate the binding from.
const BridgeBaseABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"contractIFee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimiter\",\"outputs\":[{\"internalType\":\"contractILimiter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimiterUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"isLimited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isUnlockCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFee\",\"name\":\"fee_\",\"type\":\"address\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILimiter\",\"name\":\"limiter\",\"type\":\"address\"}],\"name\":\"setLimiter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// BridgeBaseFuncSigs maps the 4-byte function signature to its string representation.
var BridgeBaseFuncSigs = map[string]string{
	"99a5d747": "calculateFee(uint256)",
	"ced72f87": "getFee()",
	"eb2a0d1f": "getLimiter()",
	"7eb76b29": "getLimiterUsage()",
	"08a90d5a": "isLimited(uint256)",
	"a4d7fa93": "isUnlockCompleted(bytes32)",
	"dd467064": "lock(uint256)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"7917fb9f": "setFee(address)",
	"7a29084c": "setLimiter(address)",
	"f2fde38b": "transferOwnership(address)",
	"b322edea": "unlock(address,uint256,bytes32)",
	"3f4ba83a": "unpause()",
}

// BridgeBase is an auto generated Go binding around an Ethereum contract.
type BridgeBase struct {
	BridgeBaseCaller     // Read-only binding to the contract
	BridgeBaseTransactor // Write-only binding to the contract
	BridgeBaseFilterer   // Log filterer for contract events
}

// BridgeBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBaseSession struct {
	Contract     *BridgeBase       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBaseCallerSession struct {
	Contract *BridgeBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBaseTransactorSession struct {
	Contract     *BridgeBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBaseRaw struct {
	Contract *BridgeBase // Generic contract binding to access the raw methods on
}

// BridgeBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBaseCallerRaw struct {
	Contract *BridgeBaseCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBaseTransactorRaw struct {
	Contract *BridgeBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBase creates a new instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBase(address common.Address, backend bind.ContractBackend) (*BridgeBase, error) {
	contract, err := bindBridgeBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBase{BridgeBaseCaller: BridgeBaseCaller{contract: contract}, BridgeBaseTransactor: BridgeBaseTransactor{contract: contract}, BridgeBaseFilterer: BridgeBaseFilterer{contract: contract}}, nil
}

// NewBridgeBaseCaller creates a new read-only instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBaseCaller(address common.Address, caller bind.ContractCaller) (*BridgeBaseCaller, error) {
	contract, err := bindBridgeBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseCaller{contract: contract}, nil
}

// NewBridgeBaseTransactor creates a new write-only instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBaseTransactor, error) {
	contract, err := bindBridgeBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseTransactor{contract: contract}, nil
}

// NewBridgeBaseFilterer creates a new log filterer instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBaseFilterer, error) {
	contract, err := bindBridgeBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseFilterer{contract: contract}, nil
}

// bindBridgeBase binds a generic wrapper to an already deployed contract.
func bindBridgeBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBase *BridgeBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBase.Contract.BridgeBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBase *BridgeBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.Contract.BridgeBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBase *BridgeBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBase.Contract.BridgeBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBase *BridgeBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBase *BridgeBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBase *BridgeBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBase.Contract.contract.Transact(opts, method, params...)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) CalculateFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "calculateFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeBase *BridgeBaseSession) CalculateFee(amount *big.Int) (*big.Int, error) {
	return _BridgeBase.Contract.CalculateFee(&_BridgeBase.CallOpts, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) CalculateFee(amount *big.Int) (*big.Int, error) {
	return _BridgeBase.Contract.CalculateFee(&_BridgeBase.CallOpts, amount)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeBase *BridgeBaseCaller) GetFee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeBase *BridgeBaseSession) GetFee() (common.Address, error) {
	return _BridgeBase.Contract.GetFee(&_BridgeBase.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeBase *BridgeBaseCallerSession) GetFee() (common.Address, error) {
	return _BridgeBase.Contract.GetFee(&_BridgeBase.CallOpts)
}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeBase *BridgeBaseCaller) GetLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "getLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeBase *BridgeBaseSession) GetLimiter() (common.Address, error) {
	return _BridgeBase.Contract.GetLimiter(&_BridgeBase.CallOpts)
}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeBase *BridgeBaseCallerSession) GetLimiter() (common.Address, error) {
	return _BridgeBase.Contract.GetLimiter(&_BridgeBase.CallOpts)
}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) GetLimiterUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "getLimiterUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeBase *BridgeBaseSession) GetLimiterUsage() (*big.Int, error) {
	return _BridgeBase.Contract.GetLimiterUsage(&_BridgeBase.CallOpts)
}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) GetLimiterUsage() (*big.Int, error) {
	return _BridgeBase.Contract.GetLimiterUsage(&_BridgeBase.CallOpts)
}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeBase *BridgeBaseCaller) IsLimited(opts *bind.CallOpts, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "isLimited", amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeBase *BridgeBaseSession) IsLimited(amount *big.Int) (bool, error) {
	return _BridgeBase.Contract.IsLimited(&_BridgeBase.CallOpts, amount)
}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeBase *BridgeBaseCallerSession) IsLimited(amount *big.Int) (bool, error) {
	return _BridgeBase.Contract.IsLimited(&_BridgeBase.CallOpts, amount)
}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeBase *BridgeBaseCaller) IsUnlockCompleted(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "isUnlockCompleted", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeBase *BridgeBaseSession) IsUnlockCompleted(hash [32]byte) (bool, error) {
	return _BridgeBase.Contract.IsUnlockCompleted(&_BridgeBase.CallOpts, hash)
}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeBase *BridgeBaseCallerSession) IsUnlockCompleted(hash [32]byte) (bool, error) {
	return _BridgeBase.Contract.IsUnlockCompleted(&_BridgeBase.CallOpts, hash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeBase *BridgeBaseCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeBase *BridgeBaseSession) Name() (string, error) {
	return _BridgeBase.Contract.Name(&_BridgeBase.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeBase *BridgeBaseCallerSession) Name() (string, error) {
	return _BridgeBase.Contract.Name(&_BridgeBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBase *BridgeBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBase *BridgeBaseSession) Owner() (common.Address, error) {
	return _BridgeBase.Contract.Owner(&_BridgeBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBase *BridgeBaseCallerSession) Owner() (common.Address, error) {
	return _BridgeBase.Contract.Owner(&_BridgeBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBase *BridgeBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBase *BridgeBaseSession) Paused() (bool, error) {
	return _BridgeBase.Contract.Paused(&_BridgeBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBase *BridgeBaseCallerSession) Paused() (bool, error) {
	return _BridgeBase.Contract.Paused(&_BridgeBase.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeBase *BridgeBaseTransactor) Lock(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "lock", amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeBase *BridgeBaseSession) Lock(amount *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.Lock(&_BridgeBase.TransactOpts, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeBase *BridgeBaseTransactorSession) Lock(amount *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.Lock(&_BridgeBase.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBase *BridgeBaseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBase *BridgeBaseSession) Pause() (*types.Transaction, error) {
	return _BridgeBase.Contract.Pause(&_BridgeBase.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBase *BridgeBaseTransactorSession) Pause() (*types.Transaction, error) {
	return _BridgeBase.Contract.Pause(&_BridgeBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBase *BridgeBaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBase *BridgeBaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeBase.Contract.RenounceOwnership(&_BridgeBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBase *BridgeBaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeBase.Contract.RenounceOwnership(&_BridgeBase.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeBase *BridgeBaseTransactor) SetFee(opts *bind.TransactOpts, fee_ common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "setFee", fee_)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeBase *BridgeBaseSession) SetFee(fee_ common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.SetFee(&_BridgeBase.TransactOpts, fee_)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeBase *BridgeBaseTransactorSession) SetFee(fee_ common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.SetFee(&_BridgeBase.TransactOpts, fee_)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeBase *BridgeBaseTransactor) SetLimiter(opts *bind.TransactOpts, limiter common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "setLimiter", limiter)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeBase *BridgeBaseSession) SetLimiter(limiter common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.SetLimiter(&_BridgeBase.TransactOpts, limiter)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeBase *BridgeBaseTransactorSession) SetLimiter(limiter common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.SetLimiter(&_BridgeBase.TransactOpts, limiter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBase *BridgeBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBase *BridgeBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.TransferOwnership(&_BridgeBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBase *BridgeBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.TransferOwnership(&_BridgeBase.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeBase *BridgeBaseTransactor) Unlock(opts *bind.TransactOpts, account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "unlock", account, amount, hash)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeBase *BridgeBaseSession) Unlock(account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeBase.Contract.Unlock(&_BridgeBase.TransactOpts, account, amount, hash)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeBase *BridgeBaseTransactorSession) Unlock(account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeBase.Contract.Unlock(&_BridgeBase.TransactOpts, account, amount, hash)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBase *BridgeBaseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBase *BridgeBaseSession) Unpause() (*types.Transaction, error) {
	return _BridgeBase.Contract.Unpause(&_BridgeBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBase *BridgeBaseTransactorSession) Unpause() (*types.Transaction, error) {
	return _BridgeBase.Contract.Unpause(&_BridgeBase.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBase *BridgeBaseTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBase *BridgeBaseSession) Receive() (*types.Transaction, error) {
	return _BridgeBase.Contract.Receive(&_BridgeBase.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBase *BridgeBaseTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeBase.Contract.Receive(&_BridgeBase.TransactOpts)
}

// BridgeBaseLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the BridgeBase contract.
type BridgeBaseLockedIterator struct {
	Event *BridgeBaseLocked // Event containing the contract specifics and raw log

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
func (it *BridgeBaseLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseLocked)
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
		it.Event = new(BridgeBaseLocked)
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
func (it *BridgeBaseLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseLocked represents a Locked event raised by the BridgeBase contract.
type BridgeBaseLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeBase *BridgeBaseFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*BridgeBaseLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseLockedIterator{contract: _BridgeBase.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeBase *BridgeBaseFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *BridgeBaseLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseLocked)
				if err := _BridgeBase.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeBase *BridgeBaseFilterer) ParseLocked(log types.Log) (*BridgeBaseLocked, error) {
	event := new(BridgeBaseLocked)
	if err := _BridgeBase.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeBase contract.
type BridgeBaseOwnershipTransferredIterator struct {
	Event *BridgeBaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseOwnershipTransferred)
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
		it.Event = new(BridgeBaseOwnershipTransferred)
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
func (it *BridgeBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeBase contract.
type BridgeBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBase *BridgeBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseOwnershipTransferredIterator{contract: _BridgeBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBase *BridgeBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseOwnershipTransferred)
				if err := _BridgeBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBase *BridgeBaseFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeBaseOwnershipTransferred, error) {
	event := new(BridgeBaseOwnershipTransferred)
	if err := _BridgeBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBasePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BridgeBase contract.
type BridgeBasePausedIterator struct {
	Event *BridgeBasePaused // Event containing the contract specifics and raw log

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
func (it *BridgeBasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBasePaused)
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
		it.Event = new(BridgeBasePaused)
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
func (it *BridgeBasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBasePaused represents a Paused event raised by the BridgeBase contract.
type BridgeBasePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeBase *BridgeBaseFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgeBasePausedIterator, error) {

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgeBasePausedIterator{contract: _BridgeBase.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeBase *BridgeBaseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgeBasePaused) (event.Subscription, error) {

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBasePaused)
				if err := _BridgeBase.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address account)
func (_BridgeBase *BridgeBaseFilterer) ParsePaused(log types.Log) (*BridgeBasePaused, error) {
	event := new(BridgeBasePaused)
	if err := _BridgeBase.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the BridgeBase contract.
type BridgeBaseUnlockedIterator struct {
	Event *BridgeBaseUnlocked // Event containing the contract specifics and raw log

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
func (it *BridgeBaseUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseUnlocked)
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
		it.Event = new(BridgeBaseUnlocked)
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
func (it *BridgeBaseUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseUnlocked represents a Unlocked event raised by the BridgeBase contract.
type BridgeBaseUnlocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeBase *BridgeBaseFilterer) FilterUnlocked(opts *bind.FilterOpts, sender []common.Address) (*BridgeBaseUnlockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Unlocked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseUnlockedIterator{contract: _BridgeBase.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeBase *BridgeBaseFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *BridgeBaseUnlocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Unlocked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseUnlocked)
				if err := _BridgeBase.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeBase *BridgeBaseFilterer) ParseUnlocked(log types.Log) (*BridgeBaseUnlocked, error) {
	event := new(BridgeBaseUnlocked)
	if err := _BridgeBase.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BridgeBase contract.
type BridgeBaseUnpausedIterator struct {
	Event *BridgeBaseUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeBaseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseUnpaused)
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
		it.Event = new(BridgeBaseUnpaused)
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
func (it *BridgeBaseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseUnpaused represents a Unpaused event raised by the BridgeBase contract.
type BridgeBaseUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeBase *BridgeBaseFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeBaseUnpausedIterator, error) {

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeBaseUnpausedIterator{contract: _BridgeBase.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeBase *BridgeBaseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeBaseUnpaused) (event.Subscription, error) {

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseUnpaused)
				if err := _BridgeBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address account)
func (_BridgeBase *BridgeBaseFilterer) ParseUnpaused(log types.Log) (*BridgeBaseUnpaused, error) {
	event := new(BridgeBaseUnpaused)
	if err := _BridgeBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBurnerABI is the input ABI used to generate the binding from.
const BridgeBurnerABI = "[{\"inputs\":[{\"internalType\":\"contractIWrappedToken\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"contractIFee\",\"name\":\"fee\",\"type\":\"address\"},{\"internalType\":\"contractILimiter\",\"name\":\"limiter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"contractIFee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimiter\",\"outputs\":[{\"internalType\":\"contractILimiter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimiterUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"isLimited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isUnlockCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFee\",\"name\":\"fee_\",\"type\":\"address\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILimiter\",\"name\":\"limiter\",\"type\":\"address\"}],\"name\":\"setLimiter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIWrappedToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// BridgeBurnerFuncSigs maps the 4-byte function signature to its string representation.
var BridgeBurnerFuncSigs = map[string]string{
	"99a5d747": "calculateFee(uint256)",
	"ced72f87": "getFee()",
	"eb2a0d1f": "getLimiter()",
	"7eb76b29": "getLimiterUsage()",
	"08a90d5a": "isLimited(uint256)",
	"a4d7fa93": "isUnlockCompleted(bytes32)",
	"dd467064": "lock(uint256)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"7917fb9f": "setFee(address)",
	"7a29084c": "setLimiter(address)",
	"fc0c546a": "token()",
	"f2fde38b": "transferOwnership(address)",
	"b322edea": "unlock(address,uint256,bytes32)",
	"3f4ba83a": "unpause()",
}

// BridgeBurnerBin is the compiled bytecode used for deploying new contracts.
var BridgeBurnerBin = "0x60806040523480156200001157600080fd5b50604051620012b8380380620012b88339810160408190526200003491620001a6565b600080546001600160a01b031916339081178255604051859285928592909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19169055600180558251620000a1906002906020860190620000ee565b50600380546001600160a01b039384166001600160a01b03199182161790915560048054928416928216929092179091556006805497909216961695909517909455506200032692505050565b828054620000fc90620002ba565b90600052602060002090601f0160209004810192826200012057600085556200016b565b82601f106200013b57805160ff19168380011785556200016b565b828001600101855582156200016b579182015b828111156200016b5782518255916020019190600101906200014e565b50620001799291506200017d565b5090565b5b808211156200017957600081556001016200017e565b8051620001a1816200030d565b919050565b60008060008060808587031215620001bc578384fd5b8451620001c9816200030d565b602086810151919550906001600160401b0380821115620001e8578586fd5b818801915088601f830112620001fc578586fd5b815181811115620002115762000211620002f7565b604051601f8201601f19908116603f011681019083821181831017156200023c576200023c620002f7565b816040528281528b8684870101111562000254578889fd5b8893505b8284101562000277578484018601518185018701529285019262000258565b828411156200028857888684830101525b8098505050505050506200029f6040860162000194565b9150620002af6060860162000194565b905092959194509250565b600181811c90821680620002cf57607f821691505b60208210811415620002f157634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146200032357600080fd5b50565b610f8280620003366000396000f3fe60806040526004361061010d5760003560e01c80638da5cb5b11610095578063ced72f8711610064578063ced72f87146102dc578063dd467064146102fa578063eb2a0d1f1461030d578063f2fde38b1461032b578063fc0c546a1461034b57600080fd5b80638da5cb5b1461023a57806399a5d7471461026c578063a4d7fa931461028c578063b322edea146102bc57600080fd5b8063715018a6116100dc578063715018a6146101ad5780637917fb9f146101c25780637a29084c146101e25780637eb76b29146102025780638456cb591461022557600080fd5b806306fdde031461011c57806308a90d5a146101475780633f4ba83a146101775780635c975abb1461018e57600080fd5b3661011757600080fd5b600080fd5b34801561012857600080fd5b50610131610369565b60405161013e9190610e74565b60405180910390f35b34801561015357600080fd5b50610167610162366004610e44565b6103fb565b604051901515815260200161013e565b34801561018357600080fd5b5061018c610489565b005b34801561019a57600080fd5b50600054600160a01b900460ff16610167565b3480156101b957600080fd5b5061018c6104c6565b3480156101ce57600080fd5b5061018c6101dd366004610dcd565b610500565b3480156101ee57600080fd5b5061018c6101fd366004610dcd565b61054c565b34801561020e57600080fd5b50610217610598565b60405190815260200161013e565b34801561023157600080fd5b5061018c61061e565b34801561024657600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200161013e565b34801561027857600080fd5b50610217610287366004610e44565b610650565b34801561029857600080fd5b506101676102a7366004610e44565b60009081526005602052604090205460ff1690565b3480156102c857600080fd5b5061018c6102d7366004610df0565b6106e7565b3480156102e857600080fd5b506003546001600160a01b0316610254565b61018c610308366004610e44565b6107c8565b34801561031957600080fd5b506004546001600160a01b0316610254565b34801561033757600080fd5b5061018c610346366004610dcd565b61088f565b34801561035757600080fd5b506006546001600160a01b0316610254565b60606002805461037890610efc565b80601f01602080910402602001604051908101604052809291908181526020018280546103a490610efc565b80156103f15780601f106103c6576101008083540402835291602001916103f1565b820191906000526020600020905b8154815290600101906020018083116103d457829003601f168201915b5050505050905090565b600480546040516323b0c65960e11b81523092810192909252602482018390526000916001600160a01b03909116906347618cb29060440160206040518083038186803b15801561044b57600080fd5b505afa15801561045f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104839190610e24565b92915050565b6000546001600160a01b031633146104bc5760405162461bcd60e51b81526004016104b390610ec7565b60405180910390fd5b6104c4610979565b565b6000546001600160a01b031633146104f05760405162461bcd60e51b81526004016104b390610ec7565b6104f8610a16565b6104c4610a9e565b6000546001600160a01b0316331461052a5760405162461bcd60e51b81526004016104b390610ec7565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146105765760405162461bcd60e51b81526004016104b390610ec7565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b60048054604051632cdcd8af60e11b815230928101929092526000916001600160a01b03909116906359b9b15e9060240160206040518083038186803b1580156105e157600080fd5b505afa1580156105f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106199190610e5c565b905090565b6000546001600160a01b031633146106485760405162461bcd60e51b81526004016104b390610ec7565b6104c4610a16565b6003546000906001600160a01b031661066b57506000919050565b60035460405163173b25bd60e31b8152600481018490526001600160a01b039091169063b9d92de89060240160206040518083038186803b1580156106af57600080fd5b505afa1580156106c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104839190610e5c565b6000546001600160a01b031633146107115760405162461bcd60e51b81526004016104b390610ec7565b61071a81610b12565b6006546040516340c10f1960e01b81526001600160a01b03858116600483015260248201859052909116906340c10f1990604401600060405180830381600087803b15801561076857600080fd5b505af115801561077c573d6000803e3d6000fd5b50505050826001600160a01b03167f0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d836040516107bb91815260200190565b60405180910390a2505050565b6107d181610b8c565b6006546001600160a01b03166379cc6790336040516001600160e01b031960e084901b1681526001600160a01b03909116600482015260248101849052604401600060405180830381600087803b15801561082b57600080fd5b505af115801561083f573d6000803e3d6000fd5b5050505061084a3390565b6001600160a01b03167f9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd600088260405161088491815260200190565b60405180910390a250565b6000546001600160a01b031633146108b95760405162461bcd60e51b81526004016104b390610ec7565b6001600160a01b03811661091e5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104b3565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600054600160a01b900460ff166109c95760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016104b3565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600054600160a01b900460ff1615610a635760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016104b3565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586109f93390565b6000546001600160a01b03163314610ac85760405162461bcd60e51b81526004016104b390610ec7565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008181526005602052604090205460ff1615610b715760405162461bcd60e51b815260206004820152601c60248201527f427269646765426173653a20616c726561647920756e6c6f636b65640000000060448201526064016104b3565b6000908152600560205260409020805460ff19166001179055565b600054600160a01b900460ff1615610bd95760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016104b3565b610be281610bee565b610beb81610c60565b50565b6004546001600160a01b0316610c015750565b6004805460405163606ecf2960e11b81529182018390526001600160a01b03169063c0dd9e5290602401600060405180830381600087803b158015610c4557600080fd5b505af1158015610c59573d6000803e3d6000fd5b5050505050565b60026001541415610cb35760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016104b3565b60026001556000610cc382610650565b905080610cd05750610dc6565b80341015610d205760405162461bcd60e51b815260206004820152601a60248201527f427269646765426173653a206e6f7420656e6f7567682066656500000000000060448201526064016104b3565b600080546040516001600160a01b039091169034908381818185875af1925050503d8060008114610d6d576040519150601f19603f3d011682016040523d82523d6000602084013e610d72565b606091505b5050905080610dc35760405162461bcd60e51b815260206004820181905260248201527f427269646765426173653a2063616e206e6f74207472616e736665722066656560448201526064016104b3565b50505b5060018055565b600060208284031215610dde578081fd5b8135610de981610f37565b9392505050565b600080600060608486031215610e04578182fd5b8335610e0f81610f37565b95602085013595506040909401359392505050565b600060208284031215610e35578081fd5b81518015158114610de9578182fd5b600060208284031215610e55578081fd5b5035919050565b600060208284031215610e6d578081fd5b5051919050565b6000602080835283518082850152825b81811015610ea057858101830151858201604001528201610e84565b81811115610eb15783604083870101525b50601f01601f1916929092016040019392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b600181811c90821680610f1057607f821691505b60208210811415610f3157634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b0381168114610beb57600080fdfea264697066735822122062ed4d3316cc903c9b783518ca42b9437f425437d1f5d6f16184b965b96f735564736f6c63430008040033"

// DeployBridgeBurner deploys a new Ethereum contract, binding an instance of BridgeBurner to it.
func DeployBridgeBurner(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address, name string, fee common.Address, limiter common.Address) (common.Address, *types.Transaction, *BridgeBurner, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeBurnerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBurnerBin), backend, token_, name, fee, limiter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeBurner{BridgeBurnerCaller: BridgeBurnerCaller{contract: contract}, BridgeBurnerTransactor: BridgeBurnerTransactor{contract: contract}, BridgeBurnerFilterer: BridgeBurnerFilterer{contract: contract}}, nil
}

// BridgeBurner is an auto generated Go binding around an Ethereum contract.
type BridgeBurner struct {
	BridgeBurnerCaller     // Read-only binding to the contract
	BridgeBurnerTransactor // Write-only binding to the contract
	BridgeBurnerFilterer   // Log filterer for contract events
}

// BridgeBurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBurnerSession struct {
	Contract     *BridgeBurner     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBurnerCallerSession struct {
	Contract *BridgeBurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BridgeBurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBurnerTransactorSession struct {
	Contract     *BridgeBurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeBurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBurnerRaw struct {
	Contract *BridgeBurner // Generic contract binding to access the raw methods on
}

// BridgeBurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBurnerCallerRaw struct {
	Contract *BridgeBurnerCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBurnerTransactorRaw struct {
	Contract *BridgeBurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBurner creates a new instance of BridgeBurner, bound to a specific deployed contract.
func NewBridgeBurner(address common.Address, backend bind.ContractBackend) (*BridgeBurner, error) {
	contract, err := bindBridgeBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBurner{BridgeBurnerCaller: BridgeBurnerCaller{contract: contract}, BridgeBurnerTransactor: BridgeBurnerTransactor{contract: contract}, BridgeBurnerFilterer: BridgeBurnerFilterer{contract: contract}}, nil
}

// NewBridgeBurnerCaller creates a new read-only instance of BridgeBurner, bound to a specific deployed contract.
func NewBridgeBurnerCaller(address common.Address, caller bind.ContractCaller) (*BridgeBurnerCaller, error) {
	contract, err := bindBridgeBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBurnerCaller{contract: contract}, nil
}

// NewBridgeBurnerTransactor creates a new write-only instance of BridgeBurner, bound to a specific deployed contract.
func NewBridgeBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBurnerTransactor, error) {
	contract, err := bindBridgeBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBurnerTransactor{contract: contract}, nil
}

// NewBridgeBurnerFilterer creates a new log filterer instance of BridgeBurner, bound to a specific deployed contract.
func NewBridgeBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBurnerFilterer, error) {
	contract, err := bindBridgeBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBurnerFilterer{contract: contract}, nil
}

// bindBridgeBurner binds a generic wrapper to an already deployed contract.
func bindBridgeBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeBurnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBurner *BridgeBurnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBurner.Contract.BridgeBurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBurner *BridgeBurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBurner.Contract.BridgeBurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBurner *BridgeBurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBurner.Contract.BridgeBurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBurner *BridgeBurnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBurner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBurner *BridgeBurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBurner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBurner *BridgeBurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBurner.Contract.contract.Transact(opts, method, params...)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeBurner *BridgeBurnerCaller) CalculateFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBurner.contract.Call(opts, &out, "calculateFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeBurner *BridgeBurnerSession) CalculateFee(amount *big.Int) (*big.Int, error) {
	return _BridgeBurner.Contract.CalculateFee(&_BridgeBurner.CallOpts, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeBurner *BridgeBurnerCallerSession) CalculateFee(amount *big.Int) (*big.Int, error) {
	return _BridgeBurner.Contract.CalculateFee(&_BridgeBurner.CallOpts, amount)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeBurner *BridgeBurnerCaller) GetFee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBurner.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeBurner *BridgeBurnerSession) GetFee() (common.Address, error) {
	return _BridgeBurner.Contract.GetFee(&_BridgeBurner.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeBurner *BridgeBurnerCallerSession) GetFee() (common.Address, error) {
	return _BridgeBurner.Contract.GetFee(&_BridgeBurner.CallOpts)
}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeBurner *BridgeBurnerCaller) GetLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBurner.contract.Call(opts, &out, "getLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeBurner *BridgeBurnerSession) GetLimiter() (common.Address, error) {
	return _BridgeBurner.Contract.GetLimiter(&_BridgeBurner.CallOpts)
}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeBurner *BridgeBurnerCallerSession) GetLimiter() (common.Address, error) {
	return _BridgeBurner.Contract.GetLimiter(&_BridgeBurner.CallOpts)
}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeBurner *BridgeBurnerCaller) GetLimiterUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBurner.contract.Call(opts, &out, "getLimiterUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeBurner *BridgeBurnerSession) GetLimiterUsage() (*big.Int, error) {
	return _BridgeBurner.Contract.GetLimiterUsage(&_BridgeBurner.CallOpts)
}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeBurner *BridgeBurnerCallerSession) GetLimiterUsage() (*big.Int, error) {
	return _BridgeBurner.Contract.GetLimiterUsage(&_BridgeBurner.CallOpts)
}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeBurner *BridgeBurnerCaller) IsLimited(opts *bind.CallOpts, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _BridgeBurner.contract.Call(opts, &out, "isLimited", amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeBurner *BridgeBurnerSession) IsLimited(amount *big.Int) (bool, error) {
	return _BridgeBurner.Contract.IsLimited(&_BridgeBurner.CallOpts, amount)
}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeBurner *BridgeBurnerCallerSession) IsLimited(amount *big.Int) (bool, error) {
	return _BridgeBurner.Contract.IsLimited(&_BridgeBurner.CallOpts, amount)
}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeBurner *BridgeBurnerCaller) IsUnlockCompleted(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _BridgeBurner.contract.Call(opts, &out, "isUnlockCompleted", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeBurner *BridgeBurnerSession) IsUnlockCompleted(hash [32]byte) (bool, error) {
	return _BridgeBurner.Contract.IsUnlockCompleted(&_BridgeBurner.CallOpts, hash)
}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeBurner *BridgeBurnerCallerSession) IsUnlockCompleted(hash [32]byte) (bool, error) {
	return _BridgeBurner.Contract.IsUnlockCompleted(&_BridgeBurner.CallOpts, hash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeBurner *BridgeBurnerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeBurner.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeBurner *BridgeBurnerSession) Name() (string, error) {
	return _BridgeBurner.Contract.Name(&_BridgeBurner.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeBurner *BridgeBurnerCallerSession) Name() (string, error) {
	return _BridgeBurner.Contract.Name(&_BridgeBurner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBurner *BridgeBurnerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBurner.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBurner *BridgeBurnerSession) Owner() (common.Address, error) {
	return _BridgeBurner.Contract.Owner(&_BridgeBurner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBurner *BridgeBurnerCallerSession) Owner() (common.Address, error) {
	return _BridgeBurner.Contract.Owner(&_BridgeBurner.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBurner *BridgeBurnerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeBurner.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBurner *BridgeBurnerSession) Paused() (bool, error) {
	return _BridgeBurner.Contract.Paused(&_BridgeBurner.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBurner *BridgeBurnerCallerSession) Paused() (bool, error) {
	return _BridgeBurner.Contract.Paused(&_BridgeBurner.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeBurner *BridgeBurnerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBurner.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeBurner *BridgeBurnerSession) Token() (common.Address, error) {
	return _BridgeBurner.Contract.Token(&_BridgeBurner.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeBurner *BridgeBurnerCallerSession) Token() (common.Address, error) {
	return _BridgeBurner.Contract.Token(&_BridgeBurner.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeBurner *BridgeBurnerTransactor) Lock(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBurner.contract.Transact(opts, "lock", amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeBurner *BridgeBurnerSession) Lock(amount *big.Int) (*types.Transaction, error) {
	return _BridgeBurner.Contract.Lock(&_BridgeBurner.TransactOpts, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeBurner *BridgeBurnerTransactorSession) Lock(amount *big.Int) (*types.Transaction, error) {
	return _BridgeBurner.Contract.Lock(&_BridgeBurner.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBurner *BridgeBurnerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBurner.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBurner *BridgeBurnerSession) Pause() (*types.Transaction, error) {
	return _BridgeBurner.Contract.Pause(&_BridgeBurner.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBurner *BridgeBurnerTransactorSession) Pause() (*types.Transaction, error) {
	return _BridgeBurner.Contract.Pause(&_BridgeBurner.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBurner *BridgeBurnerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBurner.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBurner *BridgeBurnerSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeBurner.Contract.RenounceOwnership(&_BridgeBurner.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBurner *BridgeBurnerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeBurner.Contract.RenounceOwnership(&_BridgeBurner.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeBurner *BridgeBurnerTransactor) SetFee(opts *bind.TransactOpts, fee_ common.Address) (*types.Transaction, error) {
	return _BridgeBurner.contract.Transact(opts, "setFee", fee_)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeBurner *BridgeBurnerSession) SetFee(fee_ common.Address) (*types.Transaction, error) {
	return _BridgeBurner.Contract.SetFee(&_BridgeBurner.TransactOpts, fee_)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeBurner *BridgeBurnerTransactorSession) SetFee(fee_ common.Address) (*types.Transaction, error) {
	return _BridgeBurner.Contract.SetFee(&_BridgeBurner.TransactOpts, fee_)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeBurner *BridgeBurnerTransactor) SetLimiter(opts *bind.TransactOpts, limiter common.Address) (*types.Transaction, error) {
	return _BridgeBurner.contract.Transact(opts, "setLimiter", limiter)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeBurner *BridgeBurnerSession) SetLimiter(limiter common.Address) (*types.Transaction, error) {
	return _BridgeBurner.Contract.SetLimiter(&_BridgeBurner.TransactOpts, limiter)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeBurner *BridgeBurnerTransactorSession) SetLimiter(limiter common.Address) (*types.Transaction, error) {
	return _BridgeBurner.Contract.SetLimiter(&_BridgeBurner.TransactOpts, limiter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBurner *BridgeBurnerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBurner.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBurner *BridgeBurnerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBurner.Contract.TransferOwnership(&_BridgeBurner.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBurner *BridgeBurnerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBurner.Contract.TransferOwnership(&_BridgeBurner.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeBurner *BridgeBurnerTransactor) Unlock(opts *bind.TransactOpts, account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeBurner.contract.Transact(opts, "unlock", account, amount, hash)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeBurner *BridgeBurnerSession) Unlock(account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeBurner.Contract.Unlock(&_BridgeBurner.TransactOpts, account, amount, hash)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeBurner *BridgeBurnerTransactorSession) Unlock(account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeBurner.Contract.Unlock(&_BridgeBurner.TransactOpts, account, amount, hash)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBurner *BridgeBurnerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBurner.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBurner *BridgeBurnerSession) Unpause() (*types.Transaction, error) {
	return _BridgeBurner.Contract.Unpause(&_BridgeBurner.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBurner *BridgeBurnerTransactorSession) Unpause() (*types.Transaction, error) {
	return _BridgeBurner.Contract.Unpause(&_BridgeBurner.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBurner *BridgeBurnerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBurner.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBurner *BridgeBurnerSession) Receive() (*types.Transaction, error) {
	return _BridgeBurner.Contract.Receive(&_BridgeBurner.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBurner *BridgeBurnerTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeBurner.Contract.Receive(&_BridgeBurner.TransactOpts)
}

// BridgeBurnerLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the BridgeBurner contract.
type BridgeBurnerLockedIterator struct {
	Event *BridgeBurnerLocked // Event containing the contract specifics and raw log

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
func (it *BridgeBurnerLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBurnerLocked)
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
		it.Event = new(BridgeBurnerLocked)
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
func (it *BridgeBurnerLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBurnerLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBurnerLocked represents a Locked event raised by the BridgeBurner contract.
type BridgeBurnerLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeBurner *BridgeBurnerFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*BridgeBurnerLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeBurner.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBurnerLockedIterator{contract: _BridgeBurner.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeBurner *BridgeBurnerFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *BridgeBurnerLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeBurner.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBurnerLocked)
				if err := _BridgeBurner.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeBurner *BridgeBurnerFilterer) ParseLocked(log types.Log) (*BridgeBurnerLocked, error) {
	event := new(BridgeBurnerLocked)
	if err := _BridgeBurner.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBurnerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeBurner contract.
type BridgeBurnerOwnershipTransferredIterator struct {
	Event *BridgeBurnerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeBurnerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBurnerOwnershipTransferred)
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
		it.Event = new(BridgeBurnerOwnershipTransferred)
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
func (it *BridgeBurnerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBurnerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBurnerOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeBurner contract.
type BridgeBurnerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBurner *BridgeBurnerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeBurnerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeBurner.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBurnerOwnershipTransferredIterator{contract: _BridgeBurner.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBurner *BridgeBurnerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeBurnerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeBurner.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBurnerOwnershipTransferred)
				if err := _BridgeBurner.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBurner *BridgeBurnerFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeBurnerOwnershipTransferred, error) {
	event := new(BridgeBurnerOwnershipTransferred)
	if err := _BridgeBurner.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBurnerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BridgeBurner contract.
type BridgeBurnerPausedIterator struct {
	Event *BridgeBurnerPaused // Event containing the contract specifics and raw log

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
func (it *BridgeBurnerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBurnerPaused)
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
		it.Event = new(BridgeBurnerPaused)
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
func (it *BridgeBurnerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBurnerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBurnerPaused represents a Paused event raised by the BridgeBurner contract.
type BridgeBurnerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeBurner *BridgeBurnerFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgeBurnerPausedIterator, error) {

	logs, sub, err := _BridgeBurner.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgeBurnerPausedIterator{contract: _BridgeBurner.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeBurner *BridgeBurnerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgeBurnerPaused) (event.Subscription, error) {

	logs, sub, err := _BridgeBurner.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBurnerPaused)
				if err := _BridgeBurner.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address account)
func (_BridgeBurner *BridgeBurnerFilterer) ParsePaused(log types.Log) (*BridgeBurnerPaused, error) {
	event := new(BridgeBurnerPaused)
	if err := _BridgeBurner.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBurnerUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the BridgeBurner contract.
type BridgeBurnerUnlockedIterator struct {
	Event *BridgeBurnerUnlocked // Event containing the contract specifics and raw log

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
func (it *BridgeBurnerUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBurnerUnlocked)
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
		it.Event = new(BridgeBurnerUnlocked)
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
func (it *BridgeBurnerUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBurnerUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBurnerUnlocked represents a Unlocked event raised by the BridgeBurner contract.
type BridgeBurnerUnlocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeBurner *BridgeBurnerFilterer) FilterUnlocked(opts *bind.FilterOpts, sender []common.Address) (*BridgeBurnerUnlockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeBurner.contract.FilterLogs(opts, "Unlocked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBurnerUnlockedIterator{contract: _BridgeBurner.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeBurner *BridgeBurnerFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *BridgeBurnerUnlocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeBurner.contract.WatchLogs(opts, "Unlocked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBurnerUnlocked)
				if err := _BridgeBurner.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeBurner *BridgeBurnerFilterer) ParseUnlocked(log types.Log) (*BridgeBurnerUnlocked, error) {
	event := new(BridgeBurnerUnlocked)
	if err := _BridgeBurner.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBurnerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BridgeBurner contract.
type BridgeBurnerUnpausedIterator struct {
	Event *BridgeBurnerUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeBurnerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBurnerUnpaused)
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
		it.Event = new(BridgeBurnerUnpaused)
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
func (it *BridgeBurnerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBurnerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBurnerUnpaused represents a Unpaused event raised by the BridgeBurner contract.
type BridgeBurnerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeBurner *BridgeBurnerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeBurnerUnpausedIterator, error) {

	logs, sub, err := _BridgeBurner.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeBurnerUnpausedIterator{contract: _BridgeBurner.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeBurner *BridgeBurnerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeBurnerUnpaused) (event.Subscription, error) {

	logs, sub, err := _BridgeBurner.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBurnerUnpaused)
				if err := _BridgeBurner.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address account)
func (_BridgeBurner *BridgeBurnerFilterer) ParseUnpaused(log types.Log) (*BridgeBurnerUnpaused, error) {
	event := new(BridgeBurnerUnpaused)
	if err := _BridgeBurner.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEtherABI is the input ABI used to generate the binding from.
const BridgeEtherABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"contractIFee\",\"name\":\"fee\",\"type\":\"address\"},{\"internalType\":\"contractILimiter\",\"name\":\"limiter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"contractIFee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimiter\",\"outputs\":[{\"internalType\":\"contractILimiter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimiterUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"isLimited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isUnlockCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFee\",\"name\":\"fee_\",\"type\":\"address\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILimiter\",\"name\":\"limiter\",\"type\":\"address\"}],\"name\":\"setLimiter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// BridgeEtherFuncSigs maps the 4-byte function signature to its string representation.
var BridgeEtherFuncSigs = map[string]string{
	"99a5d747": "calculateFee(uint256)",
	"ced72f87": "getFee()",
	"eb2a0d1f": "getLimiter()",
	"7eb76b29": "getLimiterUsage()",
	"08a90d5a": "isLimited(uint256)",
	"a4d7fa93": "isUnlockCompleted(bytes32)",
	"dd467064": "lock(uint256)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"7917fb9f": "setFee(address)",
	"7a29084c": "setLimiter(address)",
	"f2fde38b": "transferOwnership(address)",
	"b322edea": "unlock(address,uint256,bytes32)",
	"3f4ba83a": "unpause()",
}

// BridgeEtherBin is the compiled bytecode used for deploying new contracts.
var BridgeEtherBin = "0x60806040523480156200001157600080fd5b50604051620013633803806200136383398101604081905262000034916200019c565b600080546001600160a01b031916339081178255604051859285928592909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19169055600180558251620000a1906002906020860190620000d9565b50600380546001600160a01b039384166001600160a01b0319918216179091556004805492909316911617905550620002ec92505050565b828054620000e79062000299565b90600052602060002090601f0160209004810192826200010b576000855562000156565b82601f106200012657805160ff191683800117855562000156565b8280016001018555821562000156579182015b828111156200015657825182559160200191906001019062000139565b506200016492915062000168565b5090565b5b8082111562000164576000815560010162000169565b80516001600160a01b03811681146200019757600080fd5b919050565b600080600060608486031215620001b1578283fd5b83516001600160401b0380821115620001c8578485fd5b818601915086601f830112620001dc578485fd5b815181811115620001f157620001f1620002d6565b604051601f8201601f19908116603f011681019083821181831017156200021c576200021c620002d6565b8160405282815260209350898484870101111562000238578788fd5b8791505b828210156200025b57848201840151818301850152908301906200023c565b828211156200026c57878484830101525b96506200027e9150508682016200017f565b9350505062000290604085016200017f565b90509250925092565b600181811c90821680620002ae57607f821691505b60208210811415620002d057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b61106780620002fc6000396000f3fe6080604052600436106101025760003560e01c80638456cb5911610095578063b322edea11610064578063b322edea146102b1578063ced72f87146102d1578063dd467064146102ef578063eb2a0d1f14610302578063f2fde38b1461032057600080fd5b80638456cb591461021a5780638da5cb5b1461022f57806399a5d74714610261578063a4d7fa931461028157600080fd5b8063715018a6116100d1578063715018a6146101a25780637917fb9f146101b75780637a29084c146101d75780637eb76b29146101f757600080fd5b806306fdde031461011157806308a90d5a1461013c5780633f4ba83a1461016c5780635c975abb1461018357600080fd5b3661010c57600080fd5b600080fd5b34801561011d57600080fd5b50610126610340565b6040516101339190610eb8565b60405180910390f35b34801561014857600080fd5b5061015c610157366004610e88565b6103d2565b6040519015158152602001610133565b34801561017857600080fd5b50610181610460565b005b34801561018f57600080fd5b50600054600160a01b900460ff1661015c565b3480156101ae57600080fd5b5061018161049d565b3480156101c357600080fd5b506101816101d2366004610e11565b610582565b3480156101e357600080fd5b506101816101f2366004610e11565b6105ce565b34801561020357600080fd5b5061020c61061a565b604051908152602001610133565b34801561022657600080fd5b506101816106a0565b34801561023b57600080fd5b506000546001600160a01b03165b6040516001600160a01b039091168152602001610133565b34801561026d57600080fd5b5061020c61027c366004610e88565b6106d2565b34801561028d57600080fd5b5061015c61029c366004610e88565b60009081526005602052604090205460ff1690565b3480156102bd57600080fd5b506101816102cc366004610e34565b610769565b3480156102dd57600080fd5b506003546001600160a01b0316610249565b6101816102fd366004610e88565b6108d4565b34801561030e57600080fd5b506004546001600160a01b0316610249565b34801561032c57600080fd5b5061018161033b366004610e11565b610aa2565b60606002805461034f90610fde565b80601f016020809104026020016040519081016040528092919081815260200182805461037b90610fde565b80156103c85780601f1061039d576101008083540402835291602001916103c8565b820191906000526020600020905b8154815290600101906020018083116103ab57829003601f168201915b5050505050905090565b600480546040516323b0c65960e11b81523092810192909252602482018390526000916001600160a01b03909116906347618cb29060440160206040518083038186803b15801561042257600080fd5b505afa158015610436573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045a9190610e68565b92915050565b6000546001600160a01b031633146104935760405162461bcd60e51b815260040161048a90610f4e565b60405180910390fd5b61049b610b8c565b565b6000546001600160a01b031633146104c75760405162461bcd60e51b815260040161048a90610f4e565b600260015414156104ea5760405162461bcd60e51b815260040161048a90610f83565b600260015547801561056b57600080546040516001600160a01b039091169083908381818185875af1925050503d8060008114610543576040519150601f19603f3d011682016040523d82523d6000602084013e610548565b606091505b50509050806105695760405162461bcd60e51b815260040161048a90610f0b565b505b610573610c29565b61057b610cb1565b5060018055565b6000546001600160a01b031633146105ac5760405162461bcd60e51b815260040161048a90610f4e565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146105f85760405162461bcd60e51b815260040161048a90610f4e565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b60048054604051632cdcd8af60e11b815230928101929092526000916001600160a01b03909116906359b9b15e9060240160206040518083038186803b15801561066357600080fd5b505afa158015610677573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069b9190610ea0565b905090565b6000546001600160a01b031633146106ca5760405162461bcd60e51b815260040161048a90610f4e565b61049b610c29565b6003546000906001600160a01b03166106ed57506000919050565b60035460405163173b25bd60e31b8152600481018490526001600160a01b039091169063b9d92de89060240160206040518083038186803b15801561073157600080fd5b505afa158015610745573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045a9190610ea0565b6000546001600160a01b031633146107935760405162461bcd60e51b815260040161048a90610f4e565b600260015414156107b65760405162461bcd60e51b815260040161048a90610f83565b60026001556107c481610d25565b814710156108145760405162461bcd60e51b815260206004820152601d60248201527f42726964676545746865723a206e6f7420656e6f756768206574686572000000604482015260640161048a565b6000836001600160a01b03168360405160006040518083038185875af1925050503d8060008114610861576040519150601f19603f3d011682016040523d82523d6000602084013e610866565b606091505b50509050806108875760405162461bcd60e51b815260040161048a90610f0b565b836001600160a01b03167f0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d846040516108c291815260200190565b60405180910390a25050600180555050565b600260015414156108f75760405162461bcd60e51b815260040161048a90610f83565b6002600155600054600160a01b900460ff16156109495760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161048a565b61095281610d9f565b600061095d826106d2565b90506109698183610fba565b34146109b75760405162461bcd60e51b815260206004820152601a60248201527f42726964676545746865723a20696e76616c6964206574686572000000000000604482015260640161048a565b600080546040516001600160a01b039091169083908381818185875af1925050503d8060008114610a04576040519150601f19603f3d011682016040523d82523d6000602084013e610a09565b606091505b5050905080610a645760405162461bcd60e51b815260206004820152602160248201527f42726964676545746865723a2063616e206e6f74207472616e736665722066656044820152606560f81b606482015260840161048a565b60405183815233907f9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd600089060200160405180910390a250506001805550565b6000546001600160a01b03163314610acc5760405162461bcd60e51b815260040161048a90610f4e565b6001600160a01b038116610b315760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161048a565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600054600160a01b900460ff16610bdc5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015260640161048a565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600054600160a01b900460ff1615610c765760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161048a565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610c0c3390565b6000546001600160a01b03163314610cdb5760405162461bcd60e51b815260040161048a90610f4e565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008181526005602052604090205460ff1615610d845760405162461bcd60e51b815260206004820152601c60248201527f427269646765426173653a20616c726561647920756e6c6f636b656400000000604482015260640161048a565b6000908152600560205260409020805460ff19166001179055565b6004546001600160a01b0316610db25750565b6004805460405163606ecf2960e11b81529182018390526001600160a01b03169063c0dd9e5290602401600060405180830381600087803b158015610df657600080fd5b505af1158015610e0a573d6000803e3d6000fd5b5050505050565b600060208284031215610e22578081fd5b8135610e2d81611019565b9392505050565b600080600060608486031215610e48578182fd5b8335610e5381611019565b95602085013595506040909401359392505050565b600060208284031215610e79578081fd5b81518015158114610e2d578182fd5b600060208284031215610e99578081fd5b5035919050565b600060208284031215610eb1578081fd5b5051919050565b6000602080835283518082850152825b81811015610ee457858101830151858201604001528201610ec8565b81811115610ef55783604083870101525b50601f01601f1916929092016040019392505050565b60208082526023908201527f42726964676545746865723a2063616e206e6f74207472616e736665722065746040820152623432b960e91b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b60008219821115610fd957634e487b7160e01b81526011600452602481fd5b500190565b600181811c90821680610ff257607f821691505b6020821081141561101357634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b038116811461102e57600080fd5b5056fea2646970667358221220911552d0c7731db408d3365897107733f850ef561bf8b68220ff144dbda066e064736f6c63430008040033"

// DeployBridgeEther deploys a new Ethereum contract, binding an instance of BridgeEther to it.
func DeployBridgeEther(auth *bind.TransactOpts, backend bind.ContractBackend, name string, fee common.Address, limiter common.Address) (common.Address, *types.Transaction, *BridgeEther, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeEtherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeEtherBin), backend, name, fee, limiter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeEther{BridgeEtherCaller: BridgeEtherCaller{contract: contract}, BridgeEtherTransactor: BridgeEtherTransactor{contract: contract}, BridgeEtherFilterer: BridgeEtherFilterer{contract: contract}}, nil
}

// BridgeEther is an auto generated Go binding around an Ethereum contract.
type BridgeEther struct {
	BridgeEtherCaller     // Read-only binding to the contract
	BridgeEtherTransactor // Write-only binding to the contract
	BridgeEtherFilterer   // Log filterer for contract events
}

// BridgeEtherCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeEtherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEtherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeEtherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEtherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeEtherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEtherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeEtherSession struct {
	Contract     *BridgeEther      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeEtherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeEtherCallerSession struct {
	Contract *BridgeEtherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeEtherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeEtherTransactorSession struct {
	Contract     *BridgeEtherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeEtherRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeEtherRaw struct {
	Contract *BridgeEther // Generic contract binding to access the raw methods on
}

// BridgeEtherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeEtherCallerRaw struct {
	Contract *BridgeEtherCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeEtherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeEtherTransactorRaw struct {
	Contract *BridgeEtherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeEther creates a new instance of BridgeEther, bound to a specific deployed contract.
func NewBridgeEther(address common.Address, backend bind.ContractBackend) (*BridgeEther, error) {
	contract, err := bindBridgeEther(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeEther{BridgeEtherCaller: BridgeEtherCaller{contract: contract}, BridgeEtherTransactor: BridgeEtherTransactor{contract: contract}, BridgeEtherFilterer: BridgeEtherFilterer{contract: contract}}, nil
}

// NewBridgeEtherCaller creates a new read-only instance of BridgeEther, bound to a specific deployed contract.
func NewBridgeEtherCaller(address common.Address, caller bind.ContractCaller) (*BridgeEtherCaller, error) {
	contract, err := bindBridgeEther(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeEtherCaller{contract: contract}, nil
}

// NewBridgeEtherTransactor creates a new write-only instance of BridgeEther, bound to a specific deployed contract.
func NewBridgeEtherTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeEtherTransactor, error) {
	contract, err := bindBridgeEther(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeEtherTransactor{contract: contract}, nil
}

// NewBridgeEtherFilterer creates a new log filterer instance of BridgeEther, bound to a specific deployed contract.
func NewBridgeEtherFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeEtherFilterer, error) {
	contract, err := bindBridgeEther(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeEtherFilterer{contract: contract}, nil
}

// bindBridgeEther binds a generic wrapper to an already deployed contract.
func bindBridgeEther(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeEtherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeEther *BridgeEtherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeEther.Contract.BridgeEtherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeEther *BridgeEtherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEther.Contract.BridgeEtherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeEther *BridgeEtherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeEther.Contract.BridgeEtherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeEther *BridgeEtherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeEther.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeEther *BridgeEtherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEther.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeEther *BridgeEtherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeEther.Contract.contract.Transact(opts, method, params...)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeEther *BridgeEtherCaller) CalculateFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEther.contract.Call(opts, &out, "calculateFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeEther *BridgeEtherSession) CalculateFee(amount *big.Int) (*big.Int, error) {
	return _BridgeEther.Contract.CalculateFee(&_BridgeEther.CallOpts, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeEther *BridgeEtherCallerSession) CalculateFee(amount *big.Int) (*big.Int, error) {
	return _BridgeEther.Contract.CalculateFee(&_BridgeEther.CallOpts, amount)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeEther *BridgeEtherCaller) GetFee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeEther.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeEther *BridgeEtherSession) GetFee() (common.Address, error) {
	return _BridgeEther.Contract.GetFee(&_BridgeEther.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeEther *BridgeEtherCallerSession) GetFee() (common.Address, error) {
	return _BridgeEther.Contract.GetFee(&_BridgeEther.CallOpts)
}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeEther *BridgeEtherCaller) GetLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeEther.contract.Call(opts, &out, "getLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeEther *BridgeEtherSession) GetLimiter() (common.Address, error) {
	return _BridgeEther.Contract.GetLimiter(&_BridgeEther.CallOpts)
}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeEther *BridgeEtherCallerSession) GetLimiter() (common.Address, error) {
	return _BridgeEther.Contract.GetLimiter(&_BridgeEther.CallOpts)
}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeEther *BridgeEtherCaller) GetLimiterUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEther.contract.Call(opts, &out, "getLimiterUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeEther *BridgeEtherSession) GetLimiterUsage() (*big.Int, error) {
	return _BridgeEther.Contract.GetLimiterUsage(&_BridgeEther.CallOpts)
}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeEther *BridgeEtherCallerSession) GetLimiterUsage() (*big.Int, error) {
	return _BridgeEther.Contract.GetLimiterUsage(&_BridgeEther.CallOpts)
}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeEther *BridgeEtherCaller) IsLimited(opts *bind.CallOpts, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _BridgeEther.contract.Call(opts, &out, "isLimited", amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeEther *BridgeEtherSession) IsLimited(amount *big.Int) (bool, error) {
	return _BridgeEther.Contract.IsLimited(&_BridgeEther.CallOpts, amount)
}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeEther *BridgeEtherCallerSession) IsLimited(amount *big.Int) (bool, error) {
	return _BridgeEther.Contract.IsLimited(&_BridgeEther.CallOpts, amount)
}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeEther *BridgeEtherCaller) IsUnlockCompleted(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _BridgeEther.contract.Call(opts, &out, "isUnlockCompleted", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeEther *BridgeEtherSession) IsUnlockCompleted(hash [32]byte) (bool, error) {
	return _BridgeEther.Contract.IsUnlockCompleted(&_BridgeEther.CallOpts, hash)
}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeEther *BridgeEtherCallerSession) IsUnlockCompleted(hash [32]byte) (bool, error) {
	return _BridgeEther.Contract.IsUnlockCompleted(&_BridgeEther.CallOpts, hash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeEther *BridgeEtherCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeEther.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeEther *BridgeEtherSession) Name() (string, error) {
	return _BridgeEther.Contract.Name(&_BridgeEther.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeEther *BridgeEtherCallerSession) Name() (string, error) {
	return _BridgeEther.Contract.Name(&_BridgeEther.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeEther *BridgeEtherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeEther.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeEther *BridgeEtherSession) Owner() (common.Address, error) {
	return _BridgeEther.Contract.Owner(&_BridgeEther.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeEther *BridgeEtherCallerSession) Owner() (common.Address, error) {
	return _BridgeEther.Contract.Owner(&_BridgeEther.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeEther *BridgeEtherCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeEther.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeEther *BridgeEtherSession) Paused() (bool, error) {
	return _BridgeEther.Contract.Paused(&_BridgeEther.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeEther *BridgeEtherCallerSession) Paused() (bool, error) {
	return _BridgeEther.Contract.Paused(&_BridgeEther.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeEther *BridgeEtherTransactor) Lock(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BridgeEther.contract.Transact(opts, "lock", amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeEther *BridgeEtherSession) Lock(amount *big.Int) (*types.Transaction, error) {
	return _BridgeEther.Contract.Lock(&_BridgeEther.TransactOpts, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeEther *BridgeEtherTransactorSession) Lock(amount *big.Int) (*types.Transaction, error) {
	return _BridgeEther.Contract.Lock(&_BridgeEther.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeEther *BridgeEtherTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEther.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeEther *BridgeEtherSession) Pause() (*types.Transaction, error) {
	return _BridgeEther.Contract.Pause(&_BridgeEther.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeEther *BridgeEtherTransactorSession) Pause() (*types.Transaction, error) {
	return _BridgeEther.Contract.Pause(&_BridgeEther.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeEther *BridgeEtherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEther.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeEther *BridgeEtherSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeEther.Contract.RenounceOwnership(&_BridgeEther.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeEther *BridgeEtherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeEther.Contract.RenounceOwnership(&_BridgeEther.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeEther *BridgeEtherTransactor) SetFee(opts *bind.TransactOpts, fee_ common.Address) (*types.Transaction, error) {
	return _BridgeEther.contract.Transact(opts, "setFee", fee_)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeEther *BridgeEtherSession) SetFee(fee_ common.Address) (*types.Transaction, error) {
	return _BridgeEther.Contract.SetFee(&_BridgeEther.TransactOpts, fee_)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeEther *BridgeEtherTransactorSession) SetFee(fee_ common.Address) (*types.Transaction, error) {
	return _BridgeEther.Contract.SetFee(&_BridgeEther.TransactOpts, fee_)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeEther *BridgeEtherTransactor) SetLimiter(opts *bind.TransactOpts, limiter common.Address) (*types.Transaction, error) {
	return _BridgeEther.contract.Transact(opts, "setLimiter", limiter)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeEther *BridgeEtherSession) SetLimiter(limiter common.Address) (*types.Transaction, error) {
	return _BridgeEther.Contract.SetLimiter(&_BridgeEther.TransactOpts, limiter)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeEther *BridgeEtherTransactorSession) SetLimiter(limiter common.Address) (*types.Transaction, error) {
	return _BridgeEther.Contract.SetLimiter(&_BridgeEther.TransactOpts, limiter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeEther *BridgeEtherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeEther.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeEther *BridgeEtherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeEther.Contract.TransferOwnership(&_BridgeEther.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeEther *BridgeEtherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeEther.Contract.TransferOwnership(&_BridgeEther.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeEther *BridgeEtherTransactor) Unlock(opts *bind.TransactOpts, account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeEther.contract.Transact(opts, "unlock", account, amount, hash)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeEther *BridgeEtherSession) Unlock(account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeEther.Contract.Unlock(&_BridgeEther.TransactOpts, account, amount, hash)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeEther *BridgeEtherTransactorSession) Unlock(account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeEther.Contract.Unlock(&_BridgeEther.TransactOpts, account, amount, hash)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeEther *BridgeEtherTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEther.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeEther *BridgeEtherSession) Unpause() (*types.Transaction, error) {
	return _BridgeEther.Contract.Unpause(&_BridgeEther.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeEther *BridgeEtherTransactorSession) Unpause() (*types.Transaction, error) {
	return _BridgeEther.Contract.Unpause(&_BridgeEther.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeEther *BridgeEtherTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEther.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeEther *BridgeEtherSession) Receive() (*types.Transaction, error) {
	return _BridgeEther.Contract.Receive(&_BridgeEther.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeEther *BridgeEtherTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeEther.Contract.Receive(&_BridgeEther.TransactOpts)
}

// BridgeEtherLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the BridgeEther contract.
type BridgeEtherLockedIterator struct {
	Event *BridgeEtherLocked // Event containing the contract specifics and raw log

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
func (it *BridgeEtherLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEtherLocked)
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
		it.Event = new(BridgeEtherLocked)
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
func (it *BridgeEtherLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEtherLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEtherLocked represents a Locked event raised by the BridgeEther contract.
type BridgeEtherLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeEther *BridgeEtherFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*BridgeEtherLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeEther.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEtherLockedIterator{contract: _BridgeEther.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeEther *BridgeEtherFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *BridgeEtherLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeEther.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEtherLocked)
				if err := _BridgeEther.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeEther *BridgeEtherFilterer) ParseLocked(log types.Log) (*BridgeEtherLocked, error) {
	event := new(BridgeEtherLocked)
	if err := _BridgeEther.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEtherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeEther contract.
type BridgeEtherOwnershipTransferredIterator struct {
	Event *BridgeEtherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeEtherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEtherOwnershipTransferred)
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
		it.Event = new(BridgeEtherOwnershipTransferred)
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
func (it *BridgeEtherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEtherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEtherOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeEther contract.
type BridgeEtherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeEther *BridgeEtherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeEtherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeEther.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEtherOwnershipTransferredIterator{contract: _BridgeEther.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeEther *BridgeEtherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeEtherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeEther.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEtherOwnershipTransferred)
				if err := _BridgeEther.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeEther *BridgeEtherFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeEtherOwnershipTransferred, error) {
	event := new(BridgeEtherOwnershipTransferred)
	if err := _BridgeEther.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEtherPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BridgeEther contract.
type BridgeEtherPausedIterator struct {
	Event *BridgeEtherPaused // Event containing the contract specifics and raw log

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
func (it *BridgeEtherPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEtherPaused)
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
		it.Event = new(BridgeEtherPaused)
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
func (it *BridgeEtherPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEtherPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEtherPaused represents a Paused event raised by the BridgeEther contract.
type BridgeEtherPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeEther *BridgeEtherFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgeEtherPausedIterator, error) {

	logs, sub, err := _BridgeEther.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgeEtherPausedIterator{contract: _BridgeEther.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeEther *BridgeEtherFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgeEtherPaused) (event.Subscription, error) {

	logs, sub, err := _BridgeEther.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEtherPaused)
				if err := _BridgeEther.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address account)
func (_BridgeEther *BridgeEtherFilterer) ParsePaused(log types.Log) (*BridgeEtherPaused, error) {
	event := new(BridgeEtherPaused)
	if err := _BridgeEther.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEtherUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the BridgeEther contract.
type BridgeEtherUnlockedIterator struct {
	Event *BridgeEtherUnlocked // Event containing the contract specifics and raw log

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
func (it *BridgeEtherUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEtherUnlocked)
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
		it.Event = new(BridgeEtherUnlocked)
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
func (it *BridgeEtherUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEtherUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEtherUnlocked represents a Unlocked event raised by the BridgeEther contract.
type BridgeEtherUnlocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeEther *BridgeEtherFilterer) FilterUnlocked(opts *bind.FilterOpts, sender []common.Address) (*BridgeEtherUnlockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeEther.contract.FilterLogs(opts, "Unlocked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEtherUnlockedIterator{contract: _BridgeEther.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeEther *BridgeEtherFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *BridgeEtherUnlocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeEther.contract.WatchLogs(opts, "Unlocked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEtherUnlocked)
				if err := _BridgeEther.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeEther *BridgeEtherFilterer) ParseUnlocked(log types.Log) (*BridgeEtherUnlocked, error) {
	event := new(BridgeEtherUnlocked)
	if err := _BridgeEther.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEtherUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BridgeEther contract.
type BridgeEtherUnpausedIterator struct {
	Event *BridgeEtherUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeEtherUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEtherUnpaused)
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
		it.Event = new(BridgeEtherUnpaused)
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
func (it *BridgeEtherUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEtherUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEtherUnpaused represents a Unpaused event raised by the BridgeEther contract.
type BridgeEtherUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeEther *BridgeEtherFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeEtherUnpausedIterator, error) {

	logs, sub, err := _BridgeEther.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeEtherUnpausedIterator{contract: _BridgeEther.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeEther *BridgeEtherFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeEtherUnpaused) (event.Subscription, error) {

	logs, sub, err := _BridgeEther.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEtherUnpaused)
				if err := _BridgeEther.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address account)
func (_BridgeEther *BridgeEtherFilterer) ParseUnpaused(log types.Log) (*BridgeEtherUnpaused, error) {
	event := new(BridgeEtherUnpaused)
	if err := _BridgeEther.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLockerABI is the input ABI used to generate the binding from.
const BridgeLockerABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"contractIFee\",\"name\":\"fee\",\"type\":\"address\"},{\"internalType\":\"contractILimiter\",\"name\":\"limiter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"contractIFee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimiter\",\"outputs\":[{\"internalType\":\"contractILimiter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimiterUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"isLimited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isUnlockCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFee\",\"name\":\"fee_\",\"type\":\"address\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILimiter\",\"name\":\"limiter\",\"type\":\"address\"}],\"name\":\"setLimiter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// BridgeLockerFuncSigs maps the 4-byte function signature to its string representation.
var BridgeLockerFuncSigs = map[string]string{
	"99a5d747": "calculateFee(uint256)",
	"ced72f87": "getFee()",
	"eb2a0d1f": "getLimiter()",
	"7eb76b29": "getLimiterUsage()",
	"08a90d5a": "isLimited(uint256)",
	"a4d7fa93": "isUnlockCompleted(bytes32)",
	"dd467064": "lock(uint256)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"7917fb9f": "setFee(address)",
	"7a29084c": "setLimiter(address)",
	"fc0c546a": "token()",
	"f2fde38b": "transferOwnership(address)",
	"b322edea": "unlock(address,uint256,bytes32)",
	"3f4ba83a": "unpause()",
}

// BridgeLockerBin is the compiled bytecode used for deploying new contracts.
var BridgeLockerBin = "0x60806040523480156200001157600080fd5b50604051620015ba380380620015ba8339810160408190526200003491620001a6565b600080546001600160a01b031916339081178255604051859285928592909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19169055600180558251620000a1906002906020860190620000ee565b50600380546001600160a01b039384166001600160a01b03199182161790915560048054928416928216929092179091556006805497909216961695909517909455506200032692505050565b828054620000fc90620002ba565b90600052602060002090601f0160209004810192826200012057600085556200016b565b82601f106200013b57805160ff19168380011785556200016b565b828001600101855582156200016b579182015b828111156200016b5782518255916020019190600101906200014e565b50620001799291506200017d565b5090565b5b808211156200017957600081556001016200017e565b8051620001a1816200030d565b919050565b60008060008060808587031215620001bc578384fd5b8451620001c9816200030d565b602086810151919550906001600160401b0380821115620001e8578586fd5b818801915088601f830112620001fc578586fd5b815181811115620002115762000211620002f7565b604051601f8201601f19908116603f011681019083821181831017156200023c576200023c620002f7565b816040528281528b8684870101111562000254578889fd5b8893505b8284101562000277578484018601518185018701529285019262000258565b828411156200028857888684830101525b8098505050505050506200029f6040860162000194565b9150620002af6060860162000194565b905092959194509250565b600181811c90821680620002cf57607f821691505b60208210811415620002f157634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146200032357600080fd5b50565b61128480620003366000396000f3fe60806040526004361061010d5760003560e01c80638da5cb5b11610095578063ced72f8711610064578063ced72f87146102dc578063dd467064146102fa578063eb2a0d1f1461030d578063f2fde38b1461032b578063fc0c546a1461034b57600080fd5b80638da5cb5b1461023a57806399a5d7471461026c578063a4d7fa931461028c578063b322edea146102bc57600080fd5b8063715018a6116100dc578063715018a6146101ad5780637917fb9f146101c25780637a29084c146101e25780637eb76b29146102025780638456cb591461022557600080fd5b806306fdde031461011c57806308a90d5a146101475780633f4ba83a146101775780635c975abb1461018e57600080fd5b3661011757600080fd5b600080fd5b34801561012857600080fd5b50610131610369565b60405161013e919061116a565b60405180910390f35b34801561015357600080fd5b5061016761016236600461111e565b6103fb565b604051901515815260200161013e565b34801561018357600080fd5b5061018c610489565b005b34801561019a57600080fd5b50600054600160a01b900460ff16610167565b3480156101b957600080fd5b5061018c6104c6565b3480156101ce57600080fd5b5061018c6101dd3660046110ae565b6105af565b3480156101ee57600080fd5b5061018c6101fd3660046110ae565b6105fb565b34801561020e57600080fd5b50610217610647565b60405190815260200161013e565b34801561023157600080fd5b5061018c6106cd565b34801561024657600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200161013e565b34801561027857600080fd5b5061021761028736600461111e565b6106ff565b34801561029857600080fd5b506101676102a736600461111e565b60009081526005602052604090205460ff1690565b3480156102c857600080fd5b5061018c6102d73660046110ca565b610796565b3480156102e857600080fd5b506003546001600160a01b0316610254565b61018c61030836600461111e565b610828565b34801561031957600080fd5b506004546001600160a01b0316610254565b34801561033757600080fd5b5061018c6103463660046110ae565b610881565b34801561035757600080fd5b506006546001600160a01b0316610254565b606060028054610378906111fe565b80601f01602080910402602001604051908101604052809291908181526020018280546103a4906111fe565b80156103f15780601f106103c6576101008083540402835291602001916103f1565b820191906000526020600020905b8154815290600101906020018083116103d457829003601f168201915b5050505050905090565b600480546040516323b0c65960e11b81523092810192909252602482018390526000916001600160a01b03909116906347618cb29060440160206040518083038186803b15801561044b57600080fd5b505afa15801561045f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048391906110fe565b92915050565b6000546001600160a01b031633146104bc5760405162461bcd60e51b81526004016104b39061119d565b60405180910390fd5b6104c461096b565b565b6000546001600160a01b031633146104f05760405162461bcd60e51b81526004016104b39061119d565b6006546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a082319060240160206040518083038186803b15801561053457600080fd5b505afa158015610548573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056c9190611136565b9050801561059c5761059c6105896000546001600160a01b031690565b6006546001600160a01b03169083610a08565b6105a4610a70565b6105ac610af8565b50565b6000546001600160a01b031633146105d95760405162461bcd60e51b81526004016104b39061119d565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146106255760405162461bcd60e51b81526004016104b39061119d565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b60048054604051632cdcd8af60e11b815230928101929092526000916001600160a01b03909116906359b9b15e9060240160206040518083038186803b15801561069057600080fd5b505afa1580156106a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106c89190611136565b905090565b6000546001600160a01b031633146106f75760405162461bcd60e51b81526004016104b39061119d565b6104c4610a70565b6003546000906001600160a01b031661071a57506000919050565b60035460405163173b25bd60e31b8152600481018490526001600160a01b039091169063b9d92de89060240160206040518083038186803b15801561075e57600080fd5b505afa158015610772573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104839190611136565b6000546001600160a01b031633146107c05760405162461bcd60e51b81526004016104b39061119d565b6107c981610b6c565b6006546107e0906001600160a01b03168484610a08565b826001600160a01b03167f0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d8360405161081b91815260200190565b60405180910390a2505050565b61083181610be6565b610849336006546001600160a01b0316903084610c45565b60405181815233907f9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd600089060200160405180910390a250565b6000546001600160a01b031633146108ab5760405162461bcd60e51b81526004016104b39061119d565b6001600160a01b0381166109105760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104b3565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600054600160a01b900460ff166109bb5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016104b3565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6040516001600160a01b038316602482015260448101829052610a6b90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610c83565b505050565b600054600160a01b900460ff1615610abd5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016104b3565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586109eb3390565b6000546001600160a01b03163314610b225760405162461bcd60e51b81526004016104b39061119d565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008181526005602052604090205460ff1615610bcb5760405162461bcd60e51b815260206004820152601c60248201527f427269646765426173653a20616c726561647920756e6c6f636b65640000000060448201526064016104b3565b6000908152600560205260409020805460ff19166001179055565b600054600160a01b900460ff1615610c335760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016104b3565b610c3c81610d55565b6105ac81610dc7565b6040516001600160a01b0380851660248301528316604482015260648101829052610c7d9085906323b872dd60e01b90608401610a34565b50505050565b6000610cd8826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610f349092919063ffffffff16565b805190915015610a6b5780806020019051810190610cf691906110fe565b610a6b5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016104b3565b6004546001600160a01b0316610d685750565b6004805460405163606ecf2960e11b81529182018390526001600160a01b03169063c0dd9e5290602401600060405180830381600087803b158015610dac57600080fd5b505af1158015610dc0573d6000803e3d6000fd5b5050505050565b60026001541415610e1a5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016104b3565b60026001556000610e2a826106ff565b905080610e375750610f2d565b80341015610e875760405162461bcd60e51b815260206004820152601a60248201527f427269646765426173653a206e6f7420656e6f7567682066656500000000000060448201526064016104b3565b600080546040516001600160a01b039091169034908381818185875af1925050503d8060008114610ed4576040519150601f19603f3d011682016040523d82523d6000602084013e610ed9565b606091505b5050905080610f2a5760405162461bcd60e51b815260206004820181905260248201527f427269646765426173653a2063616e206e6f74207472616e736665722066656560448201526064016104b3565b50505b5060018055565b6060610f438484600085610f4d565b90505b9392505050565b606082471015610fae5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016104b3565b843b610ffc5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016104b3565b600080866001600160a01b03168587604051611018919061114e565b60006040518083038185875af1925050503d8060008114611055576040519150601f19603f3d011682016040523d82523d6000602084013e61105a565b606091505b509150915061106a828286611075565b979650505050505050565b60608315611084575081610f46565b8251156110945782518084602001fd5b8160405162461bcd60e51b81526004016104b3919061116a565b6000602082840312156110bf578081fd5b8135610f4681611239565b6000806000606084860312156110de578182fd5b83356110e981611239565b95602085013595506040909401359392505050565b60006020828403121561110f578081fd5b81518015158114610f46578182fd5b60006020828403121561112f578081fd5b5035919050565b600060208284031215611147578081fd5b5051919050565b600082516111608184602087016111d2565b9190910192915050565b60208152600082518060208401526111898160408501602087016111d2565b601f01601f19169190910160400192915050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60005b838110156111ed5781810151838201526020016111d5565b83811115610c7d5750506000910152565b600181811c9082168061121257607f821691505b6020821081141561123357634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b03811681146105ac57600080fdfea26469706673582212201908195b2a8a43a6de4e53d5c0772f749ec454e537e55a51bc58a261b211059664736f6c63430008040033"

// DeployBridgeLocker deploys a new Ethereum contract, binding an instance of BridgeLocker to it.
func DeployBridgeLocker(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address, name string, fee common.Address, limiter common.Address) (common.Address, *types.Transaction, *BridgeLocker, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeLockerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeLockerBin), backend, token_, name, fee, limiter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeLocker{BridgeLockerCaller: BridgeLockerCaller{contract: contract}, BridgeLockerTransactor: BridgeLockerTransactor{contract: contract}, BridgeLockerFilterer: BridgeLockerFilterer{contract: contract}}, nil
}

// BridgeLocker is an auto generated Go binding around an Ethereum contract.
type BridgeLocker struct {
	BridgeLockerCaller     // Read-only binding to the contract
	BridgeLockerTransactor // Write-only binding to the contract
	BridgeLockerFilterer   // Log filterer for contract events
}

// BridgeLockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeLockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeLockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeLockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeLockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeLockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeLockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeLockerSession struct {
	Contract     *BridgeLocker     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeLockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeLockerCallerSession struct {
	Contract *BridgeLockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BridgeLockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeLockerTransactorSession struct {
	Contract     *BridgeLockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeLockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeLockerRaw struct {
	Contract *BridgeLocker // Generic contract binding to access the raw methods on
}

// BridgeLockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeLockerCallerRaw struct {
	Contract *BridgeLockerCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeLockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeLockerTransactorRaw struct {
	Contract *BridgeLockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeLocker creates a new instance of BridgeLocker, bound to a specific deployed contract.
func NewBridgeLocker(address common.Address, backend bind.ContractBackend) (*BridgeLocker, error) {
	contract, err := bindBridgeLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeLocker{BridgeLockerCaller: BridgeLockerCaller{contract: contract}, BridgeLockerTransactor: BridgeLockerTransactor{contract: contract}, BridgeLockerFilterer: BridgeLockerFilterer{contract: contract}}, nil
}

// NewBridgeLockerCaller creates a new read-only instance of BridgeLocker, bound to a specific deployed contract.
func NewBridgeLockerCaller(address common.Address, caller bind.ContractCaller) (*BridgeLockerCaller, error) {
	contract, err := bindBridgeLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeLockerCaller{contract: contract}, nil
}

// NewBridgeLockerTransactor creates a new write-only instance of BridgeLocker, bound to a specific deployed contract.
func NewBridgeLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeLockerTransactor, error) {
	contract, err := bindBridgeLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeLockerTransactor{contract: contract}, nil
}

// NewBridgeLockerFilterer creates a new log filterer instance of BridgeLocker, bound to a specific deployed contract.
func NewBridgeLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeLockerFilterer, error) {
	contract, err := bindBridgeLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeLockerFilterer{contract: contract}, nil
}

// bindBridgeLocker binds a generic wrapper to an already deployed contract.
func bindBridgeLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeLockerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeLocker *BridgeLockerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeLocker.Contract.BridgeLockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeLocker *BridgeLockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLocker.Contract.BridgeLockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeLocker *BridgeLockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeLocker.Contract.BridgeLockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeLocker *BridgeLockerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeLocker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeLocker *BridgeLockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLocker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeLocker *BridgeLockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeLocker.Contract.contract.Transact(opts, method, params...)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeLocker *BridgeLockerCaller) CalculateFee(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeLocker.contract.Call(opts, &out, "calculateFee", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeLocker *BridgeLockerSession) CalculateFee(amount *big.Int) (*big.Int, error) {
	return _BridgeLocker.Contract.CalculateFee(&_BridgeLocker.CallOpts, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256)
func (_BridgeLocker *BridgeLockerCallerSession) CalculateFee(amount *big.Int) (*big.Int, error) {
	return _BridgeLocker.Contract.CalculateFee(&_BridgeLocker.CallOpts, amount)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeLocker *BridgeLockerCaller) GetFee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeLocker.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeLocker *BridgeLockerSession) GetFee() (common.Address, error) {
	return _BridgeLocker.Contract.GetFee(&_BridgeLocker.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(address)
func (_BridgeLocker *BridgeLockerCallerSession) GetFee() (common.Address, error) {
	return _BridgeLocker.Contract.GetFee(&_BridgeLocker.CallOpts)
}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeLocker *BridgeLockerCaller) GetLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeLocker.contract.Call(opts, &out, "getLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeLocker *BridgeLockerSession) GetLimiter() (common.Address, error) {
	return _BridgeLocker.Contract.GetLimiter(&_BridgeLocker.CallOpts)
}

// GetLimiter is a free data retrieval call binding the contract method 0xeb2a0d1f.
//
// Solidity: function getLimiter() view returns(address)
func (_BridgeLocker *BridgeLockerCallerSession) GetLimiter() (common.Address, error) {
	return _BridgeLocker.Contract.GetLimiter(&_BridgeLocker.CallOpts)
}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeLocker *BridgeLockerCaller) GetLimiterUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeLocker.contract.Call(opts, &out, "getLimiterUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeLocker *BridgeLockerSession) GetLimiterUsage() (*big.Int, error) {
	return _BridgeLocker.Contract.GetLimiterUsage(&_BridgeLocker.CallOpts)
}

// GetLimiterUsage is a free data retrieval call binding the contract method 0x7eb76b29.
//
// Solidity: function getLimiterUsage() view returns(uint256)
func (_BridgeLocker *BridgeLockerCallerSession) GetLimiterUsage() (*big.Int, error) {
	return _BridgeLocker.Contract.GetLimiterUsage(&_BridgeLocker.CallOpts)
}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeLocker *BridgeLockerCaller) IsLimited(opts *bind.CallOpts, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _BridgeLocker.contract.Call(opts, &out, "isLimited", amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeLocker *BridgeLockerSession) IsLimited(amount *big.Int) (bool, error) {
	return _BridgeLocker.Contract.IsLimited(&_BridgeLocker.CallOpts, amount)
}

// IsLimited is a free data retrieval call binding the contract method 0x08a90d5a.
//
// Solidity: function isLimited(uint256 amount) view returns(bool)
func (_BridgeLocker *BridgeLockerCallerSession) IsLimited(amount *big.Int) (bool, error) {
	return _BridgeLocker.Contract.IsLimited(&_BridgeLocker.CallOpts, amount)
}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeLocker *BridgeLockerCaller) IsUnlockCompleted(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _BridgeLocker.contract.Call(opts, &out, "isUnlockCompleted", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeLocker *BridgeLockerSession) IsUnlockCompleted(hash [32]byte) (bool, error) {
	return _BridgeLocker.Contract.IsUnlockCompleted(&_BridgeLocker.CallOpts, hash)
}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_BridgeLocker *BridgeLockerCallerSession) IsUnlockCompleted(hash [32]byte) (bool, error) {
	return _BridgeLocker.Contract.IsUnlockCompleted(&_BridgeLocker.CallOpts, hash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeLocker *BridgeLockerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeLocker.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeLocker *BridgeLockerSession) Name() (string, error) {
	return _BridgeLocker.Contract.Name(&_BridgeLocker.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeLocker *BridgeLockerCallerSession) Name() (string, error) {
	return _BridgeLocker.Contract.Name(&_BridgeLocker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeLocker *BridgeLockerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeLocker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeLocker *BridgeLockerSession) Owner() (common.Address, error) {
	return _BridgeLocker.Contract.Owner(&_BridgeLocker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeLocker *BridgeLockerCallerSession) Owner() (common.Address, error) {
	return _BridgeLocker.Contract.Owner(&_BridgeLocker.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeLocker *BridgeLockerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeLocker.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeLocker *BridgeLockerSession) Paused() (bool, error) {
	return _BridgeLocker.Contract.Paused(&_BridgeLocker.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeLocker *BridgeLockerCallerSession) Paused() (bool, error) {
	return _BridgeLocker.Contract.Paused(&_BridgeLocker.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeLocker *BridgeLockerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeLocker.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeLocker *BridgeLockerSession) Token() (common.Address, error) {
	return _BridgeLocker.Contract.Token(&_BridgeLocker.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeLocker *BridgeLockerCallerSession) Token() (common.Address, error) {
	return _BridgeLocker.Contract.Token(&_BridgeLocker.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeLocker *BridgeLockerTransactor) Lock(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BridgeLocker.contract.Transact(opts, "lock", amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeLocker *BridgeLockerSession) Lock(amount *big.Int) (*types.Transaction, error) {
	return _BridgeLocker.Contract.Lock(&_BridgeLocker.TransactOpts, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_BridgeLocker *BridgeLockerTransactorSession) Lock(amount *big.Int) (*types.Transaction, error) {
	return _BridgeLocker.Contract.Lock(&_BridgeLocker.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeLocker *BridgeLockerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLocker.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeLocker *BridgeLockerSession) Pause() (*types.Transaction, error) {
	return _BridgeLocker.Contract.Pause(&_BridgeLocker.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeLocker *BridgeLockerTransactorSession) Pause() (*types.Transaction, error) {
	return _BridgeLocker.Contract.Pause(&_BridgeLocker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeLocker *BridgeLockerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLocker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeLocker *BridgeLockerSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeLocker.Contract.RenounceOwnership(&_BridgeLocker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeLocker *BridgeLockerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeLocker.Contract.RenounceOwnership(&_BridgeLocker.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeLocker *BridgeLockerTransactor) SetFee(opts *bind.TransactOpts, fee_ common.Address) (*types.Transaction, error) {
	return _BridgeLocker.contract.Transact(opts, "setFee", fee_)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeLocker *BridgeLockerSession) SetFee(fee_ common.Address) (*types.Transaction, error) {
	return _BridgeLocker.Contract.SetFee(&_BridgeLocker.TransactOpts, fee_)
}

// SetFee is a paid mutator transaction binding the contract method 0x7917fb9f.
//
// Solidity: function setFee(address fee_) returns()
func (_BridgeLocker *BridgeLockerTransactorSession) SetFee(fee_ common.Address) (*types.Transaction, error) {
	return _BridgeLocker.Contract.SetFee(&_BridgeLocker.TransactOpts, fee_)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeLocker *BridgeLockerTransactor) SetLimiter(opts *bind.TransactOpts, limiter common.Address) (*types.Transaction, error) {
	return _BridgeLocker.contract.Transact(opts, "setLimiter", limiter)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeLocker *BridgeLockerSession) SetLimiter(limiter common.Address) (*types.Transaction, error) {
	return _BridgeLocker.Contract.SetLimiter(&_BridgeLocker.TransactOpts, limiter)
}

// SetLimiter is a paid mutator transaction binding the contract method 0x7a29084c.
//
// Solidity: function setLimiter(address limiter) returns()
func (_BridgeLocker *BridgeLockerTransactorSession) SetLimiter(limiter common.Address) (*types.Transaction, error) {
	return _BridgeLocker.Contract.SetLimiter(&_BridgeLocker.TransactOpts, limiter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeLocker *BridgeLockerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeLocker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeLocker *BridgeLockerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeLocker.Contract.TransferOwnership(&_BridgeLocker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeLocker *BridgeLockerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeLocker.Contract.TransferOwnership(&_BridgeLocker.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeLocker *BridgeLockerTransactor) Unlock(opts *bind.TransactOpts, account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeLocker.contract.Transact(opts, "unlock", account, amount, hash)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeLocker *BridgeLockerSession) Unlock(account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeLocker.Contract.Unlock(&_BridgeLocker.TransactOpts, account, amount, hash)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_BridgeLocker *BridgeLockerTransactorSession) Unlock(account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _BridgeLocker.Contract.Unlock(&_BridgeLocker.TransactOpts, account, amount, hash)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeLocker *BridgeLockerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLocker.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeLocker *BridgeLockerSession) Unpause() (*types.Transaction, error) {
	return _BridgeLocker.Contract.Unpause(&_BridgeLocker.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeLocker *BridgeLockerTransactorSession) Unpause() (*types.Transaction, error) {
	return _BridgeLocker.Contract.Unpause(&_BridgeLocker.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeLocker *BridgeLockerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLocker.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeLocker *BridgeLockerSession) Receive() (*types.Transaction, error) {
	return _BridgeLocker.Contract.Receive(&_BridgeLocker.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeLocker *BridgeLockerTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeLocker.Contract.Receive(&_BridgeLocker.TransactOpts)
}

// BridgeLockerLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the BridgeLocker contract.
type BridgeLockerLockedIterator struct {
	Event *BridgeLockerLocked // Event containing the contract specifics and raw log

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
func (it *BridgeLockerLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLockerLocked)
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
		it.Event = new(BridgeLockerLocked)
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
func (it *BridgeLockerLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLockerLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLockerLocked represents a Locked event raised by the BridgeLocker contract.
type BridgeLockerLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeLocker *BridgeLockerFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*BridgeLockerLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeLocker.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLockerLockedIterator{contract: _BridgeLocker.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeLocker *BridgeLockerFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *BridgeLockerLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeLocker.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLockerLocked)
				if err := _BridgeLocker.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_BridgeLocker *BridgeLockerFilterer) ParseLocked(log types.Log) (*BridgeLockerLocked, error) {
	event := new(BridgeLockerLocked)
	if err := _BridgeLocker.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLockerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeLocker contract.
type BridgeLockerOwnershipTransferredIterator struct {
	Event *BridgeLockerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeLockerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLockerOwnershipTransferred)
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
		it.Event = new(BridgeLockerOwnershipTransferred)
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
func (it *BridgeLockerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLockerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLockerOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeLocker contract.
type BridgeLockerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeLocker *BridgeLockerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeLockerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeLocker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLockerOwnershipTransferredIterator{contract: _BridgeLocker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeLocker *BridgeLockerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeLockerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeLocker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLockerOwnershipTransferred)
				if err := _BridgeLocker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeLocker *BridgeLockerFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeLockerOwnershipTransferred, error) {
	event := new(BridgeLockerOwnershipTransferred)
	if err := _BridgeLocker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLockerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BridgeLocker contract.
type BridgeLockerPausedIterator struct {
	Event *BridgeLockerPaused // Event containing the contract specifics and raw log

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
func (it *BridgeLockerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLockerPaused)
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
		it.Event = new(BridgeLockerPaused)
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
func (it *BridgeLockerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLockerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLockerPaused represents a Paused event raised by the BridgeLocker contract.
type BridgeLockerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeLocker *BridgeLockerFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgeLockerPausedIterator, error) {

	logs, sub, err := _BridgeLocker.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgeLockerPausedIterator{contract: _BridgeLocker.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeLocker *BridgeLockerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgeLockerPaused) (event.Subscription, error) {

	logs, sub, err := _BridgeLocker.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLockerPaused)
				if err := _BridgeLocker.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address account)
func (_BridgeLocker *BridgeLockerFilterer) ParsePaused(log types.Log) (*BridgeLockerPaused, error) {
	event := new(BridgeLockerPaused)
	if err := _BridgeLocker.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLockerUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the BridgeLocker contract.
type BridgeLockerUnlockedIterator struct {
	Event *BridgeLockerUnlocked // Event containing the contract specifics and raw log

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
func (it *BridgeLockerUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLockerUnlocked)
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
		it.Event = new(BridgeLockerUnlocked)
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
func (it *BridgeLockerUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLockerUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLockerUnlocked represents a Unlocked event raised by the BridgeLocker contract.
type BridgeLockerUnlocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeLocker *BridgeLockerFilterer) FilterUnlocked(opts *bind.FilterOpts, sender []common.Address) (*BridgeLockerUnlockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeLocker.contract.FilterLogs(opts, "Unlocked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLockerUnlockedIterator{contract: _BridgeLocker.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeLocker *BridgeLockerFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *BridgeLockerUnlocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeLocker.contract.WatchLogs(opts, "Unlocked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLockerUnlocked)
				if err := _BridgeLocker.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_BridgeLocker *BridgeLockerFilterer) ParseUnlocked(log types.Log) (*BridgeLockerUnlocked, error) {
	event := new(BridgeLockerUnlocked)
	if err := _BridgeLocker.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLockerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BridgeLocker contract.
type BridgeLockerUnpausedIterator struct {
	Event *BridgeLockerUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeLockerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLockerUnpaused)
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
		it.Event = new(BridgeLockerUnpaused)
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
func (it *BridgeLockerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLockerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLockerUnpaused represents a Unpaused event raised by the BridgeLocker contract.
type BridgeLockerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeLocker *BridgeLockerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeLockerUnpausedIterator, error) {

	logs, sub, err := _BridgeLocker.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeLockerUnpausedIterator{contract: _BridgeLocker.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeLocker *BridgeLockerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeLockerUnpaused) (event.Subscription, error) {

	logs, sub, err := _BridgeLocker.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLockerUnpaused)
				if err := _BridgeLocker.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address account)
func (_BridgeLocker *BridgeLockerFilterer) ParseUnpaused(log types.Log) (*BridgeLockerUnpaused, error) {
	event := new(BridgeLockerUnpaused)
	if err := _BridgeLocker.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x60806040523480156200001157600080fd5b5060405162000b7538038062000b758339810160408190526200003491620001c1565b81516200004990600390602085019062000068565b5080516200005f90600490602084019062000068565b5050506200027b565b828054620000769062000228565b90600052602060002090601f0160209004810192826200009a5760008555620000e5565b82601f10620000b557805160ff1916838001178555620000e5565b82800160010185558215620000e5579182015b82811115620000e5578251825591602001919060010190620000c8565b50620000f3929150620000f7565b5090565b5b80821115620000f35760008155600101620000f8565b600082601f8301126200011f578081fd5b81516001600160401b03808211156200013c576200013c62000265565b604051601f8301601f19908116603f0116810190828211818310171562000167576200016762000265565b8160405283815260209250868385880101111562000183578485fd5b8491505b83821015620001a6578582018301518183018401529082019062000187565b83821115620001b757848385830101525b9695505050505050565b60008060408385031215620001d4578182fd5b82516001600160401b0380821115620001eb578384fd5b620001f9868387016200010e565b935060208501519150808211156200020f578283fd5b506200021e858286016200010e565b9150509250929050565b600181811c908216806200023d57607f821691505b602082108114156200025f57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6108ea806200028b6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012357806370a082311461013657806395d89b411461015f578063a457c2d714610167578063a9059cbb1461017a578063dd62ed3e1461018d57600080fd5b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101c6565b6040516100c391906107e1565b60405180910390f35b6100df6100da3660046107b8565b610258565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f36600461077d565b61026e565b604051601281526020016100c3565b6100df6101313660046107b8565b610324565b6100f361014436600461072a565b6001600160a01b031660009081526020819052604090205490565b6100b661035b565b6100df6101753660046107b8565b61036a565b6100df6101883660046107b8565b610405565b6100f361019b36600461074b565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101d590610863565b80601f016020809104026020016040519081016040528092919081815260200182805461020190610863565b801561024e5780601f106102235761010080835404028352916020019161024e565b820191906000526020600020905b81548152906001019060200180831161023157829003601f168201915b5050505050905090565b6000610265338484610412565b50600192915050565b600061027b848484610536565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156103055760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b6103198533610314868561084c565b610412565b506001949350505050565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091610265918590610314908690610834565b6060600480546101d590610863565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156103ec5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016102fc565b6103fb3385610314868561084c565b5060019392505050565b6000610265338484610536565b6001600160a01b0383166104745760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016102fc565b6001600160a01b0382166104d55760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016102fc565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b03831661059a5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016102fc565b6001600160a01b0382166105fc5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016102fc565b6001600160a01b038316600090815260208190526040902054818110156106745760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016102fc565b61067e828261084c565b6001600160a01b0380861660009081526020819052604080822093909355908516815290812080548492906106b4908490610834565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161070091815260200190565b60405180910390a350505050565b80356001600160a01b038116811461072557600080fd5b919050565b60006020828403121561073b578081fd5b6107448261070e565b9392505050565b6000806040838503121561075d578081fd5b6107668361070e565b91506107746020840161070e565b90509250929050565b600080600060608486031215610791578081fd5b61079a8461070e565b92506107a86020850161070e565b9150604084013590509250925092565b600080604083850312156107ca578182fd5b6107d38361070e565b946020939093013593505050565b6000602080835283518082850152825b8181101561080d578581018301518582016040015282016107f1565b8181111561081e5783604083870101525b50601f01601f1916929092016040019392505050565b600082198211156108475761084761089e565b500190565b60008282101561085e5761085e61089e565b500390565b600181811c9082168061087757607f821691505b6020821081141561089857634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220ba1e841333051c978f5a22251e0147735129e3f406b38e4baf7b154110b85b9564736f6c63430008040033"

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableABI is the input ABI used to generate the binding from.
const ERC20BurnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20BurnableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BurnableFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"79cc6790": "burnFrom(address,uint256)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Burnable is an auto generated Go binding around an Ethereum contract.
type ERC20Burnable struct {
	ERC20BurnableCaller     // Read-only binding to the contract
	ERC20BurnableTransactor // Write-only binding to the contract
	ERC20BurnableFilterer   // Log filterer for contract events
}

// ERC20BurnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BurnableSession struct {
	Contract     *ERC20Burnable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BurnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BurnableCallerSession struct {
	Contract *ERC20BurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20BurnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BurnableTransactorSession struct {
	Contract     *ERC20BurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20BurnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BurnableRaw struct {
	Contract *ERC20Burnable // Generic contract binding to access the raw methods on
}

// ERC20BurnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BurnableCallerRaw struct {
	Contract *ERC20BurnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BurnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactorRaw struct {
	Contract *ERC20BurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Burnable creates a new instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20Burnable(address common.Address, backend bind.ContractBackend) (*ERC20Burnable, error) {
	contract, err := bindERC20Burnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Burnable{ERC20BurnableCaller: ERC20BurnableCaller{contract: contract}, ERC20BurnableTransactor: ERC20BurnableTransactor{contract: contract}, ERC20BurnableFilterer: ERC20BurnableFilterer{contract: contract}}, nil
}

// NewERC20BurnableCaller creates a new read-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20BurnableCaller, error) {
	contract, err := bindERC20Burnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableCaller{contract: contract}, nil
}

// NewERC20BurnableTransactor creates a new write-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BurnableTransactor, error) {
	contract, err := bindERC20Burnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransactor{contract: contract}, nil
}

// NewERC20BurnableFilterer creates a new log filterer instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BurnableFilterer, error) {
	contract, err := bindERC20Burnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableFilterer{contract: contract}, nil
}

// bindERC20Burnable binds a generic wrapper to an already deployed contract.
func bindERC20Burnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BurnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.ERC20BurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCallerSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// ERC20BurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Burnable contract.
type ERC20BurnableApprovalIterator struct {
	Event *ERC20BurnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableApproval)
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
		it.Event = new(ERC20BurnableApproval)
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
func (it *ERC20BurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableApproval represents a Approval event raised by the ERC20Burnable contract.
type ERC20BurnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20BurnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableApprovalIterator{contract: _ERC20Burnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BurnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableApproval)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20Burnable *ERC20BurnableFilterer) ParseApproval(log types.Log) (*ERC20BurnableApproval, error) {
	event := new(ERC20BurnableApproval)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Burnable contract.
type ERC20BurnableTransferIterator struct {
	Event *ERC20BurnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableTransfer)
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
		it.Event = new(ERC20BurnableTransfer)
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
func (it *ERC20BurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableTransfer represents a Transfer event raised by the ERC20Burnable contract.
type ERC20BurnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BurnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransferIterator{contract: _ERC20Burnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BurnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableTransfer)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Burnable *ERC20BurnableFilterer) ParseTransfer(log types.Log) (*ERC20BurnableTransfer, error) {
	event := new(ERC20BurnableTransfer)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeFixedABI is the input ABI used to generate the binding from.
const FeeFixedABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee_\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeeFixedFuncSigs maps the 4-byte function signature to its string representation.
var FeeFixedFuncSigs = map[string]string{
	"b9d92de8": "calculate(uint256)",
	"ddca3f43": "fee()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"69fe0e2d": "setFee(uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// FeeFixedBin is the compiled bytecode used for deploying new contracts.
var FeeFixedBin = "0x608060405234801561001057600080fd5b5060405161043238038061043283398101604081905261002f916100e2565b600080546001600160a01b031916339081178255604051909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506100798161007f565b506100fa565b6000546001600160a01b031633146100dd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640160405180910390fd5b600155565b6000602082840312156100f3578081fd5b5051919050565b610329806101096000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806369fe0e2d14610067578063715018a61461007c5780638da5cb5b14610084578063b9d92de8146100a4578063ddca3f43146100c7578063f2fde38b146100cf575b600080fd5b61007a6100753660046102a6565b6100e2565b005b61007a61011a565b6000546040516001600160a01b0390911681526020015b60405180910390f35b6100b96100b23660046102a6565b5060015490565b60405190815260200161009b565b6001546100b9565b61007a6100dd366004610278565b61018e565b6000546001600160a01b031633146101155760405162461bcd60e51b815260040161010c906102be565b60405180910390fd5b600155565b6000546001600160a01b031633146101445760405162461bcd60e51b815260040161010c906102be565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031633146101b85760405162461bcd60e51b815260040161010c906102be565b6001600160a01b03811661021d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161010c565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600060208284031215610289578081fd5b81356001600160a01b038116811461029f578182fd5b9392505050565b6000602082840312156102b7578081fd5b5035919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260408201526060019056fea26469706673582212208ce3c077efd27af6823dc02cd183b98b7e0998a97dca72befeb2dca108d4719864736f6c63430008040033"

// DeployFeeFixed deploys a new Ethereum contract, binding an instance of FeeFixed to it.
func DeployFeeFixed(auth *bind.TransactOpts, backend bind.ContractBackend, fee_ *big.Int) (common.Address, *types.Transaction, *FeeFixed, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeFixedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeeFixedBin), backend, fee_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeFixed{FeeFixedCaller: FeeFixedCaller{contract: contract}, FeeFixedTransactor: FeeFixedTransactor{contract: contract}, FeeFixedFilterer: FeeFixedFilterer{contract: contract}}, nil
}

// FeeFixed is an auto generated Go binding around an Ethereum contract.
type FeeFixed struct {
	FeeFixedCaller     // Read-only binding to the contract
	FeeFixedTransactor // Write-only binding to the contract
	FeeFixedFilterer   // Log filterer for contract events
}

// FeeFixedCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeFixedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeFixedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeFixedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeFixedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeFixedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeFixedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeFixedSession struct {
	Contract     *FeeFixed         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeFixedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeFixedCallerSession struct {
	Contract *FeeFixedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FeeFixedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeFixedTransactorSession struct {
	Contract     *FeeFixedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FeeFixedRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeFixedRaw struct {
	Contract *FeeFixed // Generic contract binding to access the raw methods on
}

// FeeFixedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeFixedCallerRaw struct {
	Contract *FeeFixedCaller // Generic read-only contract binding to access the raw methods on
}

// FeeFixedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeFixedTransactorRaw struct {
	Contract *FeeFixedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeFixed creates a new instance of FeeFixed, bound to a specific deployed contract.
func NewFeeFixed(address common.Address, backend bind.ContractBackend) (*FeeFixed, error) {
	contract, err := bindFeeFixed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeFixed{FeeFixedCaller: FeeFixedCaller{contract: contract}, FeeFixedTransactor: FeeFixedTransactor{contract: contract}, FeeFixedFilterer: FeeFixedFilterer{contract: contract}}, nil
}

// NewFeeFixedCaller creates a new read-only instance of FeeFixed, bound to a specific deployed contract.
func NewFeeFixedCaller(address common.Address, caller bind.ContractCaller) (*FeeFixedCaller, error) {
	contract, err := bindFeeFixed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeFixedCaller{contract: contract}, nil
}

// NewFeeFixedTransactor creates a new write-only instance of FeeFixed, bound to a specific deployed contract.
func NewFeeFixedTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeFixedTransactor, error) {
	contract, err := bindFeeFixed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeFixedTransactor{contract: contract}, nil
}

// NewFeeFixedFilterer creates a new log filterer instance of FeeFixed, bound to a specific deployed contract.
func NewFeeFixedFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeFixedFilterer, error) {
	contract, err := bindFeeFixed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeFixedFilterer{contract: contract}, nil
}

// bindFeeFixed binds a generic wrapper to an already deployed contract.
func bindFeeFixed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeFixedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeFixed *FeeFixedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeFixed.Contract.FeeFixedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeFixed *FeeFixedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeFixed.Contract.FeeFixedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeFixed *FeeFixedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeFixed.Contract.FeeFixedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeFixed *FeeFixedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeFixed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeFixed *FeeFixedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeFixed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeFixed *FeeFixedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeFixed.Contract.contract.Transact(opts, method, params...)
}

// Calculate is a free data retrieval call binding the contract method 0xb9d92de8.
//
// Solidity: function calculate(uint256 ) view returns(uint256)
func (_FeeFixed *FeeFixedCaller) Calculate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeFixed.contract.Call(opts, &out, "calculate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Calculate is a free data retrieval call binding the contract method 0xb9d92de8.
//
// Solidity: function calculate(uint256 ) view returns(uint256)
func (_FeeFixed *FeeFixedSession) Calculate(arg0 *big.Int) (*big.Int, error) {
	return _FeeFixed.Contract.Calculate(&_FeeFixed.CallOpts, arg0)
}

// Calculate is a free data retrieval call binding the contract method 0xb9d92de8.
//
// Solidity: function calculate(uint256 ) view returns(uint256)
func (_FeeFixed *FeeFixedCallerSession) Calculate(arg0 *big.Int) (*big.Int, error) {
	return _FeeFixed.Contract.Calculate(&_FeeFixed.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_FeeFixed *FeeFixedCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeFixed.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_FeeFixed *FeeFixedSession) Fee() (*big.Int, error) {
	return _FeeFixed.Contract.Fee(&_FeeFixed.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_FeeFixed *FeeFixedCallerSession) Fee() (*big.Int, error) {
	return _FeeFixed.Contract.Fee(&_FeeFixed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeFixed *FeeFixedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeFixed.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeFixed *FeeFixedSession) Owner() (common.Address, error) {
	return _FeeFixed.Contract.Owner(&_FeeFixed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeFixed *FeeFixedCallerSession) Owner() (common.Address, error) {
	return _FeeFixed.Contract.Owner(&_FeeFixed.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeFixed *FeeFixedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeFixed.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeFixed *FeeFixedSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeFixed.Contract.RenounceOwnership(&_FeeFixed.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeFixed *FeeFixedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeFixed.Contract.RenounceOwnership(&_FeeFixed.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee_) returns()
func (_FeeFixed *FeeFixedTransactor) SetFee(opts *bind.TransactOpts, fee_ *big.Int) (*types.Transaction, error) {
	return _FeeFixed.contract.Transact(opts, "setFee", fee_)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee_) returns()
func (_FeeFixed *FeeFixedSession) SetFee(fee_ *big.Int) (*types.Transaction, error) {
	return _FeeFixed.Contract.SetFee(&_FeeFixed.TransactOpts, fee_)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee_) returns()
func (_FeeFixed *FeeFixedTransactorSession) SetFee(fee_ *big.Int) (*types.Transaction, error) {
	return _FeeFixed.Contract.SetFee(&_FeeFixed.TransactOpts, fee_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeFixed *FeeFixedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeeFixed.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeFixed *FeeFixedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeFixed.Contract.TransferOwnership(&_FeeFixed.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeFixed *FeeFixedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeFixed.Contract.TransferOwnership(&_FeeFixed.TransactOpts, newOwner)
}

// FeeFixedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeeFixed contract.
type FeeFixedOwnershipTransferredIterator struct {
	Event *FeeFixedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeeFixedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeFixedOwnershipTransferred)
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
		it.Event = new(FeeFixedOwnershipTransferred)
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
func (it *FeeFixedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeFixedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeFixedOwnershipTransferred represents a OwnershipTransferred event raised by the FeeFixed contract.
type FeeFixedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeFixed *FeeFixedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeeFixedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeFixed.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeFixedOwnershipTransferredIterator{contract: _FeeFixed.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeFixed *FeeFixedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeFixedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeFixed.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeFixedOwnershipTransferred)
				if err := _FeeFixed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeFixed *FeeFixedFilterer) ParseOwnershipTransferred(log types.Log) (*FeeFixedOwnershipTransferred, error) {
	event := new(FeeFixedOwnershipTransferred)
	if err := _FeeFixed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeABI is the input ABI used to generate the binding from.
const IBridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isUnlockCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBridgeFuncSigs maps the 4-byte function signature to its string representation.
var IBridgeFuncSigs = map[string]string{
	"a4d7fa93": "isUnlockCompleted(bytes32)",
	"dd467064": "lock(uint256)",
	"b322edea": "unlock(address,uint256,bytes32)",
}

// IBridge is an auto generated Go binding around an Ethereum contract.
type IBridge struct {
	IBridgeCaller     // Read-only binding to the contract
	IBridgeTransactor // Write-only binding to the contract
	IBridgeFilterer   // Log filterer for contract events
}

// IBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeSession struct {
	Contract     *IBridge          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeCallerSession struct {
	Contract *IBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeTransactorSession struct {
	Contract     *IBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeRaw struct {
	Contract *IBridge // Generic contract binding to access the raw methods on
}

// IBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeCallerRaw struct {
	Contract *IBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeTransactorRaw struct {
	Contract *IBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridge creates a new instance of IBridge, bound to a specific deployed contract.
func NewIBridge(address common.Address, backend bind.ContractBackend) (*IBridge, error) {
	contract, err := bindIBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridge{IBridgeCaller: IBridgeCaller{contract: contract}, IBridgeTransactor: IBridgeTransactor{contract: contract}, IBridgeFilterer: IBridgeFilterer{contract: contract}}, nil
}

// NewIBridgeCaller creates a new read-only instance of IBridge, bound to a specific deployed contract.
func NewIBridgeCaller(address common.Address, caller bind.ContractCaller) (*IBridgeCaller, error) {
	contract, err := bindIBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeCaller{contract: contract}, nil
}

// NewIBridgeTransactor creates a new write-only instance of IBridge, bound to a specific deployed contract.
func NewIBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeTransactor, error) {
	contract, err := bindIBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeTransactor{contract: contract}, nil
}

// NewIBridgeFilterer creates a new log filterer instance of IBridge, bound to a specific deployed contract.
func NewIBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeFilterer, error) {
	contract, err := bindIBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeFilterer{contract: contract}, nil
}

// bindIBridge binds a generic wrapper to an already deployed contract.
func bindIBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridge *IBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridge.Contract.IBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridge *IBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.Contract.IBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridge *IBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridge.Contract.IBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridge *IBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridge *IBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridge *IBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridge.Contract.contract.Transact(opts, method, params...)
}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_IBridge *IBridgeCaller) IsUnlockCompleted(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "isUnlockCompleted", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_IBridge *IBridgeSession) IsUnlockCompleted(hash [32]byte) (bool, error) {
	return _IBridge.Contract.IsUnlockCompleted(&_IBridge.CallOpts, hash)
}

// IsUnlockCompleted is a free data retrieval call binding the contract method 0xa4d7fa93.
//
// Solidity: function isUnlockCompleted(bytes32 hash) view returns(bool)
func (_IBridge *IBridgeCallerSession) IsUnlockCompleted(hash [32]byte) (bool, error) {
	return _IBridge.Contract.IsUnlockCompleted(&_IBridge.CallOpts, hash)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_IBridge *IBridgeTransactor) Lock(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "lock", amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_IBridge *IBridgeSession) Lock(amount *big.Int) (*types.Transaction, error) {
	return _IBridge.Contract.Lock(&_IBridge.TransactOpts, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 amount) payable returns()
func (_IBridge *IBridgeTransactorSession) Lock(amount *big.Int) (*types.Transaction, error) {
	return _IBridge.Contract.Lock(&_IBridge.TransactOpts, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_IBridge *IBridgeTransactor) Unlock(opts *bind.TransactOpts, account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "unlock", account, amount, hash)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_IBridge *IBridgeSession) Unlock(account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _IBridge.Contract.Unlock(&_IBridge.TransactOpts, account, amount, hash)
}

// Unlock is a paid mutator transaction binding the contract method 0xb322edea.
//
// Solidity: function unlock(address account, uint256 amount, bytes32 hash) returns()
func (_IBridge *IBridgeTransactorSession) Unlock(account common.Address, amount *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _IBridge.Contract.Unlock(&_IBridge.TransactOpts, account, amount, hash)
}

// IBridgeLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the IBridge contract.
type IBridgeLockedIterator struct {
	Event *IBridgeLocked // Event containing the contract specifics and raw log

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
func (it *IBridgeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeLocked)
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
		it.Event = new(IBridgeLocked)
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
func (it *IBridgeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeLocked represents a Locked event raised by the IBridge contract.
type IBridgeLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_IBridge *IBridgeFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*IBridgeLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &IBridgeLockedIterator{contract: _IBridge.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_IBridge *IBridgeFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *IBridgeLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeLocked)
				if err := _IBridge.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008.
//
// Solidity: event Locked(address indexed sender, uint256 amount)
func (_IBridge *IBridgeFilterer) ParseLocked(log types.Log) (*IBridgeLocked, error) {
	event := new(IBridgeLocked)
	if err := _IBridge.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the IBridge contract.
type IBridgeUnlockedIterator struct {
	Event *IBridgeUnlocked // Event containing the contract specifics and raw log

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
func (it *IBridgeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeUnlocked)
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
		it.Event = new(IBridgeUnlocked)
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
func (it *IBridgeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeUnlocked represents a Unlocked event raised by the IBridge contract.
type IBridgeUnlocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_IBridge *IBridgeFilterer) FilterUnlocked(opts *bind.FilterOpts, sender []common.Address) (*IBridgeUnlockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "Unlocked", senderRule)
	if err != nil {
		return nil, err
	}
	return &IBridgeUnlockedIterator{contract: _IBridge.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_IBridge *IBridgeFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *IBridgeUnlocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "Unlocked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeUnlocked)
				if err := _IBridge.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed sender, uint256 amount)
func (_IBridge *IBridgeFilterer) ParseUnlocked(log types.Log) (*IBridgeUnlocked, error) {
	event := new(IBridgeUnlocked)
	if err := _IBridge.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataABI is the input ABI used to generate the binding from.
const IERC20MetadataABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20MetadataFuncSigs maps the 4-byte function signature to its string representation.
var IERC20MetadataFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20Metadata is an auto generated Go binding around an Ethereum contract.
type IERC20Metadata struct {
	IERC20MetadataCaller     // Read-only binding to the contract
	IERC20MetadataTransactor // Write-only binding to the contract
	IERC20MetadataFilterer   // Log filterer for contract events
}

// IERC20MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MetadataSession struct {
	Contract     *IERC20Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MetadataCallerSession struct {
	Contract *IERC20MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MetadataTransactorSession struct {
	Contract     *IERC20MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MetadataRaw struct {
	Contract *IERC20Metadata // Generic contract binding to access the raw methods on
}

// IERC20MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MetadataCallerRaw struct {
	Contract *IERC20MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactorRaw struct {
	Contract *IERC20MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Metadata creates a new instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20Metadata(address common.Address, backend bind.ContractBackend) (*IERC20Metadata, error) {
	contract, err := bindIERC20Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Metadata{IERC20MetadataCaller: IERC20MetadataCaller{contract: contract}, IERC20MetadataTransactor: IERC20MetadataTransactor{contract: contract}, IERC20MetadataFilterer: IERC20MetadataFilterer{contract: contract}}, nil
}

// NewIERC20MetadataCaller creates a new read-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC20MetadataCaller, error) {
	contract, err := bindIERC20Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataCaller{contract: contract}, nil
}

// NewIERC20MetadataTransactor creates a new write-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MetadataTransactor, error) {
	contract, err := bindIERC20Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransactor{contract: contract}, nil
}

// NewIERC20MetadataFilterer creates a new log filterer instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MetadataFilterer, error) {
	contract, err := bindIERC20Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataFilterer{contract: contract}, nil
}

// bindIERC20Metadata binds a generic wrapper to an already deployed contract.
func bindIERC20Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.IERC20MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCallerSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, sender, recipient, amount)
}

// IERC20MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Metadata contract.
type IERC20MetadataApprovalIterator struct {
	Event *IERC20MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataApproval)
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
		it.Event = new(IERC20MetadataApproval)
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
func (it *IERC20MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataApproval represents a Approval event raised by the IERC20Metadata contract.
type IERC20MetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataApprovalIterator{contract: _IERC20Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MetadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataApproval)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseApproval(log types.Log) (*IERC20MetadataApproval, error) {
	event := new(IERC20MetadataApproval)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Metadata contract.
type IERC20MetadataTransferIterator struct {
	Event *IERC20MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataTransfer)
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
		it.Event = new(IERC20MetadataTransfer)
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
func (it *IERC20MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataTransfer represents a Transfer event raised by the IERC20Metadata contract.
type IERC20MetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransferIterator{contract: _IERC20Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MetadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataTransfer)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseTransfer(log types.Log) (*IERC20MetadataTransfer, error) {
	event := new(IERC20MetadataTransfer)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFeeABI is the input ABI used to generate the binding from.
const IFeeABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IFeeFuncSigs maps the 4-byte function signature to its string representation.
var IFeeFuncSigs = map[string]string{
	"b9d92de8": "calculate(uint256)",
}

// IFee is an auto generated Go binding around an Ethereum contract.
type IFee struct {
	IFeeCaller     // Read-only binding to the contract
	IFeeTransactor // Write-only binding to the contract
	IFeeFilterer   // Log filterer for contract events
}

// IFeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFeeSession struct {
	Contract     *IFee             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFeeCallerSession struct {
	Contract *IFeeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IFeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFeeTransactorSession struct {
	Contract     *IFeeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFeeRaw struct {
	Contract *IFee // Generic contract binding to access the raw methods on
}

// IFeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFeeCallerRaw struct {
	Contract *IFeeCaller // Generic read-only contract binding to access the raw methods on
}

// IFeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFeeTransactorRaw struct {
	Contract *IFeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFee creates a new instance of IFee, bound to a specific deployed contract.
func NewIFee(address common.Address, backend bind.ContractBackend) (*IFee, error) {
	contract, err := bindIFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFee{IFeeCaller: IFeeCaller{contract: contract}, IFeeTransactor: IFeeTransactor{contract: contract}, IFeeFilterer: IFeeFilterer{contract: contract}}, nil
}

// NewIFeeCaller creates a new read-only instance of IFee, bound to a specific deployed contract.
func NewIFeeCaller(address common.Address, caller bind.ContractCaller) (*IFeeCaller, error) {
	contract, err := bindIFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeCaller{contract: contract}, nil
}

// NewIFeeTransactor creates a new write-only instance of IFee, bound to a specific deployed contract.
func NewIFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*IFeeTransactor, error) {
	contract, err := bindIFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeTransactor{contract: contract}, nil
}

// NewIFeeFilterer creates a new log filterer instance of IFee, bound to a specific deployed contract.
func NewIFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*IFeeFilterer, error) {
	contract, err := bindIFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFeeFilterer{contract: contract}, nil
}

// bindIFee binds a generic wrapper to an already deployed contract.
func bindIFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFeeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFee *IFeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFee.Contract.IFeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFee *IFeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFee.Contract.IFeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFee *IFeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFee.Contract.IFeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFee *IFeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFee *IFeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFee *IFeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFee.Contract.contract.Transact(opts, method, params...)
}

// Calculate is a free data retrieval call binding the contract method 0xb9d92de8.
//
// Solidity: function calculate(uint256 amount) view returns(uint256)
func (_IFee *IFeeCaller) Calculate(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IFee.contract.Call(opts, &out, "calculate", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Calculate is a free data retrieval call binding the contract method 0xb9d92de8.
//
// Solidity: function calculate(uint256 amount) view returns(uint256)
func (_IFee *IFeeSession) Calculate(amount *big.Int) (*big.Int, error) {
	return _IFee.Contract.Calculate(&_IFee.CallOpts, amount)
}

// Calculate is a free data retrieval call binding the contract method 0xb9d92de8.
//
// Solidity: function calculate(uint256 amount) view returns(uint256)
func (_IFee *IFeeCallerSession) Calculate(amount *big.Int) (*big.Int, error) {
	return _IFee.Contract.Calculate(&_IFee.CallOpts, amount)
}

// ILimiterABI is the input ABI used to generate the binding from.
const ILimiterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"}],\"name\":\"getLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"}],\"name\":\"getUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"isLimited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ILimiterFuncSigs maps the 4-byte function signature to its string representation.
var ILimiterFuncSigs = map[string]string{
	"1ce28e72": "getLimit(address)",
	"59b9b15e": "getUsage(address)",
	"c0dd9e52": "increaseUsage(uint256)",
	"47618cb2": "isLimited(address,uint256)",
}

// ILimiter is an auto generated Go binding around an Ethereum contract.
type ILimiter struct {
	ILimiterCaller     // Read-only binding to the contract
	ILimiterTransactor // Write-only binding to the contract
	ILimiterFilterer   // Log filterer for contract events
}

// ILimiterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILimiterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILimiterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILimiterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILimiterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILimiterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILimiterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILimiterSession struct {
	Contract     *ILimiter         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILimiterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILimiterCallerSession struct {
	Contract *ILimiterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ILimiterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILimiterTransactorSession struct {
	Contract     *ILimiterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ILimiterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILimiterRaw struct {
	Contract *ILimiter // Generic contract binding to access the raw methods on
}

// ILimiterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILimiterCallerRaw struct {
	Contract *ILimiterCaller // Generic read-only contract binding to access the raw methods on
}

// ILimiterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILimiterTransactorRaw struct {
	Contract *ILimiterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILimiter creates a new instance of ILimiter, bound to a specific deployed contract.
func NewILimiter(address common.Address, backend bind.ContractBackend) (*ILimiter, error) {
	contract, err := bindILimiter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILimiter{ILimiterCaller: ILimiterCaller{contract: contract}, ILimiterTransactor: ILimiterTransactor{contract: contract}, ILimiterFilterer: ILimiterFilterer{contract: contract}}, nil
}

// NewILimiterCaller creates a new read-only instance of ILimiter, bound to a specific deployed contract.
func NewILimiterCaller(address common.Address, caller bind.ContractCaller) (*ILimiterCaller, error) {
	contract, err := bindILimiter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILimiterCaller{contract: contract}, nil
}

// NewILimiterTransactor creates a new write-only instance of ILimiter, bound to a specific deployed contract.
func NewILimiterTransactor(address common.Address, transactor bind.ContractTransactor) (*ILimiterTransactor, error) {
	contract, err := bindILimiter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILimiterTransactor{contract: contract}, nil
}

// NewILimiterFilterer creates a new log filterer instance of ILimiter, bound to a specific deployed contract.
func NewILimiterFilterer(address common.Address, filterer bind.ContractFilterer) (*ILimiterFilterer, error) {
	contract, err := bindILimiter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILimiterFilterer{contract: contract}, nil
}

// bindILimiter binds a generic wrapper to an already deployed contract.
func bindILimiter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILimiterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILimiter *ILimiterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILimiter.Contract.ILimiterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILimiter *ILimiterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILimiter.Contract.ILimiterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILimiter *ILimiterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILimiter.Contract.ILimiterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILimiter *ILimiterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILimiter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILimiter *ILimiterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILimiter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILimiter *ILimiterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILimiter.Contract.contract.Transact(opts, method, params...)
}

// GetLimit is a free data retrieval call binding the contract method 0x1ce28e72.
//
// Solidity: function getLimit(address bridge) view returns(uint256)
func (_ILimiter *ILimiterCaller) GetLimit(opts *bind.CallOpts, bridge common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILimiter.contract.Call(opts, &out, "getLimit", bridge)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimit is a free data retrieval call binding the contract method 0x1ce28e72.
//
// Solidity: function getLimit(address bridge) view returns(uint256)
func (_ILimiter *ILimiterSession) GetLimit(bridge common.Address) (*big.Int, error) {
	return _ILimiter.Contract.GetLimit(&_ILimiter.CallOpts, bridge)
}

// GetLimit is a free data retrieval call binding the contract method 0x1ce28e72.
//
// Solidity: function getLimit(address bridge) view returns(uint256)
func (_ILimiter *ILimiterCallerSession) GetLimit(bridge common.Address) (*big.Int, error) {
	return _ILimiter.Contract.GetLimit(&_ILimiter.CallOpts, bridge)
}

// GetUsage is a free data retrieval call binding the contract method 0x59b9b15e.
//
// Solidity: function getUsage(address bridge) view returns(uint256)
func (_ILimiter *ILimiterCaller) GetUsage(opts *bind.CallOpts, bridge common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILimiter.contract.Call(opts, &out, "getUsage", bridge)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsage is a free data retrieval call binding the contract method 0x59b9b15e.
//
// Solidity: function getUsage(address bridge) view returns(uint256)
func (_ILimiter *ILimiterSession) GetUsage(bridge common.Address) (*big.Int, error) {
	return _ILimiter.Contract.GetUsage(&_ILimiter.CallOpts, bridge)
}

// GetUsage is a free data retrieval call binding the contract method 0x59b9b15e.
//
// Solidity: function getUsage(address bridge) view returns(uint256)
func (_ILimiter *ILimiterCallerSession) GetUsage(bridge common.Address) (*big.Int, error) {
	return _ILimiter.Contract.GetUsage(&_ILimiter.CallOpts, bridge)
}

// IsLimited is a free data retrieval call binding the contract method 0x47618cb2.
//
// Solidity: function isLimited(address bridge, uint256 amount) view returns(bool)
func (_ILimiter *ILimiterCaller) IsLimited(opts *bind.CallOpts, bridge common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _ILimiter.contract.Call(opts, &out, "isLimited", bridge, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLimited is a free data retrieval call binding the contract method 0x47618cb2.
//
// Solidity: function isLimited(address bridge, uint256 amount) view returns(bool)
func (_ILimiter *ILimiterSession) IsLimited(bridge common.Address, amount *big.Int) (bool, error) {
	return _ILimiter.Contract.IsLimited(&_ILimiter.CallOpts, bridge, amount)
}

// IsLimited is a free data retrieval call binding the contract method 0x47618cb2.
//
// Solidity: function isLimited(address bridge, uint256 amount) view returns(bool)
func (_ILimiter *ILimiterCallerSession) IsLimited(bridge common.Address, amount *big.Int) (bool, error) {
	return _ILimiter.Contract.IsLimited(&_ILimiter.CallOpts, bridge, amount)
}

// IncreaseUsage is a paid mutator transaction binding the contract method 0xc0dd9e52.
//
// Solidity: function increaseUsage(uint256 amount) returns()
func (_ILimiter *ILimiterTransactor) IncreaseUsage(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ILimiter.contract.Transact(opts, "increaseUsage", amount)
}

// IncreaseUsage is a paid mutator transaction binding the contract method 0xc0dd9e52.
//
// Solidity: function increaseUsage(uint256 amount) returns()
func (_ILimiter *ILimiterSession) IncreaseUsage(amount *big.Int) (*types.Transaction, error) {
	return _ILimiter.Contract.IncreaseUsage(&_ILimiter.TransactOpts, amount)
}

// IncreaseUsage is a paid mutator transaction binding the contract method 0xc0dd9e52.
//
// Solidity: function increaseUsage(uint256 amount) returns()
func (_ILimiter *ILimiterTransactorSession) IncreaseUsage(amount *big.Int) (*types.Transaction, error) {
	return _ILimiter.Contract.IncreaseUsage(&_ILimiter.TransactOpts, amount)
}

// IWrappedTokenABI is the input ABI used to generate the binding from.
const IWrappedTokenABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWrappedTokenFuncSigs maps the 4-byte function signature to its string representation.
var IWrappedTokenFuncSigs = map[string]string{
	"42966c68": "burn(uint256)",
	"79cc6790": "burnFrom(address,uint256)",
	"40c10f19": "mint(address,uint256)",
}

// IWrappedToken is an auto generated Go binding around an Ethereum contract.
type IWrappedToken struct {
	IWrappedTokenCaller     // Read-only binding to the contract
	IWrappedTokenTransactor // Write-only binding to the contract
	IWrappedTokenFilterer   // Log filterer for contract events
}

// IWrappedTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWrappedTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWrappedTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWrappedTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWrappedTokenSession struct {
	Contract     *IWrappedToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWrappedTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWrappedTokenCallerSession struct {
	Contract *IWrappedTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IWrappedTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWrappedTokenTransactorSession struct {
	Contract     *IWrappedTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IWrappedTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWrappedTokenRaw struct {
	Contract *IWrappedToken // Generic contract binding to access the raw methods on
}

// IWrappedTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWrappedTokenCallerRaw struct {
	Contract *IWrappedTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IWrappedTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWrappedTokenTransactorRaw struct {
	Contract *IWrappedTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWrappedToken creates a new instance of IWrappedToken, bound to a specific deployed contract.
func NewIWrappedToken(address common.Address, backend bind.ContractBackend) (*IWrappedToken, error) {
	contract, err := bindIWrappedToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWrappedToken{IWrappedTokenCaller: IWrappedTokenCaller{contract: contract}, IWrappedTokenTransactor: IWrappedTokenTransactor{contract: contract}, IWrappedTokenFilterer: IWrappedTokenFilterer{contract: contract}}, nil
}

// NewIWrappedTokenCaller creates a new read-only instance of IWrappedToken, bound to a specific deployed contract.
func NewIWrappedTokenCaller(address common.Address, caller bind.ContractCaller) (*IWrappedTokenCaller, error) {
	contract, err := bindIWrappedToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWrappedTokenCaller{contract: contract}, nil
}

// NewIWrappedTokenTransactor creates a new write-only instance of IWrappedToken, bound to a specific deployed contract.
func NewIWrappedTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IWrappedTokenTransactor, error) {
	contract, err := bindIWrappedToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWrappedTokenTransactor{contract: contract}, nil
}

// NewIWrappedTokenFilterer creates a new log filterer instance of IWrappedToken, bound to a specific deployed contract.
func NewIWrappedTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IWrappedTokenFilterer, error) {
	contract, err := bindIWrappedToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWrappedTokenFilterer{contract: contract}, nil
}

// bindIWrappedToken binds a generic wrapper to an already deployed contract.
func bindIWrappedToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWrappedTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrappedToken *IWrappedTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrappedToken.Contract.IWrappedTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrappedToken *IWrappedTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrappedToken.Contract.IWrappedTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrappedToken *IWrappedTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrappedToken.Contract.IWrappedTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrappedToken *IWrappedTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrappedToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrappedToken *IWrappedTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrappedToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrappedToken *IWrappedTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrappedToken.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IWrappedToken *IWrappedTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IWrappedToken *IWrappedTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.Burn(&_IWrappedToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IWrappedToken *IWrappedTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.Burn(&_IWrappedToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IWrappedToken *IWrappedTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IWrappedToken *IWrappedTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.BurnFrom(&_IWrappedToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IWrappedToken *IWrappedTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.BurnFrom(&_IWrappedToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IWrappedToken *IWrappedTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IWrappedToken *IWrappedTokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.Mint(&_IWrappedToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IWrappedToken *IWrappedTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.Mint(&_IWrappedToken.TransactOpts, account, amount)
}

// LimiterDailyABI is the input ABI used to generate the binding from.
const LimiterDailyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"}],\"name\":\"getLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"}],\"name\":\"getUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"isLimited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LimiterDailyFuncSigs maps the 4-byte function signature to its string representation.
var LimiterDailyFuncSigs = map[string]string{
	"1ce28e72": "getLimit(address)",
	"59b9b15e": "getUsage(address)",
	"c0dd9e52": "increaseUsage(uint256)",
	"47618cb2": "isLimited(address,uint256)",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"36db43b5": "setLimit(address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// LimiterDailyBin is the compiled bytecode used for deploying new contracts.
var LimiterDailyBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916339081178255604051909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506105b4806100616000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063715018a61161005b578063715018a6146101145780638da5cb5b1461011c578063c0dd9e5214610137578063f2fde38b1461014a57600080fd5b80631ce28e721461008d57806336db43b5146100c957806347618cb2146100de57806359b9b15e14610101575b600080fd5b6100b661009b3660046104a3565b6001600160a01b031660009081526002602052604090205490565b6040519081526020015b60405180910390f35b6100dc6100d73660046104c4565b61015d565b005b6100f16100ec3660046104c4565b6101ac565b60405190151581526020016100c0565b6100b661010f3660046104a3565b610208565b6100dc610245565b6000546040516001600160a01b0390911681526020016100c0565b6100dc6101453660046104ed565b6102b9565b6100dc6101583660046104a3565b61039d565b6000546001600160a01b031633146101905760405162461bcd60e51b815260040161018790610505565b60405180910390fd5b6001600160a01b03909116600090815260026020526040902055565b6001600160a01b0382166000908152600260205260408120546101d157506000610202565b6001600160a01b038316600090815260026020526040902054826101f485610208565b6101fe919061053a565b1190505b92915050565b600080610218620151804261055e565b6001600160a01b039093166000908152600160209081526040808320958352949052929092205492915050565b6000546001600160a01b0316331461026f5760405162461bcd60e51b815260040161018790610505565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b336000818152600260205260409020546102d1575050565b60006102e0620151804261055e565b6001600160a01b038316600090815260016020908152604080832084845290915281208054929350859290919061031890849061053a565b90915550506001600160a01b038216600090815260026020908152604080832054600183528184208585529092529091205411156103985760405162461bcd60e51b815260206004820152601c60248201527f4c696d697465724461696c793a206c696d6974206578636565646564000000006044820152606401610187565b505050565b6000546001600160a01b031633146103c75760405162461bcd60e51b815260040161018790610505565b6001600160a01b03811661042c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610187565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b80356001600160a01b038116811461049e57600080fd5b919050565b6000602082840312156104b4578081fd5b6104bd82610487565b9392505050565b600080604083850312156104d6578081fd5b6104df83610487565b946020939093013593505050565b6000602082840312156104fe578081fd5b5035919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6000821982111561055957634e487b7160e01b81526011600452602481fd5b500190565b60008261057957634e487b7160e01b81526012600452602481fd5b50049056fea2646970667358221220987900100566218745bd55d989ae81e2388db9646e88427ca6a81ef0dc01859264736f6c63430008040033"

// DeployLimiterDaily deploys a new Ethereum contract, binding an instance of LimiterDaily to it.
func DeployLimiterDaily(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LimiterDaily, error) {
	parsed, err := abi.JSON(strings.NewReader(LimiterDailyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LimiterDailyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LimiterDaily{LimiterDailyCaller: LimiterDailyCaller{contract: contract}, LimiterDailyTransactor: LimiterDailyTransactor{contract: contract}, LimiterDailyFilterer: LimiterDailyFilterer{contract: contract}}, nil
}

// LimiterDaily is an auto generated Go binding around an Ethereum contract.
type LimiterDaily struct {
	LimiterDailyCaller     // Read-only binding to the contract
	LimiterDailyTransactor // Write-only binding to the contract
	LimiterDailyFilterer   // Log filterer for contract events
}

// LimiterDailyCaller is an auto generated read-only Go binding around an Ethereum contract.
type LimiterDailyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimiterDailyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LimiterDailyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimiterDailyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LimiterDailyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimiterDailySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LimiterDailySession struct {
	Contract     *LimiterDaily     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LimiterDailyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LimiterDailyCallerSession struct {
	Contract *LimiterDailyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LimiterDailyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LimiterDailyTransactorSession struct {
	Contract     *LimiterDailyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LimiterDailyRaw is an auto generated low-level Go binding around an Ethereum contract.
type LimiterDailyRaw struct {
	Contract *LimiterDaily // Generic contract binding to access the raw methods on
}

// LimiterDailyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LimiterDailyCallerRaw struct {
	Contract *LimiterDailyCaller // Generic read-only contract binding to access the raw methods on
}

// LimiterDailyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LimiterDailyTransactorRaw struct {
	Contract *LimiterDailyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLimiterDaily creates a new instance of LimiterDaily, bound to a specific deployed contract.
func NewLimiterDaily(address common.Address, backend bind.ContractBackend) (*LimiterDaily, error) {
	contract, err := bindLimiterDaily(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LimiterDaily{LimiterDailyCaller: LimiterDailyCaller{contract: contract}, LimiterDailyTransactor: LimiterDailyTransactor{contract: contract}, LimiterDailyFilterer: LimiterDailyFilterer{contract: contract}}, nil
}

// NewLimiterDailyCaller creates a new read-only instance of LimiterDaily, bound to a specific deployed contract.
func NewLimiterDailyCaller(address common.Address, caller bind.ContractCaller) (*LimiterDailyCaller, error) {
	contract, err := bindLimiterDaily(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LimiterDailyCaller{contract: contract}, nil
}

// NewLimiterDailyTransactor creates a new write-only instance of LimiterDaily, bound to a specific deployed contract.
func NewLimiterDailyTransactor(address common.Address, transactor bind.ContractTransactor) (*LimiterDailyTransactor, error) {
	contract, err := bindLimiterDaily(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LimiterDailyTransactor{contract: contract}, nil
}

// NewLimiterDailyFilterer creates a new log filterer instance of LimiterDaily, bound to a specific deployed contract.
func NewLimiterDailyFilterer(address common.Address, filterer bind.ContractFilterer) (*LimiterDailyFilterer, error) {
	contract, err := bindLimiterDaily(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LimiterDailyFilterer{contract: contract}, nil
}

// bindLimiterDaily binds a generic wrapper to an already deployed contract.
func bindLimiterDaily(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LimiterDailyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimiterDaily *LimiterDailyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimiterDaily.Contract.LimiterDailyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimiterDaily *LimiterDailyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimiterDaily.Contract.LimiterDailyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimiterDaily *LimiterDailyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimiterDaily.Contract.LimiterDailyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimiterDaily *LimiterDailyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimiterDaily.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimiterDaily *LimiterDailyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimiterDaily.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimiterDaily *LimiterDailyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimiterDaily.Contract.contract.Transact(opts, method, params...)
}

// GetLimit is a free data retrieval call binding the contract method 0x1ce28e72.
//
// Solidity: function getLimit(address bridge) view returns(uint256)
func (_LimiterDaily *LimiterDailyCaller) GetLimit(opts *bind.CallOpts, bridge common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LimiterDaily.contract.Call(opts, &out, "getLimit", bridge)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimit is a free data retrieval call binding the contract method 0x1ce28e72.
//
// Solidity: function getLimit(address bridge) view returns(uint256)
func (_LimiterDaily *LimiterDailySession) GetLimit(bridge common.Address) (*big.Int, error) {
	return _LimiterDaily.Contract.GetLimit(&_LimiterDaily.CallOpts, bridge)
}

// GetLimit is a free data retrieval call binding the contract method 0x1ce28e72.
//
// Solidity: function getLimit(address bridge) view returns(uint256)
func (_LimiterDaily *LimiterDailyCallerSession) GetLimit(bridge common.Address) (*big.Int, error) {
	return _LimiterDaily.Contract.GetLimit(&_LimiterDaily.CallOpts, bridge)
}

// GetUsage is a free data retrieval call binding the contract method 0x59b9b15e.
//
// Solidity: function getUsage(address bridge) view returns(uint256)
func (_LimiterDaily *LimiterDailyCaller) GetUsage(opts *bind.CallOpts, bridge common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LimiterDaily.contract.Call(opts, &out, "getUsage", bridge)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsage is a free data retrieval call binding the contract method 0x59b9b15e.
//
// Solidity: function getUsage(address bridge) view returns(uint256)
func (_LimiterDaily *LimiterDailySession) GetUsage(bridge common.Address) (*big.Int, error) {
	return _LimiterDaily.Contract.GetUsage(&_LimiterDaily.CallOpts, bridge)
}

// GetUsage is a free data retrieval call binding the contract method 0x59b9b15e.
//
// Solidity: function getUsage(address bridge) view returns(uint256)
func (_LimiterDaily *LimiterDailyCallerSession) GetUsage(bridge common.Address) (*big.Int, error) {
	return _LimiterDaily.Contract.GetUsage(&_LimiterDaily.CallOpts, bridge)
}

// IsLimited is a free data retrieval call binding the contract method 0x47618cb2.
//
// Solidity: function isLimited(address bridge, uint256 amount) view returns(bool)
func (_LimiterDaily *LimiterDailyCaller) IsLimited(opts *bind.CallOpts, bridge common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _LimiterDaily.contract.Call(opts, &out, "isLimited", bridge, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLimited is a free data retrieval call binding the contract method 0x47618cb2.
//
// Solidity: function isLimited(address bridge, uint256 amount) view returns(bool)
func (_LimiterDaily *LimiterDailySession) IsLimited(bridge common.Address, amount *big.Int) (bool, error) {
	return _LimiterDaily.Contract.IsLimited(&_LimiterDaily.CallOpts, bridge, amount)
}

// IsLimited is a free data retrieval call binding the contract method 0x47618cb2.
//
// Solidity: function isLimited(address bridge, uint256 amount) view returns(bool)
func (_LimiterDaily *LimiterDailyCallerSession) IsLimited(bridge common.Address, amount *big.Int) (bool, error) {
	return _LimiterDaily.Contract.IsLimited(&_LimiterDaily.CallOpts, bridge, amount)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimiterDaily *LimiterDailyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimiterDaily.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimiterDaily *LimiterDailySession) Owner() (common.Address, error) {
	return _LimiterDaily.Contract.Owner(&_LimiterDaily.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimiterDaily *LimiterDailyCallerSession) Owner() (common.Address, error) {
	return _LimiterDaily.Contract.Owner(&_LimiterDaily.CallOpts)
}

// IncreaseUsage is a paid mutator transaction binding the contract method 0xc0dd9e52.
//
// Solidity: function increaseUsage(uint256 amount) returns()
func (_LimiterDaily *LimiterDailyTransactor) IncreaseUsage(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LimiterDaily.contract.Transact(opts, "increaseUsage", amount)
}

// IncreaseUsage is a paid mutator transaction binding the contract method 0xc0dd9e52.
//
// Solidity: function increaseUsage(uint256 amount) returns()
func (_LimiterDaily *LimiterDailySession) IncreaseUsage(amount *big.Int) (*types.Transaction, error) {
	return _LimiterDaily.Contract.IncreaseUsage(&_LimiterDaily.TransactOpts, amount)
}

// IncreaseUsage is a paid mutator transaction binding the contract method 0xc0dd9e52.
//
// Solidity: function increaseUsage(uint256 amount) returns()
func (_LimiterDaily *LimiterDailyTransactorSession) IncreaseUsage(amount *big.Int) (*types.Transaction, error) {
	return _LimiterDaily.Contract.IncreaseUsage(&_LimiterDaily.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimiterDaily *LimiterDailyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimiterDaily.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimiterDaily *LimiterDailySession) RenounceOwnership() (*types.Transaction, error) {
	return _LimiterDaily.Contract.RenounceOwnership(&_LimiterDaily.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimiterDaily *LimiterDailyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LimiterDaily.Contract.RenounceOwnership(&_LimiterDaily.TransactOpts)
}

// SetLimit is a paid mutator transaction binding the contract method 0x36db43b5.
//
// Solidity: function setLimit(address bridge, uint256 limit) returns()
func (_LimiterDaily *LimiterDailyTransactor) SetLimit(opts *bind.TransactOpts, bridge common.Address, limit *big.Int) (*types.Transaction, error) {
	return _LimiterDaily.contract.Transact(opts, "setLimit", bridge, limit)
}

// SetLimit is a paid mutator transaction binding the contract method 0x36db43b5.
//
// Solidity: function setLimit(address bridge, uint256 limit) returns()
func (_LimiterDaily *LimiterDailySession) SetLimit(bridge common.Address, limit *big.Int) (*types.Transaction, error) {
	return _LimiterDaily.Contract.SetLimit(&_LimiterDaily.TransactOpts, bridge, limit)
}

// SetLimit is a paid mutator transaction binding the contract method 0x36db43b5.
//
// Solidity: function setLimit(address bridge, uint256 limit) returns()
func (_LimiterDaily *LimiterDailyTransactorSession) SetLimit(bridge common.Address, limit *big.Int) (*types.Transaction, error) {
	return _LimiterDaily.Contract.SetLimit(&_LimiterDaily.TransactOpts, bridge, limit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimiterDaily *LimiterDailyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LimiterDaily.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimiterDaily *LimiterDailySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LimiterDaily.Contract.TransferOwnership(&_LimiterDaily.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimiterDaily *LimiterDailyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LimiterDaily.Contract.TransferOwnership(&_LimiterDaily.TransactOpts, newOwner)
}

// LimiterDailyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LimiterDaily contract.
type LimiterDailyOwnershipTransferredIterator struct {
	Event *LimiterDailyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LimiterDailyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimiterDailyOwnershipTransferred)
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
		it.Event = new(LimiterDailyOwnershipTransferred)
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
func (it *LimiterDailyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimiterDailyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimiterDailyOwnershipTransferred represents a OwnershipTransferred event raised by the LimiterDaily contract.
type LimiterDailyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LimiterDaily *LimiterDailyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LimiterDailyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LimiterDaily.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LimiterDailyOwnershipTransferredIterator{contract: _LimiterDaily.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LimiterDaily *LimiterDailyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LimiterDailyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LimiterDaily.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimiterDailyOwnershipTransferred)
				if err := _LimiterDaily.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LimiterDaily *LimiterDailyFilterer) ParseOwnershipTransferred(log types.Log) (*LimiterDailyOwnershipTransferred, error) {
	event := new(LimiterDailyOwnershipTransferred)
	if err := _LimiterDaily.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterAccessControlABI is the input ABI used to generate the binding from.
const MinterAccessControlABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"mintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"removeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MinterAccessControlFuncSigs maps the 4-byte function signature to its string representation.
var MinterAccessControlFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"59441eae": "mintable(address)",
	"8da5cb5b": "owner()",
	"3092afd5": "removeMinter(address)",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// MinterAccessControl is an auto generated Go binding around an Ethereum contract.
type MinterAccessControl struct {
	MinterAccessControlCaller     // Read-only binding to the contract
	MinterAccessControlTransactor // Write-only binding to the contract
	MinterAccessControlFilterer   // Log filterer for contract events
}

// MinterAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinterAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinterAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinterAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinterAccessControlSession struct {
	Contract     *MinterAccessControl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MinterAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinterAccessControlCallerSession struct {
	Contract *MinterAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MinterAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinterAccessControlTransactorSession struct {
	Contract     *MinterAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MinterAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinterAccessControlRaw struct {
	Contract *MinterAccessControl // Generic contract binding to access the raw methods on
}

// MinterAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinterAccessControlCallerRaw struct {
	Contract *MinterAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// MinterAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinterAccessControlTransactorRaw struct {
	Contract *MinterAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinterAccessControl creates a new instance of MinterAccessControl, bound to a specific deployed contract.
func NewMinterAccessControl(address common.Address, backend bind.ContractBackend) (*MinterAccessControl, error) {
	contract, err := bindMinterAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinterAccessControl{MinterAccessControlCaller: MinterAccessControlCaller{contract: contract}, MinterAccessControlTransactor: MinterAccessControlTransactor{contract: contract}, MinterAccessControlFilterer: MinterAccessControlFilterer{contract: contract}}, nil
}

// NewMinterAccessControlCaller creates a new read-only instance of MinterAccessControl, bound to a specific deployed contract.
func NewMinterAccessControlCaller(address common.Address, caller bind.ContractCaller) (*MinterAccessControlCaller, error) {
	contract, err := bindMinterAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterAccessControlCaller{contract: contract}, nil
}

// NewMinterAccessControlTransactor creates a new write-only instance of MinterAccessControl, bound to a specific deployed contract.
func NewMinterAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterAccessControlTransactor, error) {
	contract, err := bindMinterAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterAccessControlTransactor{contract: contract}, nil
}

// NewMinterAccessControlFilterer creates a new log filterer instance of MinterAccessControl, bound to a specific deployed contract.
func NewMinterAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*MinterAccessControlFilterer, error) {
	contract, err := bindMinterAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterAccessControlFilterer{contract: contract}, nil
}

// bindMinterAccessControl binds a generic wrapper to an already deployed contract.
func bindMinterAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinterAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterAccessControl *MinterAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinterAccessControl.Contract.MinterAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterAccessControl *MinterAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterAccessControl.Contract.MinterAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterAccessControl *MinterAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterAccessControl.Contract.MinterAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterAccessControl *MinterAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinterAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterAccessControl *MinterAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterAccessControl *MinterAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterAccessControl.Contract.contract.Transact(opts, method, params...)
}

// Mintable is a free data retrieval call binding the contract method 0x59441eae.
//
// Solidity: function mintable(address minter) view returns(bool)
func (_MinterAccessControl *MinterAccessControlCaller) Mintable(opts *bind.CallOpts, minter common.Address) (bool, error) {
	var out []interface{}
	err := _MinterAccessControl.contract.Call(opts, &out, "mintable", minter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mintable is a free data retrieval call binding the contract method 0x59441eae.
//
// Solidity: function mintable(address minter) view returns(bool)
func (_MinterAccessControl *MinterAccessControlSession) Mintable(minter common.Address) (bool, error) {
	return _MinterAccessControl.Contract.Mintable(&_MinterAccessControl.CallOpts, minter)
}

// Mintable is a free data retrieval call binding the contract method 0x59441eae.
//
// Solidity: function mintable(address minter) view returns(bool)
func (_MinterAccessControl *MinterAccessControlCallerSession) Mintable(minter common.Address) (bool, error) {
	return _MinterAccessControl.Contract.Mintable(&_MinterAccessControl.CallOpts, minter)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinterAccessControl *MinterAccessControlCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinterAccessControl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinterAccessControl *MinterAccessControlSession) Owner() (common.Address, error) {
	return _MinterAccessControl.Contract.Owner(&_MinterAccessControl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinterAccessControl *MinterAccessControlCallerSession) Owner() (common.Address, error) {
	return _MinterAccessControl.Contract.Owner(&_MinterAccessControl.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_MinterAccessControl *MinterAccessControlTransactor) AddMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _MinterAccessControl.contract.Transact(opts, "addMinter", minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_MinterAccessControl *MinterAccessControlSession) AddMinter(minter common.Address) (*types.Transaction, error) {
	return _MinterAccessControl.Contract.AddMinter(&_MinterAccessControl.TransactOpts, minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_MinterAccessControl *MinterAccessControlTransactorSession) AddMinter(minter common.Address) (*types.Transaction, error) {
	return _MinterAccessControl.Contract.AddMinter(&_MinterAccessControl.TransactOpts, minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns()
func (_MinterAccessControl *MinterAccessControlTransactor) RemoveMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _MinterAccessControl.contract.Transact(opts, "removeMinter", minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns()
func (_MinterAccessControl *MinterAccessControlSession) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _MinterAccessControl.Contract.RemoveMinter(&_MinterAccessControl.TransactOpts, minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns()
func (_MinterAccessControl *MinterAccessControlTransactorSession) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _MinterAccessControl.Contract.RemoveMinter(&_MinterAccessControl.TransactOpts, minter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MinterAccessControl *MinterAccessControlTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterAccessControl.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MinterAccessControl *MinterAccessControlSession) RenounceOwnership() (*types.Transaction, error) {
	return _MinterAccessControl.Contract.RenounceOwnership(&_MinterAccessControl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MinterAccessControl *MinterAccessControlTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MinterAccessControl.Contract.RenounceOwnership(&_MinterAccessControl.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinterAccessControl *MinterAccessControlTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MinterAccessControl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinterAccessControl *MinterAccessControlSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MinterAccessControl.Contract.TransferOwnership(&_MinterAccessControl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinterAccessControl *MinterAccessControlTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MinterAccessControl.Contract.TransferOwnership(&_MinterAccessControl.TransactOpts, newOwner)
}

// MinterAccessControlMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the MinterAccessControl contract.
type MinterAccessControlMinterAddedIterator struct {
	Event *MinterAccessControlMinterAdded // Event containing the contract specifics and raw log

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
func (it *MinterAccessControlMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterAccessControlMinterAdded)
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
		it.Event = new(MinterAccessControlMinterAdded)
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
func (it *MinterAccessControlMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterAccessControlMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterAccessControlMinterAdded represents a MinterAdded event raised by the MinterAccessControl contract.
type MinterAccessControlMinterAdded struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed minter)
func (_MinterAccessControl *MinterAccessControlFilterer) FilterMinterAdded(opts *bind.FilterOpts, minter []common.Address) (*MinterAccessControlMinterAddedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _MinterAccessControl.contract.FilterLogs(opts, "MinterAdded", minterRule)
	if err != nil {
		return nil, err
	}
	return &MinterAccessControlMinterAddedIterator{contract: _MinterAccessControl.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed minter)
func (_MinterAccessControl *MinterAccessControlFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *MinterAccessControlMinterAdded, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _MinterAccessControl.contract.WatchLogs(opts, "MinterAdded", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterAccessControlMinterAdded)
				if err := _MinterAccessControl.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed minter)
func (_MinterAccessControl *MinterAccessControlFilterer) ParseMinterAdded(log types.Log) (*MinterAccessControlMinterAdded, error) {
	event := new(MinterAccessControlMinterAdded)
	if err := _MinterAccessControl.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterAccessControlMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the MinterAccessControl contract.
type MinterAccessControlMinterRemovedIterator struct {
	Event *MinterAccessControlMinterRemoved // Event containing the contract specifics and raw log

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
func (it *MinterAccessControlMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterAccessControlMinterRemoved)
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
		it.Event = new(MinterAccessControlMinterRemoved)
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
func (it *MinterAccessControlMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterAccessControlMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterAccessControlMinterRemoved represents a MinterRemoved event raised by the MinterAccessControl contract.
type MinterAccessControlMinterRemoved struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed minter)
func (_MinterAccessControl *MinterAccessControlFilterer) FilterMinterRemoved(opts *bind.FilterOpts, minter []common.Address) (*MinterAccessControlMinterRemovedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _MinterAccessControl.contract.FilterLogs(opts, "MinterRemoved", minterRule)
	if err != nil {
		return nil, err
	}
	return &MinterAccessControlMinterRemovedIterator{contract: _MinterAccessControl.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed minter)
func (_MinterAccessControl *MinterAccessControlFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *MinterAccessControlMinterRemoved, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _MinterAccessControl.contract.WatchLogs(opts, "MinterRemoved", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterAccessControlMinterRemoved)
				if err := _MinterAccessControl.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed minter)
func (_MinterAccessControl *MinterAccessControlFilterer) ParseMinterRemoved(log types.Log) (*MinterAccessControlMinterRemoved, error) {
	event := new(MinterAccessControlMinterRemoved)
	if err := _MinterAccessControl.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterAccessControlOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MinterAccessControl contract.
type MinterAccessControlOwnershipTransferredIterator struct {
	Event *MinterAccessControlOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MinterAccessControlOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterAccessControlOwnershipTransferred)
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
		it.Event = new(MinterAccessControlOwnershipTransferred)
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
func (it *MinterAccessControlOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterAccessControlOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterAccessControlOwnershipTransferred represents a OwnershipTransferred event raised by the MinterAccessControl contract.
type MinterAccessControlOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MinterAccessControl *MinterAccessControlFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MinterAccessControlOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MinterAccessControl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MinterAccessControlOwnershipTransferredIterator{contract: _MinterAccessControl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MinterAccessControl *MinterAccessControlFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MinterAccessControlOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MinterAccessControl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterAccessControlOwnershipTransferred)
				if err := _MinterAccessControl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MinterAccessControl *MinterAccessControlFilterer) ParseOwnershipTransferred(log types.Log) (*MinterAccessControlOwnershipTransferred, error) {
	event := new(MinterAccessControlOwnershipTransferred)
	if err := _MinterAccessControl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"5c975abb": "paused()",
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[]"

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200dcff97d37d034e2be647a64767ba16d072fd3b3cd727bde85a907c06787edac64736f6c63430008040033"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// WrappedTokenABI is the input ABI used to generate the binding from.
const WrappedTokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"mintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"removeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WrappedTokenFuncSigs maps the 4-byte function signature to its string representation.
var WrappedTokenFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"79cc6790": "burnFrom(address,uint256)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"40c10f19": "mint(address,uint256)",
	"59441eae": "mintable(address)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"3092afd5": "removeMinter(address)",
	"715018a6": "renounceOwnership()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// WrappedTokenBin is the compiled bytecode used for deploying new contracts.
var WrappedTokenBin = "0x60806040523480156200001157600080fd5b506040516200156838038062001568833981016040819052620000349162000240565b8251839083906200004d906003906020850190620000e7565b50805162000063906004906020840190620000e7565b505050600062000078620000e360201b60201c565b600580546001600160a01b0319166001600160a01b038316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506007805460ff191660ff9290921691909117905550620003149050565b3390565b828054620000f590620002c1565b90600052602060002090601f01602090048101928262000119576000855562000164565b82601f106200013457805160ff191683800117855562000164565b8280016001018555821562000164579182015b828111156200016457825182559160200191906001019062000147565b506200017292915062000176565b5090565b5b8082111562000172576000815560010162000177565b600082601f8301126200019e578081fd5b81516001600160401b0380821115620001bb57620001bb620002fe565b604051601f8301601f19908116603f01168101908282118183101715620001e657620001e6620002fe565b8160405283815260209250868385880101111562000202578485fd5b8491505b8382101562000225578582018301518183018401529082019062000206565b838211156200023657848385830101525b9695505050505050565b60008060006060848603121562000255578283fd5b83516001600160401b03808211156200026c578485fd5b6200027a878388016200018d565b9450602086015191508082111562000290578384fd5b506200029f868287016200018d565b925050604084015160ff81168114620002b6578182fd5b809150509250925092565b600181811c90821680620002d657607f821691505b60208210811415620002f857634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b61124480620003246000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806370a08231116100ad578063983b2d5611610071578063983b2d561461028d578063a457c2d7146102a0578063a9059cbb146102b3578063dd62ed3e146102c6578063f2fde38b146102ff57600080fd5b806370a0823114610226578063715018a61461024f57806379cc6790146102575780638da5cb5b1461026a57806395d89b411461028557600080fd5b8063313ce567116100f4578063313ce567146101ac57806339509351146101c157806340c10f19146101d457806342966c68146101e757806359441eae146101fa57600080fd5b806306fdde0314610131578063095ea7b31461014f57806318160ddd1461017257806323b872dd146101845780633092afd514610197575b600080fd5b610139610312565b604051610146919061106e565b60405180910390f35b61016261015d36600461102d565b6103a4565b6040519015158152602001610146565b6002545b604051908152602001610146565b610162610192366004610ff2565b6103ba565b6101aa6101a5366004610f9f565b610470565b005b60075460405160ff9091168152602001610146565b6101626101cf36600461102d565b61059c565b6101aa6101e236600461102d565b6105d3565b6101aa6101f5366004611056565b610610565b610162610208366004610f9f565b6001600160a01b031660009081526006602052604090205460ff1690565b610176610234366004610f9f565b6001600160a01b031660009081526020819052604090205490565b6101aa61064b565b6101aa61026536600461102d565b6106bf565b6005546040516001600160a01b039091168152602001610146565b6101396106f8565b6101aa61029b366004610f9f565b610707565b6101626102ae36600461102d565b610833565b6101626102c136600461102d565b6108ce565b6101766102d4366004610fc0565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6101aa61030d366004610f9f565b6108db565b606060038054610321906111bd565b80601f016020809104026020016040519081016040528092919081815260200182805461034d906111bd565b801561039a5780601f1061036f5761010080835404028352916020019161039a565b820191906000526020600020905b81548152906001019060200180831161037d57829003601f168201915b5050505050905090565b60006103b13384846109c6565b50600192915050565b60006103c7848484610aeb565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156104515760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b610465853361046086856111a6565b6109c6565b506001949350505050565b6005546001600160a01b0316331461049a5760405162461bcd60e51b8152600401610448906110c1565b6001600160a01b0381166104c05760405162461bcd60e51b8152600401610448906110f6565b6001600160a01b03811660009081526006602052604090205460ff166105475760405162461bcd60e51b815260206004820152603660248201527f4d696e746572416363657373436f6e74726f6c3a206d696e74657220616c726560448201527518591e481b9bdd081a5b881858d8d95cdcc81b1a5cdd60521b6064820152608401610448565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a26001600160a01b03166000908152600660205260409020805460ff19169055565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916103b191859061046090869061118e565b3360009081526006602052604090205460ff166106025760405162461bcd60e51b815260040161044890611141565b61060c8282610cc3565b5050565b3360009081526006602052604090205460ff1661063f5760405162461bcd60e51b815260040161044890611141565b61064881610da2565b50565b6005546001600160a01b031633146106755760405162461bcd60e51b8152600401610448906110c1565b6005546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600580546001600160a01b0319169055565b3360009081526006602052604090205460ff166106ee5760405162461bcd60e51b815260040161044890611141565b61060c8282610dac565b606060048054610321906111bd565b6005546001600160a01b031633146107315760405162461bcd60e51b8152600401610448906110c1565b6001600160a01b0381166107575760405162461bcd60e51b8152600401610448906110f6565b6001600160a01b03811660009081526006602052604090205460ff16156107db5760405162461bcd60e51b815260206004820152603260248201527f4d696e746572416363657373436f6e74726f6c3a206d696e74657220616c726560448201527118591e481a5b881858d8d95cdcc81b1a5cdd60721b6064820152608401610448565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a26001600160a01b03166000908152600660205260409020805460ff19166001179055565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156108b55760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610448565b6108c4338561046086856111a6565b5060019392505050565b60006103b1338484610aeb565b6005546001600160a01b031633146109055760405162461bcd60e51b8152600401610448906110c1565b6001600160a01b03811661096a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610448565b6005546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600580546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b038316610a285760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610448565b6001600160a01b038216610a895760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610448565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b038316610b4f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610448565b6001600160a01b038216610bb15760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610448565b6001600160a01b03831660009081526020819052604090205481811015610c295760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610448565b610c3382826111a6565b6001600160a01b038086166000908152602081905260408082209390935590851681529081208054849290610c6990849061118e565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610cb591815260200190565b60405180910390a350505050565b6001600160a01b038216610d195760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610448565b8060026000828254610d2b919061118e565b90915550506001600160a01b03821660009081526020819052604081208054839290610d5890849061118e565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6106483382610e34565b6000610db883336102d4565b905081811015610e165760405162461bcd60e51b8152602060048201526024808201527f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f77604482015263616e636560e01b6064820152608401610448565b610e25833361046085856111a6565b610e2f8383610e34565b505050565b6001600160a01b038216610e945760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610448565b6001600160a01b03821660009081526020819052604090205481811015610f085760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610448565b610f1282826111a6565b6001600160a01b03841660009081526020819052604081209190915560028054849290610f409084906111a6565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610ade565b80356001600160a01b0381168114610f9a57600080fd5b919050565b600060208284031215610fb0578081fd5b610fb982610f83565b9392505050565b60008060408385031215610fd2578081fd5b610fdb83610f83565b9150610fe960208401610f83565b90509250929050565b600080600060608486031215611006578081fd5b61100f84610f83565b925061101d60208501610f83565b9150604084013590509250925092565b6000806040838503121561103f578182fd5b61104883610f83565b946020939093013593505050565b600060208284031215611067578081fd5b5035919050565b6000602080835283518082850152825b8181101561109a5785810183015185820160400152820161107e565b818111156110ab5783604083870101525b50601f01601f1916929092016040019392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252602b908201527f4d696e746572416363657373436f6e74726f6c3a2063616e206e6f742061646460408201526a206164647265737328302960a81b606082015260800190565b6020808252602d908201527f4d696e746572416363657373436f6e74726f6c3a2063616c6c6572206973206e60408201526c37ba103a34329036b4b73a32b960991b606082015260800190565b600082198211156111a1576111a16111f8565b500190565b6000828210156111b8576111b86111f8565b500390565b600181811c908216806111d157607f821691505b602082108114156111f257634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220bc5a326f7f99e2bc1b7687cde1ac1fe620bae44084b24869fa7ca3a4dfd3e6ec64736f6c63430008040033"

// DeployWrappedToken deploys a new Ethereum contract, binding an instance of WrappedToken to it.
func DeployWrappedToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals_ uint8) (common.Address, *types.Transaction, *WrappedToken, error) {
	parsed, err := abi.JSON(strings.NewReader(WrappedTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WrappedTokenBin), backend, name, symbol, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrappedToken{WrappedTokenCaller: WrappedTokenCaller{contract: contract}, WrappedTokenTransactor: WrappedTokenTransactor{contract: contract}, WrappedTokenFilterer: WrappedTokenFilterer{contract: contract}}, nil
}

// WrappedToken is an auto generated Go binding around an Ethereum contract.
type WrappedToken struct {
	WrappedTokenCaller     // Read-only binding to the contract
	WrappedTokenTransactor // Write-only binding to the contract
	WrappedTokenFilterer   // Log filterer for contract events
}

// WrappedTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappedTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappedTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappedTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappedTokenSession struct {
	Contract     *WrappedToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrappedTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappedTokenCallerSession struct {
	Contract *WrappedTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WrappedTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappedTokenTransactorSession struct {
	Contract     *WrappedTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WrappedTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappedTokenRaw struct {
	Contract *WrappedToken // Generic contract binding to access the raw methods on
}

// WrappedTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappedTokenCallerRaw struct {
	Contract *WrappedTokenCaller // Generic read-only contract binding to access the raw methods on
}

// WrappedTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappedTokenTransactorRaw struct {
	Contract *WrappedTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappedToken creates a new instance of WrappedToken, bound to a specific deployed contract.
func NewWrappedToken(address common.Address, backend bind.ContractBackend) (*WrappedToken, error) {
	contract, err := bindWrappedToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedToken{WrappedTokenCaller: WrappedTokenCaller{contract: contract}, WrappedTokenTransactor: WrappedTokenTransactor{contract: contract}, WrappedTokenFilterer: WrappedTokenFilterer{contract: contract}}, nil
}

// NewWrappedTokenCaller creates a new read-only instance of WrappedToken, bound to a specific deployed contract.
func NewWrappedTokenCaller(address common.Address, caller bind.ContractCaller) (*WrappedTokenCaller, error) {
	contract, err := bindWrappedToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenCaller{contract: contract}, nil
}

// NewWrappedTokenTransactor creates a new write-only instance of WrappedToken, bound to a specific deployed contract.
func NewWrappedTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedTokenTransactor, error) {
	contract, err := bindWrappedToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenTransactor{contract: contract}, nil
}

// NewWrappedTokenFilterer creates a new log filterer instance of WrappedToken, bound to a specific deployed contract.
func NewWrappedTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedTokenFilterer, error) {
	contract, err := bindWrappedToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenFilterer{contract: contract}, nil
}

// bindWrappedToken binds a generic wrapper to an already deployed contract.
func bindWrappedToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrappedTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedToken *WrappedTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedToken.Contract.WrappedTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedToken *WrappedTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedToken.Contract.WrappedTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedToken *WrappedTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedToken.Contract.WrappedTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedToken *WrappedTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedToken *WrappedTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedToken *WrappedTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedToken *WrappedTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedToken *WrappedTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedToken.Contract.Allowance(&_WrappedToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedToken *WrappedTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedToken.Contract.Allowance(&_WrappedToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedToken *WrappedTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedToken *WrappedTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedToken.Contract.BalanceOf(&_WrappedToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedToken *WrappedTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedToken.Contract.BalanceOf(&_WrappedToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedToken *WrappedTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrappedToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedToken *WrappedTokenSession) Decimals() (uint8, error) {
	return _WrappedToken.Contract.Decimals(&_WrappedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedToken *WrappedTokenCallerSession) Decimals() (uint8, error) {
	return _WrappedToken.Contract.Decimals(&_WrappedToken.CallOpts)
}

// Mintable is a free data retrieval call binding the contract method 0x59441eae.
//
// Solidity: function mintable(address minter) view returns(bool)
func (_WrappedToken *WrappedTokenCaller) Mintable(opts *bind.CallOpts, minter common.Address) (bool, error) {
	var out []interface{}
	err := _WrappedToken.contract.Call(opts, &out, "mintable", minter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mintable is a free data retrieval call binding the contract method 0x59441eae.
//
// Solidity: function mintable(address minter) view returns(bool)
func (_WrappedToken *WrappedTokenSession) Mintable(minter common.Address) (bool, error) {
	return _WrappedToken.Contract.Mintable(&_WrappedToken.CallOpts, minter)
}

// Mintable is a free data retrieval call binding the contract method 0x59441eae.
//
// Solidity: function mintable(address minter) view returns(bool)
func (_WrappedToken *WrappedTokenCallerSession) Mintable(minter common.Address) (bool, error) {
	return _WrappedToken.Contract.Mintable(&_WrappedToken.CallOpts, minter)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedToken *WrappedTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedToken *WrappedTokenSession) Name() (string, error) {
	return _WrappedToken.Contract.Name(&_WrappedToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedToken *WrappedTokenCallerSession) Name() (string, error) {
	return _WrappedToken.Contract.Name(&_WrappedToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrappedToken *WrappedTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrappedToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrappedToken *WrappedTokenSession) Owner() (common.Address, error) {
	return _WrappedToken.Contract.Owner(&_WrappedToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrappedToken *WrappedTokenCallerSession) Owner() (common.Address, error) {
	return _WrappedToken.Contract.Owner(&_WrappedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedToken *WrappedTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedToken *WrappedTokenSession) Symbol() (string, error) {
	return _WrappedToken.Contract.Symbol(&_WrappedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedToken *WrappedTokenCallerSession) Symbol() (string, error) {
	return _WrappedToken.Contract.Symbol(&_WrappedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedToken *WrappedTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrappedToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedToken *WrappedTokenSession) TotalSupply() (*big.Int, error) {
	return _WrappedToken.Contract.TotalSupply(&_WrappedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedToken *WrappedTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _WrappedToken.Contract.TotalSupply(&_WrappedToken.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_WrappedToken *WrappedTokenTransactor) AddMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "addMinter", minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_WrappedToken *WrappedTokenSession) AddMinter(minter common.Address) (*types.Transaction, error) {
	return _WrappedToken.Contract.AddMinter(&_WrappedToken.TransactOpts, minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_WrappedToken *WrappedTokenTransactorSession) AddMinter(minter common.Address) (*types.Transaction, error) {
	return _WrappedToken.Contract.AddMinter(&_WrappedToken.TransactOpts, minter)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedToken *WrappedTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedToken *WrappedTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.Approve(&_WrappedToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedToken *WrappedTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.Approve(&_WrappedToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_WrappedToken *WrappedTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_WrappedToken *WrappedTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.Burn(&_WrappedToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_WrappedToken *WrappedTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.Burn(&_WrappedToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_WrappedToken *WrappedTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_WrappedToken *WrappedTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.BurnFrom(&_WrappedToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_WrappedToken *WrappedTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.BurnFrom(&_WrappedToken.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedToken *WrappedTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedToken *WrappedTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.DecreaseAllowance(&_WrappedToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedToken *WrappedTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.DecreaseAllowance(&_WrappedToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedToken *WrappedTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedToken *WrappedTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.IncreaseAllowance(&_WrappedToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedToken *WrappedTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.IncreaseAllowance(&_WrappedToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_WrappedToken *WrappedTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_WrappedToken *WrappedTokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.Mint(&_WrappedToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_WrappedToken *WrappedTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.Mint(&_WrappedToken.TransactOpts, account, amount)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns()
func (_WrappedToken *WrappedTokenTransactor) RemoveMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "removeMinter", minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns()
func (_WrappedToken *WrappedTokenSession) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _WrappedToken.Contract.RemoveMinter(&_WrappedToken.TransactOpts, minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns()
func (_WrappedToken *WrappedTokenTransactorSession) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _WrappedToken.Contract.RemoveMinter(&_WrappedToken.TransactOpts, minter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WrappedToken *WrappedTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WrappedToken *WrappedTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _WrappedToken.Contract.RenounceOwnership(&_WrappedToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WrappedToken *WrappedTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WrappedToken.Contract.RenounceOwnership(&_WrappedToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WrappedToken *WrappedTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WrappedToken *WrappedTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.Transfer(&_WrappedToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WrappedToken *WrappedTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.Transfer(&_WrappedToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WrappedToken *WrappedTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WrappedToken *WrappedTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.TransferFrom(&_WrappedToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WrappedToken *WrappedTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedToken.Contract.TransferFrom(&_WrappedToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WrappedToken *WrappedTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WrappedToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WrappedToken *WrappedTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WrappedToken.Contract.TransferOwnership(&_WrappedToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WrappedToken *WrappedTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WrappedToken.Contract.TransferOwnership(&_WrappedToken.TransactOpts, newOwner)
}

// WrappedTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WrappedToken contract.
type WrappedTokenApprovalIterator struct {
	Event *WrappedTokenApproval // Event containing the contract specifics and raw log

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
func (it *WrappedTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenApproval)
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
		it.Event = new(WrappedTokenApproval)
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
func (it *WrappedTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedTokenApproval represents a Approval event raised by the WrappedToken contract.
type WrappedTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedToken *WrappedTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrappedTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenApprovalIterator{contract: _WrappedToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedToken *WrappedTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrappedTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedTokenApproval)
				if err := _WrappedToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WrappedToken *WrappedTokenFilterer) ParseApproval(log types.Log) (*WrappedTokenApproval, error) {
	event := new(WrappedTokenApproval)
	if err := _WrappedToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedTokenMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the WrappedToken contract.
type WrappedTokenMinterAddedIterator struct {
	Event *WrappedTokenMinterAdded // Event containing the contract specifics and raw log

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
func (it *WrappedTokenMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenMinterAdded)
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
		it.Event = new(WrappedTokenMinterAdded)
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
func (it *WrappedTokenMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedTokenMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedTokenMinterAdded represents a MinterAdded event raised by the WrappedToken contract.
type WrappedTokenMinterAdded struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed minter)
func (_WrappedToken *WrappedTokenFilterer) FilterMinterAdded(opts *bind.FilterOpts, minter []common.Address) (*WrappedTokenMinterAddedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _WrappedToken.contract.FilterLogs(opts, "MinterAdded", minterRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenMinterAddedIterator{contract: _WrappedToken.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed minter)
func (_WrappedToken *WrappedTokenFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *WrappedTokenMinterAdded, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _WrappedToken.contract.WatchLogs(opts, "MinterAdded", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedTokenMinterAdded)
				if err := _WrappedToken.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed minter)
func (_WrappedToken *WrappedTokenFilterer) ParseMinterAdded(log types.Log) (*WrappedTokenMinterAdded, error) {
	event := new(WrappedTokenMinterAdded)
	if err := _WrappedToken.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedTokenMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the WrappedToken contract.
type WrappedTokenMinterRemovedIterator struct {
	Event *WrappedTokenMinterRemoved // Event containing the contract specifics and raw log

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
func (it *WrappedTokenMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenMinterRemoved)
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
		it.Event = new(WrappedTokenMinterRemoved)
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
func (it *WrappedTokenMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedTokenMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedTokenMinterRemoved represents a MinterRemoved event raised by the WrappedToken contract.
type WrappedTokenMinterRemoved struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed minter)
func (_WrappedToken *WrappedTokenFilterer) FilterMinterRemoved(opts *bind.FilterOpts, minter []common.Address) (*WrappedTokenMinterRemovedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _WrappedToken.contract.FilterLogs(opts, "MinterRemoved", minterRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenMinterRemovedIterator{contract: _WrappedToken.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed minter)
func (_WrappedToken *WrappedTokenFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *WrappedTokenMinterRemoved, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _WrappedToken.contract.WatchLogs(opts, "MinterRemoved", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedTokenMinterRemoved)
				if err := _WrappedToken.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed minter)
func (_WrappedToken *WrappedTokenFilterer) ParseMinterRemoved(log types.Log) (*WrappedTokenMinterRemoved, error) {
	event := new(WrappedTokenMinterRemoved)
	if err := _WrappedToken.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WrappedToken contract.
type WrappedTokenOwnershipTransferredIterator struct {
	Event *WrappedTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WrappedTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenOwnershipTransferred)
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
		it.Event = new(WrappedTokenOwnershipTransferred)
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
func (it *WrappedTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedTokenOwnershipTransferred represents a OwnershipTransferred event raised by the WrappedToken contract.
type WrappedTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WrappedToken *WrappedTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WrappedTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WrappedToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenOwnershipTransferredIterator{contract: _WrappedToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WrappedToken *WrappedTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WrappedTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WrappedToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedTokenOwnershipTransferred)
				if err := _WrappedToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WrappedToken *WrappedTokenFilterer) ParseOwnershipTransferred(log types.Log) (*WrappedTokenOwnershipTransferred, error) {
	event := new(WrappedTokenOwnershipTransferred)
	if err := _WrappedToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WrappedToken contract.
type WrappedTokenTransferIterator struct {
	Event *WrappedTokenTransfer // Event containing the contract specifics and raw log

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
func (it *WrappedTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenTransfer)
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
		it.Event = new(WrappedTokenTransfer)
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
func (it *WrappedTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedTokenTransfer represents a Transfer event raised by the WrappedToken contract.
type WrappedTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedToken *WrappedTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenTransferIterator{contract: _WrappedToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedToken *WrappedTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrappedTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedTokenTransfer)
				if err := _WrappedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WrappedToken *WrappedTokenFilterer) ParseTransfer(log types.Log) (*WrappedTokenTransfer, error) {
	event := new(WrappedTokenTransfer)
	if err := _WrappedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

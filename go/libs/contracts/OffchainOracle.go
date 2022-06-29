// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// OffchainOracleMetaData contains all meta data concerning the OffchainOracle contract.
var OffchainOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractMultiWrapper\",\"name\":\"_multiWrapper\",\"type\":\"address\"},{\"internalType\":\"contractIOracle[]\",\"name\":\"existingOracles\",\"type\":\"address[]\"},{\"internalType\":\"enumOffchainOracle.OracleType[]\",\"name\":\"oracleTypes\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"existingConnectors\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"wBase\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"ConnectorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"ConnectorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractMultiWrapper\",\"name\":\"multiWrapper\",\"type\":\"address\"}],\"name\":\"MultiWrapperUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumOffchainOracle.OracleType\",\"name\":\"oracleType\",\"type\":\"uint8\"}],\"name\":\"OracleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumOffchainOracle.OracleType\",\"name\":\"oracleType\",\"type\":\"uint8\"}],\"name\":\"OracleRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"addConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"enumOffchainOracle.OracleType\",\"name\":\"oracleKind\",\"type\":\"uint8\"}],\"name\":\"addOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectors\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"allConnectors\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useWrappers\",\"type\":\"bool\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useSrcWrappers\",\"type\":\"bool\"}],\"name\":\"getRateToEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiWrapper\",\"outputs\":[{\"internalType\":\"contractMultiWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"contractIOracle[]\",\"name\":\"allOracles\",\"type\":\"address[]\"},{\"internalType\":\"enumOffchainOracle.OracleType[]\",\"name\":\"oracleTypes\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"removeConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"enumOffchainOracle.OracleType\",\"name\":\"oracleKind\",\"type\":\"uint8\"}],\"name\":\"removeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractMultiWrapper\",\"name\":\"_multiWrapper\",\"type\":\"address\"}],\"name\":\"setMultiWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OffchainOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OffchainOracleMetaData.ABI instead.
var OffchainOracleABI = OffchainOracleMetaData.ABI

// OffchainOracle is an auto generated Go binding around an Ethereum contract.
type OffchainOracle struct {
	OffchainOracleCaller     // Read-only binding to the contract
	OffchainOracleTransactor // Write-only binding to the contract
	OffchainOracleFilterer   // Log filterer for contract events
}

// OffchainOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OffchainOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OffchainOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OffchainOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OffchainOracleSession struct {
	Contract     *OffchainOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OffchainOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OffchainOracleCallerSession struct {
	Contract *OffchainOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OffchainOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OffchainOracleTransactorSession struct {
	Contract     *OffchainOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OffchainOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OffchainOracleRaw struct {
	Contract *OffchainOracle // Generic contract binding to access the raw methods on
}

// OffchainOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OffchainOracleCallerRaw struct {
	Contract *OffchainOracleCaller // Generic read-only contract binding to access the raw methods on
}

// OffchainOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OffchainOracleTransactorRaw struct {
	Contract *OffchainOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOffchainOracle creates a new instance of OffchainOracle, bound to a specific deployed contract.
func NewOffchainOracle(address common.Address, backend bind.ContractBackend) (*OffchainOracle, error) {
	contract, err := bindOffchainOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OffchainOracle{OffchainOracleCaller: OffchainOracleCaller{contract: contract}, OffchainOracleTransactor: OffchainOracleTransactor{contract: contract}, OffchainOracleFilterer: OffchainOracleFilterer{contract: contract}}, nil
}

// NewOffchainOracleCaller creates a new read-only instance of OffchainOracle, bound to a specific deployed contract.
func NewOffchainOracleCaller(address common.Address, caller bind.ContractCaller) (*OffchainOracleCaller, error) {
	contract, err := bindOffchainOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffchainOracleCaller{contract: contract}, nil
}

// NewOffchainOracleTransactor creates a new write-only instance of OffchainOracle, bound to a specific deployed contract.
func NewOffchainOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OffchainOracleTransactor, error) {
	contract, err := bindOffchainOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffchainOracleTransactor{contract: contract}, nil
}

// NewOffchainOracleFilterer creates a new log filterer instance of OffchainOracle, bound to a specific deployed contract.
func NewOffchainOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OffchainOracleFilterer, error) {
	contract, err := bindOffchainOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffchainOracleFilterer{contract: contract}, nil
}

// bindOffchainOracle binds a generic wrapper to an already deployed contract.
func bindOffchainOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OffchainOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OffchainOracle *OffchainOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffchainOracle.Contract.OffchainOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OffchainOracle *OffchainOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainOracle.Contract.OffchainOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OffchainOracle *OffchainOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffchainOracle.Contract.OffchainOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OffchainOracle *OffchainOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffchainOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OffchainOracle *OffchainOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OffchainOracle *OffchainOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffchainOracle.Contract.contract.Transact(opts, method, params...)
}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OffchainOracle *OffchainOracleCaller) Connectors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "connectors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OffchainOracle *OffchainOracleSession) Connectors() ([]common.Address, error) {
	return _OffchainOracle.Contract.Connectors(&_OffchainOracle.CallOpts)
}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OffchainOracle *OffchainOracleCallerSession) Connectors() ([]common.Address, error) {
	return _OffchainOracle.Contract.Connectors(&_OffchainOracle.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCaller) GetRate(opts *bind.CallOpts, srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "getRate", srcToken, dstToken, useWrappers)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleSession) GetRate(srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRate(&_OffchainOracle.CallOpts, srcToken, dstToken, useWrappers)
}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCallerSession) GetRate(srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRate(&_OffchainOracle.CallOpts, srcToken, dstToken, useWrappers)
}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCaller) GetRateToEth(opts *bind.CallOpts, srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "getRateToEth", srcToken, useSrcWrappers)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleSession) GetRateToEth(srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateToEth(&_OffchainOracle.CallOpts, srcToken, useSrcWrappers)
}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OffchainOracle *OffchainOracleCallerSession) GetRateToEth(srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	return _OffchainOracle.Contract.GetRateToEth(&_OffchainOracle.CallOpts, srcToken, useSrcWrappers)
}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OffchainOracle *OffchainOracleCaller) MultiWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "multiWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OffchainOracle *OffchainOracleSession) MultiWrapper() (common.Address, error) {
	return _OffchainOracle.Contract.MultiWrapper(&_OffchainOracle.CallOpts)
}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OffchainOracle *OffchainOracleCallerSession) MultiWrapper() (common.Address, error) {
	return _OffchainOracle.Contract.MultiWrapper(&_OffchainOracle.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OffchainOracle *OffchainOracleCaller) Oracles(opts *bind.CallOpts) (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "oracles")

	outstruct := new(struct {
		AllOracles  []common.Address
		OracleTypes []uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllOracles = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OracleTypes = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OffchainOracle *OffchainOracleSession) Oracles() (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	return _OffchainOracle.Contract.Oracles(&_OffchainOracle.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OffchainOracle *OffchainOracleCallerSession) Oracles() (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	return _OffchainOracle.Contract.Oracles(&_OffchainOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainOracle *OffchainOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainOracle *OffchainOracleSession) Owner() (common.Address, error) {
	return _OffchainOracle.Contract.Owner(&_OffchainOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainOracle *OffchainOracleCallerSession) Owner() (common.Address, error) {
	return _OffchainOracle.Contract.Owner(&_OffchainOracle.CallOpts)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleTransactor) AddConnector(opts *bind.TransactOpts, connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "addConnector", connector)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleSession) AddConnector(connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.AddConnector(&_OffchainOracle.TransactOpts, connector)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) AddConnector(connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.AddConnector(&_OffchainOracle.TransactOpts, connector)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleTransactor) AddOracle(opts *bind.TransactOpts, oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "addOracle", oracle, oracleKind)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleSession) AddOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.Contract.AddOracle(&_OffchainOracle.TransactOpts, oracle, oracleKind)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) AddOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.Contract.AddOracle(&_OffchainOracle.TransactOpts, oracle, oracleKind)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleTransactor) RemoveConnector(opts *bind.TransactOpts, connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "removeConnector", connector)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleSession) RemoveConnector(connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.RemoveConnector(&_OffchainOracle.TransactOpts, connector)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) RemoveConnector(connector common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.RemoveConnector(&_OffchainOracle.TransactOpts, connector)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleTransactor) RemoveOracle(opts *bind.TransactOpts, oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "removeOracle", oracle, oracleKind)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleSession) RemoveOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.Contract.RemoveOracle(&_OffchainOracle.TransactOpts, oracle, oracleKind)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) RemoveOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OffchainOracle.Contract.RemoveOracle(&_OffchainOracle.TransactOpts, oracle, oracleKind)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OffchainOracle *OffchainOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OffchainOracle *OffchainOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _OffchainOracle.Contract.RenounceOwnership(&_OffchainOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OffchainOracle *OffchainOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OffchainOracle.Contract.RenounceOwnership(&_OffchainOracle.TransactOpts)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OffchainOracle *OffchainOracleTransactor) SetMultiWrapper(opts *bind.TransactOpts, _multiWrapper common.Address) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "setMultiWrapper", _multiWrapper)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OffchainOracle *OffchainOracleSession) SetMultiWrapper(_multiWrapper common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.SetMultiWrapper(&_OffchainOracle.TransactOpts, _multiWrapper)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) SetMultiWrapper(_multiWrapper common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.SetMultiWrapper(&_OffchainOracle.TransactOpts, _multiWrapper)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OffchainOracle *OffchainOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OffchainOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OffchainOracle *OffchainOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.TransferOwnership(&_OffchainOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OffchainOracle *OffchainOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OffchainOracle.Contract.TransferOwnership(&_OffchainOracle.TransactOpts, newOwner)
}

// OffchainOracleConnectorAddedIterator is returned from FilterConnectorAdded and is used to iterate over the raw logs and unpacked data for ConnectorAdded events raised by the OffchainOracle contract.
type OffchainOracleConnectorAddedIterator struct {
	Event *OffchainOracleConnectorAdded // Event containing the contract specifics and raw log

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
func (it *OffchainOracleConnectorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainOracleConnectorAdded)
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
		it.Event = new(OffchainOracleConnectorAdded)
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
func (it *OffchainOracleConnectorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainOracleConnectorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainOracleConnectorAdded represents a ConnectorAdded event raised by the OffchainOracle contract.
type OffchainOracleConnectorAdded struct {
	Connector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConnectorAdded is a free log retrieval operation binding the contract event 0xff88af5d962d47fd25d87755e8267a029fad5a91740c67d0dade2bdbe5268a1d.
//
// Solidity: event ConnectorAdded(address connector)
func (_OffchainOracle *OffchainOracleFilterer) FilterConnectorAdded(opts *bind.FilterOpts) (*OffchainOracleConnectorAddedIterator, error) {

	logs, sub, err := _OffchainOracle.contract.FilterLogs(opts, "ConnectorAdded")
	if err != nil {
		return nil, err
	}
	return &OffchainOracleConnectorAddedIterator{contract: _OffchainOracle.contract, event: "ConnectorAdded", logs: logs, sub: sub}, nil
}

// WatchConnectorAdded is a free log subscription operation binding the contract event 0xff88af5d962d47fd25d87755e8267a029fad5a91740c67d0dade2bdbe5268a1d.
//
// Solidity: event ConnectorAdded(address connector)
func (_OffchainOracle *OffchainOracleFilterer) WatchConnectorAdded(opts *bind.WatchOpts, sink chan<- *OffchainOracleConnectorAdded) (event.Subscription, error) {

	logs, sub, err := _OffchainOracle.contract.WatchLogs(opts, "ConnectorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainOracleConnectorAdded)
				if err := _OffchainOracle.contract.UnpackLog(event, "ConnectorAdded", log); err != nil {
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

// ParseConnectorAdded is a log parse operation binding the contract event 0xff88af5d962d47fd25d87755e8267a029fad5a91740c67d0dade2bdbe5268a1d.
//
// Solidity: event ConnectorAdded(address connector)
func (_OffchainOracle *OffchainOracleFilterer) ParseConnectorAdded(log types.Log) (*OffchainOracleConnectorAdded, error) {
	event := new(OffchainOracleConnectorAdded)
	if err := _OffchainOracle.contract.UnpackLog(event, "ConnectorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainOracleConnectorRemovedIterator is returned from FilterConnectorRemoved and is used to iterate over the raw logs and unpacked data for ConnectorRemoved events raised by the OffchainOracle contract.
type OffchainOracleConnectorRemovedIterator struct {
	Event *OffchainOracleConnectorRemoved // Event containing the contract specifics and raw log

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
func (it *OffchainOracleConnectorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainOracleConnectorRemoved)
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
		it.Event = new(OffchainOracleConnectorRemoved)
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
func (it *OffchainOracleConnectorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainOracleConnectorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainOracleConnectorRemoved represents a ConnectorRemoved event raised by the OffchainOracle contract.
type OffchainOracleConnectorRemoved struct {
	Connector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConnectorRemoved is a free log retrieval operation binding the contract event 0x6825b26a0827e9c2ceca01d6289ce4a40e629dc074ec48ea4727d1afbff359f5.
//
// Solidity: event ConnectorRemoved(address connector)
func (_OffchainOracle *OffchainOracleFilterer) FilterConnectorRemoved(opts *bind.FilterOpts) (*OffchainOracleConnectorRemovedIterator, error) {

	logs, sub, err := _OffchainOracle.contract.FilterLogs(opts, "ConnectorRemoved")
	if err != nil {
		return nil, err
	}
	return &OffchainOracleConnectorRemovedIterator{contract: _OffchainOracle.contract, event: "ConnectorRemoved", logs: logs, sub: sub}, nil
}

// WatchConnectorRemoved is a free log subscription operation binding the contract event 0x6825b26a0827e9c2ceca01d6289ce4a40e629dc074ec48ea4727d1afbff359f5.
//
// Solidity: event ConnectorRemoved(address connector)
func (_OffchainOracle *OffchainOracleFilterer) WatchConnectorRemoved(opts *bind.WatchOpts, sink chan<- *OffchainOracleConnectorRemoved) (event.Subscription, error) {

	logs, sub, err := _OffchainOracle.contract.WatchLogs(opts, "ConnectorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainOracleConnectorRemoved)
				if err := _OffchainOracle.contract.UnpackLog(event, "ConnectorRemoved", log); err != nil {
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

// ParseConnectorRemoved is a log parse operation binding the contract event 0x6825b26a0827e9c2ceca01d6289ce4a40e629dc074ec48ea4727d1afbff359f5.
//
// Solidity: event ConnectorRemoved(address connector)
func (_OffchainOracle *OffchainOracleFilterer) ParseConnectorRemoved(log types.Log) (*OffchainOracleConnectorRemoved, error) {
	event := new(OffchainOracleConnectorRemoved)
	if err := _OffchainOracle.contract.UnpackLog(event, "ConnectorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainOracleMultiWrapperUpdatedIterator is returned from FilterMultiWrapperUpdated and is used to iterate over the raw logs and unpacked data for MultiWrapperUpdated events raised by the OffchainOracle contract.
type OffchainOracleMultiWrapperUpdatedIterator struct {
	Event *OffchainOracleMultiWrapperUpdated // Event containing the contract specifics and raw log

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
func (it *OffchainOracleMultiWrapperUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainOracleMultiWrapperUpdated)
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
		it.Event = new(OffchainOracleMultiWrapperUpdated)
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
func (it *OffchainOracleMultiWrapperUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainOracleMultiWrapperUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainOracleMultiWrapperUpdated represents a MultiWrapperUpdated event raised by the OffchainOracle contract.
type OffchainOracleMultiWrapperUpdated struct {
	MultiWrapper common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMultiWrapperUpdated is a free log retrieval operation binding the contract event 0x1030152fe2062b574a830e6b9f13c65995990df31e4dc708d142533bb3ad0f52.
//
// Solidity: event MultiWrapperUpdated(address multiWrapper)
func (_OffchainOracle *OffchainOracleFilterer) FilterMultiWrapperUpdated(opts *bind.FilterOpts) (*OffchainOracleMultiWrapperUpdatedIterator, error) {

	logs, sub, err := _OffchainOracle.contract.FilterLogs(opts, "MultiWrapperUpdated")
	if err != nil {
		return nil, err
	}
	return &OffchainOracleMultiWrapperUpdatedIterator{contract: _OffchainOracle.contract, event: "MultiWrapperUpdated", logs: logs, sub: sub}, nil
}

// WatchMultiWrapperUpdated is a free log subscription operation binding the contract event 0x1030152fe2062b574a830e6b9f13c65995990df31e4dc708d142533bb3ad0f52.
//
// Solidity: event MultiWrapperUpdated(address multiWrapper)
func (_OffchainOracle *OffchainOracleFilterer) WatchMultiWrapperUpdated(opts *bind.WatchOpts, sink chan<- *OffchainOracleMultiWrapperUpdated) (event.Subscription, error) {

	logs, sub, err := _OffchainOracle.contract.WatchLogs(opts, "MultiWrapperUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainOracleMultiWrapperUpdated)
				if err := _OffchainOracle.contract.UnpackLog(event, "MultiWrapperUpdated", log); err != nil {
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

// ParseMultiWrapperUpdated is a log parse operation binding the contract event 0x1030152fe2062b574a830e6b9f13c65995990df31e4dc708d142533bb3ad0f52.
//
// Solidity: event MultiWrapperUpdated(address multiWrapper)
func (_OffchainOracle *OffchainOracleFilterer) ParseMultiWrapperUpdated(log types.Log) (*OffchainOracleMultiWrapperUpdated, error) {
	event := new(OffchainOracleMultiWrapperUpdated)
	if err := _OffchainOracle.contract.UnpackLog(event, "MultiWrapperUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainOracleOracleAddedIterator is returned from FilterOracleAdded and is used to iterate over the raw logs and unpacked data for OracleAdded events raised by the OffchainOracle contract.
type OffchainOracleOracleAddedIterator struct {
	Event *OffchainOracleOracleAdded // Event containing the contract specifics and raw log

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
func (it *OffchainOracleOracleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainOracleOracleAdded)
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
		it.Event = new(OffchainOracleOracleAdded)
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
func (it *OffchainOracleOracleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainOracleOracleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainOracleOracleAdded represents a OracleAdded event raised by the OffchainOracle contract.
type OffchainOracleOracleAdded struct {
	Oracle     common.Address
	OracleType uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOracleAdded is a free log retrieval operation binding the contract event 0x5874b2072ff37562df54063dd700c59d45f311bdf6f9cabb5a15f0ffb2e9f622.
//
// Solidity: event OracleAdded(address oracle, uint8 oracleType)
func (_OffchainOracle *OffchainOracleFilterer) FilterOracleAdded(opts *bind.FilterOpts) (*OffchainOracleOracleAddedIterator, error) {

	logs, sub, err := _OffchainOracle.contract.FilterLogs(opts, "OracleAdded")
	if err != nil {
		return nil, err
	}
	return &OffchainOracleOracleAddedIterator{contract: _OffchainOracle.contract, event: "OracleAdded", logs: logs, sub: sub}, nil
}

// WatchOracleAdded is a free log subscription operation binding the contract event 0x5874b2072ff37562df54063dd700c59d45f311bdf6f9cabb5a15f0ffb2e9f622.
//
// Solidity: event OracleAdded(address oracle, uint8 oracleType)
func (_OffchainOracle *OffchainOracleFilterer) WatchOracleAdded(opts *bind.WatchOpts, sink chan<- *OffchainOracleOracleAdded) (event.Subscription, error) {

	logs, sub, err := _OffchainOracle.contract.WatchLogs(opts, "OracleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainOracleOracleAdded)
				if err := _OffchainOracle.contract.UnpackLog(event, "OracleAdded", log); err != nil {
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

// ParseOracleAdded is a log parse operation binding the contract event 0x5874b2072ff37562df54063dd700c59d45f311bdf6f9cabb5a15f0ffb2e9f622.
//
// Solidity: event OracleAdded(address oracle, uint8 oracleType)
func (_OffchainOracle *OffchainOracleFilterer) ParseOracleAdded(log types.Log) (*OffchainOracleOracleAdded, error) {
	event := new(OffchainOracleOracleAdded)
	if err := _OffchainOracle.contract.UnpackLog(event, "OracleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainOracleOracleRemovedIterator is returned from FilterOracleRemoved and is used to iterate over the raw logs and unpacked data for OracleRemoved events raised by the OffchainOracle contract.
type OffchainOracleOracleRemovedIterator struct {
	Event *OffchainOracleOracleRemoved // Event containing the contract specifics and raw log

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
func (it *OffchainOracleOracleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainOracleOracleRemoved)
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
		it.Event = new(OffchainOracleOracleRemoved)
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
func (it *OffchainOracleOracleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainOracleOracleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainOracleOracleRemoved represents a OracleRemoved event raised by the OffchainOracle contract.
type OffchainOracleOracleRemoved struct {
	Oracle     common.Address
	OracleType uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOracleRemoved is a free log retrieval operation binding the contract event 0x7a7f56716fe703fb190529c336e57df71ab88188ba47e8d786bac684b61ab9a6.
//
// Solidity: event OracleRemoved(address oracle, uint8 oracleType)
func (_OffchainOracle *OffchainOracleFilterer) FilterOracleRemoved(opts *bind.FilterOpts) (*OffchainOracleOracleRemovedIterator, error) {

	logs, sub, err := _OffchainOracle.contract.FilterLogs(opts, "OracleRemoved")
	if err != nil {
		return nil, err
	}
	return &OffchainOracleOracleRemovedIterator{contract: _OffchainOracle.contract, event: "OracleRemoved", logs: logs, sub: sub}, nil
}

// WatchOracleRemoved is a free log subscription operation binding the contract event 0x7a7f56716fe703fb190529c336e57df71ab88188ba47e8d786bac684b61ab9a6.
//
// Solidity: event OracleRemoved(address oracle, uint8 oracleType)
func (_OffchainOracle *OffchainOracleFilterer) WatchOracleRemoved(opts *bind.WatchOpts, sink chan<- *OffchainOracleOracleRemoved) (event.Subscription, error) {

	logs, sub, err := _OffchainOracle.contract.WatchLogs(opts, "OracleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainOracleOracleRemoved)
				if err := _OffchainOracle.contract.UnpackLog(event, "OracleRemoved", log); err != nil {
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

// ParseOracleRemoved is a log parse operation binding the contract event 0x7a7f56716fe703fb190529c336e57df71ab88188ba47e8d786bac684b61ab9a6.
//
// Solidity: event OracleRemoved(address oracle, uint8 oracleType)
func (_OffchainOracle *OffchainOracleFilterer) ParseOracleRemoved(log types.Log) (*OffchainOracleOracleRemoved, error) {
	event := new(OffchainOracleOracleRemoved)
	if err := _OffchainOracle.contract.UnpackLog(event, "OracleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OffchainOracle contract.
type OffchainOracleOwnershipTransferredIterator struct {
	Event *OffchainOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OffchainOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainOracleOwnershipTransferred)
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
		it.Event = new(OffchainOracleOwnershipTransferred)
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
func (it *OffchainOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainOracleOwnershipTransferred represents a OwnershipTransferred event raised by the OffchainOracle contract.
type OffchainOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OffchainOracle *OffchainOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OffchainOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OffchainOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OffchainOracleOwnershipTransferredIterator{contract: _OffchainOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OffchainOracle *OffchainOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffchainOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OffchainOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainOracleOwnershipTransferred)
				if err := _OffchainOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OffchainOracle *OffchainOracleFilterer) ParseOwnershipTransferred(log types.Log) (*OffchainOracleOwnershipTransferred, error) {
	event := new(OffchainOracleOwnershipTransferred)
	if err := _OffchainOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// TomatoTraceableProcessingInfo is an auto generated low-level Go binding around an user-defined struct.
type TomatoTraceableProcessingInfo struct {
	ProcessingDate     *big.Int
	SamplingEvaluation string
	Factory            common.Address
}

// TomatoTraceableTomatoBaseInfo is an auto generated low-level Go binding around an user-defined struct.
type TomatoTraceableTomatoBaseInfo struct {
	TomatoId          string
	ProductionArea    string
	HarvestDate       *big.Int
	HarvestEvaluation string
	Grower            common.Address
}

// TomatoTraceableTransportInfo is an auto generated low-level Go binding around an user-defined struct.
type TomatoTraceableTransportInfo struct {
	TransportInfo string
	Logistic      common.Address
}

// TomatoTraceableABI is the input ABI used to generate the binding from.
const TomatoTraceableABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"evaluation\",\"type\":\"string\"}],\"name\":\"HarvestInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"evaluation\",\"type\":\"string\"}],\"name\":\"ProcessingInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"evaluation\",\"type\":\"string\"}],\"name\":\"QualityInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTomatoTraceable.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"productionArea\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processingDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"transportInfo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"harvestEvaluation\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"samplingEvaluation\",\"type\":\"string\"}],\"name\":\"TomatoDataSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grower\",\"type\":\"address\"}],\"name\":\"TomatoInfoAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TomatoInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"TransportInfoUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productionArea\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"harvestDateStr\",\"type\":\"string\"}],\"name\":\"addTomatoInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allTomatoIds\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"}],\"name\":\"canViewTomatoInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTomatoIds\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserRole\",\"outputs\":[{\"internalType\":\"enumTomatoTraceable.Role\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserTomatoes\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"processingInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"processingDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"samplingEvaluation\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"}],\"name\":\"queryTomatoInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productionArea\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"harvestDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"harvestEvaluation\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"grower\",\"type\":\"address\"}],\"internalType\":\"structTomatoTraceable.TomatoBaseInfo\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"processingDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"samplingEvaluation\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structTomatoTraceable.ProcessingInfo\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"transportInfo\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"logistic\",\"type\":\"address\"}],\"internalType\":\"structTomatoTraceable.TransportInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTomatoTraceable.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"registerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tomatoBaseInfos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productionArea\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"harvestDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"harvestEvaluation\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"grower\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"transportInfos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"transportInfo\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"logistic\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"harvestEvaluation\",\"type\":\"string\"}],\"name\":\"updateHarvestEvaluation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productionArea\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"harvestDateStr\",\"type\":\"string\"}],\"name\":\"updateHarvestInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processingDateStr\",\"type\":\"string\"}],\"name\":\"updateProcessingInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"harvestEvaluation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"samplingEvaluation\",\"type\":\"string\"}],\"name\":\"updateQualityInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"samplingEvaluation\",\"type\":\"string\"}],\"name\":\"updateSamplingEvaluation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tomatoId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"transportInfo\",\"type\":\"string\"}],\"name\":\"updateTransportInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRoles\",\"outputs\":[{\"internalType\":\"enumTomatoTraceable.Role\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userTomatoes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TomatoTraceable is an auto generated Go binding around a Solidity contract.
type TomatoTraceable struct {
	TomatoTraceableCaller     // Read-only binding to the contract
	TomatoTraceableTransactor // Write-only binding to the contract
	TomatoTraceableFilterer   // Log filterer for contract events
}

// TomatoTraceableCaller is an auto generated read-only Go binding around a Solidity contract.
type TomatoTraceableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TomatoTraceableTransactor is an auto generated write-only Go binding around a Solidity contract.
type TomatoTraceableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TomatoTraceableFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TomatoTraceableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TomatoTraceableSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TomatoTraceableSession struct {
	Contract     *TomatoTraceable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TomatoTraceableCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TomatoTraceableCallerSession struct {
	Contract *TomatoTraceableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TomatoTraceableTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TomatoTraceableTransactorSession struct {
	Contract     *TomatoTraceableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TomatoTraceableRaw is an auto generated low-level Go binding around a Solidity contract.
type TomatoTraceableRaw struct {
	Contract *TomatoTraceable // Generic contract binding to access the raw methods on
}

// TomatoTraceableCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TomatoTraceableCallerRaw struct {
	Contract *TomatoTraceableCaller // Generic read-only contract binding to access the raw methods on
}

// TomatoTraceableTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TomatoTraceableTransactorRaw struct {
	Contract *TomatoTraceableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTomatoTraceable creates a new instance of TomatoTraceable, bound to a specific deployed contract.
func NewTomatoTraceable(address common.Address, backend bind.ContractBackend) (*TomatoTraceable, error) {
	contract, err := bindTomatoTraceable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TomatoTraceable{TomatoTraceableCaller: TomatoTraceableCaller{contract: contract}, TomatoTraceableTransactor: TomatoTraceableTransactor{contract: contract}, TomatoTraceableFilterer: TomatoTraceableFilterer{contract: contract}}, nil
}

// NewTomatoTraceableCaller creates a new read-only instance of TomatoTraceable, bound to a specific deployed contract.
func NewTomatoTraceableCaller(address common.Address, caller bind.ContractCaller) (*TomatoTraceableCaller, error) {
	contract, err := bindTomatoTraceable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TomatoTraceableCaller{contract: contract}, nil
}

// NewTomatoTraceableTransactor creates a new write-only instance of TomatoTraceable, bound to a specific deployed contract.
func NewTomatoTraceableTransactor(address common.Address, transactor bind.ContractTransactor) (*TomatoTraceableTransactor, error) {
	contract, err := bindTomatoTraceable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TomatoTraceableTransactor{contract: contract}, nil
}

// NewTomatoTraceableFilterer creates a new log filterer instance of TomatoTraceable, bound to a specific deployed contract.
func NewTomatoTraceableFilterer(address common.Address, filterer bind.ContractFilterer) (*TomatoTraceableFilterer, error) {
	contract, err := bindTomatoTraceable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TomatoTraceableFilterer{contract: contract}, nil
}

// bindTomatoTraceable binds a generic wrapper to an already deployed contract.
func bindTomatoTraceable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TomatoTraceableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TomatoTraceable *TomatoTraceableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TomatoTraceable.Contract.TomatoTraceableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TomatoTraceable *TomatoTraceableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.TomatoTraceableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TomatoTraceable *TomatoTraceableRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.TomatoTraceableTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TomatoTraceable *TomatoTraceableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TomatoTraceable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TomatoTraceable *TomatoTraceableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TomatoTraceable *TomatoTraceableTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// AllTomatoIds is a free data retrieval call binding the contract method 0x4c681b8a.
//
// Solidity: function allTomatoIds(uint256 ) constant returns(string)
func (_TomatoTraceable *TomatoTraceableCaller) AllTomatoIds(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TomatoTraceable.contract.Call(opts, out, "allTomatoIds", arg0)
	return *ret0, err
}

// AllTomatoIds is a free data retrieval call binding the contract method 0x4c681b8a.
//
// Solidity: function allTomatoIds(uint256 ) constant returns(string)
func (_TomatoTraceable *TomatoTraceableSession) AllTomatoIds(arg0 *big.Int) (string, error) {
	return _TomatoTraceable.Contract.AllTomatoIds(&_TomatoTraceable.CallOpts, arg0)
}

// AllTomatoIds is a free data retrieval call binding the contract method 0x4c681b8a.
//
// Solidity: function allTomatoIds(uint256 ) constant returns(string)
func (_TomatoTraceable *TomatoTraceableCallerSession) AllTomatoIds(arg0 *big.Int) (string, error) {
	return _TomatoTraceable.Contract.AllTomatoIds(&_TomatoTraceable.CallOpts, arg0)
}

// CanViewTomatoInfo is a free data retrieval call binding the contract method 0x7ad3e799.
//
// Solidity: function canViewTomatoInfo(string tomatoId) constant returns(bool)
func (_TomatoTraceable *TomatoTraceableCaller) CanViewTomatoInfo(opts *bind.CallOpts, tomatoId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TomatoTraceable.contract.Call(opts, out, "canViewTomatoInfo", tomatoId)
	return *ret0, err
}

// CanViewTomatoInfo is a free data retrieval call binding the contract method 0x7ad3e799.
//
// Solidity: function canViewTomatoInfo(string tomatoId) constant returns(bool)
func (_TomatoTraceable *TomatoTraceableSession) CanViewTomatoInfo(tomatoId string) (bool, error) {
	return _TomatoTraceable.Contract.CanViewTomatoInfo(&_TomatoTraceable.CallOpts, tomatoId)
}

// CanViewTomatoInfo is a free data retrieval call binding the contract method 0x7ad3e799.
//
// Solidity: function canViewTomatoInfo(string tomatoId) constant returns(bool)
func (_TomatoTraceable *TomatoTraceableCallerSession) CanViewTomatoInfo(tomatoId string) (bool, error) {
	return _TomatoTraceable.Contract.CanViewTomatoInfo(&_TomatoTraceable.CallOpts, tomatoId)
}

// GetAllTomatoIds is a free data retrieval call binding the contract method 0xc6ac9871.
//
// Solidity: function getAllTomatoIds() constant returns(string[])
func (_TomatoTraceable *TomatoTraceableCaller) GetAllTomatoIds(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _TomatoTraceable.contract.Call(opts, out, "getAllTomatoIds")
	return *ret0, err
}

// GetAllTomatoIds is a free data retrieval call binding the contract method 0xc6ac9871.
//
// Solidity: function getAllTomatoIds() constant returns(string[])
func (_TomatoTraceable *TomatoTraceableSession) GetAllTomatoIds() ([]string, error) {
	return _TomatoTraceable.Contract.GetAllTomatoIds(&_TomatoTraceable.CallOpts)
}

// GetAllTomatoIds is a free data retrieval call binding the contract method 0xc6ac9871.
//
// Solidity: function getAllTomatoIds() constant returns(string[])
func (_TomatoTraceable *TomatoTraceableCallerSession) GetAllTomatoIds() ([]string, error) {
	return _TomatoTraceable.Contract.GetAllTomatoIds(&_TomatoTraceable.CallOpts)
}

// GetUserRole is a free data retrieval call binding the contract method 0x27820851.
//
// Solidity: function getUserRole(address user) constant returns(uint8)
func (_TomatoTraceable *TomatoTraceableCaller) GetUserRole(opts *bind.CallOpts, user common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TomatoTraceable.contract.Call(opts, out, "getUserRole", user)
	return *ret0, err
}

// GetUserRole is a free data retrieval call binding the contract method 0x27820851.
//
// Solidity: function getUserRole(address user) constant returns(uint8)
func (_TomatoTraceable *TomatoTraceableSession) GetUserRole(user common.Address) (uint8, error) {
	return _TomatoTraceable.Contract.GetUserRole(&_TomatoTraceable.CallOpts, user)
}

// GetUserRole is a free data retrieval call binding the contract method 0x27820851.
//
// Solidity: function getUserRole(address user) constant returns(uint8)
func (_TomatoTraceable *TomatoTraceableCallerSession) GetUserRole(user common.Address) (uint8, error) {
	return _TomatoTraceable.Contract.GetUserRole(&_TomatoTraceable.CallOpts, user)
}

// GetUserTomatoes is a free data retrieval call binding the contract method 0x264b73f5.
//
// Solidity: function getUserTomatoes() constant returns(string[])
func (_TomatoTraceable *TomatoTraceableCaller) GetUserTomatoes(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _TomatoTraceable.contract.Call(opts, out, "getUserTomatoes")
	return *ret0, err
}

// GetUserTomatoes is a free data retrieval call binding the contract method 0x264b73f5.
//
// Solidity: function getUserTomatoes() constant returns(string[])
func (_TomatoTraceable *TomatoTraceableSession) GetUserTomatoes() ([]string, error) {
	return _TomatoTraceable.Contract.GetUserTomatoes(&_TomatoTraceable.CallOpts)
}

// GetUserTomatoes is a free data retrieval call binding the contract method 0x264b73f5.
//
// Solidity: function getUserTomatoes() constant returns(string[])
func (_TomatoTraceable *TomatoTraceableCallerSession) GetUserTomatoes() ([]string, error) {
	return _TomatoTraceable.Contract.GetUserTomatoes(&_TomatoTraceable.CallOpts)
}

// ProcessingInfos is a free data retrieval call binding the contract method 0x194df91e.
//
// Solidity: function processingInfos(string ) constant returns(uint256 processingDate, string samplingEvaluation, address factory)
func (_TomatoTraceable *TomatoTraceableCaller) ProcessingInfos(opts *bind.CallOpts, arg0 string) (struct {
	ProcessingDate     *big.Int
	SamplingEvaluation string
	Factory            common.Address
}, error) {
	ret := new(struct {
		ProcessingDate     *big.Int
		SamplingEvaluation string
		Factory            common.Address
	})
	out := ret
	err := _TomatoTraceable.contract.Call(opts, out, "processingInfos", arg0)
	return *ret, err
}

// ProcessingInfos is a free data retrieval call binding the contract method 0x194df91e.
//
// Solidity: function processingInfos(string ) constant returns(uint256 processingDate, string samplingEvaluation, address factory)
func (_TomatoTraceable *TomatoTraceableSession) ProcessingInfos(arg0 string) (struct {
	ProcessingDate     *big.Int
	SamplingEvaluation string
	Factory            common.Address
}, error) {
	return _TomatoTraceable.Contract.ProcessingInfos(&_TomatoTraceable.CallOpts, arg0)
}

// ProcessingInfos is a free data retrieval call binding the contract method 0x194df91e.
//
// Solidity: function processingInfos(string ) constant returns(uint256 processingDate, string samplingEvaluation, address factory)
func (_TomatoTraceable *TomatoTraceableCallerSession) ProcessingInfos(arg0 string) (struct {
	ProcessingDate     *big.Int
	SamplingEvaluation string
	Factory            common.Address
}, error) {
	return _TomatoTraceable.Contract.ProcessingInfos(&_TomatoTraceable.CallOpts, arg0)
}

// QueryTomatoInfo is a free data retrieval call binding the contract method 0xf13d9ca4.
//
// Solidity: function queryTomatoInfo(address user, string tomatoId) constant returns(TomatoTraceableTomatoBaseInfo, TomatoTraceableProcessingInfo, TomatoTraceableTransportInfo)
func (_TomatoTraceable *TomatoTraceableCaller) QueryTomatoInfo(opts *bind.CallOpts, user common.Address, tomatoId string) (TomatoTraceableTomatoBaseInfo, TomatoTraceableProcessingInfo, TomatoTraceableTransportInfo, error) {
	var (
		ret0 = new(TomatoTraceableTomatoBaseInfo)
		ret1 = new(TomatoTraceableProcessingInfo)
		ret2 = new(TomatoTraceableTransportInfo)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _TomatoTraceable.contract.Call(opts, out, "queryTomatoInfo", user, tomatoId)
	return *ret0, *ret1, *ret2, err
}

// QueryTomatoInfo is a free data retrieval call binding the contract method 0xf13d9ca4.
//
// Solidity: function queryTomatoInfo(address user, string tomatoId) constant returns(TomatoTraceableTomatoBaseInfo, TomatoTraceableProcessingInfo, TomatoTraceableTransportInfo)
func (_TomatoTraceable *TomatoTraceableSession) QueryTomatoInfo(user common.Address, tomatoId string) (TomatoTraceableTomatoBaseInfo, TomatoTraceableProcessingInfo, TomatoTraceableTransportInfo, error) {
	return _TomatoTraceable.Contract.QueryTomatoInfo(&_TomatoTraceable.CallOpts, user, tomatoId)
}

// QueryTomatoInfo is a free data retrieval call binding the contract method 0xf13d9ca4.
//
// Solidity: function queryTomatoInfo(address user, string tomatoId) constant returns(TomatoTraceableTomatoBaseInfo, TomatoTraceableProcessingInfo, TomatoTraceableTransportInfo)
func (_TomatoTraceable *TomatoTraceableCallerSession) QueryTomatoInfo(user common.Address, tomatoId string) (TomatoTraceableTomatoBaseInfo, TomatoTraceableProcessingInfo, TomatoTraceableTransportInfo, error) {
	return _TomatoTraceable.Contract.QueryTomatoInfo(&_TomatoTraceable.CallOpts, user, tomatoId)
}

// TomatoBaseInfos is a free data retrieval call binding the contract method 0x5777d006.
//
// Solidity: function tomatoBaseInfos(string ) constant returns(string tomatoId, string productionArea, uint256 harvestDate, string harvestEvaluation, address grower)
func (_TomatoTraceable *TomatoTraceableCaller) TomatoBaseInfos(opts *bind.CallOpts, arg0 string) (struct {
	TomatoId          string
	ProductionArea    string
	HarvestDate       *big.Int
	HarvestEvaluation string
	Grower            common.Address
}, error) {
	ret := new(struct {
		TomatoId          string
		ProductionArea    string
		HarvestDate       *big.Int
		HarvestEvaluation string
		Grower            common.Address
	})
	out := ret
	err := _TomatoTraceable.contract.Call(opts, out, "tomatoBaseInfos", arg0)
	return *ret, err
}

// TomatoBaseInfos is a free data retrieval call binding the contract method 0x5777d006.
//
// Solidity: function tomatoBaseInfos(string ) constant returns(string tomatoId, string productionArea, uint256 harvestDate, string harvestEvaluation, address grower)
func (_TomatoTraceable *TomatoTraceableSession) TomatoBaseInfos(arg0 string) (struct {
	TomatoId          string
	ProductionArea    string
	HarvestDate       *big.Int
	HarvestEvaluation string
	Grower            common.Address
}, error) {
	return _TomatoTraceable.Contract.TomatoBaseInfos(&_TomatoTraceable.CallOpts, arg0)
}

// TomatoBaseInfos is a free data retrieval call binding the contract method 0x5777d006.
//
// Solidity: function tomatoBaseInfos(string ) constant returns(string tomatoId, string productionArea, uint256 harvestDate, string harvestEvaluation, address grower)
func (_TomatoTraceable *TomatoTraceableCallerSession) TomatoBaseInfos(arg0 string) (struct {
	TomatoId          string
	ProductionArea    string
	HarvestDate       *big.Int
	HarvestEvaluation string
	Grower            common.Address
}, error) {
	return _TomatoTraceable.Contract.TomatoBaseInfos(&_TomatoTraceable.CallOpts, arg0)
}

// TransportInfos is a free data retrieval call binding the contract method 0xe68be622.
//
// Solidity: function transportInfos(string ) constant returns(string transportInfo, address logistic)
func (_TomatoTraceable *TomatoTraceableCaller) TransportInfos(opts *bind.CallOpts, arg0 string) (struct {
	TransportInfo string
	Logistic      common.Address
}, error) {
	ret := new(struct {
		TransportInfo string
		Logistic      common.Address
	})
	out := ret
	err := _TomatoTraceable.contract.Call(opts, out, "transportInfos", arg0)
	return *ret, err
}

// TransportInfos is a free data retrieval call binding the contract method 0xe68be622.
//
// Solidity: function transportInfos(string ) constant returns(string transportInfo, address logistic)
func (_TomatoTraceable *TomatoTraceableSession) TransportInfos(arg0 string) (struct {
	TransportInfo string
	Logistic      common.Address
}, error) {
	return _TomatoTraceable.Contract.TransportInfos(&_TomatoTraceable.CallOpts, arg0)
}

// TransportInfos is a free data retrieval call binding the contract method 0xe68be622.
//
// Solidity: function transportInfos(string ) constant returns(string transportInfo, address logistic)
func (_TomatoTraceable *TomatoTraceableCallerSession) TransportInfos(arg0 string) (struct {
	TransportInfo string
	Logistic      common.Address
}, error) {
	return _TomatoTraceable.Contract.TransportInfos(&_TomatoTraceable.CallOpts, arg0)
}

// UserRoles is a free data retrieval call binding the contract method 0x74d5e100.
//
// Solidity: function userRoles(address ) constant returns(uint8)
func (_TomatoTraceable *TomatoTraceableCaller) UserRoles(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TomatoTraceable.contract.Call(opts, out, "userRoles", arg0)
	return *ret0, err
}

// UserRoles is a free data retrieval call binding the contract method 0x74d5e100.
//
// Solidity: function userRoles(address ) constant returns(uint8)
func (_TomatoTraceable *TomatoTraceableSession) UserRoles(arg0 common.Address) (uint8, error) {
	return _TomatoTraceable.Contract.UserRoles(&_TomatoTraceable.CallOpts, arg0)
}

// UserRoles is a free data retrieval call binding the contract method 0x74d5e100.
//
// Solidity: function userRoles(address ) constant returns(uint8)
func (_TomatoTraceable *TomatoTraceableCallerSession) UserRoles(arg0 common.Address) (uint8, error) {
	return _TomatoTraceable.Contract.UserRoles(&_TomatoTraceable.CallOpts, arg0)
}

// UserTomatoes is a free data retrieval call binding the contract method 0x5277dd60.
//
// Solidity: function userTomatoes(address , uint256 ) constant returns(string)
func (_TomatoTraceable *TomatoTraceableCaller) UserTomatoes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TomatoTraceable.contract.Call(opts, out, "userTomatoes", arg0, arg1)
	return *ret0, err
}

// UserTomatoes is a free data retrieval call binding the contract method 0x5277dd60.
//
// Solidity: function userTomatoes(address , uint256 ) constant returns(string)
func (_TomatoTraceable *TomatoTraceableSession) UserTomatoes(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _TomatoTraceable.Contract.UserTomatoes(&_TomatoTraceable.CallOpts, arg0, arg1)
}

// UserTomatoes is a free data retrieval call binding the contract method 0x5277dd60.
//
// Solidity: function userTomatoes(address , uint256 ) constant returns(string)
func (_TomatoTraceable *TomatoTraceableCallerSession) UserTomatoes(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _TomatoTraceable.Contract.UserTomatoes(&_TomatoTraceable.CallOpts, arg0, arg1)
}

// AddTomatoInfo is a paid mutator transaction binding the contract method 0x33b237ea.
//
// Solidity: function addTomatoInfo(string productionArea, string harvestDateStr) returns(string)
func (_TomatoTraceable *TomatoTraceableTransactor) AddTomatoInfo(opts *bind.TransactOpts, productionArea string, harvestDateStr string) (string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	transaction, receipt, err := _TomatoTraceable.contract.TransactWithResult(opts, out, "addTomatoInfo", productionArea, harvestDateStr)
	return *ret0, transaction, receipt, err
}

func (_TomatoTraceable *TomatoTraceableTransactor) AsyncAddTomatoInfo(handler func(*types.Receipt, error), opts *bind.TransactOpts, productionArea string, harvestDateStr string) (*types.Transaction, error) {
	return _TomatoTraceable.contract.AsyncTransact(opts, handler, "addTomatoInfo", productionArea, harvestDateStr)
}

// AddTomatoInfo is a paid mutator transaction binding the contract method 0x33b237ea.
//
// Solidity: function addTomatoInfo(string productionArea, string harvestDateStr) returns(string)
func (_TomatoTraceable *TomatoTraceableSession) AddTomatoInfo(productionArea string, harvestDateStr string) (string, *types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.AddTomatoInfo(&_TomatoTraceable.TransactOpts, productionArea, harvestDateStr)
}

func (_TomatoTraceable *TomatoTraceableSession) AsyncAddTomatoInfo(handler func(*types.Receipt, error), productionArea string, harvestDateStr string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncAddTomatoInfo(handler, &_TomatoTraceable.TransactOpts, productionArea, harvestDateStr)
}

// AddTomatoInfo is a paid mutator transaction binding the contract method 0x33b237ea.
//
// Solidity: function addTomatoInfo(string productionArea, string harvestDateStr) returns(string)
func (_TomatoTraceable *TomatoTraceableTransactorSession) AddTomatoInfo(productionArea string, harvestDateStr string) (string, *types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.AddTomatoInfo(&_TomatoTraceable.TransactOpts, productionArea, harvestDateStr)
}

func (_TomatoTraceable *TomatoTraceableTransactorSession) AsyncAddTomatoInfo(handler func(*types.Receipt, error), productionArea string, harvestDateStr string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncAddTomatoInfo(handler, &_TomatoTraceable.TransactOpts, productionArea, harvestDateStr)
}

// RegisterRole is a paid mutator transaction binding the contract method 0x0a52c10a.
//
// Solidity: function registerRole(uint8 role) returns()
func (_TomatoTraceable *TomatoTraceableTransactor) RegisterRole(opts *bind.TransactOpts, role uint8) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _TomatoTraceable.contract.TransactWithResult(opts, out, "registerRole", role)
	return transaction, receipt, err
}

func (_TomatoTraceable *TomatoTraceableTransactor) AsyncRegisterRole(handler func(*types.Receipt, error), opts *bind.TransactOpts, role uint8) (*types.Transaction, error) {
	return _TomatoTraceable.contract.AsyncTransact(opts, handler, "registerRole", role)
}

// RegisterRole is a paid mutator transaction binding the contract method 0x0a52c10a.
//
// Solidity: function registerRole(uint8 role) returns()
func (_TomatoTraceable *TomatoTraceableSession) RegisterRole(role uint8) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.RegisterRole(&_TomatoTraceable.TransactOpts, role)
}

func (_TomatoTraceable *TomatoTraceableSession) AsyncRegisterRole(handler func(*types.Receipt, error), role uint8) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncRegisterRole(handler, &_TomatoTraceable.TransactOpts, role)
}

// RegisterRole is a paid mutator transaction binding the contract method 0x0a52c10a.
//
// Solidity: function registerRole(uint8 role) returns()
func (_TomatoTraceable *TomatoTraceableTransactorSession) RegisterRole(role uint8) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.RegisterRole(&_TomatoTraceable.TransactOpts, role)
}

func (_TomatoTraceable *TomatoTraceableTransactorSession) AsyncRegisterRole(handler func(*types.Receipt, error), role uint8) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncRegisterRole(handler, &_TomatoTraceable.TransactOpts, role)
}

// UpdateHarvestEvaluation is a paid mutator transaction binding the contract method 0x761716f8.
//
// Solidity: function updateHarvestEvaluation(string tomatoId, string harvestEvaluation) returns()
func (_TomatoTraceable *TomatoTraceableTransactor) UpdateHarvestEvaluation(opts *bind.TransactOpts, tomatoId string, harvestEvaluation string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _TomatoTraceable.contract.TransactWithResult(opts, out, "updateHarvestEvaluation", tomatoId, harvestEvaluation)
	return transaction, receipt, err
}

func (_TomatoTraceable *TomatoTraceableTransactor) AsyncUpdateHarvestEvaluation(handler func(*types.Receipt, error), opts *bind.TransactOpts, tomatoId string, harvestEvaluation string) (*types.Transaction, error) {
	return _TomatoTraceable.contract.AsyncTransact(opts, handler, "updateHarvestEvaluation", tomatoId, harvestEvaluation)
}

// UpdateHarvestEvaluation is a paid mutator transaction binding the contract method 0x761716f8.
//
// Solidity: function updateHarvestEvaluation(string tomatoId, string harvestEvaluation) returns()
func (_TomatoTraceable *TomatoTraceableSession) UpdateHarvestEvaluation(tomatoId string, harvestEvaluation string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateHarvestEvaluation(&_TomatoTraceable.TransactOpts, tomatoId, harvestEvaluation)
}

func (_TomatoTraceable *TomatoTraceableSession) AsyncUpdateHarvestEvaluation(handler func(*types.Receipt, error), tomatoId string, harvestEvaluation string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateHarvestEvaluation(handler, &_TomatoTraceable.TransactOpts, tomatoId, harvestEvaluation)
}

// UpdateHarvestEvaluation is a paid mutator transaction binding the contract method 0x761716f8.
//
// Solidity: function updateHarvestEvaluation(string tomatoId, string harvestEvaluation) returns()
func (_TomatoTraceable *TomatoTraceableTransactorSession) UpdateHarvestEvaluation(tomatoId string, harvestEvaluation string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateHarvestEvaluation(&_TomatoTraceable.TransactOpts, tomatoId, harvestEvaluation)
}

func (_TomatoTraceable *TomatoTraceableTransactorSession) AsyncUpdateHarvestEvaluation(handler func(*types.Receipt, error), tomatoId string, harvestEvaluation string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateHarvestEvaluation(handler, &_TomatoTraceable.TransactOpts, tomatoId, harvestEvaluation)
}

// UpdateHarvestInfo is a paid mutator transaction binding the contract method 0x791ef9ae.
//
// Solidity: function updateHarvestInfo(string tomatoId, string productionArea, string harvestDateStr) returns()
func (_TomatoTraceable *TomatoTraceableTransactor) UpdateHarvestInfo(opts *bind.TransactOpts, tomatoId string, productionArea string, harvestDateStr string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _TomatoTraceable.contract.TransactWithResult(opts, out, "updateHarvestInfo", tomatoId, productionArea, harvestDateStr)
	return transaction, receipt, err
}

func (_TomatoTraceable *TomatoTraceableTransactor) AsyncUpdateHarvestInfo(handler func(*types.Receipt, error), opts *bind.TransactOpts, tomatoId string, productionArea string, harvestDateStr string) (*types.Transaction, error) {
	return _TomatoTraceable.contract.AsyncTransact(opts, handler, "updateHarvestInfo", tomatoId, productionArea, harvestDateStr)
}

// UpdateHarvestInfo is a paid mutator transaction binding the contract method 0x791ef9ae.
//
// Solidity: function updateHarvestInfo(string tomatoId, string productionArea, string harvestDateStr) returns()
func (_TomatoTraceable *TomatoTraceableSession) UpdateHarvestInfo(tomatoId string, productionArea string, harvestDateStr string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateHarvestInfo(&_TomatoTraceable.TransactOpts, tomatoId, productionArea, harvestDateStr)
}

func (_TomatoTraceable *TomatoTraceableSession) AsyncUpdateHarvestInfo(handler func(*types.Receipt, error), tomatoId string, productionArea string, harvestDateStr string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateHarvestInfo(handler, &_TomatoTraceable.TransactOpts, tomatoId, productionArea, harvestDateStr)
}

// UpdateHarvestInfo is a paid mutator transaction binding the contract method 0x791ef9ae.
//
// Solidity: function updateHarvestInfo(string tomatoId, string productionArea, string harvestDateStr) returns()
func (_TomatoTraceable *TomatoTraceableTransactorSession) UpdateHarvestInfo(tomatoId string, productionArea string, harvestDateStr string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateHarvestInfo(&_TomatoTraceable.TransactOpts, tomatoId, productionArea, harvestDateStr)
}

func (_TomatoTraceable *TomatoTraceableTransactorSession) AsyncUpdateHarvestInfo(handler func(*types.Receipt, error), tomatoId string, productionArea string, harvestDateStr string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateHarvestInfo(handler, &_TomatoTraceable.TransactOpts, tomatoId, productionArea, harvestDateStr)
}

// UpdateProcessingInfo is a paid mutator transaction binding the contract method 0x55c9bbe8.
//
// Solidity: function updateProcessingInfo(string tomatoId, string processingDateStr) returns()
func (_TomatoTraceable *TomatoTraceableTransactor) UpdateProcessingInfo(opts *bind.TransactOpts, tomatoId string, processingDateStr string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _TomatoTraceable.contract.TransactWithResult(opts, out, "updateProcessingInfo", tomatoId, processingDateStr)
	return transaction, receipt, err
}

func (_TomatoTraceable *TomatoTraceableTransactor) AsyncUpdateProcessingInfo(handler func(*types.Receipt, error), opts *bind.TransactOpts, tomatoId string, processingDateStr string) (*types.Transaction, error) {
	return _TomatoTraceable.contract.AsyncTransact(opts, handler, "updateProcessingInfo", tomatoId, processingDateStr)
}

// UpdateProcessingInfo is a paid mutator transaction binding the contract method 0x55c9bbe8.
//
// Solidity: function updateProcessingInfo(string tomatoId, string processingDateStr) returns()
func (_TomatoTraceable *TomatoTraceableSession) UpdateProcessingInfo(tomatoId string, processingDateStr string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateProcessingInfo(&_TomatoTraceable.TransactOpts, tomatoId, processingDateStr)
}

func (_TomatoTraceable *TomatoTraceableSession) AsyncUpdateProcessingInfo(handler func(*types.Receipt, error), tomatoId string, processingDateStr string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateProcessingInfo(handler, &_TomatoTraceable.TransactOpts, tomatoId, processingDateStr)
}

// UpdateProcessingInfo is a paid mutator transaction binding the contract method 0x55c9bbe8.
//
// Solidity: function updateProcessingInfo(string tomatoId, string processingDateStr) returns()
func (_TomatoTraceable *TomatoTraceableTransactorSession) UpdateProcessingInfo(tomatoId string, processingDateStr string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateProcessingInfo(&_TomatoTraceable.TransactOpts, tomatoId, processingDateStr)
}

func (_TomatoTraceable *TomatoTraceableTransactorSession) AsyncUpdateProcessingInfo(handler func(*types.Receipt, error), tomatoId string, processingDateStr string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateProcessingInfo(handler, &_TomatoTraceable.TransactOpts, tomatoId, processingDateStr)
}

// UpdateQualityInfo is a paid mutator transaction binding the contract method 0xd2b10579.
//
// Solidity: function updateQualityInfo(string tomatoId, string harvestEvaluation, string samplingEvaluation) returns()
func (_TomatoTraceable *TomatoTraceableTransactor) UpdateQualityInfo(opts *bind.TransactOpts, tomatoId string, harvestEvaluation string, samplingEvaluation string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _TomatoTraceable.contract.TransactWithResult(opts, out, "updateQualityInfo", tomatoId, harvestEvaluation, samplingEvaluation)
	return transaction, receipt, err
}

func (_TomatoTraceable *TomatoTraceableTransactor) AsyncUpdateQualityInfo(handler func(*types.Receipt, error), opts *bind.TransactOpts, tomatoId string, harvestEvaluation string, samplingEvaluation string) (*types.Transaction, error) {
	return _TomatoTraceable.contract.AsyncTransact(opts, handler, "updateQualityInfo", tomatoId, harvestEvaluation, samplingEvaluation)
}

// UpdateQualityInfo is a paid mutator transaction binding the contract method 0xd2b10579.
//
// Solidity: function updateQualityInfo(string tomatoId, string harvestEvaluation, string samplingEvaluation) returns()
func (_TomatoTraceable *TomatoTraceableSession) UpdateQualityInfo(tomatoId string, harvestEvaluation string, samplingEvaluation string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateQualityInfo(&_TomatoTraceable.TransactOpts, tomatoId, harvestEvaluation, samplingEvaluation)
}

func (_TomatoTraceable *TomatoTraceableSession) AsyncUpdateQualityInfo(handler func(*types.Receipt, error), tomatoId string, harvestEvaluation string, samplingEvaluation string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateQualityInfo(handler, &_TomatoTraceable.TransactOpts, tomatoId, harvestEvaluation, samplingEvaluation)
}

// UpdateQualityInfo is a paid mutator transaction binding the contract method 0xd2b10579.
//
// Solidity: function updateQualityInfo(string tomatoId, string harvestEvaluation, string samplingEvaluation) returns()
func (_TomatoTraceable *TomatoTraceableTransactorSession) UpdateQualityInfo(tomatoId string, harvestEvaluation string, samplingEvaluation string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateQualityInfo(&_TomatoTraceable.TransactOpts, tomatoId, harvestEvaluation, samplingEvaluation)
}

func (_TomatoTraceable *TomatoTraceableTransactorSession) AsyncUpdateQualityInfo(handler func(*types.Receipt, error), tomatoId string, harvestEvaluation string, samplingEvaluation string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateQualityInfo(handler, &_TomatoTraceable.TransactOpts, tomatoId, harvestEvaluation, samplingEvaluation)
}

// UpdateSamplingEvaluation is a paid mutator transaction binding the contract method 0xc818d316.
//
// Solidity: function updateSamplingEvaluation(string tomatoId, string samplingEvaluation) returns()
func (_TomatoTraceable *TomatoTraceableTransactor) UpdateSamplingEvaluation(opts *bind.TransactOpts, tomatoId string, samplingEvaluation string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _TomatoTraceable.contract.TransactWithResult(opts, out, "updateSamplingEvaluation", tomatoId, samplingEvaluation)
	return transaction, receipt, err
}

func (_TomatoTraceable *TomatoTraceableTransactor) AsyncUpdateSamplingEvaluation(handler func(*types.Receipt, error), opts *bind.TransactOpts, tomatoId string, samplingEvaluation string) (*types.Transaction, error) {
	return _TomatoTraceable.contract.AsyncTransact(opts, handler, "updateSamplingEvaluation", tomatoId, samplingEvaluation)
}

// UpdateSamplingEvaluation is a paid mutator transaction binding the contract method 0xc818d316.
//
// Solidity: function updateSamplingEvaluation(string tomatoId, string samplingEvaluation) returns()
func (_TomatoTraceable *TomatoTraceableSession) UpdateSamplingEvaluation(tomatoId string, samplingEvaluation string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateSamplingEvaluation(&_TomatoTraceable.TransactOpts, tomatoId, samplingEvaluation)
}

func (_TomatoTraceable *TomatoTraceableSession) AsyncUpdateSamplingEvaluation(handler func(*types.Receipt, error), tomatoId string, samplingEvaluation string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateSamplingEvaluation(handler, &_TomatoTraceable.TransactOpts, tomatoId, samplingEvaluation)
}

// UpdateSamplingEvaluation is a paid mutator transaction binding the contract method 0xc818d316.
//
// Solidity: function updateSamplingEvaluation(string tomatoId, string samplingEvaluation) returns()
func (_TomatoTraceable *TomatoTraceableTransactorSession) UpdateSamplingEvaluation(tomatoId string, samplingEvaluation string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateSamplingEvaluation(&_TomatoTraceable.TransactOpts, tomatoId, samplingEvaluation)
}

func (_TomatoTraceable *TomatoTraceableTransactorSession) AsyncUpdateSamplingEvaluation(handler func(*types.Receipt, error), tomatoId string, samplingEvaluation string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateSamplingEvaluation(handler, &_TomatoTraceable.TransactOpts, tomatoId, samplingEvaluation)
}

// UpdateTransportInfo is a paid mutator transaction binding the contract method 0x63e38f59.
//
// Solidity: function updateTransportInfo(string tomatoId, string transportInfo) returns()
func (_TomatoTraceable *TomatoTraceableTransactor) UpdateTransportInfo(opts *bind.TransactOpts, tomatoId string, transportInfo string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _TomatoTraceable.contract.TransactWithResult(opts, out, "updateTransportInfo", tomatoId, transportInfo)
	return transaction, receipt, err
}

func (_TomatoTraceable *TomatoTraceableTransactor) AsyncUpdateTransportInfo(handler func(*types.Receipt, error), opts *bind.TransactOpts, tomatoId string, transportInfo string) (*types.Transaction, error) {
	return _TomatoTraceable.contract.AsyncTransact(opts, handler, "updateTransportInfo", tomatoId, transportInfo)
}

// UpdateTransportInfo is a paid mutator transaction binding the contract method 0x63e38f59.
//
// Solidity: function updateTransportInfo(string tomatoId, string transportInfo) returns()
func (_TomatoTraceable *TomatoTraceableSession) UpdateTransportInfo(tomatoId string, transportInfo string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateTransportInfo(&_TomatoTraceable.TransactOpts, tomatoId, transportInfo)
}

func (_TomatoTraceable *TomatoTraceableSession) AsyncUpdateTransportInfo(handler func(*types.Receipt, error), tomatoId string, transportInfo string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateTransportInfo(handler, &_TomatoTraceable.TransactOpts, tomatoId, transportInfo)
}

// UpdateTransportInfo is a paid mutator transaction binding the contract method 0x63e38f59.
//
// Solidity: function updateTransportInfo(string tomatoId, string transportInfo) returns()
func (_TomatoTraceable *TomatoTraceableTransactorSession) UpdateTransportInfo(tomatoId string, transportInfo string) (*types.Transaction, *types.Receipt, error) {
	return _TomatoTraceable.Contract.UpdateTransportInfo(&_TomatoTraceable.TransactOpts, tomatoId, transportInfo)
}

func (_TomatoTraceable *TomatoTraceableTransactorSession) AsyncUpdateTransportInfo(handler func(*types.Receipt, error), tomatoId string, transportInfo string) (*types.Transaction, error) {
	return _TomatoTraceable.Contract.AsyncUpdateTransportInfo(handler, &_TomatoTraceable.TransactOpts, tomatoId, transportInfo)
}

// TomatoTraceableHarvestInfoUpdated represents a HarvestInfoUpdated event raised by the TomatoTraceable contract.
type TomatoTraceableHarvestInfoUpdated struct {
	TomatoId   common.Hash
	Evaluation string
	Raw        types.Log // Blockchain specific contextual infos
}

// WatchHarvestInfoUpdated is a free log subscription operation binding the contract event 0xdcada125e3cdffd8174e57901629c43442f2836645056ba7428081425edf3bf0.
//
// Solidity: event HarvestInfoUpdated(string indexed tomatoId, string evaluation)
func (_TomatoTraceable *TomatoTraceableFilterer) WatchHarvestInfoUpdated(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "HarvestInfoUpdated", tomatoId)
}

func (_TomatoTraceable *TomatoTraceableFilterer) WatchAllHarvestInfoUpdated(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "HarvestInfoUpdated")
}

// ParseHarvestInfoUpdated is a log parse operation binding the contract event 0xdcada125e3cdffd8174e57901629c43442f2836645056ba7428081425edf3bf0.
//
// Solidity: event HarvestInfoUpdated(string indexed tomatoId, string evaluation)
func (_TomatoTraceable *TomatoTraceableFilterer) ParseHarvestInfoUpdated(log types.Log) (*TomatoTraceableHarvestInfoUpdated, error) {
	event := new(TomatoTraceableHarvestInfoUpdated)
	if err := _TomatoTraceable.contract.UnpackLog(event, "HarvestInfoUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchHarvestInfoUpdated is a free log subscription operation binding the contract event 0xdcada125e3cdffd8174e57901629c43442f2836645056ba7428081425edf3bf0.
//
// Solidity: event HarvestInfoUpdated(string indexed tomatoId, string evaluation)
func (_TomatoTraceable *TomatoTraceableSession) WatchHarvestInfoUpdated(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.Contract.WatchHarvestInfoUpdated(fromBlock, handler, tomatoId)
}

func (_TomatoTraceable *TomatoTraceableSession) WatchAllHarvestInfoUpdated(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.Contract.WatchAllHarvestInfoUpdated(fromBlock, handler)
}

// ParseHarvestInfoUpdated is a log parse operation binding the contract event 0xdcada125e3cdffd8174e57901629c43442f2836645056ba7428081425edf3bf0.
//
// Solidity: event HarvestInfoUpdated(string indexed tomatoId, string evaluation)
func (_TomatoTraceable *TomatoTraceableSession) ParseHarvestInfoUpdated(log types.Log) (*TomatoTraceableHarvestInfoUpdated, error) {
	return _TomatoTraceable.Contract.ParseHarvestInfoUpdated(log)
}

// TomatoTraceableProcessingInfoUpdated represents a ProcessingInfoUpdated event raised by the TomatoTraceable contract.
type TomatoTraceableProcessingInfoUpdated struct {
	TomatoId   common.Hash
	Date       *big.Int
	Evaluation string
	Raw        types.Log // Blockchain specific contextual infos
}

// WatchProcessingInfoUpdated is a free log subscription operation binding the contract event 0x4a3b3019e2a92659ba19eaa7dfee02c918414c177faf8a78953400dd75806f47.
//
// Solidity: event ProcessingInfoUpdated(string indexed tomatoId, uint256 date, string evaluation)
func (_TomatoTraceable *TomatoTraceableFilterer) WatchProcessingInfoUpdated(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "ProcessingInfoUpdated", tomatoId)
}

func (_TomatoTraceable *TomatoTraceableFilterer) WatchAllProcessingInfoUpdated(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "ProcessingInfoUpdated")
}

// ParseProcessingInfoUpdated is a log parse operation binding the contract event 0x4a3b3019e2a92659ba19eaa7dfee02c918414c177faf8a78953400dd75806f47.
//
// Solidity: event ProcessingInfoUpdated(string indexed tomatoId, uint256 date, string evaluation)
func (_TomatoTraceable *TomatoTraceableFilterer) ParseProcessingInfoUpdated(log types.Log) (*TomatoTraceableProcessingInfoUpdated, error) {
	event := new(TomatoTraceableProcessingInfoUpdated)
	if err := _TomatoTraceable.contract.UnpackLog(event, "ProcessingInfoUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchProcessingInfoUpdated is a free log subscription operation binding the contract event 0x4a3b3019e2a92659ba19eaa7dfee02c918414c177faf8a78953400dd75806f47.
//
// Solidity: event ProcessingInfoUpdated(string indexed tomatoId, uint256 date, string evaluation)
func (_TomatoTraceable *TomatoTraceableSession) WatchProcessingInfoUpdated(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.Contract.WatchProcessingInfoUpdated(fromBlock, handler, tomatoId)
}

func (_TomatoTraceable *TomatoTraceableSession) WatchAllProcessingInfoUpdated(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.Contract.WatchAllProcessingInfoUpdated(fromBlock, handler)
}

// ParseProcessingInfoUpdated is a log parse operation binding the contract event 0x4a3b3019e2a92659ba19eaa7dfee02c918414c177faf8a78953400dd75806f47.
//
// Solidity: event ProcessingInfoUpdated(string indexed tomatoId, uint256 date, string evaluation)
func (_TomatoTraceable *TomatoTraceableSession) ParseProcessingInfoUpdated(log types.Log) (*TomatoTraceableProcessingInfoUpdated, error) {
	return _TomatoTraceable.Contract.ParseProcessingInfoUpdated(log)
}

// TomatoTraceableQualityInfoUpdated represents a QualityInfoUpdated event raised by the TomatoTraceable contract.
type TomatoTraceableQualityInfoUpdated struct {
	TomatoId   common.Hash
	Evaluation string
	Raw        types.Log // Blockchain specific contextual infos
}

// WatchQualityInfoUpdated is a free log subscription operation binding the contract event 0x6a877c936f46e6d0cdab44e3414edf881183f0492db65f207365d41e8bcf099e.
//
// Solidity: event QualityInfoUpdated(string indexed tomatoId, string evaluation)
func (_TomatoTraceable *TomatoTraceableFilterer) WatchQualityInfoUpdated(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "QualityInfoUpdated", tomatoId)
}

func (_TomatoTraceable *TomatoTraceableFilterer) WatchAllQualityInfoUpdated(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "QualityInfoUpdated")
}

// ParseQualityInfoUpdated is a log parse operation binding the contract event 0x6a877c936f46e6d0cdab44e3414edf881183f0492db65f207365d41e8bcf099e.
//
// Solidity: event QualityInfoUpdated(string indexed tomatoId, string evaluation)
func (_TomatoTraceable *TomatoTraceableFilterer) ParseQualityInfoUpdated(log types.Log) (*TomatoTraceableQualityInfoUpdated, error) {
	event := new(TomatoTraceableQualityInfoUpdated)
	if err := _TomatoTraceable.contract.UnpackLog(event, "QualityInfoUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchQualityInfoUpdated is a free log subscription operation binding the contract event 0x6a877c936f46e6d0cdab44e3414edf881183f0492db65f207365d41e8bcf099e.
//
// Solidity: event QualityInfoUpdated(string indexed tomatoId, string evaluation)
func (_TomatoTraceable *TomatoTraceableSession) WatchQualityInfoUpdated(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.Contract.WatchQualityInfoUpdated(fromBlock, handler, tomatoId)
}

func (_TomatoTraceable *TomatoTraceableSession) WatchAllQualityInfoUpdated(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.Contract.WatchAllQualityInfoUpdated(fromBlock, handler)
}

// ParseQualityInfoUpdated is a log parse operation binding the contract event 0x6a877c936f46e6d0cdab44e3414edf881183f0492db65f207365d41e8bcf099e.
//
// Solidity: event QualityInfoUpdated(string indexed tomatoId, string evaluation)
func (_TomatoTraceable *TomatoTraceableSession) ParseQualityInfoUpdated(log types.Log) (*TomatoTraceableQualityInfoUpdated, error) {
	return _TomatoTraceable.Contract.ParseQualityInfoUpdated(log)
}

// TomatoTraceableRoleRegistered represents a RoleRegistered event raised by the TomatoTraceable contract.
type TomatoTraceableRoleRegistered struct {
	User common.Address
	Role uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// WatchRoleRegistered is a free log subscription operation binding the contract event 0x1e9acfe4df89a29b2be615d44fd962d8ce1fd19ff4a74ddf5a9401a6a4a2a6d2.
//
// Solidity: event RoleRegistered(address indexed user, uint8 role)
func (_TomatoTraceable *TomatoTraceableFilterer) WatchRoleRegistered(fromBlock *uint64, handler func(int, []types.Log), user common.Address) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "RoleRegistered", user)
}

func (_TomatoTraceable *TomatoTraceableFilterer) WatchAllRoleRegistered(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "RoleRegistered")
}

// ParseRoleRegistered is a log parse operation binding the contract event 0x1e9acfe4df89a29b2be615d44fd962d8ce1fd19ff4a74ddf5a9401a6a4a2a6d2.
//
// Solidity: event RoleRegistered(address indexed user, uint8 role)
func (_TomatoTraceable *TomatoTraceableFilterer) ParseRoleRegistered(log types.Log) (*TomatoTraceableRoleRegistered, error) {
	event := new(TomatoTraceableRoleRegistered)
	if err := _TomatoTraceable.contract.UnpackLog(event, "RoleRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchRoleRegistered is a free log subscription operation binding the contract event 0x1e9acfe4df89a29b2be615d44fd962d8ce1fd19ff4a74ddf5a9401a6a4a2a6d2.
//
// Solidity: event RoleRegistered(address indexed user, uint8 role)
func (_TomatoTraceable *TomatoTraceableSession) WatchRoleRegistered(fromBlock *uint64, handler func(int, []types.Log), user common.Address) (string, error) {
	return _TomatoTraceable.Contract.WatchRoleRegistered(fromBlock, handler, user)
}

func (_TomatoTraceable *TomatoTraceableSession) WatchAllRoleRegistered(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.Contract.WatchAllRoleRegistered(fromBlock, handler)
}

// ParseRoleRegistered is a log parse operation binding the contract event 0x1e9acfe4df89a29b2be615d44fd962d8ce1fd19ff4a74ddf5a9401a6a4a2a6d2.
//
// Solidity: event RoleRegistered(address indexed user, uint8 role)
func (_TomatoTraceable *TomatoTraceableSession) ParseRoleRegistered(log types.Log) (*TomatoTraceableRoleRegistered, error) {
	return _TomatoTraceable.Contract.ParseRoleRegistered(log)
}

// TomatoTraceableTomatoDataSynced represents a TomatoDataSynced event raised by the TomatoTraceable contract.
type TomatoTraceableTomatoDataSynced struct {
	TomatoId           common.Hash
	ProductionArea     string
	HarvestDate        *big.Int
	ProcessingDate     *big.Int
	TransportInfo      string
	HarvestEvaluation  string
	SamplingEvaluation string
	Raw                types.Log // Blockchain specific contextual infos
}

// WatchTomatoDataSynced is a free log subscription operation binding the contract event 0x9357633e0f3e1beea206951442103e45f2cf632c938287383b1296189b992339.
//
// Solidity: event TomatoDataSynced(string indexed tomatoId, string productionArea, uint256 harvestDate, uint256 processingDate, string transportInfo, string harvestEvaluation, string samplingEvaluation)
func (_TomatoTraceable *TomatoTraceableFilterer) WatchTomatoDataSynced(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "TomatoDataSynced", tomatoId)
}

func (_TomatoTraceable *TomatoTraceableFilterer) WatchAllTomatoDataSynced(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "TomatoDataSynced")
}

// ParseTomatoDataSynced is a log parse operation binding the contract event 0x9357633e0f3e1beea206951442103e45f2cf632c938287383b1296189b992339.
//
// Solidity: event TomatoDataSynced(string indexed tomatoId, string productionArea, uint256 harvestDate, uint256 processingDate, string transportInfo, string harvestEvaluation, string samplingEvaluation)
func (_TomatoTraceable *TomatoTraceableFilterer) ParseTomatoDataSynced(log types.Log) (*TomatoTraceableTomatoDataSynced, error) {
	event := new(TomatoTraceableTomatoDataSynced)
	if err := _TomatoTraceable.contract.UnpackLog(event, "TomatoDataSynced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTomatoDataSynced is a free log subscription operation binding the contract event 0x9357633e0f3e1beea206951442103e45f2cf632c938287383b1296189b992339.
//
// Solidity: event TomatoDataSynced(string indexed tomatoId, string productionArea, uint256 harvestDate, uint256 processingDate, string transportInfo, string harvestEvaluation, string samplingEvaluation)
func (_TomatoTraceable *TomatoTraceableSession) WatchTomatoDataSynced(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.Contract.WatchTomatoDataSynced(fromBlock, handler, tomatoId)
}

func (_TomatoTraceable *TomatoTraceableSession) WatchAllTomatoDataSynced(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.Contract.WatchAllTomatoDataSynced(fromBlock, handler)
}

// ParseTomatoDataSynced is a log parse operation binding the contract event 0x9357633e0f3e1beea206951442103e45f2cf632c938287383b1296189b992339.
//
// Solidity: event TomatoDataSynced(string indexed tomatoId, string productionArea, uint256 harvestDate, uint256 processingDate, string transportInfo, string harvestEvaluation, string samplingEvaluation)
func (_TomatoTraceable *TomatoTraceableSession) ParseTomatoDataSynced(log types.Log) (*TomatoTraceableTomatoDataSynced, error) {
	return _TomatoTraceable.Contract.ParseTomatoDataSynced(log)
}

// TomatoTraceableTomatoInfoAdded represents a TomatoInfoAdded event raised by the TomatoTraceable contract.
type TomatoTraceableTomatoInfoAdded struct {
	TomatoId common.Hash
	Grower   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchTomatoInfoAdded is a free log subscription operation binding the contract event 0x1159a6c621619e4c10470c448123c6bb895f85a6807dadc078838619e17d193d.
//
// Solidity: event TomatoInfoAdded(string indexed tomatoId, address indexed grower)
func (_TomatoTraceable *TomatoTraceableFilterer) WatchTomatoInfoAdded(fromBlock *uint64, handler func(int, []types.Log), tomatoId string, grower common.Address) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "TomatoInfoAdded", tomatoId, grower)
}

func (_TomatoTraceable *TomatoTraceableFilterer) WatchAllTomatoInfoAdded(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "TomatoInfoAdded")
}

// ParseTomatoInfoAdded is a log parse operation binding the contract event 0x1159a6c621619e4c10470c448123c6bb895f85a6807dadc078838619e17d193d.
//
// Solidity: event TomatoInfoAdded(string indexed tomatoId, address indexed grower)
func (_TomatoTraceable *TomatoTraceableFilterer) ParseTomatoInfoAdded(log types.Log) (*TomatoTraceableTomatoInfoAdded, error) {
	event := new(TomatoTraceableTomatoInfoAdded)
	if err := _TomatoTraceable.contract.UnpackLog(event, "TomatoInfoAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTomatoInfoAdded is a free log subscription operation binding the contract event 0x1159a6c621619e4c10470c448123c6bb895f85a6807dadc078838619e17d193d.
//
// Solidity: event TomatoInfoAdded(string indexed tomatoId, address indexed grower)
func (_TomatoTraceable *TomatoTraceableSession) WatchTomatoInfoAdded(fromBlock *uint64, handler func(int, []types.Log), tomatoId string, grower common.Address) (string, error) {
	return _TomatoTraceable.Contract.WatchTomatoInfoAdded(fromBlock, handler, tomatoId, grower)
}

func (_TomatoTraceable *TomatoTraceableSession) WatchAllTomatoInfoAdded(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.Contract.WatchAllTomatoInfoAdded(fromBlock, handler)
}

// ParseTomatoInfoAdded is a log parse operation binding the contract event 0x1159a6c621619e4c10470c448123c6bb895f85a6807dadc078838619e17d193d.
//
// Solidity: event TomatoInfoAdded(string indexed tomatoId, address indexed grower)
func (_TomatoTraceable *TomatoTraceableSession) ParseTomatoInfoAdded(log types.Log) (*TomatoTraceableTomatoInfoAdded, error) {
	return _TomatoTraceable.Contract.ParseTomatoInfoAdded(log)
}

// TomatoTraceableTomatoInfoUpdated represents a TomatoInfoUpdated event raised by the TomatoTraceable contract.
type TomatoTraceableTomatoInfoUpdated struct {
	TomatoId common.Hash
	Field    string
	Value    string
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchTomatoInfoUpdated is a free log subscription operation binding the contract event 0xdf32db9dfa3ab6d1c87edc56dc0755cd0f7a0069e848e5cbe7ec854f5337b4db.
//
// Solidity: event TomatoInfoUpdated(string indexed tomatoId, string field, string value)
func (_TomatoTraceable *TomatoTraceableFilterer) WatchTomatoInfoUpdated(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "TomatoInfoUpdated", tomatoId)
}

func (_TomatoTraceable *TomatoTraceableFilterer) WatchAllTomatoInfoUpdated(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "TomatoInfoUpdated")
}

// ParseTomatoInfoUpdated is a log parse operation binding the contract event 0xdf32db9dfa3ab6d1c87edc56dc0755cd0f7a0069e848e5cbe7ec854f5337b4db.
//
// Solidity: event TomatoInfoUpdated(string indexed tomatoId, string field, string value)
func (_TomatoTraceable *TomatoTraceableFilterer) ParseTomatoInfoUpdated(log types.Log) (*TomatoTraceableTomatoInfoUpdated, error) {
	event := new(TomatoTraceableTomatoInfoUpdated)
	if err := _TomatoTraceable.contract.UnpackLog(event, "TomatoInfoUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTomatoInfoUpdated is a free log subscription operation binding the contract event 0xdf32db9dfa3ab6d1c87edc56dc0755cd0f7a0069e848e5cbe7ec854f5337b4db.
//
// Solidity: event TomatoInfoUpdated(string indexed tomatoId, string field, string value)
func (_TomatoTraceable *TomatoTraceableSession) WatchTomatoInfoUpdated(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.Contract.WatchTomatoInfoUpdated(fromBlock, handler, tomatoId)
}

func (_TomatoTraceable *TomatoTraceableSession) WatchAllTomatoInfoUpdated(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.Contract.WatchAllTomatoInfoUpdated(fromBlock, handler)
}

// ParseTomatoInfoUpdated is a log parse operation binding the contract event 0xdf32db9dfa3ab6d1c87edc56dc0755cd0f7a0069e848e5cbe7ec854f5337b4db.
//
// Solidity: event TomatoInfoUpdated(string indexed tomatoId, string field, string value)
func (_TomatoTraceable *TomatoTraceableSession) ParseTomatoInfoUpdated(log types.Log) (*TomatoTraceableTomatoInfoUpdated, error) {
	return _TomatoTraceable.Contract.ParseTomatoInfoUpdated(log)
}

// TomatoTraceableTransportInfoUpdated represents a TransportInfoUpdated event raised by the TomatoTraceable contract.
type TomatoTraceableTransportInfoUpdated struct {
	TomatoId common.Hash
	Info     string
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchTransportInfoUpdated is a free log subscription operation binding the contract event 0x1477779eea63b2d4b50b5c96f66ce116ad4c9a112f5f4c335fa50b82e0b1bfa7.
//
// Solidity: event TransportInfoUpdated(string indexed tomatoId, string info)
func (_TomatoTraceable *TomatoTraceableFilterer) WatchTransportInfoUpdated(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "TransportInfoUpdated", tomatoId)
}

func (_TomatoTraceable *TomatoTraceableFilterer) WatchAllTransportInfoUpdated(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.contract.WatchLogs(fromBlock, handler, "TransportInfoUpdated")
}

// ParseTransportInfoUpdated is a log parse operation binding the contract event 0x1477779eea63b2d4b50b5c96f66ce116ad4c9a112f5f4c335fa50b82e0b1bfa7.
//
// Solidity: event TransportInfoUpdated(string indexed tomatoId, string info)
func (_TomatoTraceable *TomatoTraceableFilterer) ParseTransportInfoUpdated(log types.Log) (*TomatoTraceableTransportInfoUpdated, error) {
	event := new(TomatoTraceableTransportInfoUpdated)
	if err := _TomatoTraceable.contract.UnpackLog(event, "TransportInfoUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTransportInfoUpdated is a free log subscription operation binding the contract event 0x1477779eea63b2d4b50b5c96f66ce116ad4c9a112f5f4c335fa50b82e0b1bfa7.
//
// Solidity: event TransportInfoUpdated(string indexed tomatoId, string info)
func (_TomatoTraceable *TomatoTraceableSession) WatchTransportInfoUpdated(fromBlock *uint64, handler func(int, []types.Log), tomatoId string) (string, error) {
	return _TomatoTraceable.Contract.WatchTransportInfoUpdated(fromBlock, handler, tomatoId)
}

func (_TomatoTraceable *TomatoTraceableSession) WatchAllTransportInfoUpdated(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _TomatoTraceable.Contract.WatchAllTransportInfoUpdated(fromBlock, handler)
}

// ParseTransportInfoUpdated is a log parse operation binding the contract event 0x1477779eea63b2d4b50b5c96f66ce116ad4c9a112f5f4c335fa50b82e0b1bfa7.
//
// Solidity: event TransportInfoUpdated(string indexed tomatoId, string info)
func (_TomatoTraceable *TomatoTraceableSession) ParseTransportInfoUpdated(log types.Log) (*TomatoTraceableTransportInfoUpdated, error) {
	return _TomatoTraceable.Contract.ParseTransportInfoUpdated(log)
}

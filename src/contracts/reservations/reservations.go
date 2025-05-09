// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reservations

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
	_ = abi.ConvertType
)

// EarthfastNode is an auto generated low-level Go binding around an user-defined struct.
type EarthfastNode struct {
	Id         [32]byte
	OperatorId [32]byte
	Host       string
	Region     string
	Disabled   bool
	Prices     [2]*big.Int
	ProjectIds [2][32]byte
}

// EarthfastSlot is an auto generated low-level Go binding around an user-defined struct.
type EarthfastSlot struct {
	Last bool
	Next bool
}

// ReservationsMetaData contains all meta data concerning the Reservations contract.
var ReservationsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextPrice\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"last\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"next\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEarthfastSlot\",\"name\":\"slot\",\"type\":\"tuple\"}],\"name\":\"ReservationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextPrice\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"last\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"next\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEarthfastSlot\",\"name\":\"slot\",\"type\":\"tuple\"}],\"name\":\"ReservationDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IMPORTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"nodeIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maxPrices\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"last\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"next\",\"type\":\"bool\"}],\"internalType\":\"structEarthfastSlot\",\"name\":\"slot\",\"type\":\"tuple\"}],\"name\":\"createReservations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEarthfastNodes\",\"name\":\"allNodes\",\"type\":\"address\"},{\"internalType\":\"contractEarthfastProjects\",\"name\":\"projects\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"last\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"next\",\"type\":\"bool\"}],\"internalType\":\"structEarthfastSlot\",\"name\":\"slot\",\"type\":\"tuple\"}],\"name\":\"deleteReservationImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"nodeIds\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"last\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"next\",\"type\":\"bool\"}],\"internalType\":\"structEarthfastSlot\",\"name\":\"slot\",\"type\":\"tuple\"}],\"name\":\"deleteReservations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"contractEarthfastRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"}],\"name\":\"getReservationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"skip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getReservations\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"prices\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"projectIds\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structEarthfastNode[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"},{\"internalType\":\"contractEarthfastRegistry\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"grantImporterRole\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"}],\"name\":\"removeProjectNodeIdImpl\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"prices\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"projectIds\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structEarthfastNode[]\",\"name\":\"nodes\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"revokeImporterRole\",\"type\":\"bool\"}],\"name\":\"unsafeImportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEarthfastRegistry\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"unsafeSetRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ReservationsABI is the input ABI used to generate the binding from.
// Deprecated: Use ReservationsMetaData.ABI instead.
var ReservationsABI = ReservationsMetaData.ABI

// Reservations is an auto generated Go binding around an Ethereum contract.
type Reservations struct {
	ReservationsCaller     // Read-only binding to the contract
	ReservationsTransactor // Write-only binding to the contract
	ReservationsFilterer   // Log filterer for contract events
}

// ReservationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReservationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReservationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReservationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReservationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReservationsSession struct {
	Contract     *Reservations     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReservationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReservationsCallerSession struct {
	Contract *ReservationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ReservationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReservationsTransactorSession struct {
	Contract     *ReservationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ReservationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReservationsRaw struct {
	Contract *Reservations // Generic contract binding to access the raw methods on
}

// ReservationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReservationsCallerRaw struct {
	Contract *ReservationsCaller // Generic read-only contract binding to access the raw methods on
}

// ReservationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReservationsTransactorRaw struct {
	Contract *ReservationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReservations creates a new instance of Reservations, bound to a specific deployed contract.
func NewReservations(address common.Address, backend bind.ContractBackend) (*Reservations, error) {
	contract, err := bindReservations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reservations{ReservationsCaller: ReservationsCaller{contract: contract}, ReservationsTransactor: ReservationsTransactor{contract: contract}, ReservationsFilterer: ReservationsFilterer{contract: contract}}, nil
}

// NewReservationsCaller creates a new read-only instance of Reservations, bound to a specific deployed contract.
func NewReservationsCaller(address common.Address, caller bind.ContractCaller) (*ReservationsCaller, error) {
	contract, err := bindReservations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReservationsCaller{contract: contract}, nil
}

// NewReservationsTransactor creates a new write-only instance of Reservations, bound to a specific deployed contract.
func NewReservationsTransactor(address common.Address, transactor bind.ContractTransactor) (*ReservationsTransactor, error) {
	contract, err := bindReservations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReservationsTransactor{contract: contract}, nil
}

// NewReservationsFilterer creates a new log filterer instance of Reservations, bound to a specific deployed contract.
func NewReservationsFilterer(address common.Address, filterer bind.ContractFilterer) (*ReservationsFilterer, error) {
	contract, err := bindReservations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReservationsFilterer{contract: contract}, nil
}

// bindReservations binds a generic wrapper to an already deployed contract.
func bindReservations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReservationsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reservations *ReservationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reservations.Contract.ReservationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reservations *ReservationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reservations.Contract.ReservationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reservations *ReservationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reservations.Contract.ReservationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reservations *ReservationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reservations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reservations *ReservationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reservations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reservations *ReservationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reservations.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Reservations *ReservationsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Reservations.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Reservations *ReservationsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Reservations.Contract.DEFAULTADMINROLE(&_Reservations.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Reservations *ReservationsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Reservations.Contract.DEFAULTADMINROLE(&_Reservations.CallOpts)
}

// IMPORTERROLE is a free data retrieval call binding the contract method 0x6fa2a067.
//
// Solidity: function IMPORTER_ROLE() view returns(bytes32)
func (_Reservations *ReservationsCaller) IMPORTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Reservations.contract.Call(opts, &out, "IMPORTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IMPORTERROLE is a free data retrieval call binding the contract method 0x6fa2a067.
//
// Solidity: function IMPORTER_ROLE() view returns(bytes32)
func (_Reservations *ReservationsSession) IMPORTERROLE() ([32]byte, error) {
	return _Reservations.Contract.IMPORTERROLE(&_Reservations.CallOpts)
}

// IMPORTERROLE is a free data retrieval call binding the contract method 0x6fa2a067.
//
// Solidity: function IMPORTER_ROLE() view returns(bytes32)
func (_Reservations *ReservationsCallerSession) IMPORTERROLE() ([32]byte, error) {
	return _Reservations.Contract.IMPORTERROLE(&_Reservations.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Reservations *ReservationsCaller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reservations.contract.Call(opts, &out, "getRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Reservations *ReservationsSession) GetRegistry() (common.Address, error) {
	return _Reservations.Contract.GetRegistry(&_Reservations.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Reservations *ReservationsCallerSession) GetRegistry() (common.Address, error) {
	return _Reservations.Contract.GetRegistry(&_Reservations.CallOpts)
}

// GetReservationCount is a free data retrieval call binding the contract method 0xe5000f52.
//
// Solidity: function getReservationCount(bytes32 projectId) view returns(uint256)
func (_Reservations *ReservationsCaller) GetReservationCount(opts *bind.CallOpts, projectId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Reservations.contract.Call(opts, &out, "getReservationCount", projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReservationCount is a free data retrieval call binding the contract method 0xe5000f52.
//
// Solidity: function getReservationCount(bytes32 projectId) view returns(uint256)
func (_Reservations *ReservationsSession) GetReservationCount(projectId [32]byte) (*big.Int, error) {
	return _Reservations.Contract.GetReservationCount(&_Reservations.CallOpts, projectId)
}

// GetReservationCount is a free data retrieval call binding the contract method 0xe5000f52.
//
// Solidity: function getReservationCount(bytes32 projectId) view returns(uint256)
func (_Reservations *ReservationsCallerSession) GetReservationCount(projectId [32]byte) (*big.Int, error) {
	return _Reservations.Contract.GetReservationCount(&_Reservations.CallOpts, projectId)
}

// GetReservations is a free data retrieval call binding the contract method 0x1e8698bf.
//
// Solidity: function getReservations(bytes32 projectId, uint256 skip, uint256 size) view returns((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] result)
func (_Reservations *ReservationsCaller) GetReservations(opts *bind.CallOpts, projectId [32]byte, skip *big.Int, size *big.Int) ([]EarthfastNode, error) {
	var out []interface{}
	err := _Reservations.contract.Call(opts, &out, "getReservations", projectId, skip, size)

	if err != nil {
		return *new([]EarthfastNode), err
	}

	out0 := *abi.ConvertType(out[0], new([]EarthfastNode)).(*[]EarthfastNode)

	return out0, err

}

// GetReservations is a free data retrieval call binding the contract method 0x1e8698bf.
//
// Solidity: function getReservations(bytes32 projectId, uint256 skip, uint256 size) view returns((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] result)
func (_Reservations *ReservationsSession) GetReservations(projectId [32]byte, skip *big.Int, size *big.Int) ([]EarthfastNode, error) {
	return _Reservations.Contract.GetReservations(&_Reservations.CallOpts, projectId, skip, size)
}

// GetReservations is a free data retrieval call binding the contract method 0x1e8698bf.
//
// Solidity: function getReservations(bytes32 projectId, uint256 skip, uint256 size) view returns((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] result)
func (_Reservations *ReservationsCallerSession) GetReservations(projectId [32]byte, skip *big.Int, size *big.Int) ([]EarthfastNode, error) {
	return _Reservations.Contract.GetReservations(&_Reservations.CallOpts, projectId, skip, size)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Reservations *ReservationsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Reservations.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Reservations *ReservationsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Reservations.Contract.GetRoleAdmin(&_Reservations.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Reservations *ReservationsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Reservations.Contract.GetRoleAdmin(&_Reservations.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Reservations *ReservationsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Reservations.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Reservations *ReservationsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Reservations.Contract.HasRole(&_Reservations.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Reservations *ReservationsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Reservations.Contract.HasRole(&_Reservations.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Reservations *ReservationsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Reservations.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Reservations *ReservationsSession) Paused() (bool, error) {
	return _Reservations.Contract.Paused(&_Reservations.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Reservations *ReservationsCallerSession) Paused() (bool, error) {
	return _Reservations.Contract.Paused(&_Reservations.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Reservations *ReservationsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Reservations.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Reservations *ReservationsSession) ProxiableUUID() ([32]byte, error) {
	return _Reservations.Contract.ProxiableUUID(&_Reservations.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Reservations *ReservationsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Reservations.Contract.ProxiableUUID(&_Reservations.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Reservations *ReservationsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Reservations.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Reservations *ReservationsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Reservations.Contract.SupportsInterface(&_Reservations.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Reservations *ReservationsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Reservations.Contract.SupportsInterface(&_Reservations.CallOpts, interfaceId)
}

// CreateReservations is a paid mutator transaction binding the contract method 0x2a2d5b2e.
//
// Solidity: function createReservations(bytes32 projectId, bytes32[] nodeIds, uint256[] maxPrices, (bool,bool) slot) returns()
func (_Reservations *ReservationsTransactor) CreateReservations(opts *bind.TransactOpts, projectId [32]byte, nodeIds [][32]byte, maxPrices []*big.Int, slot EarthfastSlot) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "createReservations", projectId, nodeIds, maxPrices, slot)
}

// CreateReservations is a paid mutator transaction binding the contract method 0x2a2d5b2e.
//
// Solidity: function createReservations(bytes32 projectId, bytes32[] nodeIds, uint256[] maxPrices, (bool,bool) slot) returns()
func (_Reservations *ReservationsSession) CreateReservations(projectId [32]byte, nodeIds [][32]byte, maxPrices []*big.Int, slot EarthfastSlot) (*types.Transaction, error) {
	return _Reservations.Contract.CreateReservations(&_Reservations.TransactOpts, projectId, nodeIds, maxPrices, slot)
}

// CreateReservations is a paid mutator transaction binding the contract method 0x2a2d5b2e.
//
// Solidity: function createReservations(bytes32 projectId, bytes32[] nodeIds, uint256[] maxPrices, (bool,bool) slot) returns()
func (_Reservations *ReservationsTransactorSession) CreateReservations(projectId [32]byte, nodeIds [][32]byte, maxPrices []*big.Int, slot EarthfastSlot) (*types.Transaction, error) {
	return _Reservations.Contract.CreateReservations(&_Reservations.TransactOpts, projectId, nodeIds, maxPrices, slot)
}

// DeleteReservationImpl is a paid mutator transaction binding the contract method 0x92cb9a6c.
//
// Solidity: function deleteReservationImpl(address allNodes, address projects, bytes32 projectId, bytes32 nodeId, (bool,bool) slot) returns()
func (_Reservations *ReservationsTransactor) DeleteReservationImpl(opts *bind.TransactOpts, allNodes common.Address, projects common.Address, projectId [32]byte, nodeId [32]byte, slot EarthfastSlot) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "deleteReservationImpl", allNodes, projects, projectId, nodeId, slot)
}

// DeleteReservationImpl is a paid mutator transaction binding the contract method 0x92cb9a6c.
//
// Solidity: function deleteReservationImpl(address allNodes, address projects, bytes32 projectId, bytes32 nodeId, (bool,bool) slot) returns()
func (_Reservations *ReservationsSession) DeleteReservationImpl(allNodes common.Address, projects common.Address, projectId [32]byte, nodeId [32]byte, slot EarthfastSlot) (*types.Transaction, error) {
	return _Reservations.Contract.DeleteReservationImpl(&_Reservations.TransactOpts, allNodes, projects, projectId, nodeId, slot)
}

// DeleteReservationImpl is a paid mutator transaction binding the contract method 0x92cb9a6c.
//
// Solidity: function deleteReservationImpl(address allNodes, address projects, bytes32 projectId, bytes32 nodeId, (bool,bool) slot) returns()
func (_Reservations *ReservationsTransactorSession) DeleteReservationImpl(allNodes common.Address, projects common.Address, projectId [32]byte, nodeId [32]byte, slot EarthfastSlot) (*types.Transaction, error) {
	return _Reservations.Contract.DeleteReservationImpl(&_Reservations.TransactOpts, allNodes, projects, projectId, nodeId, slot)
}

// DeleteReservations is a paid mutator transaction binding the contract method 0xff3d9330.
//
// Solidity: function deleteReservations(bytes32 projectId, bytes32[] nodeIds, (bool,bool) slot) returns()
func (_Reservations *ReservationsTransactor) DeleteReservations(opts *bind.TransactOpts, projectId [32]byte, nodeIds [][32]byte, slot EarthfastSlot) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "deleteReservations", projectId, nodeIds, slot)
}

// DeleteReservations is a paid mutator transaction binding the contract method 0xff3d9330.
//
// Solidity: function deleteReservations(bytes32 projectId, bytes32[] nodeIds, (bool,bool) slot) returns()
func (_Reservations *ReservationsSession) DeleteReservations(projectId [32]byte, nodeIds [][32]byte, slot EarthfastSlot) (*types.Transaction, error) {
	return _Reservations.Contract.DeleteReservations(&_Reservations.TransactOpts, projectId, nodeIds, slot)
}

// DeleteReservations is a paid mutator transaction binding the contract method 0xff3d9330.
//
// Solidity: function deleteReservations(bytes32 projectId, bytes32[] nodeIds, (bool,bool) slot) returns()
func (_Reservations *ReservationsTransactorSession) DeleteReservations(projectId [32]byte, nodeIds [][32]byte, slot EarthfastSlot) (*types.Transaction, error) {
	return _Reservations.Contract.DeleteReservations(&_Reservations.TransactOpts, projectId, nodeIds, slot)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Reservations *ReservationsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Reservations *ReservationsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Reservations.Contract.GrantRole(&_Reservations.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Reservations *ReservationsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Reservations.Contract.GrantRole(&_Reservations.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x7aee1c6b.
//
// Solidity: function initialize(address[] admins, address registry, bool grantImporterRole) returns()
func (_Reservations *ReservationsTransactor) Initialize(opts *bind.TransactOpts, admins []common.Address, registry common.Address, grantImporterRole bool) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "initialize", admins, registry, grantImporterRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x7aee1c6b.
//
// Solidity: function initialize(address[] admins, address registry, bool grantImporterRole) returns()
func (_Reservations *ReservationsSession) Initialize(admins []common.Address, registry common.Address, grantImporterRole bool) (*types.Transaction, error) {
	return _Reservations.Contract.Initialize(&_Reservations.TransactOpts, admins, registry, grantImporterRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x7aee1c6b.
//
// Solidity: function initialize(address[] admins, address registry, bool grantImporterRole) returns()
func (_Reservations *ReservationsTransactorSession) Initialize(admins []common.Address, registry common.Address, grantImporterRole bool) (*types.Transaction, error) {
	return _Reservations.Contract.Initialize(&_Reservations.TransactOpts, admins, registry, grantImporterRole)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Reservations *ReservationsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Reservations *ReservationsSession) Pause() (*types.Transaction, error) {
	return _Reservations.Contract.Pause(&_Reservations.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Reservations *ReservationsTransactorSession) Pause() (*types.Transaction, error) {
	return _Reservations.Contract.Pause(&_Reservations.TransactOpts)
}

// RemoveProjectNodeIdImpl is a paid mutator transaction binding the contract method 0x9cb628be.
//
// Solidity: function removeProjectNodeIdImpl(bytes32 projectId, bytes32 nodeId) returns(bool)
func (_Reservations *ReservationsTransactor) RemoveProjectNodeIdImpl(opts *bind.TransactOpts, projectId [32]byte, nodeId [32]byte) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "removeProjectNodeIdImpl", projectId, nodeId)
}

// RemoveProjectNodeIdImpl is a paid mutator transaction binding the contract method 0x9cb628be.
//
// Solidity: function removeProjectNodeIdImpl(bytes32 projectId, bytes32 nodeId) returns(bool)
func (_Reservations *ReservationsSession) RemoveProjectNodeIdImpl(projectId [32]byte, nodeId [32]byte) (*types.Transaction, error) {
	return _Reservations.Contract.RemoveProjectNodeIdImpl(&_Reservations.TransactOpts, projectId, nodeId)
}

// RemoveProjectNodeIdImpl is a paid mutator transaction binding the contract method 0x9cb628be.
//
// Solidity: function removeProjectNodeIdImpl(bytes32 projectId, bytes32 nodeId) returns(bool)
func (_Reservations *ReservationsTransactorSession) RemoveProjectNodeIdImpl(projectId [32]byte, nodeId [32]byte) (*types.Transaction, error) {
	return _Reservations.Contract.RemoveProjectNodeIdImpl(&_Reservations.TransactOpts, projectId, nodeId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Reservations *ReservationsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Reservations *ReservationsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Reservations.Contract.RenounceRole(&_Reservations.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Reservations *ReservationsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Reservations.Contract.RenounceRole(&_Reservations.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Reservations *ReservationsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Reservations *ReservationsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Reservations.Contract.RevokeRole(&_Reservations.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Reservations *ReservationsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Reservations.Contract.RevokeRole(&_Reservations.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Reservations *ReservationsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Reservations *ReservationsSession) Unpause() (*types.Transaction, error) {
	return _Reservations.Contract.Unpause(&_Reservations.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Reservations *ReservationsTransactorSession) Unpause() (*types.Transaction, error) {
	return _Reservations.Contract.Unpause(&_Reservations.TransactOpts)
}

// UnsafeImportData is a paid mutator transaction binding the contract method 0x8d6fafd7.
//
// Solidity: function unsafeImportData((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] nodes, bool revokeImporterRole) returns()
func (_Reservations *ReservationsTransactor) UnsafeImportData(opts *bind.TransactOpts, nodes []EarthfastNode, revokeImporterRole bool) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "unsafeImportData", nodes, revokeImporterRole)
}

// UnsafeImportData is a paid mutator transaction binding the contract method 0x8d6fafd7.
//
// Solidity: function unsafeImportData((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] nodes, bool revokeImporterRole) returns()
func (_Reservations *ReservationsSession) UnsafeImportData(nodes []EarthfastNode, revokeImporterRole bool) (*types.Transaction, error) {
	return _Reservations.Contract.UnsafeImportData(&_Reservations.TransactOpts, nodes, revokeImporterRole)
}

// UnsafeImportData is a paid mutator transaction binding the contract method 0x8d6fafd7.
//
// Solidity: function unsafeImportData((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] nodes, bool revokeImporterRole) returns()
func (_Reservations *ReservationsTransactorSession) UnsafeImportData(nodes []EarthfastNode, revokeImporterRole bool) (*types.Transaction, error) {
	return _Reservations.Contract.UnsafeImportData(&_Reservations.TransactOpts, nodes, revokeImporterRole)
}

// UnsafeSetRegistry is a paid mutator transaction binding the contract method 0xb9a2adf0.
//
// Solidity: function unsafeSetRegistry(address registry) returns()
func (_Reservations *ReservationsTransactor) UnsafeSetRegistry(opts *bind.TransactOpts, registry common.Address) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "unsafeSetRegistry", registry)
}

// UnsafeSetRegistry is a paid mutator transaction binding the contract method 0xb9a2adf0.
//
// Solidity: function unsafeSetRegistry(address registry) returns()
func (_Reservations *ReservationsSession) UnsafeSetRegistry(registry common.Address) (*types.Transaction, error) {
	return _Reservations.Contract.UnsafeSetRegistry(&_Reservations.TransactOpts, registry)
}

// UnsafeSetRegistry is a paid mutator transaction binding the contract method 0xb9a2adf0.
//
// Solidity: function unsafeSetRegistry(address registry) returns()
func (_Reservations *ReservationsTransactorSession) UnsafeSetRegistry(registry common.Address) (*types.Transaction, error) {
	return _Reservations.Contract.UnsafeSetRegistry(&_Reservations.TransactOpts, registry)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Reservations *ReservationsTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Reservations *ReservationsSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Reservations.Contract.UpgradeTo(&_Reservations.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Reservations *ReservationsTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Reservations.Contract.UpgradeTo(&_Reservations.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Reservations *ReservationsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Reservations.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Reservations *ReservationsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Reservations.Contract.UpgradeToAndCall(&_Reservations.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Reservations *ReservationsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Reservations.Contract.UpgradeToAndCall(&_Reservations.TransactOpts, newImplementation, data)
}

// ReservationsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Reservations contract.
type ReservationsAdminChangedIterator struct {
	Event *ReservationsAdminChanged // Event containing the contract specifics and raw log

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
func (it *ReservationsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsAdminChanged)
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
		it.Event = new(ReservationsAdminChanged)
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
func (it *ReservationsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsAdminChanged represents a AdminChanged event raised by the Reservations contract.
type ReservationsAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Reservations *ReservationsFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ReservationsAdminChangedIterator, error) {

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ReservationsAdminChangedIterator{contract: _Reservations.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Reservations *ReservationsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ReservationsAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsAdminChanged)
				if err := _Reservations.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Reservations *ReservationsFilterer) ParseAdminChanged(log types.Log) (*ReservationsAdminChanged, error) {
	event := new(ReservationsAdminChanged)
	if err := _Reservations.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationsBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Reservations contract.
type ReservationsBeaconUpgradedIterator struct {
	Event *ReservationsBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ReservationsBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsBeaconUpgraded)
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
		it.Event = new(ReservationsBeaconUpgraded)
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
func (it *ReservationsBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsBeaconUpgraded represents a BeaconUpgraded event raised by the Reservations contract.
type ReservationsBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Reservations *ReservationsFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ReservationsBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ReservationsBeaconUpgradedIterator{contract: _Reservations.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Reservations *ReservationsFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ReservationsBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsBeaconUpgraded)
				if err := _Reservations.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Reservations *ReservationsFilterer) ParseBeaconUpgraded(log types.Log) (*ReservationsBeaconUpgraded, error) {
	event := new(ReservationsBeaconUpgraded)
	if err := _Reservations.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Reservations contract.
type ReservationsInitializedIterator struct {
	Event *ReservationsInitialized // Event containing the contract specifics and raw log

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
func (it *ReservationsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsInitialized)
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
		it.Event = new(ReservationsInitialized)
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
func (it *ReservationsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsInitialized represents a Initialized event raised by the Reservations contract.
type ReservationsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Reservations *ReservationsFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReservationsInitializedIterator, error) {

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReservationsInitializedIterator{contract: _Reservations.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Reservations *ReservationsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReservationsInitialized) (event.Subscription, error) {

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsInitialized)
				if err := _Reservations.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Reservations *ReservationsFilterer) ParseInitialized(log types.Log) (*ReservationsInitialized, error) {
	event := new(ReservationsInitialized)
	if err := _Reservations.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Reservations contract.
type ReservationsPausedIterator struct {
	Event *ReservationsPaused // Event containing the contract specifics and raw log

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
func (it *ReservationsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsPaused)
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
		it.Event = new(ReservationsPaused)
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
func (it *ReservationsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsPaused represents a Paused event raised by the Reservations contract.
type ReservationsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Reservations *ReservationsFilterer) FilterPaused(opts *bind.FilterOpts) (*ReservationsPausedIterator, error) {

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ReservationsPausedIterator{contract: _Reservations.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Reservations *ReservationsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ReservationsPaused) (event.Subscription, error) {

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsPaused)
				if err := _Reservations.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Reservations *ReservationsFilterer) ParsePaused(log types.Log) (*ReservationsPaused, error) {
	event := new(ReservationsPaused)
	if err := _Reservations.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationsReservationCreatedIterator is returned from FilterReservationCreated and is used to iterate over the raw logs and unpacked data for ReservationCreated events raised by the Reservations contract.
type ReservationsReservationCreatedIterator struct {
	Event *ReservationsReservationCreated // Event containing the contract specifics and raw log

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
func (it *ReservationsReservationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsReservationCreated)
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
		it.Event = new(ReservationsReservationCreated)
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
func (it *ReservationsReservationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsReservationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsReservationCreated represents a ReservationCreated event raised by the Reservations contract.
type ReservationsReservationCreated struct {
	NodeId     [32]byte
	OperatorId [32]byte
	ProjectId  [32]byte
	LastPrice  *big.Int
	NextPrice  *big.Int
	Slot       EarthfastSlot
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReservationCreated is a free log retrieval operation binding the contract event 0x0deae0b487e00626eb6a45c225a856d96f7d385690db189a5af957e1a33d14fa.
//
// Solidity: event ReservationCreated(bytes32 indexed nodeId, bytes32 indexed operatorId, bytes32 indexed projectId, uint256 lastPrice, uint256 nextPrice, (bool,bool) slot)
func (_Reservations *ReservationsFilterer) FilterReservationCreated(opts *bind.FilterOpts, nodeId [][32]byte, operatorId [][32]byte, projectId [][32]byte) (*ReservationsReservationCreatedIterator, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "ReservationCreated", nodeIdRule, operatorIdRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ReservationsReservationCreatedIterator{contract: _Reservations.contract, event: "ReservationCreated", logs: logs, sub: sub}, nil
}

// WatchReservationCreated is a free log subscription operation binding the contract event 0x0deae0b487e00626eb6a45c225a856d96f7d385690db189a5af957e1a33d14fa.
//
// Solidity: event ReservationCreated(bytes32 indexed nodeId, bytes32 indexed operatorId, bytes32 indexed projectId, uint256 lastPrice, uint256 nextPrice, (bool,bool) slot)
func (_Reservations *ReservationsFilterer) WatchReservationCreated(opts *bind.WatchOpts, sink chan<- *ReservationsReservationCreated, nodeId [][32]byte, operatorId [][32]byte, projectId [][32]byte) (event.Subscription, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "ReservationCreated", nodeIdRule, operatorIdRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsReservationCreated)
				if err := _Reservations.contract.UnpackLog(event, "ReservationCreated", log); err != nil {
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

// ParseReservationCreated is a log parse operation binding the contract event 0x0deae0b487e00626eb6a45c225a856d96f7d385690db189a5af957e1a33d14fa.
//
// Solidity: event ReservationCreated(bytes32 indexed nodeId, bytes32 indexed operatorId, bytes32 indexed projectId, uint256 lastPrice, uint256 nextPrice, (bool,bool) slot)
func (_Reservations *ReservationsFilterer) ParseReservationCreated(log types.Log) (*ReservationsReservationCreated, error) {
	event := new(ReservationsReservationCreated)
	if err := _Reservations.contract.UnpackLog(event, "ReservationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationsReservationDeletedIterator is returned from FilterReservationDeleted and is used to iterate over the raw logs and unpacked data for ReservationDeleted events raised by the Reservations contract.
type ReservationsReservationDeletedIterator struct {
	Event *ReservationsReservationDeleted // Event containing the contract specifics and raw log

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
func (it *ReservationsReservationDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsReservationDeleted)
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
		it.Event = new(ReservationsReservationDeleted)
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
func (it *ReservationsReservationDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsReservationDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsReservationDeleted represents a ReservationDeleted event raised by the Reservations contract.
type ReservationsReservationDeleted struct {
	NodeId     [32]byte
	OperatorId [32]byte
	ProjectId  [32]byte
	LastPrice  *big.Int
	NextPrice  *big.Int
	Slot       EarthfastSlot
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReservationDeleted is a free log retrieval operation binding the contract event 0x289b74863d2b35ba0f436a180b5ef6bb8879ee0b2afd14609b0a066e1da7bf50.
//
// Solidity: event ReservationDeleted(bytes32 indexed nodeId, bytes32 indexed operatorId, bytes32 indexed projectId, uint256 lastPrice, uint256 nextPrice, (bool,bool) slot)
func (_Reservations *ReservationsFilterer) FilterReservationDeleted(opts *bind.FilterOpts, nodeId [][32]byte, operatorId [][32]byte, projectId [][32]byte) (*ReservationsReservationDeletedIterator, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "ReservationDeleted", nodeIdRule, operatorIdRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ReservationsReservationDeletedIterator{contract: _Reservations.contract, event: "ReservationDeleted", logs: logs, sub: sub}, nil
}

// WatchReservationDeleted is a free log subscription operation binding the contract event 0x289b74863d2b35ba0f436a180b5ef6bb8879ee0b2afd14609b0a066e1da7bf50.
//
// Solidity: event ReservationDeleted(bytes32 indexed nodeId, bytes32 indexed operatorId, bytes32 indexed projectId, uint256 lastPrice, uint256 nextPrice, (bool,bool) slot)
func (_Reservations *ReservationsFilterer) WatchReservationDeleted(opts *bind.WatchOpts, sink chan<- *ReservationsReservationDeleted, nodeId [][32]byte, operatorId [][32]byte, projectId [][32]byte) (event.Subscription, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "ReservationDeleted", nodeIdRule, operatorIdRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsReservationDeleted)
				if err := _Reservations.contract.UnpackLog(event, "ReservationDeleted", log); err != nil {
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

// ParseReservationDeleted is a log parse operation binding the contract event 0x289b74863d2b35ba0f436a180b5ef6bb8879ee0b2afd14609b0a066e1da7bf50.
//
// Solidity: event ReservationDeleted(bytes32 indexed nodeId, bytes32 indexed operatorId, bytes32 indexed projectId, uint256 lastPrice, uint256 nextPrice, (bool,bool) slot)
func (_Reservations *ReservationsFilterer) ParseReservationDeleted(log types.Log) (*ReservationsReservationDeleted, error) {
	event := new(ReservationsReservationDeleted)
	if err := _Reservations.contract.UnpackLog(event, "ReservationDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Reservations contract.
type ReservationsRoleAdminChangedIterator struct {
	Event *ReservationsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ReservationsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsRoleAdminChanged)
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
		it.Event = new(ReservationsRoleAdminChanged)
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
func (it *ReservationsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsRoleAdminChanged represents a RoleAdminChanged event raised by the Reservations contract.
type ReservationsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Reservations *ReservationsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ReservationsRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ReservationsRoleAdminChangedIterator{contract: _Reservations.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Reservations *ReservationsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ReservationsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsRoleAdminChanged)
				if err := _Reservations.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Reservations *ReservationsFilterer) ParseRoleAdminChanged(log types.Log) (*ReservationsRoleAdminChanged, error) {
	event := new(ReservationsRoleAdminChanged)
	if err := _Reservations.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Reservations contract.
type ReservationsRoleGrantedIterator struct {
	Event *ReservationsRoleGranted // Event containing the contract specifics and raw log

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
func (it *ReservationsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsRoleGranted)
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
		it.Event = new(ReservationsRoleGranted)
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
func (it *ReservationsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsRoleGranted represents a RoleGranted event raised by the Reservations contract.
type ReservationsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Reservations *ReservationsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ReservationsRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ReservationsRoleGrantedIterator{contract: _Reservations.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Reservations *ReservationsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ReservationsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsRoleGranted)
				if err := _Reservations.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Reservations *ReservationsFilterer) ParseRoleGranted(log types.Log) (*ReservationsRoleGranted, error) {
	event := new(ReservationsRoleGranted)
	if err := _Reservations.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Reservations contract.
type ReservationsRoleRevokedIterator struct {
	Event *ReservationsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ReservationsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsRoleRevoked)
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
		it.Event = new(ReservationsRoleRevoked)
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
func (it *ReservationsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsRoleRevoked represents a RoleRevoked event raised by the Reservations contract.
type ReservationsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Reservations *ReservationsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ReservationsRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ReservationsRoleRevokedIterator{contract: _Reservations.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Reservations *ReservationsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ReservationsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsRoleRevoked)
				if err := _Reservations.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Reservations *ReservationsFilterer) ParseRoleRevoked(log types.Log) (*ReservationsRoleRevoked, error) {
	event := new(ReservationsRoleRevoked)
	if err := _Reservations.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Reservations contract.
type ReservationsUnpausedIterator struct {
	Event *ReservationsUnpaused // Event containing the contract specifics and raw log

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
func (it *ReservationsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsUnpaused)
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
		it.Event = new(ReservationsUnpaused)
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
func (it *ReservationsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsUnpaused represents a Unpaused event raised by the Reservations contract.
type ReservationsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Reservations *ReservationsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ReservationsUnpausedIterator, error) {

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ReservationsUnpausedIterator{contract: _Reservations.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Reservations *ReservationsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ReservationsUnpaused) (event.Subscription, error) {

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsUnpaused)
				if err := _Reservations.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Reservations *ReservationsFilterer) ParseUnpaused(log types.Log) (*ReservationsUnpaused, error) {
	event := new(ReservationsUnpaused)
	if err := _Reservations.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReservationsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Reservations contract.
type ReservationsUpgradedIterator struct {
	Event *ReservationsUpgraded // Event containing the contract specifics and raw log

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
func (it *ReservationsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReservationsUpgraded)
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
		it.Event = new(ReservationsUpgraded)
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
func (it *ReservationsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReservationsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReservationsUpgraded represents a Upgraded event raised by the Reservations contract.
type ReservationsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Reservations *ReservationsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ReservationsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Reservations.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ReservationsUpgradedIterator{contract: _Reservations.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Reservations *ReservationsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ReservationsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Reservations.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReservationsUpgraded)
				if err := _Reservations.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Reservations *ReservationsFilterer) ParseUpgraded(log types.Log) (*ReservationsUpgraded, error) {
	event := new(ReservationsUpgraded)
	if err := _Reservations.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

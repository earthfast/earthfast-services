// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// EarthfastRegistryInitializeData is an auto generated low-level Go binding around an user-defined struct.
type EarthfastRegistryInitializeData struct {
	Version         string
	Nonce           *big.Int
	EpochStart      *big.Int
	LastEpochLength *big.Int
	NextEpochLength *big.Int
	GracePeriod     *big.Int
	Usdc            common.Address
	Token           common.Address
	Billing         common.Address
	Nodes           common.Address
	Operators       common.Address
	Projects        common.Address
	Reservations    common.Address
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochStart\",\"type\":\"uint256\"}],\"name\":\"EpochAdvanced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"advanceEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"contractEarthfastBilling\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCuedEpochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochRemainder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGracePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastEpochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastEpochStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEpochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodes\",\"outputs\":[{\"internalType\":\"contractEarthfastNodes\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"contractEarthfastOperators\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProjects\",\"outputs\":[{\"internalType\":\"contractEarthfastProjects\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservations\",\"outputs\":[{\"internalType\":\"contractEarthfastReservations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractEarthfastToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDC\",\"outputs\":[{\"internalType\":\"contractERC20Permit\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastEpochLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextEpochLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gracePeriod\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Permit\",\"name\":\"usdc\",\"type\":\"address\"},{\"internalType\":\"contractEarthfastToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractEarthfastBilling\",\"name\":\"billing\",\"type\":\"address\"},{\"internalType\":\"contractEarthfastNodes\",\"name\":\"nodes\",\"type\":\"address\"},{\"internalType\":\"contractEarthfastOperators\",\"name\":\"operators\",\"type\":\"address\"},{\"internalType\":\"contractEarthfastProjects\",\"name\":\"projects\",\"type\":\"address\"},{\"internalType\":\"contractEarthfastReservations\",\"name\":\"reservations\",\"type\":\"address\"}],\"internalType\":\"structEarthfastRegistryInitializeData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newNonceImpl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireNotGracePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireNotReconciling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireReconciling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"setCuedEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setGracePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEarthfastBilling\",\"name\":\"billing\",\"type\":\"address\"}],\"name\":\"unsafeSetBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"}],\"name\":\"unsafeSetLastEpochStart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEarthfastNodes\",\"name\":\"nodes\",\"type\":\"address\"}],\"name\":\"unsafeSetNodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEarthfastOperators\",\"name\":\"operators\",\"type\":\"address\"}],\"name\":\"unsafeSetOperators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEarthfastProjects\",\"name\":\"projects\",\"type\":\"address\"}],\"name\":\"unsafeSetProjects\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEarthfastReservations\",\"name\":\"reservations\",\"type\":\"address\"}],\"name\":\"unsafeSetReservations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEarthfastToken\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unsafeSetToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20Permit\",\"name\":\"usdc\",\"type\":\"address\"}],\"name\":\"unsafeSetUSDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unsafeWithdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unsafeWithdrawUSDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Registry.Contract.DEFAULTADMINROLE(&_Registry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Registry.Contract.DEFAULTADMINROLE(&_Registry.CallOpts)
}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(address)
func (_Registry *RegistryCaller) GetBilling(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getBilling")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(address)
func (_Registry *RegistrySession) GetBilling() (common.Address, error) {
	return _Registry.Contract.GetBilling(&_Registry.CallOpts)
}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(address)
func (_Registry *RegistryCallerSession) GetBilling() (common.Address, error) {
	return _Registry.Contract.GetBilling(&_Registry.CallOpts)
}

// GetCuedEpochLength is a free data retrieval call binding the contract method 0x2368d729.
//
// Solidity: function getCuedEpochLength() view returns(uint256)
func (_Registry *RegistryCaller) GetCuedEpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getCuedEpochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCuedEpochLength is a free data retrieval call binding the contract method 0x2368d729.
//
// Solidity: function getCuedEpochLength() view returns(uint256)
func (_Registry *RegistrySession) GetCuedEpochLength() (*big.Int, error) {
	return _Registry.Contract.GetCuedEpochLength(&_Registry.CallOpts)
}

// GetCuedEpochLength is a free data retrieval call binding the contract method 0x2368d729.
//
// Solidity: function getCuedEpochLength() view returns(uint256)
func (_Registry *RegistryCallerSession) GetCuedEpochLength() (*big.Int, error) {
	return _Registry.Contract.GetCuedEpochLength(&_Registry.CallOpts)
}

// GetEpochRemainder is a free data retrieval call binding the contract method 0xc7d45634.
//
// Solidity: function getEpochRemainder() view returns(uint256)
func (_Registry *RegistryCaller) GetEpochRemainder(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getEpochRemainder")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochRemainder is a free data retrieval call binding the contract method 0xc7d45634.
//
// Solidity: function getEpochRemainder() view returns(uint256)
func (_Registry *RegistrySession) GetEpochRemainder() (*big.Int, error) {
	return _Registry.Contract.GetEpochRemainder(&_Registry.CallOpts)
}

// GetEpochRemainder is a free data retrieval call binding the contract method 0xc7d45634.
//
// Solidity: function getEpochRemainder() view returns(uint256)
func (_Registry *RegistryCallerSession) GetEpochRemainder() (*big.Int, error) {
	return _Registry.Contract.GetEpochRemainder(&_Registry.CallOpts)
}

// GetGracePeriod is a free data retrieval call binding the contract method 0xdbd18388.
//
// Solidity: function getGracePeriod() view returns(uint256)
func (_Registry *RegistryCaller) GetGracePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getGracePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGracePeriod is a free data retrieval call binding the contract method 0xdbd18388.
//
// Solidity: function getGracePeriod() view returns(uint256)
func (_Registry *RegistrySession) GetGracePeriod() (*big.Int, error) {
	return _Registry.Contract.GetGracePeriod(&_Registry.CallOpts)
}

// GetGracePeriod is a free data retrieval call binding the contract method 0xdbd18388.
//
// Solidity: function getGracePeriod() view returns(uint256)
func (_Registry *RegistryCallerSession) GetGracePeriod() (*big.Int, error) {
	return _Registry.Contract.GetGracePeriod(&_Registry.CallOpts)
}

// GetLastEpochLength is a free data retrieval call binding the contract method 0x2e51f56b.
//
// Solidity: function getLastEpochLength() view returns(uint256)
func (_Registry *RegistryCaller) GetLastEpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getLastEpochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastEpochLength is a free data retrieval call binding the contract method 0x2e51f56b.
//
// Solidity: function getLastEpochLength() view returns(uint256)
func (_Registry *RegistrySession) GetLastEpochLength() (*big.Int, error) {
	return _Registry.Contract.GetLastEpochLength(&_Registry.CallOpts)
}

// GetLastEpochLength is a free data retrieval call binding the contract method 0x2e51f56b.
//
// Solidity: function getLastEpochLength() view returns(uint256)
func (_Registry *RegistryCallerSession) GetLastEpochLength() (*big.Int, error) {
	return _Registry.Contract.GetLastEpochLength(&_Registry.CallOpts)
}

// GetLastEpochStart is a free data retrieval call binding the contract method 0xf351d5d8.
//
// Solidity: function getLastEpochStart() view returns(uint256)
func (_Registry *RegistryCaller) GetLastEpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getLastEpochStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastEpochStart is a free data retrieval call binding the contract method 0xf351d5d8.
//
// Solidity: function getLastEpochStart() view returns(uint256)
func (_Registry *RegistrySession) GetLastEpochStart() (*big.Int, error) {
	return _Registry.Contract.GetLastEpochStart(&_Registry.CallOpts)
}

// GetLastEpochStart is a free data retrieval call binding the contract method 0xf351d5d8.
//
// Solidity: function getLastEpochStart() view returns(uint256)
func (_Registry *RegistryCallerSession) GetLastEpochStart() (*big.Int, error) {
	return _Registry.Contract.GetLastEpochStart(&_Registry.CallOpts)
}

// GetNextEpochLength is a free data retrieval call binding the contract method 0xcec4191f.
//
// Solidity: function getNextEpochLength() view returns(uint256)
func (_Registry *RegistryCaller) GetNextEpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getNextEpochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextEpochLength is a free data retrieval call binding the contract method 0xcec4191f.
//
// Solidity: function getNextEpochLength() view returns(uint256)
func (_Registry *RegistrySession) GetNextEpochLength() (*big.Int, error) {
	return _Registry.Contract.GetNextEpochLength(&_Registry.CallOpts)
}

// GetNextEpochLength is a free data retrieval call binding the contract method 0xcec4191f.
//
// Solidity: function getNextEpochLength() view returns(uint256)
func (_Registry *RegistryCallerSession) GetNextEpochLength() (*big.Int, error) {
	return _Registry.Contract.GetNextEpochLength(&_Registry.CallOpts)
}

// GetNodes is a free data retrieval call binding the contract method 0xe29581aa.
//
// Solidity: function getNodes() view returns(address)
func (_Registry *RegistryCaller) GetNodes(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getNodes")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodes is a free data retrieval call binding the contract method 0xe29581aa.
//
// Solidity: function getNodes() view returns(address)
func (_Registry *RegistrySession) GetNodes() (common.Address, error) {
	return _Registry.Contract.GetNodes(&_Registry.CallOpts)
}

// GetNodes is a free data retrieval call binding the contract method 0xe29581aa.
//
// Solidity: function getNodes() view returns(address)
func (_Registry *RegistryCallerSession) GetNodes() (common.Address, error) {
	return _Registry.Contract.GetNodes(&_Registry.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Registry *RegistryCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Registry *RegistrySession) GetNonce() (*big.Int, error) {
	return _Registry.Contract.GetNonce(&_Registry.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Registry *RegistryCallerSession) GetNonce() (*big.Int, error) {
	return _Registry.Contract.GetNonce(&_Registry.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address)
func (_Registry *RegistryCaller) GetOperators(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getOperators")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address)
func (_Registry *RegistrySession) GetOperators() (common.Address, error) {
	return _Registry.Contract.GetOperators(&_Registry.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address)
func (_Registry *RegistryCallerSession) GetOperators() (common.Address, error) {
	return _Registry.Contract.GetOperators(&_Registry.CallOpts)
}

// GetProjects is a free data retrieval call binding the contract method 0xdcc60128.
//
// Solidity: function getProjects() view returns(address)
func (_Registry *RegistryCaller) GetProjects(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getProjects")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProjects is a free data retrieval call binding the contract method 0xdcc60128.
//
// Solidity: function getProjects() view returns(address)
func (_Registry *RegistrySession) GetProjects() (common.Address, error) {
	return _Registry.Contract.GetProjects(&_Registry.CallOpts)
}

// GetProjects is a free data retrieval call binding the contract method 0xdcc60128.
//
// Solidity: function getProjects() view returns(address)
func (_Registry *RegistryCallerSession) GetProjects() (common.Address, error) {
	return _Registry.Contract.GetProjects(&_Registry.CallOpts)
}

// GetReservations is a free data retrieval call binding the contract method 0x9c200b88.
//
// Solidity: function getReservations() view returns(address)
func (_Registry *RegistryCaller) GetReservations(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getReservations")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReservations is a free data retrieval call binding the contract method 0x9c200b88.
//
// Solidity: function getReservations() view returns(address)
func (_Registry *RegistrySession) GetReservations() (common.Address, error) {
	return _Registry.Contract.GetReservations(&_Registry.CallOpts)
}

// GetReservations is a free data retrieval call binding the contract method 0x9c200b88.
//
// Solidity: function getReservations() view returns(address)
func (_Registry *RegistryCallerSession) GetReservations() (common.Address, error) {
	return _Registry.Contract.GetReservations(&_Registry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Registry.Contract.GetRoleAdmin(&_Registry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Registry.Contract.GetRoleAdmin(&_Registry.CallOpts, role)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Registry *RegistryCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Registry *RegistrySession) GetToken() (common.Address, error) {
	return _Registry.Contract.GetToken(&_Registry.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Registry *RegistryCallerSession) GetToken() (common.Address, error) {
	return _Registry.Contract.GetToken(&_Registry.CallOpts)
}

// GetUSDC is a free data retrieval call binding the contract method 0x1bf01e9b.
//
// Solidity: function getUSDC() view returns(address)
func (_Registry *RegistryCaller) GetUSDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getUSDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUSDC is a free data retrieval call binding the contract method 0x1bf01e9b.
//
// Solidity: function getUSDC() view returns(address)
func (_Registry *RegistrySession) GetUSDC() (common.Address, error) {
	return _Registry.Contract.GetUSDC(&_Registry.CallOpts)
}

// GetUSDC is a free data retrieval call binding the contract method 0x1bf01e9b.
//
// Solidity: function getUSDC() view returns(address)
func (_Registry *RegistryCallerSession) GetUSDC() (common.Address, error) {
	return _Registry.Contract.GetUSDC(&_Registry.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_Registry *RegistryCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_Registry *RegistrySession) GetVersion() (string, error) {
	return _Registry.Contract.GetVersion(&_Registry.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_Registry *RegistryCallerSession) GetVersion() (string, error) {
	return _Registry.Contract.GetVersion(&_Registry.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registry.Contract.HasRole(&_Registry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registry.Contract.HasRole(&_Registry.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistrySession) Paused() (bool, error) {
	return _Registry.Contract.Paused(&_Registry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistryCallerSession) Paused() (bool, error) {
	return _Registry.Contract.Paused(&_Registry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistrySession) ProxiableUUID() ([32]byte, error) {
	return _Registry.Contract.ProxiableUUID(&_Registry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Registry.Contract.ProxiableUUID(&_Registry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// AdvanceEpoch is a paid mutator transaction binding the contract method 0x3cf80e6c.
//
// Solidity: function advanceEpoch() returns()
func (_Registry *RegistryTransactor) AdvanceEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "advanceEpoch")
}

// AdvanceEpoch is a paid mutator transaction binding the contract method 0x3cf80e6c.
//
// Solidity: function advanceEpoch() returns()
func (_Registry *RegistrySession) AdvanceEpoch() (*types.Transaction, error) {
	return _Registry.Contract.AdvanceEpoch(&_Registry.TransactOpts)
}

// AdvanceEpoch is a paid mutator transaction binding the contract method 0x3cf80e6c.
//
// Solidity: function advanceEpoch() returns()
func (_Registry *RegistryTransactorSession) AdvanceEpoch() (*types.Transaction, error) {
	return _Registry.Contract.AdvanceEpoch(&_Registry.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantRole(&_Registry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantRole(&_Registry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xaf17af0f.
//
// Solidity: function initialize(address[] admins, (string,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address,address,address) data) returns()
func (_Registry *RegistryTransactor) Initialize(opts *bind.TransactOpts, admins []common.Address, data EarthfastRegistryInitializeData) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "initialize", admins, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xaf17af0f.
//
// Solidity: function initialize(address[] admins, (string,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address,address,address) data) returns()
func (_Registry *RegistrySession) Initialize(admins []common.Address, data EarthfastRegistryInitializeData) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, admins, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xaf17af0f.
//
// Solidity: function initialize(address[] admins, (string,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address,address,address) data) returns()
func (_Registry *RegistryTransactorSession) Initialize(admins []common.Address, data EarthfastRegistryInitializeData) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, admins, data)
}

// NewNonceImpl is a paid mutator transaction binding the contract method 0x79a6a3a5.
//
// Solidity: function newNonceImpl() returns(uint256)
func (_Registry *RegistryTransactor) NewNonceImpl(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "newNonceImpl")
}

// NewNonceImpl is a paid mutator transaction binding the contract method 0x79a6a3a5.
//
// Solidity: function newNonceImpl() returns(uint256)
func (_Registry *RegistrySession) NewNonceImpl() (*types.Transaction, error) {
	return _Registry.Contract.NewNonceImpl(&_Registry.TransactOpts)
}

// NewNonceImpl is a paid mutator transaction binding the contract method 0x79a6a3a5.
//
// Solidity: function newNonceImpl() returns(uint256)
func (_Registry *RegistryTransactorSession) NewNonceImpl() (*types.Transaction, error) {
	return _Registry.Contract.NewNonceImpl(&_Registry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistrySession) Pause() (*types.Transaction, error) {
	return _Registry.Contract.Pause(&_Registry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _Registry.Contract.Pause(&_Registry.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RenounceRole(&_Registry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RenounceRole(&_Registry.TransactOpts, role, account)
}

// RequireNotGracePeriod is a paid mutator transaction binding the contract method 0x8c1d0f41.
//
// Solidity: function requireNotGracePeriod() returns()
func (_Registry *RegistryTransactor) RequireNotGracePeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "requireNotGracePeriod")
}

// RequireNotGracePeriod is a paid mutator transaction binding the contract method 0x8c1d0f41.
//
// Solidity: function requireNotGracePeriod() returns()
func (_Registry *RegistrySession) RequireNotGracePeriod() (*types.Transaction, error) {
	return _Registry.Contract.RequireNotGracePeriod(&_Registry.TransactOpts)
}

// RequireNotGracePeriod is a paid mutator transaction binding the contract method 0x8c1d0f41.
//
// Solidity: function requireNotGracePeriod() returns()
func (_Registry *RegistryTransactorSession) RequireNotGracePeriod() (*types.Transaction, error) {
	return _Registry.Contract.RequireNotGracePeriod(&_Registry.TransactOpts)
}

// RequireNotReconciling is a paid mutator transaction binding the contract method 0xbb1f4740.
//
// Solidity: function requireNotReconciling() returns()
func (_Registry *RegistryTransactor) RequireNotReconciling(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "requireNotReconciling")
}

// RequireNotReconciling is a paid mutator transaction binding the contract method 0xbb1f4740.
//
// Solidity: function requireNotReconciling() returns()
func (_Registry *RegistrySession) RequireNotReconciling() (*types.Transaction, error) {
	return _Registry.Contract.RequireNotReconciling(&_Registry.TransactOpts)
}

// RequireNotReconciling is a paid mutator transaction binding the contract method 0xbb1f4740.
//
// Solidity: function requireNotReconciling() returns()
func (_Registry *RegistryTransactorSession) RequireNotReconciling() (*types.Transaction, error) {
	return _Registry.Contract.RequireNotReconciling(&_Registry.TransactOpts)
}

// RequireReconciling is a paid mutator transaction binding the contract method 0x1abdbe93.
//
// Solidity: function requireReconciling() returns()
func (_Registry *RegistryTransactor) RequireReconciling(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "requireReconciling")
}

// RequireReconciling is a paid mutator transaction binding the contract method 0x1abdbe93.
//
// Solidity: function requireReconciling() returns()
func (_Registry *RegistrySession) RequireReconciling() (*types.Transaction, error) {
	return _Registry.Contract.RequireReconciling(&_Registry.TransactOpts)
}

// RequireReconciling is a paid mutator transaction binding the contract method 0x1abdbe93.
//
// Solidity: function requireReconciling() returns()
func (_Registry *RegistryTransactorSession) RequireReconciling() (*types.Transaction, error) {
	return _Registry.Contract.RequireReconciling(&_Registry.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRole(&_Registry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRole(&_Registry.TransactOpts, role, account)
}

// SetCuedEpochLength is a paid mutator transaction binding the contract method 0xdeeca159.
//
// Solidity: function setCuedEpochLength(uint256 length) returns()
func (_Registry *RegistryTransactor) SetCuedEpochLength(opts *bind.TransactOpts, length *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setCuedEpochLength", length)
}

// SetCuedEpochLength is a paid mutator transaction binding the contract method 0xdeeca159.
//
// Solidity: function setCuedEpochLength(uint256 length) returns()
func (_Registry *RegistrySession) SetCuedEpochLength(length *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetCuedEpochLength(&_Registry.TransactOpts, length)
}

// SetCuedEpochLength is a paid mutator transaction binding the contract method 0xdeeca159.
//
// Solidity: function setCuedEpochLength(uint256 length) returns()
func (_Registry *RegistryTransactorSession) SetCuedEpochLength(length *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetCuedEpochLength(&_Registry.TransactOpts, length)
}

// SetGracePeriod is a paid mutator transaction binding the contract method 0xf2f65960.
//
// Solidity: function setGracePeriod(uint256 period) returns()
func (_Registry *RegistryTransactor) SetGracePeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setGracePeriod", period)
}

// SetGracePeriod is a paid mutator transaction binding the contract method 0xf2f65960.
//
// Solidity: function setGracePeriod(uint256 period) returns()
func (_Registry *RegistrySession) SetGracePeriod(period *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetGracePeriod(&_Registry.TransactOpts, period)
}

// SetGracePeriod is a paid mutator transaction binding the contract method 0xf2f65960.
//
// Solidity: function setGracePeriod(uint256 period) returns()
func (_Registry *RegistryTransactorSession) SetGracePeriod(period *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetGracePeriod(&_Registry.TransactOpts, period)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string version) returns()
func (_Registry *RegistryTransactor) SetVersion(opts *bind.TransactOpts, version string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setVersion", version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string version) returns()
func (_Registry *RegistrySession) SetVersion(version string) (*types.Transaction, error) {
	return _Registry.Contract.SetVersion(&_Registry.TransactOpts, version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string version) returns()
func (_Registry *RegistryTransactorSession) SetVersion(version string) (*types.Transaction, error) {
	return _Registry.Contract.SetVersion(&_Registry.TransactOpts, version)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistrySession) Unpause() (*types.Transaction, error) {
	return _Registry.Contract.Unpause(&_Registry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _Registry.Contract.Unpause(&_Registry.TransactOpts)
}

// UnsafeSetBilling is a paid mutator transaction binding the contract method 0x0bdb5cbd.
//
// Solidity: function unsafeSetBilling(address billing) returns()
func (_Registry *RegistryTransactor) UnsafeSetBilling(opts *bind.TransactOpts, billing common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsafeSetBilling", billing)
}

// UnsafeSetBilling is a paid mutator transaction binding the contract method 0x0bdb5cbd.
//
// Solidity: function unsafeSetBilling(address billing) returns()
func (_Registry *RegistrySession) UnsafeSetBilling(billing common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetBilling(&_Registry.TransactOpts, billing)
}

// UnsafeSetBilling is a paid mutator transaction binding the contract method 0x0bdb5cbd.
//
// Solidity: function unsafeSetBilling(address billing) returns()
func (_Registry *RegistryTransactorSession) UnsafeSetBilling(billing common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetBilling(&_Registry.TransactOpts, billing)
}

// UnsafeSetLastEpochStart is a paid mutator transaction binding the contract method 0x9bcf5592.
//
// Solidity: function unsafeSetLastEpochStart(uint256 start) returns()
func (_Registry *RegistryTransactor) UnsafeSetLastEpochStart(opts *bind.TransactOpts, start *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsafeSetLastEpochStart", start)
}

// UnsafeSetLastEpochStart is a paid mutator transaction binding the contract method 0x9bcf5592.
//
// Solidity: function unsafeSetLastEpochStart(uint256 start) returns()
func (_Registry *RegistrySession) UnsafeSetLastEpochStart(start *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetLastEpochStart(&_Registry.TransactOpts, start)
}

// UnsafeSetLastEpochStart is a paid mutator transaction binding the contract method 0x9bcf5592.
//
// Solidity: function unsafeSetLastEpochStart(uint256 start) returns()
func (_Registry *RegistryTransactorSession) UnsafeSetLastEpochStart(start *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetLastEpochStart(&_Registry.TransactOpts, start)
}

// UnsafeSetNodes is a paid mutator transaction binding the contract method 0xafc612ce.
//
// Solidity: function unsafeSetNodes(address nodes) returns()
func (_Registry *RegistryTransactor) UnsafeSetNodes(opts *bind.TransactOpts, nodes common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsafeSetNodes", nodes)
}

// UnsafeSetNodes is a paid mutator transaction binding the contract method 0xafc612ce.
//
// Solidity: function unsafeSetNodes(address nodes) returns()
func (_Registry *RegistrySession) UnsafeSetNodes(nodes common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetNodes(&_Registry.TransactOpts, nodes)
}

// UnsafeSetNodes is a paid mutator transaction binding the contract method 0xafc612ce.
//
// Solidity: function unsafeSetNodes(address nodes) returns()
func (_Registry *RegistryTransactorSession) UnsafeSetNodes(nodes common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetNodes(&_Registry.TransactOpts, nodes)
}

// UnsafeSetOperators is a paid mutator transaction binding the contract method 0xd9cbbceb.
//
// Solidity: function unsafeSetOperators(address operators) returns()
func (_Registry *RegistryTransactor) UnsafeSetOperators(opts *bind.TransactOpts, operators common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsafeSetOperators", operators)
}

// UnsafeSetOperators is a paid mutator transaction binding the contract method 0xd9cbbceb.
//
// Solidity: function unsafeSetOperators(address operators) returns()
func (_Registry *RegistrySession) UnsafeSetOperators(operators common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetOperators(&_Registry.TransactOpts, operators)
}

// UnsafeSetOperators is a paid mutator transaction binding the contract method 0xd9cbbceb.
//
// Solidity: function unsafeSetOperators(address operators) returns()
func (_Registry *RegistryTransactorSession) UnsafeSetOperators(operators common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetOperators(&_Registry.TransactOpts, operators)
}

// UnsafeSetProjects is a paid mutator transaction binding the contract method 0x813f5a61.
//
// Solidity: function unsafeSetProjects(address projects) returns()
func (_Registry *RegistryTransactor) UnsafeSetProjects(opts *bind.TransactOpts, projects common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsafeSetProjects", projects)
}

// UnsafeSetProjects is a paid mutator transaction binding the contract method 0x813f5a61.
//
// Solidity: function unsafeSetProjects(address projects) returns()
func (_Registry *RegistrySession) UnsafeSetProjects(projects common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetProjects(&_Registry.TransactOpts, projects)
}

// UnsafeSetProjects is a paid mutator transaction binding the contract method 0x813f5a61.
//
// Solidity: function unsafeSetProjects(address projects) returns()
func (_Registry *RegistryTransactorSession) UnsafeSetProjects(projects common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetProjects(&_Registry.TransactOpts, projects)
}

// UnsafeSetReservations is a paid mutator transaction binding the contract method 0x5fb466b0.
//
// Solidity: function unsafeSetReservations(address reservations) returns()
func (_Registry *RegistryTransactor) UnsafeSetReservations(opts *bind.TransactOpts, reservations common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsafeSetReservations", reservations)
}

// UnsafeSetReservations is a paid mutator transaction binding the contract method 0x5fb466b0.
//
// Solidity: function unsafeSetReservations(address reservations) returns()
func (_Registry *RegistrySession) UnsafeSetReservations(reservations common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetReservations(&_Registry.TransactOpts, reservations)
}

// UnsafeSetReservations is a paid mutator transaction binding the contract method 0x5fb466b0.
//
// Solidity: function unsafeSetReservations(address reservations) returns()
func (_Registry *RegistryTransactorSession) UnsafeSetReservations(reservations common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetReservations(&_Registry.TransactOpts, reservations)
}

// UnsafeSetToken is a paid mutator transaction binding the contract method 0xe72b7cd8.
//
// Solidity: function unsafeSetToken(address token) returns()
func (_Registry *RegistryTransactor) UnsafeSetToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsafeSetToken", token)
}

// UnsafeSetToken is a paid mutator transaction binding the contract method 0xe72b7cd8.
//
// Solidity: function unsafeSetToken(address token) returns()
func (_Registry *RegistrySession) UnsafeSetToken(token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetToken(&_Registry.TransactOpts, token)
}

// UnsafeSetToken is a paid mutator transaction binding the contract method 0xe72b7cd8.
//
// Solidity: function unsafeSetToken(address token) returns()
func (_Registry *RegistryTransactorSession) UnsafeSetToken(token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetToken(&_Registry.TransactOpts, token)
}

// UnsafeSetUSDC is a paid mutator transaction binding the contract method 0xbd06eb40.
//
// Solidity: function unsafeSetUSDC(address usdc) returns()
func (_Registry *RegistryTransactor) UnsafeSetUSDC(opts *bind.TransactOpts, usdc common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsafeSetUSDC", usdc)
}

// UnsafeSetUSDC is a paid mutator transaction binding the contract method 0xbd06eb40.
//
// Solidity: function unsafeSetUSDC(address usdc) returns()
func (_Registry *RegistrySession) UnsafeSetUSDC(usdc common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetUSDC(&_Registry.TransactOpts, usdc)
}

// UnsafeSetUSDC is a paid mutator transaction binding the contract method 0xbd06eb40.
//
// Solidity: function unsafeSetUSDC(address usdc) returns()
func (_Registry *RegistryTransactorSession) UnsafeSetUSDC(usdc common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeSetUSDC(&_Registry.TransactOpts, usdc)
}

// UnsafeWithdrawToken is a paid mutator transaction binding the contract method 0x7cdf73ca.
//
// Solidity: function unsafeWithdrawToken(address to, uint256 amount) returns()
func (_Registry *RegistryTransactor) UnsafeWithdrawToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsafeWithdrawToken", to, amount)
}

// UnsafeWithdrawToken is a paid mutator transaction binding the contract method 0x7cdf73ca.
//
// Solidity: function unsafeWithdrawToken(address to, uint256 amount) returns()
func (_Registry *RegistrySession) UnsafeWithdrawToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeWithdrawToken(&_Registry.TransactOpts, to, amount)
}

// UnsafeWithdrawToken is a paid mutator transaction binding the contract method 0x7cdf73ca.
//
// Solidity: function unsafeWithdrawToken(address to, uint256 amount) returns()
func (_Registry *RegistryTransactorSession) UnsafeWithdrawToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeWithdrawToken(&_Registry.TransactOpts, to, amount)
}

// UnsafeWithdrawUSDC is a paid mutator transaction binding the contract method 0xdfb5db0f.
//
// Solidity: function unsafeWithdrawUSDC(address to, uint256 amount) returns()
func (_Registry *RegistryTransactor) UnsafeWithdrawUSDC(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsafeWithdrawUSDC", to, amount)
}

// UnsafeWithdrawUSDC is a paid mutator transaction binding the contract method 0xdfb5db0f.
//
// Solidity: function unsafeWithdrawUSDC(address to, uint256 amount) returns()
func (_Registry *RegistrySession) UnsafeWithdrawUSDC(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeWithdrawUSDC(&_Registry.TransactOpts, to, amount)
}

// UnsafeWithdrawUSDC is a paid mutator transaction binding the contract method 0xdfb5db0f.
//
// Solidity: function unsafeWithdrawUSDC(address to, uint256 amount) returns()
func (_Registry *RegistryTransactorSession) UnsafeWithdrawUSDC(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UnsafeWithdrawUSDC(&_Registry.TransactOpts, to, amount)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Registry *RegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Registry *RegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpgradeTo(&_Registry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Registry *RegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpgradeTo(&_Registry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpgradeToAndCall(&_Registry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpgradeToAndCall(&_Registry.TransactOpts, newImplementation, data)
}

// RegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Registry contract.
type RegistryAdminChangedIterator struct {
	Event *RegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *RegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryAdminChanged)
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
		it.Event = new(RegistryAdminChanged)
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
func (it *RegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryAdminChanged represents a AdminChanged event raised by the Registry contract.
type RegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Registry *RegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RegistryAdminChangedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RegistryAdminChangedIterator{contract: _Registry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Registry *RegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryAdminChanged)
				if err := _Registry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseAdminChanged(log types.Log) (*RegistryAdminChanged, error) {
	event := new(RegistryAdminChanged)
	if err := _Registry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Registry contract.
type RegistryBeaconUpgradedIterator struct {
	Event *RegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *RegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBeaconUpgraded)
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
		it.Event = new(RegistryBeaconUpgraded)
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
func (it *RegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBeaconUpgraded represents a BeaconUpgraded event raised by the Registry contract.
type RegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Registry *RegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*RegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &RegistryBeaconUpgradedIterator{contract: _Registry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Registry *RegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *RegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBeaconUpgraded)
				if err := _Registry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseBeaconUpgraded(log types.Log) (*RegistryBeaconUpgraded, error) {
	event := new(RegistryBeaconUpgraded)
	if err := _Registry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryEpochAdvancedIterator is returned from FilterEpochAdvanced and is used to iterate over the raw logs and unpacked data for EpochAdvanced events raised by the Registry contract.
type RegistryEpochAdvancedIterator struct {
	Event *RegistryEpochAdvanced // Event containing the contract specifics and raw log

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
func (it *RegistryEpochAdvancedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryEpochAdvanced)
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
		it.Event = new(RegistryEpochAdvanced)
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
func (it *RegistryEpochAdvancedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryEpochAdvancedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryEpochAdvanced represents a EpochAdvanced event raised by the Registry contract.
type RegistryEpochAdvanced struct {
	EpochStart *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEpochAdvanced is a free log retrieval operation binding the contract event 0xb7f0e37f8e78f28f1f21b98f426d307e949c22f2997ca4b562ee90c3e5b30b30.
//
// Solidity: event EpochAdvanced(uint256 epochStart)
func (_Registry *RegistryFilterer) FilterEpochAdvanced(opts *bind.FilterOpts) (*RegistryEpochAdvancedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "EpochAdvanced")
	if err != nil {
		return nil, err
	}
	return &RegistryEpochAdvancedIterator{contract: _Registry.contract, event: "EpochAdvanced", logs: logs, sub: sub}, nil
}

// WatchEpochAdvanced is a free log subscription operation binding the contract event 0xb7f0e37f8e78f28f1f21b98f426d307e949c22f2997ca4b562ee90c3e5b30b30.
//
// Solidity: event EpochAdvanced(uint256 epochStart)
func (_Registry *RegistryFilterer) WatchEpochAdvanced(opts *bind.WatchOpts, sink chan<- *RegistryEpochAdvanced) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "EpochAdvanced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryEpochAdvanced)
				if err := _Registry.contract.UnpackLog(event, "EpochAdvanced", log); err != nil {
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

// ParseEpochAdvanced is a log parse operation binding the contract event 0xb7f0e37f8e78f28f1f21b98f426d307e949c22f2997ca4b562ee90c3e5b30b30.
//
// Solidity: event EpochAdvanced(uint256 epochStart)
func (_Registry *RegistryFilterer) ParseEpochAdvanced(log types.Log) (*RegistryEpochAdvanced, error) {
	event := new(RegistryEpochAdvanced)
	if err := _Registry.contract.UnpackLog(event, "EpochAdvanced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Registry contract.
type RegistryInitializedIterator struct {
	Event *RegistryInitialized // Event containing the contract specifics and raw log

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
func (it *RegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryInitialized)
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
		it.Event = new(RegistryInitialized)
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
func (it *RegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryInitialized represents a Initialized event raised by the Registry contract.
type RegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registry *RegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistryInitializedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistryInitializedIterator{contract: _Registry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registry *RegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryInitialized)
				if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseInitialized(log types.Log) (*RegistryInitialized, error) {
	event := new(RegistryInitialized)
	if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Registry contract.
type RegistryPausedIterator struct {
	Event *RegistryPaused // Event containing the contract specifics and raw log

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
func (it *RegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPaused)
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
		it.Event = new(RegistryPaused)
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
func (it *RegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPaused represents a Paused event raised by the Registry contract.
type RegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*RegistryPausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RegistryPausedIterator{contract: _Registry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RegistryPaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPaused)
				if err := _Registry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Registry *RegistryFilterer) ParsePaused(log types.Log) (*RegistryPaused, error) {
	event := new(RegistryPaused)
	if err := _Registry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Registry contract.
type RegistryRoleAdminChangedIterator struct {
	Event *RegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleAdminChanged)
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
		it.Event = new(RegistryRoleAdminChanged)
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
func (it *RegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleAdminChanged represents a RoleAdminChanged event raised by the Registry contract.
type RegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registry *RegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleAdminChangedIterator{contract: _Registry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registry *RegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleAdminChanged)
				if err := _Registry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseRoleAdminChanged(log types.Log) (*RegistryRoleAdminChanged, error) {
	event := new(RegistryRoleAdminChanged)
	if err := _Registry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Registry contract.
type RegistryRoleGrantedIterator struct {
	Event *RegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *RegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleGranted)
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
		it.Event = new(RegistryRoleGranted)
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
func (it *RegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleGranted represents a RoleGranted event raised by the Registry contract.
type RegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleGrantedIterator{contract: _Registry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleGranted)
				if err := _Registry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseRoleGranted(log types.Log) (*RegistryRoleGranted, error) {
	event := new(RegistryRoleGranted)
	if err := _Registry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Registry contract.
type RegistryRoleRevokedIterator struct {
	Event *RegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleRevoked)
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
		it.Event = new(RegistryRoleRevoked)
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
func (it *RegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleRevoked represents a RoleRevoked event raised by the Registry contract.
type RegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleRevokedIterator{contract: _Registry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleRevoked)
				if err := _Registry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseRoleRevoked(log types.Log) (*RegistryRoleRevoked, error) {
	event := new(RegistryRoleRevoked)
	if err := _Registry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Registry contract.
type RegistryUnpausedIterator struct {
	Event *RegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *RegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUnpaused)
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
		it.Event = new(RegistryUnpaused)
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
func (it *RegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUnpaused represents a Unpaused event raised by the Registry contract.
type RegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RegistryUnpausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RegistryUnpausedIterator{contract: _Registry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUnpaused)
				if err := _Registry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseUnpaused(log types.Log) (*RegistryUnpaused, error) {
	event := new(RegistryUnpaused)
	if err := _Registry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Registry contract.
type RegistryUpgradedIterator struct {
	Event *RegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *RegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUpgraded)
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
		it.Event = new(RegistryUpgraded)
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
func (it *RegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUpgraded represents a Upgraded event raised by the Registry contract.
type RegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Registry *RegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RegistryUpgradedIterator{contract: _Registry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Registry *RegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUpgraded)
				if err := _Registry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseUpgraded(log types.Log) (*RegistryUpgraded, error) {
	event := new(RegistryUpgraded)
	if err := _Registry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

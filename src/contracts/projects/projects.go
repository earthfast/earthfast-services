// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package projects

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

// EarthfastCreateProjectData is an auto generated low-level Go binding around an user-defined struct.
type EarthfastCreateProjectData struct {
	Owner    common.Address
	Name     string
	Email    string
	Content  string
	Checksum [32]byte
	Metadata string
}

// EarthfastProject is an auto generated low-level Go binding around an user-defined struct.
type EarthfastProject struct {
	Id       [32]byte
	Owner    common.Address
	Name     string
	Email    string
	Escrow   *big.Int
	Reserve  *big.Int
	Content  string
	Checksum [32]byte
	Metadata string
}

// ProjectsMetaData contains all meta data concerning the Projects contract.
var ProjectsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldContent\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldChecksum\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newContent\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newChecksum\",\"type\":\"bytes32\"}],\"name\":\"ProjectContentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"checksum\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"ProjectCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"checksum\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"ProjectDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldEscrow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEscrow\",\"type\":\"uint256\"}],\"name\":\"ProjectEscrowChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldMetadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMetadata\",\"type\":\"string\"}],\"name\":\"ProjectMetadataChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"ProjectOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldEmail\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newEmail\",\"type\":\"string\"}],\"name\":\"ProjectPropsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IMPORTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"checksum\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structEarthfastCreateProjectData\",\"name\":\"project\",\"type\":\"tuple\"}],\"name\":\"createProject\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"}],\"name\":\"deleteProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"depositProjectEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"}],\"name\":\"getProject\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"escrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"checksum\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structEarthfastProject\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProjectCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"skip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getProjects\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"escrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"checksum\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structEarthfastProject[]\",\"name\":\"values\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"contractEarthfastRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"},{\"internalType\":\"contractEarthfastRegistry\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"grantImporterRole\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"checksum\",\"type\":\"bytes32\"}],\"name\":\"setProjectContent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"decrease\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"increase\",\"type\":\"uint256\"}],\"name\":\"setProjectEscrowImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"setProjectMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setProjectOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"setProjectProps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"decrease\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"increase\",\"type\":\"uint256\"}],\"name\":\"setProjectReserveImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"escrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"checksum\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structEarthfastProject[]\",\"name\":\"projects\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"revokeImporterRole\",\"type\":\"bool\"}],\"name\":\"unsafeImportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"skip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mul\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"div\",\"type\":\"uint256\"}],\"name\":\"unsafeSetEscrows\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEarthfastRegistry\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"unsafeSetRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawProjectEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProjectsABI is the input ABI used to generate the binding from.
// Deprecated: Use ProjectsMetaData.ABI instead.
var ProjectsABI = ProjectsMetaData.ABI

// Projects is an auto generated Go binding around an Ethereum contract.
type Projects struct {
	ProjectsCaller     // Read-only binding to the contract
	ProjectsTransactor // Write-only binding to the contract
	ProjectsFilterer   // Log filterer for contract events
}

// ProjectsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectsSession struct {
	Contract     *Projects         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProjectsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectsCallerSession struct {
	Contract *ProjectsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProjectsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectsTransactorSession struct {
	Contract     *ProjectsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProjectsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectsRaw struct {
	Contract *Projects // Generic contract binding to access the raw methods on
}

// ProjectsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectsCallerRaw struct {
	Contract *ProjectsCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectsTransactorRaw struct {
	Contract *ProjectsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProjects creates a new instance of Projects, bound to a specific deployed contract.
func NewProjects(address common.Address, backend bind.ContractBackend) (*Projects, error) {
	contract, err := bindProjects(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Projects{ProjectsCaller: ProjectsCaller{contract: contract}, ProjectsTransactor: ProjectsTransactor{contract: contract}, ProjectsFilterer: ProjectsFilterer{contract: contract}}, nil
}

// NewProjectsCaller creates a new read-only instance of Projects, bound to a specific deployed contract.
func NewProjectsCaller(address common.Address, caller bind.ContractCaller) (*ProjectsCaller, error) {
	contract, err := bindProjects(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectsCaller{contract: contract}, nil
}

// NewProjectsTransactor creates a new write-only instance of Projects, bound to a specific deployed contract.
func NewProjectsTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectsTransactor, error) {
	contract, err := bindProjects(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectsTransactor{contract: contract}, nil
}

// NewProjectsFilterer creates a new log filterer instance of Projects, bound to a specific deployed contract.
func NewProjectsFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectsFilterer, error) {
	contract, err := bindProjects(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectsFilterer{contract: contract}, nil
}

// bindProjects binds a generic wrapper to an already deployed contract.
func bindProjects(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProjectsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Projects *ProjectsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Projects.Contract.ProjectsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Projects *ProjectsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Projects.Contract.ProjectsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Projects *ProjectsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Projects.Contract.ProjectsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Projects *ProjectsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Projects.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Projects *ProjectsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Projects.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Projects *ProjectsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Projects.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Projects *ProjectsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Projects *ProjectsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Projects.Contract.DEFAULTADMINROLE(&_Projects.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Projects *ProjectsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Projects.Contract.DEFAULTADMINROLE(&_Projects.CallOpts)
}

// IMPORTERROLE is a free data retrieval call binding the contract method 0x6fa2a067.
//
// Solidity: function IMPORTER_ROLE() view returns(bytes32)
func (_Projects *ProjectsCaller) IMPORTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "IMPORTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IMPORTERROLE is a free data retrieval call binding the contract method 0x6fa2a067.
//
// Solidity: function IMPORTER_ROLE() view returns(bytes32)
func (_Projects *ProjectsSession) IMPORTERROLE() ([32]byte, error) {
	return _Projects.Contract.IMPORTERROLE(&_Projects.CallOpts)
}

// IMPORTERROLE is a free data retrieval call binding the contract method 0x6fa2a067.
//
// Solidity: function IMPORTER_ROLE() view returns(bytes32)
func (_Projects *ProjectsCallerSession) IMPORTERROLE() ([32]byte, error) {
	return _Projects.Contract.IMPORTERROLE(&_Projects.CallOpts)
}

// GetProject is a free data retrieval call binding the contract method 0x4b5f748a.
//
// Solidity: function getProject(bytes32 projectId) view returns((bytes32,address,string,string,uint256,uint256,string,bytes32,string))
func (_Projects *ProjectsCaller) GetProject(opts *bind.CallOpts, projectId [32]byte) (EarthfastProject, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "getProject", projectId)

	if err != nil {
		return *new(EarthfastProject), err
	}

	out0 := *abi.ConvertType(out[0], new(EarthfastProject)).(*EarthfastProject)

	return out0, err

}

// GetProject is a free data retrieval call binding the contract method 0x4b5f748a.
//
// Solidity: function getProject(bytes32 projectId) view returns((bytes32,address,string,string,uint256,uint256,string,bytes32,string))
func (_Projects *ProjectsSession) GetProject(projectId [32]byte) (EarthfastProject, error) {
	return _Projects.Contract.GetProject(&_Projects.CallOpts, projectId)
}

// GetProject is a free data retrieval call binding the contract method 0x4b5f748a.
//
// Solidity: function getProject(bytes32 projectId) view returns((bytes32,address,string,string,uint256,uint256,string,bytes32,string))
func (_Projects *ProjectsCallerSession) GetProject(projectId [32]byte) (EarthfastProject, error) {
	return _Projects.Contract.GetProject(&_Projects.CallOpts, projectId)
}

// GetProjectCount is a free data retrieval call binding the contract method 0x3bcff3b0.
//
// Solidity: function getProjectCount() view returns(uint256)
func (_Projects *ProjectsCaller) GetProjectCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "getProjectCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProjectCount is a free data retrieval call binding the contract method 0x3bcff3b0.
//
// Solidity: function getProjectCount() view returns(uint256)
func (_Projects *ProjectsSession) GetProjectCount() (*big.Int, error) {
	return _Projects.Contract.GetProjectCount(&_Projects.CallOpts)
}

// GetProjectCount is a free data retrieval call binding the contract method 0x3bcff3b0.
//
// Solidity: function getProjectCount() view returns(uint256)
func (_Projects *ProjectsCallerSession) GetProjectCount() (*big.Int, error) {
	return _Projects.Contract.GetProjectCount(&_Projects.CallOpts)
}

// GetProjects is a free data retrieval call binding the contract method 0xa84ce2b5.
//
// Solidity: function getProjects(uint256 skip, uint256 size) view returns((bytes32,address,string,string,uint256,uint256,string,bytes32,string)[] values)
func (_Projects *ProjectsCaller) GetProjects(opts *bind.CallOpts, skip *big.Int, size *big.Int) ([]EarthfastProject, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "getProjects", skip, size)

	if err != nil {
		return *new([]EarthfastProject), err
	}

	out0 := *abi.ConvertType(out[0], new([]EarthfastProject)).(*[]EarthfastProject)

	return out0, err

}

// GetProjects is a free data retrieval call binding the contract method 0xa84ce2b5.
//
// Solidity: function getProjects(uint256 skip, uint256 size) view returns((bytes32,address,string,string,uint256,uint256,string,bytes32,string)[] values)
func (_Projects *ProjectsSession) GetProjects(skip *big.Int, size *big.Int) ([]EarthfastProject, error) {
	return _Projects.Contract.GetProjects(&_Projects.CallOpts, skip, size)
}

// GetProjects is a free data retrieval call binding the contract method 0xa84ce2b5.
//
// Solidity: function getProjects(uint256 skip, uint256 size) view returns((bytes32,address,string,string,uint256,uint256,string,bytes32,string)[] values)
func (_Projects *ProjectsCallerSession) GetProjects(skip *big.Int, size *big.Int) ([]EarthfastProject, error) {
	return _Projects.Contract.GetProjects(&_Projects.CallOpts, skip, size)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Projects *ProjectsCaller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "getRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Projects *ProjectsSession) GetRegistry() (common.Address, error) {
	return _Projects.Contract.GetRegistry(&_Projects.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Projects *ProjectsCallerSession) GetRegistry() (common.Address, error) {
	return _Projects.Contract.GetRegistry(&_Projects.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Projects *ProjectsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Projects *ProjectsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Projects.Contract.GetRoleAdmin(&_Projects.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Projects *ProjectsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Projects.Contract.GetRoleAdmin(&_Projects.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Projects *ProjectsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Projects *ProjectsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Projects.Contract.HasRole(&_Projects.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Projects *ProjectsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Projects.Contract.HasRole(&_Projects.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Projects *ProjectsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Projects *ProjectsSession) Paused() (bool, error) {
	return _Projects.Contract.Paused(&_Projects.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Projects *ProjectsCallerSession) Paused() (bool, error) {
	return _Projects.Contract.Paused(&_Projects.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Projects *ProjectsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Projects *ProjectsSession) ProxiableUUID() ([32]byte, error) {
	return _Projects.Contract.ProxiableUUID(&_Projects.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Projects *ProjectsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Projects.Contract.ProxiableUUID(&_Projects.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Projects *ProjectsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Projects.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Projects *ProjectsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Projects.Contract.SupportsInterface(&_Projects.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Projects *ProjectsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Projects.Contract.SupportsInterface(&_Projects.CallOpts, interfaceId)
}

// CreateProject is a paid mutator transaction binding the contract method 0x4db87362.
//
// Solidity: function createProject((address,string,string,string,bytes32,string) project) returns(bytes32 projectId)
func (_Projects *ProjectsTransactor) CreateProject(opts *bind.TransactOpts, project EarthfastCreateProjectData) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "createProject", project)
}

// CreateProject is a paid mutator transaction binding the contract method 0x4db87362.
//
// Solidity: function createProject((address,string,string,string,bytes32,string) project) returns(bytes32 projectId)
func (_Projects *ProjectsSession) CreateProject(project EarthfastCreateProjectData) (*types.Transaction, error) {
	return _Projects.Contract.CreateProject(&_Projects.TransactOpts, project)
}

// CreateProject is a paid mutator transaction binding the contract method 0x4db87362.
//
// Solidity: function createProject((address,string,string,string,bytes32,string) project) returns(bytes32 projectId)
func (_Projects *ProjectsTransactorSession) CreateProject(project EarthfastCreateProjectData) (*types.Transaction, error) {
	return _Projects.Contract.CreateProject(&_Projects.TransactOpts, project)
}

// DeleteProject is a paid mutator transaction binding the contract method 0xd0a18958.
//
// Solidity: function deleteProject(bytes32 projectId) returns()
func (_Projects *ProjectsTransactor) DeleteProject(opts *bind.TransactOpts, projectId [32]byte) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "deleteProject", projectId)
}

// DeleteProject is a paid mutator transaction binding the contract method 0xd0a18958.
//
// Solidity: function deleteProject(bytes32 projectId) returns()
func (_Projects *ProjectsSession) DeleteProject(projectId [32]byte) (*types.Transaction, error) {
	return _Projects.Contract.DeleteProject(&_Projects.TransactOpts, projectId)
}

// DeleteProject is a paid mutator transaction binding the contract method 0xd0a18958.
//
// Solidity: function deleteProject(bytes32 projectId) returns()
func (_Projects *ProjectsTransactorSession) DeleteProject(projectId [32]byte) (*types.Transaction, error) {
	return _Projects.Contract.DeleteProject(&_Projects.TransactOpts, projectId)
}

// DepositProjectEscrow is a paid mutator transaction binding the contract method 0x491a8bbd.
//
// Solidity: function depositProjectEscrow(bytes32 projectId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Projects *ProjectsTransactor) DepositProjectEscrow(opts *bind.TransactOpts, projectId [32]byte, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "depositProjectEscrow", projectId, amount, deadline, v, r, s)
}

// DepositProjectEscrow is a paid mutator transaction binding the contract method 0x491a8bbd.
//
// Solidity: function depositProjectEscrow(bytes32 projectId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Projects *ProjectsSession) DepositProjectEscrow(projectId [32]byte, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Projects.Contract.DepositProjectEscrow(&_Projects.TransactOpts, projectId, amount, deadline, v, r, s)
}

// DepositProjectEscrow is a paid mutator transaction binding the contract method 0x491a8bbd.
//
// Solidity: function depositProjectEscrow(bytes32 projectId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Projects *ProjectsTransactorSession) DepositProjectEscrow(projectId [32]byte, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Projects.Contract.DepositProjectEscrow(&_Projects.TransactOpts, projectId, amount, deadline, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Projects *ProjectsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Projects *ProjectsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Projects.Contract.GrantRole(&_Projects.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Projects *ProjectsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Projects.Contract.GrantRole(&_Projects.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x7aee1c6b.
//
// Solidity: function initialize(address[] admins, address registry, bool grantImporterRole) returns()
func (_Projects *ProjectsTransactor) Initialize(opts *bind.TransactOpts, admins []common.Address, registry common.Address, grantImporterRole bool) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "initialize", admins, registry, grantImporterRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x7aee1c6b.
//
// Solidity: function initialize(address[] admins, address registry, bool grantImporterRole) returns()
func (_Projects *ProjectsSession) Initialize(admins []common.Address, registry common.Address, grantImporterRole bool) (*types.Transaction, error) {
	return _Projects.Contract.Initialize(&_Projects.TransactOpts, admins, registry, grantImporterRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x7aee1c6b.
//
// Solidity: function initialize(address[] admins, address registry, bool grantImporterRole) returns()
func (_Projects *ProjectsTransactorSession) Initialize(admins []common.Address, registry common.Address, grantImporterRole bool) (*types.Transaction, error) {
	return _Projects.Contract.Initialize(&_Projects.TransactOpts, admins, registry, grantImporterRole)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Projects *ProjectsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Projects *ProjectsSession) Pause() (*types.Transaction, error) {
	return _Projects.Contract.Pause(&_Projects.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Projects *ProjectsTransactorSession) Pause() (*types.Transaction, error) {
	return _Projects.Contract.Pause(&_Projects.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Projects *ProjectsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Projects *ProjectsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Projects.Contract.RenounceRole(&_Projects.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Projects *ProjectsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Projects.Contract.RenounceRole(&_Projects.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Projects *ProjectsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Projects *ProjectsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Projects.Contract.RevokeRole(&_Projects.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Projects *ProjectsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Projects.Contract.RevokeRole(&_Projects.TransactOpts, role, account)
}

// SetProjectContent is a paid mutator transaction binding the contract method 0x67a89099.
//
// Solidity: function setProjectContent(bytes32 projectId, string content, bytes32 checksum) returns()
func (_Projects *ProjectsTransactor) SetProjectContent(opts *bind.TransactOpts, projectId [32]byte, content string, checksum [32]byte) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "setProjectContent", projectId, content, checksum)
}

// SetProjectContent is a paid mutator transaction binding the contract method 0x67a89099.
//
// Solidity: function setProjectContent(bytes32 projectId, string content, bytes32 checksum) returns()
func (_Projects *ProjectsSession) SetProjectContent(projectId [32]byte, content string, checksum [32]byte) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectContent(&_Projects.TransactOpts, projectId, content, checksum)
}

// SetProjectContent is a paid mutator transaction binding the contract method 0x67a89099.
//
// Solidity: function setProjectContent(bytes32 projectId, string content, bytes32 checksum) returns()
func (_Projects *ProjectsTransactorSession) SetProjectContent(projectId [32]byte, content string, checksum [32]byte) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectContent(&_Projects.TransactOpts, projectId, content, checksum)
}

// SetProjectEscrowImpl is a paid mutator transaction binding the contract method 0x83945d4b.
//
// Solidity: function setProjectEscrowImpl(bytes32 projectId, uint256 decrease, uint256 increase) returns()
func (_Projects *ProjectsTransactor) SetProjectEscrowImpl(opts *bind.TransactOpts, projectId [32]byte, decrease *big.Int, increase *big.Int) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "setProjectEscrowImpl", projectId, decrease, increase)
}

// SetProjectEscrowImpl is a paid mutator transaction binding the contract method 0x83945d4b.
//
// Solidity: function setProjectEscrowImpl(bytes32 projectId, uint256 decrease, uint256 increase) returns()
func (_Projects *ProjectsSession) SetProjectEscrowImpl(projectId [32]byte, decrease *big.Int, increase *big.Int) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectEscrowImpl(&_Projects.TransactOpts, projectId, decrease, increase)
}

// SetProjectEscrowImpl is a paid mutator transaction binding the contract method 0x83945d4b.
//
// Solidity: function setProjectEscrowImpl(bytes32 projectId, uint256 decrease, uint256 increase) returns()
func (_Projects *ProjectsTransactorSession) SetProjectEscrowImpl(projectId [32]byte, decrease *big.Int, increase *big.Int) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectEscrowImpl(&_Projects.TransactOpts, projectId, decrease, increase)
}

// SetProjectMetadata is a paid mutator transaction binding the contract method 0x755dc597.
//
// Solidity: function setProjectMetadata(bytes32 projectId, string metadata) returns()
func (_Projects *ProjectsTransactor) SetProjectMetadata(opts *bind.TransactOpts, projectId [32]byte, metadata string) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "setProjectMetadata", projectId, metadata)
}

// SetProjectMetadata is a paid mutator transaction binding the contract method 0x755dc597.
//
// Solidity: function setProjectMetadata(bytes32 projectId, string metadata) returns()
func (_Projects *ProjectsSession) SetProjectMetadata(projectId [32]byte, metadata string) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectMetadata(&_Projects.TransactOpts, projectId, metadata)
}

// SetProjectMetadata is a paid mutator transaction binding the contract method 0x755dc597.
//
// Solidity: function setProjectMetadata(bytes32 projectId, string metadata) returns()
func (_Projects *ProjectsTransactorSession) SetProjectMetadata(projectId [32]byte, metadata string) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectMetadata(&_Projects.TransactOpts, projectId, metadata)
}

// SetProjectOwner is a paid mutator transaction binding the contract method 0xe6416b21.
//
// Solidity: function setProjectOwner(bytes32 projectId, address owner) returns()
func (_Projects *ProjectsTransactor) SetProjectOwner(opts *bind.TransactOpts, projectId [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "setProjectOwner", projectId, owner)
}

// SetProjectOwner is a paid mutator transaction binding the contract method 0xe6416b21.
//
// Solidity: function setProjectOwner(bytes32 projectId, address owner) returns()
func (_Projects *ProjectsSession) SetProjectOwner(projectId [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectOwner(&_Projects.TransactOpts, projectId, owner)
}

// SetProjectOwner is a paid mutator transaction binding the contract method 0xe6416b21.
//
// Solidity: function setProjectOwner(bytes32 projectId, address owner) returns()
func (_Projects *ProjectsTransactorSession) SetProjectOwner(projectId [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectOwner(&_Projects.TransactOpts, projectId, owner)
}

// SetProjectProps is a paid mutator transaction binding the contract method 0x862514b3.
//
// Solidity: function setProjectProps(bytes32 projectId, string name, string email) returns()
func (_Projects *ProjectsTransactor) SetProjectProps(opts *bind.TransactOpts, projectId [32]byte, name string, email string) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "setProjectProps", projectId, name, email)
}

// SetProjectProps is a paid mutator transaction binding the contract method 0x862514b3.
//
// Solidity: function setProjectProps(bytes32 projectId, string name, string email) returns()
func (_Projects *ProjectsSession) SetProjectProps(projectId [32]byte, name string, email string) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectProps(&_Projects.TransactOpts, projectId, name, email)
}

// SetProjectProps is a paid mutator transaction binding the contract method 0x862514b3.
//
// Solidity: function setProjectProps(bytes32 projectId, string name, string email) returns()
func (_Projects *ProjectsTransactorSession) SetProjectProps(projectId [32]byte, name string, email string) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectProps(&_Projects.TransactOpts, projectId, name, email)
}

// SetProjectReserveImpl is a paid mutator transaction binding the contract method 0xda481dd6.
//
// Solidity: function setProjectReserveImpl(bytes32 projectId, uint256 decrease, uint256 increase) returns()
func (_Projects *ProjectsTransactor) SetProjectReserveImpl(opts *bind.TransactOpts, projectId [32]byte, decrease *big.Int, increase *big.Int) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "setProjectReserveImpl", projectId, decrease, increase)
}

// SetProjectReserveImpl is a paid mutator transaction binding the contract method 0xda481dd6.
//
// Solidity: function setProjectReserveImpl(bytes32 projectId, uint256 decrease, uint256 increase) returns()
func (_Projects *ProjectsSession) SetProjectReserveImpl(projectId [32]byte, decrease *big.Int, increase *big.Int) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectReserveImpl(&_Projects.TransactOpts, projectId, decrease, increase)
}

// SetProjectReserveImpl is a paid mutator transaction binding the contract method 0xda481dd6.
//
// Solidity: function setProjectReserveImpl(bytes32 projectId, uint256 decrease, uint256 increase) returns()
func (_Projects *ProjectsTransactorSession) SetProjectReserveImpl(projectId [32]byte, decrease *big.Int, increase *big.Int) (*types.Transaction, error) {
	return _Projects.Contract.SetProjectReserveImpl(&_Projects.TransactOpts, projectId, decrease, increase)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Projects *ProjectsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Projects *ProjectsSession) Unpause() (*types.Transaction, error) {
	return _Projects.Contract.Unpause(&_Projects.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Projects *ProjectsTransactorSession) Unpause() (*types.Transaction, error) {
	return _Projects.Contract.Unpause(&_Projects.TransactOpts)
}

// UnsafeImportData is a paid mutator transaction binding the contract method 0x737f0372.
//
// Solidity: function unsafeImportData((bytes32,address,string,string,uint256,uint256,string,bytes32,string)[] projects, bool revokeImporterRole) returns()
func (_Projects *ProjectsTransactor) UnsafeImportData(opts *bind.TransactOpts, projects []EarthfastProject, revokeImporterRole bool) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "unsafeImportData", projects, revokeImporterRole)
}

// UnsafeImportData is a paid mutator transaction binding the contract method 0x737f0372.
//
// Solidity: function unsafeImportData((bytes32,address,string,string,uint256,uint256,string,bytes32,string)[] projects, bool revokeImporterRole) returns()
func (_Projects *ProjectsSession) UnsafeImportData(projects []EarthfastProject, revokeImporterRole bool) (*types.Transaction, error) {
	return _Projects.Contract.UnsafeImportData(&_Projects.TransactOpts, projects, revokeImporterRole)
}

// UnsafeImportData is a paid mutator transaction binding the contract method 0x737f0372.
//
// Solidity: function unsafeImportData((bytes32,address,string,string,uint256,uint256,string,bytes32,string)[] projects, bool revokeImporterRole) returns()
func (_Projects *ProjectsTransactorSession) UnsafeImportData(projects []EarthfastProject, revokeImporterRole bool) (*types.Transaction, error) {
	return _Projects.Contract.UnsafeImportData(&_Projects.TransactOpts, projects, revokeImporterRole)
}

// UnsafeSetEscrows is a paid mutator transaction binding the contract method 0xf1bd1b8a.
//
// Solidity: function unsafeSetEscrows(uint256 skip, uint256 size, uint256 mul, uint256 div) returns()
func (_Projects *ProjectsTransactor) UnsafeSetEscrows(opts *bind.TransactOpts, skip *big.Int, size *big.Int, mul *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "unsafeSetEscrows", skip, size, mul, div)
}

// UnsafeSetEscrows is a paid mutator transaction binding the contract method 0xf1bd1b8a.
//
// Solidity: function unsafeSetEscrows(uint256 skip, uint256 size, uint256 mul, uint256 div) returns()
func (_Projects *ProjectsSession) UnsafeSetEscrows(skip *big.Int, size *big.Int, mul *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Projects.Contract.UnsafeSetEscrows(&_Projects.TransactOpts, skip, size, mul, div)
}

// UnsafeSetEscrows is a paid mutator transaction binding the contract method 0xf1bd1b8a.
//
// Solidity: function unsafeSetEscrows(uint256 skip, uint256 size, uint256 mul, uint256 div) returns()
func (_Projects *ProjectsTransactorSession) UnsafeSetEscrows(skip *big.Int, size *big.Int, mul *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Projects.Contract.UnsafeSetEscrows(&_Projects.TransactOpts, skip, size, mul, div)
}

// UnsafeSetRegistry is a paid mutator transaction binding the contract method 0xb9a2adf0.
//
// Solidity: function unsafeSetRegistry(address registry) returns()
func (_Projects *ProjectsTransactor) UnsafeSetRegistry(opts *bind.TransactOpts, registry common.Address) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "unsafeSetRegistry", registry)
}

// UnsafeSetRegistry is a paid mutator transaction binding the contract method 0xb9a2adf0.
//
// Solidity: function unsafeSetRegistry(address registry) returns()
func (_Projects *ProjectsSession) UnsafeSetRegistry(registry common.Address) (*types.Transaction, error) {
	return _Projects.Contract.UnsafeSetRegistry(&_Projects.TransactOpts, registry)
}

// UnsafeSetRegistry is a paid mutator transaction binding the contract method 0xb9a2adf0.
//
// Solidity: function unsafeSetRegistry(address registry) returns()
func (_Projects *ProjectsTransactorSession) UnsafeSetRegistry(registry common.Address) (*types.Transaction, error) {
	return _Projects.Contract.UnsafeSetRegistry(&_Projects.TransactOpts, registry)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Projects *ProjectsTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Projects *ProjectsSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Projects.Contract.UpgradeTo(&_Projects.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Projects *ProjectsTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Projects.Contract.UpgradeTo(&_Projects.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Projects *ProjectsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Projects *ProjectsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Projects.Contract.UpgradeToAndCall(&_Projects.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Projects *ProjectsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Projects.Contract.UpgradeToAndCall(&_Projects.TransactOpts, newImplementation, data)
}

// WithdrawProjectEscrow is a paid mutator transaction binding the contract method 0x12565dd0.
//
// Solidity: function withdrawProjectEscrow(bytes32 projectId, uint256 amount, address to) returns()
func (_Projects *ProjectsTransactor) WithdrawProjectEscrow(opts *bind.TransactOpts, projectId [32]byte, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Projects.contract.Transact(opts, "withdrawProjectEscrow", projectId, amount, to)
}

// WithdrawProjectEscrow is a paid mutator transaction binding the contract method 0x12565dd0.
//
// Solidity: function withdrawProjectEscrow(bytes32 projectId, uint256 amount, address to) returns()
func (_Projects *ProjectsSession) WithdrawProjectEscrow(projectId [32]byte, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Projects.Contract.WithdrawProjectEscrow(&_Projects.TransactOpts, projectId, amount, to)
}

// WithdrawProjectEscrow is a paid mutator transaction binding the contract method 0x12565dd0.
//
// Solidity: function withdrawProjectEscrow(bytes32 projectId, uint256 amount, address to) returns()
func (_Projects *ProjectsTransactorSession) WithdrawProjectEscrow(projectId [32]byte, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Projects.Contract.WithdrawProjectEscrow(&_Projects.TransactOpts, projectId, amount, to)
}

// ProjectsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Projects contract.
type ProjectsAdminChangedIterator struct {
	Event *ProjectsAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProjectsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsAdminChanged)
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
		it.Event = new(ProjectsAdminChanged)
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
func (it *ProjectsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsAdminChanged represents a AdminChanged event raised by the Projects contract.
type ProjectsAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Projects *ProjectsFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ProjectsAdminChangedIterator, error) {

	logs, sub, err := _Projects.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ProjectsAdminChangedIterator{contract: _Projects.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Projects *ProjectsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ProjectsAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Projects.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsAdminChanged)
				if err := _Projects.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Projects *ProjectsFilterer) ParseAdminChanged(log types.Log) (*ProjectsAdminChanged, error) {
	event := new(ProjectsAdminChanged)
	if err := _Projects.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Projects contract.
type ProjectsBeaconUpgradedIterator struct {
	Event *ProjectsBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ProjectsBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsBeaconUpgraded)
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
		it.Event = new(ProjectsBeaconUpgraded)
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
func (it *ProjectsBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsBeaconUpgraded represents a BeaconUpgraded event raised by the Projects contract.
type ProjectsBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Projects *ProjectsFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ProjectsBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Projects.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsBeaconUpgradedIterator{contract: _Projects.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Projects *ProjectsFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ProjectsBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Projects.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsBeaconUpgraded)
				if err := _Projects.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Projects *ProjectsFilterer) ParseBeaconUpgraded(log types.Log) (*ProjectsBeaconUpgraded, error) {
	event := new(ProjectsBeaconUpgraded)
	if err := _Projects.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Projects contract.
type ProjectsInitializedIterator struct {
	Event *ProjectsInitialized // Event containing the contract specifics and raw log

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
func (it *ProjectsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsInitialized)
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
		it.Event = new(ProjectsInitialized)
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
func (it *ProjectsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsInitialized represents a Initialized event raised by the Projects contract.
type ProjectsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Projects *ProjectsFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProjectsInitializedIterator, error) {

	logs, sub, err := _Projects.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProjectsInitializedIterator{contract: _Projects.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Projects *ProjectsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProjectsInitialized) (event.Subscription, error) {

	logs, sub, err := _Projects.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsInitialized)
				if err := _Projects.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Projects *ProjectsFilterer) ParseInitialized(log types.Log) (*ProjectsInitialized, error) {
	event := new(ProjectsInitialized)
	if err := _Projects.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Projects contract.
type ProjectsPausedIterator struct {
	Event *ProjectsPaused // Event containing the contract specifics and raw log

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
func (it *ProjectsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsPaused)
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
		it.Event = new(ProjectsPaused)
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
func (it *ProjectsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsPaused represents a Paused event raised by the Projects contract.
type ProjectsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Projects *ProjectsFilterer) FilterPaused(opts *bind.FilterOpts) (*ProjectsPausedIterator, error) {

	logs, sub, err := _Projects.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ProjectsPausedIterator{contract: _Projects.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Projects *ProjectsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ProjectsPaused) (event.Subscription, error) {

	logs, sub, err := _Projects.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsPaused)
				if err := _Projects.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Projects *ProjectsFilterer) ParsePaused(log types.Log) (*ProjectsPaused, error) {
	event := new(ProjectsPaused)
	if err := _Projects.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsProjectContentChangedIterator is returned from FilterProjectContentChanged and is used to iterate over the raw logs and unpacked data for ProjectContentChanged events raised by the Projects contract.
type ProjectsProjectContentChangedIterator struct {
	Event *ProjectsProjectContentChanged // Event containing the contract specifics and raw log

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
func (it *ProjectsProjectContentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsProjectContentChanged)
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
		it.Event = new(ProjectsProjectContentChanged)
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
func (it *ProjectsProjectContentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsProjectContentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsProjectContentChanged represents a ProjectContentChanged event raised by the Projects contract.
type ProjectsProjectContentChanged struct {
	ProjectId   [32]byte
	OldContent  string
	OldChecksum [32]byte
	NewContent  string
	NewChecksum [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProjectContentChanged is a free log retrieval operation binding the contract event 0x34fd38e6eebd32ffe18208627966f06ce8f8756fa72a8d8dd2ab1750c0f11e23.
//
// Solidity: event ProjectContentChanged(bytes32 indexed projectId, string oldContent, bytes32 oldChecksum, string newContent, bytes32 newChecksum)
func (_Projects *ProjectsFilterer) FilterProjectContentChanged(opts *bind.FilterOpts, projectId [][32]byte) (*ProjectsProjectContentChangedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Projects.contract.FilterLogs(opts, "ProjectContentChanged", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsProjectContentChangedIterator{contract: _Projects.contract, event: "ProjectContentChanged", logs: logs, sub: sub}, nil
}

// WatchProjectContentChanged is a free log subscription operation binding the contract event 0x34fd38e6eebd32ffe18208627966f06ce8f8756fa72a8d8dd2ab1750c0f11e23.
//
// Solidity: event ProjectContentChanged(bytes32 indexed projectId, string oldContent, bytes32 oldChecksum, string newContent, bytes32 newChecksum)
func (_Projects *ProjectsFilterer) WatchProjectContentChanged(opts *bind.WatchOpts, sink chan<- *ProjectsProjectContentChanged, projectId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Projects.contract.WatchLogs(opts, "ProjectContentChanged", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsProjectContentChanged)
				if err := _Projects.contract.UnpackLog(event, "ProjectContentChanged", log); err != nil {
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

// ParseProjectContentChanged is a log parse operation binding the contract event 0x34fd38e6eebd32ffe18208627966f06ce8f8756fa72a8d8dd2ab1750c0f11e23.
//
// Solidity: event ProjectContentChanged(bytes32 indexed projectId, string oldContent, bytes32 oldChecksum, string newContent, bytes32 newChecksum)
func (_Projects *ProjectsFilterer) ParseProjectContentChanged(log types.Log) (*ProjectsProjectContentChanged, error) {
	event := new(ProjectsProjectContentChanged)
	if err := _Projects.contract.UnpackLog(event, "ProjectContentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsProjectCreatedIterator is returned from FilterProjectCreated and is used to iterate over the raw logs and unpacked data for ProjectCreated events raised by the Projects contract.
type ProjectsProjectCreatedIterator struct {
	Event *ProjectsProjectCreated // Event containing the contract specifics and raw log

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
func (it *ProjectsProjectCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsProjectCreated)
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
		it.Event = new(ProjectsProjectCreated)
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
func (it *ProjectsProjectCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsProjectCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsProjectCreated represents a ProjectCreated event raised by the Projects contract.
type ProjectsProjectCreated struct {
	ProjectId [32]byte
	Owner     common.Address
	Name      string
	Email     string
	Content   string
	Checksum  [32]byte
	Metadata  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectCreated is a free log retrieval operation binding the contract event 0xc9a2563ef6e863a3e90b633d77bef27cc637d965cadcc8e629740f13d6c5b83e.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, string name, string email, string content, bytes32 checksum, string metadata)
func (_Projects *ProjectsFilterer) FilterProjectCreated(opts *bind.FilterOpts, projectId [][32]byte, owner []common.Address) (*ProjectsProjectCreatedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Projects.contract.FilterLogs(opts, "ProjectCreated", projectIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsProjectCreatedIterator{contract: _Projects.contract, event: "ProjectCreated", logs: logs, sub: sub}, nil
}

// WatchProjectCreated is a free log subscription operation binding the contract event 0xc9a2563ef6e863a3e90b633d77bef27cc637d965cadcc8e629740f13d6c5b83e.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, string name, string email, string content, bytes32 checksum, string metadata)
func (_Projects *ProjectsFilterer) WatchProjectCreated(opts *bind.WatchOpts, sink chan<- *ProjectsProjectCreated, projectId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Projects.contract.WatchLogs(opts, "ProjectCreated", projectIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsProjectCreated)
				if err := _Projects.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
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

// ParseProjectCreated is a log parse operation binding the contract event 0xc9a2563ef6e863a3e90b633d77bef27cc637d965cadcc8e629740f13d6c5b83e.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, string name, string email, string content, bytes32 checksum, string metadata)
func (_Projects *ProjectsFilterer) ParseProjectCreated(log types.Log) (*ProjectsProjectCreated, error) {
	event := new(ProjectsProjectCreated)
	if err := _Projects.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsProjectDeletedIterator is returned from FilterProjectDeleted and is used to iterate over the raw logs and unpacked data for ProjectDeleted events raised by the Projects contract.
type ProjectsProjectDeletedIterator struct {
	Event *ProjectsProjectDeleted // Event containing the contract specifics and raw log

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
func (it *ProjectsProjectDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsProjectDeleted)
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
		it.Event = new(ProjectsProjectDeleted)
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
func (it *ProjectsProjectDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsProjectDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsProjectDeleted represents a ProjectDeleted event raised by the Projects contract.
type ProjectsProjectDeleted struct {
	ProjectId [32]byte
	Owner     common.Address
	Name      string
	Email     string
	Content   string
	Checksum  [32]byte
	Metadata  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectDeleted is a free log retrieval operation binding the contract event 0x0dfede72020edf94c02fe864fbd28e7615a81fb1ddde6c3224c7236b5b520a67.
//
// Solidity: event ProjectDeleted(bytes32 indexed projectId, address indexed owner, string name, string email, string content, bytes32 checksum, string metadata)
func (_Projects *ProjectsFilterer) FilterProjectDeleted(opts *bind.FilterOpts, projectId [][32]byte, owner []common.Address) (*ProjectsProjectDeletedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Projects.contract.FilterLogs(opts, "ProjectDeleted", projectIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsProjectDeletedIterator{contract: _Projects.contract, event: "ProjectDeleted", logs: logs, sub: sub}, nil
}

// WatchProjectDeleted is a free log subscription operation binding the contract event 0x0dfede72020edf94c02fe864fbd28e7615a81fb1ddde6c3224c7236b5b520a67.
//
// Solidity: event ProjectDeleted(bytes32 indexed projectId, address indexed owner, string name, string email, string content, bytes32 checksum, string metadata)
func (_Projects *ProjectsFilterer) WatchProjectDeleted(opts *bind.WatchOpts, sink chan<- *ProjectsProjectDeleted, projectId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Projects.contract.WatchLogs(opts, "ProjectDeleted", projectIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsProjectDeleted)
				if err := _Projects.contract.UnpackLog(event, "ProjectDeleted", log); err != nil {
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

// ParseProjectDeleted is a log parse operation binding the contract event 0x0dfede72020edf94c02fe864fbd28e7615a81fb1ddde6c3224c7236b5b520a67.
//
// Solidity: event ProjectDeleted(bytes32 indexed projectId, address indexed owner, string name, string email, string content, bytes32 checksum, string metadata)
func (_Projects *ProjectsFilterer) ParseProjectDeleted(log types.Log) (*ProjectsProjectDeleted, error) {
	event := new(ProjectsProjectDeleted)
	if err := _Projects.contract.UnpackLog(event, "ProjectDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsProjectEscrowChangedIterator is returned from FilterProjectEscrowChanged and is used to iterate over the raw logs and unpacked data for ProjectEscrowChanged events raised by the Projects contract.
type ProjectsProjectEscrowChangedIterator struct {
	Event *ProjectsProjectEscrowChanged // Event containing the contract specifics and raw log

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
func (it *ProjectsProjectEscrowChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsProjectEscrowChanged)
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
		it.Event = new(ProjectsProjectEscrowChanged)
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
func (it *ProjectsProjectEscrowChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsProjectEscrowChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsProjectEscrowChanged represents a ProjectEscrowChanged event raised by the Projects contract.
type ProjectsProjectEscrowChanged struct {
	ProjectId [32]byte
	OldEscrow *big.Int
	NewEscrow *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectEscrowChanged is a free log retrieval operation binding the contract event 0x2a5f5e56e89c57ee03d89ec4718d7ae9911b09486c8ac692e4908ef267eb2804.
//
// Solidity: event ProjectEscrowChanged(bytes32 indexed projectId, uint256 oldEscrow, uint256 newEscrow)
func (_Projects *ProjectsFilterer) FilterProjectEscrowChanged(opts *bind.FilterOpts, projectId [][32]byte) (*ProjectsProjectEscrowChangedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Projects.contract.FilterLogs(opts, "ProjectEscrowChanged", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsProjectEscrowChangedIterator{contract: _Projects.contract, event: "ProjectEscrowChanged", logs: logs, sub: sub}, nil
}

// WatchProjectEscrowChanged is a free log subscription operation binding the contract event 0x2a5f5e56e89c57ee03d89ec4718d7ae9911b09486c8ac692e4908ef267eb2804.
//
// Solidity: event ProjectEscrowChanged(bytes32 indexed projectId, uint256 oldEscrow, uint256 newEscrow)
func (_Projects *ProjectsFilterer) WatchProjectEscrowChanged(opts *bind.WatchOpts, sink chan<- *ProjectsProjectEscrowChanged, projectId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Projects.contract.WatchLogs(opts, "ProjectEscrowChanged", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsProjectEscrowChanged)
				if err := _Projects.contract.UnpackLog(event, "ProjectEscrowChanged", log); err != nil {
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

// ParseProjectEscrowChanged is a log parse operation binding the contract event 0x2a5f5e56e89c57ee03d89ec4718d7ae9911b09486c8ac692e4908ef267eb2804.
//
// Solidity: event ProjectEscrowChanged(bytes32 indexed projectId, uint256 oldEscrow, uint256 newEscrow)
func (_Projects *ProjectsFilterer) ParseProjectEscrowChanged(log types.Log) (*ProjectsProjectEscrowChanged, error) {
	event := new(ProjectsProjectEscrowChanged)
	if err := _Projects.contract.UnpackLog(event, "ProjectEscrowChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsProjectMetadataChangedIterator is returned from FilterProjectMetadataChanged and is used to iterate over the raw logs and unpacked data for ProjectMetadataChanged events raised by the Projects contract.
type ProjectsProjectMetadataChangedIterator struct {
	Event *ProjectsProjectMetadataChanged // Event containing the contract specifics and raw log

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
func (it *ProjectsProjectMetadataChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsProjectMetadataChanged)
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
		it.Event = new(ProjectsProjectMetadataChanged)
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
func (it *ProjectsProjectMetadataChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsProjectMetadataChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsProjectMetadataChanged represents a ProjectMetadataChanged event raised by the Projects contract.
type ProjectsProjectMetadataChanged struct {
	ProjectId   [32]byte
	OldMetadata string
	NewMetadata string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProjectMetadataChanged is a free log retrieval operation binding the contract event 0xc31c69f7d2faef3482aac0691da695aa5839c98a11dbbe8797fd91d7f8a4f6e2.
//
// Solidity: event ProjectMetadataChanged(bytes32 indexed projectId, string oldMetadata, string newMetadata)
func (_Projects *ProjectsFilterer) FilterProjectMetadataChanged(opts *bind.FilterOpts, projectId [][32]byte) (*ProjectsProjectMetadataChangedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Projects.contract.FilterLogs(opts, "ProjectMetadataChanged", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsProjectMetadataChangedIterator{contract: _Projects.contract, event: "ProjectMetadataChanged", logs: logs, sub: sub}, nil
}

// WatchProjectMetadataChanged is a free log subscription operation binding the contract event 0xc31c69f7d2faef3482aac0691da695aa5839c98a11dbbe8797fd91d7f8a4f6e2.
//
// Solidity: event ProjectMetadataChanged(bytes32 indexed projectId, string oldMetadata, string newMetadata)
func (_Projects *ProjectsFilterer) WatchProjectMetadataChanged(opts *bind.WatchOpts, sink chan<- *ProjectsProjectMetadataChanged, projectId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Projects.contract.WatchLogs(opts, "ProjectMetadataChanged", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsProjectMetadataChanged)
				if err := _Projects.contract.UnpackLog(event, "ProjectMetadataChanged", log); err != nil {
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

// ParseProjectMetadataChanged is a log parse operation binding the contract event 0xc31c69f7d2faef3482aac0691da695aa5839c98a11dbbe8797fd91d7f8a4f6e2.
//
// Solidity: event ProjectMetadataChanged(bytes32 indexed projectId, string oldMetadata, string newMetadata)
func (_Projects *ProjectsFilterer) ParseProjectMetadataChanged(log types.Log) (*ProjectsProjectMetadataChanged, error) {
	event := new(ProjectsProjectMetadataChanged)
	if err := _Projects.contract.UnpackLog(event, "ProjectMetadataChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsProjectOwnerChangedIterator is returned from FilterProjectOwnerChanged and is used to iterate over the raw logs and unpacked data for ProjectOwnerChanged events raised by the Projects contract.
type ProjectsProjectOwnerChangedIterator struct {
	Event *ProjectsProjectOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ProjectsProjectOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsProjectOwnerChanged)
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
		it.Event = new(ProjectsProjectOwnerChanged)
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
func (it *ProjectsProjectOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsProjectOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsProjectOwnerChanged represents a ProjectOwnerChanged event raised by the Projects contract.
type ProjectsProjectOwnerChanged struct {
	ProjectId [32]byte
	OldOwner  common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectOwnerChanged is a free log retrieval operation binding the contract event 0x1e33df2f19cae9abe592944dff51706da0aa04f97851c3113235a7056a338e92.
//
// Solidity: event ProjectOwnerChanged(bytes32 indexed projectId, address indexed oldOwner, address indexed newOwner)
func (_Projects *ProjectsFilterer) FilterProjectOwnerChanged(opts *bind.FilterOpts, projectId [][32]byte, oldOwner []common.Address, newOwner []common.Address) (*ProjectsProjectOwnerChangedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Projects.contract.FilterLogs(opts, "ProjectOwnerChanged", projectIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsProjectOwnerChangedIterator{contract: _Projects.contract, event: "ProjectOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchProjectOwnerChanged is a free log subscription operation binding the contract event 0x1e33df2f19cae9abe592944dff51706da0aa04f97851c3113235a7056a338e92.
//
// Solidity: event ProjectOwnerChanged(bytes32 indexed projectId, address indexed oldOwner, address indexed newOwner)
func (_Projects *ProjectsFilterer) WatchProjectOwnerChanged(opts *bind.WatchOpts, sink chan<- *ProjectsProjectOwnerChanged, projectId [][32]byte, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Projects.contract.WatchLogs(opts, "ProjectOwnerChanged", projectIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsProjectOwnerChanged)
				if err := _Projects.contract.UnpackLog(event, "ProjectOwnerChanged", log); err != nil {
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

// ParseProjectOwnerChanged is a log parse operation binding the contract event 0x1e33df2f19cae9abe592944dff51706da0aa04f97851c3113235a7056a338e92.
//
// Solidity: event ProjectOwnerChanged(bytes32 indexed projectId, address indexed oldOwner, address indexed newOwner)
func (_Projects *ProjectsFilterer) ParseProjectOwnerChanged(log types.Log) (*ProjectsProjectOwnerChanged, error) {
	event := new(ProjectsProjectOwnerChanged)
	if err := _Projects.contract.UnpackLog(event, "ProjectOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsProjectPropsChangedIterator is returned from FilterProjectPropsChanged and is used to iterate over the raw logs and unpacked data for ProjectPropsChanged events raised by the Projects contract.
type ProjectsProjectPropsChangedIterator struct {
	Event *ProjectsProjectPropsChanged // Event containing the contract specifics and raw log

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
func (it *ProjectsProjectPropsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsProjectPropsChanged)
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
		it.Event = new(ProjectsProjectPropsChanged)
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
func (it *ProjectsProjectPropsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsProjectPropsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsProjectPropsChanged represents a ProjectPropsChanged event raised by the Projects contract.
type ProjectsProjectPropsChanged struct {
	ProjectId [32]byte
	OldName   string
	OldEmail  string
	NewName   string
	NewEmail  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectPropsChanged is a free log retrieval operation binding the contract event 0xfbe7b9226d0807aa630b0494cc4a185400061b1a4c1d0df3bc556f094498b16e.
//
// Solidity: event ProjectPropsChanged(bytes32 indexed projectId, string oldName, string oldEmail, string newName, string newEmail)
func (_Projects *ProjectsFilterer) FilterProjectPropsChanged(opts *bind.FilterOpts, projectId [][32]byte) (*ProjectsProjectPropsChangedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Projects.contract.FilterLogs(opts, "ProjectPropsChanged", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsProjectPropsChangedIterator{contract: _Projects.contract, event: "ProjectPropsChanged", logs: logs, sub: sub}, nil
}

// WatchProjectPropsChanged is a free log subscription operation binding the contract event 0xfbe7b9226d0807aa630b0494cc4a185400061b1a4c1d0df3bc556f094498b16e.
//
// Solidity: event ProjectPropsChanged(bytes32 indexed projectId, string oldName, string oldEmail, string newName, string newEmail)
func (_Projects *ProjectsFilterer) WatchProjectPropsChanged(opts *bind.WatchOpts, sink chan<- *ProjectsProjectPropsChanged, projectId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Projects.contract.WatchLogs(opts, "ProjectPropsChanged", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsProjectPropsChanged)
				if err := _Projects.contract.UnpackLog(event, "ProjectPropsChanged", log); err != nil {
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

// ParseProjectPropsChanged is a log parse operation binding the contract event 0xfbe7b9226d0807aa630b0494cc4a185400061b1a4c1d0df3bc556f094498b16e.
//
// Solidity: event ProjectPropsChanged(bytes32 indexed projectId, string oldName, string oldEmail, string newName, string newEmail)
func (_Projects *ProjectsFilterer) ParseProjectPropsChanged(log types.Log) (*ProjectsProjectPropsChanged, error) {
	event := new(ProjectsProjectPropsChanged)
	if err := _Projects.contract.UnpackLog(event, "ProjectPropsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Projects contract.
type ProjectsRoleAdminChangedIterator struct {
	Event *ProjectsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProjectsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsRoleAdminChanged)
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
		it.Event = new(ProjectsRoleAdminChanged)
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
func (it *ProjectsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsRoleAdminChanged represents a RoleAdminChanged event raised by the Projects contract.
type ProjectsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Projects *ProjectsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ProjectsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Projects.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsRoleAdminChangedIterator{contract: _Projects.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Projects *ProjectsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ProjectsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Projects.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsRoleAdminChanged)
				if err := _Projects.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Projects *ProjectsFilterer) ParseRoleAdminChanged(log types.Log) (*ProjectsRoleAdminChanged, error) {
	event := new(ProjectsRoleAdminChanged)
	if err := _Projects.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Projects contract.
type ProjectsRoleGrantedIterator struct {
	Event *ProjectsRoleGranted // Event containing the contract specifics and raw log

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
func (it *ProjectsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsRoleGranted)
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
		it.Event = new(ProjectsRoleGranted)
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
func (it *ProjectsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsRoleGranted represents a RoleGranted event raised by the Projects contract.
type ProjectsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Projects *ProjectsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProjectsRoleGrantedIterator, error) {

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

	logs, sub, err := _Projects.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsRoleGrantedIterator{contract: _Projects.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Projects *ProjectsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ProjectsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Projects.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsRoleGranted)
				if err := _Projects.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Projects *ProjectsFilterer) ParseRoleGranted(log types.Log) (*ProjectsRoleGranted, error) {
	event := new(ProjectsRoleGranted)
	if err := _Projects.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Projects contract.
type ProjectsRoleRevokedIterator struct {
	Event *ProjectsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ProjectsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsRoleRevoked)
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
		it.Event = new(ProjectsRoleRevoked)
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
func (it *ProjectsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsRoleRevoked represents a RoleRevoked event raised by the Projects contract.
type ProjectsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Projects *ProjectsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProjectsRoleRevokedIterator, error) {

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

	logs, sub, err := _Projects.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsRoleRevokedIterator{contract: _Projects.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Projects *ProjectsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ProjectsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Projects.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsRoleRevoked)
				if err := _Projects.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Projects *ProjectsFilterer) ParseRoleRevoked(log types.Log) (*ProjectsRoleRevoked, error) {
	event := new(ProjectsRoleRevoked)
	if err := _Projects.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Projects contract.
type ProjectsUnpausedIterator struct {
	Event *ProjectsUnpaused // Event containing the contract specifics and raw log

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
func (it *ProjectsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsUnpaused)
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
		it.Event = new(ProjectsUnpaused)
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
func (it *ProjectsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsUnpaused represents a Unpaused event raised by the Projects contract.
type ProjectsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Projects *ProjectsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ProjectsUnpausedIterator, error) {

	logs, sub, err := _Projects.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ProjectsUnpausedIterator{contract: _Projects.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Projects *ProjectsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ProjectsUnpaused) (event.Subscription, error) {

	logs, sub, err := _Projects.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsUnpaused)
				if err := _Projects.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Projects *ProjectsFilterer) ParseUnpaused(log types.Log) (*ProjectsUnpaused, error) {
	event := new(ProjectsUnpaused)
	if err := _Projects.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Projects contract.
type ProjectsUpgradedIterator struct {
	Event *ProjectsUpgraded // Event containing the contract specifics and raw log

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
func (it *ProjectsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectsUpgraded)
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
		it.Event = new(ProjectsUpgraded)
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
func (it *ProjectsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectsUpgraded represents a Upgraded event raised by the Projects contract.
type ProjectsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Projects *ProjectsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ProjectsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Projects.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ProjectsUpgradedIterator{contract: _Projects.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Projects *ProjectsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ProjectsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Projects.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectsUpgraded)
				if err := _Projects.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Projects *ProjectsFilterer) ParseUpgraded(log types.Log) (*ProjectsUpgraded, error) {
	event := new(ProjectsUpgraded)
	if err := _Projects.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

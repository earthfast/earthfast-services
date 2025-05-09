// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nodes

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

// EarthfastCreateNodeData is an auto generated low-level Go binding around an user-defined struct.
type EarthfastCreateNodeData struct {
	Host     string
	Region   string
	Disabled bool
	Price    *big.Int
}

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

// NodesMetaData contains all meta data concerning the Nodes contract.
var NodesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NodeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"oldDisabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newDisabled\",\"type\":\"bool\"}],\"name\":\"NodeDisabledChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldHost\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldRegion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newHost\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newRegion\",\"type\":\"string\"}],\"name\":\"NodeHostChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLastPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldNextPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"last\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"next\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEarthfastSlot\",\"name\":\"slot\",\"type\":\"tuple\"}],\"name\":\"NodePriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IMPORTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"}],\"name\":\"advanceNodeEpochImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structEarthfastCreateNodeData[]\",\"name\":\"nodes\",\"type\":\"tuple[]\"}],\"name\":\"createNodes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"nodeIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"nodeIds\",\"type\":\"bytes32[]\"}],\"name\":\"deleteNodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"}],\"name\":\"getNode\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"prices\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"projectIds\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structEarthfastNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorIdOrZero\",\"type\":\"bytes32\"}],\"name\":\"getNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorIdOrZero\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"skip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getNodes\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"prices\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"projectIds\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structEarthfastNode[]\",\"name\":\"values\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"contractEarthfastRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"},{\"internalType\":\"contractEarthfastRegistry\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"grantImporterRole\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"nodeIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"disabled\",\"type\":\"bool[]\"}],\"name\":\"setNodeDisabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"nodeIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"string[]\",\"name\":\"hosts\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"regions\",\"type\":\"string[]\"}],\"name\":\"setNodeHosts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"epochSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setNodePriceImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"nodeIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"last\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"next\",\"type\":\"bool\"}],\"internalType\":\"structEarthfastSlot\",\"name\":\"slot\",\"type\":\"tuple\"}],\"name\":\"setNodePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"epochSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"}],\"name\":\"setNodeProjectImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"prices\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"projectIds\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structEarthfastNode[]\",\"name\":\"nodes\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"revokeImporterRole\",\"type\":\"bool\"}],\"name\":\"unsafeImportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"skip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mul\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"div\",\"type\":\"uint256\"}],\"name\":\"unsafeSetPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEarthfastRegistry\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"unsafeSetRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NodesABI is the input ABI used to generate the binding from.
// Deprecated: Use NodesMetaData.ABI instead.
var NodesABI = NodesMetaData.ABI

// Nodes is an auto generated Go binding around an Ethereum contract.
type Nodes struct {
	NodesCaller     // Read-only binding to the contract
	NodesTransactor // Write-only binding to the contract
	NodesFilterer   // Log filterer for contract events
}

// NodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodesSession struct {
	Contract     *Nodes            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodesCallerSession struct {
	Contract *NodesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodesTransactorSession struct {
	Contract     *NodesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodesRaw struct {
	Contract *Nodes // Generic contract binding to access the raw methods on
}

// NodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodesCallerRaw struct {
	Contract *NodesCaller // Generic read-only contract binding to access the raw methods on
}

// NodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodesTransactorRaw struct {
	Contract *NodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodes creates a new instance of Nodes, bound to a specific deployed contract.
func NewNodes(address common.Address, backend bind.ContractBackend) (*Nodes, error) {
	contract, err := bindNodes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nodes{NodesCaller: NodesCaller{contract: contract}, NodesTransactor: NodesTransactor{contract: contract}, NodesFilterer: NodesFilterer{contract: contract}}, nil
}

// NewNodesCaller creates a new read-only instance of Nodes, bound to a specific deployed contract.
func NewNodesCaller(address common.Address, caller bind.ContractCaller) (*NodesCaller, error) {
	contract, err := bindNodes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodesCaller{contract: contract}, nil
}

// NewNodesTransactor creates a new write-only instance of Nodes, bound to a specific deployed contract.
func NewNodesTransactor(address common.Address, transactor bind.ContractTransactor) (*NodesTransactor, error) {
	contract, err := bindNodes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodesTransactor{contract: contract}, nil
}

// NewNodesFilterer creates a new log filterer instance of Nodes, bound to a specific deployed contract.
func NewNodesFilterer(address common.Address, filterer bind.ContractFilterer) (*NodesFilterer, error) {
	contract, err := bindNodes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodesFilterer{contract: contract}, nil
}

// bindNodes binds a generic wrapper to an already deployed contract.
func bindNodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodes *NodesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nodes.Contract.NodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodes *NodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodes.Contract.NodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodes *NodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodes.Contract.NodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodes *NodesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodes *NodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodes *NodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodes.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Nodes *NodesCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Nodes *NodesSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Nodes.Contract.DEFAULTADMINROLE(&_Nodes.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Nodes *NodesCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Nodes.Contract.DEFAULTADMINROLE(&_Nodes.CallOpts)
}

// IMPORTERROLE is a free data retrieval call binding the contract method 0x6fa2a067.
//
// Solidity: function IMPORTER_ROLE() view returns(bytes32)
func (_Nodes *NodesCaller) IMPORTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "IMPORTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IMPORTERROLE is a free data retrieval call binding the contract method 0x6fa2a067.
//
// Solidity: function IMPORTER_ROLE() view returns(bytes32)
func (_Nodes *NodesSession) IMPORTERROLE() ([32]byte, error) {
	return _Nodes.Contract.IMPORTERROLE(&_Nodes.CallOpts)
}

// IMPORTERROLE is a free data retrieval call binding the contract method 0x6fa2a067.
//
// Solidity: function IMPORTER_ROLE() view returns(bytes32)
func (_Nodes *NodesCallerSession) IMPORTERROLE() ([32]byte, error) {
	return _Nodes.Contract.IMPORTERROLE(&_Nodes.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(bytes32 nodeId) view returns((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2]))
func (_Nodes *NodesCaller) GetNode(opts *bind.CallOpts, nodeId [32]byte) (EarthfastNode, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getNode", nodeId)

	if err != nil {
		return *new(EarthfastNode), err
	}

	out0 := *abi.ConvertType(out[0], new(EarthfastNode)).(*EarthfastNode)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(bytes32 nodeId) view returns((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2]))
func (_Nodes *NodesSession) GetNode(nodeId [32]byte) (EarthfastNode, error) {
	return _Nodes.Contract.GetNode(&_Nodes.CallOpts, nodeId)
}

// GetNode is a free data retrieval call binding the contract method 0x50c946fe.
//
// Solidity: function getNode(bytes32 nodeId) view returns((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2]))
func (_Nodes *NodesCallerSession) GetNode(nodeId [32]byte) (EarthfastNode, error) {
	return _Nodes.Contract.GetNode(&_Nodes.CallOpts, nodeId)
}

// GetNodeCount is a free data retrieval call binding the contract method 0xda0549ab.
//
// Solidity: function getNodeCount(bytes32 operatorIdOrZero) view returns(uint256 count)
func (_Nodes *NodesCaller) GetNodeCount(opts *bind.CallOpts, operatorIdOrZero [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getNodeCount", operatorIdOrZero)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeCount is a free data retrieval call binding the contract method 0xda0549ab.
//
// Solidity: function getNodeCount(bytes32 operatorIdOrZero) view returns(uint256 count)
func (_Nodes *NodesSession) GetNodeCount(operatorIdOrZero [32]byte) (*big.Int, error) {
	return _Nodes.Contract.GetNodeCount(&_Nodes.CallOpts, operatorIdOrZero)
}

// GetNodeCount is a free data retrieval call binding the contract method 0xda0549ab.
//
// Solidity: function getNodeCount(bytes32 operatorIdOrZero) view returns(uint256 count)
func (_Nodes *NodesCallerSession) GetNodeCount(operatorIdOrZero [32]byte) (*big.Int, error) {
	return _Nodes.Contract.GetNodeCount(&_Nodes.CallOpts, operatorIdOrZero)
}

// GetNodes is a free data retrieval call binding the contract method 0x79ca4a39.
//
// Solidity: function getNodes(bytes32 operatorIdOrZero, uint256 skip, uint256 size) view returns((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] values)
func (_Nodes *NodesCaller) GetNodes(opts *bind.CallOpts, operatorIdOrZero [32]byte, skip *big.Int, size *big.Int) ([]EarthfastNode, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getNodes", operatorIdOrZero, skip, size)

	if err != nil {
		return *new([]EarthfastNode), err
	}

	out0 := *abi.ConvertType(out[0], new([]EarthfastNode)).(*[]EarthfastNode)

	return out0, err

}

// GetNodes is a free data retrieval call binding the contract method 0x79ca4a39.
//
// Solidity: function getNodes(bytes32 operatorIdOrZero, uint256 skip, uint256 size) view returns((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] values)
func (_Nodes *NodesSession) GetNodes(operatorIdOrZero [32]byte, skip *big.Int, size *big.Int) ([]EarthfastNode, error) {
	return _Nodes.Contract.GetNodes(&_Nodes.CallOpts, operatorIdOrZero, skip, size)
}

// GetNodes is a free data retrieval call binding the contract method 0x79ca4a39.
//
// Solidity: function getNodes(bytes32 operatorIdOrZero, uint256 skip, uint256 size) view returns((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] values)
func (_Nodes *NodesCallerSession) GetNodes(operatorIdOrZero [32]byte, skip *big.Int, size *big.Int) ([]EarthfastNode, error) {
	return _Nodes.Contract.GetNodes(&_Nodes.CallOpts, operatorIdOrZero, skip, size)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Nodes *NodesCaller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Nodes *NodesSession) GetRegistry() (common.Address, error) {
	return _Nodes.Contract.GetRegistry(&_Nodes.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Nodes *NodesCallerSession) GetRegistry() (common.Address, error) {
	return _Nodes.Contract.GetRegistry(&_Nodes.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Nodes *NodesCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Nodes *NodesSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Nodes.Contract.GetRoleAdmin(&_Nodes.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Nodes *NodesCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Nodes.Contract.GetRoleAdmin(&_Nodes.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Nodes *NodesCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Nodes *NodesSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Nodes.Contract.HasRole(&_Nodes.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Nodes *NodesCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Nodes.Contract.HasRole(&_Nodes.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Nodes *NodesCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Nodes *NodesSession) Paused() (bool, error) {
	return _Nodes.Contract.Paused(&_Nodes.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Nodes *NodesCallerSession) Paused() (bool, error) {
	return _Nodes.Contract.Paused(&_Nodes.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Nodes *NodesCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Nodes *NodesSession) ProxiableUUID() ([32]byte, error) {
	return _Nodes.Contract.ProxiableUUID(&_Nodes.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Nodes *NodesCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Nodes.Contract.ProxiableUUID(&_Nodes.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nodes *NodesCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Nodes.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nodes *NodesSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nodes.Contract.SupportsInterface(&_Nodes.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nodes *NodesCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nodes.Contract.SupportsInterface(&_Nodes.CallOpts, interfaceId)
}

// AdvanceNodeEpochImpl is a paid mutator transaction binding the contract method 0xfea4a675.
//
// Solidity: function advanceNodeEpochImpl(bytes32 nodeId) returns()
func (_Nodes *NodesTransactor) AdvanceNodeEpochImpl(opts *bind.TransactOpts, nodeId [32]byte) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "advanceNodeEpochImpl", nodeId)
}

// AdvanceNodeEpochImpl is a paid mutator transaction binding the contract method 0xfea4a675.
//
// Solidity: function advanceNodeEpochImpl(bytes32 nodeId) returns()
func (_Nodes *NodesSession) AdvanceNodeEpochImpl(nodeId [32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.AdvanceNodeEpochImpl(&_Nodes.TransactOpts, nodeId)
}

// AdvanceNodeEpochImpl is a paid mutator transaction binding the contract method 0xfea4a675.
//
// Solidity: function advanceNodeEpochImpl(bytes32 nodeId) returns()
func (_Nodes *NodesTransactorSession) AdvanceNodeEpochImpl(nodeId [32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.AdvanceNodeEpochImpl(&_Nodes.TransactOpts, nodeId)
}

// CreateNodes is a paid mutator transaction binding the contract method 0xd204ade1.
//
// Solidity: function createNodes(bytes32 operatorId, (string,string,bool,uint256)[] nodes) returns(bytes32[] nodeIds)
func (_Nodes *NodesTransactor) CreateNodes(opts *bind.TransactOpts, operatorId [32]byte, nodes []EarthfastCreateNodeData) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "createNodes", operatorId, nodes)
}

// CreateNodes is a paid mutator transaction binding the contract method 0xd204ade1.
//
// Solidity: function createNodes(bytes32 operatorId, (string,string,bool,uint256)[] nodes) returns(bytes32[] nodeIds)
func (_Nodes *NodesSession) CreateNodes(operatorId [32]byte, nodes []EarthfastCreateNodeData) (*types.Transaction, error) {
	return _Nodes.Contract.CreateNodes(&_Nodes.TransactOpts, operatorId, nodes)
}

// CreateNodes is a paid mutator transaction binding the contract method 0xd204ade1.
//
// Solidity: function createNodes(bytes32 operatorId, (string,string,bool,uint256)[] nodes) returns(bytes32[] nodeIds)
func (_Nodes *NodesTransactorSession) CreateNodes(operatorId [32]byte, nodes []EarthfastCreateNodeData) (*types.Transaction, error) {
	return _Nodes.Contract.CreateNodes(&_Nodes.TransactOpts, operatorId, nodes)
}

// DeleteNodes is a paid mutator transaction binding the contract method 0xa6fdda15.
//
// Solidity: function deleteNodes(bytes32 operatorId, bytes32[] nodeIds) returns()
func (_Nodes *NodesTransactor) DeleteNodes(opts *bind.TransactOpts, operatorId [32]byte, nodeIds [][32]byte) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "deleteNodes", operatorId, nodeIds)
}

// DeleteNodes is a paid mutator transaction binding the contract method 0xa6fdda15.
//
// Solidity: function deleteNodes(bytes32 operatorId, bytes32[] nodeIds) returns()
func (_Nodes *NodesSession) DeleteNodes(operatorId [32]byte, nodeIds [][32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.DeleteNodes(&_Nodes.TransactOpts, operatorId, nodeIds)
}

// DeleteNodes is a paid mutator transaction binding the contract method 0xa6fdda15.
//
// Solidity: function deleteNodes(bytes32 operatorId, bytes32[] nodeIds) returns()
func (_Nodes *NodesTransactorSession) DeleteNodes(operatorId [32]byte, nodeIds [][32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.DeleteNodes(&_Nodes.TransactOpts, operatorId, nodeIds)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Nodes *NodesTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Nodes *NodesSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.GrantRole(&_Nodes.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Nodes *NodesTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.GrantRole(&_Nodes.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x7aee1c6b.
//
// Solidity: function initialize(address[] admins, address registry, bool grantImporterRole) returns()
func (_Nodes *NodesTransactor) Initialize(opts *bind.TransactOpts, admins []common.Address, registry common.Address, grantImporterRole bool) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "initialize", admins, registry, grantImporterRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x7aee1c6b.
//
// Solidity: function initialize(address[] admins, address registry, bool grantImporterRole) returns()
func (_Nodes *NodesSession) Initialize(admins []common.Address, registry common.Address, grantImporterRole bool) (*types.Transaction, error) {
	return _Nodes.Contract.Initialize(&_Nodes.TransactOpts, admins, registry, grantImporterRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x7aee1c6b.
//
// Solidity: function initialize(address[] admins, address registry, bool grantImporterRole) returns()
func (_Nodes *NodesTransactorSession) Initialize(admins []common.Address, registry common.Address, grantImporterRole bool) (*types.Transaction, error) {
	return _Nodes.Contract.Initialize(&_Nodes.TransactOpts, admins, registry, grantImporterRole)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Nodes *NodesTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Nodes *NodesSession) Pause() (*types.Transaction, error) {
	return _Nodes.Contract.Pause(&_Nodes.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Nodes *NodesTransactorSession) Pause() (*types.Transaction, error) {
	return _Nodes.Contract.Pause(&_Nodes.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Nodes *NodesTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Nodes *NodesSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.RenounceRole(&_Nodes.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Nodes *NodesTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.RenounceRole(&_Nodes.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Nodes *NodesTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Nodes *NodesSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.RevokeRole(&_Nodes.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Nodes *NodesTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.RevokeRole(&_Nodes.TransactOpts, role, account)
}

// SetNodeDisabled is a paid mutator transaction binding the contract method 0x87477979.
//
// Solidity: function setNodeDisabled(bytes32 operatorId, bytes32[] nodeIds, bool[] disabled) returns()
func (_Nodes *NodesTransactor) SetNodeDisabled(opts *bind.TransactOpts, operatorId [32]byte, nodeIds [][32]byte, disabled []bool) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "setNodeDisabled", operatorId, nodeIds, disabled)
}

// SetNodeDisabled is a paid mutator transaction binding the contract method 0x87477979.
//
// Solidity: function setNodeDisabled(bytes32 operatorId, bytes32[] nodeIds, bool[] disabled) returns()
func (_Nodes *NodesSession) SetNodeDisabled(operatorId [32]byte, nodeIds [][32]byte, disabled []bool) (*types.Transaction, error) {
	return _Nodes.Contract.SetNodeDisabled(&_Nodes.TransactOpts, operatorId, nodeIds, disabled)
}

// SetNodeDisabled is a paid mutator transaction binding the contract method 0x87477979.
//
// Solidity: function setNodeDisabled(bytes32 operatorId, bytes32[] nodeIds, bool[] disabled) returns()
func (_Nodes *NodesTransactorSession) SetNodeDisabled(operatorId [32]byte, nodeIds [][32]byte, disabled []bool) (*types.Transaction, error) {
	return _Nodes.Contract.SetNodeDisabled(&_Nodes.TransactOpts, operatorId, nodeIds, disabled)
}

// SetNodeHosts is a paid mutator transaction binding the contract method 0x471f8e3a.
//
// Solidity: function setNodeHosts(bytes32 operatorId, bytes32[] nodeIds, string[] hosts, string[] regions) returns()
func (_Nodes *NodesTransactor) SetNodeHosts(opts *bind.TransactOpts, operatorId [32]byte, nodeIds [][32]byte, hosts []string, regions []string) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "setNodeHosts", operatorId, nodeIds, hosts, regions)
}

// SetNodeHosts is a paid mutator transaction binding the contract method 0x471f8e3a.
//
// Solidity: function setNodeHosts(bytes32 operatorId, bytes32[] nodeIds, string[] hosts, string[] regions) returns()
func (_Nodes *NodesSession) SetNodeHosts(operatorId [32]byte, nodeIds [][32]byte, hosts []string, regions []string) (*types.Transaction, error) {
	return _Nodes.Contract.SetNodeHosts(&_Nodes.TransactOpts, operatorId, nodeIds, hosts, regions)
}

// SetNodeHosts is a paid mutator transaction binding the contract method 0x471f8e3a.
//
// Solidity: function setNodeHosts(bytes32 operatorId, bytes32[] nodeIds, string[] hosts, string[] regions) returns()
func (_Nodes *NodesTransactorSession) SetNodeHosts(operatorId [32]byte, nodeIds [][32]byte, hosts []string, regions []string) (*types.Transaction, error) {
	return _Nodes.Contract.SetNodeHosts(&_Nodes.TransactOpts, operatorId, nodeIds, hosts, regions)
}

// SetNodePriceImpl is a paid mutator transaction binding the contract method 0x1f9fb2bb.
//
// Solidity: function setNodePriceImpl(bytes32 nodeId, uint256 epochSlot, uint256 price) returns()
func (_Nodes *NodesTransactor) SetNodePriceImpl(opts *bind.TransactOpts, nodeId [32]byte, epochSlot *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "setNodePriceImpl", nodeId, epochSlot, price)
}

// SetNodePriceImpl is a paid mutator transaction binding the contract method 0x1f9fb2bb.
//
// Solidity: function setNodePriceImpl(bytes32 nodeId, uint256 epochSlot, uint256 price) returns()
func (_Nodes *NodesSession) SetNodePriceImpl(nodeId [32]byte, epochSlot *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Nodes.Contract.SetNodePriceImpl(&_Nodes.TransactOpts, nodeId, epochSlot, price)
}

// SetNodePriceImpl is a paid mutator transaction binding the contract method 0x1f9fb2bb.
//
// Solidity: function setNodePriceImpl(bytes32 nodeId, uint256 epochSlot, uint256 price) returns()
func (_Nodes *NodesTransactorSession) SetNodePriceImpl(nodeId [32]byte, epochSlot *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Nodes.Contract.SetNodePriceImpl(&_Nodes.TransactOpts, nodeId, epochSlot, price)
}

// SetNodePrices is a paid mutator transaction binding the contract method 0x1e074fb6.
//
// Solidity: function setNodePrices(bytes32 operatorId, bytes32[] nodeIds, uint256[] prices, (bool,bool) slot) returns()
func (_Nodes *NodesTransactor) SetNodePrices(opts *bind.TransactOpts, operatorId [32]byte, nodeIds [][32]byte, prices []*big.Int, slot EarthfastSlot) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "setNodePrices", operatorId, nodeIds, prices, slot)
}

// SetNodePrices is a paid mutator transaction binding the contract method 0x1e074fb6.
//
// Solidity: function setNodePrices(bytes32 operatorId, bytes32[] nodeIds, uint256[] prices, (bool,bool) slot) returns()
func (_Nodes *NodesSession) SetNodePrices(operatorId [32]byte, nodeIds [][32]byte, prices []*big.Int, slot EarthfastSlot) (*types.Transaction, error) {
	return _Nodes.Contract.SetNodePrices(&_Nodes.TransactOpts, operatorId, nodeIds, prices, slot)
}

// SetNodePrices is a paid mutator transaction binding the contract method 0x1e074fb6.
//
// Solidity: function setNodePrices(bytes32 operatorId, bytes32[] nodeIds, uint256[] prices, (bool,bool) slot) returns()
func (_Nodes *NodesTransactorSession) SetNodePrices(operatorId [32]byte, nodeIds [][32]byte, prices []*big.Int, slot EarthfastSlot) (*types.Transaction, error) {
	return _Nodes.Contract.SetNodePrices(&_Nodes.TransactOpts, operatorId, nodeIds, prices, slot)
}

// SetNodeProjectImpl is a paid mutator transaction binding the contract method 0xf2be8ee7.
//
// Solidity: function setNodeProjectImpl(bytes32 nodeId, uint256 epochSlot, bytes32 projectId) returns()
func (_Nodes *NodesTransactor) SetNodeProjectImpl(opts *bind.TransactOpts, nodeId [32]byte, epochSlot *big.Int, projectId [32]byte) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "setNodeProjectImpl", nodeId, epochSlot, projectId)
}

// SetNodeProjectImpl is a paid mutator transaction binding the contract method 0xf2be8ee7.
//
// Solidity: function setNodeProjectImpl(bytes32 nodeId, uint256 epochSlot, bytes32 projectId) returns()
func (_Nodes *NodesSession) SetNodeProjectImpl(nodeId [32]byte, epochSlot *big.Int, projectId [32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.SetNodeProjectImpl(&_Nodes.TransactOpts, nodeId, epochSlot, projectId)
}

// SetNodeProjectImpl is a paid mutator transaction binding the contract method 0xf2be8ee7.
//
// Solidity: function setNodeProjectImpl(bytes32 nodeId, uint256 epochSlot, bytes32 projectId) returns()
func (_Nodes *NodesTransactorSession) SetNodeProjectImpl(nodeId [32]byte, epochSlot *big.Int, projectId [32]byte) (*types.Transaction, error) {
	return _Nodes.Contract.SetNodeProjectImpl(&_Nodes.TransactOpts, nodeId, epochSlot, projectId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Nodes *NodesTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Nodes *NodesSession) Unpause() (*types.Transaction, error) {
	return _Nodes.Contract.Unpause(&_Nodes.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Nodes *NodesTransactorSession) Unpause() (*types.Transaction, error) {
	return _Nodes.Contract.Unpause(&_Nodes.TransactOpts)
}

// UnsafeImportData is a paid mutator transaction binding the contract method 0x8d6fafd7.
//
// Solidity: function unsafeImportData((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] nodes, bool revokeImporterRole) returns()
func (_Nodes *NodesTransactor) UnsafeImportData(opts *bind.TransactOpts, nodes []EarthfastNode, revokeImporterRole bool) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "unsafeImportData", nodes, revokeImporterRole)
}

// UnsafeImportData is a paid mutator transaction binding the contract method 0x8d6fafd7.
//
// Solidity: function unsafeImportData((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] nodes, bool revokeImporterRole) returns()
func (_Nodes *NodesSession) UnsafeImportData(nodes []EarthfastNode, revokeImporterRole bool) (*types.Transaction, error) {
	return _Nodes.Contract.UnsafeImportData(&_Nodes.TransactOpts, nodes, revokeImporterRole)
}

// UnsafeImportData is a paid mutator transaction binding the contract method 0x8d6fafd7.
//
// Solidity: function unsafeImportData((bytes32,bytes32,string,string,bool,uint256[2],bytes32[2])[] nodes, bool revokeImporterRole) returns()
func (_Nodes *NodesTransactorSession) UnsafeImportData(nodes []EarthfastNode, revokeImporterRole bool) (*types.Transaction, error) {
	return _Nodes.Contract.UnsafeImportData(&_Nodes.TransactOpts, nodes, revokeImporterRole)
}

// UnsafeSetPrices is a paid mutator transaction binding the contract method 0x73c3ae65.
//
// Solidity: function unsafeSetPrices(uint256 skip, uint256 size, uint256 mul, uint256 div) returns()
func (_Nodes *NodesTransactor) UnsafeSetPrices(opts *bind.TransactOpts, skip *big.Int, size *big.Int, mul *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "unsafeSetPrices", skip, size, mul, div)
}

// UnsafeSetPrices is a paid mutator transaction binding the contract method 0x73c3ae65.
//
// Solidity: function unsafeSetPrices(uint256 skip, uint256 size, uint256 mul, uint256 div) returns()
func (_Nodes *NodesSession) UnsafeSetPrices(skip *big.Int, size *big.Int, mul *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Nodes.Contract.UnsafeSetPrices(&_Nodes.TransactOpts, skip, size, mul, div)
}

// UnsafeSetPrices is a paid mutator transaction binding the contract method 0x73c3ae65.
//
// Solidity: function unsafeSetPrices(uint256 skip, uint256 size, uint256 mul, uint256 div) returns()
func (_Nodes *NodesTransactorSession) UnsafeSetPrices(skip *big.Int, size *big.Int, mul *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Nodes.Contract.UnsafeSetPrices(&_Nodes.TransactOpts, skip, size, mul, div)
}

// UnsafeSetRegistry is a paid mutator transaction binding the contract method 0xb9a2adf0.
//
// Solidity: function unsafeSetRegistry(address registry) returns()
func (_Nodes *NodesTransactor) UnsafeSetRegistry(opts *bind.TransactOpts, registry common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "unsafeSetRegistry", registry)
}

// UnsafeSetRegistry is a paid mutator transaction binding the contract method 0xb9a2adf0.
//
// Solidity: function unsafeSetRegistry(address registry) returns()
func (_Nodes *NodesSession) UnsafeSetRegistry(registry common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.UnsafeSetRegistry(&_Nodes.TransactOpts, registry)
}

// UnsafeSetRegistry is a paid mutator transaction binding the contract method 0xb9a2adf0.
//
// Solidity: function unsafeSetRegistry(address registry) returns()
func (_Nodes *NodesTransactorSession) UnsafeSetRegistry(registry common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.UnsafeSetRegistry(&_Nodes.TransactOpts, registry)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Nodes *NodesTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Nodes *NodesSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.UpgradeTo(&_Nodes.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Nodes *NodesTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Nodes.Contract.UpgradeTo(&_Nodes.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Nodes *NodesTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Nodes.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Nodes *NodesSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Nodes.Contract.UpgradeToAndCall(&_Nodes.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Nodes *NodesTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Nodes.Contract.UpgradeToAndCall(&_Nodes.TransactOpts, newImplementation, data)
}

// NodesAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Nodes contract.
type NodesAdminChangedIterator struct {
	Event *NodesAdminChanged // Event containing the contract specifics and raw log

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
func (it *NodesAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesAdminChanged)
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
		it.Event = new(NodesAdminChanged)
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
func (it *NodesAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesAdminChanged represents a AdminChanged event raised by the Nodes contract.
type NodesAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Nodes *NodesFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NodesAdminChangedIterator, error) {

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NodesAdminChangedIterator{contract: _Nodes.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Nodes *NodesFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NodesAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesAdminChanged)
				if err := _Nodes.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Nodes *NodesFilterer) ParseAdminChanged(log types.Log) (*NodesAdminChanged, error) {
	event := new(NodesAdminChanged)
	if err := _Nodes.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Nodes contract.
type NodesBeaconUpgradedIterator struct {
	Event *NodesBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NodesBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesBeaconUpgraded)
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
		it.Event = new(NodesBeaconUpgraded)
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
func (it *NodesBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesBeaconUpgraded represents a BeaconUpgraded event raised by the Nodes contract.
type NodesBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Nodes *NodesFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NodesBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NodesBeaconUpgradedIterator{contract: _Nodes.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Nodes *NodesFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NodesBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesBeaconUpgraded)
				if err := _Nodes.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Nodes *NodesFilterer) ParseBeaconUpgraded(log types.Log) (*NodesBeaconUpgraded, error) {
	event := new(NodesBeaconUpgraded)
	if err := _Nodes.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Nodes contract.
type NodesInitializedIterator struct {
	Event *NodesInitialized // Event containing the contract specifics and raw log

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
func (it *NodesInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesInitialized)
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
		it.Event = new(NodesInitialized)
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
func (it *NodesInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesInitialized represents a Initialized event raised by the Nodes contract.
type NodesInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nodes *NodesFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodesInitializedIterator, error) {

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodesInitializedIterator{contract: _Nodes.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nodes *NodesFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodesInitialized) (event.Subscription, error) {

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesInitialized)
				if err := _Nodes.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Nodes *NodesFilterer) ParseInitialized(log types.Log) (*NodesInitialized, error) {
	event := new(NodesInitialized)
	if err := _Nodes.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesNodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the Nodes contract.
type NodesNodeCreatedIterator struct {
	Event *NodesNodeCreated // Event containing the contract specifics and raw log

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
func (it *NodesNodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesNodeCreated)
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
		it.Event = new(NodesNodeCreated)
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
func (it *NodesNodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesNodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesNodeCreated represents a NodeCreated event raised by the Nodes contract.
type NodesNodeCreated struct {
	NodeId     [32]byte
	OperatorId [32]byte
	Host       string
	Region     string
	Disabled   bool
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0x85f443b741d17b6e22b10c766b89a7d244f820b4d72ba784e3c8f51c0c5ee548.
//
// Solidity: event NodeCreated(bytes32 indexed nodeId, bytes32 indexed operatorId, string host, string region, bool disabled, uint256 price)
func (_Nodes *NodesFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeId [][32]byte, operatorId [][32]byte) (*NodesNodeCreatedIterator, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "NodeCreated", nodeIdRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodesNodeCreatedIterator{contract: _Nodes.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0x85f443b741d17b6e22b10c766b89a7d244f820b4d72ba784e3c8f51c0c5ee548.
//
// Solidity: event NodeCreated(bytes32 indexed nodeId, bytes32 indexed operatorId, string host, string region, bool disabled, uint256 price)
func (_Nodes *NodesFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *NodesNodeCreated, nodeId [][32]byte, operatorId [][32]byte) (event.Subscription, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "NodeCreated", nodeIdRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesNodeCreated)
				if err := _Nodes.contract.UnpackLog(event, "NodeCreated", log); err != nil {
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

// ParseNodeCreated is a log parse operation binding the contract event 0x85f443b741d17b6e22b10c766b89a7d244f820b4d72ba784e3c8f51c0c5ee548.
//
// Solidity: event NodeCreated(bytes32 indexed nodeId, bytes32 indexed operatorId, string host, string region, bool disabled, uint256 price)
func (_Nodes *NodesFilterer) ParseNodeCreated(log types.Log) (*NodesNodeCreated, error) {
	event := new(NodesNodeCreated)
	if err := _Nodes.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesNodeDeletedIterator is returned from FilterNodeDeleted and is used to iterate over the raw logs and unpacked data for NodeDeleted events raised by the Nodes contract.
type NodesNodeDeletedIterator struct {
	Event *NodesNodeDeleted // Event containing the contract specifics and raw log

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
func (it *NodesNodeDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesNodeDeleted)
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
		it.Event = new(NodesNodeDeleted)
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
func (it *NodesNodeDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesNodeDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesNodeDeleted represents a NodeDeleted event raised by the Nodes contract.
type NodesNodeDeleted struct {
	NodeId     [32]byte
	OperatorId [32]byte
	Host       string
	Region     string
	Disabled   bool
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeDeleted is a free log retrieval operation binding the contract event 0x6fded3d147f7c9b713583f6260c2b67c18fd11587c272b135a8c438e8a288f2f.
//
// Solidity: event NodeDeleted(bytes32 indexed nodeId, bytes32 indexed operatorId, string host, string region, bool disabled, uint256 price)
func (_Nodes *NodesFilterer) FilterNodeDeleted(opts *bind.FilterOpts, nodeId [][32]byte, operatorId [][32]byte) (*NodesNodeDeletedIterator, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "NodeDeleted", nodeIdRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &NodesNodeDeletedIterator{contract: _Nodes.contract, event: "NodeDeleted", logs: logs, sub: sub}, nil
}

// WatchNodeDeleted is a free log subscription operation binding the contract event 0x6fded3d147f7c9b713583f6260c2b67c18fd11587c272b135a8c438e8a288f2f.
//
// Solidity: event NodeDeleted(bytes32 indexed nodeId, bytes32 indexed operatorId, string host, string region, bool disabled, uint256 price)
func (_Nodes *NodesFilterer) WatchNodeDeleted(opts *bind.WatchOpts, sink chan<- *NodesNodeDeleted, nodeId [][32]byte, operatorId [][32]byte) (event.Subscription, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "NodeDeleted", nodeIdRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesNodeDeleted)
				if err := _Nodes.contract.UnpackLog(event, "NodeDeleted", log); err != nil {
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

// ParseNodeDeleted is a log parse operation binding the contract event 0x6fded3d147f7c9b713583f6260c2b67c18fd11587c272b135a8c438e8a288f2f.
//
// Solidity: event NodeDeleted(bytes32 indexed nodeId, bytes32 indexed operatorId, string host, string region, bool disabled, uint256 price)
func (_Nodes *NodesFilterer) ParseNodeDeleted(log types.Log) (*NodesNodeDeleted, error) {
	event := new(NodesNodeDeleted)
	if err := _Nodes.contract.UnpackLog(event, "NodeDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesNodeDisabledChangedIterator is returned from FilterNodeDisabledChanged and is used to iterate over the raw logs and unpacked data for NodeDisabledChanged events raised by the Nodes contract.
type NodesNodeDisabledChangedIterator struct {
	Event *NodesNodeDisabledChanged // Event containing the contract specifics and raw log

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
func (it *NodesNodeDisabledChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesNodeDisabledChanged)
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
		it.Event = new(NodesNodeDisabledChanged)
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
func (it *NodesNodeDisabledChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesNodeDisabledChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesNodeDisabledChanged represents a NodeDisabledChanged event raised by the Nodes contract.
type NodesNodeDisabledChanged struct {
	NodeId      [32]byte
	OldDisabled bool
	NewDisabled bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeDisabledChanged is a free log retrieval operation binding the contract event 0xbd8bc64d24321dcae6059f4a6108793a3c9613c0964637b8f811de2d3b2affc2.
//
// Solidity: event NodeDisabledChanged(bytes32 indexed nodeId, bool oldDisabled, bool newDisabled)
func (_Nodes *NodesFilterer) FilterNodeDisabledChanged(opts *bind.FilterOpts, nodeId [][32]byte) (*NodesNodeDisabledChangedIterator, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "NodeDisabledChanged", nodeIdRule)
	if err != nil {
		return nil, err
	}
	return &NodesNodeDisabledChangedIterator{contract: _Nodes.contract, event: "NodeDisabledChanged", logs: logs, sub: sub}, nil
}

// WatchNodeDisabledChanged is a free log subscription operation binding the contract event 0xbd8bc64d24321dcae6059f4a6108793a3c9613c0964637b8f811de2d3b2affc2.
//
// Solidity: event NodeDisabledChanged(bytes32 indexed nodeId, bool oldDisabled, bool newDisabled)
func (_Nodes *NodesFilterer) WatchNodeDisabledChanged(opts *bind.WatchOpts, sink chan<- *NodesNodeDisabledChanged, nodeId [][32]byte) (event.Subscription, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "NodeDisabledChanged", nodeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesNodeDisabledChanged)
				if err := _Nodes.contract.UnpackLog(event, "NodeDisabledChanged", log); err != nil {
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

// ParseNodeDisabledChanged is a log parse operation binding the contract event 0xbd8bc64d24321dcae6059f4a6108793a3c9613c0964637b8f811de2d3b2affc2.
//
// Solidity: event NodeDisabledChanged(bytes32 indexed nodeId, bool oldDisabled, bool newDisabled)
func (_Nodes *NodesFilterer) ParseNodeDisabledChanged(log types.Log) (*NodesNodeDisabledChanged, error) {
	event := new(NodesNodeDisabledChanged)
	if err := _Nodes.contract.UnpackLog(event, "NodeDisabledChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesNodeHostChangedIterator is returned from FilterNodeHostChanged and is used to iterate over the raw logs and unpacked data for NodeHostChanged events raised by the Nodes contract.
type NodesNodeHostChangedIterator struct {
	Event *NodesNodeHostChanged // Event containing the contract specifics and raw log

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
func (it *NodesNodeHostChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesNodeHostChanged)
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
		it.Event = new(NodesNodeHostChanged)
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
func (it *NodesNodeHostChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesNodeHostChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesNodeHostChanged represents a NodeHostChanged event raised by the Nodes contract.
type NodesNodeHostChanged struct {
	NodeId    [32]byte
	OldHost   string
	OldRegion string
	NewHost   string
	NewRegion string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodeHostChanged is a free log retrieval operation binding the contract event 0xf939895421dc7faad365f1a6c3add39f1dbe41938423d98182e2df720718c8b2.
//
// Solidity: event NodeHostChanged(bytes32 indexed nodeId, string oldHost, string oldRegion, string newHost, string newRegion)
func (_Nodes *NodesFilterer) FilterNodeHostChanged(opts *bind.FilterOpts, nodeId [][32]byte) (*NodesNodeHostChangedIterator, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "NodeHostChanged", nodeIdRule)
	if err != nil {
		return nil, err
	}
	return &NodesNodeHostChangedIterator{contract: _Nodes.contract, event: "NodeHostChanged", logs: logs, sub: sub}, nil
}

// WatchNodeHostChanged is a free log subscription operation binding the contract event 0xf939895421dc7faad365f1a6c3add39f1dbe41938423d98182e2df720718c8b2.
//
// Solidity: event NodeHostChanged(bytes32 indexed nodeId, string oldHost, string oldRegion, string newHost, string newRegion)
func (_Nodes *NodesFilterer) WatchNodeHostChanged(opts *bind.WatchOpts, sink chan<- *NodesNodeHostChanged, nodeId [][32]byte) (event.Subscription, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "NodeHostChanged", nodeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesNodeHostChanged)
				if err := _Nodes.contract.UnpackLog(event, "NodeHostChanged", log); err != nil {
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

// ParseNodeHostChanged is a log parse operation binding the contract event 0xf939895421dc7faad365f1a6c3add39f1dbe41938423d98182e2df720718c8b2.
//
// Solidity: event NodeHostChanged(bytes32 indexed nodeId, string oldHost, string oldRegion, string newHost, string newRegion)
func (_Nodes *NodesFilterer) ParseNodeHostChanged(log types.Log) (*NodesNodeHostChanged, error) {
	event := new(NodesNodeHostChanged)
	if err := _Nodes.contract.UnpackLog(event, "NodeHostChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesNodePriceChangedIterator is returned from FilterNodePriceChanged and is used to iterate over the raw logs and unpacked data for NodePriceChanged events raised by the Nodes contract.
type NodesNodePriceChangedIterator struct {
	Event *NodesNodePriceChanged // Event containing the contract specifics and raw log

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
func (it *NodesNodePriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesNodePriceChanged)
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
		it.Event = new(NodesNodePriceChanged)
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
func (it *NodesNodePriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesNodePriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesNodePriceChanged represents a NodePriceChanged event raised by the Nodes contract.
type NodesNodePriceChanged struct {
	NodeId       [32]byte
	OldLastPrice *big.Int
	OldNextPrice *big.Int
	NewPrice     *big.Int
	Slot         EarthfastSlot
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNodePriceChanged is a free log retrieval operation binding the contract event 0x055b91d27a22fa7278a8976495eb88088023f4b0f8862aa797c7ceeb61f2eba3.
//
// Solidity: event NodePriceChanged(bytes32 indexed nodeId, uint256 oldLastPrice, uint256 oldNextPrice, uint256 newPrice, (bool,bool) slot)
func (_Nodes *NodesFilterer) FilterNodePriceChanged(opts *bind.FilterOpts, nodeId [][32]byte) (*NodesNodePriceChangedIterator, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "NodePriceChanged", nodeIdRule)
	if err != nil {
		return nil, err
	}
	return &NodesNodePriceChangedIterator{contract: _Nodes.contract, event: "NodePriceChanged", logs: logs, sub: sub}, nil
}

// WatchNodePriceChanged is a free log subscription operation binding the contract event 0x055b91d27a22fa7278a8976495eb88088023f4b0f8862aa797c7ceeb61f2eba3.
//
// Solidity: event NodePriceChanged(bytes32 indexed nodeId, uint256 oldLastPrice, uint256 oldNextPrice, uint256 newPrice, (bool,bool) slot)
func (_Nodes *NodesFilterer) WatchNodePriceChanged(opts *bind.WatchOpts, sink chan<- *NodesNodePriceChanged, nodeId [][32]byte) (event.Subscription, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "NodePriceChanged", nodeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesNodePriceChanged)
				if err := _Nodes.contract.UnpackLog(event, "NodePriceChanged", log); err != nil {
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

// ParseNodePriceChanged is a log parse operation binding the contract event 0x055b91d27a22fa7278a8976495eb88088023f4b0f8862aa797c7ceeb61f2eba3.
//
// Solidity: event NodePriceChanged(bytes32 indexed nodeId, uint256 oldLastPrice, uint256 oldNextPrice, uint256 newPrice, (bool,bool) slot)
func (_Nodes *NodesFilterer) ParseNodePriceChanged(log types.Log) (*NodesNodePriceChanged, error) {
	event := new(NodesNodePriceChanged)
	if err := _Nodes.contract.UnpackLog(event, "NodePriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Nodes contract.
type NodesPausedIterator struct {
	Event *NodesPaused // Event containing the contract specifics and raw log

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
func (it *NodesPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesPaused)
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
		it.Event = new(NodesPaused)
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
func (it *NodesPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesPaused represents a Paused event raised by the Nodes contract.
type NodesPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Nodes *NodesFilterer) FilterPaused(opts *bind.FilterOpts) (*NodesPausedIterator, error) {

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NodesPausedIterator{contract: _Nodes.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Nodes *NodesFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NodesPaused) (event.Subscription, error) {

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesPaused)
				if err := _Nodes.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Nodes *NodesFilterer) ParsePaused(log types.Log) (*NodesPaused, error) {
	event := new(NodesPaused)
	if err := _Nodes.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Nodes contract.
type NodesRoleAdminChangedIterator struct {
	Event *NodesRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NodesRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRoleAdminChanged)
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
		it.Event = new(NodesRoleAdminChanged)
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
func (it *NodesRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRoleAdminChanged represents a RoleAdminChanged event raised by the Nodes contract.
type NodesRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Nodes *NodesFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NodesRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NodesRoleAdminChangedIterator{contract: _Nodes.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Nodes *NodesFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NodesRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRoleAdminChanged)
				if err := _Nodes.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Nodes *NodesFilterer) ParseRoleAdminChanged(log types.Log) (*NodesRoleAdminChanged, error) {
	event := new(NodesRoleAdminChanged)
	if err := _Nodes.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Nodes contract.
type NodesRoleGrantedIterator struct {
	Event *NodesRoleGranted // Event containing the contract specifics and raw log

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
func (it *NodesRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRoleGranted)
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
		it.Event = new(NodesRoleGranted)
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
func (it *NodesRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRoleGranted represents a RoleGranted event raised by the Nodes contract.
type NodesRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Nodes *NodesFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NodesRoleGrantedIterator, error) {

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

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NodesRoleGrantedIterator{contract: _Nodes.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Nodes *NodesFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NodesRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRoleGranted)
				if err := _Nodes.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Nodes *NodesFilterer) ParseRoleGranted(log types.Log) (*NodesRoleGranted, error) {
	event := new(NodesRoleGranted)
	if err := _Nodes.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Nodes contract.
type NodesRoleRevokedIterator struct {
	Event *NodesRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NodesRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRoleRevoked)
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
		it.Event = new(NodesRoleRevoked)
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
func (it *NodesRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRoleRevoked represents a RoleRevoked event raised by the Nodes contract.
type NodesRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Nodes *NodesFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NodesRoleRevokedIterator, error) {

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

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NodesRoleRevokedIterator{contract: _Nodes.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Nodes *NodesFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NodesRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRoleRevoked)
				if err := _Nodes.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Nodes *NodesFilterer) ParseRoleRevoked(log types.Log) (*NodesRoleRevoked, error) {
	event := new(NodesRoleRevoked)
	if err := _Nodes.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Nodes contract.
type NodesUnpausedIterator struct {
	Event *NodesUnpaused // Event containing the contract specifics and raw log

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
func (it *NodesUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesUnpaused)
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
		it.Event = new(NodesUnpaused)
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
func (it *NodesUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesUnpaused represents a Unpaused event raised by the Nodes contract.
type NodesUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Nodes *NodesFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NodesUnpausedIterator, error) {

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NodesUnpausedIterator{contract: _Nodes.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Nodes *NodesFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NodesUnpaused) (event.Subscription, error) {

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesUnpaused)
				if err := _Nodes.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Nodes *NodesFilterer) ParseUnpaused(log types.Log) (*NodesUnpaused, error) {
	event := new(NodesUnpaused)
	if err := _Nodes.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Nodes contract.
type NodesUpgradedIterator struct {
	Event *NodesUpgraded // Event containing the contract specifics and raw log

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
func (it *NodesUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesUpgraded)
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
		it.Event = new(NodesUpgraded)
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
func (it *NodesUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesUpgraded represents a Upgraded event raised by the Nodes contract.
type NodesUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Nodes *NodesFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NodesUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Nodes.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NodesUpgradedIterator{contract: _Nodes.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Nodes *NodesFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NodesUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Nodes.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesUpgraded)
				if err := _Nodes.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Nodes *NodesFilterer) ParseUpgraded(log types.Log) (*NodesUpgraded, error) {
	event := new(NodesUpgraded)
	if err := _Nodes.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

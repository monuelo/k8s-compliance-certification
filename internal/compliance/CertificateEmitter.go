// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compliance

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

// ComplianceMetaData contains all meta data concerning the Compliance contract.
var ComplianceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"certificateHash\",\"type\":\"bytes32\"}],\"name\":\"CertificateGenerated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"certificates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"certificateHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"name\":\"emitCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"getCertificate\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610eeb8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610079575f3560e01c80637ebd1b30116100585780637ebd1b30146100e35780638ab1d6811461010e578063e43252d714610121578063fd531e931461017f575f80fd5b806216e5261461007d5780633af32abf146100ab5780635d50bc04146100ce575b5f80fd5b61009061008b3660046109e6565b610192565b6040516100a296959493929190610a53565b60405180910390f35b6100be6100b93660046109e6565b6103dd565b60405190151581526020016100a2565b6100e16100dc366004610b52565b610443565b005b6100f66100f1366004610c08565b6105e3565b6040516001600160a01b0390911681526020016100a2565b6100e161011c3660046109e6565b61060b565b6100e161012f3660046109e6565b6001805480820182555f919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0392909216919091179055565b61009061018d3660046109e6565b61070d565b5f602081905290815260409020805481906101ac90610c1f565b80601f01602080910402602001604051908101604052809291908181526020018280546101d890610c1f565b80156102235780601f106101fa57610100808354040283529160200191610223565b820191905f5260205f20905b81548152906001019060200180831161020657829003601f168201915b50505050509080600101805461023890610c1f565b80601f016020809104026020016040519081016040528092919081815260200182805461026490610c1f565b80156102af5780601f10610286576101008083540402835291602001916102af565b820191905f5260205f20905b81548152906001019060200180831161029257829003601f168201915b5050505050908060020180546102c490610c1f565b80601f01602080910402602001604051908101604052809291908181526020018280546102f090610c1f565b801561033b5780601f106103125761010080835404028352916020019161033b565b820191905f5260205f20905b81548152906001019060200180831161031e57829003601f168201915b50505050509080600301805461035090610c1f565b80601f016020809104026020016040519081016040528092919081815260200182805461037c90610c1f565b80156103c75780601f1061039e576101008083540402835291602001916103c7565b820191905f5260205f20905b8154815290600101906020018083116103aa57829003601f168201915b5050505050908060040154908060050154905086565b5f805b60015481101561043b57826001600160a01b03166001828154811061040757610407610c57565b5f918252602090912001546001600160a01b0316036104295750600192915050565b8061043381610c7f565b9150506103e0565b505f92915050565b61044c336103dd565b6104ba5760405162461bcd60e51b815260206004820152603560248201527f4f6e6c792077686974656c6973746564206164647265737365732063616e206760448201527432b732b930ba329031b2b93a34b334b1b0ba32b99760591b606482015260840160405180910390fd5b5f6040518060c0016040528086815260200185815260200184815260200183815260200142815260200186868686428c6040516020016104ff96959493929190610c97565b60408051601f1981840301815291815281516020928301209092526001600160a01b0389165f90815290819052208151919250829181906105409082610d58565b50602082015160018201906105559082610d58565b506040820151600282019061056a9082610d58565b506060820151600382019061057f9082610d58565b506080820151600482015560a0918201516005909101558101516040517fc3003285d7eb1adf5519eca26085ea2121ccb24e407251469b5761a42b0e72c8916105d391899189918991899189914291610e14565b60405180910390a1505050505050565b600181815481106105f2575f80fd5b5f918252602090912001546001600160a01b0316905081565b5f5b60015481101561070957816001600160a01b03166001828154811061063457610634610c57565b5f918252602090912001546001600160a01b0316036106f7576001805461065c908290610e88565b8154811061066c5761066c610c57565b5f91825260209091200154600180546001600160a01b03909216918390811061069757610697610c57565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060018054806106d3576106d3610ea1565b5f8281526020902081015f1990810180546001600160a01b03191690550190555050565b8061070181610c7f565b91505061060d565b5050565b6060806060805f805f805f896001600160a01b03166001600160a01b031681526020019081526020015f206040518060c00160405290815f8201805461075290610c1f565b80601f016020809104026020016040519081016040528092919081815260200182805461077e90610c1f565b80156107c95780601f106107a0576101008083540402835291602001916107c9565b820191905f5260205f20905b8154815290600101906020018083116107ac57829003601f168201915b505050505081526020016001820180546107e290610c1f565b80601f016020809104026020016040519081016040528092919081815260200182805461080e90610c1f565b80156108595780601f1061083057610100808354040283529160200191610859565b820191905f5260205f20905b81548152906001019060200180831161083c57829003601f168201915b5050505050815260200160028201805461087290610c1f565b80601f016020809104026020016040519081016040528092919081815260200182805461089e90610c1f565b80156108e95780601f106108c0576101008083540402835291602001916108e9565b820191905f5260205f20905b8154815290600101906020018083116108cc57829003601f168201915b5050505050815260200160038201805461090290610c1f565b80601f016020809104026020016040519081016040528092919081815260200182805461092e90610c1f565b80156109795780601f1061095057610100808354040283529160200191610979565b820191905f5260205f20905b81548152906001019060200180831161095c57829003601f168201915b50505050508152602001600482015481526020016005820154815250509050805f015181602001518260400151836060015184608001518560a001519650965096509650965096505091939550919395565b80356001600160a01b03811681146109e1575f80fd5b919050565b5f602082840312156109f6575f80fd5b6109ff826109cb565b9392505050565b5f5b83811015610a20578181015183820152602001610a08565b50505f910152565b5f8151808452610a3f816020860160208601610a06565b601f01601f19169290920160200192915050565b60c081525f610a6560c0830189610a28565b8281036020840152610a778189610a28565b90508281036040840152610a8b8188610a28565b90508281036060840152610a9f8187610a28565b6080840195909552505060a00152949350505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610ad8575f80fd5b813567ffffffffffffffff80821115610af357610af3610ab5565b604051601f8301601f19908116603f01168101908282118183101715610b1b57610b1b610ab5565b81604052838152866020858801011115610b33575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f805f60a08688031215610b66575f80fd5b610b6f866109cb565b9450602086013567ffffffffffffffff80821115610b8b575f80fd5b610b9789838a01610ac9565b95506040880135915080821115610bac575f80fd5b610bb889838a01610ac9565b94506060880135915080821115610bcd575f80fd5b610bd989838a01610ac9565b93506080880135915080821115610bee575f80fd5b50610bfb88828901610ac9565b9150509295509295909350565b5f60208284031215610c18575f80fd5b5035919050565b600181811c90821680610c3357607f821691505b602082108103610c5157634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f60018201610c9057610c90610c6b565b5060010190565b5f8751610ca8818460208c01610a06565b875190830190610cbc818360208c01610a06565b8751910190610ccf818360208b01610a06565b8651910190610ce2818360208a01610a06565b01938452505060601b6bffffffffffffffffffffffff19166020820152603401949350505050565b601f821115610d53575f81815260208120601f850160051c81016020861015610d305750805b601f850160051c820191505b81811015610d4f57828155600101610d3c565b5050505b505050565b815167ffffffffffffffff811115610d7257610d72610ab5565b610d8681610d808454610c1f565b84610d0a565b602080601f831160018114610db9575f8415610da25750858301515b5f19600386901b1c1916600185901b178555610d4f565b5f85815260208120601f198616915b82811015610de757888601518255948401946001909101908401610dc8565b5085821015610e0457878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b038816815260e0602082018190525f90610e3790830189610a28565b8281036040840152610e498189610a28565b90508281036060840152610e5d8188610a28565b90508281036080840152610e718187610a28565b60a0840195909552505060c0015295945050505050565b81810381811115610e9b57610e9b610c6b565b92915050565b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220f567e26829cf1f8c3aeeb991cba7104a178fd3c30ce2c028e2430eb3c7e6d76c64736f6c63430008140033",
}

// ComplianceABI is the input ABI used to generate the binding from.
// Deprecated: Use ComplianceMetaData.ABI instead.
var ComplianceABI = ComplianceMetaData.ABI

// ComplianceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ComplianceMetaData.Bin instead.
var ComplianceBin = ComplianceMetaData.Bin

// DeployCompliance deploys a new Ethereum contract, binding an instance of Compliance to it.
func DeployCompliance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Compliance, error) {
	parsed, err := ComplianceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ComplianceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Compliance{ComplianceCaller: ComplianceCaller{contract: contract}, ComplianceTransactor: ComplianceTransactor{contract: contract}, ComplianceFilterer: ComplianceFilterer{contract: contract}}, nil
}

// Compliance is an auto generated Go binding around an Ethereum contract.
type Compliance struct {
	ComplianceCaller     // Read-only binding to the contract
	ComplianceTransactor // Write-only binding to the contract
	ComplianceFilterer   // Log filterer for contract events
}

// ComplianceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComplianceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComplianceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComplianceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComplianceSession struct {
	Contract     *Compliance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComplianceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComplianceCallerSession struct {
	Contract *ComplianceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ComplianceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComplianceTransactorSession struct {
	Contract     *ComplianceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ComplianceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComplianceRaw struct {
	Contract *Compliance // Generic contract binding to access the raw methods on
}

// ComplianceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComplianceCallerRaw struct {
	Contract *ComplianceCaller // Generic read-only contract binding to access the raw methods on
}

// ComplianceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComplianceTransactorRaw struct {
	Contract *ComplianceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompliance creates a new instance of Compliance, bound to a specific deployed contract.
func NewCompliance(address common.Address, backend bind.ContractBackend) (*Compliance, error) {
	contract, err := bindCompliance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compliance{ComplianceCaller: ComplianceCaller{contract: contract}, ComplianceTransactor: ComplianceTransactor{contract: contract}, ComplianceFilterer: ComplianceFilterer{contract: contract}}, nil
}

// NewComplianceCaller creates a new read-only instance of Compliance, bound to a specific deployed contract.
func NewComplianceCaller(address common.Address, caller bind.ContractCaller) (*ComplianceCaller, error) {
	contract, err := bindCompliance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComplianceCaller{contract: contract}, nil
}

// NewComplianceTransactor creates a new write-only instance of Compliance, bound to a specific deployed contract.
func NewComplianceTransactor(address common.Address, transactor bind.ContractTransactor) (*ComplianceTransactor, error) {
	contract, err := bindCompliance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComplianceTransactor{contract: contract}, nil
}

// NewComplianceFilterer creates a new log filterer instance of Compliance, bound to a specific deployed contract.
func NewComplianceFilterer(address common.Address, filterer bind.ContractFilterer) (*ComplianceFilterer, error) {
	contract, err := bindCompliance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComplianceFilterer{contract: contract}, nil
}

// bindCompliance binds a generic wrapper to an already deployed contract.
func bindCompliance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ComplianceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compliance *ComplianceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compliance.Contract.ComplianceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compliance *ComplianceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compliance.Contract.ComplianceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compliance *ComplianceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compliance.Contract.ComplianceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compliance *ComplianceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compliance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compliance *ComplianceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compliance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compliance *ComplianceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compliance.Contract.contract.Transact(opts, method, params...)
}

// Certificates is a free data retrieval call binding the contract method 0x0016e526.
//
// Solidity: function certificates(address ) view returns(string name, string date, string issuer, string status, uint256 timestamp, bytes32 certificateHash)
func (_Compliance *ComplianceCaller) Certificates(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name            string
	Date            string
	Issuer          string
	Status          string
	Timestamp       *big.Int
	CertificateHash [32]byte
}, error) {
	var out []interface{}
	err := _Compliance.contract.Call(opts, &out, "certificates", arg0)

	outstruct := new(struct {
		Name            string
		Date            string
		Issuer          string
		Status          string
		Timestamp       *big.Int
		CertificateHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Date = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Issuer = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CertificateHash = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Certificates is a free data retrieval call binding the contract method 0x0016e526.
//
// Solidity: function certificates(address ) view returns(string name, string date, string issuer, string status, uint256 timestamp, bytes32 certificateHash)
func (_Compliance *ComplianceSession) Certificates(arg0 common.Address) (struct {
	Name            string
	Date            string
	Issuer          string
	Status          string
	Timestamp       *big.Int
	CertificateHash [32]byte
}, error) {
	return _Compliance.Contract.Certificates(&_Compliance.CallOpts, arg0)
}

// Certificates is a free data retrieval call binding the contract method 0x0016e526.
//
// Solidity: function certificates(address ) view returns(string name, string date, string issuer, string status, uint256 timestamp, bytes32 certificateHash)
func (_Compliance *ComplianceCallerSession) Certificates(arg0 common.Address) (struct {
	Name            string
	Date            string
	Issuer          string
	Status          string
	Timestamp       *big.Int
	CertificateHash [32]byte
}, error) {
	return _Compliance.Contract.Certificates(&_Compliance.CallOpts, arg0)
}

// GetCertificate is a free data retrieval call binding the contract method 0xfd531e93.
//
// Solidity: function getCertificate(address recipient) view returns(string, string, string, string, uint256, bytes32)
func (_Compliance *ComplianceCaller) GetCertificate(opts *bind.CallOpts, recipient common.Address) (string, string, string, string, *big.Int, [32]byte, error) {
	var out []interface{}
	err := _Compliance.contract.Call(opts, &out, "getCertificate", recipient)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(string), *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return out0, out1, out2, out3, out4, out5, err

}

// GetCertificate is a free data retrieval call binding the contract method 0xfd531e93.
//
// Solidity: function getCertificate(address recipient) view returns(string, string, string, string, uint256, bytes32)
func (_Compliance *ComplianceSession) GetCertificate(recipient common.Address) (string, string, string, string, *big.Int, [32]byte, error) {
	return _Compliance.Contract.GetCertificate(&_Compliance.CallOpts, recipient)
}

// GetCertificate is a free data retrieval call binding the contract method 0xfd531e93.
//
// Solidity: function getCertificate(address recipient) view returns(string, string, string, string, uint256, bytes32)
func (_Compliance *ComplianceCallerSession) GetCertificate(recipient common.Address) (string, string, string, string, *big.Int, [32]byte, error) {
	return _Compliance.Contract.GetCertificate(&_Compliance.CallOpts, recipient)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Compliance *ComplianceCaller) IsWhitelisted(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Compliance.contract.Call(opts, &out, "isWhitelisted", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Compliance *ComplianceSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Compliance.Contract.IsWhitelisted(&_Compliance.CallOpts, _address)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Compliance *ComplianceCallerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Compliance.Contract.IsWhitelisted(&_Compliance.CallOpts, _address)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) view returns(address)
func (_Compliance *ComplianceCaller) Whitelist(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Compliance.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) view returns(address)
func (_Compliance *ComplianceSession) Whitelist(arg0 *big.Int) (common.Address, error) {
	return _Compliance.Contract.Whitelist(&_Compliance.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) view returns(address)
func (_Compliance *ComplianceCallerSession) Whitelist(arg0 *big.Int) (common.Address, error) {
	return _Compliance.Contract.Whitelist(&_Compliance.CallOpts, arg0)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _address) returns()
func (_Compliance *ComplianceTransactor) AddToWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Compliance.contract.Transact(opts, "addToWhitelist", _address)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _address) returns()
func (_Compliance *ComplianceSession) AddToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Compliance.Contract.AddToWhitelist(&_Compliance.TransactOpts, _address)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _address) returns()
func (_Compliance *ComplianceTransactorSession) AddToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Compliance.Contract.AddToWhitelist(&_Compliance.TransactOpts, _address)
}

// EmitCertificate is a paid mutator transaction binding the contract method 0x5d50bc04.
//
// Solidity: function emitCertificate(address recipient, string name, string date, string issuer, string status) returns()
func (_Compliance *ComplianceTransactor) EmitCertificate(opts *bind.TransactOpts, recipient common.Address, name string, date string, issuer string, status string) (*types.Transaction, error) {
	return _Compliance.contract.Transact(opts, "emitCertificate", recipient, name, date, issuer, status)
}

// EmitCertificate is a paid mutator transaction binding the contract method 0x5d50bc04.
//
// Solidity: function emitCertificate(address recipient, string name, string date, string issuer, string status) returns()
func (_Compliance *ComplianceSession) EmitCertificate(recipient common.Address, name string, date string, issuer string, status string) (*types.Transaction, error) {
	return _Compliance.Contract.EmitCertificate(&_Compliance.TransactOpts, recipient, name, date, issuer, status)
}

// EmitCertificate is a paid mutator transaction binding the contract method 0x5d50bc04.
//
// Solidity: function emitCertificate(address recipient, string name, string date, string issuer, string status) returns()
func (_Compliance *ComplianceTransactorSession) EmitCertificate(recipient common.Address, name string, date string, issuer string, status string) (*types.Transaction, error) {
	return _Compliance.Contract.EmitCertificate(&_Compliance.TransactOpts, recipient, name, date, issuer, status)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _address) returns()
func (_Compliance *ComplianceTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Compliance.contract.Transact(opts, "removeFromWhitelist", _address)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _address) returns()
func (_Compliance *ComplianceSession) RemoveFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Compliance.Contract.RemoveFromWhitelist(&_Compliance.TransactOpts, _address)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _address) returns()
func (_Compliance *ComplianceTransactorSession) RemoveFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Compliance.Contract.RemoveFromWhitelist(&_Compliance.TransactOpts, _address)
}

// ComplianceCertificateGeneratedIterator is returned from FilterCertificateGenerated and is used to iterate over the raw logs and unpacked data for CertificateGenerated events raised by the Compliance contract.
type ComplianceCertificateGeneratedIterator struct {
	Event *ComplianceCertificateGenerated // Event containing the contract specifics and raw log

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
func (it *ComplianceCertificateGeneratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComplianceCertificateGenerated)
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
		it.Event = new(ComplianceCertificateGenerated)
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
func (it *ComplianceCertificateGeneratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComplianceCertificateGeneratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComplianceCertificateGenerated represents a CertificateGenerated event raised by the Compliance contract.
type ComplianceCertificateGenerated struct {
	Recipient       common.Address
	Name            string
	Date            string
	Issuer          string
	Status          string
	Timestamp       *big.Int
	CertificateHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCertificateGenerated is a free log retrieval operation binding the contract event 0xc3003285d7eb1adf5519eca26085ea2121ccb24e407251469b5761a42b0e72c8.
//
// Solidity: event CertificateGenerated(address recipient, string name, string date, string issuer, string status, uint256 timestamp, bytes32 certificateHash)
func (_Compliance *ComplianceFilterer) FilterCertificateGenerated(opts *bind.FilterOpts) (*ComplianceCertificateGeneratedIterator, error) {

	logs, sub, err := _Compliance.contract.FilterLogs(opts, "CertificateGenerated")
	if err != nil {
		return nil, err
	}
	return &ComplianceCertificateGeneratedIterator{contract: _Compliance.contract, event: "CertificateGenerated", logs: logs, sub: sub}, nil
}

// WatchCertificateGenerated is a free log subscription operation binding the contract event 0xc3003285d7eb1adf5519eca26085ea2121ccb24e407251469b5761a42b0e72c8.
//
// Solidity: event CertificateGenerated(address recipient, string name, string date, string issuer, string status, uint256 timestamp, bytes32 certificateHash)
func (_Compliance *ComplianceFilterer) WatchCertificateGenerated(opts *bind.WatchOpts, sink chan<- *ComplianceCertificateGenerated) (event.Subscription, error) {

	logs, sub, err := _Compliance.contract.WatchLogs(opts, "CertificateGenerated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComplianceCertificateGenerated)
				if err := _Compliance.contract.UnpackLog(event, "CertificateGenerated", log); err != nil {
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

// ParseCertificateGenerated is a log parse operation binding the contract event 0xc3003285d7eb1adf5519eca26085ea2121ccb24e407251469b5761a42b0e72c8.
//
// Solidity: event CertificateGenerated(address recipient, string name, string date, string issuer, string status, uint256 timestamp, bytes32 certificateHash)
func (_Compliance *ComplianceFilterer) ParseCertificateGenerated(log types.Log) (*ComplianceCertificateGenerated, error) {
	event := new(ComplianceCertificateGenerated)
	if err := _Compliance.contract.UnpackLog(event, "CertificateGenerated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

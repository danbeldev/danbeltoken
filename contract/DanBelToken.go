// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"investor1\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getReferralCode\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"enumDanBelToken.UserRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"referalCode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"enumDanBelToken.UserRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060400160405280600b81526020017f44616e42656c546f6b656e000000000000000000000000000000000000000000815250600090805190602001906200005192919062000609565b50620f42406001556602aa1efb94e0006002553480156200007157600080fd5b5060405162001b6538038062001b658339818101604052810190620000979190620006d0565b81600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000111600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166298968060036200013060201b60201c565b6200012881620493e060006200013060201b60201c565b505062000a92565b60006200014384620002c460201b60201c565b9050604051806060016040528082815260200184815260200183600381111562000196577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815250600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190620001f692919062000609565b506020820151816001015560408201518160020160006101000a81548160ff0219169083600381111562000253577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055509050506004849080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b606060008260601b90506000600467ffffffffffffffff81111562000312577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015620003455781602001600182028036833780820191505090505b50905060006040518060400160405280601081526020017f3031323334353637383961626364656600000000000000000000000000000000815250905060005b60028160ff161015620005dc57818251858360ff1660028110620003d2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b60f81c60ff16620003e8919062000853565b8151811062000420577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b836002836200043b91906200088b565b60ff168151811062000476577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350818251858360ff1660028110620004e5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b60f81c60ff16620004fb9190620009b3565b8151811062000533577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b8360016002846200055091906200088b565b6200055c919062000815565b60ff168151811062000597577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080620005d39062000984565b91505062000385565b5081604051602001620005f09190620007cc565b6040516020818303038152906040529350505050919050565b82805462000617906200094e565b90600052602060002090601f0160209004810192826200063b576000855562000687565b82601f106200065657805160ff191683800117855562000687565b8280016001018555821562000687579182015b828111156200068657825182559160200191906001019062000669565b5b5090506200069691906200069a565b5090565b5b80821115620006b55760008160009055506001016200069b565b5090565b600081519050620006ca8162000a78565b92915050565b60008060408385031215620006e457600080fd5b6000620006f485828601620006b9565b92505060206200070785828601620006b9565b9150509250929050565b60006200071e82620007ff565b6200072a81856200080a565b93506200073c81856020860162000918565b80840191505092915050565b6000620007576008836200080a565b91507f50524f4649202d200000000000000000000000000000000000000000000000006000830152600882019050919050565b6000620007996004836200080a565b91507f32303234000000000000000000000000000000000000000000000000000000006000830152600482019050919050565b6000620007d98262000748565b9150620007e7828462000711565b9150620007f4826200078a565b915081905092915050565b600081519050919050565b600081905092915050565b600062000822826200090b565b91506200082f836200090b565b92508260ff03821115620008485762000847620009eb565b5b828201905092915050565b6000620008608262000901565b91506200086d8362000901565b92508262000880576200087f62000a1a565b5b828204905092915050565b600062000898826200090b565b9150620008a5836200090b565b92508160ff0483118215151615620008c257620008c1620009eb565b5b828202905092915050565b6000620008da82620008e1565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015620009385780820151818401526020810190506200091b565b8381111562000948576000848401525b50505050565b600060028204905060018216806200096757607f821691505b602082108114156200097e576200097d62000a49565b5b50919050565b600062000991826200090b565b915060ff821415620009a857620009a7620009eb565b5b600182019050919050565b6000620009c08262000901565b9150620009cd8362000901565b925082620009e057620009df62000a1a565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b62000a8381620008cd565b811462000a8f57600080fd5b50565b6110c38062000aa26000396000f3fe6080604052600436106100865760003560e01c80638da5cb5b116100595780638da5cb5b1461013557806392c40344146101605780639674a0ac1461019d578063a87430ba146101da578063d96a094a1461021957610086565b806306fdde031461008b57806313faede6146100b657806318160ddd146100e157806361c28c361461010c575b600080fd5b34801561009757600080fd5b506100a0610235565b6040516100ad9190610c71565b60405180910390f35b3480156100c257600080fd5b506100cb6102c3565b6040516100d89190610cd1565b60405180910390f35b3480156100ed57600080fd5b506100f66102c9565b6040516101039190610cd1565b60405180910390f35b34801561011857600080fd5b50610133600480360381019061012e9190610a9a565b6102cf565b005b34801561014157600080fd5b5061014a610457565b6040516101579190610c56565b60405180910390f35b34801561016c57600080fd5b5061018760048036038101906101829190610a71565b61047d565b6040516101949190610c71565b60405180910390f35b3480156101a957600080fd5b506101c460048036038101906101bf9190610ae9565b6107aa565b6040516101d19190610c56565b60405180910390f35b3480156101e657600080fd5b5061020160048036038101906101fc9190610a71565b6107e9565b60405161021093929190610c93565b60405180910390f35b610233600480360381019061022e9190610ae9565b6108a8565b005b6000805461024290610ee1565b80601f016020809104026020016040519081016040528092919081815260200182805461026e90610ee1565b80156102bb5780601f10610290576101008083540402835291602001916102bb565b820191906000526020600020905b81548152906001019060200180831161029e57829003601f168201915b505050505081565b60025481565b60015481565b60006102da8461047d565b9050604051806060016040528082815260200184815260200183600381111561032c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815250600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001908051906020019061038a92919061098f565b506020820151816001015560408201518160020160006101000a81548160ff021916908360038111156103e6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055509050506004849080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060008260601b90506000600467ffffffffffffffff8111156104ca577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156104fc5781602001600182028036833780820191505090505b50905060006040518060400160405280601081526020017f3031323334353637383961626364656600000000000000000000000000000000815250905060005b60028160ff16101561077f57818251858360ff1660028110610587577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b60f81c60ff1661059b9190610da0565b815181106105d2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b836002836105eb9190610dd1565b60ff1681518110610625577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350818251858360ff1660028110610693577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b60f81c60ff166106a79190610f3d565b815181106106de577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b8360016002846106f99190610dd1565b6107039190610d69565b60ff168151811061073d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808061077790610f13565b91505061053c565b50816040516020016107919190610c29565b6040516020818303038152906040529350505050919050565b600481815481106107ba57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360205280600052604060002060009150905080600001805461080c90610ee1565b80601f016020809104026020016040519081016040528092919081815260200182805461083890610ee1565b80156108855780601f1061085a57610100808354040283529160200191610885565b820191906000526020600020905b81548152906001019060200180831161086857829003601f168201915b5050505050908060010154908060020160009054906101000a900460ff16905083565b6108d5600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1633836108d8565b50565b80600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600082825461092a9190610e0c565b9250508190555080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008282546109839190610d13565b92505081905550505050565b82805461099b90610ee1565b90600052602060002090601f0160209004810192826109bd5760008555610a04565b82601f106109d657805160ff1916838001178555610a04565b82800160010185558215610a04579182015b82811115610a035782518255916020019190600101906109e8565b5b509050610a119190610a15565b5090565b5b80821115610a2e576000816000905550600101610a16565b5090565b600081359050610a418161104f565b92915050565b600081359050610a5681611066565b92915050565b600081359050610a6b81611076565b92915050565b600060208284031215610a8357600080fd5b6000610a9184828501610a32565b91505092915050565b600080600060608486031215610aaf57600080fd5b6000610abd86828701610a32565b9350506020610ace86828701610a5c565b9250506040610adf86828701610a47565b9150509250925092565b600060208284031215610afb57600080fd5b6000610b0984828501610a5c565b91505092915050565b610b1b81610e40565b82525050565b610b2a81610e9c565b82525050565b6000610b3b82610cec565b610b458185610cf7565b9350610b55818560208601610eae565b610b5e8161102a565b840191505092915050565b6000610b7482610cec565b610b7e8185610d08565b9350610b8e818560208601610eae565b80840191505092915050565b6000610ba7600883610d08565b91507f50524f4649202d200000000000000000000000000000000000000000000000006000830152600882019050919050565b6000610be7600483610d08565b91507f32303234000000000000000000000000000000000000000000000000000000006000830152600482019050919050565b610c2381610e85565b82525050565b6000610c3482610b9a565b9150610c408284610b69565b9150610c4b82610bda565b915081905092915050565b6000602082019050610c6b6000830184610b12565b92915050565b60006020820190508181036000830152610c8b8184610b30565b905092915050565b60006060820190508181036000830152610cad8186610b30565b9050610cbc6020830185610c1a565b610cc96040830184610b21565b949350505050565b6000602082019050610ce66000830184610c1a565b92915050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b6000610d1e82610e85565b9150610d2983610e85565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610d5e57610d5d610f6e565b5b828201905092915050565b6000610d7482610e8f565b9150610d7f83610e8f565b92508260ff03821115610d9557610d94610f6e565b5b828201905092915050565b6000610dab82610e85565b9150610db683610e85565b925082610dc657610dc5610f9d565b5b828204905092915050565b6000610ddc82610e8f565b9150610de783610e8f565b92508160ff0483118215151615610e0157610e00610f6e565b5b828202905092915050565b6000610e1782610e85565b9150610e2283610e85565b925082821015610e3557610e34610f6e565b5b828203905092915050565b6000610e4b82610e65565b9050919050565b6000819050610e608261103b565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000610ea782610e52565b9050919050565b60005b83811015610ecc578082015181840152602081019050610eb1565b83811115610edb576000848401525b50505050565b60006002820490506001821680610ef957607f821691505b60208210811415610f0d57610f0c610ffb565b5b50919050565b6000610f1e82610e8f565b915060ff821415610f3257610f31610f6e565b5b600182019050919050565b6000610f4882610e85565b9150610f5383610e85565b925082610f6357610f62610f9d565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b6004811061104c5761104b610fcc565b5b50565b61105881610e40565b811461106357600080fd5b50565b6004811061107357600080fd5b50565b61107f81610e85565b811461108a57600080fd5b5056fea264697066735822122058475a848af5fb43f1c5cc7fcb564a2218df871cf686a070f6105b63cf854c4564736f6c63430008000033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, investor1 common.Address) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _owner, investor1)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_Contract *ContractCaller) Cost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "cost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_Contract *ContractSession) Cost() (*big.Int, error) {
	return _Contract.Contract.Cost(&_Contract.CallOpts)
}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_Contract *ContractCallerSession) Cost() (*big.Int, error) {
	return _Contract.Contract.Cost(&_Contract.CallOpts)
}

// GetReferralCode is a free data retrieval call binding the contract method 0x92c40344.
//
// Solidity: function getReferralCode(address addr) pure returns(string)
func (_Contract *ContractCaller) GetReferralCode(opts *bind.CallOpts, addr common.Address) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getReferralCode", addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetReferralCode is a free data retrieval call binding the contract method 0x92c40344.
//
// Solidity: function getReferralCode(address addr) pure returns(string)
func (_Contract *ContractSession) GetReferralCode(addr common.Address) (string, error) {
	return _Contract.Contract.GetReferralCode(&_Contract.CallOpts, addr)
}

// GetReferralCode is a free data retrieval call binding the contract method 0x92c40344.
//
// Solidity: function getReferralCode(address addr) pure returns(string)
func (_Contract *ContractCallerSession) GetReferralCode(addr common.Address) (string, error) {
	return _Contract.Contract.GetReferralCode(&_Contract.CallOpts, addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// UserAddress is a free data retrieval call binding the contract method 0x9674a0ac.
//
// Solidity: function userAddress(uint256 ) view returns(address)
func (_Contract *ContractCaller) UserAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserAddress is a free data retrieval call binding the contract method 0x9674a0ac.
//
// Solidity: function userAddress(uint256 ) view returns(address)
func (_Contract *ContractSession) UserAddress(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.UserAddress(&_Contract.CallOpts, arg0)
}

// UserAddress is a free data retrieval call binding the contract method 0x9674a0ac.
//
// Solidity: function userAddress(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) UserAddress(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.UserAddress(&_Contract.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(string referalCode, uint256 balance, uint8 role)
func (_Contract *ContractCaller) Users(opts *bind.CallOpts, arg0 common.Address) (struct {
	ReferalCode string
	Balance     *big.Int
	Role        uint8
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		ReferalCode string
		Balance     *big.Int
		Role        uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReferalCode = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Role = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(string referalCode, uint256 balance, uint8 role)
func (_Contract *ContractSession) Users(arg0 common.Address) (struct {
	ReferalCode string
	Balance     *big.Int
	Role        uint8
}, error) {
	return _Contract.Contract.Users(&_Contract.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(string referalCode, uint256 balance, uint8 role)
func (_Contract *ContractCallerSession) Users(arg0 common.Address) (struct {
	ReferalCode string
	Balance     *big.Int
	Role        uint8
}, error) {
	return _Contract.Contract.Users(&_Contract.CallOpts, arg0)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 amount) payable returns()
func (_Contract *ContractTransactor) Buy(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "buy", amount)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 amount) payable returns()
func (_Contract *ContractSession) Buy(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Buy(&_Contract.TransactOpts, amount)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 amount) payable returns()
func (_Contract *ContractTransactorSession) Buy(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Buy(&_Contract.TransactOpts, amount)
}

// SetUser is a paid mutator transaction binding the contract method 0x61c28c36.
//
// Solidity: function setUser(address addr, uint256 balance, uint8 role) returns()
func (_Contract *ContractTransactor) SetUser(opts *bind.TransactOpts, addr common.Address, balance *big.Int, role uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setUser", addr, balance, role)
}

// SetUser is a paid mutator transaction binding the contract method 0x61c28c36.
//
// Solidity: function setUser(address addr, uint256 balance, uint8 role) returns()
func (_Contract *ContractSession) SetUser(addr common.Address, balance *big.Int, role uint8) (*types.Transaction, error) {
	return _Contract.Contract.SetUser(&_Contract.TransactOpts, addr, balance, role)
}

// SetUser is a paid mutator transaction binding the contract method 0x61c28c36.
//
// Solidity: function setUser(address addr, uint256 balance, uint8 role) returns()
func (_Contract *ContractTransactorSession) SetUser(addr common.Address, balance *big.Int, role uint8) (*types.Transaction, error) {
	return _Contract.Contract.SetUser(&_Contract.TransactOpts, addr, balance, role)
}

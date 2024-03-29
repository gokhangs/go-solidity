// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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

// TodoFactoryMetaData contains all meta data concerning the TodoFactory contract.
var TodoFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"createTodoContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_todoIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"tFAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_todoIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_whichTodo\",\"type\":\"uint256\"}],\"name\":\"tFGet\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"todoArray\",\"outputs\":[{\"internalType\":\"contractTodo\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611b59806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063385135171461005157806348a313621461008157806391d531f91461008b578063a0de3055146100bb575b600080fd5b61006b60048036038101906100669190610378565b6100d7565b6040516100789190610424565b60405180910390f35b610089610116565b005b6100a560048036038101906100a0919061043f565b6101a9565b6040516100b2919061050f565b60405180910390f35b6100d560048036038101906100d09190610666565b610274565b005b600081815481106100e757600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060405161012490610321565b604051809103906000f080158015610140573d6000803e3d6000fd5b5090506000819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6060600083815481106101bf576101be6106c2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639507d39a836040518263ffffffff1660e01b81526004016102229190610700565b600060405180830381865afa15801561023f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102689190610839565b60000151905092915050565b60008281548110610288576102876106c2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b0c8f9dc826040518263ffffffff1660e01b81526004016102eb919061050f565b600060405180830381600087803b15801561030557600080fd5b505af1158015610319573d6000803e3d6000fd5b505050505050565b6112a18061088383390190565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61035581610342565b811461036057600080fd5b50565b6000813590506103728161034c565b92915050565b60006020828403121561038e5761038d610338565b5b600061039c84828501610363565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006103ea6103e56103e0846103a5565b6103c5565b6103a5565b9050919050565b60006103fc826103cf565b9050919050565b600061040e826103f1565b9050919050565b61041e81610403565b82525050565b60006020820190506104396000830184610415565b92915050565b6000806040838503121561045657610455610338565b5b600061046485828601610363565b925050602061047585828601610363565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104b957808201518184015260208101905061049e565b60008484015250505050565b6000601f19601f8301169050919050565b60006104e18261047f565b6104eb818561048a565b93506104fb81856020860161049b565b610504816104c5565b840191505092915050565b6000602082019050818103600083015261052981846104d6565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610573826104c5565b810181811067ffffffffffffffff821117156105925761059161053b565b5b80604052505050565b60006105a561032e565b90506105b1828261056a565b919050565b600067ffffffffffffffff8211156105d1576105d061053b565b5b6105da826104c5565b9050602081019050919050565b82818337600083830152505050565b6000610609610604846105b6565b61059b565b90508281526020810184848401111561062557610624610536565b5b6106308482856105e7565b509392505050565b600082601f83011261064d5761064c610531565b5b813561065d8482602086016105f6565b91505092915050565b6000806040838503121561067d5761067c610338565b5b600061068b85828601610363565b925050602083013567ffffffffffffffff8111156106ac576106ab61033d565b5b6106b885828601610638565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6106fa81610342565b82525050565b600060208201905061071560008301846106f1565b92915050565b600080fd5b600080fd5b6000610738610733846105b6565b61059b565b90508281526020810184848401111561075457610753610536565b5b61075f84828561049b565b509392505050565b600082601f83011261077c5761077b610531565b5b815161078c848260208601610725565b91505092915050565b60008115159050919050565b6107aa81610795565b81146107b557600080fd5b50565b6000815190506107c7816107a1565b92915050565b6000604082840312156107e3576107e261071b565b5b6107ed604061059b565b9050600082015167ffffffffffffffff81111561080d5761080c610720565b5b61081984828501610767565b600083015250602061082d848285016107b8565b60208301525092915050565b60006020828403121561084f5761084e610338565b5b600082015167ffffffffffffffff81111561086d5761086c61033d565b5b610879848285016107cd565b9150509291505056fe608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611240806100616000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063751ef7531161005b578063751ef753146100d85780638da5cb5b146100f45780639507d39a14610112578063b0c8f9dc146101425761007d565b80630f560cd7146100825780634cc82215146100a05780634e324847146100bc575b600080fd5b61008a61015e565b6040516100979190610984565b60405180910390f35b6100ba60048036038101906100b591906109f0565b61026a565b005b6100d660048036038101906100d19190610b7e565b6103de565b005b6100f260048036038101906100ed91906109f0565b6104ad565b005b6100fc61057a565b6040516101099190610c2e565b60405180910390f35b61012c600480360381019061012791906109f0565b6105a0565b6040516101399190610c86565b60405180910390f35b61015c60048036038101906101579190610ca8565b610689565b005b60606000805480602002602001604051908101604052809291908181526020016000905b8282101561026157838290600052602060002090600202016040518060400160405290816000820180546101b590610d20565b80601f01602080910402602001604051908101604052809291908181526020018280546101e190610d20565b801561022e5780601f106102035761010080835404028352916020019161022e565b820191906000526020600020905b81548152906001019060200180831161021157829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152505081526020019060010190610182565b50505050905090565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102c457600080fd5b60008190505b60016000805490506102dc9190610d80565b81101561038b5760006001836102f29190610db4565b8154811061030357610302610de8565b5b90600052602060002090600202016000838154811061032557610324610de8565b5b9060005260206000209060020201600082018160000190816103479190610fd9565b506001820160009054906101000a900460ff168160010160006101000a81548160ff0219169083151502179055509050508080610383906110c1565b9150506102ca565b50600080548061039e5761039d611109565b5b6001900381819060005260206000209060020201600080820160006103c39190610761565b6001820160006101000a81549060ff02191690555050905550565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461043857600080fd5b806000848154811061044d5761044c610de8565b5b906000526020600020906002020160010160006101000a81548160ff021916908315150217905550816000848154811061048a57610489610de8565b5b906000526020600020906002020160000190816104a79190611138565b50505050565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461050757600080fd5b6000818154811061051b5761051a610de8565b5b906000526020600020906002020160010160009054906101000a900460ff16156000828154811061054f5761054e610de8565b5b906000526020600020906002020160010160006101000a81548160ff02191690831515021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6105a86107a1565b600082815481106105bc576105bb610de8565b5b90600052602060002090600202016040518060400160405290816000820180546105e590610d20565b80601f016020809104026020016040519081016040528092919081815260200182805461061190610d20565b801561065e5780601f106106335761010080835404028352916020019161065e565b820191906000526020600020905b81548152906001019060200180831161064157829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106e357600080fd5b60006040518060400160405280838152602001600015158152509080600181540180825580915050600190039060005260206000209060020201600090919091909150600082015181600001908161073b9190611138565b5060208201518160010160006101000a81548160ff021916908315150217905550505050565b50805461076d90610d20565b6000825580601f1061077f575061079e565b601f01602090049060005260206000209081019061079d91906107bd565b5b50565b6040518060400160405280606081526020016000151581525090565b5b808211156107d65760008160009055506001016107be565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610840578082015181840152602081019050610825565b60008484015250505050565b6000601f19601f8301169050919050565b600061086882610806565b6108728185610811565b9350610882818560208601610822565b61088b8161084c565b840191505092915050565b60008115159050919050565b6108ab81610896565b82525050565b600060408301600083015184820360008601526108ce828261085d565b91505060208301516108e360208601826108a2565b508091505092915050565b60006108fa83836108b1565b905092915050565b6000602082019050919050565b600061091a826107da565b61092481856107e5565b935083602082028501610936856107f6565b8060005b85811015610972578484038952815161095385826108ee565b945061095e83610902565b925060208a0199505060018101905061093a565b50829750879550505050505092915050565b6000602082019050818103600083015261099e818461090f565b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6109cd816109ba565b81146109d857600080fd5b50565b6000813590506109ea816109c4565b92915050565b600060208284031215610a0657610a056109b0565b5b6000610a14848285016109db565b91505092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610a5f8261084c565b810181811067ffffffffffffffff82111715610a7e57610a7d610a27565b5b80604052505050565b6000610a916109a6565b9050610a9d8282610a56565b919050565b600067ffffffffffffffff821115610abd57610abc610a27565b5b610ac68261084c565b9050602081019050919050565b82818337600083830152505050565b6000610af5610af084610aa2565b610a87565b905082815260208101848484011115610b1157610b10610a22565b5b610b1c848285610ad3565b509392505050565b600082601f830112610b3957610b38610a1d565b5b8135610b49848260208601610ae2565b91505092915050565b610b5b81610896565b8114610b6657600080fd5b50565b600081359050610b7881610b52565b92915050565b600080600060608486031215610b9757610b966109b0565b5b6000610ba5868287016109db565b935050602084013567ffffffffffffffff811115610bc657610bc56109b5565b5b610bd286828701610b24565b9250506040610be386828701610b69565b9150509250925092565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c1882610bed565b9050919050565b610c2881610c0d565b82525050565b6000602082019050610c436000830184610c1f565b92915050565b60006040830160008301518482036000860152610c66828261085d565b9150506020830151610c7b60208601826108a2565b508091505092915050565b60006020820190508181036000830152610ca08184610c49565b905092915050565b600060208284031215610cbe57610cbd6109b0565b5b600082013567ffffffffffffffff811115610cdc57610cdb6109b5565b5b610ce884828501610b24565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610d3857607f821691505b602082108103610d4b57610d4a610cf1565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d8b826109ba565b9150610d96836109ba565b9250828203905081811115610dae57610dad610d51565b5b92915050565b6000610dbf826109ba565b9150610dca836109ba565b9250828201905080821115610de257610de1610d51565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081549050610e2681610d20565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610e8f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610e52565b610e998683610e52565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610ed6610ed1610ecc846109ba565b610eb1565b6109ba565b9050919050565b6000819050919050565b610ef083610ebb565b610f04610efc82610edd565b848454610e5f565b825550505050565b600090565b610f19610f0c565b610f24818484610ee7565b505050565b5b81811015610f4857610f3d600082610f11565b600181019050610f2a565b5050565b601f821115610f8d57610f5e81610e2d565b610f6784610e42565b81016020851015610f76578190505b610f8a610f8285610e42565b830182610f29565b50505b505050565b600082821c905092915050565b6000610fb060001984600802610f92565b1980831691505092915050565b6000610fc98383610f9f565b9150826002028217905092915050565b818103610fe75750506110bf565b610ff082610e17565b67ffffffffffffffff81111561100957611008610a27565b5b6110138254610d20565b61101e828285610f4c565b6000601f83116001811461104d576000841561103b578287015490505b6110458582610fbd565b8655506110b8565b601f19841661105b87610e2d565b965061106686610e2d565b60005b8281101561108e57848901548255600182019150600185019450602081019050611069565b868310156110ab57848901546110a7601f891682610f9f565b8355505b6001600288020188555050505b5050505050505b565b60006110cc826109ba565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110fe576110fd610d51565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61114182610806565b67ffffffffffffffff81111561115a57611159610a27565b5b6111648254610d20565b61116f828285610f4c565b600060209050601f8311600181146111a25760008415611190578287015190505b61119a8582610fbd565b865550611202565b601f1984166111b086610e2d565b60005b828110156111d8578489015182556001820191506020850194506020810190506111b3565b868310156111f557848901516111f1601f891682610f9f565b8355505b6001600288020188555050505b50505050505056fea26469706673582212209cacea2e874e89cc8941e3f39b21c79c96a7619be271d4dba82dedf2717397f264736f6c63430008110033a264697066735822122006a7b9e71f32819b098402f9105472c4ed069bbdc8c2d9772e8054cdd026404e64736f6c63430008110033",
}

// TodoFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoFactoryMetaData.ABI instead.
var TodoFactoryABI = TodoFactoryMetaData.ABI

// TodoFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoFactoryMetaData.Bin instead.
var TodoFactoryBin = TodoFactoryMetaData.Bin

// DeployTodoFactory deploys a new Ethereum contract, binding an instance of TodoFactory to it.
func DeployTodoFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TodoFactory, error) {
	parsed, err := TodoFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TodoFactory{TodoFactoryCaller: TodoFactoryCaller{contract: contract}, TodoFactoryTransactor: TodoFactoryTransactor{contract: contract}, TodoFactoryFilterer: TodoFactoryFilterer{contract: contract}}, nil
}

// TodoFactory is an auto generated Go binding around an Ethereum contract.
type TodoFactory struct {
	TodoFactoryCaller     // Read-only binding to the contract
	TodoFactoryTransactor // Write-only binding to the contract
	TodoFactoryFilterer   // Log filterer for contract events
}

// TodoFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoFactorySession struct {
	Contract     *TodoFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoFactoryCallerSession struct {
	Contract *TodoFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TodoFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoFactoryTransactorSession struct {
	Contract     *TodoFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TodoFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoFactoryRaw struct {
	Contract *TodoFactory // Generic contract binding to access the raw methods on
}

// TodoFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoFactoryCallerRaw struct {
	Contract *TodoFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TodoFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoFactoryTransactorRaw struct {
	Contract *TodoFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodoFactory creates a new instance of TodoFactory, bound to a specific deployed contract.
func NewTodoFactory(address common.Address, backend bind.ContractBackend) (*TodoFactory, error) {
	contract, err := bindTodoFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TodoFactory{TodoFactoryCaller: TodoFactoryCaller{contract: contract}, TodoFactoryTransactor: TodoFactoryTransactor{contract: contract}, TodoFactoryFilterer: TodoFactoryFilterer{contract: contract}}, nil
}

// NewTodoFactoryCaller creates a new read-only instance of TodoFactory, bound to a specific deployed contract.
func NewTodoFactoryCaller(address common.Address, caller bind.ContractCaller) (*TodoFactoryCaller, error) {
	contract, err := bindTodoFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoFactoryCaller{contract: contract}, nil
}

// NewTodoFactoryTransactor creates a new write-only instance of TodoFactory, bound to a specific deployed contract.
func NewTodoFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoFactoryTransactor, error) {
	contract, err := bindTodoFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoFactoryTransactor{contract: contract}, nil
}

// NewTodoFactoryFilterer creates a new log filterer instance of TodoFactory, bound to a specific deployed contract.
func NewTodoFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFactoryFilterer, error) {
	contract, err := bindTodoFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFactoryFilterer{contract: contract}, nil
}

// bindTodoFactory binds a generic wrapper to an already deployed contract.
func bindTodoFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TodoFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TodoFactory *TodoFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TodoFactory.Contract.TodoFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TodoFactory *TodoFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TodoFactory.Contract.TodoFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TodoFactory *TodoFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TodoFactory.Contract.TodoFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TodoFactory *TodoFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TodoFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TodoFactory *TodoFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TodoFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TodoFactory *TodoFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TodoFactory.Contract.contract.Transact(opts, method, params...)
}

// TFGet is a free data retrieval call binding the contract method 0x91d531f9.
//
// Solidity: function tFGet(uint256 _todoIndex, uint256 _whichTodo) view returns(string)
func (_TodoFactory *TodoFactoryCaller) TFGet(opts *bind.CallOpts, _todoIndex *big.Int, _whichTodo *big.Int) (string, error) {
	var out []interface{}
	err := _TodoFactory.contract.Call(opts, &out, "tFGet", _todoIndex, _whichTodo)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TFGet is a free data retrieval call binding the contract method 0x91d531f9.
//
// Solidity: function tFGet(uint256 _todoIndex, uint256 _whichTodo) view returns(string)
func (_TodoFactory *TodoFactorySession) TFGet(_todoIndex *big.Int, _whichTodo *big.Int) (string, error) {
	return _TodoFactory.Contract.TFGet(&_TodoFactory.CallOpts, _todoIndex, _whichTodo)
}

// TFGet is a free data retrieval call binding the contract method 0x91d531f9.
//
// Solidity: function tFGet(uint256 _todoIndex, uint256 _whichTodo) view returns(string)
func (_TodoFactory *TodoFactoryCallerSession) TFGet(_todoIndex *big.Int, _whichTodo *big.Int) (string, error) {
	return _TodoFactory.Contract.TFGet(&_TodoFactory.CallOpts, _todoIndex, _whichTodo)
}

// TodoArray is a free data retrieval call binding the contract method 0x38513517.
//
// Solidity: function todoArray(uint256 ) view returns(address)
func (_TodoFactory *TodoFactoryCaller) TodoArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TodoFactory.contract.Call(opts, &out, "todoArray", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TodoArray is a free data retrieval call binding the contract method 0x38513517.
//
// Solidity: function todoArray(uint256 ) view returns(address)
func (_TodoFactory *TodoFactorySession) TodoArray(arg0 *big.Int) (common.Address, error) {
	return _TodoFactory.Contract.TodoArray(&_TodoFactory.CallOpts, arg0)
}

// TodoArray is a free data retrieval call binding the contract method 0x38513517.
//
// Solidity: function todoArray(uint256 ) view returns(address)
func (_TodoFactory *TodoFactoryCallerSession) TodoArray(arg0 *big.Int) (common.Address, error) {
	return _TodoFactory.Contract.TodoArray(&_TodoFactory.CallOpts, arg0)
}

// CreateTodoContract is a paid mutator transaction binding the contract method 0x48a31362.
//
// Solidity: function createTodoContract() returns()
func (_TodoFactory *TodoFactoryTransactor) CreateTodoContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TodoFactory.contract.Transact(opts, "createTodoContract")
}

// CreateTodoContract is a paid mutator transaction binding the contract method 0x48a31362.
//
// Solidity: function createTodoContract() returns()
func (_TodoFactory *TodoFactorySession) CreateTodoContract() (*types.Transaction, error) {
	return _TodoFactory.Contract.CreateTodoContract(&_TodoFactory.TransactOpts)
}

// CreateTodoContract is a paid mutator transaction binding the contract method 0x48a31362.
//
// Solidity: function createTodoContract() returns()
func (_TodoFactory *TodoFactoryTransactorSession) CreateTodoContract() (*types.Transaction, error) {
	return _TodoFactory.Contract.CreateTodoContract(&_TodoFactory.TransactOpts)
}

// TFAdd is a paid mutator transaction binding the contract method 0xa0de3055.
//
// Solidity: function tFAdd(uint256 _todoIndex, string content) returns()
func (_TodoFactory *TodoFactoryTransactor) TFAdd(opts *bind.TransactOpts, _todoIndex *big.Int, content string) (*types.Transaction, error) {
	return _TodoFactory.contract.Transact(opts, "tFAdd", _todoIndex, content)
}

// TFAdd is a paid mutator transaction binding the contract method 0xa0de3055.
//
// Solidity: function tFAdd(uint256 _todoIndex, string content) returns()
func (_TodoFactory *TodoFactorySession) TFAdd(_todoIndex *big.Int, content string) (*types.Transaction, error) {
	return _TodoFactory.Contract.TFAdd(&_TodoFactory.TransactOpts, _todoIndex, content)
}

// TFAdd is a paid mutator transaction binding the contract method 0xa0de3055.
//
// Solidity: function tFAdd(uint256 _todoIndex, string content) returns()
func (_TodoFactory *TodoFactoryTransactorSession) TFAdd(_todoIndex *big.Int, content string) (*types.Transaction, error) {
	return _TodoFactory.Contract.TFAdd(&_TodoFactory.TransactOpts, _todoIndex, content)
}

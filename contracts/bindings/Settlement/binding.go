// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSettlement

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

// IOrderBookOrder is an auto generated low-level Go binding around an user-defined struct.
type IOrderBookOrder struct {
	OrderId             uint32
	Maker               common.Address
	Taker               common.Address
	InputToken          common.Address
	InputAmount         *big.Int
	OutputToken         common.Address
	OutputAmount        *big.Int
	Expiry              *big.Int
	TargetNetworkNumber uint32
}

// IOrderBookOrderResponse is an auto generated low-level Go binding around an user-defined struct.
type IOrderBookOrderResponse struct {
	Recipient           common.Address
	ReferenceOrderIndex uint32
}

// IOrderBookOrderResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IOrderBookOrderResponseMetadata struct {
	OrderResponsedBlock uint32
	HashOfNonSigners    [32]byte
}

// ContractSettlementMetaData contains all meta data concerning the ContractSettlement contract.
var ContractSettlementMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"orderBookAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signChainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allOrderExecutions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfill\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIOrderBook.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hashOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIOrderBook.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIOrderBook.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FulfillEvent\",\"inputs\":[{\"name\":\"sig\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"order\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOrderBook.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderRespondedEvent\",\"inputs\":[{\"name\":\"orderResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOrderBook.OrderResponse\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"OrderResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOrderBook.OrderResponseMetadata\",\"components\":[{\"name\":\"orderResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderFulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TakerMismatch\",\"inputs\":[]}]",
	Bin: "0x6101006040527fcd190519e6feb4299db38636bfbcef24f017d944ee1188f74093454f21ed61f160c05234801561003557600080fd5b506040516110f83803806110f883398101604081905261005491610105565b60408051808201825260098152684f72646572426f6f6b60b81b6020808301918252835180850190945260038452620312e360ec1b908401528151902060808190527fe6bbd6277e1bf288eed5e8d1780f9a50b239e86b153736bceebccf4ea79d90b360a081905263ffffffff851660e052919291849186916001600160a01b038316156100f857600080546001600160a01b0319166001600160a01b0385161790555b5050505050505050610154565b6000806040838503121561011857600080fd5b82516001600160a01b038116811461012f57600080fd5b602084015190925063ffffffff8116811461014957600080fd5b809150509250929050565b60805160a05160c05160e051610f6b61018d600039600061082e015260006101ff0152600061080d015260006107ec0152610f6b6000f3fe608060405234801561001057600080fd5b506004361061004b5760003560e01c80629e8b5714610050578063298fb5691461007b5780633b692c091461009e578063658d6341146100bf575b600080fd5b61006361005e366004610aa5565b6100d4565b60405161007293929190610ac7565b60405180910390f35b61008e610089366004610c67565b61018d565b6040519015158152602001610072565b6100b16100ac366004610cbd565b6101f8565b604051908152602001610072565b6100d26100cd366004610cda565b6102e0565b005b60016020819052600091825260409091208054918101805463ffffffff909316926100fe90610d70565b80601f016020809104026020016040519081016040528092919081815260200182805461012a90610d70565b80156101775780601f1061014c57610100808354040283529160200191610177565b820191906000526020600020905b81548152906001019060200180831161015a57829003601f168201915b5050506002909301549192505063ffffffff1683565b6000806101db84848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506101d592508991506101f89050565b90610529565b60208601516001600160a01b039182169116149150509392505050565b60006102da7f0000000000000000000000000000000000000000000000000000000000000000836000015184602001518560400151866060015187608001518860a001518960c001518a60e001518b61010001516040516020016102bf9a99989796959493929190998a5263ffffffff98891660208b01526001600160a01b0397881660408b015295871660608a0152938616608089015260a088019290925290931660c086015260e0850192909252610100840191909152166101208201526101400190565b6040516020818303038152906040528051906020012061054d565b92915050565b855163ffffffff908116600090815260016020526040902060020154161561031b576040516356f1733f60e01b815260040160405180910390fd5b4663ffffffff1686610100015163ffffffff161461034c5760405163057f3fa760e51b815260040160405180910390fd5b428660e001511015610371576040516362b439dd60e11b815260040160405180910390fd5b61037c86838361018d565b61039957604051638baa579f60e01b815260040160405180910390fd5b60a0860151602087015160c08801516040516323b872dd60e01b81523360048201526001600160a01b03928316602482015260448101919091529116906323b872dd906064016020604051808303816000875af11580156103fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104229190610dab565b5060405180606001604052808663ffffffff16815260200185858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050504363ffffffff908116602093840152895181168252600180845260409092208451815463ffffffff1916921691909117815583830151805191936104bb938501929101906109f3565b50604091820151600291909101805463ffffffff191663ffffffff909216919091179055517f712dd31697b327ff70b0442a703553fe27c70cce6ec06c292a9cbba1ada6d7809061051990849084908a908a908a908a903390610df6565b60405180910390a1505050505050565b60008060006105388585610594565b9150915061054581610604565b509392505050565b60006105576107c7565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b6000808251604114156105cb5760208301516040840151606085015160001a6105bf878285856108a6565b945094505050506105fd565b8251604014156105f557602083015160408401516105ea868383610993565b9350935050506105fd565b506000905060025b9250929050565b600081600481111561061857610618610ef9565b14156106215750565b600181600481111561063557610635610ef9565b14156106885760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064015b60405180910390fd5b600281600481111561069c5761069c610ef9565b14156106ea5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161067f565b60038160048111156106fe576106fe610ef9565b14156107575760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161067f565b600481600481111561076b5761076b610ef9565b14156107c45760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161067f565b50565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006108556109cc565b604080516020810196909652850193909352606084019190915263ffffffff1660808301526001600160a01b031660a082015260c00160405160208183030381529060405280519060200120905090565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156108dd575060009050600361098a565b8460ff16601b141580156108f557508460ff16601c14155b15610906575060009050600461098a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561095a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166109835760006001925092505061098a565b9150600090505b94509492505050565b6000806001600160ff1b038316816109b060ff86901c601b610f0f565b90506109be878288856108a6565b935093505050935093915050565b600080546001600160a01b0316156109ee57506000546001600160a01b031690565b503090565b8280546109ff90610d70565b90600052602060002090601f016020900481019282610a215760008555610a67565b82601f10610a3a57805160ff1916838001178555610a67565b82800160010185558215610a67579182015b82811115610a67578251825591602001919060010190610a4c565b50610a73929150610a77565b5090565b5b80821115610a735760008155600101610a78565b803563ffffffff81168114610aa057600080fd5b919050565b600060208284031215610ab757600080fd5b610ac082610a8c565b9392505050565b600063ffffffff80861683526020606081850152855180606086015260005b81811015610b0257878101830151868201608001528201610ae6565b81811115610b14576000608083880101525b509490911660408401525050601f91909101601f19160160800192915050565b604051610120810167ffffffffffffffff81118282101715610b6657634e487b7160e01b600052604160045260246000fd5b60405290565b80356001600160a01b0381168114610aa057600080fd5b60006101208284031215610b9657600080fd5b610b9e610b34565b9050610ba982610a8c565b8152610bb760208301610b6c565b6020820152610bc860408301610b6c565b6040820152610bd960608301610b6c565b606082015260808201356080820152610bf460a08301610b6c565b60a082015260c082013560c082015260e082013560e0820152610100610c1b818401610a8c565b9082015292915050565b60008083601f840112610c3757600080fd5b50813567ffffffffffffffff811115610c4f57600080fd5b6020830191508360208285010111156105fd57600080fd5b60008060006101408486031215610c7d57600080fd5b610c878585610b83565b925061012084013567ffffffffffffffff811115610ca457600080fd5b610cb086828701610c25565b9497909650939450505050565b60006101208284031215610cd057600080fd5b610ac08383610b83565b6000806000806000806101808789031215610cf457600080fd5b610cfe8888610b83565b9550610d0d6101208801610a8c565b945061014087013567ffffffffffffffff80821115610d2b57600080fd5b610d378a838b01610c25565b9096509450610160890135915080821115610d5157600080fd5b50610d5e89828a01610c25565b979a9699509497509295939492505050565b600181811c90821680610d8457607f821691505b60208210811415610da557634e487b7160e01b600052602260045260246000fd5b50919050565b600060208284031215610dbd57600080fd5b81518015158114610ac057600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60006101a0808352610e0b8184018a8c610dcd565b905063ffffffff885116602084015260018060a01b0360208901511660408401526040880151610e4660608501826001600160a01b03169052565b5060608801516001600160a01b038116608085015250608088015160a084015260a0880151610e8060c08501826001600160a01b03169052565b5060c088015160e084015260e08801516101008181860152808a0151915050610eb261012085018263ffffffff169052565b5063ffffffff8716610140840152828103610160840152610ed4818688610dcd565b915050610eed6101808301846001600160a01b03169052565b98975050505050505050565b634e487b7160e01b600052602160045260246000fd5b60008219821115610f3057634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220975fa42ed18549c14770dedbb1116e8a7014db4936a40f5e7adfded6fbdd601264736f6c634300080c0033",
}

// ContractSettlementABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSettlementMetaData.ABI instead.
var ContractSettlementABI = ContractSettlementMetaData.ABI

// ContractSettlementBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSettlementMetaData.Bin instead.
var ContractSettlementBin = ContractSettlementMetaData.Bin

// DeployContractSettlement deploys a new Ethereum contract, binding an instance of ContractSettlement to it.
func DeployContractSettlement(auth *bind.TransactOpts, backend bind.ContractBackend, orderBookAddr common.Address, signChainId uint32) (common.Address, *types.Transaction, *ContractSettlement, error) {
	parsed, err := ContractSettlementMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSettlementBin), backend, orderBookAddr, signChainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSettlement{ContractSettlementCaller: ContractSettlementCaller{contract: contract}, ContractSettlementTransactor: ContractSettlementTransactor{contract: contract}, ContractSettlementFilterer: ContractSettlementFilterer{contract: contract}}, nil
}

// ContractSettlement is an auto generated Go binding around an Ethereum contract.
type ContractSettlement struct {
	ContractSettlementCaller     // Read-only binding to the contract
	ContractSettlementTransactor // Write-only binding to the contract
	ContractSettlementFilterer   // Log filterer for contract events
}

// ContractSettlementCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSettlementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSettlementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSettlementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSettlementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSettlementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSettlementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSettlementSession struct {
	Contract     *ContractSettlement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractSettlementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSettlementCallerSession struct {
	Contract *ContractSettlementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContractSettlementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSettlementTransactorSession struct {
	Contract     *ContractSettlementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractSettlementRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSettlementRaw struct {
	Contract *ContractSettlement // Generic contract binding to access the raw methods on
}

// ContractSettlementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSettlementCallerRaw struct {
	Contract *ContractSettlementCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSettlementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSettlementTransactorRaw struct {
	Contract *ContractSettlementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSettlement creates a new instance of ContractSettlement, bound to a specific deployed contract.
func NewContractSettlement(address common.Address, backend bind.ContractBackend) (*ContractSettlement, error) {
	contract, err := bindContractSettlement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSettlement{ContractSettlementCaller: ContractSettlementCaller{contract: contract}, ContractSettlementTransactor: ContractSettlementTransactor{contract: contract}, ContractSettlementFilterer: ContractSettlementFilterer{contract: contract}}, nil
}

// NewContractSettlementCaller creates a new read-only instance of ContractSettlement, bound to a specific deployed contract.
func NewContractSettlementCaller(address common.Address, caller bind.ContractCaller) (*ContractSettlementCaller, error) {
	contract, err := bindContractSettlement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementCaller{contract: contract}, nil
}

// NewContractSettlementTransactor creates a new write-only instance of ContractSettlement, bound to a specific deployed contract.
func NewContractSettlementTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSettlementTransactor, error) {
	contract, err := bindContractSettlement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementTransactor{contract: contract}, nil
}

// NewContractSettlementFilterer creates a new log filterer instance of ContractSettlement, bound to a specific deployed contract.
func NewContractSettlementFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSettlementFilterer, error) {
	contract, err := bindContractSettlement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementFilterer{contract: contract}, nil
}

// bindContractSettlement binds a generic wrapper to an already deployed contract.
func bindContractSettlement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSettlementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSettlement *ContractSettlementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSettlement.Contract.ContractSettlementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSettlement *ContractSettlementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSettlement.Contract.ContractSettlementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSettlement *ContractSettlementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSettlement.Contract.ContractSettlementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSettlement *ContractSettlementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSettlement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSettlement *ContractSettlementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSettlement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSettlement *ContractSettlementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSettlement.Contract.contract.Transact(opts, method, params...)
}

// AllOrderExecutions is a free data retrieval call binding the contract method 0x009e8b57.
//
// Solidity: function allOrderExecutions(uint32 ) view returns(uint32 quorumThresholdPercentage, bytes quorumNumbers, uint32 createdBlock)
func (_ContractSettlement *ContractSettlementCaller) AllOrderExecutions(opts *bind.CallOpts, arg0 uint32) (struct {
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	CreatedBlock              uint32
}, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "allOrderExecutions", arg0)

	outstruct := new(struct {
		QuorumThresholdPercentage uint32
		QuorumNumbers             []byte
		CreatedBlock              uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.QuorumThresholdPercentage = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.QuorumNumbers = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.CreatedBlock = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// AllOrderExecutions is a free data retrieval call binding the contract method 0x009e8b57.
//
// Solidity: function allOrderExecutions(uint32 ) view returns(uint32 quorumThresholdPercentage, bytes quorumNumbers, uint32 createdBlock)
func (_ContractSettlement *ContractSettlementSession) AllOrderExecutions(arg0 uint32) (struct {
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	CreatedBlock              uint32
}, error) {
	return _ContractSettlement.Contract.AllOrderExecutions(&_ContractSettlement.CallOpts, arg0)
}

// AllOrderExecutions is a free data retrieval call binding the contract method 0x009e8b57.
//
// Solidity: function allOrderExecutions(uint32 ) view returns(uint32 quorumThresholdPercentage, bytes quorumNumbers, uint32 createdBlock)
func (_ContractSettlement *ContractSettlementCallerSession) AllOrderExecutions(arg0 uint32) (struct {
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	CreatedBlock              uint32
}, error) {
	return _ContractSettlement.Contract.AllOrderExecutions(&_ContractSettlement.CallOpts, arg0)
}

// HashOrder is a free data retrieval call binding the contract method 0x3b692c09.
//
// Solidity: function hashOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
func (_ContractSettlement *ContractSettlementCaller) HashOrder(opts *bind.CallOpts, order IOrderBookOrder) ([32]byte, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x3b692c09.
//
// Solidity: function hashOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
func (_ContractSettlement *ContractSettlementSession) HashOrder(order IOrderBookOrder) ([32]byte, error) {
	return _ContractSettlement.Contract.HashOrder(&_ContractSettlement.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x3b692c09.
//
// Solidity: function hashOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
func (_ContractSettlement *ContractSettlementCallerSession) HashOrder(order IOrderBookOrder) ([32]byte, error) {
	return _ContractSettlement.Contract.HashOrder(&_ContractSettlement.CallOpts, order)
}

// Verify is a free data retrieval call binding the contract method 0x298fb569.
//
// Solidity: function verify((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractSettlement *ContractSettlementCaller) Verify(opts *bind.CallOpts, order IOrderBookOrder, signature []byte) (bool, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "verify", order, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x298fb569.
//
// Solidity: function verify((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractSettlement *ContractSettlementSession) Verify(order IOrderBookOrder, signature []byte) (bool, error) {
	return _ContractSettlement.Contract.Verify(&_ContractSettlement.CallOpts, order, signature)
}

// Verify is a free data retrieval call binding the contract method 0x298fb569.
//
// Solidity: function verify((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractSettlement *ContractSettlementCallerSession) Verify(order IOrderBookOrder, signature []byte) (bool, error) {
	return _ContractSettlement.Contract.Verify(&_ContractSettlement.CallOpts, order, signature)
}

// Fulfill is a paid mutator transaction binding the contract method 0x658d6341.
//
// Solidity: function fulfill((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, bytes sig) returns()
func (_ContractSettlement *ContractSettlementTransactor) Fulfill(opts *bind.TransactOpts, order IOrderBookOrder, quorumThresholdPercentage uint32, quorumNumbers []byte, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "fulfill", order, quorumThresholdPercentage, quorumNumbers, sig)
}

// Fulfill is a paid mutator transaction binding the contract method 0x658d6341.
//
// Solidity: function fulfill((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, bytes sig) returns()
func (_ContractSettlement *ContractSettlementSession) Fulfill(order IOrderBookOrder, quorumThresholdPercentage uint32, quorumNumbers []byte, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, order, quorumThresholdPercentage, quorumNumbers, sig)
}

// Fulfill is a paid mutator transaction binding the contract method 0x658d6341.
//
// Solidity: function fulfill((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, bytes sig) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) Fulfill(order IOrderBookOrder, quorumThresholdPercentage uint32, quorumNumbers []byte, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, order, quorumThresholdPercentage, quorumNumbers, sig)
}

// ContractSettlementFulfillEventIterator is returned from FilterFulfillEvent and is used to iterate over the raw logs and unpacked data for FulfillEvent events raised by the ContractSettlement contract.
type ContractSettlementFulfillEventIterator struct {
	Event *ContractSettlementFulfillEvent // Event containing the contract specifics and raw log

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
func (it *ContractSettlementFulfillEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementFulfillEvent)
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
		it.Event = new(ContractSettlementFulfillEvent)
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
func (it *ContractSettlementFulfillEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementFulfillEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementFulfillEvent represents a FulfillEvent event raised by the ContractSettlement contract.
type ContractSettlementFulfillEvent struct {
	Sig                       []byte
	Order                     IOrderBookOrder
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	Recipient                 common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterFulfillEvent is a free log retrieval operation binding the contract event 0x712dd31697b327ff70b0442a703553fe27c70cce6ec06c292a9cbba1ada6d780.
//
// Solidity: event FulfillEvent(bytes sig, (uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
func (_ContractSettlement *ContractSettlementFilterer) FilterFulfillEvent(opts *bind.FilterOpts) (*ContractSettlementFulfillEventIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "FulfillEvent")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementFulfillEventIterator{contract: _ContractSettlement.contract, event: "FulfillEvent", logs: logs, sub: sub}, nil
}

// WatchFulfillEvent is a free log subscription operation binding the contract event 0x712dd31697b327ff70b0442a703553fe27c70cce6ec06c292a9cbba1ada6d780.
//
// Solidity: event FulfillEvent(bytes sig, (uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
func (_ContractSettlement *ContractSettlementFilterer) WatchFulfillEvent(opts *bind.WatchOpts, sink chan<- *ContractSettlementFulfillEvent) (event.Subscription, error) {

	logs, sub, err := _ContractSettlement.contract.WatchLogs(opts, "FulfillEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementFulfillEvent)
				if err := _ContractSettlement.contract.UnpackLog(event, "FulfillEvent", log); err != nil {
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

// ParseFulfillEvent is a log parse operation binding the contract event 0x712dd31697b327ff70b0442a703553fe27c70cce6ec06c292a9cbba1ada6d780.
//
// Solidity: event FulfillEvent(bytes sig, (uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
func (_ContractSettlement *ContractSettlementFilterer) ParseFulfillEvent(log types.Log) (*ContractSettlementFulfillEvent, error) {
	event := new(ContractSettlementFulfillEvent)
	if err := _ContractSettlement.contract.UnpackLog(event, "FulfillEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementOrderRespondedEventIterator is returned from FilterOrderRespondedEvent and is used to iterate over the raw logs and unpacked data for OrderRespondedEvent events raised by the ContractSettlement contract.
type ContractSettlementOrderRespondedEventIterator struct {
	Event *ContractSettlementOrderRespondedEvent // Event containing the contract specifics and raw log

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
func (it *ContractSettlementOrderRespondedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementOrderRespondedEvent)
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
		it.Event = new(ContractSettlementOrderRespondedEvent)
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
func (it *ContractSettlementOrderRespondedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementOrderRespondedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementOrderRespondedEvent represents a OrderRespondedEvent event raised by the ContractSettlement contract.
type ContractSettlementOrderRespondedEvent struct {
	OrderResponse         IOrderBookOrderResponse
	OrderResponseMetadata IOrderBookOrderResponseMetadata
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterOrderRespondedEvent is a free log retrieval operation binding the contract event 0xa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd.
//
// Solidity: event OrderRespondedEvent((address,uint32) orderResponse, (uint32,bytes32) OrderResponseMetadata)
func (_ContractSettlement *ContractSettlementFilterer) FilterOrderRespondedEvent(opts *bind.FilterOpts) (*ContractSettlementOrderRespondedEventIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "OrderRespondedEvent")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementOrderRespondedEventIterator{contract: _ContractSettlement.contract, event: "OrderRespondedEvent", logs: logs, sub: sub}, nil
}

// WatchOrderRespondedEvent is a free log subscription operation binding the contract event 0xa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd.
//
// Solidity: event OrderRespondedEvent((address,uint32) orderResponse, (uint32,bytes32) OrderResponseMetadata)
func (_ContractSettlement *ContractSettlementFilterer) WatchOrderRespondedEvent(opts *bind.WatchOpts, sink chan<- *ContractSettlementOrderRespondedEvent) (event.Subscription, error) {

	logs, sub, err := _ContractSettlement.contract.WatchLogs(opts, "OrderRespondedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementOrderRespondedEvent)
				if err := _ContractSettlement.contract.UnpackLog(event, "OrderRespondedEvent", log); err != nil {
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

// ParseOrderRespondedEvent is a log parse operation binding the contract event 0xa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd.
//
// Solidity: event OrderRespondedEvent((address,uint32) orderResponse, (uint32,bytes32) OrderResponseMetadata)
func (_ContractSettlement *ContractSettlementFilterer) ParseOrderRespondedEvent(log types.Log) (*ContractSettlementOrderRespondedEvent, error) {
	event := new(ContractSettlementOrderRespondedEvent)
	if err := _ContractSettlement.contract.UnpackLog(event, "OrderRespondedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractFluxLayerServiceManager

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

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractFluxLayerServiceManagerMetaData contains all meta data concerning the ContractFluxLayerServiceManager contract.
var ContractFluxLayerServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_settlement\",\"type\":\"address\",\"internalType\":\"contractSettlement\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settlement\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSettlement\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001b3638038062001b3683398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e051610100516118bf620002776000396000818161010f015261093f01526000818161065e015281816107ba0152818161085101528181610c4d01528181610dd10152610e700152600081816104890152818161051801528181610598015281816109df01528181610a7501528181610b8b0152610d2c01526000818161031901526103f701526000818161014b01528181610a3301528181610ad10152610b5001526118bf6000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638da5cb5b116100715780638da5cb5b146101775780639926ee7d14610188578063a364f4da1461019b578063a98fb355146101ae578063e481af9d146101c1578063f2fde38b146101c957600080fd5b80631b445516146100b957806333cfb7b7146100ce57806338c8ee64146100f7578063511606301461010a5780636b3aa72e14610149578063715018a61461016f575b600080fd5b6100cc6100c7366004611165565b6101dc565b005b6100e16100dc3660046111ef565b610464565b6040516100ee9190611213565b60405180910390f35b6100cc6101053660046111ef565b610934565b6101317f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ee565b7f0000000000000000000000000000000000000000000000000000000000000000610131565b6100cc6109c0565b6033546001600160a01b0316610131565b6100cc610196366004611315565b6109d4565b6100cc6101a93660046111ef565b610a6a565b6100cc6101bc3660046113c0565b610b31565b6100e1610b85565b6100cc6101d73660046111ef565b610f4f565b6101e4610fc5565b60005b818110156103df5782828281811061020157610201611411565b90506020028101906102139190611427565b6102249060408101906020016111ef565b6001600160a01b03166323b872dd333086868681811061024657610246611411565b90506020028101906102589190611427565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156102af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d39190611457565b508282828181106102e6576102e6611411565b90506020028101906102f89190611427565b6103099060408101906020016111ef565b6001600160a01b031663095ea7b37f000000000000000000000000000000000000000000000000000000000000000085858581811061034a5761034a611411565b905060200281019061035c9190611427565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af11580156103aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ce9190611457565b506103d88161148f565b90506101e7565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061042e9085908590600401611543565b600060405180830381600087803b15801561044857600080fd5b505af115801561045c573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156104d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f49190611651565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561055f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610583919061166a565b90506001600160c01b038116158061061d57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106189190611693565b60ff16155b1561063957505060408051600081526020810190915292915050565b600061064d826001600160c01b031661101f565b90506000805b8251811015610723577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061069d5761069d611411565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156106e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107059190611651565b61070f90836116b6565b91508061071b8161148f565b915050610653565b5060008167ffffffffffffffff81111561073f5761073f611260565b604051908082528060200260200182016040528015610768578160200160208202803683370190505b5090506000805b845181101561092757600085828151811061078c5761078c611411565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610801573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108259190611651565b905060005b81811015610911576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561089f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c391906116ce565b600001518686815181106108d9576108d9611411565b6001600160a01b0390921660209283029190910190910152846108fb8161148f565b95505080806109099061148f565b91505061082a565b505050808061091f9061148f565b91505061076f565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109bd5760405162461bcd60e51b815260206004820152602360248201527f6f6e6c79536574746c656d656e743a206e6f742066726f6d20736574746c656d604482015262195b9d60ea1b60648201526084015b60405180910390fd5b50565b6109c8610fc5565b6109d260006110e2565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a1c5760405162461bcd60e51b81526004016109b49061172d565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061042e90859085906004016117f2565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ab25760405162461bcd60e51b81526004016109b49061172d565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610b1657600080fd5b505af1158015610b2a573d6000803e3d6000fd5b5050505050565b610b39610fc5565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610afc90849060040161183d565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610be7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0b9190611693565b60ff16905080610c2957505060408051600081526020810190915290565b6000805b82811015610cde57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610c9c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc09190611651565b610cca90836116b6565b915080610cd68161148f565b915050610c2d565b5060008167ffffffffffffffff811115610cfa57610cfa611260565b604051908082528060200260200182016040528015610d23578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dac9190611693565b60ff16811015610f4557604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610e20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e449190611651565b905060005b81811015610f30576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610ebe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee291906116ce565b60000151858581518110610ef857610ef8611411565b6001600160a01b039092166020928302919091019091015283610f1a8161148f565b9450508080610f289061148f565b915050610e49565b50508080610f3d9061148f565b915050610d2a565b5090949350505050565b610f57610fc5565b6001600160a01b038116610fbc5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109b4565b6109bd816110e2565b6033546001600160a01b031633146109d25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109b4565b606060008061102d84611134565b61ffff1667ffffffffffffffff81111561104957611049611260565b6040519080825280601f01601f191660200182016040528015611073576020820181803683370190505b5090506000805b82518210801561108b575061010081105b15610f45576001811b9350858416156110d2578060f81b8383815181106110b4576110b4611411565b60200101906001600160f81b031916908160001a9053508160010191505b6110db8161148f565b905061107a565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b821561115f57611149600184611850565b909216918061115781611867565b915050611138565b92915050565b6000806020838503121561117857600080fd5b823567ffffffffffffffff8082111561119057600080fd5b818501915085601f8301126111a457600080fd5b8135818111156111b357600080fd5b8660208260051b85010111156111c857600080fd5b60209290920196919550909350505050565b6001600160a01b03811681146109bd57600080fd5b60006020828403121561120157600080fd5b813561120c816111da565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156112545783516001600160a01b03168352928401929184019160010161122f565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561129957611299611260565b60405290565b600067ffffffffffffffff808411156112ba576112ba611260565b604051601f8501601f19908116603f011681019082821181831017156112e2576112e2611260565b816040528093508581528686860111156112fb57600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561132857600080fd5b8235611333816111da565b9150602083013567ffffffffffffffff8082111561135057600080fd5b908401906060828703121561136457600080fd5b61136c611276565b82358281111561137b57600080fd5b83019150601f8201871361138e57600080fd5b61139d8783356020850161129f565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156113d257600080fd5b813567ffffffffffffffff8111156113e957600080fd5b8201601f810184136113fa57600080fd5b6114098482356020840161129f565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261143d57600080fd5b9190910192915050565b8035611452816111da565b919050565b60006020828403121561146957600080fd5b8151801515811461120c57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114a3576114a3611479565b5060010190565b6bffffffffffffffffffffffff811681146109bd57600080fd5b8183526000602080850194508260005b858110156115245781356114e7816111da565b6001600160a01b03168752818301356114ff816114aa565b6bffffffffffffffffffffffff168784015260409687019691909101906001016114d4565b509495945050505050565b803563ffffffff8116811461145257600080fd5b60208082528181018390526000906040808401600586901b8501820187855b8881101561164357878303603f190184528135368b9003609e1901811261158857600080fd5b8a0160a0813536839003601e190181126115a157600080fd5b8201803567ffffffffffffffff8111156115ba57600080fd5b8060061b36038413156115cc57600080fd5b8287526115de838801828c85016114c4565b925050506115ed888301611447565b6001600160a01b0316888601528187013587860152606061160f81840161152f565b63ffffffff1690860152608061162683820161152f565b63ffffffff16950194909452509285019290850190600101611562565b509098975050505050505050565b60006020828403121561166357600080fd5b5051919050565b60006020828403121561167c57600080fd5b81516001600160c01b038116811461120c57600080fd5b6000602082840312156116a557600080fd5b815160ff8116811461120c57600080fd5b600082198211156116c9576116c9611479565b500190565b6000604082840312156116e057600080fd5b6040516040810181811067ffffffffffffffff8211171561170357611703611260565b6040528251611711816111da565b81526020830151611721816114aa565b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156117cb576020818501810151868301820152016117af565b818111156117dd576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b038316815260406020820152600082516060604084015261181c60a08401826117a5565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061120c60208301846117a5565b60008282101561186257611862611479565b500390565b600061ffff8083168181141561187f5761187f611479565b600101939250505056fea26469706673582212204c2037a11cfc83aeb41dced0ea2f41f51b9822fe70c9477509fa4a6059553d1764736f6c634300080c0033",
}

// ContractFluxLayerServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractFluxLayerServiceManagerMetaData.ABI instead.
var ContractFluxLayerServiceManagerABI = ContractFluxLayerServiceManagerMetaData.ABI

// ContractFluxLayerServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractFluxLayerServiceManagerMetaData.Bin instead.
var ContractFluxLayerServiceManagerBin = ContractFluxLayerServiceManagerMetaData.Bin

// DeployContractFluxLayerServiceManager deploys a new Ethereum contract, binding an instance of ContractFluxLayerServiceManager to it.
func DeployContractFluxLayerServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _settlement common.Address) (common.Address, *types.Transaction, *ContractFluxLayerServiceManager, error) {
	parsed, err := ContractFluxLayerServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractFluxLayerServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _settlement)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractFluxLayerServiceManager{ContractFluxLayerServiceManagerCaller: ContractFluxLayerServiceManagerCaller{contract: contract}, ContractFluxLayerServiceManagerTransactor: ContractFluxLayerServiceManagerTransactor{contract: contract}, ContractFluxLayerServiceManagerFilterer: ContractFluxLayerServiceManagerFilterer{contract: contract}}, nil
}

// ContractFluxLayerServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractFluxLayerServiceManager struct {
	ContractFluxLayerServiceManagerCaller     // Read-only binding to the contract
	ContractFluxLayerServiceManagerTransactor // Write-only binding to the contract
	ContractFluxLayerServiceManagerFilterer   // Log filterer for contract events
}

// ContractFluxLayerServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractFluxLayerServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFluxLayerServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractFluxLayerServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFluxLayerServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFluxLayerServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFluxLayerServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractFluxLayerServiceManagerSession struct {
	Contract     *ContractFluxLayerServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractFluxLayerServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractFluxLayerServiceManagerCallerSession struct {
	Contract *ContractFluxLayerServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// ContractFluxLayerServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractFluxLayerServiceManagerTransactorSession struct {
	Contract     *ContractFluxLayerServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// ContractFluxLayerServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractFluxLayerServiceManagerRaw struct {
	Contract *ContractFluxLayerServiceManager // Generic contract binding to access the raw methods on
}

// ContractFluxLayerServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractFluxLayerServiceManagerCallerRaw struct {
	Contract *ContractFluxLayerServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractFluxLayerServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractFluxLayerServiceManagerTransactorRaw struct {
	Contract *ContractFluxLayerServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractFluxLayerServiceManager creates a new instance of ContractFluxLayerServiceManager, bound to a specific deployed contract.
func NewContractFluxLayerServiceManager(address common.Address, backend bind.ContractBackend) (*ContractFluxLayerServiceManager, error) {
	contract, err := bindContractFluxLayerServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractFluxLayerServiceManager{ContractFluxLayerServiceManagerCaller: ContractFluxLayerServiceManagerCaller{contract: contract}, ContractFluxLayerServiceManagerTransactor: ContractFluxLayerServiceManagerTransactor{contract: contract}, ContractFluxLayerServiceManagerFilterer: ContractFluxLayerServiceManagerFilterer{contract: contract}}, nil
}

// NewContractFluxLayerServiceManagerCaller creates a new read-only instance of ContractFluxLayerServiceManager, bound to a specific deployed contract.
func NewContractFluxLayerServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractFluxLayerServiceManagerCaller, error) {
	contract, err := bindContractFluxLayerServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractFluxLayerServiceManagerCaller{contract: contract}, nil
}

// NewContractFluxLayerServiceManagerTransactor creates a new write-only instance of ContractFluxLayerServiceManager, bound to a specific deployed contract.
func NewContractFluxLayerServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractFluxLayerServiceManagerTransactor, error) {
	contract, err := bindContractFluxLayerServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractFluxLayerServiceManagerTransactor{contract: contract}, nil
}

// NewContractFluxLayerServiceManagerFilterer creates a new log filterer instance of ContractFluxLayerServiceManager, bound to a specific deployed contract.
func NewContractFluxLayerServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFluxLayerServiceManagerFilterer, error) {
	contract, err := bindContractFluxLayerServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFluxLayerServiceManagerFilterer{contract: contract}, nil
}

// bindContractFluxLayerServiceManager binds a generic wrapper to an already deployed contract.
func bindContractFluxLayerServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractFluxLayerServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractFluxLayerServiceManager.Contract.ContractFluxLayerServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.ContractFluxLayerServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.ContractFluxLayerServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractFluxLayerServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractFluxLayerServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractFluxLayerServiceManager.Contract.AvsDirectory(&_ContractFluxLayerServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractFluxLayerServiceManager.Contract.AvsDirectory(&_ContractFluxLayerServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractFluxLayerServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractFluxLayerServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractFluxLayerServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractFluxLayerServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractFluxLayerServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractFluxLayerServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractFluxLayerServiceManager.Contract.GetRestakeableStrategies(&_ContractFluxLayerServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractFluxLayerServiceManager.Contract.GetRestakeableStrategies(&_ContractFluxLayerServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractFluxLayerServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) Owner() (common.Address, error) {
	return _ContractFluxLayerServiceManager.Contract.Owner(&_ContractFluxLayerServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractFluxLayerServiceManager.Contract.Owner(&_ContractFluxLayerServiceManager.CallOpts)
}

// Settlement is a free data retrieval call binding the contract method 0x51160630.
//
// Solidity: function settlement() view returns(address)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCaller) Settlement(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractFluxLayerServiceManager.contract.Call(opts, &out, "settlement")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Settlement is a free data retrieval call binding the contract method 0x51160630.
//
// Solidity: function settlement() view returns(address)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) Settlement() (common.Address, error) {
	return _ContractFluxLayerServiceManager.Contract.Settlement(&_ContractFluxLayerServiceManager.CallOpts)
}

// Settlement is a free data retrieval call binding the contract method 0x51160630.
//
// Solidity: function settlement() view returns(address)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerCallerSession) Settlement() (common.Address, error) {
	return _ContractFluxLayerServiceManager.Contract.Settlement(&_ContractFluxLayerServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractFluxLayerServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractFluxLayerServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.FreezeOperator(&_ContractFluxLayerServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.FreezeOperator(&_ContractFluxLayerServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.PayForRange(&_ContractFluxLayerServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.PayForRange(&_ContractFluxLayerServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.RegisterOperatorToAVS(&_ContractFluxLayerServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.RegisterOperatorToAVS(&_ContractFluxLayerServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.RenounceOwnership(&_ContractFluxLayerServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.RenounceOwnership(&_ContractFluxLayerServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.TransferOwnership(&_ContractFluxLayerServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.TransferOwnership(&_ContractFluxLayerServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.UpdateAVSMetadataURI(&_ContractFluxLayerServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractFluxLayerServiceManager.Contract.UpdateAVSMetadataURI(&_ContractFluxLayerServiceManager.TransactOpts, _metadataURI)
}

// ContractFluxLayerServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractFluxLayerServiceManager contract.
type ContractFluxLayerServiceManagerInitializedIterator struct {
	Event *ContractFluxLayerServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractFluxLayerServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFluxLayerServiceManagerInitialized)
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
		it.Event = new(ContractFluxLayerServiceManagerInitialized)
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
func (it *ContractFluxLayerServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFluxLayerServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFluxLayerServiceManagerInitialized represents a Initialized event raised by the ContractFluxLayerServiceManager contract.
type ContractFluxLayerServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractFluxLayerServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractFluxLayerServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractFluxLayerServiceManagerInitializedIterator{contract: _ContractFluxLayerServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractFluxLayerServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractFluxLayerServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFluxLayerServiceManagerInitialized)
				if err := _ContractFluxLayerServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractFluxLayerServiceManagerInitialized, error) {
	event := new(ContractFluxLayerServiceManagerInitialized)
	if err := _ContractFluxLayerServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFluxLayerServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractFluxLayerServiceManager contract.
type ContractFluxLayerServiceManagerOwnershipTransferredIterator struct {
	Event *ContractFluxLayerServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractFluxLayerServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFluxLayerServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractFluxLayerServiceManagerOwnershipTransferred)
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
func (it *ContractFluxLayerServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFluxLayerServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFluxLayerServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractFluxLayerServiceManager contract.
type ContractFluxLayerServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractFluxLayerServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractFluxLayerServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractFluxLayerServiceManagerOwnershipTransferredIterator{contract: _ContractFluxLayerServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractFluxLayerServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractFluxLayerServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFluxLayerServiceManagerOwnershipTransferred)
				if err := _ContractFluxLayerServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractFluxLayerServiceManager *ContractFluxLayerServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractFluxLayerServiceManagerOwnershipTransferred, error) {
	event := new(ContractFluxLayerServiceManagerOwnershipTransferred)
	if err := _ContractFluxLayerServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

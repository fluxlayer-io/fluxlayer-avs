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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// SettlementOrder is an auto generated low-level Go binding around an user-defined struct.
type SettlementOrder struct {
	OrderId                   uint32
	InputToken                common.Address
	InputAmount               *big.Int
	OutputToken               common.Address
	OutputAmount              *big.Int
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	CreatedBlock              uint32
}

// SettlementOrderResponse is an auto generated low-level Go binding around an user-defined struct.
type SettlementOrderResponse struct {
	ReferenceOrderIndex uint32
	TxSuccess           bool
}

// SettlementOrderResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type SettlementOrderResponseMetadata struct {
	OrderResponsedBlock uint32
	HashOfNonSigners    [32]byte
}

// ContractSettlementMetaData contains all meta data concerning the ContractSettlement contract.
var ContractSettlementMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderDetails\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfill\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestOrderNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orderNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToFulfill\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structSettlement.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"orderResponse\",\"type\":\"tuple\",\"internalType\":\"structSettlement.OrderResponse\",\"components\":[{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"txSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderRespondedEvent\",\"inputs\":[{\"name\":\"orderResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSettlement.OrderResponse\",\"components\":[{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"txSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"OrderResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSettlement.OrderResponseMetadata\",\"components\":[{\"name\":\"orderResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"fulfillEvent\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"orderNum\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162005a8938038062005a898339810160408190526200003591620001ed565b80806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001ed565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620001ed565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b39190620001ed565b6001600160a01b031660e05250506097805460ff1916600117905562000214565b6001600160a01b0381168114620001ea57600080fd5b50565b6000602082840312156200020057600080fd5b81516200020d81620001d4565b9392505050565b60805160a05160c05160e05161580362000286600039600081816104bd0152611f6501526000818161033101526121470152600081816103580152818161231d01526124df01526000818161037f01528181610d5501528181611c3001528181611dc8015261200201526158036000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806370acd7e91161010f578063cefdc1d4116100a2578063e772935f11610071578063e772935f146104df578063f2fde38b146104ff578063fabc1cbc14610512578063fd3d73bd1461052557600080fd5b8063cefdc1d414610467578063cfe3074e14610488578063dd3285b314610498578063df5cf723146104b857600080fd5b8063aa4f5cc6116100de578063aa4f5cc614610401578063b04dd81b14610428578063b98d090814610447578063c0c53b8b1461045457600080fd5b806370acd7e9146103c2578063715018a6146103d5578063886f1195146103dd5780638da5cb5b146103f057600080fd5b8063595c6a67116101875780635df45946116101565780635df459461461032c57806368304835146103535780636d14a9871461037a5780636efb4636146103a157600080fd5b8063595c6a67146102bf5780635ac86ab7146102c75780635c155662146102fa5780635c975abb1461031a57600080fd5b8063245a7bfc116101c3578063245a7bfc146102415780633563b0d11461026c578063416c7e5e1461028c5780634f739f741461029f57600080fd5b806310d67a2f146101ea578063136439dd146101ff578063171f1d5b14610212575b600080fd5b6101fd6101f83660046141ae565b610538565b005b6101fd61020d3660046141cb565b6105f4565b610225610220366004614349565b610733565b6040805192151583529015156020830152015b60405180910390f35b60cd54610254906001600160a01b031681565b6040516001600160a01b039091168152602001610238565b61027f61027a3660046143b7565b6108bd565b6040516102389190614512565b6101fd61029a36600461453a565b610d53565b6102b26102ad36600461459f565b610ec8565b60405161023891906146a3565b6101fd6115ee565b6102ea6102d536600461476d565b606654600160ff9092169190911b9081161490565b6040519015158152602001610238565b61030d6103083660046147ad565b6116b5565b604051610238919061485e565b6066545b604051908152602001610238565b6102547f000000000000000000000000000000000000000000000000000000000000000081565b6102547f000000000000000000000000000000000000000000000000000000000000000081565b6102547f000000000000000000000000000000000000000000000000000000000000000081565b6103b46103af366004614b1e565b61187d565b604051610238929190614bde565b6101fd6103d0366004614c32565b612794565b6101fd6129b9565b606554610254906001600160a01b031681565b6033546001600160a01b0316610254565b61041461040f366004614ce8565b6129cd565b604051610238989796959493929190614d52565b60c95463ffffffff165b60405163ffffffff9091168152602001610238565b6097546102ea9060ff1681565b6101fd610462366004614dbb565b612ab5565b61047a610475366004614e06565b612bf0565b604051610238929190614e3d565b60c9546104329063ffffffff1681565b61031e6104a6366004614ce8565b60ca6020526000908152604090205481565b6102547f000000000000000000000000000000000000000000000000000000000000000081565b61031e6104ed366004614ce8565b60cc6020526000908152604090205481565b6101fd61050d3660046141ae565b612d82565b6101fd6105203660046141cb565b612df8565b6101fd610533366004614e5e565b612f54565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561058b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105af9190614ee7565b6001600160a01b0316336001600160a01b0316146105e85760405162461bcd60e51b81526004016105df90614f04565b60405180910390fd5b6105f18161333a565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561063c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106609190614f4e565b61067c5760405162461bcd60e51b81526004016105df90614f6b565b606654818116146106f55760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016105df565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061077b5761077b614fb3565b60200201518951600160200201518a602001516000600281106107a0576107a0614fb3565b60200201518b602001516001600281106107bc576107bc614fb3565b602090810291909101518c518d8301516040516108199a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c61083c9190614fc9565b90506108af61085561084e8884613431565b86906134c8565b61085d61355c565b6108a561089685610890604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613431565b61089f8c61361c565b906134c8565b886201d4c06136ac565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109239190614ee7565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610965573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109899190614ee7565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ef9190614ee7565b9050600086516001600160401b03811115610a0c57610a0c6141e4565b604051908082528060200260200182016040528015610a3f57816020015b6060815260200190600190039081610a2a5790505b50905060005b8751811015610d47576000888281518110610a6257610a62614fb3565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610ac3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610aeb9190810190614feb565b905080516001600160401b03811115610b0657610b066141e4565b604051908082528060200260200182016040528015610b5157816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610b245790505b50848481518110610b6457610b64614fb3565b602002602001018190525060005b8151811015610d31576040518060600160405280876001600160a01b03166347b314e8858581518110610ba757610ba7614fb3565b60200260200101516040518263ffffffff1660e01b8152600401610bcd91815260200190565b602060405180830381865afa158015610bea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0e9190614ee7565b6001600160a01b03168152602001838381518110610c2e57610c2e614fb3565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610c5c57610c5c614fb3565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610cb8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cdc919061507b565b6001600160601b0316815250858581518110610cfa57610cfa614fb3565b60200260200101518281518110610d1357610d13614fb3565b60200260200101819052508080610d29906150ba565b915050610b72565b5050508080610d3f906150ba565b915050610a45565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610db1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd59190614ee7565b6001600160a01b0316336001600160a01b031614610e815760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a4016105df565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b610ef36040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f579190614ee7565b9050610f846040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90610fb4908b90899089906004016150d5565b600060405180830381865afa158015610fd1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ff9919081019061511f565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061102b908b908b908b906004016151d6565b600060405180830381865afa158015611048573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611070919081019061511f565b6040820152856001600160401b0381111561108d5761108d6141e4565b6040519080825280602002602001820160405280156110c057816020015b60608152602001906001900390816110ab5790505b50606082015260005b60ff81168711156114ff576000856001600160401b038111156110ee576110ee6141e4565b604051908082528060200260200182016040528015611117578160200160208202803683370190505b5083606001518360ff168151811061113157611131614fb3565b602002602001018190525060005b868110156113ff5760008c6001600160a01b03166304ec63518a8a8581811061116a5761116a614fb3565b905060200201358e8860000151868151811061118857611188614fb3565b60200260200101516040518463ffffffff1660e01b81526004016111c59392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156111e2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061120691906151ff565b90506001600160c01b0381166112aa5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a4016105df565b8a8a8560ff168181106112bf576112bf614fb3565b6001600160c01b03841692013560f81c9190911c6001908116141590506113ec57856001600160a01b031663dd9846b98a8a8581811061130157611301614fb3565b905060200201358d8d8860ff1681811061131d5761131d614fb3565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611373573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113979190615228565b85606001518560ff16815181106113b0576113b0614fb3565b602002602001015184815181106113c9576113c9614fb3565b63ffffffff90921660209283029190910190910152826113e8816150ba565b9350505b50806113f7816150ba565b91505061113f565b506000816001600160401b0381111561141a5761141a6141e4565b604051908082528060200260200182016040528015611443578160200160208202803683370190505b50905060005b828110156114c45784606001518460ff168151811061146a5761146a614fb3565b6020026020010151818151811061148357611483614fb3565b602002602001015182828151811061149d5761149d614fb3565b63ffffffff90921660209283029190910190910152806114bc816150ba565b915050611449565b508084606001518460ff16815181106114df576114df614fb3565b6020026020010181905250505080806114f790615245565b9150506110c9565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611540573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115649190614ee7565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611597908b908b908e90600401615265565b600060405180830381865afa1580156115b4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115dc919081019061511f565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611636573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061165a9190614f4e565b6116765760405162461bcd60e51b81526004016105df90614f6b565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b81526004016116e792919061528f565b600060405180830381865afa158015611704573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261172c919081019061511f565b9050600084516001600160401b03811115611749576117496141e4565b604051908082528060200260200182016040528015611772578160200160208202803683370190505b50905060005b855181101561187357866001600160a01b03166304ec63518783815181106117a2576117a2614fb3565b6020026020010151878685815181106117bd576117bd614fb3565b60200260200101516040518463ffffffff1660e01b81526004016117fa9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611817573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061183b91906151ff565b6001600160c01b031682828151811061185657611856614fb3565b60209081029190910101528061186b816150ba565b915050611778565b5095945050505050565b60408051808201909152606080825260208201526000846118f45760405162461bcd60e51b815260206004820152603760248201526000805160206157ae83398151915260448201527f7265733a20656d7074792071756f72756d20696e70757400000000000000000060648201526084016105df565b6040830151518514801561190c575060a08301515185145b801561191c575060c08301515185145b801561192c575060e08301515185145b6119965760405162461bcd60e51b815260206004820152604160248201526000805160206157ae83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016105df565b82515160208401515114611a0e5760405162461bcd60e51b8152602060048201526044602482018190526000805160206157ae833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016105df565b4363ffffffff168463ffffffff1610611a7d5760405162461bcd60e51b815260206004820152603c60248201526000805160206157ae83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b0000000060648201526084016105df565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611abe57611abe6141e4565b604051908082528060200260200182016040528015611ae7578160200160208202803683370190505b506020820152866001600160401b03811115611b0557611b056141e4565b604051908082528060200260200182016040528015611b2e578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611b6257611b626141e4565b604051908082528060200260200182016040528015611b8b578160200160208202803683370190505b5081526020860151516001600160401b03811115611bab57611bab6141e4565b604051908082528060200260200182016040528015611bd4578160200160208202803683370190505b5081602001819052506000611ca68a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611c7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ca191906152e3565b6138d0565b905060005b876020015151811015611f4157611cf088602001518281518110611cd157611cd1614fb3565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110611d0657611d06614fb3565b60209081029190910101528015611dc6576020830151611d27600183615300565b81518110611d3757611d37614fb3565b602002602001015160001c83602001518281518110611d5857611d58614fb3565b602002602001015160001c11611dc6576040805162461bcd60e51b81526020600482015260248101919091526000805160206157ae83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016105df565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110611e0b57611e0b614fb3565b60200260200101518b8b600001518581518110611e2a57611e2a614fb3565b60200260200101516040518463ffffffff1660e01b8152600401611e679392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611e84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ea891906151ff565b6001600160c01b031683600001518281518110611ec757611ec7614fb3565b602002602001018181525050611f2d61084e611f018486600001518581518110611ef357611ef3614fb3565b602002602001015116613963565b8a602001518481518110611f1757611f17614fb3565b602002602001015161398e90919063ffffffff16565b945080611f39816150ba565b915050611cab565b5050611f4c83613a72565b60975490935060ff16600081611f63576000611fe5565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fe59190615317565b905060005b8a811015612663578215612145578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061204157612041614fb3565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612081573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120a59190615317565b6120af9190615330565b116121455760405162461bcd60e51b815260206004820152606660248201526000805160206157ae83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016105df565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061218657612186614fb3565b9050013560f81c60f81b60f81c8c8c60a0015185815181106121aa576121aa614fb3565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612206573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061222a9190615348565b6001600160401b03191661224d8a604001518381518110611cd157611cd1614fb3565b67ffffffffffffffff1916146122e95760405162461bcd60e51b815260206004820152606160248201526000805160206157ae83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016105df565b6123198960400151828151811061230257612302614fb3565b6020026020010151876134c890919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061235c5761235c614fb3565b9050013560f81c60f81b60f81c8c8c60c00151858151811061238057612380614fb3565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156123dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612400919061507b565b8560200151828151811061241657612416614fb3565b6001600160601b0390921660209283029190910182015285015180518290811061244257612442614fb3565b60200260200101518560000151828151811061246057612460614fb3565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561264e576124d8866000015182815181106124aa576124aa614fb3565b60200260200101518f8f868181106124c4576124c4614fb3565b600192013560f81c9290921c811614919050565b1561263c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061251e5761251e614fb3565b9050013560f81c60f81b60f81c8e8960200151858151811061254257612542614fb3565b60200260200101518f60e00151888151811061256057612560614fb3565b6020026020010151878151811061257957612579614fb3565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156125dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612601919061507b565b875180518590811061261557612615614fb3565b602002602001018181516126299190615373565b6001600160601b03169052506001909101905b80612646816150ba565b915050612484565b5050808061265b906150ba565b915050611fea565b50505060008061267d8c868a606001518b60800151610733565b91509150816126ee5760405162461bcd60e51b815260206004820152604360248201526000805160206157ae83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016105df565b8061274f5760405162461bcd60e51b815260206004820152603960248201526000805160206157ae83398151915260448201527f7265733a207369676e617475726520697320696e76616c69640000000000000060648201526084016105df565b5050600087826020015160405160200161276a92919061539b565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b60006040518061010001604052808b63ffffffff168152602001896001600160a01b03168152602001888152602001876001600160a01b031681526020018681526020018563ffffffff16815260200184848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525063ffffffff431660209182015260405191925061283c918391016153e3565b60408051808303601f19018152918152815160209283012060c9805463ffffffff908116600090815260ca86528481209390935590548116825260cb84529082902084518154868601519184166001600160c01b0319909116176401000000006001600160a01b039283160217825592850151600182015560608501516002820180546001600160a01b03191691909416179092556080840151600383015560a084015160048301805463ffffffff19169190921617905560c08301518051849361290e926005850192910190614026565b5060e091909101516006909101805463ffffffff191663ffffffff92831617905560c9546040517f6f375196052db01a32a94e2166b4d273355741c797cb8eb067534c606f1a50fa92612977928e929116908d9033908e908e908e908e908e908e908e90615483565b60405180910390a160c9546129939063ffffffff1660016154ff565b60c9805463ffffffff191663ffffffff9290921691909117905550505050505050505050565b6129c1613b0d565b6129cb6000613b67565b565b60cb6020526000908152604090208054600182015460028301546003840154600485015460058601805463ffffffff808816986401000000009098046001600160a01b03908116989616959316929190612a2690615527565b80601f0160208091040260200160405190810160405280929190818152602001828054612a5290615527565b8015612a9f5780601f10612a7457610100808354040283529160200191612a9f565b820191906000526020600020905b815481529060010190602001808311612a8257829003601f168201915b5050506006909301549192505063ffffffff1688565b600054610100900460ff1615808015612ad55750600054600160ff909116105b80612aef5750303b158015612aef575060005460ff166001145b612b525760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016105df565b6000805460ff191660011790558015612b75576000805461ff0019166101001790555b612b80846000613bb9565b612b8983613b67565b60cd80546001600160a01b0319166001600160a01b0384161790558015612bea576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612c2b57612c2b614fb3565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612c67908890869060040161528f565b600060405180830381865afa158015612c84573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612cac919081019061511f565b600081518110612cbe57612cbe614fb3565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612d2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d4e91906151ff565b6001600160c01b031690506000612d6482613ca3565b905081612d728a838a6108bd565b9550955050505050935093915050565b612d8a613b0d565b6001600160a01b038116612def5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105df565b6105f181613b67565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e6f9190614ee7565b6001600160a01b0316336001600160a01b031614612e9f5760405162461bcd60e51b81526004016105df90614f04565b606654198119606654191614612f1d5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016105df565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610728565b60cd546001600160a01b03163314612fae5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064016105df565b366000612fbe60c0860186615562565b90925090506000612fd560c0870160a08801614ce8565b90506000612fea610100880160e08901614ce8565b905060ca6000612ffd6020890189614ce8565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161302991906155ed565b60405160208183030381529060405280519060200120146130b25760405162461bcd60e51b815260206004820152603e60248201527f737570706c696564206f7264657220646f6573206e6f74206d6174636820746860448201527f65206f6e65207265636f7264656420696e2074686520636f6e7472616374000060648201526084016105df565b600060cc816130c460208a018a614ce8565b63ffffffff1663ffffffff16815260200190815260200160002054146131425760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526c103a37903a34329037b93232b960991b60648201526084016105df565b60008660405160200161315591906156e3565b60405160208183030381529060405280519060200120905060008061317d838888878c61187d565b9150915060005b8681101561327c578560ff16836020015182815181106131a6576131a6614fb3565b60200260200101516131b891906156f1565b6001600160601b03166064846000015183815181106131d9576131d9614fb3565b60200260200101516001600160601b03166131f49190615720565b101561326a576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d60648201526084016105df565b80613274816150ba565b915050613184565b5060408051808201825263ffffffff431681526020808201849052915190916132a9918c9184910161573f565b6040516020818303038152906040528051906020012060cc60008c60000160208101906132d69190614ce8565b63ffffffff1663ffffffff168152602001908152602001600020819055507ffc76f8e0b77f0d3f2776538c1e70df0b30f8dce6075db04d3ed8bb6f3d4e5f968a8260405161332592919061573f565b60405180910390a15050505050505050505050565b6001600160a01b0381166133c85760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016105df565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261344d6140aa565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561348057613482565bfe5b50806134c05760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016105df565b505092915050565b60408051808201909152600080825260208201526134e46140c8565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156134805750806134c05760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016105df565b6135646140e6565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061364c60008051602061578e83398151915286614fc9565b90505b61365881613d6f565b909350915060008051602061578e833981519152828309831415613692576040805180820190915290815260208101919091529392505050565b60008051602061578e83398151915260018208905061364f565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906136de61410b565b60005b60028110156138a35760006136f7826006615720565b905084826002811061370b5761370b614fb3565b6020020151518361371d836000615330565b600c811061372d5761372d614fb3565b602002015284826002811061374457613744614fb3565b6020020151602001518382600161375b9190615330565b600c811061376b5761376b614fb3565b602002015283826002811061378257613782614fb3565b6020020151515183613795836002615330565b600c81106137a5576137a5614fb3565b60200201528382600281106137bc576137bc614fb3565b60200201515160016020020151836137d5836003615330565b600c81106137e5576137e5614fb3565b60200201528382600281106137fc576137fc614fb3565b60200201516020015160006002811061381757613817614fb3565b602002015183613828836004615330565b600c811061383857613838614fb3565b602002015283826002811061384f5761384f614fb3565b60200201516020015160016002811061386a5761386a614fb3565b60200201518361387b836005615330565b600c811061388b5761388b614fb3565b6020020152508061389b816150ba565b9150506136e1565b506138ac61412a565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6000806138dc84613df1565b9050808360ff166001901b1161395a5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016105df565b90505b92915050565b6000805b821561395d57613978600184615300565b90921691806139868161576b565b915050613967565b60408051808201909152600080825260208201526102008261ffff16106139ea5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016105df565b8161ffff16600114156139fe57508161395d565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613a6757600161ffff871660ff83161c81161415613a4a57613a4784846134c8565b93505b613a5483846134c8565b92506201fffe600192831b169101613a1a565b509195945050505050565b60408051808201909152600080825260208201528151158015613a9757506020820151155b15613ab5575050604080518082019091526000808252602082015290565b60405180604001604052808360000151815260200160008051602061578e8339815191528460200151613ae89190614fc9565b613b009060008051602061578e833981519152615300565b905292915050565b919050565b6033546001600160a01b031633146129cb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105df565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546001600160a01b0316158015613bda57506001600160a01b03821615155b613c5c5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016105df565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613c9f8261333a565b5050565b6060600080613cb184613963565b61ffff166001600160401b03811115613ccc57613ccc6141e4565b6040519080825280601f01601f191660200182016040528015613cf6576020820181803683370190505b5090506000805b825182108015613d0e575061010081105b15613d65576001811b935085841615613d55578060f81b838381518110613d3757613d37614fb3565b60200101906001600160f81b031916908160001a9053508160010191505b613d5e816150ba565b9050613cfd565b5090949350505050565b6000808060008051602061578e833981519152600360008051602061578e8339815191528660008051602061578e833981519152888909090890506000613de5827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260008051602061578e833981519152613f7e565b91959194509092505050565b600061010082511115613e7a5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016105df565b8151613e8857506000919050565b60008083600081518110613e9e57613e9e614fb3565b0160200151600160f89190911c81901b92505b8451811015613f7557848181518110613ecc57613ecc614fb3565b0160200151600160f89190911c1b9150828211613f615760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016105df565b91811791613f6e816150ba565b9050613eb1565b50909392505050565b600080613f8961412a565b613f91614148565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561348057508261401b5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016105df565b505195945050505050565b82805461403290615527565b90600052602060002090601f016020900481019282614054576000855561409a565b82601f1061406d57805160ff191683800117855561409a565b8280016001018555821561409a579182015b8281111561409a57825182559160200191906001019061407f565b506140a6929150614166565b5090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806140f961417b565b815260200161410661417b565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b5b808211156140a65760008155600101614167565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146105f157600080fd5b6000602082840312156141c057600080fd5b813561395a81614199565b6000602082840312156141dd57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561421c5761421c6141e4565b60405290565b60405161010081016001600160401b038111828210171561421c5761421c6141e4565b604051601f8201601f191681016001600160401b038111828210171561426d5761426d6141e4565b604052919050565b60006040828403121561428757600080fd5b61428f6141fa565b9050813581526020820135602082015292915050565b600082601f8301126142b657600080fd5b604051604081018181106001600160401b03821117156142d8576142d86141e4565b80604052508060408401858111156142ef57600080fd5b845b81811015613a675780358352602092830192016142f1565b60006080828403121561431b57600080fd5b6143236141fa565b905061432f83836142a5565b815261433e83604084016142a5565b602082015292915050565b600080600080610120858703121561436057600080fd5b843593506143718660208701614275565b92506143808660608701614309565b915061438f8660e08701614275565b905092959194509250565b63ffffffff811681146105f157600080fd5b8035613b088161439a565b6000806000606084860312156143cc57600080fd5b83356143d781614199565b92506020848101356001600160401b03808211156143f457600080fd5b818701915087601f83011261440857600080fd5b81358181111561441a5761441a6141e4565b61442c601f8201601f19168501614245565b9150808252888482850101111561444257600080fd5b8084840185840137600084828401015250809450505050614465604085016143ac565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614504578385038a52825180518087529087019087870190845b818110156144ef57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016144ab565b50509a87019a9550509185019160010161448d565b509298975050505050505050565b602081526000614525602083018461446e565b9392505050565b80151581146105f157600080fd5b60006020828403121561454c57600080fd5b813561395a8161452c565b60008083601f84011261456957600080fd5b5081356001600160401b0381111561458057600080fd5b60208301915083602082850101111561459857600080fd5b9250929050565b600080600080600080608087890312156145b857600080fd5b86356145c381614199565b955060208701356145d38161439a565b945060408701356001600160401b03808211156145ef57600080fd5b6145fb8a838b01614557565b9096509450606089013591508082111561461457600080fd5b818901915089601f83011261462857600080fd5b81358181111561463757600080fd5b8a60208260051b850101111561464c57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561469857815163ffffffff1687529582019590820190600101614676565b509495945050505050565b6000602080835283516080828501526146bf60a0850182614662565b905081850151601f19808684030160408701526146dc8383614662565b925060408701519150808684030160608701526146f98383614662565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614750578487830301845261473e828751614662565b95880195938801939150600101614724565b509998505050505050505050565b60ff811681146105f157600080fd5b60006020828403121561477f57600080fd5b813561395a8161475e565b60006001600160401b038211156147a3576147a36141e4565b5060051b60200190565b6000806000606084860312156147c257600080fd5b83356147cd81614199565b92506020848101356001600160401b038111156147e957600080fd5b8501601f810187136147fa57600080fd5b803561480d6148088261478a565b614245565b81815260059190911b8201830190838101908983111561482c57600080fd5b928401925b8284101561484a57833582529284019290840190614831565b8096505050505050614465604085016143ac565b6020808252825182820181905260009190848201906040850190845b818110156148965783518352928401929184019160010161487a565b50909695505050505050565b600082601f8301126148b357600080fd5b813560206148c36148088361478a565b82815260059290921b840181019181810190868411156148e257600080fd5b8286015b848110156149065780356148f98161439a565b83529183019183016148e6565b509695505050505050565b600082601f83011261492257600080fd5b813560206149326148088361478a565b82815260069290921b8401810191818101908684111561495157600080fd5b8286015b84811015614906576149678882614275565b835291830191604001614955565b600082601f83011261498657600080fd5b813560206149966148088361478a565b82815260059290921b840181019181810190868411156149b557600080fd5b8286015b848110156149065780356001600160401b038111156149d85760008081fd5b6149e68986838b01016148a2565b8452509183019183016149b9565b60006101808284031215614a0757600080fd5b614a0f614222565b905081356001600160401b0380821115614a2857600080fd5b614a34858386016148a2565b83526020840135915080821115614a4a57600080fd5b614a5685838601614911565b60208401526040840135915080821115614a6f57600080fd5b614a7b85838601614911565b6040840152614a8d8560608601614309565b6060840152614a9f8560e08601614275565b6080840152610120840135915080821115614ab957600080fd5b614ac5858386016148a2565b60a0840152610140840135915080821115614adf57600080fd5b614aeb858386016148a2565b60c0840152610160840135915080821115614b0557600080fd5b50614b1284828501614975565b60e08301525092915050565b600080600080600060808688031215614b3657600080fd5b8535945060208601356001600160401b0380821115614b5457600080fd5b614b6089838a01614557565b909650945060408801359150614b758261439a565b90925060608701359080821115614b8b57600080fd5b50614b98888289016149f4565b9150509295509295909350565b600081518084526020808501945080840160005b838110156146985781516001600160601b031687529582019590820190600101614bb9565b6040815260008351604080840152614bf96080840182614ba5565b90506020850151603f19848303016060850152614c168282614ba5565b925050508260208301529392505050565b8035613b0881614199565b60008060008060008060008060006101008a8c031215614c5157600080fd5b8935614c5c8161439a565b985060208a0135614c6c81614199565b975060408a0135614c7c81614199565b965060608a0135955060808a0135614c9381614199565b945060a08a0135935060c08a0135614caa8161439a565b925060e08a01356001600160401b03811115614cc557600080fd5b614cd18c828d01614557565b915080935050809150509295985092959850929598565b600060208284031215614cfa57600080fd5b813561395a8161439a565b6000815180845260005b81811015614d2b57602081850181015186830182015201614d0f565b81811115614d3d576000602083870101525b50601f01601f19169290920160200192915050565b63ffffffff89811682526001600160a01b03898116602084015260408301899052871660608301526080820186905284811660a083015261010060c08301819052600091614da284830187614d05565b925080851660e085015250509998505050505050505050565b600080600060608486031215614dd057600080fd5b8335614ddb81614199565b92506020840135614deb81614199565b91506040840135614dfb81614199565b809150509250925092565b600080600060608486031215614e1b57600080fd5b8335614e2681614199565b9250602084013591506040840135614dfb8161439a565b828152604060208201526000614e56604083018461446e565b949350505050565b60008060008385036080811215614e7457600080fd5b84356001600160401b0380821115614e8b57600080fd5b908601906101008289031215614ea057600080fd5b8195506040601f1984011215614eb557600080fd5b6020870194506060870135925080831115614ecf57600080fd5b5050614edd868287016149f4565b9150509250925092565b600060208284031215614ef957600080fd5b815161395a81614199565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614f6057600080fd5b815161395a8161452c565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082614fe657634e487b7160e01b600052601260045260246000fd5b500690565b60006020808385031215614ffe57600080fd5b82516001600160401b0381111561501457600080fd5b8301601f8101851361502557600080fd5b80516150336148088261478a565b81815260059190911b8201830190838101908783111561505257600080fd5b928401925b8284101561507057835182529284019290840190615057565b979650505050505050565b60006020828403121561508d57600080fd5b81516001600160601b038116811461395a57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156150ce576150ce6150a4565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561510257600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561513257600080fd5b82516001600160401b0381111561514857600080fd5b8301601f8101851361515957600080fd5b80516151676148088261478a565b81815260059190911b8201830190838101908783111561518657600080fd5b928401925b8284101561507057835161519e8161439a565b8252928401929084019061518b565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006151f66040830184866151ad565b95945050505050565b60006020828403121561521157600080fd5b81516001600160c01b038116811461395a57600080fd5b60006020828403121561523a57600080fd5b815161395a8161439a565b600060ff821660ff81141561525c5761525c6150a4565b60010192915050565b6040815260006152796040830185876151ad565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156152d6578451835293830193918301916001016152ba565b5090979650505050505050565b6000602082840312156152f557600080fd5b815161395a8161475e565b600082821015615312576153126150a4565b500390565b60006020828403121561532957600080fd5b5051919050565b60008219821115615343576153436150a4565b500190565b60006020828403121561535a57600080fd5b815167ffffffffffffffff198116811461395a57600080fd5b60006001600160601b0383811690831681811015615393576153936150a4565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156153d6578151855293820193908201906001016153ba565b5092979650505050505050565b6020815263ffffffff825116602082015260018060a01b036020830151166040820152604082015160608201526000606083015161542c60808401826001600160a01b03169052565b50608083015160a083015260a083015161544e60c084018263ffffffff169052565b5060c08301516101008060e085015261546b610120850183614d05565b915060e0850151613d658286018263ffffffff169052565b63ffffffff8c811682528b811660208301526001600160a01b038b811660408401528a81166060840152898116608084015260a08301899052871660c083015260e08201869052841661010082015261014061012082018190526000906154ed83820185876151ad565b9e9d5050505050505050505050505050565b600063ffffffff80831681851680830382111561551e5761551e6150a4565b01949350505050565b600181811c9082168061553b57607f821691505b6020821081141561555c57634e487b7160e01b600052602260045260246000fd5b50919050565b6000808335601e1984360301811261557957600080fd5b8301803591506001600160401b0382111561559357600080fd5b60200191503681900382131561459857600080fd5b6000808335601e198436030181126155bf57600080fd5b83016020810192503590506001600160401b038111156155de57600080fd5b80360383131561459857600080fd5b60208152600082356155fe8161439a565b63ffffffff811660208401525061561760208401614c27565b6001600160a01b0381166040840152506040830135606083015261563d60608401614c27565b6001600160a01b038116608084015250608083013560a083015261566360a084016143ac565b63ffffffff811660c08401525061567d60c08401846155a8565b6101008060e0860152615695610120860183856151ad565b92506156a360e087016143ac565b63ffffffff8116868301529150613d65565b80356156c08161439a565b63ffffffff16825260208101356156d68161452c565b8015156020840152505050565b6040810161395d82846156b5565b60006001600160601b0380831681851681830481118215151615615717576157176150a4565b02949350505050565b600081600019048311821515161561573a5761573a6150a4565b500290565b6080810161574d82856156b5565b63ffffffff8351166040830152602083015160608301529392505050565b600061ffff80831681811415615783576157836150a4565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212201362a7df25093d424e5ea74e390f8636a3921ee6a249cdbbea6beedc01dae7ab64736f6c634300080c0033",
}

// ContractSettlementABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSettlementMetaData.ABI instead.
var ContractSettlementABI = ContractSettlementMetaData.ABI

// ContractSettlementBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSettlementMetaData.Bin instead.
var ContractSettlementBin = ContractSettlementMetaData.Bin

// DeployContractSettlement deploys a new Ethereum contract, binding an instance of ContractSettlement to it.
func DeployContractSettlement(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address) (common.Address, *types.Transaction, *ContractSettlement, error) {
	parsed, err := ContractSettlementMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSettlementBin), backend, _registryCoordinator)
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

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSettlement *ContractSettlementCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSettlement *ContractSettlementSession) Aggregator() (common.Address, error) {
	return _ContractSettlement.Contract.Aggregator(&_ContractSettlement.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSettlement *ContractSettlementCallerSession) Aggregator() (common.Address, error) {
	return _ContractSettlement.Contract.Aggregator(&_ContractSettlement.CallOpts)
}

// AllOrderDetails is a free data retrieval call binding the contract method 0xaa4f5cc6.
//
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers, uint32 createdBlock)
func (_ContractSettlement *ContractSettlementCaller) AllOrderDetails(opts *bind.CallOpts, arg0 uint32) (struct {
	OrderId                   uint32
	InputToken                common.Address
	InputAmount               *big.Int
	OutputToken               common.Address
	OutputAmount              *big.Int
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	CreatedBlock              uint32
}, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "allOrderDetails", arg0)

	outstruct := new(struct {
		OrderId                   uint32
		InputToken                common.Address
		InputAmount               *big.Int
		OutputToken               common.Address
		OutputAmount              *big.Int
		QuorumThresholdPercentage uint32
		QuorumNumbers             []byte
		CreatedBlock              uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.InputToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.InputAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OutputToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.OutputAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.QuorumThresholdPercentage = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.QuorumNumbers = *abi.ConvertType(out[6], new([]byte)).(*[]byte)
	outstruct.CreatedBlock = *abi.ConvertType(out[7], new(uint32)).(*uint32)

	return *outstruct, err

}

// AllOrderDetails is a free data retrieval call binding the contract method 0xaa4f5cc6.
//
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers, uint32 createdBlock)
func (_ContractSettlement *ContractSettlementSession) AllOrderDetails(arg0 uint32) (struct {
	OrderId                   uint32
	InputToken                common.Address
	InputAmount               *big.Int
	OutputToken               common.Address
	OutputAmount              *big.Int
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	CreatedBlock              uint32
}, error) {
	return _ContractSettlement.Contract.AllOrderDetails(&_ContractSettlement.CallOpts, arg0)
}

// AllOrderDetails is a free data retrieval call binding the contract method 0xaa4f5cc6.
//
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers, uint32 createdBlock)
func (_ContractSettlement *ContractSettlementCallerSession) AllOrderDetails(arg0 uint32) (struct {
	OrderId                   uint32
	InputToken                common.Address
	InputAmount               *big.Int
	OutputToken               common.Address
	OutputAmount              *big.Int
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	CreatedBlock              uint32
}, error) {
	return _ContractSettlement.Contract.AllOrderDetails(&_ContractSettlement.CallOpts, arg0)
}

// AllOrderHashes is a free data retrieval call binding the contract method 0xdd3285b3.
//
// Solidity: function allOrderHashes(uint32 ) view returns(bytes32)
func (_ContractSettlement *ContractSettlementCaller) AllOrderHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "allOrderHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllOrderHashes is a free data retrieval call binding the contract method 0xdd3285b3.
//
// Solidity: function allOrderHashes(uint32 ) view returns(bytes32)
func (_ContractSettlement *ContractSettlementSession) AllOrderHashes(arg0 uint32) ([32]byte, error) {
	return _ContractSettlement.Contract.AllOrderHashes(&_ContractSettlement.CallOpts, arg0)
}

// AllOrderHashes is a free data retrieval call binding the contract method 0xdd3285b3.
//
// Solidity: function allOrderHashes(uint32 ) view returns(bytes32)
func (_ContractSettlement *ContractSettlementCallerSession) AllOrderHashes(arg0 uint32) ([32]byte, error) {
	return _ContractSettlement.Contract.AllOrderHashes(&_ContractSettlement.CallOpts, arg0)
}

// AllOrderResponses is a free data retrieval call binding the contract method 0xe772935f.
//
// Solidity: function allOrderResponses(uint32 ) view returns(bytes32)
func (_ContractSettlement *ContractSettlementCaller) AllOrderResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "allOrderResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllOrderResponses is a free data retrieval call binding the contract method 0xe772935f.
//
// Solidity: function allOrderResponses(uint32 ) view returns(bytes32)
func (_ContractSettlement *ContractSettlementSession) AllOrderResponses(arg0 uint32) ([32]byte, error) {
	return _ContractSettlement.Contract.AllOrderResponses(&_ContractSettlement.CallOpts, arg0)
}

// AllOrderResponses is a free data retrieval call binding the contract method 0xe772935f.
//
// Solidity: function allOrderResponses(uint32 ) view returns(bytes32)
func (_ContractSettlement *ContractSettlementCallerSession) AllOrderResponses(arg0 uint32) ([32]byte, error) {
	return _ContractSettlement.Contract.AllOrderResponses(&_ContractSettlement.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSettlement *ContractSettlementCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSettlement *ContractSettlementSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSettlement.Contract.BlsApkRegistry(&_ContractSettlement.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSettlement *ContractSettlementCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSettlement.Contract.BlsApkRegistry(&_ContractSettlement.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSettlement *ContractSettlementCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSettlement *ContractSettlementSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractSettlement.Contract.CheckSignatures(&_ContractSettlement.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSettlement *ContractSettlementCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractSettlement.Contract.CheckSignatures(&_ContractSettlement.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSettlement *ContractSettlementCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSettlement *ContractSettlementSession) Delegation() (common.Address, error) {
	return _ContractSettlement.Contract.Delegation(&_ContractSettlement.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSettlement *ContractSettlementCallerSession) Delegation() (common.Address, error) {
	return _ContractSettlement.Contract.Delegation(&_ContractSettlement.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSettlement *ContractSettlementCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSettlement *ContractSettlementSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractSettlement.Contract.GetCheckSignaturesIndices(&_ContractSettlement.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSettlement *ContractSettlementCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractSettlement.Contract.GetCheckSignaturesIndices(&_ContractSettlement.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSettlement *ContractSettlementCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSettlement *ContractSettlementSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractSettlement.Contract.GetOperatorState(&_ContractSettlement.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSettlement *ContractSettlementCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractSettlement.Contract.GetOperatorState(&_ContractSettlement.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSettlement *ContractSettlementCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSettlement *ContractSettlementSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractSettlement.Contract.GetOperatorState0(&_ContractSettlement.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSettlement *ContractSettlementCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractSettlement.Contract.GetOperatorState0(&_ContractSettlement.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractSettlement *ContractSettlementCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractSettlement *ContractSettlementSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractSettlement.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractSettlement.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractSettlement *ContractSettlementCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractSettlement.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractSettlement.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// LatestOrderNum is a free data retrieval call binding the contract method 0xcfe3074e.
//
// Solidity: function latestOrderNum() view returns(uint32)
func (_ContractSettlement *ContractSettlementCaller) LatestOrderNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "latestOrderNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestOrderNum is a free data retrieval call binding the contract method 0xcfe3074e.
//
// Solidity: function latestOrderNum() view returns(uint32)
func (_ContractSettlement *ContractSettlementSession) LatestOrderNum() (uint32, error) {
	return _ContractSettlement.Contract.LatestOrderNum(&_ContractSettlement.CallOpts)
}

// LatestOrderNum is a free data retrieval call binding the contract method 0xcfe3074e.
//
// Solidity: function latestOrderNum() view returns(uint32)
func (_ContractSettlement *ContractSettlementCallerSession) LatestOrderNum() (uint32, error) {
	return _ContractSettlement.Contract.LatestOrderNum(&_ContractSettlement.CallOpts)
}

// OrderNumber is a free data retrieval call binding the contract method 0xb04dd81b.
//
// Solidity: function orderNumber() view returns(uint32)
func (_ContractSettlement *ContractSettlementCaller) OrderNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "orderNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OrderNumber is a free data retrieval call binding the contract method 0xb04dd81b.
//
// Solidity: function orderNumber() view returns(uint32)
func (_ContractSettlement *ContractSettlementSession) OrderNumber() (uint32, error) {
	return _ContractSettlement.Contract.OrderNumber(&_ContractSettlement.CallOpts)
}

// OrderNumber is a free data retrieval call binding the contract method 0xb04dd81b.
//
// Solidity: function orderNumber() view returns(uint32)
func (_ContractSettlement *ContractSettlementCallerSession) OrderNumber() (uint32, error) {
	return _ContractSettlement.Contract.OrderNumber(&_ContractSettlement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSettlement *ContractSettlementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSettlement *ContractSettlementSession) Owner() (common.Address, error) {
	return _ContractSettlement.Contract.Owner(&_ContractSettlement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSettlement *ContractSettlementCallerSession) Owner() (common.Address, error) {
	return _ContractSettlement.Contract.Owner(&_ContractSettlement.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSettlement *ContractSettlementCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSettlement *ContractSettlementSession) Paused(index uint8) (bool, error) {
	return _ContractSettlement.Contract.Paused(&_ContractSettlement.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSettlement *ContractSettlementCallerSession) Paused(index uint8) (bool, error) {
	return _ContractSettlement.Contract.Paused(&_ContractSettlement.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSettlement *ContractSettlementCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSettlement *ContractSettlementSession) Paused0() (*big.Int, error) {
	return _ContractSettlement.Contract.Paused0(&_ContractSettlement.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSettlement *ContractSettlementCallerSession) Paused0() (*big.Int, error) {
	return _ContractSettlement.Contract.Paused0(&_ContractSettlement.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSettlement *ContractSettlementCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSettlement *ContractSettlementSession) PauserRegistry() (common.Address, error) {
	return _ContractSettlement.Contract.PauserRegistry(&_ContractSettlement.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSettlement *ContractSettlementCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractSettlement.Contract.PauserRegistry(&_ContractSettlement.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSettlement *ContractSettlementCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSettlement *ContractSettlementSession) RegistryCoordinator() (common.Address, error) {
	return _ContractSettlement.Contract.RegistryCoordinator(&_ContractSettlement.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSettlement *ContractSettlementCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractSettlement.Contract.RegistryCoordinator(&_ContractSettlement.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSettlement *ContractSettlementCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSettlement *ContractSettlementSession) StakeRegistry() (common.Address, error) {
	return _ContractSettlement.Contract.StakeRegistry(&_ContractSettlement.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSettlement *ContractSettlementCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractSettlement.Contract.StakeRegistry(&_ContractSettlement.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSettlement *ContractSettlementCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSettlement *ContractSettlementSession) StaleStakesForbidden() (bool, error) {
	return _ContractSettlement.Contract.StaleStakesForbidden(&_ContractSettlement.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSettlement *ContractSettlementCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractSettlement.Contract.StaleStakesForbidden(&_ContractSettlement.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSettlement *ContractSettlementCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSettlement *ContractSettlementSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractSettlement.Contract.TrySignatureAndApkVerification(&_ContractSettlement.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSettlement *ContractSettlementCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractSettlement.Contract.TrySignatureAndApkVerification(&_ContractSettlement.CallOpts, msgHash, apk, apkG2, sigma)
}

// Fulfill is a paid mutator transaction binding the contract method 0x70acd7e9.
//
// Solidity: function fulfill(uint32 orderId, address maker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractSettlement *ContractSettlementTransactor) Fulfill(opts *bind.TransactOpts, orderId uint32, maker common.Address, inputToken common.Address, inputAmount *big.Int, outputToken common.Address, outputAmount *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "fulfill", orderId, maker, inputToken, inputAmount, outputToken, outputAmount, quorumThresholdPercentage, quorumNumbers)
}

// Fulfill is a paid mutator transaction binding the contract method 0x70acd7e9.
//
// Solidity: function fulfill(uint32 orderId, address maker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractSettlement *ContractSettlementSession) Fulfill(orderId uint32, maker common.Address, inputToken common.Address, inputAmount *big.Int, outputToken common.Address, outputAmount *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, orderId, maker, inputToken, inputAmount, outputToken, outputAmount, quorumThresholdPercentage, quorumNumbers)
}

// Fulfill is a paid mutator transaction binding the contract method 0x70acd7e9.
//
// Solidity: function fulfill(uint32 orderId, address maker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) Fulfill(orderId uint32, maker common.Address, inputToken common.Address, inputAmount *big.Int, outputToken common.Address, outputAmount *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, orderId, maker, inputToken, inputAmount, outputToken, outputAmount, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractSettlement *ContractSettlementTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractSettlement *ContractSettlementSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Initialize(&_ContractSettlement.TransactOpts, _pauserRegistry, initialOwner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Initialize(&_ContractSettlement.TransactOpts, _pauserRegistry, initialOwner, _aggregator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSettlement *ContractSettlementTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSettlement *ContractSettlementSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Pause(&_ContractSettlement.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Pause(&_ContractSettlement.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSettlement *ContractSettlementTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSettlement *ContractSettlementSession) PauseAll() (*types.Transaction, error) {
	return _ContractSettlement.Contract.PauseAll(&_ContractSettlement.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSettlement *ContractSettlementTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractSettlement.Contract.PauseAll(&_ContractSettlement.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSettlement *ContractSettlementTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSettlement *ContractSettlementSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSettlement.Contract.RenounceOwnership(&_ContractSettlement.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSettlement *ContractSettlementTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSettlement.Contract.RenounceOwnership(&_ContractSettlement.TransactOpts)
}

// RespondToFulfill is a paid mutator transaction binding the contract method 0xfd3d73bd.
//
// Solidity: function respondToFulfill((uint32,address,uint256,address,uint256,uint32,bytes,uint32) order, (uint32,bool) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSettlement *ContractSettlementTransactor) RespondToFulfill(opts *bind.TransactOpts, order SettlementOrder, orderResponse SettlementOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "respondToFulfill", order, orderResponse, nonSignerStakesAndSignature)
}

// RespondToFulfill is a paid mutator transaction binding the contract method 0xfd3d73bd.
//
// Solidity: function respondToFulfill((uint32,address,uint256,address,uint256,uint32,bytes,uint32) order, (uint32,bool) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSettlement *ContractSettlementSession) RespondToFulfill(order SettlementOrder, orderResponse SettlementOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSettlement.Contract.RespondToFulfill(&_ContractSettlement.TransactOpts, order, orderResponse, nonSignerStakesAndSignature)
}

// RespondToFulfill is a paid mutator transaction binding the contract method 0xfd3d73bd.
//
// Solidity: function respondToFulfill((uint32,address,uint256,address,uint256,uint32,bytes,uint32) order, (uint32,bool) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) RespondToFulfill(order SettlementOrder, orderResponse SettlementOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSettlement.Contract.RespondToFulfill(&_ContractSettlement.TransactOpts, order, orderResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSettlement *ContractSettlementTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSettlement *ContractSettlementSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSettlement.Contract.SetPauserRegistry(&_ContractSettlement.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSettlement.Contract.SetPauserRegistry(&_ContractSettlement.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSettlement *ContractSettlementTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSettlement *ContractSettlementSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractSettlement.Contract.SetStaleStakesForbidden(&_ContractSettlement.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractSettlement.Contract.SetStaleStakesForbidden(&_ContractSettlement.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSettlement *ContractSettlementTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSettlement *ContractSettlementSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSettlement.Contract.TransferOwnership(&_ContractSettlement.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSettlement.Contract.TransferOwnership(&_ContractSettlement.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSettlement *ContractSettlementTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSettlement *ContractSettlementSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Unpause(&_ContractSettlement.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Unpause(&_ContractSettlement.TransactOpts, newPausedStatus)
}

// ContractSettlementInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractSettlement contract.
type ContractSettlementInitializedIterator struct {
	Event *ContractSettlementInitialized // Event containing the contract specifics and raw log

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
func (it *ContractSettlementInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementInitialized)
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
		it.Event = new(ContractSettlementInitialized)
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
func (it *ContractSettlementInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementInitialized represents a Initialized event raised by the ContractSettlement contract.
type ContractSettlementInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSettlement *ContractSettlementFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractSettlementInitializedIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementInitializedIterator{contract: _ContractSettlement.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSettlement *ContractSettlementFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSettlementInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractSettlement.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementInitialized)
				if err := _ContractSettlement.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractSettlement *ContractSettlementFilterer) ParseInitialized(log types.Log) (*ContractSettlementInitialized, error) {
	event := new(ContractSettlementInitialized)
	if err := _ContractSettlement.contract.UnpackLog(event, "Initialized", log); err != nil {
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
	OrderResponse         SettlementOrderResponse
	OrderResponseMetadata SettlementOrderResponseMetadata
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterOrderRespondedEvent is a free log retrieval operation binding the contract event 0xfc76f8e0b77f0d3f2776538c1e70df0b30f8dce6075db04d3ed8bb6f3d4e5f96.
//
// Solidity: event OrderRespondedEvent((uint32,bool) orderResponse, (uint32,bytes32) OrderResponseMetadata)
func (_ContractSettlement *ContractSettlementFilterer) FilterOrderRespondedEvent(opts *bind.FilterOpts) (*ContractSettlementOrderRespondedEventIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "OrderRespondedEvent")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementOrderRespondedEventIterator{contract: _ContractSettlement.contract, event: "OrderRespondedEvent", logs: logs, sub: sub}, nil
}

// WatchOrderRespondedEvent is a free log subscription operation binding the contract event 0xfc76f8e0b77f0d3f2776538c1e70df0b30f8dce6075db04d3ed8bb6f3d4e5f96.
//
// Solidity: event OrderRespondedEvent((uint32,bool) orderResponse, (uint32,bytes32) OrderResponseMetadata)
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

// ParseOrderRespondedEvent is a log parse operation binding the contract event 0xfc76f8e0b77f0d3f2776538c1e70df0b30f8dce6075db04d3ed8bb6f3d4e5f96.
//
// Solidity: event OrderRespondedEvent((uint32,bool) orderResponse, (uint32,bytes32) OrderResponseMetadata)
func (_ContractSettlement *ContractSettlementFilterer) ParseOrderRespondedEvent(log types.Log) (*ContractSettlementOrderRespondedEvent, error) {
	event := new(ContractSettlementOrderRespondedEvent)
	if err := _ContractSettlement.contract.UnpackLog(event, "OrderRespondedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractSettlement contract.
type ContractSettlementOwnershipTransferredIterator struct {
	Event *ContractSettlementOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractSettlementOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementOwnershipTransferred)
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
		it.Event = new(ContractSettlementOwnershipTransferred)
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
func (it *ContractSettlementOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementOwnershipTransferred represents a OwnershipTransferred event raised by the ContractSettlement contract.
type ContractSettlementOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSettlement *ContractSettlementFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractSettlementOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementOwnershipTransferredIterator{contract: _ContractSettlement.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSettlement *ContractSettlementFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractSettlementOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSettlement.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementOwnershipTransferred)
				if err := _ContractSettlement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractSettlement *ContractSettlementFilterer) ParseOwnershipTransferred(log types.Log) (*ContractSettlementOwnershipTransferred, error) {
	event := new(ContractSettlementOwnershipTransferred)
	if err := _ContractSettlement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractSettlement contract.
type ContractSettlementPausedIterator struct {
	Event *ContractSettlementPaused // Event containing the contract specifics and raw log

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
func (it *ContractSettlementPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementPaused)
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
		it.Event = new(ContractSettlementPaused)
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
func (it *ContractSettlementPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementPaused represents a Paused event raised by the ContractSettlement contract.
type ContractSettlementPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSettlement *ContractSettlementFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractSettlementPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementPausedIterator{contract: _ContractSettlement.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSettlement *ContractSettlementFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractSettlementPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSettlement.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementPaused)
				if err := _ContractSettlement.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSettlement *ContractSettlementFilterer) ParsePaused(log types.Log) (*ContractSettlementPaused, error) {
	event := new(ContractSettlementPaused)
	if err := _ContractSettlement.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractSettlement contract.
type ContractSettlementPauserRegistrySetIterator struct {
	Event *ContractSettlementPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractSettlementPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementPauserRegistrySet)
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
		it.Event = new(ContractSettlementPauserRegistrySet)
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
func (it *ContractSettlementPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementPauserRegistrySet represents a PauserRegistrySet event raised by the ContractSettlement contract.
type ContractSettlementPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSettlement *ContractSettlementFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractSettlementPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementPauserRegistrySetIterator{contract: _ContractSettlement.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSettlement *ContractSettlementFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractSettlementPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractSettlement.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementPauserRegistrySet)
				if err := _ContractSettlement.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSettlement *ContractSettlementFilterer) ParsePauserRegistrySet(log types.Log) (*ContractSettlementPauserRegistrySet, error) {
	event := new(ContractSettlementPauserRegistrySet)
	if err := _ContractSettlement.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractSettlement contract.
type ContractSettlementStaleStakesForbiddenUpdateIterator struct {
	Event *ContractSettlementStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractSettlementStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractSettlementStaleStakesForbiddenUpdate)
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
func (it *ContractSettlementStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractSettlement contract.
type ContractSettlementStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSettlement *ContractSettlementFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractSettlementStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementStaleStakesForbiddenUpdateIterator{contract: _ContractSettlement.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSettlement *ContractSettlementFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractSettlementStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractSettlement.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementStaleStakesForbiddenUpdate)
				if err := _ContractSettlement.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSettlement *ContractSettlementFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractSettlementStaleStakesForbiddenUpdate, error) {
	event := new(ContractSettlementStaleStakesForbiddenUpdate)
	if err := _ContractSettlement.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSettlementUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractSettlement contract.
type ContractSettlementUnpausedIterator struct {
	Event *ContractSettlementUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractSettlementUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementUnpaused)
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
		it.Event = new(ContractSettlementUnpaused)
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
func (it *ContractSettlementUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementUnpaused represents a Unpaused event raised by the ContractSettlement contract.
type ContractSettlementUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSettlement *ContractSettlementFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractSettlementUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementUnpausedIterator{contract: _ContractSettlement.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSettlement *ContractSettlementFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractSettlementUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSettlement.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementUnpaused)
				if err := _ContractSettlement.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSettlement *ContractSettlementFilterer) ParseUnpaused(log types.Log) (*ContractSettlementUnpaused, error) {
	event := new(ContractSettlementUnpaused)
	if err := _ContractSettlement.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	OrderId                   uint32
	OrderNum                  uint32
	Maker                     common.Address
	Taker                     common.Address
	InputToken                common.Address
	InputAmount               *big.Int
	OutputToken               common.Address
	OutputAmount              *big.Int
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterFulfillEvent is a free log retrieval operation binding the contract event 0x6f375196052db01a32a94e2166b4d273355741c797cb8eb067534c606f1a50fa.
//
// Solidity: event fulfillEvent(uint32 orderId, uint32 orderNum, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers)
func (_ContractSettlement *ContractSettlementFilterer) FilterFulfillEvent(opts *bind.FilterOpts) (*ContractSettlementFulfillEventIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "fulfillEvent")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementFulfillEventIterator{contract: _ContractSettlement.contract, event: "fulfillEvent", logs: logs, sub: sub}, nil
}

// WatchFulfillEvent is a free log subscription operation binding the contract event 0x6f375196052db01a32a94e2166b4d273355741c797cb8eb067534c606f1a50fa.
//
// Solidity: event fulfillEvent(uint32 orderId, uint32 orderNum, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers)
func (_ContractSettlement *ContractSettlementFilterer) WatchFulfillEvent(opts *bind.WatchOpts, sink chan<- *ContractSettlementFulfillEvent) (event.Subscription, error) {

	logs, sub, err := _ContractSettlement.contract.WatchLogs(opts, "fulfillEvent")
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
				if err := _ContractSettlement.contract.UnpackLog(event, "fulfillEvent", log); err != nil {
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

// ParseFulfillEvent is a log parse operation binding the contract event 0x6f375196052db01a32a94e2166b4d273355741c797cb8eb067534c606f1a50fa.
//
// Solidity: event fulfillEvent(uint32 orderId, uint32 orderNum, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers)
func (_ContractSettlement *ContractSettlementFilterer) ParseFulfillEvent(log types.Log) (*ContractSettlementFulfillEvent, error) {
	event := new(ContractSettlementFulfillEvent)
	if err := _ContractSettlement.contract.UnpackLog(event, "fulfillEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

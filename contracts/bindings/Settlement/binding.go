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
	Maker                     common.Address
	Taker                     common.Address
	InputToken                common.Address
	InputAmount               *big.Int
	OutputToken               common.Address
	OutputAmount              *big.Int
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	Expiry                    *big.Int
	TargetNetworkNumber       uint32
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderDetails\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfill\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structSettlement.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToFulfill\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structSettlement.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"orderResponse\",\"type\":\"tuple\",\"internalType\":\"structSettlement.OrderResponse\",\"components\":[{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"txSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderRespondedEvent\",\"inputs\":[{\"name\":\"orderResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSettlement.OrderResponse\",\"components\":[{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"txSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"OrderResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSettlement.OrderResponseMetadata\",\"components\":[{\"name\":\"orderResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"fulfillEvent\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSettlement.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TakerMismatch\",\"inputs\":[]}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162006199380380620061998339810160408190526200003591620001ed565b80806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001ed565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620001ed565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b39190620001ed565b6001600160a01b031660e05250506097805460ff1916600117905562000214565b6001600160a01b0381168114620001ea57600080fd5b50565b6000602082840312156200020057600080fd5b81516200020d81620001d4565b9392505050565b60805160a05160c05160e051615f13620002866000396000818161047c015261230f01526000818161032e01526124f1015260008181610355015281816126c7015261288901526000818161037c01528181610d1601528181611fda0152818161217201526123ac0152615f136000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80636d14a98711610104578063c0c53b8b116100a2578063e772935f11610071578063e772935f1461049e578063ec381235146104be578063f2fde38b146104d1578063fabc1cbc146104e457600080fd5b8063c0c53b8b14610423578063cefdc1d414610436578063dd3285b314610457578063df5cf7231461047757600080fd5b8063886f1195116100de578063886f1195146103c75780638da5cb5b146103da578063aa4f5cc6146103eb578063b98d09081461041657600080fd5b80636d14a987146103775780636efb46361461039e578063715018a6146103bf57600080fd5b80634f739f74116101715780635c1556621161014b5780635c155662146102f75780635c975abb146103175780635df4594614610329578063683048351461035057600080fd5b80634f739f741461029c578063595c6a67146102bc5780635ac86ab7146102c457600080fd5b8063245a7bfc116101ad578063245a7bfc1461022b5780633563b0d114610256578063416c7e5e1461027657806347f1385e1461028957600080fd5b806310d67a2f146101d4578063136439dd146101e9578063171f1d5b146101fc575b600080fd5b6101e76101e2366004614706565b6104f7565b005b6101e76101f7366004614723565b6105b3565b61020f61020a3660046148c4565b6106f2565b6040805192151583529015156020830152015b60405180910390f35b60cc5461023e906001600160a01b031681565b6040516001600160a01b039091168152602001610222565b6102696102643660046149a1565b61087c565b6040516102229190614aa8565b6101e7610284366004614ac9565b610d14565b6101e7610297366004614d8a565b610e89565b6102af6102aa366004614e54565b611272565b6040516102229190614f58565b6101e7611998565b6102e76102d2366004615022565b606654600160ff9092169190911b9081161490565b6040519015158152602001610222565b61030a61030536600461503f565b611a5f565b60405161022291906150f4565b6066545b604051908152602001610222565b61023e7f000000000000000000000000000000000000000000000000000000000000000081565b61023e7f000000000000000000000000000000000000000000000000000000000000000081565b61023e7f000000000000000000000000000000000000000000000000000000000000000081565b6103b16103ac366004615138565b611c27565b6040516102229291906151f8565b6101e7612b3e565b60655461023e906001600160a01b031681565b6033546001600160a01b031661023e565b6103fe6103f9366004615241565b612b52565b6040516102229c9b9a999897969594939291906152b6565b6097546102e79060ff1681565b6101e7610431366004615351565b612c5f565b610449610444366004615391565b612d9a565b6040516102229291906153c8565b61031b610465366004615241565b60c96020526000908152604090205481565b61023e7f000000000000000000000000000000000000000000000000000000000000000081565b61031b6104ac366004615241565b60cb6020526000908152604090205481565b6101e76104cc3660046153e9565b612f2c565b6101e76104df366004614706565b6133e8565b6101e76104f2366004614723565b61345e565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561054a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056e919061552c565b6001600160a01b0316336001600160a01b0316146105a75760405162461bcd60e51b815260040161059e90615549565b60405180910390fd5b6105b0816135ba565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156105fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061f9190615593565b61063b5760405162461bcd60e51b815260040161059e906155b0565b606654818116146106b45760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161059e565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061073a5761073a6155f8565b60200201518951600160200201518a6020015160006002811061075f5761075f6155f8565b60200201518b6020015160016002811061077b5761077b6155f8565b602090810291909101518c518d8301516040516107d89a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6107fb919061560e565b905061086e61081461080d88846136b1565b8690613748565b61081c6137dc565b6108646108558561084f604080518082018252600080825260209182015281518083019092526001825260029082015290565b906136b1565b61085e8c61389c565b90613748565b886201d4c061392c565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e2919061552c565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610924573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610948919061552c565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561098a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ae919061552c565b9050600086516001600160401b038111156109cb576109cb61473c565b6040519080825280602002602001820160405280156109fe57816020015b60608152602001906001900390816109e95790505b50905060005b8751811015610d06576000888281518110610a2157610a216155f8565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610a82573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610aaa9190810190615630565b905080516001600160401b03811115610ac557610ac561473c565b604051908082528060200260200182016040528015610b1057816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610ae35790505b50848481518110610b2357610b236155f8565b602002602001018190525060005b8151811015610cf0576040518060600160405280876001600160a01b03166347b314e8858581518110610b6657610b666155f8565b60200260200101516040518263ffffffff1660e01b8152600401610b8c91815260200190565b602060405180830381865afa158015610ba9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bcd919061552c565b6001600160a01b03168152602001838381518110610bed57610bed6155f8565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610c1b57610c1b6155f8565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610c77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c9b91906156c0565b6001600160601b0316815250858581518110610cb957610cb96155f8565b60200260200101518281518110610cd257610cd26155f8565b60200260200101819052508080610ce8906156ff565b915050610b31565b5050508080610cfe906156ff565b915050610a04565b5093505050505b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d96919061552c565b6001600160a01b0316336001600160a01b031614610e425760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a40161059e565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b60cc546001600160a01b03163314610ee35760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c6572000000604482015260640161059e565b366000610ef461010086018661571a565b90925090506000610f0c610100870160e08801615241565b90506000610f2261018088016101608901615241565b905060c96000610f356020890189615241565b63ffffffff1663ffffffff1681526020019081526020016000205487604051602001610f6191906157ce565b6040516020818303038152906040528051906020012014610fea5760405162461bcd60e51b815260206004820152603e60248201527f737570706c696564206f7264657220646f6573206e6f74206d6174636820746860448201527f65206f6e65207265636f7264656420696e2074686520636f6e74726163740000606482015260840161059e565b600060cb81610ffc60208a018a615241565b63ffffffff1663ffffffff168152602001908152602001600020541461107a5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526c103a37903a34329037b93232b960991b606482015260840161059e565b60008660405160200161108d919061593a565b6040516020818303038152906040528051906020012090506000806110b5838888878c611c27565b9150915060005b868110156111b4578560ff16836020015182815181106110de576110de6155f8565b60200260200101516110f09190615948565b6001600160601b0316606484600001518381518110611111576111116155f8565b60200260200101516001600160601b031661112c9190615977565b10156111a2576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d606482015260840161059e565b806111ac816156ff565b9150506110bc565b5060408051808201825263ffffffff431681526020808201849052915190916111e1918c91849101615996565b6040516020818303038152906040528051906020012060cb60008c600001602081019061120e9190615241565b63ffffffff1663ffffffff168152602001908152602001600020819055507ffc76f8e0b77f0d3f2776538c1e70df0b30f8dce6075db04d3ed8bb6f3d4e5f968a8260405161125d929190615996565b60405180910390a15050505050505050505050565b61129d6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611301919061552c565b905061132e6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061135e908b90899089906004016159c2565b600060405180830381865afa15801561137b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526113a39190810190615a0c565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906113d5908b908b908b90600401615a9a565b600060405180830381865afa1580156113f2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261141a9190810190615a0c565b6040820152856001600160401b038111156114375761143761473c565b60405190808252806020026020018201604052801561146a57816020015b60608152602001906001900390816114555790505b50606082015260005b60ff81168711156118a9576000856001600160401b038111156114985761149861473c565b6040519080825280602002602001820160405280156114c1578160200160208202803683370190505b5083606001518360ff16815181106114db576114db6155f8565b602002602001018190525060005b868110156117a95760008c6001600160a01b03166304ec63518a8a85818110611514576115146155f8565b905060200201358e88600001518681518110611532576115326155f8565b60200260200101516040518463ffffffff1660e01b815260040161156f9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561158c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115b09190615ac3565b90506001600160c01b0381166116545760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a40161059e565b8a8a8560ff16818110611669576116696155f8565b6001600160c01b03841692013560f81c9190911c60019081161415905061179657856001600160a01b031663dd9846b98a8a858181106116ab576116ab6155f8565b905060200201358d8d8860ff168181106116c7576116c76155f8565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa15801561171d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117419190615aec565b85606001518560ff168151811061175a5761175a6155f8565b60200260200101518481518110611773576117736155f8565b63ffffffff9092166020928302919091019091015282611792816156ff565b9350505b50806117a1816156ff565b9150506114e9565b506000816001600160401b038111156117c4576117c461473c565b6040519080825280602002602001820160405280156117ed578160200160208202803683370190505b50905060005b8281101561186e5784606001518460ff1681518110611814576118146155f8565b6020026020010151818151811061182d5761182d6155f8565b6020026020010151828281518110611847576118476155f8565b63ffffffff9092166020928302919091019091015280611866816156ff565b9150506117f3565b508084606001518460ff1681518110611889576118896155f8565b6020026020010181905250505080806118a190615b09565b915050611473565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061190e919061552c565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611941908b908b908e90600401615b29565b600060405180830381865afa15801561195e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119869190810190615a0c565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156119e0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a049190615593565b611a205760405162461bcd60e51b815260040161059e906155b0565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611a91929190615b53565b600060405180830381865afa158015611aae573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ad69190810190615a0c565b9050600084516001600160401b03811115611af357611af361473c565b604051908082528060200260200182016040528015611b1c578160200160208202803683370190505b50905060005b8551811015611c1d57866001600160a01b03166304ec6351878381518110611b4c57611b4c6155f8565b602002602001015187868581518110611b6757611b676155f8565b60200260200101516040518463ffffffff1660e01b8152600401611ba49392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611bc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611be59190615ac3565b6001600160c01b0316828281518110611c0057611c006155f8565b602090810291909101015280611c15816156ff565b915050611b22565b5095945050505050565b6040805180820190915260608082526020820152600084611c9e5760405162461bcd60e51b81526020600482015260376024820152600080516020615ebe83398151915260448201527f7265733a20656d7074792071756f72756d20696e707574000000000000000000606482015260840161059e565b60408301515185148015611cb6575060a08301515185145b8015611cc6575060c08301515185145b8015611cd6575060e08301515185145b611d405760405162461bcd60e51b81526020600482015260416024820152600080516020615ebe83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a40161059e565b82515160208401515114611db85760405162461bcd60e51b815260206004820152604460248201819052600080516020615ebe833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a40161059e565b4363ffffffff168463ffffffff1610611e275760405162461bcd60e51b815260206004820152603c6024820152600080516020615ebe83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b00000000606482015260840161059e565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611e6857611e6861473c565b604051908082528060200260200182016040528015611e91578160200160208202803683370190505b506020820152866001600160401b03811115611eaf57611eaf61473c565b604051908082528060200260200182016040528015611ed8578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611f0c57611f0c61473c565b604051908082528060200260200182016040528015611f35578160200160208202803683370190505b5081526020860151516001600160401b03811115611f5557611f5561473c565b604051908082528060200260200182016040528015611f7e578160200160208202803683370190505b50816020018190525060006120508a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612027573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061204b9190615ba7565b613b50565b905060005b8760200151518110156122eb5761209a8860200151828151811061207b5761207b6155f8565b6020026020010151805160009081526020918201519091526040902090565b836020015182815181106120b0576120b06155f8565b602090810291909101015280156121705760208301516120d1600183615bc4565b815181106120e1576120e16155f8565b602002602001015160001c83602001518281518110612102576121026155f8565b602002602001015160001c11612170576040805162461bcd60e51b8152602060048201526024810191909152600080516020615ebe83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f72746564606482015260840161059e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106121b5576121b56155f8565b60200260200101518b8b6000015185815181106121d4576121d46155f8565b60200260200101516040518463ffffffff1660e01b81526004016122119392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561222e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122529190615ac3565b6001600160c01b031683600001518281518110612271576122716155f8565b6020026020010181815250506122d761080d6122ab848660000151858151811061229d5761229d6155f8565b602002602001015116613be3565b8a6020015184815181106122c1576122c16155f8565b6020026020010151613c0e90919063ffffffff16565b9450806122e3816156ff565b915050612055565b50506122f683613cf2565b60975490935060ff1660008161230d57600061238f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa15801561236b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061238f9190615bdb565b905060005b8a811015612a0d5782156124ef578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106123eb576123eb6155f8565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa15801561242b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061244f9190615bdb565b6124599190615bf4565b116124ef5760405162461bcd60e51b81526020600482015260666024820152600080516020615ebe83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c40161059e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612530576125306155f8565b9050013560f81c60f81b60f81c8c8c60a001518581518110612554576125546155f8565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156125b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125d49190615c0c565b6001600160401b0319166125f78a60400151838151811061207b5761207b6155f8565b67ffffffffffffffff1916146126935760405162461bcd60e51b81526020600482015260616024820152600080516020615ebe83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c40161059e565b6126c3896040015182815181106126ac576126ac6155f8565b60200260200101518761374890919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612706576127066155f8565b9050013560f81c60f81b60f81c8c8c60c00151858151811061272a5761272a6155f8565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612786573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127aa91906156c0565b856020015182815181106127c0576127c06155f8565b6001600160601b039092166020928302919091018201528501518051829081106127ec576127ec6155f8565b60200260200101518560000151828151811061280a5761280a6155f8565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156129f85761288286600001518281518110612854576128546155f8565b60200260200101518f8f8681811061286e5761286e6155f8565b600192013560f81c9290921c811614919050565b156129e6577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106128c8576128c86155f8565b9050013560f81c60f81b60f81c8e896020015185815181106128ec576128ec6155f8565b60200260200101518f60e00151888151811061290a5761290a6155f8565b60200260200101518781518110612923576129236155f8565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612987573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129ab91906156c0565b87518051859081106129bf576129bf6155f8565b602002602001018181516129d39190615c37565b6001600160601b03169052506001909101905b806129f0816156ff565b91505061282e565b50508080612a05906156ff565b915050612394565b505050600080612a278c868a606001518b608001516106f2565b9150915081612a985760405162461bcd60e51b81526020600482015260436024820152600080516020615ebe83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a40161059e565b80612af95760405162461bcd60e51b81526020600482015260396024820152600080516020615ebe83398151915260448201527f7265733a207369676e617475726520697320696e76616c696400000000000000606482015260840161059e565b50506000878260200151604051602001612b14929190615c5f565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612b46613d8d565b612b506000613de7565b565b60ca602052600090815260409020805460018201546002830154600384015460048501546005860154600687015460078801805463ffffffff808a169a640100000000909a046001600160a01b039081169a99811699988116989616959316929190612bbd90615ca7565b80601f0160208091040260200160405190810160405280929190818152602001828054612be990615ca7565b8015612c365780601f10612c0b57610100808354040283529160200191612c36565b820191906000526020600020905b815481529060010190602001808311612c1957829003601f168201915b50505050600883015460099093015491929163ffffffff8082169250640100000000909104168c565b600054610100900460ff1615808015612c7f5750600054600160ff909116105b80612c995750303b158015612c99575060005460ff166001145b612cfc5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161059e565b6000805460ff191660011790558015612d1f576000805461ff0019166101001790555b612d2a846000613e39565b612d3383613de7565b60cc80546001600160a01b0319166001600160a01b0384161790558015612d94576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612dd557612dd56155f8565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612e119088908690600401615b53565b600060405180830381865afa158015612e2e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612e569190810190615a0c565b600081518110612e6857612e686155f8565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612ed4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ef89190615ac3565b6001600160c01b031690506000612f0e82613f23565b905081612f1c8a838a61087c565b9550955050505050935093915050565b60408301516001600160a01b031615801590612f5e575082604001516001600160a01b0316336001600160a01b031614155b15612f7c57604051630d9d754f60e41b815260040160405180910390fd5b428361012001511015612fa2576040516362b439dd60e11b815260040160405180910390fd5b82516020808501516040808701516060808901516080808b015160a08c015160c08d01516101208e01516101408f0151985160e09c8d1b6001600160e01b03199081169c82019c909c5299871b6bffffffffffffffffffffffff1990811660248c015297871b881660388b015294861b8716604c8a01528886019290925290931b90931692850192909252609484015260b483015290921b1660d482015260009060d80160405160208183030381529060405280519060200120905060006130fb85602001516130bf846040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b86868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613fef92505050565b90508061311b57604051638baa579f60e01b815260040160405180910390fd5b6060850151602086015160808701516040516323b872dd60e01b81526001600160a01b03928316600482015233602482015260448101919091529116906323b872dd906064016020604051808303816000875af1158015613180573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131a49190615593565b5060a0850151602086015160c08701516040516323b872dd60e01b81523360048201526001600160a01b03928316602482015260448101919091529116906323b872dd906064016020604051808303816000875af115801561320a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061322e9190615593565b503360408087019190915263ffffffff431661016087015251613255908690602001615de2565b60408051808303601f190181529181528151602092830120875163ffffffff908116600090815260c985528381209290925588518116825260ca845290829020885181548a8601519184166001600160c01b0319909116176401000000006001600160a01b0392831602178255928901516001820180546001600160a01b031990811692861692909217905560608a0151600283018054831691861691909117905560808a0151600383015560a08a015160048301805490921694169390931790925560c0880151600583015560e088015160068301805463ffffffff1916919092161790556101008701518051889361335692600785019291019061457e565b506101208201516008820155610140820151600990910180546101609093015163ffffffff9081166401000000000267ffffffffffffffff1990941692169190911791909117905560e08501516101008601516040517f9a5d42e3ee6260573aeaa5309b575a6abe8be1d8490569b1400fba374bca74ef926133d9928992615df5565b60405180910390a15050505050565b6133f0613d8d565b6001600160a01b0381166134555760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161059e565b6105b081613de7565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134d5919061552c565b6001600160a01b0316336001600160a01b0316146135055760405162461bcd60e51b815260040161059e90615549565b6066541981196066541916146135835760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161059e565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016106e7565b6001600160a01b0381166136485760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161059e565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526136cd614602565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561370057613702565bfe5b50806137405760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b604482015260640161059e565b505092915050565b6040805180820190915260008082526020820152613764614620565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156137005750806137405760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b604482015260640161059e565b6137e461463e565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806138cc600080516020615e9e8339815191528661560e565b90505b6138d881614131565b9093509150600080516020615e9e833981519152828309831415613912576040805180820190915290815260208101919091529392505050565b600080516020615e9e8339815191526001820890506138cf565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061395e614663565b60005b6002811015613b23576000613977826006615977565b905084826002811061398b5761398b6155f8565b6020020151518361399d836000615bf4565b600c81106139ad576139ad6155f8565b60200201528482600281106139c4576139c46155f8565b602002015160200151838260016139db9190615bf4565b600c81106139eb576139eb6155f8565b6020020152838260028110613a0257613a026155f8565b6020020151515183613a15836002615bf4565b600c8110613a2557613a256155f8565b6020020152838260028110613a3c57613a3c6155f8565b6020020151516001602002015183613a55836003615bf4565b600c8110613a6557613a656155f8565b6020020152838260028110613a7c57613a7c6155f8565b602002015160200151600060028110613a9757613a976155f8565b602002015183613aa8836004615bf4565b600c8110613ab857613ab86155f8565b6020020152838260028110613acf57613acf6155f8565b602002015160200151600160028110613aea57613aea6155f8565b602002015183613afb836005615bf4565b600c8110613b0b57613b0b6155f8565b60200201525080613b1b816156ff565b915050613961565b50613b2c614682565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613b5c846141b3565b9050808360ff166001901b11613bda5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c756500606482015260840161059e565b90505b92915050565b6000805b8215613bdd57613bf8600184615bc4565b9092169180613c0681615e30565b915050613be7565b60408051808201909152600080825260208201526102008261ffff1610613c6a5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b604482015260640161059e565b8161ffff1660011415613c7e575081613bdd565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613ce757600161ffff871660ff83161c81161415613cca57613cc78484613748565b93505b613cd48384613748565b92506201fffe600192831b169101613c9a565b509195945050505050565b60408051808201909152600080825260208201528151158015613d1757506020820151155b15613d35575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615e9e8339815191528460200151613d68919061560e565b613d8090600080516020615e9e833981519152615bc4565b905292915050565b919050565b6033546001600160a01b03163314612b505760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161059e565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546001600160a01b0316158015613e5a57506001600160a01b03821615155b613edc5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161059e565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613f1f826135ba565b5050565b6060600080613f3184613be3565b61ffff166001600160401b03811115613f4c57613f4c61473c565b6040519080825280601f01601f191660200182016040528015613f76576020820181803683370190505b5090506000805b825182108015613f8e575061010081105b15613fe5576001811b935085841615613fd5578060f81b838381518110613fb757613fb76155f8565b60200101906001600160f81b031916908160001a9053508160010191505b613fde816156ff565b9050613f7d565b5090949350505050565b6000806000613ffe8585614340565b9092509050600081600481111561401757614017615e52565b1480156140355750856001600160a01b0316826001600160a01b0316145b1561404557600192505050610d0d565b600080876001600160a01b0316631626ba7e60e01b888860405160240161406d929190615e68565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516140ab9190615e81565b600060405180830381855afa9150503d80600081146140e6576040519150601f19603f3d011682016040523d82523d6000602084013e6140eb565b606091505b50915091508180156140fe575080516020145b801561412557508051630b135d3f60e11b906141239083016020908101908401615bdb565b145b98975050505050505050565b60008080600080516020615e9e8339815191526003600080516020615e9e83398151915286600080516020615e9e8339815191528889090908905060006141a7827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615e9e8339815191526143b0565b91959194509092505050565b60006101008251111561423c5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a40161059e565b815161424a57506000919050565b60008083600081518110614260576142606155f8565b0160200151600160f89190911c81901b92505b84518110156143375784818151811061428e5761428e6155f8565b0160200151600160f89190911c1b91508282116143235760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a40161059e565b91811791614330816156ff565b9050614273565b50909392505050565b6000808251604114156143775760208301516040840151606085015160001a61436b87828585614458565b945094505050506143a9565b8251604014156143a15760208301516040840151614396868383614545565b9350935050506143a9565b506000905060025b9250929050565b6000806143bb614682565b6143c36146a0565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561370057508261444d5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c757265000000000000604482015260640161059e565b505195945050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561448f575060009050600361453c565b8460ff16601b141580156144a757508460ff16601c14155b156144b8575060009050600461453c565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561450c573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166145355760006001925092505061453c565b9150600090505b94509492505050565b6000806001600160ff1b0383168161456260ff86901c601b615bf4565b905061457087828885614458565b935093505050935093915050565b82805461458a90615ca7565b90600052602060002090601f0160209004810192826145ac57600085556145f2565b82601f106145c557805160ff19168380011785556145f2565b828001600101855582156145f2579182015b828111156145f25782518255916020019190600101906145d7565b506145fe9291506146be565b5090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806146516146d3565b815260200161465e6146d3565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b5b808211156145fe57600081556001016146bf565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146105b057600080fd5b60006020828403121561471857600080fd5b8135613bda816146f1565b60006020828403121561473557600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156147745761477461473c565b60405290565b60405161010081016001600160401b03811182821017156147745761477461473c565b60405161018081016001600160401b03811182821017156147745761477461473c565b604051601f8201601f191681016001600160401b03811182821017156147e8576147e861473c565b604052919050565b60006040828403121561480257600080fd5b61480a614752565b9050813581526020820135602082015292915050565b600082601f83011261483157600080fd5b604051604081018181106001600160401b03821117156148535761485361473c565b806040525080604084018581111561486a57600080fd5b845b81811015613ce757803583526020928301920161486c565b60006080828403121561489657600080fd5b61489e614752565b90506148aa8383614820565b81526148b98360408401614820565b602082015292915050565b60008060008061012085870312156148db57600080fd5b843593506148ec86602087016147f0565b92506148fb8660608701614884565b915061490a8660e087016147f0565b905092959194509250565b600082601f83011261492657600080fd5b81356001600160401b0381111561493f5761493f61473c565b614952601f8201601f19166020016147c0565b81815284602083860101111561496757600080fd5b816020850160208301376000918101602001919091529392505050565b63ffffffff811681146105b057600080fd5b8035613d8881614984565b6000806000606084860312156149b657600080fd5b83356149c1816146f1565b925060208401356001600160401b038111156149dc57600080fd5b6149e886828701614915565b92505060408401356149f981614984565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614a9a578385038a52825180518087529087019087870190845b81811015614a8557835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614a41565b50509a87019a95505091850191600101614a23565b509298975050505050505050565b602081526000610d0d6020830184614a04565b80151581146105b057600080fd5b600060208284031215614adb57600080fd5b8135613bda81614abb565b60006001600160401b03821115614aff57614aff61473c565b5060051b60200190565b600082601f830112614b1a57600080fd5b81356020614b2f614b2a83614ae6565b6147c0565b82815260059290921b84018101918181019086841115614b4e57600080fd5b8286015b84811015614b72578035614b6581614984565b8352918301918301614b52565b509695505050505050565b600082601f830112614b8e57600080fd5b81356020614b9e614b2a83614ae6565b82815260069290921b84018101918181019086841115614bbd57600080fd5b8286015b84811015614b7257614bd388826147f0565b835291830191604001614bc1565b600082601f830112614bf257600080fd5b81356020614c02614b2a83614ae6565b82815260059290921b84018101918181019086841115614c2157600080fd5b8286015b84811015614b725780356001600160401b03811115614c445760008081fd5b614c528986838b0101614b09565b845250918301918301614c25565b60006101808284031215614c7357600080fd5b614c7b61477a565b905081356001600160401b0380821115614c9457600080fd5b614ca085838601614b09565b83526020840135915080821115614cb657600080fd5b614cc285838601614b7d565b60208401526040840135915080821115614cdb57600080fd5b614ce785838601614b7d565b6040840152614cf98560608601614884565b6060840152614d0b8560e086016147f0565b6080840152610120840135915080821115614d2557600080fd5b614d3185838601614b09565b60a0840152610140840135915080821115614d4b57600080fd5b614d5785838601614b09565b60c0840152610160840135915080821115614d7157600080fd5b50614d7e84828501614be1565b60e08301525092915050565b60008060008385036080811215614da057600080fd5b84356001600160401b0380821115614db757600080fd5b908601906101808289031215614dcc57600080fd5b8195506040601f1984011215614de157600080fd5b6020870194506060870135925080831115614dfb57600080fd5b5050614e0986828701614c60565b9150509250925092565b60008083601f840112614e2557600080fd5b5081356001600160401b03811115614e3c57600080fd5b6020830191508360208285010111156143a957600080fd5b60008060008060008060808789031215614e6d57600080fd5b8635614e78816146f1565b95506020870135614e8881614984565b945060408701356001600160401b0380821115614ea457600080fd5b614eb08a838b01614e13565b90965094506060890135915080821115614ec957600080fd5b818901915089601f830112614edd57600080fd5b813581811115614eec57600080fd5b8a60208260051b8501011115614f0157600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614f4d57815163ffffffff1687529582019590820190600101614f2b565b509495945050505050565b600060208083528351608082850152614f7460a0850182614f17565b905081850151601f1980868403016040870152614f918383614f17565b92506040870151915080868403016060870152614fae8383614f17565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156150055784878303018452614ff3828751614f17565b95880195938801939150600101614fd9565b509998505050505050505050565b60ff811681146105b057600080fd5b60006020828403121561503457600080fd5b8135613bda81615013565b60008060006060848603121561505457600080fd5b833561505f816146f1565b92506020848101356001600160401b0381111561507b57600080fd5b8501601f8101871361508c57600080fd5b803561509a614b2a82614ae6565b81815260059190911b820183019083810190898311156150b957600080fd5b928401925b828410156150d7578335825292840192908401906150be565b80965050505050506150eb60408501614996565b90509250925092565b6020808252825182820181905260009190848201906040850190845b8181101561512c57835183529284019291840191600101615110565b50909695505050505050565b60008060008060006080868803121561515057600080fd5b8535945060208601356001600160401b038082111561516e57600080fd5b61517a89838a01614e13565b90965094506040880135915061518f82614984565b909250606087013590808211156151a557600080fd5b506151b288828901614c60565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614f4d5781516001600160601b0316875295820195908201906001016151d3565b604081526000835160408084015261521360808401826151bf565b90506020850151603f1984830301606085015261523082826151bf565b925050508260208301529392505050565b60006020828403121561525357600080fd5b8135613bda81614984565b60005b83811015615279578181015183820152602001615261565b83811115612d945750506000910152565b600081518084526152a281602086016020860161525e565b601f01601f19169290920160200192915050565b600063ffffffff808f16835260018060a01b03808f166020850152808e166040850152808d1660608501528b6080850152808b1660a0850152508860c084015280881660e084015261018061010084015261531561018084018861528a565b610120840187905290851661014084015263ffffffff841661016084015290509d9c50505050505050505050505050565b8035613d88816146f1565b60008060006060848603121561536657600080fd5b8335615371816146f1565b92506020840135615381816146f1565b915060408401356149f9816146f1565b6000806000606084860312156153a657600080fd5b83356153b1816146f1565b92506020840135915060408401356149f981614984565b8281526040602082015260006153e16040830184614a04565b949350505050565b6000806000604084860312156153fe57600080fd5b83356001600160401b038082111561541557600080fd5b90850190610180828803121561542a57600080fd5b61543261479d565b61543b83614996565b815261544960208401615346565b602082015261545a60408401615346565b604082015261546b60608401615346565b60608201526080830135608082015261548660a08401615346565b60a082015260c083013560c08201526154a160e08401614996565b60e082015261010080840135838111156154ba57600080fd5b6154c68a828701614915565b8284015250506101208084013581830152506101406154e6818501614996565b908201526101606154f8848201614996565b908201529450602086013591508082111561551257600080fd5b5061551f86828701614e13565b9497909650939450505050565b60006020828403121561553e57600080fd5b8151613bda816146f1565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156155a557600080fd5b8151613bda81614abb565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261562b57634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561564357600080fd5b82516001600160401b0381111561565957600080fd5b8301601f8101851361566a57600080fd5b8051615678614b2a82614ae6565b81815260059190911b8201830190838101908783111561569757600080fd5b928401925b828410156156b55783518252928401929084019061569c565b979650505050505050565b6000602082840312156156d257600080fd5b81516001600160601b0381168114613bda57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415615713576157136156e9565b5060010190565b6000808335601e1984360301811261573157600080fd5b8301803591506001600160401b0382111561574b57600080fd5b6020019150368190038213156143a957600080fd5b6000808335601e1984360301811261577757600080fd5b83016020810192503590506001600160401b0381111561579657600080fd5b8036038313156143a957600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b602081526157ec602082016157e284614996565b63ffffffff169052565b60006157fa60208401615346565b6001600160a01b03811660408401525061581660408401615346565b6001600160a01b03811660608401525061583260608401615346565b6001600160a01b038116608084015250608083013560a083015261585860a08401615346565b6001600160a01b03811660c08401525060c083013560e083015261587e60e08401614996565b6101006158928185018363ffffffff169052565b61589e81860186615760565b9250905061018061012081818701526158bc6101a0870185856157a5565b935061014092508087013583870152506158d7828701614996565b91506101606158ed8187018463ffffffff169052565b6158f8818801614996565b925050613fe58186018363ffffffff169052565b803561591781614984565b63ffffffff168252602081013561592d81614abb565b8015156020840152505050565b60408101613bdd828461590c565b60006001600160601b038083168185168183048111821515161561596e5761596e6156e9565b02949350505050565b6000816000190483118215151615615991576159916156e9565b500290565b608081016159a4828561590c565b63ffffffff8351166040830152602083015160608301529392505050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156159ef57600080fd5b8260051b8085606085013760009201606001918252509392505050565b60006020808385031215615a1f57600080fd5b82516001600160401b03811115615a3557600080fd5b8301601f81018513615a4657600080fd5b8051615a54614b2a82614ae6565b81815260059190911b82018301908381019087831115615a7357600080fd5b928401925b828410156156b5578351615a8b81614984565b82529284019290840190615a78565b63ffffffff84168152604060208201526000615aba6040830184866157a5565b95945050505050565b600060208284031215615ad557600080fd5b81516001600160c01b0381168114613bda57600080fd5b600060208284031215615afe57600080fd5b8151613bda81614984565b600060ff821660ff811415615b2057615b206156e9565b60010192915050565b604081526000615b3d6040830185876157a5565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b81811015615b9a57845183529383019391830191600101615b7e565b5090979650505050505050565b600060208284031215615bb957600080fd5b8151613bda81615013565b600082821015615bd657615bd66156e9565b500390565b600060208284031215615bed57600080fd5b5051919050565b60008219821115615c0757615c076156e9565b500190565b600060208284031215615c1e57600080fd5b815167ffffffffffffffff1981168114613bda57600080fd5b60006001600160601b0383811690831681811015615c5757615c576156e9565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615c9a57815185529382019390820190600101615c7e565b5092979650505050505050565b600181811c90821680615cbb57607f821691505b60208210811415615cdc57634e487b7160e01b600052602260045260246000fd5b50919050565b805163ffffffff16825260006101806020830151615d0b60208601826001600160a01b03169052565b506040830151615d2660408601826001600160a01b03169052565b506060830151615d4160608601826001600160a01b03169052565b506080830151608085015260a0830151615d6660a08601826001600160a01b03169052565b5060c083015160c085015260e0830151615d8860e086018263ffffffff169052565b50610100808401518282870152615da18387018261528a565b9250505061012080840151818601525061014080840151615dc98287018263ffffffff169052565b50506101608381015163ffffffff811686830152613fe5565b602081526000610d0d6020830184615ce2565b606081526000615e086060830186615ce2565b63ffffffff851660208401528281036040840152615e26818561528a565b9695505050505050565b600061ffff80831681811415615e4857615e486156e9565b6001019392505050565b634e487b7160e01b600052602160045260246000fd5b8281526040602082015260006153e1604083018461528a565b60008251615e9381846020870161525e565b919091019291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122099c5e018ae0df5cb8c720f8a6f561c45932e9b3fb5760ab0249638fedceffb8864736f6c634300080c0033",
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
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers, uint256 expiry, uint32 targetNetworkNumber, uint32 createdBlock)
func (_ContractSettlement *ContractSettlementCaller) AllOrderDetails(opts *bind.CallOpts, arg0 uint32) (struct {
	OrderId                   uint32
	Maker                     common.Address
	Taker                     common.Address
	InputToken                common.Address
	InputAmount               *big.Int
	OutputToken               common.Address
	OutputAmount              *big.Int
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	Expiry                    *big.Int
	TargetNetworkNumber       uint32
	CreatedBlock              uint32
}, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "allOrderDetails", arg0)

	outstruct := new(struct {
		OrderId                   uint32
		Maker                     common.Address
		Taker                     common.Address
		InputToken                common.Address
		InputAmount               *big.Int
		OutputToken               common.Address
		OutputAmount              *big.Int
		QuorumThresholdPercentage uint32
		QuorumNumbers             []byte
		Expiry                    *big.Int
		TargetNetworkNumber       uint32
		CreatedBlock              uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Maker = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Taker = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.InputToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.InputAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OutputToken = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.OutputAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.QuorumThresholdPercentage = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.QuorumNumbers = *abi.ConvertType(out[8], new([]byte)).(*[]byte)
	outstruct.Expiry = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.TargetNetworkNumber = *abi.ConvertType(out[10], new(uint32)).(*uint32)
	outstruct.CreatedBlock = *abi.ConvertType(out[11], new(uint32)).(*uint32)

	return *outstruct, err

}

// AllOrderDetails is a free data retrieval call binding the contract method 0xaa4f5cc6.
//
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers, uint256 expiry, uint32 targetNetworkNumber, uint32 createdBlock)
func (_ContractSettlement *ContractSettlementSession) AllOrderDetails(arg0 uint32) (struct {
	OrderId                   uint32
	Maker                     common.Address
	Taker                     common.Address
	InputToken                common.Address
	InputAmount               *big.Int
	OutputToken               common.Address
	OutputAmount              *big.Int
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	Expiry                    *big.Int
	TargetNetworkNumber       uint32
	CreatedBlock              uint32
}, error) {
	return _ContractSettlement.Contract.AllOrderDetails(&_ContractSettlement.CallOpts, arg0)
}

// AllOrderDetails is a free data retrieval call binding the contract method 0xaa4f5cc6.
//
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint32 quorumThresholdPercentage, bytes quorumNumbers, uint256 expiry, uint32 targetNetworkNumber, uint32 createdBlock)
func (_ContractSettlement *ContractSettlementCallerSession) AllOrderDetails(arg0 uint32) (struct {
	OrderId                   uint32
	Maker                     common.Address
	Taker                     common.Address
	InputToken                common.Address
	InputAmount               *big.Int
	OutputToken               common.Address
	OutputAmount              *big.Int
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	Expiry                    *big.Int
	TargetNetworkNumber       uint32
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

// Fulfill is a paid mutator transaction binding the contract method 0xec381235.
//
// Solidity: function fulfill((uint32,address,address,address,uint256,address,uint256,uint32,bytes,uint256,uint32,uint32) order, bytes sig) returns()
func (_ContractSettlement *ContractSettlementTransactor) Fulfill(opts *bind.TransactOpts, order SettlementOrder, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "fulfill", order, sig)
}

// Fulfill is a paid mutator transaction binding the contract method 0xec381235.
//
// Solidity: function fulfill((uint32,address,address,address,uint256,address,uint256,uint32,bytes,uint256,uint32,uint32) order, bytes sig) returns()
func (_ContractSettlement *ContractSettlementSession) Fulfill(order SettlementOrder, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, order, sig)
}

// Fulfill is a paid mutator transaction binding the contract method 0xec381235.
//
// Solidity: function fulfill((uint32,address,address,address,uint256,address,uint256,uint32,bytes,uint256,uint32,uint32) order, bytes sig) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) Fulfill(order SettlementOrder, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, order, sig)
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

// RespondToFulfill is a paid mutator transaction binding the contract method 0x47f1385e.
//
// Solidity: function respondToFulfill((uint32,address,address,address,uint256,address,uint256,uint32,bytes,uint256,uint32,uint32) order, (uint32,bool) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSettlement *ContractSettlementTransactor) RespondToFulfill(opts *bind.TransactOpts, order SettlementOrder, orderResponse SettlementOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "respondToFulfill", order, orderResponse, nonSignerStakesAndSignature)
}

// RespondToFulfill is a paid mutator transaction binding the contract method 0x47f1385e.
//
// Solidity: function respondToFulfill((uint32,address,address,address,uint256,address,uint256,uint32,bytes,uint256,uint32,uint32) order, (uint32,bool) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSettlement *ContractSettlementSession) RespondToFulfill(order SettlementOrder, orderResponse SettlementOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSettlement.Contract.RespondToFulfill(&_ContractSettlement.TransactOpts, order, orderResponse, nonSignerStakesAndSignature)
}

// RespondToFulfill is a paid mutator transaction binding the contract method 0x47f1385e.
//
// Solidity: function respondToFulfill((uint32,address,address,address,uint256,address,uint256,uint32,bytes,uint256,uint32,uint32) order, (uint32,bool) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
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
	Order                     SettlementOrder
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterFulfillEvent is a free log retrieval operation binding the contract event 0x9a5d42e3ee6260573aeaa5309b575a6abe8be1d8490569b1400fba374bca74ef.
//
// Solidity: event fulfillEvent((uint32,address,address,address,uint256,address,uint256,uint32,bytes,uint256,uint32,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers)
func (_ContractSettlement *ContractSettlementFilterer) FilterFulfillEvent(opts *bind.FilterOpts) (*ContractSettlementFulfillEventIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "fulfillEvent")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementFulfillEventIterator{contract: _ContractSettlement.contract, event: "fulfillEvent", logs: logs, sub: sub}, nil
}

// WatchFulfillEvent is a free log subscription operation binding the contract event 0x9a5d42e3ee6260573aeaa5309b575a6abe8be1d8490569b1400fba374bca74ef.
//
// Solidity: event fulfillEvent((uint32,address,address,address,uint256,address,uint256,uint32,bytes,uint256,uint32,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers)
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

// ParseFulfillEvent is a log parse operation binding the contract event 0x9a5d42e3ee6260573aeaa5309b575a6abe8be1d8490569b1400fba374bca74ef.
//
// Solidity: event fulfillEvent((uint32,address,address,address,uint256,address,uint256,uint32,bytes,uint256,uint32,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers)
func (_ContractSettlement *ContractSettlementFilterer) ParseFulfillEvent(log types.Log) (*ContractSettlementFulfillEvent, error) {
	event := new(ContractSettlementFulfillEvent)
	if err := _ContractSettlement.contract.UnpackLog(event, "fulfillEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

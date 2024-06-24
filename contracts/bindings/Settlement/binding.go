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

// ISettlementOrder is an auto generated low-level Go binding around an user-defined struct.
type ISettlementOrder struct {
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

// ISettlementOrderResponse is an auto generated low-level Go binding around an user-defined struct.
type ISettlementOrderResponse struct {
	Recipient           common.Address
	ReferenceOrderIndex uint32
}

// ISettlementOrderResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type ISettlementOrderResponseMetadata struct {
	OrderResponsedBlock uint32
	HashOfNonSigners    [32]byte
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

// ContractSettlementMetaData contains all meta data concerning the ContractSettlement contract.
var ContractSettlementMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderDetails\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderExecutions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfill\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structISettlement.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structISettlement.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToFulfill\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderResponse\",\"type\":\"tuple\",\"internalType\":\"structISettlement.OrderResponse\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structISettlement.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FulfillEvent\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISettlement.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderRespondedEvent\",\"inputs\":[{\"name\":\"orderResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISettlement.OrderResponse\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"OrderResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISettlement.OrderResponseMetadata\",\"components\":[{\"name\":\"orderResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderFulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TakerMismatch\",\"inputs\":[]}]",
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162006367380380620063678339810160408190526200003591620002d4565b6040518060400160405280600a81526020016914d95d1d1b195b595b9d60b21b815250604051806040016040528060038152602001620312e360ec1b81525082806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000ce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000f49190620002d4565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200014c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001729190620002d4565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001cc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001f29190620002d4565b6001600160a01b031660e052506097805460ff19166001179055815160208084019190912082519183019190912061016082905261018081905246610120527f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f620002a28184846040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b6101005230610140526101a05250620002fb9350505050565b6001600160a01b0381168114620002d157600080fd5b50565b600060208284031215620002e757600080fd5b8151620002f481620002bb565b9392505050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a051615fb8620003af600039600061441901526000614468015260006144430152600061439c015260006143c6015260006143f00152600081816105a1015261295b01526000818161039e0152612b3d0152600081816103d801528181612d130152612ed50152600081816103ff015281816113a301528181612626015281816127be01526129f80152615fb86000f3fe608060405234801561001057600080fd5b50600436106101ef5760003560e01c8063658d63411161010f578063b98d0908116100a2578063df5cf72311610071578063df5cf7231461059c578063e772935f146105c3578063f2fde38b146105e3578063fabc1cbc146105f657600080fd5b8063b98d09081461053b578063c0c53b8b14610548578063cefdc1d41461055b578063dd3285b31461057c57600080fd5b8063715018a6116100de578063715018a614610442578063886f11951461044a5780638da5cb5b1461045d578063aa4f5cc61461046e57600080fd5b8063658d6341146103c057806368304835146103d35780636d14a987146103fa5780636efb46361461042157600080fd5b80633b692c09116101875780635ac86ab7116101565780635ac86ab71461034e5780635c155662146103715780635c975abb146103915780635df459461461039957600080fd5b80633b692c09146102f2578063416c7e5e146103135780634f739f7414610326578063595c6a671461034657600080fd5b8063245a7bfc116101c3578063245a7bfc14610271578063298fb5691461029c5780633563b0d1146102bf578063377a91a6146102df57600080fd5b80629e8b57146101f457806310d67a2f1461021f578063136439dd14610234578063171f1d5b14610247575b600080fd5b6102076102023660046149a1565b610609565b604051610216939291906149be565b60405180910390f35b61023261022d366004614a40565b6106c1565b005b610232610242366004614a5d565b61077d565b61025a610255366004614bdb565b6108bc565b604080519215158352901515602083015201610216565b60cd54610284906001600160a01b031681565b6040516001600160a01b039091168152602001610216565b6102af6102aa366004614c86565b610a46565b6040519015158152602001610216565b6102d26102cd366004614cdb565b610abc565b6040516102169190614e36565b6102326102ed3660046150ed565b610f52565b61030561030036600461519b565b611287565b604051908152602001610216565b6102326103213660046151c6565b6113a1565b6103396103343660046151e3565b611516565b60405161021691906152e7565b610232611c3c565b6102af61035c3660046153b1565b606654600160ff9092169190911b9081161490565b61038461037f3660046153ce565b611d03565b604051610216919061547a565b606654610305565b6102847f000000000000000000000000000000000000000000000000000000000000000081565b6102326103ce3660046154be565b611ecb565b6102847f000000000000000000000000000000000000000000000000000000000000000081565b6102847f000000000000000000000000000000000000000000000000000000000000000081565b61043461042f366004615555565b612273565b604051610216929190615615565b61023261318a565b606554610284906001600160a01b031681565b6033546001600160a01b0316610284565b6104de61047c3660046149a1565b60ca602052600090815260409020805460018201546002830154600384015460048501546005860154600687015460079097015463ffffffff808816986401000000009098046001600160a01b03908116989781169796811696941693911689565b6040805163ffffffff9a8b1681526001600160a01b03998a1660208201529789169088015294871660608701526080860193909352941660a084015260c083019390935260e0820192909252911661010082015261012001610216565b6097546102af9060ff1681565b610232610556366004615669565b61319e565b61056e6105693660046156b4565b6132d9565b6040516102169291906156eb565b61030561058a3660046149a1565b60c96020526000908152604090205481565b6102847f000000000000000000000000000000000000000000000000000000000000000081565b6103056105d13660046149a1565b60cc6020526000908152604090205481565b6102326105f1366004614a40565b61346b565b610232610604366004614a5d565b6134e1565b60cb602052600090815260409020805460018201805463ffffffff90921692916106329061570c565b80601f016020809104026020016040519081016040528092919081815260200182805461065e9061570c565b80156106ab5780601f10610680576101008083540402835291602001916106ab565b820191906000526020600020905b81548152906001019060200180831161068e57829003601f168201915b5050506002909301549192505063ffffffff1683565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610714573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107389190615741565b6001600160a01b0316336001600160a01b0316146107715760405162461bcd60e51b81526004016107689061575e565b60405180910390fd5b61077a8161363d565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156107c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e991906157a8565b6108055760405162461bcd60e51b8152600401610768906157c5565b6066548181161461087e5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610768565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106109045761090461580d565b60200201518951600160200201518a602001516000600281106109295761092961580d565b60200201518b602001516001600281106109455761094561580d565b602090810291909101518c518d8301516040516109a29a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6109c59190615823565b9050610a386109de6109d78884613734565b86906137cb565b6109e661385f565b610a2e610a1f85610a19604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613734565b610a288c61391f565b906137cb565b886201d4c06139af565b909890975095505050505050565b600080610a9484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610a8e92508991506112879050565b90613bd3565b9050610aa66040860160208701614a40565b6001600160a01b03918216911614949350505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610afe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b229190615741565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b889190615741565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bee9190615741565b9050600086516001600160401b03811115610c0b57610c0b614a76565b604051908082528060200260200182016040528015610c3e57816020015b6060815260200190600190039081610c295790505b50905060005b8751811015610f46576000888281518110610c6157610c6161580d565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610cc2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610cea9190810190615845565b905080516001600160401b03811115610d0557610d05614a76565b604051908082528060200260200182016040528015610d5057816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610d235790505b50848481518110610d6357610d6361580d565b602002602001018190525060005b8151811015610f30576040518060600160405280876001600160a01b03166347b314e8858581518110610da657610da661580d565b60200260200101516040518263ffffffff1660e01b8152600401610dcc91815260200190565b602060405180830381865afa158015610de9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e0d9190615741565b6001600160a01b03168152602001838381518110610e2d57610e2d61580d565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610e5b57610e5b61580d565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610eb7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610edb91906158d5565b6001600160601b0316815250858581518110610ef957610ef961580d565b60200260200101518281518110610f1257610f1261580d565b60200260200101819052508080610f2890615914565b915050610d71565b5050508080610f3e90615914565b915050610c44565b50979650505050505050565b60cd546001600160a01b03163314610fac5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610768565b600060cc81610fc160408601602087016149a1565b63ffffffff1663ffffffff168152602001908152602001600020541461103f5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526c103a37903a34329037b93232b960991b6064820152608401610768565b6000826040516020016110529190615964565b60405160208183030381529060405280519060200120905060008061107a838a8a8988612273565b91509150600060405180604001604052804363ffffffff16815260200183815250905085816040516020016110b0929190615972565b6040516020818303038152906040528051906020012060cc60008860200160208101906110dd91906149a1565b63ffffffff1663ffffffff16815260200190815260200160002081905550600060ca600088602001602081019061111491906149a1565b63ffffffff90811682526020808301939093526040918201600020825161012081018452815480841682526001600160a01b036401000000009091048116828701526001830154811694820194909452600282015484166060820181905260038301546080830152600483015490941660a0820152600582015460c0820152600682015460e08201526007909101549091166101008201529250906323b872dd9030906111c3908b018b614a40565b60808501516040516001600160e01b031960e086901b1681526001600160a01b03938416600482015292909116602483015260448201526064016020604051808303816000875af115801561121c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124091906157a8565b507fa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd8783604051611272929190615972565b60405180910390a15050505050505050505050565b600061139b7fcd190519e6feb4299db38636bfbcef24f017d944ee1188f74093454f21ed61f16112ba60208501856149a1565b6112ca6040860160208701614a40565b6112da6060870160408801614a40565b6112ea6080880160608901614a40565b60808801356112ff60c08a0160a08b01614a40565b60c08a013560e08b013561131b6101208d016101008e016149a1565b60408051602081019b909b5263ffffffff998a16908b01526001600160a01b0397881660608b015295871660808a015293861660a089015260c088019290925290931660e0860152610100850192909252610120840191909152166101408201526101600160405160208183030381529060405280519060200120613bf7565b92915050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114239190615741565b6001600160a01b0316336001600160a01b0316146114cf5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610768565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b6115416040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611581573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115a59190615741565b90506115d26040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611602908b908990899060040161599e565b600060405180830381865afa15801561161f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261164791908101906159e8565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611679908b908b908b90600401615a9f565b600060405180830381865afa158015611696573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116be91908101906159e8565b6040820152856001600160401b038111156116db576116db614a76565b60405190808252806020026020018201604052801561170e57816020015b60608152602001906001900390816116f95790505b50606082015260005b60ff8116871115611b4d576000856001600160401b0381111561173c5761173c614a76565b604051908082528060200260200182016040528015611765578160200160208202803683370190505b5083606001518360ff168151811061177f5761177f61580d565b602002602001018190525060005b86811015611a4d5760008c6001600160a01b03166304ec63518a8a858181106117b8576117b861580d565b905060200201358e886000015186815181106117d6576117d661580d565b60200260200101516040518463ffffffff1660e01b81526004016118139392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611830573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118549190615ac8565b90506001600160c01b0381166118f85760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610768565b8a8a8560ff1681811061190d5761190d61580d565b6001600160c01b03841692013560f81c9190911c600190811614159050611a3a57856001600160a01b031663dd9846b98a8a8581811061194f5761194f61580d565b905060200201358d8d8860ff1681811061196b5761196b61580d565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156119c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119e59190615af1565b85606001518560ff16815181106119fe576119fe61580d565b60200260200101518481518110611a1757611a1761580d565b63ffffffff9092166020928302919091019091015282611a3681615914565b9350505b5080611a4581615914565b91505061178d565b506000816001600160401b03811115611a6857611a68614a76565b604051908082528060200260200182016040528015611a91578160200160208202803683370190505b50905060005b82811015611b125784606001518460ff1681518110611ab857611ab861580d565b60200260200101518181518110611ad157611ad161580d565b6020026020010151828281518110611aeb57611aeb61580d565b63ffffffff9092166020928302919091019091015280611b0a81615914565b915050611a97565b508084606001518460ff1681518110611b2d57611b2d61580d565b602002602001018190525050508080611b4590615b0e565b915050611717565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bb29190615741565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611be5908b908b908e90600401615b2e565b600060405180830381865afa158015611c02573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c2a91908101906159e8565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611c84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ca891906157a8565b611cc45760405162461bcd60e51b8152600401610768906157c5565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611d35929190615b58565b600060405180830381865afa158015611d52573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611d7a91908101906159e8565b9050600084516001600160401b03811115611d9757611d97614a76565b604051908082528060200260200182016040528015611dc0578160200160208202803683370190505b50905060005b8551811015611ec157866001600160a01b03166304ec6351878381518110611df057611df061580d565b602002602001015187868581518110611e0b57611e0b61580d565b60200260200101516040518463ffffffff1660e01b8152600401611e489392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611e65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e899190615ac8565b6001600160c01b0316828281518110611ea457611ea461580d565b602090810291909101015280611eb981615914565b915050611dc6565b5095945050505050565b6000611eda60208801886149a1565b63ffffffff8116600090815260c9602052604090205490915015611f11576040516356f1733f60e01b815260040160405180910390fd5b6000611f236060890160408a01614a40565b6001600160a01b031614158015611f5b5750611f456060880160408901614a40565b6001600160a01b0316336001600160a01b031614155b15611f7957604051630d9d754f60e41b815260040160405180910390fd5b428760e001351015611f9e576040516362b439dd60e11b815260040160405180910390fd5b611fa9878484610a46565b611fc657604051638baa579f60e01b815260040160405180910390fd5b611fd66080880160608901614a40565b6001600160a01b03166323b872dd611ff460408a0160208b01614a40565b6040516001600160e01b031960e084901b1681526001600160a01b03909116600482015230602482015260808a013560448201526064016020604051808303816000875af115801561204a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061206e91906157a8565b5061207f60c0880160a08901614a40565b6001600160a01b03166323b872dd3361209e60408b0160208c01614a40565b6040516001600160e01b031960e085901b1681526001600160a01b0392831660048201529116602482015260c08a013560448201526064016020604051808303816000875af11580156120f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061211991906157a8565b5060405180606001604052808763ffffffff16815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050504363ffffffff908116602093840152848116825260cb835260409091208351815463ffffffff1916921691909117815582820151805191926121b392600185019290910190614811565b50604091820151600291909101805463ffffffff191663ffffffff909216919091179055516121e6908890602001615c67565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260c984528281209190915560ca909252902087906122288282615cb0565b9050507f94a6dd0de15ac83f39634005b9c02f69524f7797c69b9f7c1f8a95a5d527391c8787878733604051612262959493929190615d9f565b60405180910390a150505050505050565b60408051808201909152606080825260208201526000846122ea5760405162461bcd60e51b81526020600482015260376024820152600080516020615f6383398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610768565b60408301515185148015612302575060a08301515185145b8015612312575060c08301515185145b8015612322575060e08301515185145b61238c5760405162461bcd60e51b81526020600482015260416024820152600080516020615f6383398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610768565b825151602084015151146124045760405162461bcd60e51b815260206004820152604460248201819052600080516020615f63833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610768565b4363ffffffff168463ffffffff16106124735760405162461bcd60e51b815260206004820152603c6024820152600080516020615f6383398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610768565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b038111156124b4576124b4614a76565b6040519080825280602002602001820160405280156124dd578160200160208202803683370190505b506020820152866001600160401b038111156124fb576124fb614a76565b604051908082528060200260200182016040528015612524578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561255857612558614a76565b604051908082528060200260200182016040528015612581578160200160208202803683370190505b5081526020860151516001600160401b038111156125a1576125a1614a76565b6040519080825280602002602001820160405280156125ca578160200160208202803683370190505b508160200181905250600061269c8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612673573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126979190615deb565b613c45565b905060005b876020015151811015612937576126e6886020015182815181106126c7576126c761580d565b6020026020010151805160009081526020918201519091526040902090565b836020015182815181106126fc576126fc61580d565b602090810291909101015280156127bc57602083015161271d600183615e08565b8151811061272d5761272d61580d565b602002602001015160001c8360200151828151811061274e5761274e61580d565b602002602001015160001c116127bc576040805162461bcd60e51b8152602060048201526024810191909152600080516020615f6383398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610768565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106128015761280161580d565b60200260200101518b8b6000015185815181106128205761282061580d565b60200260200101516040518463ffffffff1660e01b815260040161285d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561287a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061289e9190615ac8565b6001600160c01b0316836000015182815181106128bd576128bd61580d565b6020026020010181815250506129236109d76128f784866000015185815181106128e9576128e961580d565b602002602001015116613cd6565b8a60200151848151811061290d5761290d61580d565b6020026020010151613d0190919063ffffffff16565b94508061292f81615914565b9150506126a1565b505061294283613de5565b60975490935060ff166000816129595760006129db565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129db9190615e1f565b905060005b8a811015613059578215612b3b578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612a3757612a3761580d565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612a77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a9b9190615e1f565b612aa59190615e38565b11612b3b5760405162461bcd60e51b81526020600482015260666024820152600080516020615f6383398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610768565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612b7c57612b7c61580d565b9050013560f81c60f81b60f81c8c8c60a001518581518110612ba057612ba061580d565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612bfc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c209190615e50565b6001600160401b031916612c438a6040015183815181106126c7576126c761580d565b67ffffffffffffffff191614612cdf5760405162461bcd60e51b81526020600482015260616024820152600080516020615f6383398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610768565b612d0f89604001518281518110612cf857612cf861580d565b6020026020010151876137cb90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612d5257612d5261580d565b9050013560f81c60f81b60f81c8c8c60c001518581518110612d7657612d7661580d565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612dd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612df691906158d5565b85602001518281518110612e0c57612e0c61580d565b6001600160601b03909216602092830291909101820152850151805182908110612e3857612e3861580d565b602002602001015185600001518281518110612e5657612e5661580d565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561304457612ece86600001518281518110612ea057612ea061580d565b60200260200101518f8f86818110612eba57612eba61580d565b600192013560f81c9290921c811614919050565b15613032577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612f1457612f1461580d565b9050013560f81c60f81b60f81c8e89602001518581518110612f3857612f3861580d565b60200260200101518f60e001518881518110612f5657612f5661580d565b60200260200101518781518110612f6f57612f6f61580d565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612fd3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ff791906158d5565b875180518590811061300b5761300b61580d565b6020026020010181815161301f9190615e7b565b6001600160601b03169052506001909101905b8061303c81615914565b915050612e7a565b5050808061305190615914565b9150506129e0565b5050506000806130738c868a606001518b608001516108bc565b91509150816130e45760405162461bcd60e51b81526020600482015260436024820152600080516020615f6383398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610768565b806131455760405162461bcd60e51b81526020600482015260396024820152600080516020615f6383398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610768565b50506000878260200151604051602001613160929190615ea3565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b613192613e80565b61319c6000613eda565b565b600054610100900460ff16158080156131be5750600054600160ff909116105b806131d85750303b1580156131d8575060005460ff166001145b61323b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610768565b6000805460ff19166001179055801561325e576000805461ff0019166101001790555b613269846000613f2c565b61327283613eda565b60cd80546001600160a01b0319166001600160a01b03841617905580156132d3576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106133145761331461580d565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906133509088908690600401615b58565b600060405180830381865afa15801561336d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261339591908101906159e8565b6000815181106133a7576133a761580d565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015613413573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134379190615ac8565b6001600160c01b03169050600061344d82614016565b90508161345b8a838a610abc565b9550955050505050935093915050565b613473613e80565b6001600160a01b0381166134d85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610768565b61077a81613eda565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613534573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135589190615741565b6001600160a01b0316336001600160a01b0316146135885760405162461bcd60e51b81526004016107689061575e565b6066541981196066541916146136065760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610768565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016108b1565b6001600160a01b0381166136cb5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610768565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613750614895565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561378357613785565bfe5b50806137c35760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610768565b505092915050565b60408051808201909152600080825260208201526137e76148b3565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156137835750806137c35760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610768565b6138676148d1565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061394f600080516020615f4383398151915286615823565b90505b61395b816140e2565b9093509150600080516020615f43833981519152828309831415613995576040805180820190915290815260208101919091529392505050565b600080516020615f43833981519152600182089050613952565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906139e16148f6565b60005b6002811015613ba65760006139fa826006615eeb565b9050848260028110613a0e57613a0e61580d565b60200201515183613a20836000615e38565b600c8110613a3057613a3061580d565b6020020152848260028110613a4757613a4761580d565b60200201516020015183826001613a5e9190615e38565b600c8110613a6e57613a6e61580d565b6020020152838260028110613a8557613a8561580d565b6020020151515183613a98836002615e38565b600c8110613aa857613aa861580d565b6020020152838260028110613abf57613abf61580d565b6020020151516001602002015183613ad8836003615e38565b600c8110613ae857613ae861580d565b6020020152838260028110613aff57613aff61580d565b602002015160200151600060028110613b1a57613b1a61580d565b602002015183613b2b836004615e38565b600c8110613b3b57613b3b61580d565b6020020152838260028110613b5257613b5261580d565b602002015160200151600160028110613b6d57613b6d61580d565b602002015183613b7e836005615e38565b600c8110613b8e57613b8e61580d565b60200201525080613b9e81615914565b9150506139e4565b50613baf614915565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6000806000613be28585614164565b91509150613bef816141d4565b509392505050565b600061139b613c0461438f565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b600080613c51846144b6565b9050808360ff166001901b11613ccf5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610768565b9392505050565b6000805b821561139b57613ceb600184615e08565b9092169180613cf981615f0a565b915050613cda565b60408051808201909152600080825260208201526102008261ffff1610613d5d5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610768565b8161ffff1660011415613d7157508161139b565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613dda57600161ffff871660ff83161c81161415613dbd57613dba84846137cb565b93505b613dc783846137cb565b92506201fffe600192831b169101613d8d565b509195945050505050565b60408051808201909152600080825260208201528151158015613e0a57506020820151155b15613e28575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615f438339815191528460200151613e5b9190615823565b613e7390600080516020615f43833981519152615e08565b905292915050565b919050565b6033546001600160a01b0316331461319c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610768565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546001600160a01b0316158015613f4d57506001600160a01b03821615155b613fcf5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610768565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26140128261363d565b5050565b606060008061402484613cd6565b61ffff166001600160401b0381111561403f5761403f614a76565b6040519080825280601f01601f191660200182016040528015614069576020820181803683370190505b5090506000805b825182108015614081575061010081105b156140d8576001811b9350858416156140c8578060f81b8383815181106140aa576140aa61580d565b60200101906001600160f81b031916908160001a9053508160010191505b6140d181615914565b9050614070565b5090949350505050565b60008080600080516020615f438339815191526003600080516020615f4383398151915286600080516020615f43833981519152888909090890506000614158827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615f43833981519152614643565b91959194509092505050565b60008082516041141561419b5760208301516040840151606085015160001a61418f878285856146eb565b945094505050506141cd565b8251604014156141c557602083015160408401516141ba8683836147d8565b9350935050506141cd565b506000905060025b9250929050565b60008160048111156141e8576141e8615f2c565b14156141f15750565b600181600481111561420557614205615f2c565b14156142535760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610768565b600281600481111561426757614267615f2c565b14156142b55760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610768565b60038160048111156142c9576142c9615f2c565b14156143225760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610768565b600481600481111561433657614336615f2c565b141561077a5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610768565b6000306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156143e857507f000000000000000000000000000000000000000000000000000000000000000046145b1561441257507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b60006101008251111561453f5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610768565b815161454d57506000919050565b600080836000815181106145635761456361580d565b0160200151600160f89190911c81901b92505b845181101561463a578481815181106145915761459161580d565b0160200151600160f89190911c1b91508282116146265760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610768565b9181179161463381615914565b9050614576565b50909392505050565b60008061464e614915565b614656614933565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156137835750826146e05760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610768565b505195945050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561472257506000905060036147cf565b8460ff16601b1415801561473a57508460ff16601c14155b1561474b57506000905060046147cf565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561479f573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166147c8576000600192509250506147cf565b9150600090505b94509492505050565b6000806001600160ff1b038316816147f560ff86901c601b615e38565b9050614803878288856146eb565b935093505050935093915050565b82805461481d9061570c565b90600052602060002090601f01602090048101928261483f5760008555614885565b82601f1061485857805160ff1916838001178555614885565b82800160010185558215614885579182015b8281111561488557825182559160200191906001019061486a565b50614891929150614951565b5090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806148e4614966565b81526020016148f1614966565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b5b808211156148915760008155600101614952565b60405180604001604052806002906020820280368337509192915050565b63ffffffff8116811461077a57600080fd5b8035613e7b81614984565b6000602082840312156149b357600080fd5b8135613ccf81614984565b600063ffffffff80861683526020606081850152855180606086015260005b818110156149f9578781018301518682016080015282016149dd565b81811115614a0b576000608083880101525b509490911660408401525050601f91909101601f19160160800192915050565b6001600160a01b038116811461077a57600080fd5b600060208284031215614a5257600080fd5b8135613ccf81614a2b565b600060208284031215614a6f57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614aae57614aae614a76565b60405290565b60405161010081016001600160401b0381118282101715614aae57614aae614a76565b604051601f8201601f191681016001600160401b0381118282101715614aff57614aff614a76565b604052919050565b600060408284031215614b1957600080fd5b614b21614a8c565b9050813581526020820135602082015292915050565b600082601f830112614b4857600080fd5b604051604081018181106001600160401b0382111715614b6a57614b6a614a76565b8060405250806040840185811115614b8157600080fd5b845b81811015613dda578035835260209283019201614b83565b600060808284031215614bad57600080fd5b614bb5614a8c565b9050614bc18383614b37565b8152614bd08360408401614b37565b602082015292915050565b6000806000806101208587031215614bf257600080fd5b84359350614c038660208701614b07565b9250614c128660608701614b9b565b9150614c218660e08701614b07565b905092959194509250565b60006101208284031215614c3f57600080fd5b50919050565b60008083601f840112614c5757600080fd5b5081356001600160401b03811115614c6e57600080fd5b6020830191508360208285010111156141cd57600080fd5b60008060006101408486031215614c9c57600080fd5b614ca68585614c2c565b92506101208401356001600160401b03811115614cc257600080fd5b614cce86828701614c45565b9497909650939450505050565b600080600060608486031215614cf057600080fd5b8335614cfb81614a2b565b92506020848101356001600160401b0380821115614d1857600080fd5b818701915087601f830112614d2c57600080fd5b813581811115614d3e57614d3e614a76565b614d50601f8201601f19168501614ad7565b91508082528884828501011115614d6657600080fd5b8084840185840137600084828401015250809450505050614d8960408501614996565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614e28578385038a52825180518087529087019087870190845b81811015614e1357835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614dcf565b50509a87019a95505091850191600101614db1565b509298975050505050505050565b602081526000613ccf6020830184614d92565b60006001600160401b03821115614e6257614e62614a76565b5060051b60200190565b600082601f830112614e7d57600080fd5b81356020614e92614e8d83614e49565b614ad7565b82815260059290921b84018101918181019086841115614eb157600080fd5b8286015b84811015614ed5578035614ec881614984565b8352918301918301614eb5565b509695505050505050565b600082601f830112614ef157600080fd5b81356020614f01614e8d83614e49565b82815260069290921b84018101918181019086841115614f2057600080fd5b8286015b84811015614ed557614f368882614b07565b835291830191604001614f24565b600082601f830112614f5557600080fd5b81356020614f65614e8d83614e49565b82815260059290921b84018101918181019086841115614f8457600080fd5b8286015b84811015614ed55780356001600160401b03811115614fa75760008081fd5b614fb58986838b0101614e6c565b845250918301918301614f88565b60006101808284031215614fd657600080fd5b614fde614ab4565b905081356001600160401b0380821115614ff757600080fd5b61500385838601614e6c565b8352602084013591508082111561501957600080fd5b61502585838601614ee0565b6020840152604084013591508082111561503e57600080fd5b61504a85838601614ee0565b604084015261505c8560608601614b9b565b606084015261506e8560e08601614b07565b608084015261012084013591508082111561508857600080fd5b61509485838601614e6c565b60a08401526101408401359150808211156150ae57600080fd5b6150ba85838601614e6c565b60c08401526101608401359150808211156150d457600080fd5b506150e184828501614f44565b60e08301525092915050565b60008060008060008086880360c081121561510757600080fd5b87356001600160401b038082111561511e57600080fd5b61512a8b838c01614c45565b909950975060208a0135915061513f82614984565b90955060408901359061515182614984565b8195506040605f198401121561516657600080fd5b60608a01945060a08a013592508083111561518057600080fd5b505061518e89828a01614fc3565b9150509295509295509295565b600061012082840312156151ae57600080fd5b613ccf8383614c2c565b801515811461077a57600080fd5b6000602082840312156151d857600080fd5b8135613ccf816151b8565b600080600080600080608087890312156151fc57600080fd5b863561520781614a2b565b9550602087013561521781614984565b945060408701356001600160401b038082111561523357600080fd5b61523f8a838b01614c45565b9096509450606089013591508082111561525857600080fd5b818901915089601f83011261526c57600080fd5b81358181111561527b57600080fd5b8a60208260051b850101111561529057600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b838110156152dc57815163ffffffff16875295820195908201906001016152ba565b509495945050505050565b60006020808352835160808285015261530360a08501826152a6565b905081850151601f198086840301604087015261532083836152a6565b9250604087015191508086840301606087015261533d83836152a6565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561539457848783030184526153828287516152a6565b95880195938801939150600101615368565b509998505050505050505050565b60ff8116811461077a57600080fd5b6000602082840312156153c357600080fd5b8135613ccf816153a2565b6000806000606084860312156153e357600080fd5b83356153ee81614a2b565b92506020848101356001600160401b0381111561540a57600080fd5b8501601f8101871361541b57600080fd5b8035615429614e8d82614e49565b81815260059190911b8201830190838101908983111561544857600080fd5b928401925b828410156154665783358252928401929084019061544d565b8096505050505050614d8960408501614996565b6020808252825182820181905260009190848201906040850190845b818110156154b257835183529284019291840191600101615496565b50909695505050505050565b60008060008060008061018087890312156154d857600080fd5b6154e28888614c2c565b95506101208701356154f381614984565b94506101408701356001600160401b038082111561551057600080fd5b61551c8a838b01614c45565b909650945061016089013591508082111561553657600080fd5b5061554389828a01614c45565b979a9699509497509295939492505050565b60008060008060006080868803121561556d57600080fd5b8535945060208601356001600160401b038082111561558b57600080fd5b61559789838a01614c45565b9096509450604088013591506155ac82614984565b909250606087013590808211156155c257600080fd5b506155cf88828901614fc3565b9150509295509295909350565b600081518084526020808501945080840160005b838110156152dc5781516001600160601b0316875295820195908201906001016155f0565b604081526000835160408084015261563060808401826155dc565b90506020850151603f1984830301606085015261564d82826155dc565b925050508260208301529392505050565b8035613e7b81614a2b565b60008060006060848603121561567e57600080fd5b833561568981614a2b565b9250602084013561569981614a2b565b915060408401356156a981614a2b565b809150509250925092565b6000806000606084860312156156c957600080fd5b83356156d481614a2b565b92506020840135915060408401356156a981614984565b8281526040602082015260006157046040830184614d92565b949350505050565b600181811c9082168061572057607f821691505b60208210811415614c3f57634e487b7160e01b600052602260045260246000fd5b60006020828403121561575357600080fd5b8151613ccf81614a2b565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156157ba57600080fd5b8151613ccf816151b8565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261584057634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561585857600080fd5b82516001600160401b0381111561586e57600080fd5b8301601f8101851361587f57600080fd5b805161588d614e8d82614e49565b81815260059190911b820183019083810190878311156158ac57600080fd5b928401925b828410156158ca578351825292840192908401906158b1565b979650505050505050565b6000602082840312156158e757600080fd5b81516001600160601b0381168114613ccf57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415615928576159286158fe565b5060010190565b803561593a81614a2b565b6001600160a01b03168252602081013561595381614984565b63ffffffff81166020840152505050565b6040810161139b828461592f565b60808101615980828561592f565b63ffffffff8351166040830152602083015160608301529392505050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156159cb57600080fd5b8260051b8085606085013760009201606001918252509392505050565b600060208083850312156159fb57600080fd5b82516001600160401b03811115615a1157600080fd5b8301601f81018513615a2257600080fd5b8051615a30614e8d82614e49565b81815260059190911b82018301908381019087831115615a4f57600080fd5b928401925b828410156158ca578351615a6781614984565b82529284019290840190615a54565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615abf604083018486615a76565b95945050505050565b600060208284031215615ada57600080fd5b81516001600160c01b0381168114613ccf57600080fd5b600060208284031215615b0357600080fd5b8151613ccf81614984565b600060ff821660ff811415615b2557615b256158fe565b60010192915050565b604081526000615b42604083018587615a76565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b81811015615b9f57845183529383019391830191600101615b83565b5090979650505050505050565b8035615bb781614984565b63ffffffff1682526020810135615bcd81614a2b565b6001600160a01b03166020830152615be76040820161565e565b6001600160a01b03166040830152615c016060820161565e565b6001600160a01b0316606083015260808181013590830152615c2560a0820161565e565b6001600160a01b03811660a08401525060c081013560c083015260e081013560e0830152610100615c57818301614996565b63ffffffff8116848301526132d3565b610120810161139b8284615bac565b6000813561139b81614984565b6000813561139b81614a2b565b80546001600160a01b0319166001600160a01b0392909216919091179055565b615cd3615cbc83615c76565b825463ffffffff191663ffffffff91909116178255565b615d0e615ce260208401615c83565b828054640100000000600160c01b03191660209290921b640100000000600160c01b0316919091179055565b615d26615d1d60408401615c83565b60018301615c90565b615d3e615d3560608401615c83565b60028301615c90565b60808201356003820155615d60615d5760a08401615c83565b60048301615c90565b60c0820135600582015560e08201356006820155614012615d846101008401615c76565b6007830163ffffffff821663ffffffff198254161781555050565b6000610180615dae8389615bac565b63ffffffff871661012084015280610140840152615dcf8184018688615a76565b91505060018060a01b0383166101608301529695505050505050565b600060208284031215615dfd57600080fd5b8151613ccf816153a2565b600082821015615e1a57615e1a6158fe565b500390565b600060208284031215615e3157600080fd5b5051919050565b60008219821115615e4b57615e4b6158fe565b500190565b600060208284031215615e6257600080fd5b815167ffffffffffffffff1981168114613ccf57600080fd5b60006001600160601b0383811690831681811015615e9b57615e9b6158fe565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615ede57815185529382019390820190600101615ec2565b5092979650505050505050565b6000816000190483118215151615615f0557615f056158fe565b500290565b600061ffff80831681811415615f2257615f226158fe565b6001019392505050565b634e487b7160e01b600052602160045260246000fdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212203628425a44b7e8e38bcd882a0c5c766d0f25232f800010b5d72178b0144ba24264736f6c634300080c0033",
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
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint256 expiry, uint32 targetNetworkNumber)
func (_ContractSettlement *ContractSettlementCaller) AllOrderDetails(opts *bind.CallOpts, arg0 uint32) (struct {
	OrderId             uint32
	Maker               common.Address
	Taker               common.Address
	InputToken          common.Address
	InputAmount         *big.Int
	OutputToken         common.Address
	OutputAmount        *big.Int
	Expiry              *big.Int
	TargetNetworkNumber uint32
}, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "allOrderDetails", arg0)

	outstruct := new(struct {
		OrderId             uint32
		Maker               common.Address
		Taker               common.Address
		InputToken          common.Address
		InputAmount         *big.Int
		OutputToken         common.Address
		OutputAmount        *big.Int
		Expiry              *big.Int
		TargetNetworkNumber uint32
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
	outstruct.Expiry = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TargetNetworkNumber = *abi.ConvertType(out[8], new(uint32)).(*uint32)

	return *outstruct, err

}

// AllOrderDetails is a free data retrieval call binding the contract method 0xaa4f5cc6.
//
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint256 expiry, uint32 targetNetworkNumber)
func (_ContractSettlement *ContractSettlementSession) AllOrderDetails(arg0 uint32) (struct {
	OrderId             uint32
	Maker               common.Address
	Taker               common.Address
	InputToken          common.Address
	InputAmount         *big.Int
	OutputToken         common.Address
	OutputAmount        *big.Int
	Expiry              *big.Int
	TargetNetworkNumber uint32
}, error) {
	return _ContractSettlement.Contract.AllOrderDetails(&_ContractSettlement.CallOpts, arg0)
}

// AllOrderDetails is a free data retrieval call binding the contract method 0xaa4f5cc6.
//
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint256 expiry, uint32 targetNetworkNumber)
func (_ContractSettlement *ContractSettlementCallerSession) AllOrderDetails(arg0 uint32) (struct {
	OrderId             uint32
	Maker               common.Address
	Taker               common.Address
	InputToken          common.Address
	InputAmount         *big.Int
	OutputToken         common.Address
	OutputAmount        *big.Int
	Expiry              *big.Int
	TargetNetworkNumber uint32
}, error) {
	return _ContractSettlement.Contract.AllOrderDetails(&_ContractSettlement.CallOpts, arg0)
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

// HashOrder is a free data retrieval call binding the contract method 0x3b692c09.
//
// Solidity: function hashOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
func (_ContractSettlement *ContractSettlementCaller) HashOrder(opts *bind.CallOpts, order ISettlementOrder) ([32]byte, error) {
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
func (_ContractSettlement *ContractSettlementSession) HashOrder(order ISettlementOrder) ([32]byte, error) {
	return _ContractSettlement.Contract.HashOrder(&_ContractSettlement.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x3b692c09.
//
// Solidity: function hashOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
func (_ContractSettlement *ContractSettlementCallerSession) HashOrder(order ISettlementOrder) ([32]byte, error) {
	return _ContractSettlement.Contract.HashOrder(&_ContractSettlement.CallOpts, order)
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

// Verify is a free data retrieval call binding the contract method 0x298fb569.
//
// Solidity: function verify((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractSettlement *ContractSettlementCaller) Verify(opts *bind.CallOpts, order ISettlementOrder, signature []byte) (bool, error) {
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
func (_ContractSettlement *ContractSettlementSession) Verify(order ISettlementOrder, signature []byte) (bool, error) {
	return _ContractSettlement.Contract.Verify(&_ContractSettlement.CallOpts, order, signature)
}

// Verify is a free data retrieval call binding the contract method 0x298fb569.
//
// Solidity: function verify((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractSettlement *ContractSettlementCallerSession) Verify(order ISettlementOrder, signature []byte) (bool, error) {
	return _ContractSettlement.Contract.Verify(&_ContractSettlement.CallOpts, order, signature)
}

// Fulfill is a paid mutator transaction binding the contract method 0x658d6341.
//
// Solidity: function fulfill((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, bytes sig) returns()
func (_ContractSettlement *ContractSettlementTransactor) Fulfill(opts *bind.TransactOpts, order ISettlementOrder, quorumThresholdPercentage uint32, quorumNumbers []byte, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "fulfill", order, quorumThresholdPercentage, quorumNumbers, sig)
}

// Fulfill is a paid mutator transaction binding the contract method 0x658d6341.
//
// Solidity: function fulfill((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, bytes sig) returns()
func (_ContractSettlement *ContractSettlementSession) Fulfill(order ISettlementOrder, quorumThresholdPercentage uint32, quorumNumbers []byte, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, order, quorumThresholdPercentage, quorumNumbers, sig)
}

// Fulfill is a paid mutator transaction binding the contract method 0x658d6341.
//
// Solidity: function fulfill((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, bytes sig) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) Fulfill(order ISettlementOrder, quorumThresholdPercentage uint32, quorumNumbers []byte, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, order, quorumThresholdPercentage, quorumNumbers, sig)
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

// RespondToFulfill is a paid mutator transaction binding the contract method 0x377a91a6.
//
// Solidity: function respondToFulfill(bytes quorumNumbers, uint32 quorumThresholdPercentage, uint32 createdBlock, (address,uint32) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSettlement *ContractSettlementTransactor) RespondToFulfill(opts *bind.TransactOpts, quorumNumbers []byte, quorumThresholdPercentage uint32, createdBlock uint32, orderResponse ISettlementOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "respondToFulfill", quorumNumbers, quorumThresholdPercentage, createdBlock, orderResponse, nonSignerStakesAndSignature)
}

// RespondToFulfill is a paid mutator transaction binding the contract method 0x377a91a6.
//
// Solidity: function respondToFulfill(bytes quorumNumbers, uint32 quorumThresholdPercentage, uint32 createdBlock, (address,uint32) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSettlement *ContractSettlementSession) RespondToFulfill(quorumNumbers []byte, quorumThresholdPercentage uint32, createdBlock uint32, orderResponse ISettlementOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSettlement.Contract.RespondToFulfill(&_ContractSettlement.TransactOpts, quorumNumbers, quorumThresholdPercentage, createdBlock, orderResponse, nonSignerStakesAndSignature)
}

// RespondToFulfill is a paid mutator transaction binding the contract method 0x377a91a6.
//
// Solidity: function respondToFulfill(bytes quorumNumbers, uint32 quorumThresholdPercentage, uint32 createdBlock, (address,uint32) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) RespondToFulfill(quorumNumbers []byte, quorumThresholdPercentage uint32, createdBlock uint32, orderResponse ISettlementOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSettlement.Contract.RespondToFulfill(&_ContractSettlement.TransactOpts, quorumNumbers, quorumThresholdPercentage, createdBlock, orderResponse, nonSignerStakesAndSignature)
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
	Order                     ISettlementOrder
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	Recipient                 common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterFulfillEvent is a free log retrieval operation binding the contract event 0x94a6dd0de15ac83f39634005b9c02f69524f7797c69b9f7c1f8a95a5d527391c.
//
// Solidity: event FulfillEvent((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
func (_ContractSettlement *ContractSettlementFilterer) FilterFulfillEvent(opts *bind.FilterOpts) (*ContractSettlementFulfillEventIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "FulfillEvent")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementFulfillEventIterator{contract: _ContractSettlement.contract, event: "FulfillEvent", logs: logs, sub: sub}, nil
}

// WatchFulfillEvent is a free log subscription operation binding the contract event 0x94a6dd0de15ac83f39634005b9c02f69524f7797c69b9f7c1f8a95a5d527391c.
//
// Solidity: event FulfillEvent((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
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

// ParseFulfillEvent is a log parse operation binding the contract event 0x94a6dd0de15ac83f39634005b9c02f69524f7797c69b9f7c1f8a95a5d527391c.
//
// Solidity: event FulfillEvent((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
func (_ContractSettlement *ContractSettlementFilterer) ParseFulfillEvent(log types.Log) (*ContractSettlementFulfillEvent, error) {
	event := new(ContractSettlementFulfillEvent)
	if err := _ContractSettlement.contract.UnpackLog(event, "FulfillEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	OrderResponse         ISettlementOrderResponse
	OrderResponseMetadata ISettlementOrderResponseMetadata
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

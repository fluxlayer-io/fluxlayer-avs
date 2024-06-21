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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderDetails\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderExecutions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfill\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structISettlement.Order\",\"components\":[{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structISettlement.Order\",\"components\":[{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToFulfill\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderResponse\",\"type\":\"tuple\",\"internalType\":\"structISettlement.OrderResponse\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structISettlement.Order\",\"components\":[{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FulfillEvent\",\"inputs\":[{\"name\":\"sig\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"orderId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"order\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISettlement.Order\",\"components\":[{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderRespondedEvent\",\"inputs\":[{\"name\":\"orderResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISettlement.OrderResponse\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"OrderResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISettlement.OrderResponseMetadata\",\"components\":[{\"name\":\"orderResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TakerMismatch\",\"inputs\":[]}]",
	Bin: "0x6101c06040523480156200001257600080fd5b5060405162006303380380620063038339810160408190526200003591620002d4565b6040518060400160405280600a81526020016914d95d1d1b195b595b9d60b21b815250604051806040016040528060038152602001620312e360ec1b81525082806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000ce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000f49190620002d4565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200014c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001729190620002d4565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001cc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001f29190620002d4565b6001600160a01b031660e052506097805460ff19166001179055815160208084019190912082519183019190912061016082905261018081905246610120527f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f620002a28184846040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b6101005230610140526101a05250620002fb9350505050565b6001600160a01b0381168114620002d157600080fd5b50565b600060208284031215620002e757600080fd5b8151620002f481620002bb565b9392505050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a051615f54620003af60003960006141b601526000614205015260006141e001526000614139015260006141630152600061418d01526000818161059301526128b001526000818161039e0152612a920152600081816103c501528181612c680152612e2a0152600081816103ec015281816116a00152818161257b01528181612713015261294d0152615f546000f3fe608060405234801561001057600080fd5b50600436106101ef5760003560e01c8063683048351161010f578063b98d0908116100a2578063df5cf72311610071578063df5cf7231461058e578063e772935f146105b5578063f2fde38b146105d5578063fabc1cbc146105e857600080fd5b8063b98d09081461052d578063c0c53b8b1461053a578063cefdc1d41461054d578063dd3285b31461056e57600080fd5b8063886f1195116100de578063886f1195146104375780638da5cb5b1461044a5780638f9b1a371461045b578063aa4f5cc61461046e57600080fd5b806368304835146103c05780636d14a987146103e75780636efb46361461040e578063715018a61461042f57600080fd5b8063377a91a6116101875780635ac86ab7116101565780635ac86ab71461033e5780635c155662146103715780635c975abb146103915780635df459461461039957600080fd5b8063377a91a6146102f0578063416c7e5e146103035780634f739f7414610316578063595c6a671461033657600080fd5b8063245a7bfc116101c3578063245a7bfc1461027157806329d81f491461029c578063322ed59d146102bd5780633563b0d1146102d057600080fd5b80629e8b57146101f457806310d67a2f1461021f578063136439dd14610234578063171f1d5b14610247575b600080fd5b610207610202366004614969565b6105fb565b60405161021693929190614986565b60405180910390f35b61023261022d366004614a08565b6106b3565b005b610232610242366004614a25565b61076f565b61025a610255366004614ba3565b6108ae565b604080519215158352901515602083015201610216565b60ce54610284906001600160a01b031681565b6040516001600160a01b039091168152602001610216565b6102af6102aa366004614c0d565b610a38565b604051908152602001610216565b6102326102cb366004614c6b565b610b38565b6102e36102de366004614d02565b610ee9565b6040516102169190614e5d565b6102326102fe366004615114565b61137f565b6102326103113660046151d0565b61169e565b6103296103243660046151ed565b611813565b60405161021691906152f1565b610232611f39565b61036161034c3660046153bb565b606654600160ff9092169190911b9081161490565b6040519015158152602001610216565b61038461037f3660046153d8565b612000565b6040516102169190615484565b6066546102af565b6102847f000000000000000000000000000000000000000000000000000000000000000081565b6102847f000000000000000000000000000000000000000000000000000000000000000081565b6102847f000000000000000000000000000000000000000000000000000000000000000081565b61042161041c3660046154c8565b6121c8565b604051610216929190615588565b6102326130df565b606554610284906001600160a01b031681565b6033546001600160a01b0316610284565b6103616104693660046155d1565b6130f3565b6104d561047c366004614969565b60cb60205260009081526040902080546001820154600283015460038401546004850154600586015460068701546007909701546001600160a01b03968716979587169694851695939490921692909163ffffffff1688565b604080516001600160a01b03998a16815297891660208901529588169587019590955260608601939093529416608084015260a083019390935260c082019290925263ffffffff90911660e082015261010001610216565b6097546103619060ff1681565b610232610548366004615626565b613166565b61056061055b366004615671565b6132a1565b6040516102169291906156a8565b6102af61057c366004614969565b60ca6020526000908152604090205481565b6102847f000000000000000000000000000000000000000000000000000000000000000081565b6102af6105c3366004614969565b60cd6020526000908152604090205481565b6102326105e3366004614a08565b613433565b6102326105f6366004614a25565b6134a9565b60cc602052600090815260409020805460018201805463ffffffff9092169291610624906156c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610650906156c9565b801561069d5780601f106106725761010080835404028352916020019161069d565b820191906000526020600020905b81548152906001019060200180831161068057829003601f168201915b5050506002909301549192505063ffffffff1683565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610706573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072a91906156fe565b6001600160a01b0316336001600160a01b0316146107635760405162461bcd60e51b815260040161075a9061571b565b60405180910390fd5b61076c81613605565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156107b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107db9190615765565b6107f75760405162461bcd60e51b815260040161075a90615782565b606654818116146108705760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161075a565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106108f6576108f66157ca565b60200201518951600160200201518a6020015160006002811061091b5761091b6157ca565b60200201518b60200151600160028110610937576109376157ca565b602090810291909101518c518d8301516040516109949a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6109b791906157e0565b9050610a2a6109d06109c988846136fc565b8690613793565b6109d8613827565b610a20610a1185610a0b604080518082018252600080825260209182015281518083019092526001825260029082015290565b906136fc565b610a1a8c6138e7565b90613793565b886201d4c0613977565b909890975095505050505050565b6000610b327f47d2f645193c5bc1f8819c38495c8e2c7f69ea073ec2139744d2d6ff5208583d610a6b6020850185614a08565b610a7b6040860160208701614a08565b610a8b6060870160408801614a08565b6060870135610aa060a0890160808a01614a08565b60a089013560c08a0135610abb6101008c0160e08d01614969565b60408051602081019a909a526001600160a01b03988916908a01529587166060890152938616608088015260a087019290925290931660c085015260e084019290925261010083019190915263ffffffff166101208201526101400160405160208183030381529060405280519060200120613b9b565b92915050565b6000610b4a6040880160208901614a08565b6001600160a01b031614158015610b825750610b6c6040870160208801614a08565b6001600160a01b0316336001600160a01b031614155b15610ba057604051630d9d754f60e41b815260040160405180910390fd5b428660c001351015610bc5576040516362b439dd60e11b815260040160405180910390fd5b610bd08683836130f3565b610bed57604051638baa579f60e01b815260040160405180910390fd5b610bfd6060870160408801614a08565b6001600160a01b03166323b872dd610c186020890189614a08565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152606089013560448201526064016020604051808303816000875af1158015610c6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c929190615765565b50610ca360a0870160808801614a08565b6001600160a01b03166323b872dd33610cbf60208a018a614a08565b6040516001600160e01b031960e085901b1681526001600160a01b0392831660048201529116602482015260a089013560448201526064016020604051808303816000875af1158015610d16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d3a9190615765565b5060405180606001604052808663ffffffff16815260200185858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050504363ffffffff90811660209384015260c9548116825260cc835260409091208351815463ffffffff191692169190911781558282015180519192610dd6926001850192909101906147d9565b50604091820151600291909101805463ffffffff191663ffffffff90921691909117905551610e09908790602001615899565b60408051601f19818403018152918152815160209283012060c9805463ffffffff908116600090815260ca865284812093909355905416815260cb90925290208690610e5582826158e2565b505060c9546040517f4c04f413243ecfd9c1ddddf799c525affe9c8fe07aa9fa7ce49bc953ad89e25891610e9f918591859163ffffffff909116908b908b908b908b9033906159b5565b60405180910390a160c9805460019190600090610ec390849063ffffffff16615a37565b92506101000a81548163ffffffff021916908363ffffffff160217905550505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4f91906156fe565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fb591906156fe565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ff7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061101b91906156fe565b9050600086516001600160401b0381111561103857611038614a3e565b60405190808252806020026020018201604052801561106b57816020015b60608152602001906001900390816110565790505b50905060005b875181101561137357600088828151811061108e5761108e6157ca565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa1580156110ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111179190810190615a5f565b905080516001600160401b0381111561113257611132614a3e565b60405190808252806020026020018201604052801561117d57816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816111505790505b50848481518110611190576111906157ca565b602002602001018190525060005b815181101561135d576040518060600160405280876001600160a01b03166347b314e88585815181106111d3576111d36157ca565b60200260200101516040518263ffffffff1660e01b81526004016111f991815260200190565b602060405180830381865afa158015611216573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123a91906156fe565b6001600160a01b0316815260200183838151811061125a5761125a6157ca565b60200260200101518152602001896001600160a01b031663fa28c627858581518110611288576112886157ca565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa1580156112e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113089190615aef565b6001600160601b0316815250858581518110611326576113266157ca565b6020026020010151828151811061133f5761133f6157ca565b6020026020010181905250808061135590615b18565b91505061119e565b505050808061136b90615b18565b915050611071565b50979650505050505050565b60ce546001600160a01b031633146113d95760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c6572000000604482015260640161075a565b600060cd816113ee6040860160208701614969565b63ffffffff1663ffffffff168152602001908152602001600020541461146c5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526c103a37903a34329037b93232b960991b606482015260840161075a565b60008260405160200161147f9190615b68565b6040516020818303038152906040528051906020012090506000806114a7838a8a89886121c8565b91509150600060405180604001604052804363ffffffff16815260200183815250905085816040516020016114dd929190615b76565b6040516020818303038152906040528051906020012060cd600088602001602081019061150a9190614969565b63ffffffff1663ffffffff16815260200190815260200160002081905550600060cb60008860200160208101906115419190614969565b63ffffffff9081168252602080830193909352604091820160002082516101008101845281546001600160a01b039081168252600183015481168287015260028301548116948201859052600383015460608301526004830154166080820152600582015460a0820152600682015460c082015260079091015490911660e08201529250906323b872dd9030906115da908b018b614a08565b60608501516040516001600160e01b031960e086901b1681526001600160a01b03938416600482015292909116602483015260448201526064016020604051808303816000875af1158015611633573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116579190615765565b507fa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd8783604051611689929190615b76565b60405180910390a15050505050505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116fc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061172091906156fe565b6001600160a01b0316336001600160a01b0316146117cc5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a40161075a565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b61183e6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561187e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118a291906156fe565b90506118cf6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906118ff908b9089908990600401615ba2565b600060405180830381865afa15801561191c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119449190810190615bec565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611976908b908b908b90600401615c7a565b600060405180830381865afa158015611993573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119bb9190810190615bec565b6040820152856001600160401b038111156119d8576119d8614a3e565b604051908082528060200260200182016040528015611a0b57816020015b60608152602001906001900390816119f65790505b50606082015260005b60ff8116871115611e4a576000856001600160401b03811115611a3957611a39614a3e565b604051908082528060200260200182016040528015611a62578160200160208202803683370190505b5083606001518360ff1681518110611a7c57611a7c6157ca565b602002602001018190525060005b86811015611d4a5760008c6001600160a01b03166304ec63518a8a85818110611ab557611ab56157ca565b905060200201358e88600001518681518110611ad357611ad36157ca565b60200260200101516040518463ffffffff1660e01b8152600401611b109392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611b2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b519190615ca3565b90506001600160c01b038116611bf55760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a40161075a565b8a8a8560ff16818110611c0a57611c0a6157ca565b6001600160c01b03841692013560f81c9190911c600190811614159050611d3757856001600160a01b031663dd9846b98a8a85818110611c4c57611c4c6157ca565b905060200201358d8d8860ff16818110611c6857611c686157ca565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611cbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ce29190615ccc565b85606001518560ff1681518110611cfb57611cfb6157ca565b60200260200101518481518110611d1457611d146157ca565b63ffffffff9092166020928302919091019091015282611d3381615b18565b9350505b5080611d4281615b18565b915050611a8a565b506000816001600160401b03811115611d6557611d65614a3e565b604051908082528060200260200182016040528015611d8e578160200160208202803683370190505b50905060005b82811015611e0f5784606001518460ff1681518110611db557611db56157ca565b60200260200101518181518110611dce57611dce6157ca565b6020026020010151828281518110611de857611de86157ca565b63ffffffff9092166020928302919091019091015280611e0781615b18565b915050611d94565b508084606001518460ff1681518110611e2a57611e2a6157ca565b602002602001018190525050508080611e4290615ce9565b915050611a14565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eaf91906156fe565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611ee2908b908b908e90600401615d09565b600060405180830381865afa158015611eff573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611f279190810190615bec565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611f81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fa59190615765565b611fc15760405162461bcd60e51b815260040161075a90615782565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401612032929190615d33565b600060405180830381865afa15801561204f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120779190810190615bec565b9050600084516001600160401b0381111561209457612094614a3e565b6040519080825280602002602001820160405280156120bd578160200160208202803683370190505b50905060005b85518110156121be57866001600160a01b03166304ec63518783815181106120ed576120ed6157ca565b602002602001015187868581518110612108576121086157ca565b60200260200101516040518463ffffffff1660e01b81526004016121459392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612162573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121869190615ca3565b6001600160c01b03168282815181106121a1576121a16157ca565b6020908102919091010152806121b681615b18565b9150506120c3565b5095945050505050565b604080518082019091526060808252602082015260008461223f5760405162461bcd60e51b81526020600482015260376024820152600080516020615eff83398151915260448201527f7265733a20656d7074792071756f72756d20696e707574000000000000000000606482015260840161075a565b60408301515185148015612257575060a08301515185145b8015612267575060c08301515185145b8015612277575060e08301515185145b6122e15760405162461bcd60e51b81526020600482015260416024820152600080516020615eff83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a40161075a565b825151602084015151146123595760405162461bcd60e51b815260206004820152604460248201819052600080516020615eff833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a40161075a565b4363ffffffff168463ffffffff16106123c85760405162461bcd60e51b815260206004820152603c6024820152600080516020615eff83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b00000000606482015260840161075a565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561240957612409614a3e565b604051908082528060200260200182016040528015612432578160200160208202803683370190505b506020820152866001600160401b0381111561245057612450614a3e565b604051908082528060200260200182016040528015612479578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156124ad576124ad614a3e565b6040519080825280602002602001820160405280156124d6578160200160208202803683370190505b5081526020860151516001600160401b038111156124f6576124f6614a3e565b60405190808252806020026020018201604052801561251f578160200160208202803683370190505b50816020018190525060006125f18a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156125c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125ec9190615d87565b613be9565b905060005b87602001515181101561288c5761263b8860200151828151811061261c5761261c6157ca565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110612651576126516157ca565b60209081029190910101528015612711576020830151612672600183615da4565b81518110612682576126826157ca565b602002602001015160001c836020015182815181106126a3576126a36157ca565b602002602001015160001c11612711576040805162461bcd60e51b8152602060048201526024810191909152600080516020615eff83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f72746564606482015260840161075a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110612756576127566157ca565b60200260200101518b8b600001518581518110612775576127756157ca565b60200260200101516040518463ffffffff1660e01b81526004016127b29392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156127cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127f39190615ca3565b6001600160c01b031683600001518281518110612812576128126157ca565b6020026020010181815250506128786109c961284c848660000151858151811061283e5761283e6157ca565b602002602001015116613c7a565b8a602001518481518110612862576128626157ca565b6020026020010151613ca590919063ffffffff16565b94508061288481615b18565b9150506125f6565b505061289783613d89565b60975490935060ff166000816128ae576000612930565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa15801561290c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129309190615dbb565b905060005b8a811015612fae578215612a90578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061298c5761298c6157ca565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156129cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129f09190615dbb565b6129fa9190615dd4565b11612a905760405162461bcd60e51b81526020600482015260666024820152600080516020615eff83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c40161075a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612ad157612ad16157ca565b9050013560f81c60f81b60f81c8c8c60a001518581518110612af557612af56157ca565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612b51573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b759190615dec565b6001600160401b031916612b988a60400151838151811061261c5761261c6157ca565b67ffffffffffffffff191614612c345760405162461bcd60e51b81526020600482015260616024820152600080516020615eff83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c40161075a565b612c6489604001518281518110612c4d57612c4d6157ca565b60200260200101518761379390919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612ca757612ca76157ca565b9050013560f81c60f81b60f81c8c8c60c001518581518110612ccb57612ccb6157ca565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612d27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d4b9190615aef565b85602001518281518110612d6157612d616157ca565b6001600160601b03909216602092830291909101820152850151805182908110612d8d57612d8d6157ca565b602002602001015185600001518281518110612dab57612dab6157ca565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612f9957612e2386600001518281518110612df557612df56157ca565b60200260200101518f8f86818110612e0f57612e0f6157ca565b600192013560f81c9290921c811614919050565b15612f87577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612e6957612e696157ca565b9050013560f81c60f81b60f81c8e89602001518581518110612e8d57612e8d6157ca565b60200260200101518f60e001518881518110612eab57612eab6157ca565b60200260200101518781518110612ec457612ec46157ca565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612f28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f4c9190615aef565b8751805185908110612f6057612f606157ca565b60200260200101818151612f749190615e17565b6001600160601b03169052506001909101905b80612f9181615b18565b915050612dcf565b50508080612fa690615b18565b915050612935565b505050600080612fc88c868a606001518b608001516108ae565b91509150816130395760405162461bcd60e51b81526020600482015260436024820152600080516020615eff83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a40161075a565b8061309a5760405162461bcd60e51b81526020600482015260396024820152600080516020615eff83398151915260448201527f7265733a207369676e617475726520697320696e76616c696400000000000000606482015260840161075a565b505060008782602001516040516020016130b5929190615e3f565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6130e7613e24565b6130f16000613e7e565b565b60008061314184848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061313b9250899150610a389050565b90613ed0565b90506131506020860186614a08565b6001600160a01b03918216911614949350505050565b600054610100900460ff16158080156131865750600054600160ff909116105b806131a05750303b1580156131a0575060005460ff166001145b6132035760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161075a565b6000805460ff191660011790558015613226576000805461ff0019166101001790555b613231846000613ef4565b61323a83613e7e565b60ce80546001600160a01b0319166001600160a01b038416179055801561329b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106132dc576132dc6157ca565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906133189088908690600401615d33565b600060405180830381865afa158015613335573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261335d9190810190615bec565b60008151811061336f5761336f6157ca565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156133db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133ff9190615ca3565b6001600160c01b03169050600061341582613fde565b9050816134238a838a610ee9565b9550955050505050935093915050565b61343b613e24565b6001600160a01b0381166134a05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161075a565b61076c81613e7e565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134fc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061352091906156fe565b6001600160a01b0316336001600160a01b0316146135505760405162461bcd60e51b815260040161075a9061571b565b6066541981196066541916146135ce5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161075a565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016108a3565b6001600160a01b0381166136935760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161075a565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261371861485d565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561374b5761374d565bfe5b508061378b5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b604482015260640161075a565b505092915050565b60408051808201909152600080825260208201526137af61487b565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561374b57508061378b5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b604482015260640161075a565b61382f614899565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613917600080516020615edf833981519152866157e0565b90505b613923816140aa565b9093509150600080516020615edf83398151915282830983141561395d576040805180820190915290815260208101919091529392505050565b600080516020615edf83398151915260018208905061391a565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906139a96148be565b60005b6002811015613b6e5760006139c2826006615e87565b90508482600281106139d6576139d66157ca565b602002015151836139e8836000615dd4565b600c81106139f8576139f86157ca565b6020020152848260028110613a0f57613a0f6157ca565b60200201516020015183826001613a269190615dd4565b600c8110613a3657613a366157ca565b6020020152838260028110613a4d57613a4d6157ca565b6020020151515183613a60836002615dd4565b600c8110613a7057613a706157ca565b6020020152838260028110613a8757613a876157ca565b6020020151516001602002015183613aa0836003615dd4565b600c8110613ab057613ab06157ca565b6020020152838260028110613ac757613ac76157ca565b602002015160200151600060028110613ae257613ae26157ca565b602002015183613af3836004615dd4565b600c8110613b0357613b036157ca565b6020020152838260028110613b1a57613b1a6157ca565b602002015160200151600160028110613b3557613b356157ca565b602002015183613b46836005615dd4565b600c8110613b5657613b566157ca565b60200201525080613b6681615b18565b9150506139ac565b50613b776148dd565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6000610b32613ba861412c565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b600080613bf584614253565b9050808360ff166001901b11613c735760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c756500606482015260840161075a565b9392505050565b6000805b8215610b3257613c8f600184615da4565b9092169180613c9d81615ea6565b915050613c7e565b60408051808201909152600080825260208201526102008261ffff1610613d015760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b604482015260640161075a565b8161ffff1660011415613d15575081610b32565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613d7e57600161ffff871660ff83161c81161415613d6157613d5e8484613793565b93505b613d6b8384613793565b92506201fffe600192831b169101613d31565b509195945050505050565b60408051808201909152600080825260208201528151158015613dae57506020820151155b15613dcc575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615edf8339815191528460200151613dff91906157e0565b613e1790600080516020615edf833981519152615da4565b905292915050565b919050565b6033546001600160a01b031633146130f15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161075a565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000806000613edf85856143e0565b91509150613eec81614450565b509392505050565b6065546001600160a01b0316158015613f1557506001600160a01b03821615155b613f975760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161075a565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613fda82613605565b5050565b6060600080613fec84613c7a565b61ffff166001600160401b0381111561400757614007614a3e565b6040519080825280601f01601f191660200182016040528015614031576020820181803683370190505b5090506000805b825182108015614049575061010081105b156140a0576001811b935085841615614090578060f81b838381518110614072576140726157ca565b60200101906001600160f81b031916908160001a9053508160010191505b61409981615b18565b9050614038565b5090949350505050565b60008080600080516020615edf8339815191526003600080516020615edf83398151915286600080516020615edf833981519152888909090890506000614120827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615edf83398151915261460b565b91959194509092505050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561418557507f000000000000000000000000000000000000000000000000000000000000000046145b156141af57507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b6000610100825111156142dc5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a40161075a565b81516142ea57506000919050565b60008083600081518110614300576143006157ca565b0160200151600160f89190911c81901b92505b84518110156143d75784818151811061432e5761432e6157ca565b0160200151600160f89190911c1b91508282116143c35760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a40161075a565b918117916143d081615b18565b9050614313565b50909392505050565b6000808251604114156144175760208301516040840151606085015160001a61440b878285856146b3565b94509450505050614449565b82516040141561444157602083015160408401516144368683836147a0565b935093505050614449565b506000905060025b9250929050565b600081600481111561446457614464615ec8565b141561446d5750565b600181600481111561448157614481615ec8565b14156144cf5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161075a565b60028160048111156144e3576144e3615ec8565b14156145315760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161075a565b600381600481111561454557614545615ec8565b141561459e5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161075a565b60048160048111156145b2576145b2615ec8565b141561076c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161075a565b6000806146166148dd565b61461e6148fb565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561374b5750826146a85760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c757265000000000000604482015260640161075a565b505195945050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156146ea5750600090506003614797565b8460ff16601b1415801561470257508460ff16601c14155b156147135750600090506004614797565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614767573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661479057600060019250925050614797565b9150600090505b94509492505050565b6000806001600160ff1b038316816147bd60ff86901c601b615dd4565b90506147cb878288856146b3565b935093505050935093915050565b8280546147e5906156c9565b90600052602060002090601f016020900481019282614807576000855561484d565b82601f1061482057805160ff191683800117855561484d565b8280016001018555821561484d579182015b8281111561484d578251825591602001919060010190614832565b50614859929150614919565b5090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806148ac61492e565b81526020016148b961492e565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b5b80821115614859576000815560010161491a565b60405180604001604052806002906020820280368337509192915050565b63ffffffff8116811461076c57600080fd5b8035613e1f8161494c565b60006020828403121561497b57600080fd5b8135613c738161494c565b600063ffffffff80861683526020606081850152855180606086015260005b818110156149c1578781018301518682016080015282016149a5565b818111156149d3576000608083880101525b509490911660408401525050601f91909101601f19160160800192915050565b6001600160a01b038116811461076c57600080fd5b600060208284031215614a1a57600080fd5b8135613c73816149f3565b600060208284031215614a3757600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614a7657614a76614a3e565b60405290565b60405161010081016001600160401b0381118282101715614a7657614a76614a3e565b604051601f8201601f191681016001600160401b0381118282101715614ac757614ac7614a3e565b604052919050565b600060408284031215614ae157600080fd5b614ae9614a54565b9050813581526020820135602082015292915050565b600082601f830112614b1057600080fd5b604051604081018181106001600160401b0382111715614b3257614b32614a3e565b8060405250806040840185811115614b4957600080fd5b845b81811015613d7e578035835260209283019201614b4b565b600060808284031215614b7557600080fd5b614b7d614a54565b9050614b898383614aff565b8152614b988360408401614aff565b602082015292915050565b6000806000806101208587031215614bba57600080fd5b84359350614bcb8660208701614acf565b9250614bda8660608701614b63565b9150614be98660e08701614acf565b905092959194509250565b60006101008284031215614c0757600080fd5b50919050565b60006101008284031215614c2057600080fd5b613c738383614bf4565b60008083601f840112614c3c57600080fd5b5081356001600160401b03811115614c5357600080fd5b60208301915083602082850101111561444957600080fd5b6000806000806000806101608789031215614c8557600080fd5b614c8f8888614bf4565b9550610100870135614ca08161494c565b94506101208701356001600160401b0380821115614cbd57600080fd5b614cc98a838b01614c2a565b9096509450610140890135915080821115614ce357600080fd5b50614cf089828a01614c2a565b979a9699509497509295939492505050565b600080600060608486031215614d1757600080fd5b8335614d22816149f3565b92506020848101356001600160401b0380821115614d3f57600080fd5b818701915087601f830112614d5357600080fd5b813581811115614d6557614d65614a3e565b614d77601f8201601f19168501614a9f565b91508082528884828501011115614d8d57600080fd5b8084840185840137600084828401015250809450505050614db06040850161495e565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614e4f578385038a52825180518087529087019087870190845b81811015614e3a57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614df6565b50509a87019a95505091850191600101614dd8565b509298975050505050505050565b602081526000613c736020830184614db9565b60006001600160401b03821115614e8957614e89614a3e565b5060051b60200190565b600082601f830112614ea457600080fd5b81356020614eb9614eb483614e70565b614a9f565b82815260059290921b84018101918181019086841115614ed857600080fd5b8286015b84811015614efc578035614eef8161494c565b8352918301918301614edc565b509695505050505050565b600082601f830112614f1857600080fd5b81356020614f28614eb483614e70565b82815260069290921b84018101918181019086841115614f4757600080fd5b8286015b84811015614efc57614f5d8882614acf565b835291830191604001614f4b565b600082601f830112614f7c57600080fd5b81356020614f8c614eb483614e70565b82815260059290921b84018101918181019086841115614fab57600080fd5b8286015b84811015614efc5780356001600160401b03811115614fce5760008081fd5b614fdc8986838b0101614e93565b845250918301918301614faf565b60006101808284031215614ffd57600080fd5b615005614a7c565b905081356001600160401b038082111561501e57600080fd5b61502a85838601614e93565b8352602084013591508082111561504057600080fd5b61504c85838601614f07565b6020840152604084013591508082111561506557600080fd5b61507185838601614f07565b60408401526150838560608601614b63565b60608401526150958560e08601614acf565b60808401526101208401359150808211156150af57600080fd5b6150bb85838601614e93565b60a08401526101408401359150808211156150d557600080fd5b6150e185838601614e93565b60c08401526101608401359150808211156150fb57600080fd5b5061510884828501614f6b565b60e08301525092915050565b60008060008060008086880360c081121561512e57600080fd5b87356001600160401b038082111561514557600080fd5b6151518b838c01614c2a565b909950975060208a013591506151668261494c565b9095506040890135906151788261494c565b8195506040605f198401121561518d57600080fd5b60608a01945060a08a01359250808311156151a757600080fd5b50506151b589828a01614fea565b9150509295509295509295565b801515811461076c57600080fd5b6000602082840312156151e257600080fd5b8135613c73816151c2565b6000806000806000806080878903121561520657600080fd5b8635615211816149f3565b955060208701356152218161494c565b945060408701356001600160401b038082111561523d57600080fd5b6152498a838b01614c2a565b9096509450606089013591508082111561526257600080fd5b818901915089601f83011261527657600080fd5b81358181111561528557600080fd5b8a60208260051b850101111561529a57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b838110156152e657815163ffffffff16875295820195908201906001016152c4565b509495945050505050565b60006020808352835160808285015261530d60a08501826152b0565b905081850151601f198086840301604087015261532a83836152b0565b9250604087015191508086840301606087015261534783836152b0565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561539e578487830301845261538c8287516152b0565b95880195938801939150600101615372565b509998505050505050505050565b60ff8116811461076c57600080fd5b6000602082840312156153cd57600080fd5b8135613c73816153ac565b6000806000606084860312156153ed57600080fd5b83356153f8816149f3565b92506020848101356001600160401b0381111561541457600080fd5b8501601f8101871361542557600080fd5b8035615433614eb482614e70565b81815260059190911b8201830190838101908983111561545257600080fd5b928401925b8284101561547057833582529284019290840190615457565b8096505050505050614db06040850161495e565b6020808252825182820181905260009190848201906040850190845b818110156154bc578351835292840192918401916001016154a0565b50909695505050505050565b6000806000806000608086880312156154e057600080fd5b8535945060208601356001600160401b03808211156154fe57600080fd5b61550a89838a01614c2a565b90965094506040880135915061551f8261494c565b9092506060870135908082111561553557600080fd5b5061554288828901614fea565b9150509295509295909350565b600081518084526020808501945080840160005b838110156152e65781516001600160601b031687529582019590820190600101615563565b60408152600083516040808401526155a3608084018261554f565b90506020850151603f198483030160608501526155c0828261554f565b925050508260208301529392505050565b600080600061012084860312156155e757600080fd5b6155f18585614bf4565b92506101008401356001600160401b0381111561560d57600080fd5b61561986828701614c2a565b9497909650939450505050565b60008060006060848603121561563b57600080fd5b8335615646816149f3565b92506020840135615656816149f3565b91506040840135615666816149f3565b809150509250925092565b60008060006060848603121561568657600080fd5b8335615691816149f3565b92506020840135915060408401356156668161494c565b8281526040602082015260006156c16040830184614db9565b949350505050565b600181811c908216806156dd57607f821691505b60208210811415614c0757634e487b7160e01b600052602260045260246000fd5b60006020828403121561571057600080fd5b8151613c73816149f3565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561577757600080fd5b8151613c73816151c2565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000826157fd57634e487b7160e01b600052601260045260246000fd5b500690565b803561580d816149f3565b6001600160a01b039081168352602082013590615829826149f3565b9081166020840152604082013590615840826149f3565b80821660408501526060830135606085015260808301359150615862826149f3565b16608083015260a0818101359083015260c0808201359083015261588860e0820161495e565b63ffffffff811660e0840152505050565b6101008101610b328284615802565b60008135610b32816149f3565b80546001600160a01b0319166001600160a01b0392909216919091179055565b60008135610b328161494c565b81356158ed816149f3565b6158f781836158b5565b506020820135615906816149f3565b61591381600184016158b5565b5061592c615923604084016158a8565b600283016158b5565b6060820135600382015561594e615945608084016158a8565b600483016158b5565b60a0820135600582015560c08201356006820155613fda61597160e084016158d5565b6007830163ffffffff821663ffffffff198254161781555050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60006101a08083526159ca8184018b8d61598c565b905063ffffffff808a1660208501526159e6604085018a615802565b8716610140840152828103610160840152615a0281868861598c565b91505060018060a01b0383166101808301529998505050505050505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff808316818516808303821115615a5657615a56615a21565b01949350505050565b60006020808385031215615a7257600080fd5b82516001600160401b03811115615a8857600080fd5b8301601f81018513615a9957600080fd5b8051615aa7614eb482614e70565b81815260059190911b82018301908381019087831115615ac657600080fd5b928401925b82841015615ae457835182529284019290840190615acb565b979650505050505050565b600060208284031215615b0157600080fd5b81516001600160601b0381168114613c7357600080fd5b6000600019821415615b2c57615b2c615a21565b5060010190565b8035615b3e816149f3565b6001600160a01b031682526020810135615b578161494c565b63ffffffff81166020840152505050565b60408101610b328284615b33565b60808101615b848285615b33565b63ffffffff8351166040830152602083015160608301529392505050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115615bcf57600080fd5b8260051b8085606085013760009201606001918252509392505050565b60006020808385031215615bff57600080fd5b82516001600160401b03811115615c1557600080fd5b8301601f81018513615c2657600080fd5b8051615c34614eb482614e70565b81815260059190911b82018301908381019087831115615c5357600080fd5b928401925b82841015615ae4578351615c6b8161494c565b82529284019290840190615c58565b63ffffffff84168152604060208201526000615c9a60408301848661598c565b95945050505050565b600060208284031215615cb557600080fd5b81516001600160c01b0381168114613c7357600080fd5b600060208284031215615cde57600080fd5b8151613c738161494c565b600060ff821660ff811415615d0057615d00615a21565b60010192915050565b604081526000615d1d60408301858761598c565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b81811015615d7a57845183529383019391830191600101615d5e565b5090979650505050505050565b600060208284031215615d9957600080fd5b8151613c73816153ac565b600082821015615db657615db6615a21565b500390565b600060208284031215615dcd57600080fd5b5051919050565b60008219821115615de757615de7615a21565b500190565b600060208284031215615dfe57600080fd5b815167ffffffffffffffff1981168114613c7357600080fd5b60006001600160601b0383811690831681811015615e3757615e37615a21565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615e7a57815185529382019390820190600101615e5e565b5092979650505050505050565b6000816000190483118215151615615ea157615ea1615a21565b500290565b600061ffff80831681811415615ebe57615ebe615a21565b6001019392505050565b634e487b7160e01b600052602160045260246000fdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220eee60dce702fb66e037deb80f77ae0b75331bb1812898c39a98c42bd006ff63c64736f6c634300080c0033",
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
// Solidity: function allOrderDetails(uint32 ) view returns(address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint256 expiry, uint32 targetNetworkNumber)
func (_ContractSettlement *ContractSettlementCaller) AllOrderDetails(opts *bind.CallOpts, arg0 uint32) (struct {
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

	outstruct.Maker = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Taker = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.InputToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.InputAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OutputToken = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.OutputAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Expiry = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TargetNetworkNumber = *abi.ConvertType(out[7], new(uint32)).(*uint32)

	return *outstruct, err

}

// AllOrderDetails is a free data retrieval call binding the contract method 0xaa4f5cc6.
//
// Solidity: function allOrderDetails(uint32 ) view returns(address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint256 expiry, uint32 targetNetworkNumber)
func (_ContractSettlement *ContractSettlementSession) AllOrderDetails(arg0 uint32) (struct {
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
// Solidity: function allOrderDetails(uint32 ) view returns(address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint256 expiry, uint32 targetNetworkNumber)
func (_ContractSettlement *ContractSettlementCallerSession) AllOrderDetails(arg0 uint32) (struct {
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

// HashOrder is a free data retrieval call binding the contract method 0x29d81f49.
//
// Solidity: function hashOrder((address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
func (_ContractSettlement *ContractSettlementCaller) HashOrder(opts *bind.CallOpts, order ISettlementOrder) ([32]byte, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x29d81f49.
//
// Solidity: function hashOrder((address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
func (_ContractSettlement *ContractSettlementSession) HashOrder(order ISettlementOrder) ([32]byte, error) {
	return _ContractSettlement.Contract.HashOrder(&_ContractSettlement.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x29d81f49.
//
// Solidity: function hashOrder((address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
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

// Verify is a free data retrieval call binding the contract method 0x8f9b1a37.
//
// Solidity: function verify((address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractSettlement *ContractSettlementCaller) Verify(opts *bind.CallOpts, order ISettlementOrder, signature []byte) (bool, error) {
	var out []interface{}
	err := _ContractSettlement.contract.Call(opts, &out, "verify", order, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x8f9b1a37.
//
// Solidity: function verify((address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractSettlement *ContractSettlementSession) Verify(order ISettlementOrder, signature []byte) (bool, error) {
	return _ContractSettlement.Contract.Verify(&_ContractSettlement.CallOpts, order, signature)
}

// Verify is a free data retrieval call binding the contract method 0x8f9b1a37.
//
// Solidity: function verify((address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractSettlement *ContractSettlementCallerSession) Verify(order ISettlementOrder, signature []byte) (bool, error) {
	return _ContractSettlement.Contract.Verify(&_ContractSettlement.CallOpts, order, signature)
}

// Fulfill is a paid mutator transaction binding the contract method 0x322ed59d.
//
// Solidity: function fulfill((address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, bytes sig) returns()
func (_ContractSettlement *ContractSettlementTransactor) Fulfill(opts *bind.TransactOpts, order ISettlementOrder, quorumThresholdPercentage uint32, quorumNumbers []byte, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "fulfill", order, quorumThresholdPercentage, quorumNumbers, sig)
}

// Fulfill is a paid mutator transaction binding the contract method 0x322ed59d.
//
// Solidity: function fulfill((address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, bytes sig) returns()
func (_ContractSettlement *ContractSettlementSession) Fulfill(order ISettlementOrder, quorumThresholdPercentage uint32, quorumNumbers []byte, sig []byte) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, order, quorumThresholdPercentage, quorumNumbers, sig)
}

// Fulfill is a paid mutator transaction binding the contract method 0x322ed59d.
//
// Solidity: function fulfill((address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, bytes sig) returns()
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
	Sig                       []byte
	OrderId                   uint32
	Order                     ISettlementOrder
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	Recipient                 common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterFulfillEvent is a free log retrieval operation binding the contract event 0x4c04f413243ecfd9c1ddddf799c525affe9c8fe07aa9fa7ce49bc953ad89e258.
//
// Solidity: event FulfillEvent(bytes sig, uint32 orderId, (address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
func (_ContractSettlement *ContractSettlementFilterer) FilterFulfillEvent(opts *bind.FilterOpts) (*ContractSettlementFulfillEventIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "FulfillEvent")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementFulfillEventIterator{contract: _ContractSettlement.contract, event: "FulfillEvent", logs: logs, sub: sub}, nil
}

// WatchFulfillEvent is a free log subscription operation binding the contract event 0x4c04f413243ecfd9c1ddddf799c525affe9c8fe07aa9fa7ce49bc953ad89e258.
//
// Solidity: event FulfillEvent(bytes sig, uint32 orderId, (address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
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

// ParseFulfillEvent is a log parse operation binding the contract event 0x4c04f413243ecfd9c1ddddf799c525affe9c8fe07aa9fa7ce49bc953ad89e258.
//
// Solidity: event FulfillEvent(bytes sig, uint32 orderId, (address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
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

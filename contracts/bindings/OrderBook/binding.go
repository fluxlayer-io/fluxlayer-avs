// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOrderBook

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

// ContractOrderBookMetaData contains all meta data concerning the ContractOrderBook contract.
var ContractOrderBookMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"signChainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderDetails\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allOrderResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIOrderBook.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIOrderBook.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToFulfill\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"createdBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderResponse\",\"type\":\"tuple\",\"internalType\":\"structIOrderBook.OrderResponse\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIOrderBook.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FulfillEvent\",\"inputs\":[{\"name\":\"sig\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"order\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOrderBook.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderRespondedEvent\",\"inputs\":[{\"name\":\"orderResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOrderBook.OrderResponse\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"OrderResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOrderBook.OrderResponseMetadata\",\"components\":[{\"name\":\"orderResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderFulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TakerMismatch\",\"inputs\":[]}]",
	Bin: "0x6101806040527fcd190519e6feb4299db38636bfbcef24f017d944ee1188f74093454f21ed61f1610140523480156200003757600080fd5b5060405162005dff38038062005dff8339810160408190526200005a91620002ae565b604051806040016040528060098152602001684f72646572426f6f6b60b81b815250604051806040016040528060038152602001620312e360ec1b81525082600085806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200011b9190620002f5565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000173573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001999190620002f5565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001f3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002199190620002f5565b6001600160a01b0390811660e0526097805460ff19166001179055855160208088019190912086519187019190912061010082905261012081905263ffffffff86166101605290925090831615620002875760c980546001600160a01b0319166001600160a01b0385161790555b50505050505050506200031c565b6001600160a01b0381168114620002ab57600080fd5b50565b60008060408385031215620002c257600080fd5b8251620002cf8162000295565b602084015190925063ffffffff81168114620002ea57600080fd5b809150509250929050565b6000602082840312156200030857600080fd5b8151620003158162000295565b9392505050565b60805160a05160c05160e05161010051610120516101405161016051615a45620003ba600039600061415d015260006111740152600061413c0152600061411b015260008181610571015261246701526000818161036e01526126490152600081816103950152818161281f01526129e10152600081816103bc0152818161125701528181612132015281816122ca01526125040152615a456000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c8063683048351161010f578063c0c53b8b116100a2578063df5cf72311610071578063df5cf7231461056c578063e772935f14610593578063f2fde38b146105b3578063fabc1cbc146105c657600080fd5b8063c0c53b8b14610505578063c33e1fd314610518578063cefdc1d41461052b578063dd3285b31461054c57600080fd5b8063886f1195116100de578063886f1195146104075780638da5cb5b1461041a578063aa4f5cc61461042b578063b98d0908146104f857600080fd5b806368304835146103905780636d14a987146103b75780636efb4636146103de578063715018a6146103ff57600080fd5b80633b692c09116101875780635ac86ab7116101565780635ac86ab71461031e5780635c155662146103415780635c975abb146103615780635df459461461036957600080fd5b80633b692c09146102c2578063416c7e5e146102e35780634f739f74146102f6578063595c6a671461031657600080fd5b8063245a7bfc116101c3578063245a7bfc14610241578063298fb5691461026c5780633563b0d11461028f578063377a91a6146102af57600080fd5b806310d67a2f146101ea578063136439dd146101ff578063171f1d5b14610212575b600080fd5b6101fd6101f8366004614648565b6105d9565b005b6101fd61020d366004614665565b610695565b610225610220366004614806565b6107d4565b6040805192151583529015156020830152015b60405180910390f35b60cd54610254906001600160a01b031681565b6040516001600160a01b039091168152602001610238565b61027f61027a366004614962565b61095e565b6040519015158152602001610238565b6102a261029d3660046149b7565b6109c9565b6040516102389190614b12565b6101fd6102bd366004614dc9565b610e5f565b6102d56102d0366004614e77565b61116d565b604051908152602001610238565b6101fd6102f1366004614ea2565b611255565b610309610304366004614ebf565b6113ca565b6040516102389190614fc3565b6101fd611af0565b61027f61032c36600461508d565b606654600160ff9092169190911b9081161490565b61035461034f3660046150aa565b611bb7565b6040516102389190615156565b6066546102d5565b6102547f000000000000000000000000000000000000000000000000000000000000000081565b6102547f000000000000000000000000000000000000000000000000000000000000000081565b6102547f000000000000000000000000000000000000000000000000000000000000000081565b6103f16103ec36600461519a565b611d7f565b60405161023892919061525a565b6101fd612c96565b606554610254906001600160a01b031681565b6033546001600160a01b0316610254565b61049b6104393660046152a3565b60cb602052600090815260409020805460018201546002830154600384015460048501546005860154600687015460079097015463ffffffff808816986401000000009098046001600160a01b03908116989781169796811696941693911689565b6040805163ffffffff9a8b1681526001600160a01b03998a1660208201529789169088015294871660608701526080860193909352941660a084015260c083019390935260e0820192909252911661010082015261012001610238565b60975461027f9060ff1681565b6101fd6105133660046152c0565b612caa565b6101fd610526366004614962565b612de5565b61053e61053936600461530b565b613047565b604051610238929190615342565b6102d561055a3660046152a3565b60ca6020526000908152604090205481565b6102547f000000000000000000000000000000000000000000000000000000000000000081565b6102d56105a13660046152a3565b60cc6020526000908152604090205481565b6101fd6105c1366004614648565b6131d9565b6101fd6105d4366004614665565b61324f565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561062c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106509190615363565b6001600160a01b0316336001600160a01b0316146106895760405162461bcd60e51b815260040161068090615380565b60405180910390fd5b610692816133ab565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106dd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070191906153ca565b61071d5760405162461bcd60e51b8152600401610680906153e7565b606654818116146107965760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610680565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061081c5761081c61542f565b60200201518951600160200201518a602001516000600281106108415761084161542f565b60200201518b6020015160016002811061085d5761085d61542f565b602090810291909101518c518d8301516040516108ba9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108dd9190615445565b90506109506108f66108ef88846134a2565b8690613539565b6108fe6135cd565b61094661093785610931604080518082018252600080825260209182015281518083019092526001825260029082015290565b906134a2565b6109408c61368d565b90613539565b886201d4c061371d565b909890975095505050505050565b6000806109ac84848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506109a6925089915061116d9050565b90613941565b60208601516001600160a01b039182169116149150509392505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a2f9190615363565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a959190615363565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ad7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610afb9190615363565b9050600086516001600160401b03811115610b1857610b1861467e565b604051908082528060200260200182016040528015610b4b57816020015b6060815260200190600190039081610b365790505b50905060005b8751811015610e53576000888281518110610b6e57610b6e61542f565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610bcf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610bf79190810190615467565b905080516001600160401b03811115610c1257610c1261467e565b604051908082528060200260200182016040528015610c5d57816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610c305790505b50848481518110610c7057610c7061542f565b602002602001018190525060005b8151811015610e3d576040518060600160405280876001600160a01b03166347b314e8858581518110610cb357610cb361542f565b60200260200101516040518263ffffffff1660e01b8152600401610cd991815260200190565b602060405180830381865afa158015610cf6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1a9190615363565b6001600160a01b03168152602001838381518110610d3a57610d3a61542f565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610d6857610d6861542f565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610dc4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de891906154f7565b6001600160601b0316815250858581518110610e0657610e0661542f565b60200260200101518281518110610e1f57610e1f61542f565b60200260200101819052508080610e3590615536565b915050610c7e565b5050508080610e4b90615536565b915050610b51565b50979650505050505050565b60cd546001600160a01b03163314610eb95760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610680565b600060cc81610ece60408601602087016152a3565b63ffffffff1663ffffffff1681526020019081526020016000205414610f4c5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526c103a37903a34329037b93232b960991b6064820152608401610680565b600082604051602001610f5f9190615586565b604051602081830303815290604052805190602001209050600080610f87838a8a8988611d7f565b91509150600060405180604001604052804363ffffffff1681526020018381525090508581604051602001610fbd929190615594565b6040516020818303038152906040528051906020012060cc6000886020016020810190610fea91906152a3565b63ffffffff1663ffffffff16815260200190815260200160002081905550600060cb600088602001602081019061102191906152a3565b63ffffffff908116825260208083019390935260409182016000208251610120810184528154808416825264010000000090046001600160a01b039081169582019590955260018201548516818501819052600283015486166060830181905260038401546080840181905260048086015490981660a0850152600585015460c0850152600685015460e0850152600790940154909416610100830152935163a9059cbb60e01b81529485019390935260248401529092509063a9059cbb906044016020604051808303816000875af1158015611102573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061112691906153ca565b507fa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd8783604051611158929190615594565b60405180910390a15050505050505050505050565b600061124f7f0000000000000000000000000000000000000000000000000000000000000000836000015184602001518560400151866060015187608001518860a001518960c001518a60e001518b61010001516040516020016112349a99989796959493929190998a5263ffffffff98891660208b01526001600160a01b0397881660408b015295871660608a0152938616608089015260a088019290925290931660c086015260e0850192909252610100840191909152166101208201526101400190565b60405160208183030381529060405280519060200120613965565b92915050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112d79190615363565b6001600160a01b0316336001600160a01b0316146113835760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610680565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b6113f56040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611435573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114599190615363565b90506114866040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906114b6908b90899089906004016155c0565b600060405180830381865afa1580156114d3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114fb919081019061560a565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061152d908b908b908b906004016156c1565b600060405180830381865afa15801561154a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611572919081019061560a565b6040820152856001600160401b0381111561158f5761158f61467e565b6040519080825280602002602001820160405280156115c257816020015b60608152602001906001900390816115ad5790505b50606082015260005b60ff8116871115611a01576000856001600160401b038111156115f0576115f061467e565b604051908082528060200260200182016040528015611619578160200160208202803683370190505b5083606001518360ff16815181106116335761163361542f565b602002602001018190525060005b868110156119015760008c6001600160a01b03166304ec63518a8a8581811061166c5761166c61542f565b905060200201358e8860000151868151811061168a5761168a61542f565b60200260200101516040518463ffffffff1660e01b81526004016116c79392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156116e4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061170891906156ea565b90506001600160c01b0381166117ac5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610680565b8a8a8560ff168181106117c1576117c161542f565b6001600160c01b03841692013560f81c9190911c6001908116141590506118ee57856001600160a01b031663dd9846b98a8a858181106118035761180361542f565b905060200201358d8d8860ff1681811061181f5761181f61542f565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611875573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118999190615713565b85606001518560ff16815181106118b2576118b261542f565b602002602001015184815181106118cb576118cb61542f565b63ffffffff90921660209283029190910190910152826118ea81615536565b9350505b50806118f981615536565b915050611641565b506000816001600160401b0381111561191c5761191c61467e565b604051908082528060200260200182016040528015611945578160200160208202803683370190505b50905060005b828110156119c65784606001518460ff168151811061196c5761196c61542f565b602002602001015181815181106119855761198561542f565b602002602001015182828151811061199f5761199f61542f565b63ffffffff90921660209283029190910190910152806119be81615536565b91505061194b565b508084606001518460ff16815181106119e1576119e161542f565b6020026020010181905250505080806119f990615730565b9150506115cb565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a669190615363565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611a99908b908b908e90600401615750565b600060405180830381865afa158015611ab6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ade919081019061560a565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611b38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b5c91906153ca565b611b785760405162461bcd60e51b8152600401610680906153e7565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611be992919061577a565b600060405180830381865afa158015611c06573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c2e919081019061560a565b9050600084516001600160401b03811115611c4b57611c4b61467e565b604051908082528060200260200182016040528015611c74578160200160208202803683370190505b50905060005b8551811015611d7557866001600160a01b03166304ec6351878381518110611ca457611ca461542f565b602002602001015187868581518110611cbf57611cbf61542f565b60200260200101516040518463ffffffff1660e01b8152600401611cfc9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611d19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d3d91906156ea565b6001600160c01b0316828281518110611d5857611d5861542f565b602090810291909101015280611d6d81615536565b915050611c7a565b5095945050505050565b6040805180820190915260608082526020820152600084611df65760405162461bcd60e51b815260206004820152603760248201526000805160206159f083398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610680565b60408301515185148015611e0e575060a08301515185145b8015611e1e575060c08301515185145b8015611e2e575060e08301515185145b611e985760405162461bcd60e51b815260206004820152604160248201526000805160206159f083398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610680565b82515160208401515114611f105760405162461bcd60e51b8152602060048201526044602482018190526000805160206159f0833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610680565b4363ffffffff168463ffffffff1610611f7f5760405162461bcd60e51b815260206004820152603c60248201526000805160206159f083398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610680565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611fc057611fc061467e565b604051908082528060200260200182016040528015611fe9578160200160208202803683370190505b506020820152866001600160401b038111156120075761200761467e565b604051908082528060200260200182016040528015612030578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156120645761206461467e565b60405190808252806020026020018201604052801561208d578160200160208202803683370190505b5081526020860151516001600160401b038111156120ad576120ad61467e565b6040519080825280602002602001820160405280156120d6578160200160208202803683370190505b50816020018190525060006121a88a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa15801561217f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121a391906157ce565b6139ac565b905060005b876020015151811015612443576121f2886020015182815181106121d3576121d361542f565b6020026020010151805160009081526020918201519091526040902090565b836020015182815181106122085761220861542f565b602090810291909101015280156122c85760208301516122296001836157eb565b815181106122395761223961542f565b602002602001015160001c8360200151828151811061225a5761225a61542f565b602002602001015160001c116122c8576040805162461bcd60e51b81526020600482015260248101919091526000805160206159f083398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610680565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061230d5761230d61542f565b60200260200101518b8b60000151858151811061232c5761232c61542f565b60200260200101516040518463ffffffff1660e01b81526004016123699392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612386573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123aa91906156ea565b6001600160c01b0316836000015182815181106123c9576123c961542f565b60200260200101818152505061242f6108ef61240384866000015185815181106123f5576123f561542f565b602002602001015116613a3d565b8a6020015184815181106124195761241961542f565b6020026020010151613a6890919063ffffffff16565b94508061243b81615536565b9150506121ad565b505061244e83613b4c565b60975490935060ff166000816124655760006124e7565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156124c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124e79190615802565b905060005b8a811015612b65578215612647578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106125435761254361542f565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612583573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125a79190615802565b6125b1919061581b565b116126475760405162461bcd60e51b815260206004820152606660248201526000805160206159f083398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610680565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106126885761268861542f565b9050013560f81c60f81b60f81c8c8c60a0015185815181106126ac576126ac61542f565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612708573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061272c9190615833565b6001600160401b03191661274f8a6040015183815181106121d3576121d361542f565b67ffffffffffffffff1916146127eb5760405162461bcd60e51b815260206004820152606160248201526000805160206159f083398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610680565b61281b896040015182815181106128045761280461542f565b60200260200101518761353990919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061285e5761285e61542f565b9050013560f81c60f81b60f81c8c8c60c0015185815181106128825761288261542f565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156128de573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061290291906154f7565b856020015182815181106129185761291861542f565b6001600160601b039092166020928302919091018201528501518051829081106129445761294461542f565b6020026020010151856000015182815181106129625761296261542f565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612b50576129da866000015182815181106129ac576129ac61542f565b60200260200101518f8f868181106129c6576129c661542f565b600192013560f81c9290921c811614919050565b15612b3e577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612a2057612a2061542f565b9050013560f81c60f81b60f81c8e89602001518581518110612a4457612a4461542f565b60200260200101518f60e001518881518110612a6257612a6261542f565b60200260200101518781518110612a7b57612a7b61542f565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612adf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b0391906154f7565b8751805185908110612b1757612b1761542f565b60200260200101818151612b2b919061585e565b6001600160601b03169052506001909101905b80612b4881615536565b915050612986565b50508080612b5d90615536565b9150506124ec565b505050600080612b7f8c868a606001518b608001516107d4565b9150915081612bf05760405162461bcd60e51b815260206004820152604360248201526000805160206159f083398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610680565b80612c515760405162461bcd60e51b815260206004820152603960248201526000805160206159f083398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610680565b50506000878260200151604051602001612c6c929190615886565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612c9e613be7565b612ca86000613c41565b565b600054610100900460ff1615808015612cca5750600054600160ff909116105b80612ce45750303b158015612ce4575060005460ff166001145b612d475760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610680565b6000805460ff191660011790558015612d6a576000805461ff0019166101001790555b612d75846000613c93565b612d7e83613c41565b60cd80546001600160a01b0319166001600160a01b0384161790558015612ddf576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b825163ffffffff8116600090815260ca602052604090205415612e1b57604051630e59146360e41b815260040160405180910390fd5b60408401516001600160a01b031615801590612e44575060408401516001600160a01b03163314155b15612e6257604051630d9d754f60e41b815260040160405180910390fd5b428460e001511015612e87576040516362b439dd60e11b815260040160405180910390fd5b612e9284848461095e565b612eaf57604051638baa579f60e01b815260040160405180910390fd5b6060840151602085015160808601516040516323b872dd60e01b81526001600160a01b03928316600482015230602482015260448101919091529116906323b872dd906064016020604051808303816000875af1158015612f14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f3891906153ca565b503360408086019190915251612f529085906020016158ce565b60408051808303601f19018152918152815160209283012063ffffffff938416600090815260ca84528281209190915560cb835281902086518154938801519085166001600160c01b0319909416939093176401000000006001600160a01b0394851602178155908601516001820180546001600160a01b0319908116928516929092179055606087015160028301805483169185169190911790556080870151600383015560a087015160048301805490921693169290921790915560c0850151600582015560e08501516006820155610100909401516007909401805463ffffffff191694909116939093179092555050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106130825761308261542f565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906130be908890869060040161577a565b600060405180830381865afa1580156130db573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613103919081019061560a565b6000815181106131155761311561542f565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015613181573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131a591906156ea565b6001600160c01b0316905060006131bb82613d7d565b9050816131c98a838a6109c9565b9550955050505050935093915050565b6131e1613be7565b6001600160a01b0381166132465760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610680565b61069281613c41565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156132a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132c69190615363565b6001600160a01b0316336001600160a01b0316146132f65760405162461bcd60e51b815260040161068090615380565b6066541981196066541916146133745760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610680565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107c9565b6001600160a01b0381166134395760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610680565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526134be614559565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156134f1576134f3565bfe5b50806135315760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610680565b505092915050565b6040805180820190915260008082526020820152613555614577565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156134f15750806135315760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610680565b6135d5614595565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806136bd6000805160206159d083398151915286615445565b90505b6136c981613e49565b90935091506000805160206159d0833981519152828309831415613703576040805180820190915290815260208101919091529392505050565b6000805160206159d08339815191526001820890506136c0565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061374f6145ba565b60005b6002811015613914576000613768826006615978565b905084826002811061377c5761377c61542f565b6020020151518361378e83600061581b565b600c811061379e5761379e61542f565b60200201528482600281106137b5576137b561542f565b602002015160200151838260016137cc919061581b565b600c81106137dc576137dc61542f565b60200201528382600281106137f3576137f361542f565b602002015151518361380683600261581b565b600c81106138165761381661542f565b602002015283826002811061382d5761382d61542f565b602002015151600160200201518361384683600361581b565b600c81106138565761385661542f565b602002015283826002811061386d5761386d61542f565b6020020151602001516000600281106138885761388861542f565b60200201518361389983600461581b565b600c81106138a9576138a961542f565b60200201528382600281106138c0576138c061542f565b6020020151602001516001600281106138db576138db61542f565b6020020151836138ec83600561581b565b600c81106138fc576138fc61542f565b6020020152508061390c81615536565b915050613752565b5061391d6145d9565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60008060006139508585613ecb565b9150915061395d81613f3b565b509392505050565b600061396f6140f6565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b6000806139b8846141d5565b9050808360ff166001901b11613a365760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610680565b9392505050565b6000805b821561124f57613a526001846157eb565b9092169180613a6081615997565b915050613a41565b60408051808201909152600080825260208201526102008261ffff1610613ac45760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610680565b8161ffff1660011415613ad857508161124f565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613b4157600161ffff871660ff83161c81161415613b2457613b218484613539565b93505b613b2e8384613539565b92506201fffe600192831b169101613af4565b509195945050505050565b60408051808201909152600080825260208201528151158015613b7157506020820151155b15613b8f575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206159d08339815191528460200151613bc29190615445565b613bda906000805160206159d08339815191526157eb565b905292915050565b919050565b6033546001600160a01b03163314612ca85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610680565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546001600160a01b0316158015613cb457506001600160a01b03821615155b613d365760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610680565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613d79826133ab565b5050565b6060600080613d8b84613a3d565b61ffff166001600160401b03811115613da657613da661467e565b6040519080825280601f01601f191660200182016040528015613dd0576020820181803683370190505b5090506000805b825182108015613de8575061010081105b15613e3f576001811b935085841615613e2f578060f81b838381518110613e1157613e1161542f565b60200101906001600160f81b031916908160001a9053508160010191505b613e3881615536565b9050613dd7565b5090949350505050565b600080806000805160206159d083398151915260036000805160206159d0833981519152866000805160206159d0833981519152888909090890506000613ebf827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206159d0833981519152614362565b91959194509092505050565b600080825160411415613f025760208301516040840151606085015160001a613ef68782858561440a565b94509450505050613f34565b825160401415613f2c5760208301516040840151613f218683836144f7565b935093505050613f34565b506000905060025b9250929050565b6000816004811115613f4f57613f4f6159b9565b1415613f585750565b6001816004811115613f6c57613f6c6159b9565b1415613fba5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610680565b6002816004811115613fce57613fce6159b9565b141561401c5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610680565b6003816004811115614030576140306159b9565b14156140895760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610680565b600481600481111561409d5761409d6159b9565b14156106925760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610680565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000614184614530565b604080516020810196909652850193909352606084019190915263ffffffff1660808301526001600160a01b031660a082015260c00160405160208183030381529060405280519060200120905090565b60006101008251111561425e5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610680565b815161426c57506000919050565b600080836000815181106142825761428261542f565b0160200151600160f89190911c81901b92505b8451811015614359578481815181106142b0576142b061542f565b0160200151600160f89190911c1b91508282116143455760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610680565b9181179161435281615536565b9050614295565b50909392505050565b60008061436d6145d9565b6143756145f7565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156134f15750826143ff5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610680565b505195945050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561444157506000905060036144ee565b8460ff16601b1415801561445957508460ff16601c14155b1561446a57506000905060046144ee565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156144be573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166144e7576000600192509250506144ee565b9150600090505b94509492505050565b6000806001600160ff1b0383168161451460ff86901c601b61581b565b90506145228782888561440a565b935093505050935093915050565b60c9546000906001600160a01b031615614554575060c9546001600160a01b031690565b503090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806145a8614615565b81526020016145b5614615565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461069257600080fd5b60006020828403121561465a57600080fd5b8135613a3681614633565b60006020828403121561467757600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156146b6576146b661467e565b60405290565b60405161012081016001600160401b03811182821017156146b6576146b661467e565b60405161010081016001600160401b03811182821017156146b6576146b661467e565b604051601f8201601f191681016001600160401b038111828210171561472a5761472a61467e565b604052919050565b60006040828403121561474457600080fd5b61474c614694565b9050813581526020820135602082015292915050565b600082601f83011261477357600080fd5b604051604081018181106001600160401b03821117156147955761479561467e565b80604052508060408401858111156147ac57600080fd5b845b81811015613b415780358352602092830192016147ae565b6000608082840312156147d857600080fd5b6147e0614694565b90506147ec8383614762565b81526147fb8360408401614762565b602082015292915050565b600080600080610120858703121561481d57600080fd5b8435935061482e8660208701614732565b925061483d86606087016147c6565b915061484c8660e08701614732565b905092959194509250565b63ffffffff8116811461069257600080fd5b8035613be281614857565b8035613be281614633565b6000610120828403121561489257600080fd5b61489a6146bc565b90506148a582614869565b81526148b360208301614874565b60208201526148c460408301614874565b60408201526148d560608301614874565b6060820152608082013560808201526148f060a08301614874565b60a082015260c082013560c082015260e082013560e0820152610100614917818401614869565b9082015292915050565b60008083601f84011261493357600080fd5b5081356001600160401b0381111561494a57600080fd5b602083019150836020828501011115613f3457600080fd5b6000806000610140848603121561497857600080fd5b614982858561487f565b92506101208401356001600160401b0381111561499e57600080fd5b6149aa86828701614921565b9497909650939450505050565b6000806000606084860312156149cc57600080fd5b83356149d781614633565b92506020848101356001600160401b03808211156149f457600080fd5b818701915087601f830112614a0857600080fd5b813581811115614a1a57614a1a61467e565b614a2c601f8201601f19168501614702565b91508082528884828501011115614a4257600080fd5b8084840185840137600084828401015250809450505050614a6560408501614869565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614b04578385038a52825180518087529087019087870190845b81811015614aef57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614aab565b50509a87019a95505091850191600101614a8d565b509298975050505050505050565b602081526000613a366020830184614a6e565b60006001600160401b03821115614b3e57614b3e61467e565b5060051b60200190565b600082601f830112614b5957600080fd5b81356020614b6e614b6983614b25565b614702565b82815260059290921b84018101918181019086841115614b8d57600080fd5b8286015b84811015614bb1578035614ba481614857565b8352918301918301614b91565b509695505050505050565b600082601f830112614bcd57600080fd5b81356020614bdd614b6983614b25565b82815260069290921b84018101918181019086841115614bfc57600080fd5b8286015b84811015614bb157614c128882614732565b835291830191604001614c00565b600082601f830112614c3157600080fd5b81356020614c41614b6983614b25565b82815260059290921b84018101918181019086841115614c6057600080fd5b8286015b84811015614bb15780356001600160401b03811115614c835760008081fd5b614c918986838b0101614b48565b845250918301918301614c64565b60006101808284031215614cb257600080fd5b614cba6146df565b905081356001600160401b0380821115614cd357600080fd5b614cdf85838601614b48565b83526020840135915080821115614cf557600080fd5b614d0185838601614bbc565b60208401526040840135915080821115614d1a57600080fd5b614d2685838601614bbc565b6040840152614d3885606086016147c6565b6060840152614d4a8560e08601614732565b6080840152610120840135915080821115614d6457600080fd5b614d7085838601614b48565b60a0840152610140840135915080821115614d8a57600080fd5b614d9685838601614b48565b60c0840152610160840135915080821115614db057600080fd5b50614dbd84828501614c20565b60e08301525092915050565b60008060008060008086880360c0811215614de357600080fd5b87356001600160401b0380821115614dfa57600080fd5b614e068b838c01614921565b909950975060208a01359150614e1b82614857565b909550604089013590614e2d82614857565b8195506040605f1984011215614e4257600080fd5b60608a01945060a08a0135925080831115614e5c57600080fd5b5050614e6a89828a01614c9f565b9150509295509295509295565b60006101208284031215614e8a57600080fd5b613a36838361487f565b801515811461069257600080fd5b600060208284031215614eb457600080fd5b8135613a3681614e94565b60008060008060008060808789031215614ed857600080fd5b8635614ee381614633565b95506020870135614ef381614857565b945060408701356001600160401b0380821115614f0f57600080fd5b614f1b8a838b01614921565b90965094506060890135915080821115614f3457600080fd5b818901915089601f830112614f4857600080fd5b813581811115614f5757600080fd5b8a60208260051b8501011115614f6c57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614fb857815163ffffffff1687529582019590820190600101614f96565b509495945050505050565b600060208083528351608082850152614fdf60a0850182614f82565b905081850151601f1980868403016040870152614ffc8383614f82565b925060408701519150808684030160608701526150198383614f82565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015615070578487830301845261505e828751614f82565b95880195938801939150600101615044565b509998505050505050505050565b60ff8116811461069257600080fd5b60006020828403121561509f57600080fd5b8135613a368161507e565b6000806000606084860312156150bf57600080fd5b83356150ca81614633565b92506020848101356001600160401b038111156150e657600080fd5b8501601f810187136150f757600080fd5b8035615105614b6982614b25565b81815260059190911b8201830190838101908983111561512457600080fd5b928401925b8284101561514257833582529284019290840190615129565b8096505050505050614a6560408501614869565b6020808252825182820181905260009190848201906040850190845b8181101561518e57835183529284019291840191600101615172565b50909695505050505050565b6000806000806000608086880312156151b257600080fd5b8535945060208601356001600160401b03808211156151d057600080fd5b6151dc89838a01614921565b9096509450604088013591506151f182614857565b9092506060870135908082111561520757600080fd5b5061521488828901614c9f565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614fb85781516001600160601b031687529582019590820190600101615235565b60408152600083516040808401526152756080840182615221565b90506020850151603f198483030160608501526152928282615221565b925050508260208301529392505050565b6000602082840312156152b557600080fd5b8135613a3681614857565b6000806000606084860312156152d557600080fd5b83356152e081614633565b925060208401356152f081614633565b9150604084013561530081614633565b809150509250925092565b60008060006060848603121561532057600080fd5b833561532b81614633565b925060208401359150604084013561530081614857565b82815260406020820152600061535b6040830184614a6e565b949350505050565b60006020828403121561537557600080fd5b8151613a3681614633565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156153dc57600080fd5b8151613a3681614e94565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261546257634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561547a57600080fd5b82516001600160401b0381111561549057600080fd5b8301601f810185136154a157600080fd5b80516154af614b6982614b25565b81815260059190911b820183019083810190878311156154ce57600080fd5b928401925b828410156154ec578351825292840192908401906154d3565b979650505050505050565b60006020828403121561550957600080fd5b81516001600160601b0381168114613a3657600080fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561554a5761554a615520565b5060010190565b803561555c81614633565b6001600160a01b03168252602081013561557581614857565b63ffffffff81166020840152505050565b6040810161124f8284615551565b608081016155a28285615551565b63ffffffff8351166040830152602083015160608301529392505050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156155ed57600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561561d57600080fd5b82516001600160401b0381111561563357600080fd5b8301601f8101851361564457600080fd5b8051615652614b6982614b25565b81815260059190911b8201830190838101908783111561567157600080fd5b928401925b828410156154ec57835161568981614857565b82529284019290840190615676565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006156e1604083018486615698565b95945050505050565b6000602082840312156156fc57600080fd5b81516001600160c01b0381168114613a3657600080fd5b60006020828403121561572557600080fd5b8151613a3681614857565b600060ff821660ff81141561574757615747615520565b60010192915050565b604081526000615764604083018587615698565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156157c1578451835293830193918301916001016157a5565b5090979650505050505050565b6000602082840312156157e057600080fd5b8151613a368161507e565b6000828210156157fd576157fd615520565b500390565b60006020828403121561581457600080fd5b5051919050565b6000821982111561582e5761582e615520565b500190565b60006020828403121561584557600080fd5b815167ffffffffffffffff1981168114613a3657600080fd5b60006001600160601b038381169083168181101561587e5761587e615520565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156158c1578151855293820193908201906001016158a5565b5092979650505050505050565b815163ffffffff1681526020808301516001600160a01b03169082015260408083015161012083019161590b908401826001600160a01b03169052565b50606083015161592660608401826001600160a01b03169052565b506080830151608083015260a083015161594b60a08401826001600160a01b03169052565b5060c083015160c083015260e083015160e0830152610100808401516135318285018263ffffffff169052565b600081600019048311821515161561599257615992615520565b500290565b600061ffff808316818114156159af576159af615520565b6001019392505050565b634e487b7160e01b600052602160045260246000fdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122023035a3e84bdab087d84f78ca9b786027a1a487f83e8fc9e099315fb435eb64f64736f6c634300080c0033",
}

// ContractOrderBookABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOrderBookMetaData.ABI instead.
var ContractOrderBookABI = ContractOrderBookMetaData.ABI

// ContractOrderBookBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOrderBookMetaData.Bin instead.
var ContractOrderBookBin = ContractOrderBookMetaData.Bin

// DeployContractOrderBook deploys a new Ethereum contract, binding an instance of ContractOrderBook to it.
func DeployContractOrderBook(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, signChainId uint32) (common.Address, *types.Transaction, *ContractOrderBook, error) {
	parsed, err := ContractOrderBookMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOrderBookBin), backend, _registryCoordinator, signChainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOrderBook{ContractOrderBookCaller: ContractOrderBookCaller{contract: contract}, ContractOrderBookTransactor: ContractOrderBookTransactor{contract: contract}, ContractOrderBookFilterer: ContractOrderBookFilterer{contract: contract}}, nil
}

// ContractOrderBook is an auto generated Go binding around an Ethereum contract.
type ContractOrderBook struct {
	ContractOrderBookCaller     // Read-only binding to the contract
	ContractOrderBookTransactor // Write-only binding to the contract
	ContractOrderBookFilterer   // Log filterer for contract events
}

// ContractOrderBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOrderBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOrderBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOrderBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOrderBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOrderBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOrderBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOrderBookSession struct {
	Contract     *ContractOrderBook // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractOrderBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOrderBookCallerSession struct {
	Contract *ContractOrderBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractOrderBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOrderBookTransactorSession struct {
	Contract     *ContractOrderBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractOrderBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOrderBookRaw struct {
	Contract *ContractOrderBook // Generic contract binding to access the raw methods on
}

// ContractOrderBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOrderBookCallerRaw struct {
	Contract *ContractOrderBookCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOrderBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOrderBookTransactorRaw struct {
	Contract *ContractOrderBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOrderBook creates a new instance of ContractOrderBook, bound to a specific deployed contract.
func NewContractOrderBook(address common.Address, backend bind.ContractBackend) (*ContractOrderBook, error) {
	contract, err := bindContractOrderBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOrderBook{ContractOrderBookCaller: ContractOrderBookCaller{contract: contract}, ContractOrderBookTransactor: ContractOrderBookTransactor{contract: contract}, ContractOrderBookFilterer: ContractOrderBookFilterer{contract: contract}}, nil
}

// NewContractOrderBookCaller creates a new read-only instance of ContractOrderBook, bound to a specific deployed contract.
func NewContractOrderBookCaller(address common.Address, caller bind.ContractCaller) (*ContractOrderBookCaller, error) {
	contract, err := bindContractOrderBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookCaller{contract: contract}, nil
}

// NewContractOrderBookTransactor creates a new write-only instance of ContractOrderBook, bound to a specific deployed contract.
func NewContractOrderBookTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOrderBookTransactor, error) {
	contract, err := bindContractOrderBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookTransactor{contract: contract}, nil
}

// NewContractOrderBookFilterer creates a new log filterer instance of ContractOrderBook, bound to a specific deployed contract.
func NewContractOrderBookFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOrderBookFilterer, error) {
	contract, err := bindContractOrderBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookFilterer{contract: contract}, nil
}

// bindContractOrderBook binds a generic wrapper to an already deployed contract.
func bindContractOrderBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOrderBookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOrderBook *ContractOrderBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOrderBook.Contract.ContractOrderBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOrderBook *ContractOrderBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.ContractOrderBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOrderBook *ContractOrderBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.ContractOrderBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOrderBook *ContractOrderBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOrderBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOrderBook *ContractOrderBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOrderBook *ContractOrderBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOrderBook *ContractOrderBookCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOrderBook *ContractOrderBookSession) Aggregator() (common.Address, error) {
	return _ContractOrderBook.Contract.Aggregator(&_ContractOrderBook.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOrderBook *ContractOrderBookCallerSession) Aggregator() (common.Address, error) {
	return _ContractOrderBook.Contract.Aggregator(&_ContractOrderBook.CallOpts)
}

// AllOrderDetails is a free data retrieval call binding the contract method 0xaa4f5cc6.
//
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint256 expiry, uint32 targetNetworkNumber)
func (_ContractOrderBook *ContractOrderBookCaller) AllOrderDetails(opts *bind.CallOpts, arg0 uint32) (struct {
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
	err := _ContractOrderBook.contract.Call(opts, &out, "allOrderDetails", arg0)

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
func (_ContractOrderBook *ContractOrderBookSession) AllOrderDetails(arg0 uint32) (struct {
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
	return _ContractOrderBook.Contract.AllOrderDetails(&_ContractOrderBook.CallOpts, arg0)
}

// AllOrderDetails is a free data retrieval call binding the contract method 0xaa4f5cc6.
//
// Solidity: function allOrderDetails(uint32 ) view returns(uint32 orderId, address maker, address taker, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount, uint256 expiry, uint32 targetNetworkNumber)
func (_ContractOrderBook *ContractOrderBookCallerSession) AllOrderDetails(arg0 uint32) (struct {
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
	return _ContractOrderBook.Contract.AllOrderDetails(&_ContractOrderBook.CallOpts, arg0)
}

// AllOrderHashes is a free data retrieval call binding the contract method 0xdd3285b3.
//
// Solidity: function allOrderHashes(uint32 ) view returns(bytes32)
func (_ContractOrderBook *ContractOrderBookCaller) AllOrderHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "allOrderHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllOrderHashes is a free data retrieval call binding the contract method 0xdd3285b3.
//
// Solidity: function allOrderHashes(uint32 ) view returns(bytes32)
func (_ContractOrderBook *ContractOrderBookSession) AllOrderHashes(arg0 uint32) ([32]byte, error) {
	return _ContractOrderBook.Contract.AllOrderHashes(&_ContractOrderBook.CallOpts, arg0)
}

// AllOrderHashes is a free data retrieval call binding the contract method 0xdd3285b3.
//
// Solidity: function allOrderHashes(uint32 ) view returns(bytes32)
func (_ContractOrderBook *ContractOrderBookCallerSession) AllOrderHashes(arg0 uint32) ([32]byte, error) {
	return _ContractOrderBook.Contract.AllOrderHashes(&_ContractOrderBook.CallOpts, arg0)
}

// AllOrderResponses is a free data retrieval call binding the contract method 0xe772935f.
//
// Solidity: function allOrderResponses(uint32 ) view returns(bytes32)
func (_ContractOrderBook *ContractOrderBookCaller) AllOrderResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "allOrderResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllOrderResponses is a free data retrieval call binding the contract method 0xe772935f.
//
// Solidity: function allOrderResponses(uint32 ) view returns(bytes32)
func (_ContractOrderBook *ContractOrderBookSession) AllOrderResponses(arg0 uint32) ([32]byte, error) {
	return _ContractOrderBook.Contract.AllOrderResponses(&_ContractOrderBook.CallOpts, arg0)
}

// AllOrderResponses is a free data retrieval call binding the contract method 0xe772935f.
//
// Solidity: function allOrderResponses(uint32 ) view returns(bytes32)
func (_ContractOrderBook *ContractOrderBookCallerSession) AllOrderResponses(arg0 uint32) ([32]byte, error) {
	return _ContractOrderBook.Contract.AllOrderResponses(&_ContractOrderBook.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOrderBook *ContractOrderBookCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOrderBook *ContractOrderBookSession) BlsApkRegistry() (common.Address, error) {
	return _ContractOrderBook.Contract.BlsApkRegistry(&_ContractOrderBook.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOrderBook *ContractOrderBookCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractOrderBook.Contract.BlsApkRegistry(&_ContractOrderBook.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractOrderBook *ContractOrderBookCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

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
func (_ContractOrderBook *ContractOrderBookSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractOrderBook.Contract.CheckSignatures(&_ContractOrderBook.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractOrderBook *ContractOrderBookCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractOrderBook.Contract.CheckSignatures(&_ContractOrderBook.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOrderBook *ContractOrderBookCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOrderBook *ContractOrderBookSession) Delegation() (common.Address, error) {
	return _ContractOrderBook.Contract.Delegation(&_ContractOrderBook.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOrderBook *ContractOrderBookCallerSession) Delegation() (common.Address, error) {
	return _ContractOrderBook.Contract.Delegation(&_ContractOrderBook.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOrderBook *ContractOrderBookCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOrderBook *ContractOrderBookSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOrderBook.Contract.GetCheckSignaturesIndices(&_ContractOrderBook.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOrderBook *ContractOrderBookCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOrderBook.Contract.GetCheckSignaturesIndices(&_ContractOrderBook.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOrderBook *ContractOrderBookCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOrderBook *ContractOrderBookSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOrderBook.Contract.GetOperatorState(&_ContractOrderBook.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOrderBook *ContractOrderBookCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOrderBook.Contract.GetOperatorState(&_ContractOrderBook.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOrderBook *ContractOrderBookCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

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
func (_ContractOrderBook *ContractOrderBookSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOrderBook.Contract.GetOperatorState0(&_ContractOrderBook.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOrderBook *ContractOrderBookCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOrderBook.Contract.GetOperatorState0(&_ContractOrderBook.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOrderBook *ContractOrderBookCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOrderBook *ContractOrderBookSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOrderBook.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOrderBook.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOrderBook *ContractOrderBookCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOrderBook.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOrderBook.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// HashOrder is a free data retrieval call binding the contract method 0x3b692c09.
//
// Solidity: function hashOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
func (_ContractOrderBook *ContractOrderBookCaller) HashOrder(opts *bind.CallOpts, order IOrderBookOrder) ([32]byte, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x3b692c09.
//
// Solidity: function hashOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
func (_ContractOrderBook *ContractOrderBookSession) HashOrder(order IOrderBookOrder) ([32]byte, error) {
	return _ContractOrderBook.Contract.HashOrder(&_ContractOrderBook.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x3b692c09.
//
// Solidity: function hashOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order) view returns(bytes32)
func (_ContractOrderBook *ContractOrderBookCallerSession) HashOrder(order IOrderBookOrder) ([32]byte, error) {
	return _ContractOrderBook.Contract.HashOrder(&_ContractOrderBook.CallOpts, order)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOrderBook *ContractOrderBookCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOrderBook *ContractOrderBookSession) Owner() (common.Address, error) {
	return _ContractOrderBook.Contract.Owner(&_ContractOrderBook.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOrderBook *ContractOrderBookCallerSession) Owner() (common.Address, error) {
	return _ContractOrderBook.Contract.Owner(&_ContractOrderBook.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOrderBook *ContractOrderBookCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOrderBook *ContractOrderBookSession) Paused(index uint8) (bool, error) {
	return _ContractOrderBook.Contract.Paused(&_ContractOrderBook.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOrderBook *ContractOrderBookCallerSession) Paused(index uint8) (bool, error) {
	return _ContractOrderBook.Contract.Paused(&_ContractOrderBook.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOrderBook *ContractOrderBookCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOrderBook *ContractOrderBookSession) Paused0() (*big.Int, error) {
	return _ContractOrderBook.Contract.Paused0(&_ContractOrderBook.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOrderBook *ContractOrderBookCallerSession) Paused0() (*big.Int, error) {
	return _ContractOrderBook.Contract.Paused0(&_ContractOrderBook.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOrderBook *ContractOrderBookCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOrderBook *ContractOrderBookSession) PauserRegistry() (common.Address, error) {
	return _ContractOrderBook.Contract.PauserRegistry(&_ContractOrderBook.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOrderBook *ContractOrderBookCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractOrderBook.Contract.PauserRegistry(&_ContractOrderBook.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOrderBook *ContractOrderBookCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOrderBook *ContractOrderBookSession) RegistryCoordinator() (common.Address, error) {
	return _ContractOrderBook.Contract.RegistryCoordinator(&_ContractOrderBook.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOrderBook *ContractOrderBookCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractOrderBook.Contract.RegistryCoordinator(&_ContractOrderBook.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOrderBook *ContractOrderBookCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOrderBook *ContractOrderBookSession) StakeRegistry() (common.Address, error) {
	return _ContractOrderBook.Contract.StakeRegistry(&_ContractOrderBook.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOrderBook *ContractOrderBookCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractOrderBook.Contract.StakeRegistry(&_ContractOrderBook.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOrderBook *ContractOrderBookCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOrderBook *ContractOrderBookSession) StaleStakesForbidden() (bool, error) {
	return _ContractOrderBook.Contract.StaleStakesForbidden(&_ContractOrderBook.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOrderBook *ContractOrderBookCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractOrderBook.Contract.StaleStakesForbidden(&_ContractOrderBook.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractOrderBook *ContractOrderBookCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

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
func (_ContractOrderBook *ContractOrderBookSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractOrderBook.Contract.TrySignatureAndApkVerification(&_ContractOrderBook.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractOrderBook *ContractOrderBookCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractOrderBook.Contract.TrySignatureAndApkVerification(&_ContractOrderBook.CallOpts, msgHash, apk, apkG2, sigma)
}

// Verify is a free data retrieval call binding the contract method 0x298fb569.
//
// Solidity: function verify((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractOrderBook *ContractOrderBookCaller) Verify(opts *bind.CallOpts, order IOrderBookOrder, signature []byte) (bool, error) {
	var out []interface{}
	err := _ContractOrderBook.contract.Call(opts, &out, "verify", order, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x298fb569.
//
// Solidity: function verify((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractOrderBook *ContractOrderBookSession) Verify(order IOrderBookOrder, signature []byte) (bool, error) {
	return _ContractOrderBook.Contract.Verify(&_ContractOrderBook.CallOpts, order, signature)
}

// Verify is a free data retrieval call binding the contract method 0x298fb569.
//
// Solidity: function verify((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes signature) view returns(bool)
func (_ContractOrderBook *ContractOrderBookCallerSession) Verify(order IOrderBookOrder, signature []byte) (bool, error) {
	return _ContractOrderBook.Contract.Verify(&_ContractOrderBook.CallOpts, order, signature)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xc33e1fd3.
//
// Solidity: function createOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes sig) returns()
func (_ContractOrderBook *ContractOrderBookTransactor) CreateOrder(opts *bind.TransactOpts, order IOrderBookOrder, sig []byte) (*types.Transaction, error) {
	return _ContractOrderBook.contract.Transact(opts, "createOrder", order, sig)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xc33e1fd3.
//
// Solidity: function createOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes sig) returns()
func (_ContractOrderBook *ContractOrderBookSession) CreateOrder(order IOrderBookOrder, sig []byte) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.CreateOrder(&_ContractOrderBook.TransactOpts, order, sig)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xc33e1fd3.
//
// Solidity: function createOrder((uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, bytes sig) returns()
func (_ContractOrderBook *ContractOrderBookTransactorSession) CreateOrder(order IOrderBookOrder, sig []byte) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.CreateOrder(&_ContractOrderBook.TransactOpts, order, sig)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractOrderBook *ContractOrderBookTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOrderBook.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractOrderBook *ContractOrderBookSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.Initialize(&_ContractOrderBook.TransactOpts, _pauserRegistry, initialOwner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator) returns()
func (_ContractOrderBook *ContractOrderBookTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.Initialize(&_ContractOrderBook.TransactOpts, _pauserRegistry, initialOwner, _aggregator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOrderBook *ContractOrderBookTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOrderBook.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOrderBook *ContractOrderBookSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.Pause(&_ContractOrderBook.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOrderBook *ContractOrderBookTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.Pause(&_ContractOrderBook.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOrderBook *ContractOrderBookTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOrderBook.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOrderBook *ContractOrderBookSession) PauseAll() (*types.Transaction, error) {
	return _ContractOrderBook.Contract.PauseAll(&_ContractOrderBook.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOrderBook *ContractOrderBookTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractOrderBook.Contract.PauseAll(&_ContractOrderBook.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOrderBook *ContractOrderBookTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOrderBook.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOrderBook *ContractOrderBookSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOrderBook.Contract.RenounceOwnership(&_ContractOrderBook.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOrderBook *ContractOrderBookTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOrderBook.Contract.RenounceOwnership(&_ContractOrderBook.TransactOpts)
}

// RespondToFulfill is a paid mutator transaction binding the contract method 0x377a91a6.
//
// Solidity: function respondToFulfill(bytes quorumNumbers, uint32 quorumThresholdPercentage, uint32 createdBlock, (address,uint32) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOrderBook *ContractOrderBookTransactor) RespondToFulfill(opts *bind.TransactOpts, quorumNumbers []byte, quorumThresholdPercentage uint32, createdBlock uint32, orderResponse IOrderBookOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOrderBook.contract.Transact(opts, "respondToFulfill", quorumNumbers, quorumThresholdPercentage, createdBlock, orderResponse, nonSignerStakesAndSignature)
}

// RespondToFulfill is a paid mutator transaction binding the contract method 0x377a91a6.
//
// Solidity: function respondToFulfill(bytes quorumNumbers, uint32 quorumThresholdPercentage, uint32 createdBlock, (address,uint32) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOrderBook *ContractOrderBookSession) RespondToFulfill(quorumNumbers []byte, quorumThresholdPercentage uint32, createdBlock uint32, orderResponse IOrderBookOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.RespondToFulfill(&_ContractOrderBook.TransactOpts, quorumNumbers, quorumThresholdPercentage, createdBlock, orderResponse, nonSignerStakesAndSignature)
}

// RespondToFulfill is a paid mutator transaction binding the contract method 0x377a91a6.
//
// Solidity: function respondToFulfill(bytes quorumNumbers, uint32 quorumThresholdPercentage, uint32 createdBlock, (address,uint32) orderResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOrderBook *ContractOrderBookTransactorSession) RespondToFulfill(quorumNumbers []byte, quorumThresholdPercentage uint32, createdBlock uint32, orderResponse IOrderBookOrderResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.RespondToFulfill(&_ContractOrderBook.TransactOpts, quorumNumbers, quorumThresholdPercentage, createdBlock, orderResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOrderBook *ContractOrderBookTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOrderBook.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOrderBook *ContractOrderBookSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.SetPauserRegistry(&_ContractOrderBook.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOrderBook *ContractOrderBookTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.SetPauserRegistry(&_ContractOrderBook.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOrderBook *ContractOrderBookTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractOrderBook.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOrderBook *ContractOrderBookSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.SetStaleStakesForbidden(&_ContractOrderBook.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOrderBook *ContractOrderBookTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.SetStaleStakesForbidden(&_ContractOrderBook.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOrderBook *ContractOrderBookTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractOrderBook.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOrderBook *ContractOrderBookSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.TransferOwnership(&_ContractOrderBook.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOrderBook *ContractOrderBookTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.TransferOwnership(&_ContractOrderBook.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOrderBook *ContractOrderBookTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOrderBook.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOrderBook *ContractOrderBookSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.Unpause(&_ContractOrderBook.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOrderBook *ContractOrderBookTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOrderBook.Contract.Unpause(&_ContractOrderBook.TransactOpts, newPausedStatus)
}

// ContractOrderBookFulfillEventIterator is returned from FilterFulfillEvent and is used to iterate over the raw logs and unpacked data for FulfillEvent events raised by the ContractOrderBook contract.
type ContractOrderBookFulfillEventIterator struct {
	Event *ContractOrderBookFulfillEvent // Event containing the contract specifics and raw log

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
func (it *ContractOrderBookFulfillEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderBookFulfillEvent)
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
		it.Event = new(ContractOrderBookFulfillEvent)
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
func (it *ContractOrderBookFulfillEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderBookFulfillEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderBookFulfillEvent represents a FulfillEvent event raised by the ContractOrderBook contract.
type ContractOrderBookFulfillEvent struct {
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
func (_ContractOrderBook *ContractOrderBookFilterer) FilterFulfillEvent(opts *bind.FilterOpts) (*ContractOrderBookFulfillEventIterator, error) {

	logs, sub, err := _ContractOrderBook.contract.FilterLogs(opts, "FulfillEvent")
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookFulfillEventIterator{contract: _ContractOrderBook.contract, event: "FulfillEvent", logs: logs, sub: sub}, nil
}

// WatchFulfillEvent is a free log subscription operation binding the contract event 0x712dd31697b327ff70b0442a703553fe27c70cce6ec06c292a9cbba1ada6d780.
//
// Solidity: event FulfillEvent(bytes sig, (uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
func (_ContractOrderBook *ContractOrderBookFilterer) WatchFulfillEvent(opts *bind.WatchOpts, sink chan<- *ContractOrderBookFulfillEvent) (event.Subscription, error) {

	logs, sub, err := _ContractOrderBook.contract.WatchLogs(opts, "FulfillEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderBookFulfillEvent)
				if err := _ContractOrderBook.contract.UnpackLog(event, "FulfillEvent", log); err != nil {
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
func (_ContractOrderBook *ContractOrderBookFilterer) ParseFulfillEvent(log types.Log) (*ContractOrderBookFulfillEvent, error) {
	event := new(ContractOrderBookFulfillEvent)
	if err := _ContractOrderBook.contract.UnpackLog(event, "FulfillEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderBookInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractOrderBook contract.
type ContractOrderBookInitializedIterator struct {
	Event *ContractOrderBookInitialized // Event containing the contract specifics and raw log

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
func (it *ContractOrderBookInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderBookInitialized)
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
		it.Event = new(ContractOrderBookInitialized)
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
func (it *ContractOrderBookInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderBookInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderBookInitialized represents a Initialized event raised by the ContractOrderBook contract.
type ContractOrderBookInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOrderBook *ContractOrderBookFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractOrderBookInitializedIterator, error) {

	logs, sub, err := _ContractOrderBook.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookInitializedIterator{contract: _ContractOrderBook.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOrderBook *ContractOrderBookFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractOrderBookInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractOrderBook.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderBookInitialized)
				if err := _ContractOrderBook.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractOrderBook *ContractOrderBookFilterer) ParseInitialized(log types.Log) (*ContractOrderBookInitialized, error) {
	event := new(ContractOrderBookInitialized)
	if err := _ContractOrderBook.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderBookOrderRespondedEventIterator is returned from FilterOrderRespondedEvent and is used to iterate over the raw logs and unpacked data for OrderRespondedEvent events raised by the ContractOrderBook contract.
type ContractOrderBookOrderRespondedEventIterator struct {
	Event *ContractOrderBookOrderRespondedEvent // Event containing the contract specifics and raw log

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
func (it *ContractOrderBookOrderRespondedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderBookOrderRespondedEvent)
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
		it.Event = new(ContractOrderBookOrderRespondedEvent)
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
func (it *ContractOrderBookOrderRespondedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderBookOrderRespondedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderBookOrderRespondedEvent represents a OrderRespondedEvent event raised by the ContractOrderBook contract.
type ContractOrderBookOrderRespondedEvent struct {
	OrderResponse         IOrderBookOrderResponse
	OrderResponseMetadata IOrderBookOrderResponseMetadata
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterOrderRespondedEvent is a free log retrieval operation binding the contract event 0xa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd.
//
// Solidity: event OrderRespondedEvent((address,uint32) orderResponse, (uint32,bytes32) OrderResponseMetadata)
func (_ContractOrderBook *ContractOrderBookFilterer) FilterOrderRespondedEvent(opts *bind.FilterOpts) (*ContractOrderBookOrderRespondedEventIterator, error) {

	logs, sub, err := _ContractOrderBook.contract.FilterLogs(opts, "OrderRespondedEvent")
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookOrderRespondedEventIterator{contract: _ContractOrderBook.contract, event: "OrderRespondedEvent", logs: logs, sub: sub}, nil
}

// WatchOrderRespondedEvent is a free log subscription operation binding the contract event 0xa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd.
//
// Solidity: event OrderRespondedEvent((address,uint32) orderResponse, (uint32,bytes32) OrderResponseMetadata)
func (_ContractOrderBook *ContractOrderBookFilterer) WatchOrderRespondedEvent(opts *bind.WatchOpts, sink chan<- *ContractOrderBookOrderRespondedEvent) (event.Subscription, error) {

	logs, sub, err := _ContractOrderBook.contract.WatchLogs(opts, "OrderRespondedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderBookOrderRespondedEvent)
				if err := _ContractOrderBook.contract.UnpackLog(event, "OrderRespondedEvent", log); err != nil {
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
func (_ContractOrderBook *ContractOrderBookFilterer) ParseOrderRespondedEvent(log types.Log) (*ContractOrderBookOrderRespondedEvent, error) {
	event := new(ContractOrderBookOrderRespondedEvent)
	if err := _ContractOrderBook.contract.UnpackLog(event, "OrderRespondedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderBookOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractOrderBook contract.
type ContractOrderBookOwnershipTransferredIterator struct {
	Event *ContractOrderBookOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOrderBookOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderBookOwnershipTransferred)
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
		it.Event = new(ContractOrderBookOwnershipTransferred)
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
func (it *ContractOrderBookOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderBookOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderBookOwnershipTransferred represents a OwnershipTransferred event raised by the ContractOrderBook contract.
type ContractOrderBookOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOrderBook *ContractOrderBookFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOrderBookOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOrderBook.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookOwnershipTransferredIterator{contract: _ContractOrderBook.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOrderBook *ContractOrderBookFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOrderBookOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOrderBook.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderBookOwnershipTransferred)
				if err := _ContractOrderBook.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractOrderBook *ContractOrderBookFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOrderBookOwnershipTransferred, error) {
	event := new(ContractOrderBookOwnershipTransferred)
	if err := _ContractOrderBook.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderBookPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractOrderBook contract.
type ContractOrderBookPausedIterator struct {
	Event *ContractOrderBookPaused // Event containing the contract specifics and raw log

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
func (it *ContractOrderBookPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderBookPaused)
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
		it.Event = new(ContractOrderBookPaused)
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
func (it *ContractOrderBookPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderBookPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderBookPaused represents a Paused event raised by the ContractOrderBook contract.
type ContractOrderBookPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOrderBook *ContractOrderBookFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractOrderBookPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOrderBook.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookPausedIterator{contract: _ContractOrderBook.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOrderBook *ContractOrderBookFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractOrderBookPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOrderBook.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderBookPaused)
				if err := _ContractOrderBook.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractOrderBook *ContractOrderBookFilterer) ParsePaused(log types.Log) (*ContractOrderBookPaused, error) {
	event := new(ContractOrderBookPaused)
	if err := _ContractOrderBook.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderBookPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractOrderBook contract.
type ContractOrderBookPauserRegistrySetIterator struct {
	Event *ContractOrderBookPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractOrderBookPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderBookPauserRegistrySet)
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
		it.Event = new(ContractOrderBookPauserRegistrySet)
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
func (it *ContractOrderBookPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderBookPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderBookPauserRegistrySet represents a PauserRegistrySet event raised by the ContractOrderBook contract.
type ContractOrderBookPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOrderBook *ContractOrderBookFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractOrderBookPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractOrderBook.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookPauserRegistrySetIterator{contract: _ContractOrderBook.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOrderBook *ContractOrderBookFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractOrderBookPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractOrderBook.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderBookPauserRegistrySet)
				if err := _ContractOrderBook.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractOrderBook *ContractOrderBookFilterer) ParsePauserRegistrySet(log types.Log) (*ContractOrderBookPauserRegistrySet, error) {
	event := new(ContractOrderBookPauserRegistrySet)
	if err := _ContractOrderBook.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderBookStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractOrderBook contract.
type ContractOrderBookStaleStakesForbiddenUpdateIterator struct {
	Event *ContractOrderBookStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractOrderBookStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderBookStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractOrderBookStaleStakesForbiddenUpdate)
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
func (it *ContractOrderBookStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderBookStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderBookStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractOrderBook contract.
type ContractOrderBookStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractOrderBook *ContractOrderBookFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractOrderBookStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractOrderBook.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookStaleStakesForbiddenUpdateIterator{contract: _ContractOrderBook.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractOrderBook *ContractOrderBookFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractOrderBookStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractOrderBook.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderBookStaleStakesForbiddenUpdate)
				if err := _ContractOrderBook.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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
func (_ContractOrderBook *ContractOrderBookFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractOrderBookStaleStakesForbiddenUpdate, error) {
	event := new(ContractOrderBookStaleStakesForbiddenUpdate)
	if err := _ContractOrderBook.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderBookUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractOrderBook contract.
type ContractOrderBookUnpausedIterator struct {
	Event *ContractOrderBookUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractOrderBookUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderBookUnpaused)
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
		it.Event = new(ContractOrderBookUnpaused)
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
func (it *ContractOrderBookUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderBookUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderBookUnpaused represents a Unpaused event raised by the ContractOrderBook contract.
type ContractOrderBookUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOrderBook *ContractOrderBookFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractOrderBookUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOrderBook.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractOrderBookUnpausedIterator{contract: _ContractOrderBook.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOrderBook *ContractOrderBookFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractOrderBookUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOrderBook.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderBookUnpaused)
				if err := _ContractOrderBook.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractOrderBook *ContractOrderBookFilterer) ParseUnpaused(log types.Log) (*ContractOrderBookUnpaused, error) {
	event := new(ContractOrderBookUnpaused)
	if err := _ContractOrderBook.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

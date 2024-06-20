import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "../lib/eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@openzeppelin/contracts/utils/cryptography/SignatureChecker.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Settlement is
Initializable,
OwnableUpgradeable,
Pausable,
BLSSignatureChecker,
OperatorStateRetriever
{
    using BN254 for BN254.G1Point;
    // fulfill event
    event fulfillEvent(
        Order order,
        uint32 quorumThresholdPercentage,
        bytes quorumNumbers
    );
    event OrderRespondedEvent(
        OrderResponse orderResponse,
        OrderResponseMetadata OrderResponseMetadata
    );
    /* STORAGE */
    // mapping of order indices to all order hashes
    // when an order is fulfilled, order hash is stored here,
    // and responses need to pass the actual order,
    // which is hashed onchain and checked against this mapping
    mapping(uint32 => bytes32) public allOrderHashes;
    mapping(uint32 => Order) public allOrderDetails;
    // mapping of order indices to hash of abi.encode(taskResponse, taskResponseMetadata)
    mapping(uint32 => bytes32) public allOrderResponses;

    /* CONSTANT */
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;
    address public aggregator;

    error TakerMismatch();
    error OrderExpired();
    error InvalidSignature();

    // STRUCTS
    struct Order {
        uint32 orderId;
        address maker;
        address taker;
        address inputToken;
        uint256 inputAmount;
        address outputToken;
        uint256 outputAmount;
        uint32 quorumThresholdPercentage;
        bytes quorumNumbers;
        uint256 expiry;
        uint32 targetNetworkNumber;
        uint32 createdBlock;
    }

    // Order response is hashed and signed by operators.
    // these signatures are aggregated and sent to the contract as response.
    struct OrderResponse {
        uint32 referenceOrderIndex;
        // This is just the response that the operator has to compute by itself.
        bool txSuccess;
    }

    // Extra information related to orderResponse, which is filled inside the contract.
    // It thus cannot be signed by operators, so we keep it in a separate struct than OrderResponse
    // This metadata is needed by the challenger, so we emit it in the OrderResponsed event
    struct OrderResponseMetadata {
        uint32 orderResponsedBlock;
        bytes32 hashOfNonSigners;
    }

    /* MODIFIERS */
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator
    ) BLSSignatureChecker(_registryCoordinator) {
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        address _aggregator
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
        aggregator = _aggregator;
    }

    function fulfill(Order memory order, bytes calldata sig) public {
        // check taker address
        if (order.taker != address(0) && msg.sender != order.taker) revert TakerMismatch();
        if (order.expiry < block.timestamp) revert OrderExpired();
        // prepare data to verify signature
        bytes32 hash = keccak256(abi.encodePacked(order.orderId, order.maker, order.taker, order.inputToken, order.inputAmount, order.outputToken, order.outputAmount, order.expiry, order.targetNetworkNumber));
        // recover the signer's address
        bool validSig = SignatureChecker.isValidSignatureNow(order.maker, ECDSA.toEthSignedMessageHash(hash), sig);
        // check that the signer is the maker
        if (!validSig) revert InvalidSignature();
        // transfer the tokens
        IERC20(order.inputToken).transferFrom(order.maker, msg.sender, order.inputAmount);
        IERC20(order.outputToken).transferFrom(msg.sender, order.maker, order.outputAmount);
        // set order taker
        order.taker = msg.sender;
        order.createdBlock = uint32(block.number);
        allOrderHashes[order.orderId] = keccak256(abi.encode(order));
        allOrderDetails[order.orderId] = order;
        emit fulfillEvent(order, order.quorumThresholdPercentage, order.quorumNumbers);
    }

    // NOTE: this function responds to existing tasks.
    function respondToFulfill(
        Order calldata order,
        OrderResponse calldata orderResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        bytes calldata quorumNumbers = order.quorumNumbers;
        uint32 quorumThresholdPercentage = order.quorumThresholdPercentage;
        uint32 createdBlock = order.createdBlock;

        // check that the task is valid, hasn't been responsed yet, and is being responsed in time
        require(
            keccak256(abi.encode(order)) ==
            allOrderHashes[orderResponse.referenceOrderIndex],
            "supplied order does not match the one recorded in the contract"
        );
        // some logical checks
        require(
            allOrderResponses[orderResponse.referenceOrderIndex] == bytes32(0),
            "Aggregator has already responded to the order"
        );
        /* CHECKING SIGNATURES & WHETHER THRESHOLD IS MET OR NOT */
        // calculate message which operators signed
        bytes32 message = keccak256(abi.encode(orderResponse));

        // check the BLS signature
        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
            message,
            quorumNumbers,
            createdBlock,
            nonSignerStakesAndSignature
        );

        // check that signatories own at least a threshold percentage of each quourm
        for (uint i = 0; i < quorumNumbers.length; i++) {
            // we don't check that the quorumThresholdPercentages are not >100 because a greater value would trivially fail the check, implying
            // signed stake > total stake
            require(
                quorumStakeTotals.signedStakeForQuorum[i] *
                _THRESHOLD_DENOMINATOR >=
                quorumStakeTotals.totalStakeForQuorum[i] *
                uint8(quorumThresholdPercentage),
                "Signatories do not own at least threshold percentage of a quorum"
            );
        }

        OrderResponseMetadata memory orderResponseMetadata = OrderResponseMetadata(
            uint32(block.number),
            hashOfNonSigners
        );
        // updating the storage with order response
        allOrderResponses[orderResponse.referenceOrderIndex] = keccak256(
            abi.encode(orderResponse, orderResponseMetadata)
        );

        // emitting event
        emit OrderRespondedEvent(orderResponse, orderResponseMetadata);
    }

}

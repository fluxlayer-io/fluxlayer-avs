contract ISettlement {
    // EVENTS
    event FulfillEvent(
        Order order,
        uint32 quorumThresholdPercentage,
        bytes quorumNumbers,
        address recipient
    );
    event OrderRespondedEvent(
        OrderResponse orderResponse,
        OrderResponseMetadata OrderResponseMetadata
    );

    // ERRORS
    error OrderFulfilled();
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
        uint256 expiry;
        uint32 targetNetworkNumber;
    }

    struct OrderExecution {
        uint32 quorumThresholdPercentage;
        bytes quorumNumbers;
        uint32 createdBlock;
    }

    // Order response is hashed and signed by operators.
    // these signatures are aggregated and sent to the contract as response.
    struct OrderResponse {
        address recipient;
        uint32 referenceOrderIndex;
    }

    // Extra information related to orderResponse, which is filled inside the contract.
    // It thus cannot be signed by operators, so we keep it in a separate struct than OrderResponse
    // This metadata is needed by the challenger, so we emit it in the OrderResponsed event
    struct OrderResponseMetadata {
        uint32 orderResponsedBlock;
        bytes32 hashOfNonSigners;
    }
}

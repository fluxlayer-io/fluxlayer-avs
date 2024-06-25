// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";
import "./OrderBook.sol";

/**
 * @title Primary entrypoint for procuring services from IncredibleSquaring.
 * @author Layr Labs, Inc.
 */
contract FluxLayerServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    OrderBook
        public immutable orderBook;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlySettlement() {
        require(
            msg.sender == address(orderBook),
            "onlyOrderBook: not from orderBook"
        );
        _;
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        OrderBook _orderBook
    )
        ServiceManagerBase(
            _avsDirectory,
            IPaymentCoordinator(address(0)), // inc-sq doesn't need to deal with payments
            _registryCoordinator,
            _stakeRegistry
        )
    {
        orderBook = _orderBook;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(
        address operatorAddr
    ) external onlySettlement {
        // slasher.freezeOperator(operatorAddr);
    }
}

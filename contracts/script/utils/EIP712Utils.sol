pragma solidity ^0.8.12;

import {IOrderBook} from "../../src/IOrderBook.sol";
import "forge-std/console.sol";

contract EIP712Utils {
    bytes32 public constant ORDER_TYPEHASH = keccak256("Order(uint32 orderId,address maker,address taker,address inputToken,uint256 inputAmount,address outputToken,uint256 outputAmount,uint256 expiry,uint32 targetNetworkNumber)");
    bytes32 DOMAIN_SEPARATOR;

    constructor(string memory name, string memory version, uint32 chainId, address target) {
        DOMAIN_SEPARATOR = _buildDomainSeparator(name, version, chainId, target);
        console.log("DOMAIN_SEPARATOR");
        console.logBytes32(DOMAIN_SEPARATOR);
        console.log("ORDER_TYPEHASH");
        console.logBytes32(ORDER_TYPEHASH);
    }


    function _buildDomainSeparator(
        string memory name,
        string memory version,
        uint32 chainId,
        address verifyingContract
    ) private view returns (bytes32) {
        bytes32 typeHash = keccak256(
            "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
        );
        bytes32 nameHash = keccak256(bytes(name));
        bytes32 versionHash = keccak256(bytes(version));
        return keccak256(abi.encode(typeHash, nameHash, versionHash, chainId, verifyingContract));
    }

    // computes the hash of a permit
    function getStructHash(IOrderBook.Order memory order)
    internal
    pure
    returns (bytes32)
    {
        return
            keccak256(
            abi.encode(
                ORDER_TYPEHASH,
                order.orderId,
                order.maker,
                order.taker,
                order.inputToken,
                order.inputAmount,
                order.outputToken,
                order.outputAmount,
                order.expiry,
                order.targetNetworkNumber
            )
        );
    }

    // computes the hash of the fully encoded EIP-712 message for the domain, which can be used to recover the signer
    function getTypedDataHash(IOrderBook.Order memory order)
    public
    view
    returns (bytes32)
    {
        return
            keccak256(
            abi.encodePacked(
                "\x19\x01",
                DOMAIN_SEPARATOR,
                getStructHash(order)
            )
        );
    }
}


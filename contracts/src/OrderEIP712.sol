pragma solidity ^0.8.12;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./IOrderBook.sol";
import "forge-std/console.sol";


contract OrderEIP712 {
    using ECDSA for bytes32;

    bytes32 private constant _DOMAIN_TYPE_HASH = keccak256(
        "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
    );
    bytes32 private immutable _HASHED_NAME;
    bytes32 private immutable _HASHED_VERSION;
    bytes32 private immutable _ORDER_TYPE_HASH = keccak256("Order(uint32 orderId,address maker,address taker,address inputToken,uint256 inputAmount,address outputToken,uint256 outputAmount,uint256 expiry,uint32 targetNetworkNumber)");
    uint32 private immutable _CHAIN_ID;
    address _ORDERBOOK_ADDR;

    constructor(string memory name, string memory version, uint32 chainId, address orderBookAddr) {
        bytes32 hashedName = keccak256(bytes(name));
        bytes32 hashedVersion = keccak256(bytes(version));
        _HASHED_NAME = hashedName;
        _HASHED_VERSION = hashedVersion;
        _CHAIN_ID = chainId;
        if (orderBookAddr != address(0)) {
            _ORDERBOOK_ADDR = orderBookAddr;
        }
    }

    function _getOrderAddr() internal view returns (address) {
        if (_ORDERBOOK_ADDR != address(0)) {
            return _ORDERBOOK_ADDR;
        } else {
            return address(this);
        }
    }


    function _domainSeparatorV4() internal view returns (bytes32) {
        return keccak256(abi.encode(
            _DOMAIN_TYPE_HASH,
            _HASHED_NAME,
            _HASHED_VERSION,
            _CHAIN_ID,
            _getOrderAddr()
        ));
    }

    function _hashTypedDataV4(bytes32 structHash) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", _domainSeparatorV4(), structHash));
    }

    function hashOrder(IOrderBook.Order memory order) public view returns (bytes32) {
        return _hashTypedDataV4(keccak256(abi.encode(
            _ORDER_TYPE_HASH,
            order.orderId,
            order.maker,
            order.taker,
            order.inputToken,
            order.inputAmount,
            order.outputToken,
            order.outputAmount,
            order.expiry,
            order.targetNetworkNumber
        )));
    }

    function verify(IOrderBook.Order memory order, bytes calldata signature) public view returns (bool) {
        address signer = hashOrder(order).recover(signature);
        return signer == order.maker;
    }

}
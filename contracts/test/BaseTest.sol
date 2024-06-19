import "forge-std/Test.sol";

contract BaseTest is Test {
    function signatureToBytes(bytes32 r, bytes32 s, uint8 v) internal pure returns (bytes memory) {
        bytes memory signature = new bytes(65);
        assembly {
            mstore(add(signature, 0x20), r)
            mstore(add(signature, 0x40), s)
            mstore(add(signature, 0x60), shl(248, v))
        }
        return signature;
    }

    function signContent(uint256 pk, bytes memory content) internal pure returns (bytes memory) {
        bytes32 hash = keccak256(content);
        hash = toEthSignedMessageHash(hash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, hash);
        return signatureToBytes(r, s, v);
    }

    function toEthSignedMessageHash(bytes32 hash) internal pure returns (bytes32) {
        // 32 is the length in bytes of hash,
        // enforced by the type signature above
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }
}
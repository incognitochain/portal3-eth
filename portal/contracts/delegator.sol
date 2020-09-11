pragma solidity 0.6.6;

import "./pause.sol";

contract Delegator is AdminPausable {

    address constant public ETH_TOKEN = 0x0000000000000000000000000000000000000000;
    address public delegator;

    constructor(address _admin, address _delegator) public {
        require(_admin != address(0) && _delegator != address(0));
        admin = _admin;
        paused = false;
        expire = block.timestamp + 365 * 1 days;
        delegator = _delegator;
    }

    /**
     * @dev Payable receive function to receive Ether from oldVault when migrating
     */
    receive() external payable {}

    fallback() external payable {
        address _target = delegator;
        assembly {
            // 0x40 is the address where the next free memory slot is stored in Solidity
            let _calldataMemoryOffset := mload(0x40)
            let _calldatasize := calldatasize()
            // new "memory end" including padding. The bitwise operations here ensure we get rounded up to the nearest 32 byte boundary
            let _size := and(add(_calldatasize, 0x1f), not(0x1f))
            // Update the pointer at 0x40 to point at new free memory location so any theoretical allocation doesn't stomp our memory in this call
            mstore(0x40, add(_calldataMemoryOffset, _size))
            // Copy method signature and parameters of this call into memory
            calldatacopy(_calldataMemoryOffset, 0x0, _calldatasize)
            // Call the actual method via delegation
            let _retval := delegatecall(gas(), _target, _calldataMemoryOffset, _calldatasize, 0, 0)
            switch _retval
            case 0 {
                // 0 == it threw, so we revert
                revert(0,0)
            } default {
                // If the call succeeded return the return data from the delegate call
                let _returndataMemoryOffset := mload(0x40)
                let _returndatasize := returndatasize()
                // Update the pointer at 0x40 again to point at new free memory location so any theoretical allocation doesn't stomp our memory in this call
                mstore(0x40, add(_returndataMemoryOffset, _returndatasize))
                returndatacopy(_returndataMemoryOffset, 0x0, _returndatasize)
                return(_returndataMemoryOffset, returndatasize())
            }
        }
    }
}
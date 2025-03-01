// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PrintMyAsset {
    function printEth() public view returns (uint256) {
        return address(msg.sender).balance;
    }
}
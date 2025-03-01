pragma solidity ^0.8.0;

contract CheckAsset {
    function checkBalance() public view returns (uint256) {
        return address(this).balance;
    }
}
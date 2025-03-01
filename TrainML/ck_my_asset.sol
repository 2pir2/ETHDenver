// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AssetChecker {
    uint256 private _assetAmount;

    constructor(uint256 initialAmount) {
        _assetAmount = initialAmount;
    }

    function checkAssetAmount() public view returns (uint256) {
        return _assetAmount;
    }
}
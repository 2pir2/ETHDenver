// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AssetManager {
    struct Asset {
        uint256 id;
        string name;
        uint256 value;
        address owner;
    }

    mapping(uint256 => Asset) private assets;
    uint256 private nextAssetId = 1;

    event AssetCreated(uint256 assetId, string name, uint256 value, address owner);
    event AssetInfo(uint256 assetId, string name, uint256 value, address owner);

    function createAsset(string memory name, uint256 value) public {
        uint256 assetId = nextAssetId++;
        assets[assetId] = Asset(assetId, name, value, msg.sender);
        emit AssetCreated(assetId, name, value, msg.sender);
    }

    function getAsset(uint256 assetId) public view returns (Asset memory) {
        Asset memory asset = assets[assetId];
        require(asset.id != 0, "Asset does not exist");
        return asset;
    }

    function emitAssetInfo(uint256 assetId) public {
        Asset memory asset = getAsset(assetId);
        emit AssetInfo(asset.id, asset.name, asset.value, asset.owner);
    }
}
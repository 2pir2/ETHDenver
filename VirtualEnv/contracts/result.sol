// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AssetManager {
    // Define an Asset structure
    struct Asset {
        uint256 id;
        string description;
        address owner;
    }

    // Mapping to store assets
    mapping(uint256 => Asset) private assets;

    // Variable to keep track of the next asset ID
    uint256 private nextAssetId = 1;

    // Event to log the addition of a new asset
    event AssetAdded(uint256 indexed assetId, string description, address owner);

    // Function to add a new asset
    function addAsset(string memory description) public {
        // Create a new asset
        Asset memory newAsset = Asset({
            id: nextAssetId,
            description: description,
            owner: msg.sender
        });

        // Store the asset in the mapping
        assets[nextAssetId] = newAsset;

        // Emit an event
        emit AssetAdded(nextAssetId, description, msg.sender);

        // Increment the asset ID for the next asset
        nextAssetId++;
    }

    // Function to get an asset by ID
    function getAsset(uint256 assetId) public view returns (uint256, string memory, address) {
        // Check if the asset exists
        require(assetId > 0 && assetId < nextAssetId, "Asset does not exist");

        // Get the asset from the mapping
        Asset memory asset = assets[assetId];

        // Return the asset details
        return (asset.id, asset.description, asset.owner);
    }

    // Function to get the total number of assets
    function getTotalAssets() public view returns (uint256) {
        return nextAssetId - 1;
    }
}
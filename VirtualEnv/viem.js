import { createWalletClient, createPublicClient, webSocket, parseEther } from 'viem'
import { privateKeyToAccount } from 'viem/accounts'
import { sepolia } from 'viem/chains'
import fs from 'fs'

// ✅ Load Contract ABI & Bytecode
const contractJsonPath = "/Users/hanxu/Desktop/ETHDenver/VirtualEnv/artifacts/contracts/result.sol/AssetManager.json"
const contractJson = JSON.parse(fs.readFileSync(contractJsonPath, 'utf8'))

const contractAbi = contractJson.abi
const contractBytecode = contractJson.bytecode

// ✅ Private Key Setup
const privateKey = "0x0fcf0b8c984a46d9344f465a9af73d9edeb835e358dfbcf5925591b0a7a43551"
const account = privateKeyToAccount(privateKey)

console.log("Using Wallet Address:", account.address)

// ✅ Clients for Transactions & Events
const walletClient = createWalletClient({
    account,
    chain: sepolia,
    transport: webSocket("wss://sepolia.drpc.org")
})

const publicClient = createPublicClient({
    chain: sepolia,
    transport: webSocket("wss://sepolia.drpc.org")
})

// ✅ Function to Deploy Contract and Interact
async function deployAndInteract() {
    console.log("🚀 Deploying contract...")

    // ✅ Deploy Contract
    const hash = await walletClient.deployContract({
        abi: contractAbi,
        bytecode: contractBytecode,
    })

    console.log("📜 Deployment Transaction Hash:", hash)

    // ✅ Wait for contract to be deployed
    const receipt = await publicClient.waitForTransactionReceipt({ hash })
    const contractAddress = receipt.contractAddress

    console.log("✅ Contract Deployed at:", contractAddress)

    // ✅ Interact with Contract (Call addAsset)
    console.log("📌 Creating an asset on deployed contract...")

    const createAssetHash = await walletClient.writeContract({
        address: contractAddress,
        abi: contractAbi,
        functionName: "addAsset",
        args: ["This is a test description"], // ✅ Fix: Only passing `description`
    })

    console.log("📜 Asset Creation Transaction Hash:", createAssetHash)

    // ✅ Wait for transaction confirmation
    const assetReceipt = await publicClient.waitForTransactionReceipt({ hash: createAssetHash })

    console.log("✅ Asset Created")

    // ✅ Extract logs from events
    for (const log of assetReceipt.logs) {
        console.log("📜 Event Log:", log)
    }

    // ✅ Save logs for ZK-SNARK proof
    fs.writeFileSync(
        "deployment_logs.json",
        JSON.stringify(assetReceipt.logs, (key, value) =>
          typeof value === "bigint" ? value.toString() : value, // Convert BigInt to String
        null, 2)
    )
    console.log("Log Saved!")
}

deployAndInteract().catch(console.error)

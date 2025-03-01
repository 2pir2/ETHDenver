import { createWalletClient, createPublicClient, webSocket, parseEther } from 'viem'
import { privateKeyToAccount } from 'viem/accounts'
import { sepolia } from 'viem/chains'
import fs from 'fs'

// ✅ Load Contract ABI & Bytecode
const contractJsonPath = "/Users/hanxu/Desktop/ETHDenver/VirtualEnv/artifacts/contracts/printasset.sol/PrintAsset.json"
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

// ✅ Function to Deploy Contract & Capture Logs
async function deployContract() {
    console.log("🚀 Deploying contract...")

    const hash = await walletClient.sendTransaction({
        account,
        data: contractBytecode, // ✅ Bytecode for deployment
    })

    console.log("📜 Deployment Transaction Hash:", hash)

    // ✅ Wait for transaction confirmation
    const receipt = await publicClient.waitForTransactionReceipt({ hash })

    console.log("✅ Contract Deployed at:", receipt.contractAddress)

    // ✅ Extract logs from events
    for (const log of receipt.logs) {
        console.log("📜 Event Log:", log)
    }

    // ✅ Save logs for ZK-SNARK proof
    fs.writeFileSync(
        "deployment_logs.json",
        JSON.stringify(receipt.logs, (key, value) =>
          typeof value === "bigint" ? value.toString() : value, // Convert BigInt to String
        null, 2)
      )
      
}

deployContract().catch(console.error)

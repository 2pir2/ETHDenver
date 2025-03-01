from web3 import Web3

# Connect to Ethereum (Sepolia Testnet)
w3 = Web3(Web3.HTTPProvider("https://sepolia.infura.io/v3/0xbcaeb05D15c61E5ABf4dB475D8A459449fcD22df"))

# Smart contract address
contract_address = "0xYourSwapContractAddress"
contract_abi = "your_contract_abi.json"

# Load contract
contract = w3.eth.contract(address=contract_address, abi=contract_abi)

# Fetch contract BTC balance before swap
contract_btc_before = contract.functions.getContractBTCBalance().call()

# Fetch user BTC balance before swap
user_address = "0xUserEthereumAddress"
user_btc_before = contract.functions.getUserBTCBalance(user_address).call()

print(f"Contract BTC Before Swap: {contract_btc_before}")
print(f"User BTC Before Swap: {user_btc_before}")

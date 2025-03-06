from openai import OpenAI

client = OpenAI(api_key="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")


# Step 1: Ask thes user what kind of smart contract they want
contract_type = input("What kind of smart contract do you want to create? (e.g., Loan, Flash Loan, Token, etc.) ").strip()

# Step 2: Customize questions based on contract type
additional_details = ""
if contract_type.lower() == "loan":
    loan_type = input("Is it a simple loan, collateralized loan, or flash loan? ").strip()
    interest_rate = input("Enter interest rate (e.g., 5%): ").strip()
    repayment_terms = input("Describe repayment terms (e.g., monthly, on-demand, etc.): ").strip()
    
    additional_details = f"Type: {loan_type}, Interest Rate: {interest_rate}, Repayment Terms: {repayment_terms}"

elif contract_type.lower() == "token":
    token_name = input("Enter token name: ").strip()
    symbol = input("Enter token symbol: ").strip()
    supply = input("Enter total supply: ").strip()

    additional_details = f"Token Name: {token_name}, Symbol: {symbol}, Total Supply: {supply}"

else:
    additional_details = input("Describe the functionality of your smart contract: ").strip()

# Step 3: Generate the Solidity contract using OpenAI API
prompt = f"Write a Solidity smart contract for a {contract_type} with the following details: {additional_details}."

response = client.chat.completions.create(
    model="ft:gpt-4o-2024-08-06:personal::B5jlW6nC",
    messages=[{"role": "user", "content": prompt}],
    temperature=0.7
)

# Extract Solidity code
solidity_code = response.choices[0].message.content

# Step 4: Save the contract to a .sol file
file_name = f"{contract_type.lower().replace(' ', '_')}.sol"
with open(file_name, "w") as file:
    file.write(solidity_code)

print(f"Solidity smart contract saved as {file_name}")

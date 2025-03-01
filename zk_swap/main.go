package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

// ✅ zk-SNARK Circuit for Verification
type ContractVerificationCircuit struct {
	DeployerAddr frontend.Variable
	ContractAddr frontend.Variable
	TxHash       frontend.Variable
	EventHash    frontend.Variable `gnark:",public"` // ✅ Public Proof Output
}

// ✅ Define Circuit Logic
func (circuit *ContractVerificationCircuit) Define(api frontend.API) error {
	// ✅ Convert frontend.Variable to *big.Int properly
	deployerBig := toBigInt(circuit.DeployerAddr)
	contractBig := toBigInt(circuit.ContractAddr)
	txHashBig := toBigInt(circuit.TxHash)

	// ✅ Compute Poseidon Hash
	hashedValue, err := poseidon.Hash([]*big.Int{
		deployerBig,
		contractBig,
		txHashBig,
	})
	if err != nil {
		return err
	}

	// ✅ Ensure computed hash matches the event log hash
	api.AssertIsEqual(frontend.Variable(hashedValue), circuit.EventHash)

	return nil
}

// ✅ Convert frontend.Variable to *big.Int properly (ENSURES NON-NIL RETURN)
func toBigInt(v frontend.Variable) *big.Int {
	if v == nil {

		return big.NewInt(0) // Default to 0 instead of nil
	}

	b, ok := new(big.Int).SetString(fmt.Sprintf("%v", v), 10)
	if !ok {

		return big.NewInt(0) // Default to 0
	}
	return b
}

// ✅ Convert Hex String to frontend.Variable using BigInt (ENSURES NON-NIL RETURN)
func stringToVariable(str string) frontend.Variable {
	bigInt := new(big.Int)
	a, success := bigInt.SetString(str, 16) // Convert hex string to BigInt
	fmt.Print("This is A", a)
	if !success {

		return frontend.Variable(big.NewInt(0)) // Return 0 instead of nil
	}
	return frontend.Variable(bigInt)
}

func main() {
	// ✅ Load the Deployment Logs (only necessary data)
	logFilePath := "/Users/hanxu/Desktop/ETHDenver/VirtualEnv/deployment_logs.json"
	logData, _ := os.ReadFile(logFilePath)

	// ✅ Parse JSON as an array of logs
	var logs []map[string]interface{}
	_ = json.Unmarshal(logData, &logs)

	logEntry := logs[0]                                        // ✅ Use the first log entry
	contract := logEntry["address"].(string)                   // Contract Address (Hex)
	txHash := logEntry["transactionHash"].(string)             // Deployment Transaction Hash (Hex)
	deployer := logEntry["topics"].([]interface{})[1].(string) // Deployer Address (Hex)

	// ✅ Convert Hex String → frontend.Variable
	deployerNum := stringToVariable(deployer)
	contractNum := stringToVariable(contract)
	txHashNum := stringToVariable(txHash)

	// ✅ Compute Poseidon Hash
	hashInputs := []*big.Int{
		toBigInt(deployerNum),
		toBigInt(contractNum),
		toBigInt(txHashNum),
	}
	hashedValue, _ := poseidon.Hash(hashInputs)

	// ✅ Create zk-SNARK Circuit Instance
	var circuit ContractVerificationCircuit
	assignment := &ContractVerificationCircuit{
		DeployerAddr: deployerNum,
		ContractAddr: contractNum,
		TxHash:       txHashNum,
		EventHash:    frontend.Variable(hashedValue),
	}

	// ✅ Fix: Explicitly define BN254 Scalar Field during Compilation
	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		fmt.Println("❌ Error compiling circuit:", err)
		return
	}

	// ✅ Generate zk-SNARK Proof
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		fmt.Println("❌ Error setting up SNARK:", err)
		return
	}

	witness, err := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	if err != nil {
		fmt.Println("❌ Error creating witness:", err)
		return
	}

	// ✅ Extract `publicWitness` Correctly
	publicWitness, err := witness.Public()
	if err != nil {
		fmt.Println("❌ Error getting public witness:", err)
		return
	}

	// ✅ Generate Proof
	proof, err := groth16.Prove(r1cs, pk, witness)
	if err != nil {
		fmt.Println("❌ Error generating proof:", err)
		return
	}

	// ✅ Verify Proof
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		fmt.Println("❌ Verification failed")
	} else {
		fmt.Println("✅ Deployment Verified Successfully")
		fmt.Println("🔹 ZK-SNARK Proof:", proof)
	}
}

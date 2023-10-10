package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"github.com/aurora-is-near/near-api-go"
	"log"
	"math/big"
	"os"
)

func main() {
	// Set environment variable for network choice
	os.Setenv("NEAR_ENV", "testnet")

	config := near.GetConfig()

	accountID := "c_test.testnet"
	privateKeyData := "yourPrivateKeyInBytesFormat" // Convert your private key to bytes format if it's in string
	privateKey := ed25519.NewKeyFromSeed(privateKeyData)
	// Create a new NEAR connection using the configuration
	connection := near.NewConnection(config.NodeURL) 

	// Set up the contract details
	contractID := "your.deployed.contract.id"
	methodName := "send_message"
	args := map[string]interface{}{
		"message": "Hello from Go!",
	}

	encodedArgs, err := near.Serialize(args) // Assuming Serialize exists in the near package to encode the arguments
	if err != nil {
		log.Fatal(err)
	}
	functionCallAction := near.Action{
		Enum: 3,  // For FunctionCall
		FunctionCall: near.FunctionCall{
			MethodName: methodName,
			Args:       encodedArgs,
			Gas:        10000000000000,
			Deposit:    big.NewInt(0),
		},
	}

	// Create and sign the transaction
	blockHash, _ := hex.DecodeString("YOUR_BLOCK_HASH") // Replace with an appropriate block hash
	actions := []near.Action{functionCallAction}
	publicKeyBytes, _ := hex.DecodeString("YOUR_PUBLIC_KEY_HEX") // Replace with your account's public key in hex format
	txHash, signedTx, err := near.signTransaction(
		contractID,
		account.Nonce + 1, // Incrementing the nonce by 1 for this transaction
		actions,
		blockHash,
		ed25519.PublicKey(publicKeyBytes),
		account.PrivateKey,  // Assuming this is how you retrieve the private key
		accountID,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Send the transaction
	result, err := near.SendTransaction(connection, signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// Handle the result
	log.Printf("Result: %v", result)
}


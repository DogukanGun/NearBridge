package contractCaller

import (
	"encoding/json"
	"errors"
	"github.com/aurora-is-near/near-api-go"
	"log"
	"math/big"
	"os"
)

type MyFunctionArgs struct {
	Message string `json:"message"`
}

func InteractWithContract(methodName string, message string) (interface{}, error) {

	// Setting up the environment variables for testing
	os.Setenv("NEAR_ENV", "testnet")
	os.Setenv("NEAR_ACCOUNT_ID", "c_test.testnet")

	//TODO: check more usage of putting keypath in config
	var config = near.GetConfig()
	var connection = near.NewConnection(config.NodeURL)
	accountID := os.Getenv("NEAR_ACCOUNT_ID")
	if accountID == "" {
		return nil, errors.New("environment variable NEAR_ACCOUNT_ID is not set")
	}

	// Load the account using the connection and the provided credentials
	account, err := near.LoadAccount(connection, config, accountID)
	if err != nil {
		return nil, err
	}

	// Convert your arguments to a serialized format
	argsStruct := MyFunctionArgs{Message: message}
	args, err := json.Marshal(argsStruct) //borsh.Serialize(argsStruct)
	if err != nil {
		log.Fatalf("Error serializing arguments: %v", err)
	}

	// Using the provided FunctionCall inside acccount.go
	amount := *big.NewInt(0)
	result, err := account.FunctionCall("c_test.testnet", methodName, args, 10000000000000, amount)
	if err != nil {
		return nil, err
	}

	return result, nil
}

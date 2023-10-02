package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitializeENV(varNames []string, envFileName string) (bool, string, map[string]string) {
	//Initialize the map object
	theMap := map[string]string{}
	targetVar := ""

	// First check if env file exists & load the file if it exists
	dotenvErr := godotenv.Load(envFileName)

	if dotenvErr == nil {
		log.Println(".env file located and loaded. Running on Development Mode")
	}

	if dotenvErr != nil {
		log.Println("Failed to load .env file: ", dotenvErr)
		log.Println("Switching to the Production Mode")
	}

	for _, varName := range varNames {
		// Get the variable from the env
		envVar := os.Getenv(varName)

		// Check if the variable has been initialized
		if envVar != "" {
			// Add the variable to the map
			theMap[varName] = envVar
		} else {
			// Indicate that there is a variable that has not been initialized and leave the handling to the user
			targetVar = varName
			return false, targetVar, theMap
		}
	}

	return true, targetVar, theMap
}

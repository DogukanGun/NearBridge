package main

import (
	"github.com/DogukanGun/NearBridge/bridge/nearclient/contract_caller"
	"log"
)

func main() {

	result, err := contract_caller.InteractWithContract("send_message", "Hello from Go! new structure")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Result: %v", result)
}

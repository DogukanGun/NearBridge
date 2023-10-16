package contractCaller

import (
	"github.com/DogukanGun/NearBridge/bridge/nearclient/contractCaller"
	"log"
)

func runExample() {

	result, err := contractCaller.InteractWithContract("send_message", "Hello from Go! new structure")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Result: %v", result)
}

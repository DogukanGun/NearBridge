package main

import (
	"fmt"
	"net/http"
	"networklistener/chainlistener"
	"networklistener/router"
	"networklistener/utils"
	"os"
)

func main() {
	envNames := []string{"SECRET", "HOST", "NEAR_RPC", "NEAR_WEB_SOCKET", "PRIVATE_KEY"}
	status, missingName, envMap := utils.InitializeENV(envNames, "main.env")

	if !status {
		fmt.Println(missingName + " was not initialized... Aborting!")
	}

	rtr := router.New(envMap["SECRET"])
	err := http.ListenAndServe(envMap["HOST"], rtr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	chainlistener.StartNetworkListening()
}

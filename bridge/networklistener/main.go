package main

import (
	"fmt"
	"net/http"
	"networklistener/router"
	"networklistener/utils"
	"os"
)

func main() {
	envNames := []string{"SECRET", "HOST"}
	status, missingName, envMap := utils.InitializeENV(envNames, "rtm.env")

	if !status {
		fmt.Println(missingName + " was not initialized... Aborting!")
	}

	rtr := router.New(envMap["SECRET"])
	err := http.ListenAndServe(envMap["HOST"], rtr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

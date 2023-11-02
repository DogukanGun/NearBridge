package controller

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"networklistener/contracts"
	bridgeControllerData "networklistener/data"
	"networklistener/db"
	httpResponse "networklistener/response"
	"networklistener/utils"
)

/*
BridgeController type
*/
type BridgeController struct{}

func (bc *BridgeController) CreateBridge() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData bridgeControllerData.CreateBridgeRequest
		if err := c.ShouldBindJSON(&reqData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		dbHandler := db.DatabaseHandler{
			TargetTable: "Bridges",
		}
		client := dbHandler.ConnectMongo()
		status, err := dbHandler.InsertData(client, reqData)
		envNames := []string{"PRIVATE_KEY"}
		status, missingName, envMap := utils.InitializeENV(envNames, "../.env")
		if !status {
			fmt.Println(missingName + " was not initialized... Aborting!")
		}
		bridgeInst := bridgeControllerData.Chain{}
		bridgeFilter := bridgeControllerData.Chain{
			ChainID: reqData.ChainID,
		}
		err = dbHandler.GetData(client, "Chain", bridgeFilter, &bridgeInst)
		if err != nil {
			fmt.Println(err)
			return
		}
		err, contractInstance, _, res, _ := contracts.CreateFunctionRequirementsForBridge(reqData.TargetURL, "", envMap["PRIVATE_KEY"])
		targetContractAddress := common.HexToAddress(reqData.TargetContractAddress)
		result, err := contractInstance.AddUserContract(res, reqData.UserEmail, targetContractAddress)
		fmt.Println(result)
		if err != nil {
			return
		}
		if err != nil {
			fmt.Println(err)
		}
		if status {
			httpResponse.CreateSuccessfulResponse(c, 200, "Your bridge is created", nil)
		}

	}
}

func (bc *BridgeController) SendMessage() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

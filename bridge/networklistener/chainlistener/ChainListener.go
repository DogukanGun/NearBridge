package chainlistener

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	chainListenerEvents "networklistener/chainlistener/structs"
	"strings"
)

type NetworkListener struct {
	url                     string
	contractAddress         common.Address
	crossChainUrl           string
	crossChainTargetAddress string
}

func StartNetworkListening() {
}

func (c *NetworkListener) listenNetwork() {
	client, err := ethclient.Dial(c.url)
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if vLog.Topics[0] == common.HexToHash("0xdb43b563e614fca7a30f9dd93dc8340650747bbd94e07818eaa4de02eafa6fdf") {
				fmt.Println("Topic", vLog.Topics)
				fmt.Println("Data", vLog.Data)
				fmt.Println("Adresss", vLog.Address)
				abiFile, err := ioutil.ReadFile("Bridge.abi")
				contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
				if err != nil {
					fmt.Println(err)
				}
				event := chainListenerEvents.SolidityBridgeEvent{}
				err = contractAbi.UnpackIntoInterface(&event, "BridgeCall", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				// 0=>Wallet Address 1=>Amount 2=>ChainId
				fmt.Println("Event", event)
				c.sendRequestToNearChain(event)
			}
		}
	}
}

func (c *NetworkListener) sendRequestToNearChain(event chainListenerEvents.SolidityBridgeEvent) {

}

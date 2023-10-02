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
	"math/big"
	"networklistener/contracts"
	"strings"
)

type NetworkListener struct {
	url                     string
	contractAddress         common.Address
	crossChainUrl           string
	crossChainTargetAddress string
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
			if vLog.Topics[0] == common.HexToHash("0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e") {
				fmt.Println("Topic", vLog.Topics)
				fmt.Println("Data", vLog.Data)
				fmt.Println("Adresss", vLog.Address)
				abiFile, err := ioutil.ReadFile("Bridge.abi")
				contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
				if err != nil {
					// Handle error
					fmt.Println(err)
				}
				event := struct {
					WalletAddress common.Address
					Amount        string
					ChainID       *big.Int
				}{}
				err = contractAbi.UnpackIntoInterface(&event, "Deposit", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				// 0=>Wallet Address 1=>Amount 2=>ChainId
				fmt.Println("Event", event)
				c.sendRequestToNearChain()
			}
		}
	}
}

func (c *NetworkListener) sendRequestToNearChain() {
	err, contractInstance, _, _, _ := contracts.CreateFunctionRequirementsForBridge(c.crossChainUrl, c.contractAddress.String(), "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(contractInstance)
}

package chainlistener

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	api "github.com/textileio/near-api-go"
	"github.com/textileio/near-api-go/keys"
	"github.com/textileio/near-api-go/types"
)

type NearListener struct {
	nearRpc                 string
	nearWebSocket           string
	secretKey               string
	contractAddress         common.Address
	crossChainUrl           string
	crossChainTargetAddress string
}

func (nl *NearListener) CreateNearClient(network string) *api.Client {

	rpcClient, _ := rpc.DialContext(context.TODO(), nl.nearRpc)

	keyPair, _ := keys.NewKeyPairFromString(
		nl.secretKey,
	)

	config := &types.Config{
		RPCClient: rpcClient,
		Signer:    keyPair,
		NetworkID: network, //testnet - mainnet
	}

	client, _ := api.NewClient(config)

	return client
}

func (nl *NearListener) CallFunction(client *api.Client, accountId string, methodName string) *api.CallFunctionResponse {
	res, err := client.CallFunction(
		context.TODO(),
		accountId,
		methodName,
		api.CallFunctionWithFinality("final"),
	)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res
}

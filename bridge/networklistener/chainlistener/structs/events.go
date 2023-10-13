package chainListenerEvents

import "github.com/ethereum/go-ethereum/common"

type SolidityBridgeEvent struct {
	ContractAddress common.Address
	Message         string
}

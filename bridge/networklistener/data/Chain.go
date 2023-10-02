package bridgeControllerData

type Chain struct {
	ChainName     string `json:"chain_name"`
	Websocket     string `json:"websocket"`
	ChainID       string `json:"chain_id"`
	BridgeAddress string `json:"bridge_address"`
}

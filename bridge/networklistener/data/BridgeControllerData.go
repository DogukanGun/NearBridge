package bridgeControllerData

type CreateBridgeRequest struct {
	NearAddress           string `json:"nearAddress"`
	TargetContractAddress string `json:"targetContractAddress"`
	UserEmail             string `json:"userEmail"`
	TargetURL             string `json:"targetUrl"`
	ChainID               string `json:"chainId"`
}

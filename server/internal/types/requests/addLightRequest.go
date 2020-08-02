package requests

type AddLightRequest struct {
	LightAddrs []string `json:"ipAddresses"`
}

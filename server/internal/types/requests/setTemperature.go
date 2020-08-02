package requests

type SetTemperatureRequest struct {
	*BaseTransitionRequest
	Ct int `json:"ct"`
}

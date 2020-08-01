package requests

type SetBrightnessRequest struct {
	*BaseTransitionRequest
	Brightness int `json:"brigtness"`
}

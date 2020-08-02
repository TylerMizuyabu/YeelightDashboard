package requests

type SetRgbRequest struct {
	*BaseTransitionRequest
	Rgb int `json: "rgb"`
}

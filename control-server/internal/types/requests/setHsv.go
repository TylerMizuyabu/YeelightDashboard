package requests

type SetHsvRequest struct {
	*BaseTransitionRequest
	Hue int `json:"hue"`
	Sat int `json:"sat"`
}

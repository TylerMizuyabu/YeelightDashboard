package requests

type BaseRequest struct {
	LightIds []string `json:"lightIds"`
}

type BaseTransitionRequest struct {
	*BaseRequest
	Smooth   bool `json:"smooth"`
	Duration int  `json:"duration"`
}

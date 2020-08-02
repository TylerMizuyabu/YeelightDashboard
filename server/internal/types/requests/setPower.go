package requests

import "yeelight-server/internal/types"

type SetPowerRequest struct {
	*BaseTransitionRequest
	PowerOn bool            `json:"on"`
	Mode    types.LightMode `json:"mode"`
}

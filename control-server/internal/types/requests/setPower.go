package requests

import "yeelight-control-server/internal/types"

type SetPowerRequest struct {
	*BaseTransitionRequest
	PowerOn bool            `json:"on"`
	Mode    types.LightMode `json:"mode"`
}

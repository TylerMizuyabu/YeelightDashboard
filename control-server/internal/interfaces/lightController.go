package interfaces

import "yeelight-control-server/internal/types"

type ILightController interface {
	UpdateLight(l *types.Yeelight) error
}

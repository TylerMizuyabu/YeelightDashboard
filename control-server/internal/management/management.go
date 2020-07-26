package management

import (
	"fmt"
	"yeelight-control-server/internal/types"
)

type YeelightManager struct {
	lights map[string]*types.Yeelight
}

func NewYeelightManager() *YeelightManager {
	return &YeelightManager{
		lights: make(map[string]*types.Yeelight, 0),
	}
}

func (ym *YeelightManager) Start(discoveredLights chan *types.Yeelight) {
	for light := range discoveredLights {
		if _, ok := ym.lights[light.Id]; ok {
			fmt.Println("Adding received light", light)
			ym.lights[light.Id] = light
		}
	}
}

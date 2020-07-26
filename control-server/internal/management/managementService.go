package management

import (
	"bufio"
	"fmt"
	"net"
	"time"
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
		fmt.Println(light)
		if _, ok := ym.lights[light.Id]; !ok {
			fmt.Println("Adding received light", light)
			ym.lights[light.Id] = light
		}
	}
}

func (ym *YeelightManager) WatchLight(ipAddr string) {
	conn, err := net.DialTimeout("tcp", ipAddr, time.Second*3)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		data, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}
		fmt.Println("Data: ", data)

	}
}

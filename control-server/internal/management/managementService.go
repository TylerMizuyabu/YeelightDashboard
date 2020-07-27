package management

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
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
			go ym.MonitorLight(light.GetAddress(), light.Id)
		}
	}
}

func (ym *YeelightManager) MonitorLight(ipAddr string, id string) {
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
		var nr types.NotificationResponse
		err = json.Unmarshal([]byte(data), &nr)
		if err != nil {
			fmt.Println("Error unmarshaling data", err.Error())
		}
		ym.UpdateLight(ym.lights[id], nr.Params)
		fmt.Println(ym.lights[id])
	}
}

// All of the code bellow should be refactored at some point

func (ym *YeelightManager) UpdateLight(y *types.Yeelight, params types.NotificationResponseParams) {
	// Uuuuuh I'm not sure how to do this better but will have to look into it -_-
	if params.Power != nil {
		if *params.Power == types.On {
			y.IsOn = true
		} else {
			y.IsOn = false
		}
	}

	if params.Brightness != nil {
		y.Brightness = *params.Brightness
	}

	if params.Mode != nil {
		y.Mode = *params.Mode
	}

	if params.Ct != nil {
		y.Ct = *params.Ct
	}

	if params.Rgb != nil {
		y.Rgb = *params.Rgb
	}

	if params.Hue != nil {
		y.Hue = *params.Hue
	}

	if params.Sat != nil {
		y.Sat = *params.Sat
	}

	if params.Name != nil {
		y.Name = *params.Name
	}

	if params.Flowing != nil {
		y.Flowing = *params.Flowing
	}

	if params.FlowParameters != nil {
		flowParams, err := parseFlowParameters(*params.FlowParameters)
		if err != nil {
			fmt.Println("Error updating flow params")
		} else {
			y.FlowParameters = *flowParams
		}
	}
}

func parseFlowParameters(flowParams string) (*types.FlowParams, error) {
	tuples := make([]types.FlowTuple, 0)
	params := strings.Split(flowParams, ",")

	count, err := strconv.Atoi(params[0])
	if err != nil {
		fmt.Println("Error occurred attempting to parse count param")
		return nil, err
	}

	action, err := strconv.Atoi(params[1])
	if err != nil {
		fmt.Println("Error occurred attempting to parse action param")
		return nil, err
	}

	for i := 0; i < len(params[2:])/4; i++ {
		startIndex := 2 + i*4
		endIndex := 6 + i*4
		flowTuple, err := createFlowTuple(params[startIndex:endIndex])
		if err != nil {
			fmt.Println("Error occurred attempting to parse flow tuple")
			return nil, err
		}
		tuples = append(tuples, *flowTuple)
	}

	return &types.FlowParams{
		Count:  uint8(count),
		Action: types.FlowAction(action),
		Tuples: tuples,
	}, nil
}

func createFlowTuple(tuple []string) (*types.FlowTuple, error) {
	duration, err := strconv.Atoi(tuple[0])
	if err != nil {
		return nil, err
	}
	mode, err := strconv.Atoi(tuple[1])
	if err != nil {
		return nil, err
	}
	value, err := strconv.Atoi(tuple[2])
	if err != nil {
		return nil, err
	}
	brightness, err := strconv.Atoi(tuple[3])
	if err != nil {
		return nil, err
	}
	return &types.FlowTuple{
		Duration:   uint64(duration),
		Mode:       uint8(mode),
		Value:      uint64(value),
		Brightness: int8(brightness),
	}, nil
}

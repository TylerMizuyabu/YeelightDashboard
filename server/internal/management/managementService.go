package management

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
	"yeelight-control-server/internal/types"
)

var lightsMutex = &sync.Mutex{}
var lightConnsMutex = &sync.Mutex{}

type YeelightManager struct {
	lights     map[string]*types.Yeelight
	lightConns map[string]net.Conn
}

func NewYeelightManager() *YeelightManager {
	return &YeelightManager{
		lights:     make(map[string]*types.Yeelight, 0),
		lightConns: make(map[string]net.Conn, 0),
	}
}

func (ym *YeelightManager) Start(discoveredLights chan *types.Yeelight) {
	for light := range discoveredLights {
		if _, ok := ym.getLight(light.Id); !ok {
			ym.addLight(light)
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
	ym.addLightConn(id, conn)
	r := bufio.NewReader(conn)
	for {
		data, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}
		var message interface{}
		err = json.Unmarshal([]byte(data), &message)
		if err != nil {
			fmt.Println("Error unmarshaling data", err.Error())
		}
		switch result := message.(type) {
		case types.CommandSuccessResponse:
			fmt.Println("Command success: ", result)
		case types.CommandErrorResponse:
			fmt.Println("Command error: ", result.Error.Message)
		case types.NotificationResponse:
			ym.UpdateLightRecord(ym.lights[id], result.Params)
		}
	}
}

func (ym *YeelightManager) RunCommand(command *types.Command, lightIds []string) string {
	id := 1
	command.SetId(id)
	for _, id := range lightIds {
		if conn, has := ym.getLightConn(id); has {
			if payload, err := json.Marshal(*command); err != nil {
				fmt.Println("Error occurred attempting to parse command json ", command)
			} else {
				payload = append(payload, '\r', '\n')
				fmt.Println("Executing command: ", string(payload))
				conn.Write(payload)
			}
		}
	}

	return fmt.Sprintf("%d", id)
}

// All of the code bellow should be refactored at some point

func (ym *YeelightManager) UpdateLightRecord(y *types.Yeelight, params types.NotificationResponseParams) {
	// Uuuuuh I'm not sure how to do this better but will have to look into it -_-
	ifIntNotNilSetField(&y.Brightness, params.Brightness)
	ifIntNotNilSetField(&y.Ct, params.Ct)
	ifIntNotNilSetField(&y.Rgb, params.Rgb)
	ifIntNotNilSetField(&y.Hue, params.Hue)
	ifIntNotNilSetField(&y.Sat, params.Sat)
	ifStringNotNilSetField(&y.Name, params.Name)

	if params.Power != nil {
		if *params.Power == types.On {
			y.IsOn = true
		} else {
			y.IsOn = false
		}
	}
	if params.Mode != nil {
		y.Mode = *params.Mode
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

func (ym *YeelightManager) addLight(light *types.Yeelight) {
	lightsMutex.Lock()
	ym.lights[light.Id] = light
	lightsMutex.Unlock()
}

func (ym *YeelightManager) getLight(id string) (*types.Yeelight, bool) {
	lightsMutex.Lock()
	light, exists := ym.lights[id]
	lightsMutex.Unlock()
	return light, exists
}

func (ym *YeelightManager) addLightConn(id string, conn net.Conn) {
	lightConnsMutex.Lock()
	ym.lightConns[id] = conn
	lightConnsMutex.Unlock()
}

func (ym *YeelightManager) getLightConn(id string) (net.Conn, bool) {
	lightConnsMutex.Lock()
	conn, exists := ym.lightConns[id]
	lightConnsMutex.Unlock()
	return conn, exists
}

// I want generics T_T

func ifIntNotNilSetField(field *int, value *int) {
	if value != nil {
		*field = *value
	}
}

func ifStringNotNilSetField(field *string, value *string) {
	if value != nil {
		*field = *value
	}
}

func parseFlowParameters(flowParams string) (*types.FlowParams, error) {
	fp := new(types.FlowParams)
	fp.Tuples = make([]types.FlowTuple, 0)
	params := strings.Split(flowParams, ",")

	count, err := strconv.Atoi(params[0])
	if err != nil {
		fmt.Println("Error occurred attempting to parse count param")
		return nil, err
	}
	fp.Count = uint(count)

	action, err := strconv.Atoi(params[1])
	if err != nil {
		fmt.Println("Error occurred attempting to parse action param")
		return nil, err
	}
	fp.Action = types.FlowAction(action)

	for i := 0; i < len(params[2:])/4; i++ {
		startIndex := 2 + i*4
		endIndex := 6 + i*4
		flowTuple, err := createFlowTuple(params[startIndex:endIndex])
		if err != nil {
			fmt.Println("Error occurred attempting to parse flow tuple")
			return nil, err
		}
		fp.Tuples = append(fp.Tuples, *flowTuple)
	}

	return fp, nil
}

func parseIntAndSetField(field *int, str string) error {
	result, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	*field = result
	return nil
}

func createFlowTuple(tuple []string) (*types.FlowTuple, error) {
	ft := new(types.FlowTuple)
	if err := parseIntAndSetField(&ft.Duration, tuple[0]); err != nil {
		return nil, err
	}
	if err := parseIntAndSetField(&ft.Mode, tuple[1]); err != nil {
		return nil, err
	}
	if err := parseIntAndSetField(&ft.Value, tuple[2]); err != nil {
		return nil, err
	}
	if err := parseIntAndSetField(&ft.Brightness, tuple[3]); err != nil {
		return nil, err
	}
	return ft, nil
}

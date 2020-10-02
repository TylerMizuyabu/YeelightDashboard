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
	"yeelight-server/internal/types"

	"github.com/mitchellh/mapstructure"
)

var lightsMutex = &sync.Mutex{}
var lightConnsMutex = &sync.Mutex{}
var requestIdMutex = &sync.Mutex{}

var lightIdField = "lightId"
var getStatusFields = []string{"power", "bright", "ct", "rgb", "hue", "sat", "color_mode", "flowing", "flow_params", "name"}

type YeelightManager struct {
	// the maps and uint should probably be their own custom types that manage their own mutexes
	// the data stored in them would then be private and only exposed via functions that make use of the mutex
	// at the moment if you aren't careful you can still circumvent the mutex. Which kind of defeats the purpose of
	// having them.
	lights     map[string]*types.Yeelight
	lightConns map[string]net.Conn
	broadcast  chan []byte
	requestId  uint
}

func NewYeelightManager(broadcast chan []byte) *YeelightManager {
	return &YeelightManager{
		lights:     make(map[string]*types.Yeelight, 0),
		lightConns: make(map[string]net.Conn, 0),
		broadcast:  broadcast,
		requestId:  0,
	}
}

func (ym *YeelightManager) Start(discoveredLights chan *types.Yeelight) {
	for light := range discoveredLights {
		if ok := ym.AddLight(light); !ok {
			go ym.MonitorLight(light.GetAddress(), light.Id)
		}
	}
}

//This method can be broken up into simpler smaller parts

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
		message := make(map[string]interface{}, 0)
		err = json.Unmarshal([]byte(data), &message)
		if err != nil {
			fmt.Println("Error unmarshaling data", err.Error())
		}
		//Look into a better way to determine type dynamically
		if _, exists := message["result"]; exists {
			var successResponse types.CommandSuccessResponse
			if err := mapstructure.Decode(message, &successResponse); err != nil {
				fmt.Println("Error occurred attempting to convert message into CommandSuccessResponse struct", err)
				continue
			}
			if len(successResponse.Result) > 1 {
				// its a get_prop command.
				var results map[string]interface{}
				for i, field := range getStatusFields {
					results[field] = successResponse.Result[i]
				}
				var params types.NotificationResponseParams
				if err := mapstructure.Decode(results, &params); err != nil {
					fmt.Println("Error occurred attempting to convert get_props result into NotificationResponseParams struct", err)
				} else {
					ym.UpdateLight(id, params)
				}
				message[lightIdField] = id
				message["params"] = params
				messageStr, err := json.Marshal(message)
				if err != nil {
					fmt.Println("Error converting notification message to string", err)
					continue
				}
				ym.broadcast <- messageStr
			}
		} else if _, exists := message["error"]; exists {
			fmt.Println("Command error: ", data)
		} else {
			var notification types.NotificationResponse
			if err := mapstructure.Decode(message, &notification); err != nil {
				fmt.Println("Error occurred attempting to convert message into NotificationResponse struct", err)
				continue
			}
			ym.UpdateLight(id, notification.Params)
			message[lightIdField] = id
			messageStr, err := json.Marshal(message)
			if err != nil {
				fmt.Println("Error converting notification message to string", err)
				continue
			}
			ym.broadcast <- messageStr
		}
	}
}

func (ym *YeelightManager) FetchLightStatus(lightIds []string) string {
	var command = NewGetPropsCommand(getStatusFields...)
	return ym.RunCommand(command, lightIds)
}

func (ym *YeelightManager) RunCommand(command *types.Command, lightIds []string) string {
	requestId := ym.getRequestId()
	command.SetId(requestId)
	for _, lightId := range lightIds {
		if conn, has := ym.getLightConn(lightId); has {
			if payload, err := json.Marshal(*command); err != nil {
				fmt.Println("Error occurred attempting to parse command json ", command)
			} else {
				payload = append(payload, '\r', '\n')
				fmt.Println("Executing command: ", string(payload))
				conn.Write(payload)
			}
		}
	}

	return fmt.Sprintf("%d", requestId)
}

// All of the code bellow should be refactored at some point

func (ym *YeelightManager) UpdateLight(id string, params types.NotificationResponseParams) {
	lightsMutex.Lock()
	defer lightsMutex.Unlock()
	y := ym.lights[id]
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

func (ym *YeelightManager) AddLight(light *types.Yeelight) bool {
	lightsMutex.Lock()
	defer lightsMutex.Unlock()
	fmt.Println("Adding light ", light)
	if _, ok := ym.lights[light.Id]; !ok {
		ym.lights[light.Id] = light
		return ok
	}
	return true
}

func (ym *YeelightManager) GetLights() []*types.Yeelight {
	lights := make([]*types.Yeelight, 0)
	for _, light := range ym.lights {
		lights = append(lights, light)
	}
	return lights
}

func (ym *YeelightManager) getLight(id string) (*types.Yeelight, bool) {
	lightsMutex.Lock()
	defer lightsMutex.Unlock()
	light, exists := ym.lights[id]
	return light, exists
}

func (ym *YeelightManager) addLightConn(id string, conn net.Conn) {
	lightConnsMutex.Lock()
	defer lightConnsMutex.Unlock()
	ym.lightConns[id] = conn
}

func (ym *YeelightManager) getLightConn(id string) (net.Conn, bool) {
	lightConnsMutex.Lock()
	defer lightConnsMutex.Unlock()
	conn, exists := ym.lightConns[id]
	return conn, exists
}

func (ym *YeelightManager) getRequestId() uint {
	requestIdMutex.Lock()
	defer requestIdMutex.Unlock()
	var returnValue = ym.requestId
	ym.requestId++
	return returnValue
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

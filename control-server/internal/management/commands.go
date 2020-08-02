package management

import (
	"fmt"
	"yeelight-control-server/internal/types"
)

var minDuration int = 30
var maxRgb int = 16777215
var minRgb int = 0
var maxHue int = 359
var minHue int = 0
var minTemp int = 1700
var maxTemp int = 6500
var minBrightness int = 1
var maxBrightness int = 100

func NewGetPropsCommand(props ...string) *types.Command {
	return types.NewCommand(0, "get_prop", []interface{}{props})
}

func NewSetCtAbxCommand(temp int, smooth bool, duration int) *types.Command {
	return types.NewCommand(0, "set_ct_abx", []interface{}{withinRange(temp, &minTemp, &maxTemp), smoothOrSudden(smooth), withinRange(duration, &minDuration, nil)})
}

func NewSetRgbCommand(rgb int, smooth bool, duration int) *types.Command {
	return types.NewCommand(0, "set_rgb", []interface{}{withinRange(rgb, &minRgb, &maxRgb), smoothOrSudden(smooth), withinRange(duration, &minDuration, nil)})
}

func NewSetHsvCommand(hue int, sat int, smooth bool, duration int) *types.Command {
	return types.NewCommand(0, "set_hsv", []interface{}{withinRange(hue, &minHue, &maxHue), sat, smoothOrSudden(smooth), withinRange(duration, &minDuration, nil)})
}

func NewSetBrightnessCommand(brightness int, smooth bool, duration int) *types.Command {
	return types.NewCommand(0, "set_bright", []interface{}{withinRange(brightness, &minBrightness, &maxBrightness), smoothOrSudden(smooth), withinRange(duration, &minDuration, nil)})
}

func NewSetPowerCommand(on bool, smooth bool, duration int, mode *types.LightMode) *types.Command {
	if mode == nil {
		mode = new(types.LightMode)
		*mode = types.DefaultLightMode
	}
	return types.NewCommand(0, "set_power", []interface{}{onOrOff(on), smoothOrSudden(smooth), withinRange(duration, &minDuration, nil), mode})
}

func NewSetDefaultCommand() *types.Command {
	return types.NewCommand(0, "set_default", []interface{}{})
}

func NewStartColorFlowCommand(params types.FlowParams) *types.Command {
	return types.NewCommand(0, "start_cf", []interface{}{params.Count, params.Action, flowTupleSliceToString(&params.Tuples)})
}

func NewStopColorFlowCommand() *types.Command {
	return types.NewCommand(0, "stop_cf", []interface{}{})
}

func smoothOrSudden(isSmooth bool) string {
	if isSmooth {
		return "smooth"
	} else {
		return "sudden"
	}
}

func onOrOff(isOn bool) string {
	if isOn {
		return types.On
	}
	return types.Off
}

func withinRange(value int, min *int, max *int) int {
	if min != nil && value < *min {
		return *min
	} else if  max != nil && value > *max {
		return *max
	} else {
		return value
	}
}

func flowTupleSliceToString(ft *[]types.FlowTuple) string {
	str := ""
	for _, t := range *ft {
		// TODO: ToParamsString could benefit from functions such as withinRange()
		if len(str) == 0 {
			str = t.ToParamsString()
		} else {
			str = fmt.Sprintf("%s,%s", str, t.ToParamsString())
		}
	}
	return str
}

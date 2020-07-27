package types

import "fmt"

type FlowParams struct {
	Count  uint8
	Action FlowAction
	Tuples []FlowTuple
}

type FlowTuple struct {
	Duration   uint64
	Mode       uint8
	Value      uint64
	Brightness int8
}

type FlowAction uint8

const (
	RecoverToPreviousState FlowAction = iota
	KeepCurrentState
	TurnOff
)

func (fp *FlowParams) ToParamsArray() []interface{} {
	arr := []interface{}{fp.Count, fp.Action}
	flowTuplesString := ""
	for _, tuple := range fp.Tuples {
		flowTuplesString = fmt.Sprintf("%s,%s", flowTuplesString, tuple.ToParamsString())
	}
	return append(arr, flowTuplesString)
}

func (ft *FlowTuple) ToParamsString() string {
	return fmt.Sprintf("%d,%d,%d,%d", ft.Duration, ft.Mode, ft.Value, ft.Brightness)
}

package types

import "fmt"

type FlowParams struct {
	Count  uint
	Action FlowAction
	Tuples []FlowTuple
}

type FlowTuple struct {
	Duration   int
	Mode       int
	Value      int
	Brightness int
}

type FlowAction int

const (
	RecoverToPreviousState int = iota
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

package types

import (
	"errors"
	"fmt"
)

var ErrorInvalidResponseMessage error = errors.New("Invalid Response Message")

type Yeelight struct {
	addr       string
	Id         string
	Model      LightModel
	IsOn       bool
	Brightness uint8
	Mode       LightMode
	Ct         uint64
	Rgb        uint64
	Hue        uint16
	Sat        uint8
	Name       string
}

type LightModel string

const (
	Mono    LightModel = "mono"
	Color              = "color"
	Stripe             = "stripe"
	Ceiling            = "ceiling"
	BsLamp             = "bslamp"
)

type LightMode uint8

const (
	ColorMode            LightMode = 1
	ColorTemperatureMode           = 2
	HSVMode                        = 3
)

func NewYeelight(responseMessage string) (*Yeelight, error) {
	parser := NewParser(responseMessage)
	errs := make([]error, 0)
	y := new(Yeelight)
	handleParserError(parser.ParseAddr(&y.addr), &errs)
	handleParserError(parser.ParseHeader("ID", &y.Id), &errs)
	handleParserError(parser.ParseHeader("NAME", &y.Name))
	if len(errs) > 0 {
		logParserErrors(&errs)
		return nil, ErrorInvalidResponseMessage
	}
	return y, nil
}

func handleParserError(err error, errs *[]error) {
	if err != nil {
		*errs = append(*errs, err)
	}
}

func logParserErrors(errs *[]error) {
	fmt.Println("The following errors occurred while parsing the response message:")
	for _, err := range *errs {
		fmt.Println("\t", err.Error())
	}
}

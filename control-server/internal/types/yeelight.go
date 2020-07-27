package types

import (
	"errors"
	"fmt"
)

var ErrorInvalidResponseMessage error = errors.New("Invalid Response Message")
var ErrorUnableToParseNotificationMessage error = errors.New("Unable To Parse Yeelight Notification Message")

type Yeelight struct {
	addr           string
	Id             string
	Model          LightModel
	IsOn           bool
	Brightness     uint8
	Mode           LightMode
	Ct             uint64
	Rgb            uint64
	Hue            uint16
	Sat            uint8
	Name           string
	Flowing        FlowMode
	FlowParameters FlowParams
}

type LightModel string

const (
	Mono    LightModel = "mono"
	Color              = "color"
	Stripe             = "stripe"
	Ceiling            = "ceiling"
	BsLamp             = "bslamp"
)

func NewYeelightFromDiscoveryResponse(responseMessage string) (*Yeelight, error) {
	parser := NewParser(responseMessage)
	errs := make([]error, 0)
	y := new(Yeelight)
	handleParserError(parser.ParseAddr(&y.addr), &errs)
	handleParserError(parser.ParseHeader("id", false, &y.Id), &errs)
	handleParserError(parser.ParseHeader("name", true, &y.Name), &errs)
	if len(errs) > 0 {
		logErrors(&errs)
		return nil, ErrorInvalidResponseMessage
	}
	return y, nil
}

func (y *Yeelight) GetAddress() string {
	return y.addr
}

func handleParserError(err error, errs *[]error) {
	if err != nil {
		*errs = append(*errs, err)
	}
}

func logErrors(errs *[]error) {
	fmt.Println("The following errors occurred while parsing the response message:")
	for _, err := range *errs {
		fmt.Println("\t", err.Error())
	}
}

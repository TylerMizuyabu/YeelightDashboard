package types

type Yeelight struct {
	addr string
	Id string
	Model LightModel
	IsOn bool
	Brightness uint8
	Mode LightMode
	Ct uint64
	Rgb uint64
	Hue uint16
	Sat uint8
	Name string

}

type LightModel string

const (
	Mono LightModel= "mono"
	Color = "color"
	Stripe = "stripe"
	Ceiling = "ceiling"
	BsLamp = "bslamp"
)

type LightMode uint8

const (
	ColorMode LightMode = 1
	ColorTemperatureMode = 2
	HSVMode = 3
)

package types

type LightMode uint8

const (
	DefaultLightMode LightMode = iota
	ColorMode
	ColorTemperatureMode
	HSVMode
	ColorFlowMode
	NightLightMode
)

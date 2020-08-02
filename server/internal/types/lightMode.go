package types

type LightMode uint8

var (
	DefaultLightMode     LightMode = 0
	ColorMode                      = 1
	ColorTemperatureMode           = 2
	HSVMode                        = 3
	ColorFlowMode                  = 4
	NightLightMode                 = 5
)

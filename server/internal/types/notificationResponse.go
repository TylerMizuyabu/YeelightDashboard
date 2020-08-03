package types

type NotificationResponse struct {
	Method string                     `json:"method"`
	Params NotificationResponseParams `json:"params"`
}

type NotificationResponseParams struct {
	Power          *string    `json:"power"`
	Brightness     *int       `json:"bright" mapstructure:"bright"`
	Mode           *LightMode `json:"color_mode" mapstructure:"color_mode"`
	Ct             *int       `json:"ct"`
	Rgb            *int       `json:"rgb"`
	Hue            *int       `json:"hue"`
	Sat            *int       `json:"sat"`
	Name           *string    `json:"name"`
	Flowing        *FlowMode  `json:"flowing"`
	FlowParameters *string    `json:"flow_params" mapstructure:"flow_params"`
	// I'm not bothering with the other properties... for now
}

const (
	On  string = "on"
	Off        = "off"
)

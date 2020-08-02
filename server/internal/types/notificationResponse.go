package types

type NotificationResponse struct {
	Method string                     `json:"method"`
	Params NotificationResponseParams `json:"params"`
}

type NotificationResponseParams struct {
	Power          *string    `json:"power,omitempty"`
	Brightness     *int       `json:"bright,omitempty"`
	Mode           *LightMode `json:"color_mode,omitempty"`
	Ct             *int       `json:"ct,omitempty"`
	Rgb            *int       `json:"rgb,omitempty"`
	Hue            *int       `json:"hue,omitempty"`
	Sat            *int       `json:"sat,omitempty"`
	Name           *string    `json:"name,omitempty"`
	Flowing        *FlowMode  `json:"flowing,omitempty"`
	FlowParameters *string    `json:"flow_params,omitempty"`
	// I'm not bothering with the other properties... for now
}

const (
	On  string = "on"
	Off        = "off"
)

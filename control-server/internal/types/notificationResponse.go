package types

type NotificationResponse struct {
	Method string                     `json:"method"`
	Params NotificationResponseParams `json:"params"`
}

type NotificationResponseParams struct {
	Power      string    `json:"power,omitempty"`
	Brightness uint8     `json:"bright,omitempty"`
	Mode       LightMode `json:"color_mode,omitempty"`
	Ct         uint64    `json:"ct,omitempty"`
	Rgb        uint64    `json:"rgb,omitempty"`
	Hue        uint16    `json:"hue,omitempty"`
	Sat        uint8     `json:"sat,omitempty"`
	Name       string    `json:"name,omitempty"`
	Flowing    FlowMode  `json:"flowing,omitempty"`
	FlowParams string    `json:"flow_params,omitempty"`
	DelayOff   uint8     `json:"delayoff,omitempty"`
}

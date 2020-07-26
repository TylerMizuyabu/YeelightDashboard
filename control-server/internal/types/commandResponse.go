package types

// Need to look more into the actual return value of this

type CommandResponse struct {
	id     int
	result []string
	error  map[string]interface{}
}

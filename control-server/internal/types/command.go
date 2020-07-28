package types

type Command struct {
	Id     string
	Method string
	Params []interface{}
}

func NewCommand(id string, method string, params []interface{}) *Command {
	return &Command{
		Id:     id,
		Method: method,
		Params: params,
	}
}

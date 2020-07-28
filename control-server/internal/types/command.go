package types

type Command struct {
	Id     string        `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

func NewCommand(id string, method string, params []interface{}) *Command {
	return &Command{
		Id:     id,
		Method: method,
		Params: params,
	}
}

func (c *Command) SetId(id string) {
	c.Id = id
}

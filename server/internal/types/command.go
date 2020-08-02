package types

type Command struct {
	Id     int           `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

func NewCommand(id int, method string, params []interface{}) *Command {
	return &Command{
		Id:     id,
		Method: method,
		Params: params,
	}
}

func (c *Command) SetId(id int) {
	c.Id = id
}

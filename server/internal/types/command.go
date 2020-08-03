package types

type Command struct {
	Id     uint          `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

func NewCommand(id uint, method string, params []interface{}) *Command {
	return &Command{
		Id:     id,
		Method: method,
		Params: params,
	}
}

func (c *Command) SetId(id uint) {
	c.Id = id
}

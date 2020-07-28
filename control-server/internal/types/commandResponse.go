package types

type CommandSuccessResponse struct {
	Id     int           `json:"id"`
	Result []interface{} `json:"result"`
}

type CommandErrorResponse struct {
	Id    int `json:"id"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

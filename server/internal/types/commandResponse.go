package types

type CommandSuccessResponse struct {
	Id     uint          `json:"id"`
	Result []interface{} `json:"result"`
}

type CommandErrorResponse struct {
	Id    uint `json:"id"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

package types

import (
	"encoding/json"
	"errors"
)

var ErrorMessageNotRecognized = errors.New("Unable to parse incomming light message")

type LightMessageUnpacker struct {
	Data interface{}
}

func (u *LightMessageUnpacker) UnmarshalJSON(b []byte) error {
	cmdSuccess := new(CommandSuccessResponse)
	if err := json.Unmarshal(b, cmdSuccess); err == nil {
		u.Data = cmdSuccess
		return nil
	}
	cmdError := new(CommandErrorResponse)
	if err := json.Unmarshal(b, cmdError); err == nil {
		u.Data = cmdError
		return nil
	}
	notification := new(NotificationResponse)
	if err := json.Unmarshal(b, notification); err == nil {
		u.Data = notification
		return nil
	}

	return ErrorMessageNotRecognized
}

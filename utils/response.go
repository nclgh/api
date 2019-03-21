package utils

import "github.com/sirupsen/logrus"

type Response struct {
	Success bool                   `json:"success"`
	Code    int                    `json:"error_code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewCommonResponse(code int) *Response {
	msg, ok := msgMap[code]
	if ok != true {
		logrus.Errorf("code not installed code:%d", code)
		return nil
	}
	rsp := &Response{
		Success: code == CodeSuccess,
		Code:    code,
		Message: msg,
		Data:    make(map[string]interface{}),
	}
	return rsp
}

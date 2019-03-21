package utils

import "github.com/sirupsen/logrus"

type Response struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func NewCommonResponse(code int) *Response {
	msg, ok := msgMap[code]
	if ok != true {
		logrus.Errorf("code not installed code:%d", code)
		return nil
	}
	rsp := &Response{
		Code: code,
		Msg:  msg,
		Data: make(map[string]interface{}),
	}
	return rsp
}

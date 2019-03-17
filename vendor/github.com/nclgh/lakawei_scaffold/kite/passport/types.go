package passport

import (
	"encoding/json"
	"github.com/nclgh/lakawei_scaffold/kite/kite_common"
)

type CreateSessionRequest struct {
	UserId int64 `json:"user_id"`
}

type CreateSessionResponse struct {
	SessionId string              `json:"session_id"`
	Code      kite_common.RspCode `json:"code"`
	Msg       string              `json:"msg"`
}

func UnmarshalCreateSessionRequest(r string) (req *CreateSessionRequest, err error) {
	if len(r) < 2 {
		return nil, kite_common.ErrReq
	}
	r = r[1 : len(r)-1]
	req = &CreateSessionRequest{}
	err = json.Unmarshal([]byte(r), req)
	return req, err
}

func GetCreateSessionResponse(code kite_common.RspCode, msg string) *CreateSessionResponse {
	rsp := &CreateSessionResponse{
		Code: code,
		Msg:  msg,
	}
	if msg == "" {
		dMsg := kite_common.UnknownCodeMsg
		m, ok := kite_common.CodeMsg[code]
		if ok {
			dMsg = m
		}
		rsp.Msg = dMsg
	}
	return rsp
}

type GetSessionRequest struct {
	SessionId string `json:"session_id"`
}

type GetSessionResponse struct {
	UserId int64               `json:"user_id"`
	Code   kite_common.RspCode `json:"code"`
	Msg    string              `json:"msg"`
}

func UnmarshalGetSessionRequest(r string) (req *GetSessionRequest, err error) {
	if len(r) < 2 {
		return nil, kite_common.ErrReq
	}
	r = r[1 : len(r)-1]
	req = &GetSessionRequest{}
	err = json.Unmarshal([]byte(r), req)
	return req, err
}

func GetGetSessionResponse(code kite_common.RspCode, msg string) *GetSessionResponse {
	rsp := &GetSessionResponse{
		Code: code,
		Msg:  msg,
	}
	if msg == "" {
		dMsg := kite_common.UnknownCodeMsg
		m, ok := kite_common.CodeMsg[code]
		if ok {
			dMsg = m
		}
		rsp.Msg = dMsg
	}
	return rsp
}

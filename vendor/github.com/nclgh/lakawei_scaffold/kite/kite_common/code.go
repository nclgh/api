package kite_common

type RspCode int64

const (
	CodeSuccess RspCode = 0
	CodeReqErr  RspCode = 1
)

var (
	UnknownCodeMsg = "unknown code"

	CodeMsg = map[RspCode]string{
		CodeSuccess: "success",
		CodeReqErr:  "parser req err",
	}
)

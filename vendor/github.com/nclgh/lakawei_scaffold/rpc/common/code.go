package common

type RspCode int64

const (
	CodeSuccess RspCode = 0
	CodeFailed  RspCode = 500
	CodeReqErr  RspCode = 1
)

var (
	UnknownCodeMsg = "unknown code"

	CodeMsg = map[RspCode]string{
		CodeSuccess: "success",
		CodeFailed:  "failed",
		CodeReqErr:  "parser req err",
	}
)

package utils

var (
	CodeSuccess  = 0
	CodePARAMERR = 1

	CodeFailed       = 4
	CodeLoginRequire = 5
)

var (
	msgMap = map[int]string{
		CodeSuccess:      "Success",
		CodePARAMERR:     "参数出现错误",
		CodeFailed:       "failed",
		CodeLoginRequire: "Login Required",
	}
)

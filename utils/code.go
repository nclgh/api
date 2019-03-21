package utils

var (
	CodeSuccess      = 0
	CodeFailed       = 1
	CodePARAMERR     = 2
	CodeLoginRequire = 3
)

var (
	msgMap = map[int]string{
		CodeSuccess:      "Success",
		CodePARAMERR:     "参数出现错误",
		CodeFailed:       "failed",
		CodeLoginRequire: "Login Required",
	}
)

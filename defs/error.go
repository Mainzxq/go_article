package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"` // 系统内的err非webstatus
}

type ErrorRespones struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorRespones{HttpSC: 400, Error: Err{Error:"Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser = ErrorRespones{HttpSC: 401, Error:Err{Error:"user authentication failed", ErrorCode:"002"}}
)

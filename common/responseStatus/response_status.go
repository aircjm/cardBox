package responseStatus

const (
	SUCCESS                        = 0
	SYSTEM_ERROR                   = 9001
	ERROR_PARAMS_ERROR             = 9002
	ERROR_LOGIN                    = 9004
	INVALID_PARAMS                 = 9005
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 9006
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 9007
	ERROR_AUTH_TOKEN               = 9008
)

var statusMsg = map[int]string{
	SUCCESS:                        "success",
	SYSTEM_ERROR:                   "system error",
	ERROR_PARAMS_ERROR:             "param error",
	ERROR_LOGIN:                    "login error",
	INVALID_PARAMS:                 "invalid param",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "error auth check token timeout",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "error auth check token fail",
	ERROR_AUTH_TOKEN:               "error auth token",
}

// 获取当前status的msg内容
func GetStatusMsg(code int) string {
	msg, ok := statusMsg[code]
	if ok {
		return msg
	}
	return statusMsg[SYSTEM_ERROR]
}

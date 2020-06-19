package common

const (
	Success          = 0
	SystemError      = 1
	ErrorParamsError = 2
	ErrorLogin       = 4
)

var statusMsg = map[int]string{
	Success:          "success",
	SystemError:      "system error",
	ErrorParamsError: "param error",
	ErrorLogin:       "login error",
}

// 获取当前status的msg内容
func GetStatusMsg(code int) string {
	msg, ok := statusMsg[code]
	if ok {
		return msg
	}
	return statusMsg[SystemError]
}

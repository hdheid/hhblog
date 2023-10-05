package common

type ErrorCode int

//封装一下错误码

const (
	SettingsErr = 1001 //系统错误
)

var (
	ErrMap = map[ErrorCode]string{
		SettingsErr: "系统错误",
	}
)

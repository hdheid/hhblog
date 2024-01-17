package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 // qq登录
	SignGitee SignStatus = 2 // 码云登录
	SignEmail SignStatus = 3 // 邮箱登录
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	var str string
	switch s {
	case SignQQ:
		str = "qq"
	case SignGitee:
		str = "码云"
	case SignEmail:
		str = "邮箱"
	default:
		str = "其他"
	}
	return json.Marshal(str)
}

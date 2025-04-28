package socketMsg

import (
	"im/global"
	"im/utils"
)

// MsgStr 加密结构
type MsgStr string

type MsgStrT interface {
	string | []byte
}

func NewMsgStr[T MsgStrT](str T) MsgStr {
	switch v := any(str).(type) {
	case string:
		s, _ := utils.AesBase64Encrypt([]byte(v), global.G_CONFIG.Gin.WebSocketKey)
		return MsgStr(s)
	case []byte:
		s, _ := utils.AesBase64Encrypt(v, global.G_CONFIG.Gin.WebSocketKey)
		return MsgStr(s)
	}
	return ""
}

func (m MsgStr) String() string {
	s, _ := utils.AesBase64Decrypt(string(m), global.G_CONFIG.Gin.WebSocketKey)
	return string(s)
}

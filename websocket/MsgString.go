package socketMsg

import (
	"encoding/json"
	"im/global"
	"im/utils"
	"time"
)

// MsgStr 加密结构
type MsgStr string

type MsgStrT interface {
	string | []byte | any
}

func NewMsgStr[T MsgStrT](str T) (MsgStr, error) {
	switch v := any(str).(type) {
	case string:
		s, err := utils.AesBase64Encrypt([]byte(v), global.G_CONFIG.Gin.WebSocketKey)
		return MsgStr(s), err
	case []byte:
		s, err := utils.AesBase64Encrypt(v, global.G_CONFIG.Gin.WebSocketKey)
		return MsgStr(s), err
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		s, err := utils.AesBase64Encrypt(b, global.G_CONFIG.Gin.WebSocketKey)
		if err != nil {
			return "", err
		}
		return MsgStr(s), nil
	}
}

type TypeCode2Any interface {
	NotifyConst | GroupConst | FriendConst | ErrConst
}

func NewMstT[TypeCode2 TypeCode2Any](typeCode1 MsgConst, typeCode2 TypeCode2, d TMsgBase, msg any) (TMsg, error) {
	m, err := NewMsgStr(msg)
	if err != nil {
		return TMsg{}, err
	}
	ret := TMsg{
		TMsgBase:  d,
		TypeCode:  typeCode1,
		TypeCode2: uint(typeCode2),
		T:         time.Now(),
		Msg:       m,
	}
	return ret, nil
}

func NewErr(msg TMsg, code ErrConst, sendmsg any) (TMsg, error) {
	return NewMstT(
		MsgErr,
		code,
		TMsgBase{
			ToID: msg.FromID,
		},
		Err{
			LastMsg: msg,
			Msg:     sendmsg,
		},
	)
}
func (m MsgStr) String() string {
	s, _ := utils.AesBase64Decrypt(string(m), global.G_CONFIG.Gin.WebSocketKey)
	return string(s)
}

func (m MsgStr) ToFriendCall() (ret FriendCall, err error) {
	s := []byte(m.String())
	err = json.Unmarshal(s, &ret)
	return
}

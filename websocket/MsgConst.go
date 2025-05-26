package socketMsg

import "time"

// MsgConst 消息类型常量
type MsgConst uint

// 消息类型常量
const (
	MsgErr    MsgConst = iota //ping消息
	MsgFriend                 //好友消息 FriendConst
	MsgGroup                  //群组消息 GroupConst
	MsgNotify                 //通知 添加好友请求 群邀请等 NotifyConst
	MsgProxy                  //消息转发
)

/*
	{
		"typecode": 0,
		"typecode2": 0,
		"t": 0,
		"userid": 0,
		"toid": 0,
		"id": 0,
		"msg": "string"
	}
*/
type TMsgBase struct {
	FromID  uint `json:"fromid,omitempty"`  //来自
	ToID    uint `json:"toid,omitempty"`    //发送到
	GroupID uint `json:"gourpid,omitempty"` //相关群组ID
	ID      uint `json:"id,omitempty"`      //消息ID
}

// TMsg 消息结构体
type TMsg struct {
	TMsgBase
	TypeCode  MsgConst  `json:"typecode"`      //消息类型
	TypeCode2 uint      `json:"typecode2"`     //消息类型 FriendConst GroupConst ErrConst NotifyConst
	T         time.Time `json:"t"`             //时间
	Msg       MsgStr    `json:"msg,omitempty"` //具体消息 aes加密后的base64
}

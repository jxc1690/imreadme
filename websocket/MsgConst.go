package socketMsg

import "time"

type MsgConst uint

const (
	MsgErr    MsgConst = iota //ping消息
	MsgFriend                 //好友消息
	MsgGroup                  //群组消息
	MsgNotify                 //通知 添加好友请求 群邀请等
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
type TMsg struct {
	TypeCode  MsgConst  `json:"typecode,omitempty"`  //消息类型
	TypeCode2 uint      `json:"typecode2,omitempty"` //消息类型2
	T         time.Time `json:"t"`                   //时间
	UserID    uint      `json:"userid,omitempty"`    //来自
	ToID      uint      `json:"toid,omitempty"`      //发送到
	ID        uint      `json:"id,omitempty"`        //消息ID
	Msg       string    `json:"msg,omitempty"`       //具体消息 aes加密后的base64
}

type ProxyConst uint

const (
	ProxyFriend ProxyConst = iota //好友消息
	ProxyGroup                    //群组消息
)

type Proxy struct {
	TypeCode MsgConst `json:"typecode,omitempty"`
	Msg      []TMsg   `json:"msgs"`
}

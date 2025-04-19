package socketMsg

import "time"

type MsgConst uint

const (
	MsgErr    MsgConst = iota //ping消息
	MsgFriend                 //好友消息
	MsgGroup                  //群组消息
	MsgNotify                 //通知 添加好友请求 群邀请等
)

type TMsg struct {
	TypeCode MsgConst  `json:"typecode,omitempty"` //消息类型
	T        time.Time `json:"t"`                  //时间
	Msg      string    `json:"msg,omitempty"`      //具体消息 aes加密后的base64
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

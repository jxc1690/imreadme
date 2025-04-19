package socketMsg

type NotifyConst uint

const (
	NotifyFriend  NotifyConst = iota //添加好友请求 msg NotifyADD
	NotifyGroup                      //群邀请 msg NotifyADD
	NotifyGConfig                    //全局设置变更 收到该消息之后需要重新拉取设置 msg nil
)

type Notify struct {
	Type uint `json:"type"`          //消息类型
	Msg  any  `json:"msg,omitempty"` //消息内容
}

type NotifyADD struct {
	Msg     string `json:"msg"`               //消息内容
	Form    uint   `json:"form"`              //消息来源
	GroupID uint   `json:"groupID,omitempty"` //邀请加群ID
	To      uint   `json:"to"`                //消息发送至
}

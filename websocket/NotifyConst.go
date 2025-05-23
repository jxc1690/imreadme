package socketMsg

type NotifyConst uint

const (
	NotifyFriend   NotifyConst = iota //添加好友请求 msg NotifyADD
	NotifyGroup                       //群邀请 msg NotifyADD
	NotifyGroupAdd                    //群申请 msg NotifyADD
	NotifyGConfig                     //全局设置变更 收到该消息之后需要重新拉取设置 msg nil
	NotifyPass                        //通过 msg NotifyPassT
)

type NotifyADD struct {
	Msg     string `json:"msg"`               //消息内容
	GroupID uint   `json:"groupID,omitempty"` //邀请加群ID
}

type FGConst uint

const (
	FGFriend FGConst = iota
	FGGroup
)

type NotifyPassT struct {
	ID   uint    //model.FriendDelay model.GroupDelay
	Type FGConst //friend or group
}

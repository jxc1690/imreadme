package socketMsg

type FriendConst uint

const (
	FriendStr    FriendConst = iota //文本 FriendMsg
	FriendMis                       //音频 FriendMsg
	FriendImg                       //图片 FriendMsg
	FriendVideo                     //视频 FriendMsg
	FriendProxy                     //转发消息 FriendProxyMsg
	FriendRet                       //撤回 FriendRetMsg
	FriendFriend                    //好友名片 FriendFriendMsg
	FriendNotify                    //通知 FriendMsg
)

type Friend struct {
	Msg   string `json:"msg,omitempty"`   //消息内容
	Proxy []TMsg `json:"proxy,omitempty"` //转发消息``
	Ret   uint   `json:"ret,omitempty"`   //撤回
	Other any    `json:"other,omitempty"` //其它 前端自定义
}

type FriendMsg struct {
	Msg   string `json:"msg"`             //消息内容
	Other any    `json:"other,omitempty"` //其它 前端自定义
}

type FriendRetMsg struct {
	Ret   uint `json:"ret"`             //撤回
	Other any  `json:"other,omitempty"` //其它 前端自定义
}

type FriendProxyMsg struct {
	Proxy []TMsg `json:"proxy"`           //转发消息``
	Other any    `json:"other,omitempty"` //其它 前端自定义
}

type FriendFriendMsg struct {
	UserID uint `json:"user_id"`         //用户ID
	Other  any  `json:"other,omitempty"` //其它 前端自定义 (带上头像 名称)
}

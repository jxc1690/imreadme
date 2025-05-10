package websocket

type FriendConst uint

const (
	FriendStr   FriendConst = iota //文本
	FriendMis                      //音频
	FriendImg                      //图片
	FriendVideo                    //视频
	FriendProxy                    //转发消息
	FriendRet                      //撤回
)

type Friend struct {
	Msg   string `json:"msg,omitempty"`   //消息内容
	Proxy Proxy  `json:"proxy,omitempty"` //转发消息``
	Ret   uint   `json:"ret,omitempty"`   //撤回
}

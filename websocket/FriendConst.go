package socketMsg

type FriendConst uint

const (
	FriendStr   FriendConst = iota //文本 msg String
	FriendMis                      //音频 msg String
	FriendImg                      //图片 msg String
	FriendVideo                    //视频 msg String
	FriendProxy                    //转发消息 msg Proxy
	FriendRet                      //撤回 msg uint
)

type Friend struct {
	Form uint `json:"form"` //来自
	To   uint `json:"to"`   //发送到
	Msg  any  `json:"msg"`  //消息内容
}

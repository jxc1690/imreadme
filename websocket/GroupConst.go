package socketMsg

type GroupConst uint

const (
	GroupStr   GroupConst = iota //文本 msg String
	GroupMis                     //音频 msg String
	GroupImg                     //图片 msg String
	GroupVideo                   //视频 msg String
	GroupProxy                   //转发消息 msg Proxy
)

type Group struct {
	Type GroupConst `json:"type"` //消息类型
	Form uint       `json:"form"` //来自
	To   uint       `json:"to"`   //发送到
	Msg  any        `json:"msg"`  //消息内容
}

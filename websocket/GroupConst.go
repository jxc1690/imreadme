package websocket

type GroupConst uint

const (
	GroupStr   GroupConst = iota //文本
	GroupMis                     //音频
	GroupImg                     //图片
	GroupVideo                   //视频
	GroupProxy                   //转发消息
	GroupRet                     //撤回
)

type Group struct {
	Msg   string `json:"msg,omitempty"`   //消息内容
	Ret   uint   `json:"ret,omitempty"`   //撤回消息
	Proxy Proxy  `json:"proxy,omitempty"` //转发消息
}

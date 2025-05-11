package socketMsg

type GroupConst uint

const (
	GroupStr   GroupConst = iota //文本 GroupMsg
	GroupMis                     //音频 GroupMsg
	GroupImg                     //图片 GroupMsg
	GroupVideo                   //视频 GroupMsg
	GroupProxy                   //转发消息 GroupProxyMsg
	GroupRet                     //撤回  GroupRetMsg
)

type Group struct {
	Msg   string `json:"msg,omitempty"`   //消息内容
	Proxy []TMsg `json:"proxy,omitempty"` //转发消息
	Ret   uint   `json:"ret,omitempty"`   //撤回消息
	Other any    `json:"other,omitempty"` //其它 前端自定义
}

type GroupMsg struct {
	Msg   string `json:"msg"`             //消息内容
	Other any    `json:"other,omitempty"` //其它 前端自定义
}

type GroupRetMsg struct {
	Ret   uint `json:"ret"`             //撤回
	Other any  `json:"other,omitempty"` //其它 前端自定义
}

type GroupProxyMsg struct {
	Proxy []TMsg `json:"proxy"`           //转发消息``
	Other any    `json:"other,omitempty"` //其它 前端自定义
}

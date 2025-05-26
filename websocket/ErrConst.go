package socketMsg

type ErrConst uint

const (
	ErrAny ErrConst = iota //任何错误都可能 Err
)

type Err struct {
	LastMsg any `json:"last_msg"` //上一条消息
	Msg     any `json:"msg"`      //消息内容
}

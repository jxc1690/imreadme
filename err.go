package response

type ResErr int

const (
	SUCCESS         ResErr = iota
	ErrOther               //未知错误
	ErrTokenExpired        //登录已失效或者已过期
	ErrPushData            //参数错误
	ErrWhere               //查询错误
	ErrUpData              //更新错误
	ErrNoUser              //不是自己发送的消息
	ErrNoSend              //发送失败
	ErrNoRet               //撤回失败
	ErrUpFile              //上传文件失败
)

const (
	ErrNoCode  ResErr = iota + 100 //验证码错误
	ErrGetCode                     //获取验证码失败
)
const (
	ErrUserNoWhere ResErr = iota + 200 //用户不存在
)

const (
	ErrFriend     ResErr = iota + 300 //添加好友失败
	ErrFriendOver                     //已经是好友
	ErrFriendMore                     //好友已满
)
const (
	ErrGroup ResErr = iota + 400 //群组错误
)

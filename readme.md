## **为强类型语言的后端 接口的参数必须按照类型来**
# 大致界面
## [第一次打开的界面](1.md)

## [登录/注册](2.md)

## [主界面](3.md)


# token
> 设置token 在请求中返回协议头中有一个set-token 建议在axios的拦截器中写 将token存储在store之类的地方 每次发生请求的时候在axios 加上x-token的协议头
# 流程
## 进入主界面的前面步骤
### 第一次
1. 向服务器请求对应邀请码的HOST
2. 调用ping接口测试是否联通
3. 获取验证码图片并将附带的token带入后面的请求
4. 向HOST发送登录请求 成功将有新的token（之前的数据可以存在本地）（token需要存在本地 没有token跳登录界面 有token跳主页面）

### 后面
1. ping测试 如果未成功 则重复上面第一次的1 2
2. 验证登录 是否在登录状态

## 主界面


## 联系人界面
1. 获取好友列表
2. 获取群组列表

## 频道
1. 获取频道列表

## 我的
1. 获取自己的信息（需要存在缓存里 如果有就不请求了）

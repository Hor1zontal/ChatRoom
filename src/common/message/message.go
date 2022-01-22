package common

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

const (
	LoginMessageType 				= "LoginMessageType"
	RegisterMessageType				= "RegisterMessageType"
	LoginResponseMessageType		= "LoginResponseMessageType"
	RegisterResponseMessageType		= "ResponseMessageType"
)

type LoginMessage struct {
	UserName string
	Password string
}

type ResponseMessage struct {
	Type  string
	Code  int    // 404 用户没找到， 403 账号或者密码错误, 200 登陆成功, 500 服务端错误
	Error string // 错误消息
	Data  string
}

type RegisterMessage struct {
	UserName        string
	Password        string
	PasswordConfirm string
}

type UserSendGroupMessage struct {
	GroupID  int    // target group id, 0 => all users
	UserName string // current user name
	Content  string // message content
}

type SendGroupMessageToClient struct {
	GroupID  int    // group id, 0 => all users
	UserName string // current user
	Content  string // message
}

type PointToPointMessage struct {
	SourceUserID   int
	SourceUserName string
	TargetUserID   int
	TargetUserName string
	Content        string
}

// on line user info
type UserInfo struct {
	ID       int
	UserName string
}
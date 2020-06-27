package message

const (
	//登录消息类型
	LoginMesType = "LoginMes"
	//登录返回消息类型
	LoginResMesType = "LoginResMes"
	//注册消息类型
	RegisterMesType = "registerMes"

	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

//传输消息
type Message struct {
	//消息类型
	Type string `json:"type"`
	//消息数据
	Data string `json:"data"`
}

//登录-ID和passwd
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

//登录-状态码和返回错误信息
type LoginResMes struct {
	Code    int    `json:"code"`
	UsersId []int  // 增加字段，保存用户id的切片
	Error   string `json:"error"`
}

type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体.
}
type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回状态码 400 表示该用户已经占有 200表示注册成功
	Error string `json:"error"` // 返回错误信息
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户的状态
}

//增加一个SmsMes //发送的消息
type SmsMes struct {
	Content string `json:"content"` //内容
	User           //匿名结构体，继承
}

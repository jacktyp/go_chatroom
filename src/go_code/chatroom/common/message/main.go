package message

const (
	//登录消息类型
	LoginMesType = "LoginMes"
	//登录返回消息类型
	LoginResMesType = "LoginResMes"
	//注册消息类型
	RegisterMesType = "registerMes"
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
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

//登录-状态码和返回错误信息
type LoginResMes struct {
	Code int `json:"code"`
	Error string `json:"error"`
}
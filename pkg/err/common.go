package err

var (
	ParamErr    = NewError(10100, "参数错误")
	SystemErr   = NewError(10101, "系统错误")
	LoginErr    = NewError(10102, "登录失败")
	PasswordErr = NewError(10101, "登录密码错误")
)

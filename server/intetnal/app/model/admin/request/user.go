package request

type Login struct {
	Username    string `form:"username" json:"username"`       // 用户名
	Password    string `form:"password" json:"password"`       // 密码
	CaptchaId   string `form:"captchaId" json:"captchaId"`     // 验证码ID
	VerifyValue string `form:"verifyValue" json:"verifyValue"` // 验证码
}
type Register struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type ManagerDoAdd struct {
	Username string `form:"username" json:"username"` // 用户名
	Password string `form:"password" json:"password"` // 密码
	Email    string `form:"email" json:"email"`
	Mobile   string `form:"mobile" json:"mobile"`
	RoleId   string `form:"role_id" json:"role_id"`
}

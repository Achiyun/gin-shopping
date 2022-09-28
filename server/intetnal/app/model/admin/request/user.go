package request

type Login struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	CaptchaId   string `json:"captchaId"`   // 验证码ID
	VerifyValue string `json:"VerifyValue"` // 验证码
}

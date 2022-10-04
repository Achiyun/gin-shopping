package request

type Role struct {
	Id          string `form:"id" json:"id"`
	Title       string `form:"title" json:"title"`             // 标题
	Description string `form:"description" json:"description"` // 描述
}

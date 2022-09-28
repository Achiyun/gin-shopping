package utils

import (
	"html/template"
)

// 把字符串解析成html
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}

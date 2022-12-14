package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

// md5加密
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

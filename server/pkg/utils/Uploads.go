package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 上传图片
func UploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)
	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、创建图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "./static/upload/" + day

	err1 := os.MkdirAll(dir, 0666)
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil

}

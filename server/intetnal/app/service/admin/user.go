package service

import (
	"net/http"

	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/system"
	service "github.com/Achiyun/gin-shopping/server/intetnal/app/service"
	sql "github.com/Achiyun/gin-shopping/server/intetnal/pkg/db"
	utils "github.com/Achiyun/gin-shopping/server/intetnal/pkg/utils"

	"github.com/gin-gonic/gin"
)

type LoginService struct {
	service.BaseApi
}

// CreateApi
// @Tags      LoginApi
// @Summary   登陆api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Router    / [post]
func (con LoginService) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "张三",
	})
}
func (con LoginService) Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	//2、查询数据库 判断用户以及密码是否存在
	userinfoList := []models.Manager{}
	password = utils.Md5(password)

	sql.DB.Where("username=? AND password=?", username, password).Find(&userinfoList)

	if len(userinfoList) > 0 {
		c.JSON(200, userinfoList)
	}
}

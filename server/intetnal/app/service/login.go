package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginService struct {
}

func (con LoginService) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "张三",
	})
}

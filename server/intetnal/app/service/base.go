package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (con BaseApi) Success(c *gin.Context, message string, redirectUrl string) {

	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}

func (con BaseApi) Error(c *gin.Context, message string, redirectUrl string) {

	c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}

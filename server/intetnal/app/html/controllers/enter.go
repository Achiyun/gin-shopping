package controllers

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers/admin"
	"github.com/Achiyun/gin-shopping/server/intetnal/app/html/controllers/base"
)

type ConGroup struct {
	AdminConGroup admin.AdminController
	BaseConGroup  base.BaseController
}

var ConGroupApp = new(ConGroup)

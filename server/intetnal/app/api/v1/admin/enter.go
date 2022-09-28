package admin

import (
	service "github.com/Achiyun/gin-shopping/server/intetnal/app/service"
)

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)

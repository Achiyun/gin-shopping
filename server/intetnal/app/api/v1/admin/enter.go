package admin

import (
	service "github.com/Achiyun/gin-shopping/server/intetnal/app/service"
)

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.AdminServiceGroup.UserService
)

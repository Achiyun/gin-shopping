package admin

import "github.com/Achiyun/gin-shopping/server/intetnal/app/service"

type AdminController struct {
	UserController
	MainController
}

var (
	userService = service.ServiceGroupApp.AdminServiceGroup.UserService
)

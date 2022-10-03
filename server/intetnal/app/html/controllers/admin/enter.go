package admin

import "github.com/Achiyun/gin-shopping/server/intetnal/app/service"

type AdminController struct {
	UserController
	MainController
	ManagerController
}

var (
	userService     = service.ServiceGroupApp.AdminServiceGroup.UserService
	mainService     = service.ServiceGroupApp.AdminServiceGroup.MainService
	manangerService = service.ServiceGroupApp.AdminServiceGroup.ManangerService
)

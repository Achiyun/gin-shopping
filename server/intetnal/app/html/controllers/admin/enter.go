package admin

import "github.com/Achiyun/gin-shopping/server/intetnal/app/service"

type AdminController struct {
	UserController
	MainController
	ManagerController
	RoleController
	AccessController
	FocusController
	GoodsController
	GoodsCateController
	GoodsTypeController
	GoodsTypeAttributeController
	NavController
	SettingController
}

var (
	userService     = service.ServiceGroupApp.AdminServiceGroup.UserService
	mainService     = service.ServiceGroupApp.AdminServiceGroup.MainService
	manangerService = service.ServiceGroupApp.AdminServiceGroup.ManangerService
	roleService     = service.ServiceGroupApp.AdminServiceGroup.RoleService
	accessService   = service.ServiceGroupApp.AdminServiceGroup.AccessService
)

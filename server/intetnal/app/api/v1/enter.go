package service

import "github.com/Achiyun/gin-shopping/server/intetnal/app/api/v1/admin"

type ServiceGroup struct {
	AdminServiceGroup admin.ApiGroup
}

var ServiceGroupApp = new(ServiceGroup)

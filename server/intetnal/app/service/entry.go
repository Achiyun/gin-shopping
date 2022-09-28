package service

import "github.com/Achiyun/gin-shopping/server/intetnal/app/service/admin"

type ServiceGroup struct {
	SystemServiceGroup admin.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

package service

import (
	"errors"

	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	sql "github.com/Achiyun/gin-shopping/server/intetnal/pkg/db"
	"github.com/Achiyun/gin-shopping/server/pkg/utils"
)

type UserService struct {
}

func (userService *UserService) Login(m *models.Manager) (managerInter *models.Manager, err error) {

	var manager models.Manager

	err = sql.DB.Where("username=?", m.Username).First(&manager).Error
	if err == nil {
		if ok := utils.BcryptCheck(m.Password, manager.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &manager, err
}

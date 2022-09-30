package admin

import (
	"errors"

	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	sql "github.com/Achiyun/gin-shopping/server/intetnal/pkg/db"
	uuid "github.com/satori/go.uuid"

	"github.com/Achiyun/gin-shopping/server/pkg/utils"
	"gorm.io/gorm"
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
func (userService *UserService) Register(m models.Manager) (managerInter models.Manager, err error) {
	var manager models.Manager
	if !errors.Is(sql.DB.Where("username = ?", m.Username).First(&manager).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return managerInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	m.Password = utils.BcryptHash(m.Password)
	m.UUID = uuid.NewV4()
	err = sql.DB.Create(&m).Error
	return m, err
}

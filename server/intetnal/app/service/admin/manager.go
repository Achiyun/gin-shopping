package admin

import (
	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"
	uuid "github.com/satori/go.uuid"
)

type ManangerService struct {
}

func (mananger *ManangerService) List(m *[]models.Manager) (err error) {

	err = global.GVA_DB.Preload("Role").Find(&m).Error
	return err
}
func (mananger *ManangerService) Info(username string, m *[]models.Manager) (err error) {

	err = global.GVA_DB.Where("username=?", username).Find(&m).Error
	return err
}
func (mananger *ManangerService) Create(m *models.Manager) (err error) {
	// uuid.NewV4(),
	m.UUID, _ = uuid.NewV4()
	err = global.GVA_DB.Create(&m).Error
	return err
}

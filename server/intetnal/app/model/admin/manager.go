package models

import uuid "github.com/satori/go.uuid"

type Manager struct {
	Id       int
	UUID     uuid.UUID
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
	Role     Role `gorm:"foreignKey:RoleId;references:Id"`
}

func (Manager) TableName() string {
	return "manager"
}

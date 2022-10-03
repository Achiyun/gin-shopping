package config

import "github.com/Achiyun/gin-shopping/server/intetnal/pkg/db"

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System System `mapstructure:"system" json:"system" yaml:"system"`

	// Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis db.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`

	// gorm
	Mysql db.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}

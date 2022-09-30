package config

type Server struct {
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	// Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}

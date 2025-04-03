package config

type Config struct {
	Zap    zap    `json:"zap" yaml:"zap"`
	Jwt    jwt    `json:"jwt" yaml:"jwt"`
	Mysql  mysql  `json:"mysql" yaml:"mysql"`
	Redis  redis  `json:"redis" yaml:"redis"`
	ES     es     `json:"es" yaml:"es"`
	System system `json:"system" yaml:"system"`
	Upload upload `json:"upload" yaml:"upload"`
}

package config

type Config struct {
	Zap   Zap   `json:"zap" yaml:"zap"`
	Jwt   Jwt   `json:"jwt" yaml:"jwt"`
	Mysql Mysql `json:"mysql" yaml:"mysql"`
	Redis Redis `json:"redis" yaml:"redis"`
	ES    ES    `json:"es" yaml:"es"`
}

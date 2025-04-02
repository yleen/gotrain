package config

import "fmt"

type System struct {
	Host           string `json:"-" yaml:"host"`
	Port           string `json:"-" yaml:"port"`
	Env            string `json:"-" yaml:"env"`
	RouterPrefix   string `json:"-" yaml:"router_prefix"`
	UseMultipoint  bool   `json:"use-multipoint" yaml:"use_multipoint"`
	SessionsSecret string `json:"sessions_secret" yaml:"sessions_secret"`
	OssType        string `json:"oss_type" yaml:"oss_type"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

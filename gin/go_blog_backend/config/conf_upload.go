package config

type Upload struct {
	Size int64  `json:"size" yaml:"size"`
	Path string `json:"path" yaml:"path"`
}

package core

import (
	"blog_backend/config"
	"blog_backend/utils"
	"gopkg.in/yaml.v3"
	"log"
)

func InitConf() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadYAML()

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Failed to unmarshal configuration: %v", err)
	}

	return c
}

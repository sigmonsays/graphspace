package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func GetDefaultConfig() *ApplicationConfig {
	c := &ApplicationConfig{}
	return c
}

func (c *ApplicationConfig) LoadYaml(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return err
	}
	return nil
}

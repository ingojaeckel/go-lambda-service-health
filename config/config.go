package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	Timeout int `yaml:"timeout"`
}

func LoadConfiguration(path string) (*Configuration, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Configuration
	if err := yaml.Unmarshal(c, &conf); err != nil {
		return nil, err
	}
	return &conf, err
}

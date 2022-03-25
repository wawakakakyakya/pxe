package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type port int

func LoadYamlConfig(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	var yc *Config
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(f, &yc)
	if err != nil {
		return nil, err
	}

	return yc, nil
}

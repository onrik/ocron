package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TZ    string            `yaml:"tz"`
	Env   map[string]string `yaml:"env"`
	Tasks []struct {
		Name      string            `yaml:"name"`
		Spec      string            `yaml:"spec"`
		Script    []string          `yaml:"script"`
		Finally   []string          `yaml:"finally"`
		OnError   []string          `yaml:"on_error"`
		OnSuccess []string          `yaml:"on_success"`
		Env       map[string]string `yaml:"env"`
	} `yaml:"tasks"`
}

func readConfig(filename string) (Config, error) {
	config := Config{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}

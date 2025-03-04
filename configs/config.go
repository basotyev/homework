package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App struct {
		Port int    `yaml:"port"`
		Env  string `yaml:"env"`
	}
	Db struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	}
	Token struct {
		AccessSecret  string `yaml:"access"`
		RefreshSecret string `yaml:"refresh"`
		AccessExpire  int    `yaml:"access_expire"`
		RefreshExpire int    `yaml:"refresh_expire"`
	}
	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
}

func NewConfig(path string) (*Config, error) {
	var cfg Config
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)

	if err := d.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

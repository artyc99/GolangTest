package configs

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type HttpServerConfig struct {
	ConfigFilePath string `env:"CONFIG_FILE_PATH,required" validate:"required"`

	LogLevel   string     `yaml:"log_level" validate:"required"`
	EchoConfig EchoConfig `yaml:"echo_config" validate:"required"`

	PostgresConfig PostgresConfig
}

type EchoConfig struct {
	Host string `yaml:"host" validate:"required"`
	Port uint   `yaml:"port" validate:"required"`
}

func GetHttpConfig() (cfg HttpServerConfig, err error) {
	err = env.Parse(&cfg)
	if err != nil {
		return HttpServerConfig{}, err
	}

	filename, err := filepath.Abs(cfg.ConfigFilePath)
	if err != nil {
		return HttpServerConfig{}, err
	}

	yamlFileIo, err := os.ReadFile(filename)
	if err != nil {
		return HttpServerConfig{}, err
	}

	err = yaml.Unmarshal(yamlFileIo, &cfg)
	if err != nil {
		return HttpServerConfig{}, err
	}

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return HttpServerConfig{}, err
	}

	return cfg, nil
}

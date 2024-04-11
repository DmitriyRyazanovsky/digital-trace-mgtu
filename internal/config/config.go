package config

import (
	"flag"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
)

type Config struct {
	Service  Service  `yaml:"service"`
	Logger   Logger   `yaml:"logger"`
	Database Database `yaml:"database"`
	Jwt      Jwt      `yaml:"jwt"`
	Mail     Mail     `yaml:"mail"`
}

func New() (*Config, error) {
	b, err := readFile()
	if err != nil {
		return nil, err
	}

	var out Config
	err = yaml.Unmarshal(b, &out)
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal bytes to config")
	}

	return &out, nil
}

func readFile() ([]byte, error) {
	var configPath string
	flag.StringVar(&configPath, "config_path", "./configs/config.yaml", "path to configuration file")
	flag.Parse()

	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read file with configuration")
	}

	return b, nil
}

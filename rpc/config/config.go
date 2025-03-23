package config

import (
	_ "embed"
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Secrets struct {
		AccessTokenSecret  string `yaml:"accessTokenSecret"`
		RefreshTokenSecret string `yaml:"refreshTokenSecret"`
	} `yaml:"Secrets"`

	TokenExpiry struct {
		AccessTokenExpiry  int64 `yaml:"accessTokenExpiry"`
		RefreshTokenExpiry int64 `yaml:"refreshTokenExpiry"`
	} `yaml:"TokenExpiry"`

	Redis struct {
		Address  string `yaml:"address"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"Redis"`

	Registry struct {
		RegistryAddress string `yaml:"registry_address"`
		UserName        string `yaml:"username"`
		Password        string `yaml:"password"`
	} `yaml:"Registry"`
}

var (
	configInstance *Config
	once           sync.Once
)

//go:embed config.yaml
var data []byte

func GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		configInstance, err = loadConfig()
	})
	return configInstance, err
}

func loadConfig() (*Config, error) {
	var conf Config

	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	etcdHostEnv := os.Getenv("ETCD_HOST")
	if etcdHostEnv != "" {
		conf.Registry.RegistryAddress = etcdHostEnv
	}

	return &conf, nil
}

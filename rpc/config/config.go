package config

import (
	_ "embed"
	"fmt"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Secrets struct {
		AccessTokenSecret  string `yaml:"AccessTokenSecret"`
		RefreshTokenSecret string `yaml:"RefreshTokenSecret"`
	} `yaml:"Secrets"`

	TokenExpiry struct {
		AccessTokenExpiry  int64 `yaml:"AccessTokenExpiry"`
		RefreshTokenExpiry int64 `yaml:"RefreshTokenExpiry"`
	} `yaml:"TokenExpiry"`

	Redis struct {
		Address  string `yaml:"Address"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		DB       int    `yaml:"DB"`
	} `yaml:"Redis"`

	Registry struct {
		RegistryAddress []string `yaml:"RegistryAddress"`
		UserName        string   `yaml:"Username"`
		Password        string   `yaml:"Password"`
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

	//etcdHostEnv := os.Getenv("ETCD_HOST")
	//if etcdHostEnv != "" {
	//	conf.Registry.RegistryAddress = etcdHostEnv
	//}

	return &conf, nil
}

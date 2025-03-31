package config

import (
	_ "embed"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"

	"gopkg.in/yaml.v3"
)

type Kitex struct {
	Service       string `yaml:"Service"`
	Address       string `yaml:"Address"`
	LogLevel      string `yaml:"LogLevel"`
	LogFileName   string `yaml:"LogFileName"`
	LogMaxSize    int    `yaml:"LogMaxSize"`
	LogMaxBackups int    `yaml:"LogMaxBackups"`
	LogMaxAge     int    `yaml:"LogMaxAge"`
}
type Config struct {
	Kitex   Kitex `yaml:"Kitex"`
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

func GetConfig() *Config {

	once.Do(func() {
		configInstance = loadConfig()
	})
	return configInstance
}

func loadConfig() *Config {
	var conf Config

	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}

	//etcdHostEnv := os.Getenv("ETCD_HOST")
	//if etcdHostEnv != "" {
	//	conf.Registry.RegistryAddress = etcdHostEnv
	//}

	return &conf
}

func LogLevel() klog.Level {
	level := GetConfig().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}

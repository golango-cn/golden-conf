package conf

import (
	"github.com/micro/go-micro/v2/web"
	"sync"
)

type Config struct {
	sync.RWMutex
	Services map[string]interface{}
}

var config *Config

func New() *Config {
	config = &Config{
		Services: make(map[string]interface{}),
	}
	return config
}

func GetConfig() *Config {
	return config
}

type registerFunc func(name string, service web.Service, config *Config) error

// 注册服务
func (c *Config) Register(name string, service web.Service, fc registerFunc) error {
	if fc == nil {
		return nil
	}
	return fc(name, service, c)
}

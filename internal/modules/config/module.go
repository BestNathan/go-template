package config

import (
	"errors"
	"go-template/internal/pkg/config"

	"github.com/spf13/viper"
)

var cm *ConfigModule

func init() {
	c := config.New()

	var cc Config
	c.GetViper().Unmarshal(&cc)

	cm = &ConfigModule{
		c:  c,
		cc: &cc,
	}
}

type ConfigModule struct {
	c  *config.Config
	cc *Config
}

func CM() *ConfigModule {
	if cm == nil {
		panic(errors.New("config module is not inited"))
	}

	return cm
}

func (cm *ConfigModule) GetViper() *viper.Viper {
	return cm.c.GetViper()
}

func (cm *ConfigModule) Config() *Config {
	return cm.cc
}

// func (cm *ConfigModule) Unmarshal(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
// 	// return cm.c.Unmarshal(key, rawVal, opts...)
// 	return cm.c.GetViper().UnmarshalKey(key, rawVal, opts...)
// }

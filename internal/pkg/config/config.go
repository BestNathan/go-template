package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var EnvConfigFileName = os.Getenv("GO_ENV")

var DefaultConfigFileName = "default"
var DefaultConfigEnvFileName = ".env"
var DefaultConfigFileType = "yaml"
var DefaultConfigFilePath = "./configs"

type ConfigOption func(*Config)

func ConfigFileNameOption(f string) ConfigOption {
	return func(c *Config) {
		c.fileName = f
	}
}

func ConfigEnvFileNameOption(ef string) ConfigOption {
	return func(c *Config) {
		c.envFileName = ef
	}
}

func ConfigFilePathOption(p string) ConfigOption {
	return func(c *Config) {
		c.filePath = p
	}
}

func ConfigFileTypeOption(ft string) ConfigOption {
	return func(c *Config) {
		c.fileType = ft
	}
}

type Config struct {
	mu          sync.Mutex
	filePath    string
	fileType    string
	fileName    string
	envFileName string
	viper       *viper.Viper
}

func New(opts ...ConfigOption) *Config {
	c := &Config{
		fileName:    DefaultConfigFileName,
		envFileName: DefaultConfigEnvFileName,
		fileType:    DefaultConfigFileType,
		filePath:    DefaultConfigFilePath,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Config) setPath(v *viper.Viper) {
	// path
	v.AddConfigPath(c.filePath)

	// filename
	if EnvConfigFileName != "" {
		v.SetConfigName(EnvConfigFileName)
	} else {
		v.SetConfigName(c.fileName)
	}

	// type
	v.SetConfigType(c.fileType)
}

func (c *Config) bindEnv(v *viper.Viper) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fn := path.Join(pwd, c.filePath, c.envFileName)
	fmt.Println("bindEnv: env file path=", fn)

	b, err := ioutil.ReadFile(fn)
	if err != nil {
		// panic(err)
		fmt.Println("bindEnv: read env config err: ", err)
		return
	}

	str := string(b)

	for _, raw := range strings.Split(str, "\n") {

		s := strings.Split(raw, " ")
		if len(s) != 2 {
			continue
		}

		// fmt.Printf("bindEnv: env=%s to field=%s\n", s[1], s[0])
		err := v.BindEnv(s[0], s[1])
		if err != nil {
			fmt.Println("bindEnv: bind err: ", err)
		}
	}

	v.AutomaticEnv()
}

func (c *Config) initViper() {
	c.mu.Lock()
	defer c.mu.Unlock()

	v := viper.New()

	// path
	c.setPath(v)
	// env
	c.bindEnv(v)

	// 查找并读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	c.viper = v
}

func (c *Config) GetViper() *viper.Viper {
	if c.viper == nil {
		c.initViper()
	}

	return c.viper
}

func (c *Config) Unmarshal(key string, rawVal interface{}, opt ...viper.DecoderConfigOption) error {
	return c.GetViper().Sub(key).Unmarshal(rawVal, opt...)
}

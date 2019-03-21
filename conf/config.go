package conf

import (
	"sync"
	"github.com/jinzhu/configor"
	"os"
)

type c struct {
	ServerDomain string `yaml:"ServerDomain"`
}

var config *c

func initConfig() {
	config = &c{}
	configor.New(&configor.Config{}).Load(config, "./conf/config.yml")
}

var once sync.Once

func GetConfig() *c {
	once.Do(func() { initConfig() })
	return config
}

func GetDomain() string {
	if domain := os.Getenv("ServerDomain"); domain != "" {
		return domain
	}
	return GetConfig().ServerDomain
}

package hihttp

import (
	"fmt"
	"github.com/tradjick/hiconfig"
)

type serverConfig struct {
	Host string
	Port string
}

func (sc *serverConfig) AsString() string {
	return fmt.Sprintf("%s:%s", sc.Host, sc.Port)
}

func fetchConfig(p string, l hiconfig.ConfigLoader) *serverConfig {
	c := new(serverConfig)
	l(p, c)
	return c
}

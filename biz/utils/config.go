package utils

import (
	"github.com/spf13/viper"
	"sync"
)

type ConfigResolve struct {
	Viper *viper.Viper
}

var (
	Resolve     *ConfigResolve
	ResolveOnce sync.Once
)

func GetConfigResolve() *ConfigResolve {
	ResolveOnce.Do(func() {
		Resolve = &ConfigResolve{
			Viper: viper.New(),
		}
		Resolve.Viper.SetConfigFile("/root/go/src/chat/conf/config.yml")
		Resolve.Viper.SetConfigType("yml")
		err := Resolve.Viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	})
	return Resolve
}

func (ins *ConfigResolve) GetConfigString(target string) string {
	return Resolve.Viper.GetString(target)
}

func (ins *ConfigResolve) GetConfigIntSlice(target string) []int {
	return Resolve.Viper.GetIntSlice(target)
}

func (ins *ConfigResolve) GetConfigStringSlice(target string) []string {
	return Resolve.Viper.GetStringSlice(target)
}

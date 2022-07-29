package utils

import (
	"github.com/spf13/viper"
	"os"
)

var ConfigResolve *viper.Viper

func init() {
	ConfigResolve = viper.New()
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ConfigResolve.SetConfigFile(pwd + "/conf/config.yml")
	ConfigResolve.SetConfigType("yml")
	err = ConfigResolve.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetConfigString(target string) string {
	return ConfigResolve.GetString(target)
}

func GetConfigIntSlice(target string) []int {
	return ConfigResolve.GetIntSlice(target)
}

func GetConfigStringSlice(target string) []string {
	return ConfigResolve.GetStringSlice(target)
}

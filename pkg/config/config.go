package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"log"
)

var ProviderSet = wire.NewSet(New)

func init() {
	viper.AutomaticEnv() // 绑定环境变量
}

func New(confPath string) (v *viper.Viper, err error) {
	v = viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigFile(confPath) //要在最后设置,否则不生效

	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}

	log.Printf("use config file -> %s\n", v.ConfigFileUsed())
	return v, err
}

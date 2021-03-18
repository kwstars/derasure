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
	v.AddConfigPath("./configs/derasure/")
	v.AddConfigPath("../../../../configs/derasure") //测试用
	v.SetConfigFile(confPath)
	v.SetConfigName("config")

	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}

	log.Printf("use config file -> %s\n", v.ConfigFileUsed())
	return v, err
}

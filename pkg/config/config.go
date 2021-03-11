package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"log"
)

var GlobalProviderSet = wire.NewSet(New)

func New(confPath string) (v *viper.Viper, err error) {
	v = viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath(" ../../configs/derasure/")
	v.SetConfigFile(confPath)
	v.SetConfigName("config")

	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}

	log.Printf("use config file -> %s\n", v.ConfigFileUsed())
	return v, err
}

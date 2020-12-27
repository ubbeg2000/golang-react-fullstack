package config

import "github.com/spf13/viper"

func New() {
	viper.SetConfigFile(`./config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
}

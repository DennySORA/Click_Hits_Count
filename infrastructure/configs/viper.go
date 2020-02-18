package configs

import "github.com/spf13/viper"

func InitializationViper() {
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error config file: " + err.Error())
	}
}

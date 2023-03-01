package initializer

import (
	"atm/global"
	"github.com/spf13/viper"
)

func loadConf() {
	viper.AddConfigPath("./")
	viper.SetConfigName("_config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic("[ERROR] invalid configure file.")
	}
	if err := viper.Unmarshal(&global.Conf); err != nil {
		panic("[ERROR] invalid configuration.")
	}
}

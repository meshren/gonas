package config

import (
	"github.com/spf13/viper"
	"gonas/utils"
	"log"
)

func init()  {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file not found.")
		}else{
			log.Println("read config file error.")
		}
		utils.ErrDetail(err)
	}
}


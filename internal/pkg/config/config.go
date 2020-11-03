package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("default")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

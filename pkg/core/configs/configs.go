package configs

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(path string) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./bin")
	viper.AddConfigPath("..")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

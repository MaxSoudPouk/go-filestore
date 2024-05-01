package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fmt.Println("START_READ_ENVIRONMENT_SUCCESSFULLY")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("READ_ENVIRONMENT_SUCCESSFULLY")
}

func GetEnv(key, defaultValue string) string {
	readValue := viper.GetString(key)
	if readValue == "" {
		return defaultValue
	}
	return readValue
}

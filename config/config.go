package config

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("local")   // Name of the config file (without extension)
	viper.SetConfigType("yaml")    // Type of the config file (yaml)
	viper.AddConfigPath("config/") // Search for the config file in the current directory

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

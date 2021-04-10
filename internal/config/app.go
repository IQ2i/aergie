package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitAppConfig() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("An error occured during config initialization. Check that you can write into your home dir.")
		os.Exit(1)
	}

	configPath := filepath.Join(userHomeDir, ".aergie")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.Mkdir(configPath, 0775)
	}

	configFile := filepath.Join(userHomeDir, ".aergie", "config.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		viper.WriteConfig()
	}

	viper.SetConfigFile(configFile)
	viper.ReadInConfig()

	viper.SetDefault("update.latest_version", "")
	viper.SetDefault("update.latest_check", "")
	viper.WriteConfig()
}

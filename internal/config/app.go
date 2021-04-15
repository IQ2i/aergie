package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitAppConfig() {
	userConfig()
	envConfig()
}

func userConfig() {
	// get config from file store in user home directory
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("An error occured during config initialization. Check that you can write into your home dir.")
		os.Exit(1)
	}

	configPath := filepath.Join(userHomeDir, ".aergie")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.Mkdir(configPath, 0775); err != nil {
			fmt.Println("An error occured during config initialization. Check that you can write into your home dir.")
			os.Exit(1)
		}
	}

	configFile := filepath.Join(userHomeDir, ".aergie", "config.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if err := viper.WriteConfig(); err != nil {
			fmt.Println("An error occured during config initialization. Check that you can write into your home dir.")
			os.Exit(1)
		}
	}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("An error occured during config initialization. Check that you can write into your home dir.")
		os.Exit(1)
	}

	// set default config values
	viper.SetDefault("update.latest_version", "")
	viper.SetDefault("update.latest_check", "")
	if err := viper.WriteConfig(); err != nil {
		fmt.Println("An error occured during config initialization. Check that you can write into your home dir.")
		os.Exit(1)
	}
}

func envConfig() {
	env := "prod"
	if os.Getenv("AE_VERSION") == "DEV" {
		env = "dev"
	}
	os.Setenv("AE_ENV", env)

	updateDomain := "https://get.aergie.com"
	if env == "dev" {
		updateDomain = "https://localhost"
	}
	os.Setenv("AE_UPDATE_DOMAIN", updateDomain)
}

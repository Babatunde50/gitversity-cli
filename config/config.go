package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Token struct {
	FileName   string `mapstructure:"FileName"`
	AppName    string `mapstructure:"AppName"`
	FolderName string `mapstructure:"FolderName"`
}

type Service struct {
	Host string `mapstructure:"host"`
}

type Config struct {
	Token    Token              `mapstructure:"token"`
	Services map[string]Service `mapstructure:"services"`
}

var C Config

func LoadConfig() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rootDir := filepath.Dir(exePath)
	configPath := filepath.Join(rootDir, "config.yaml")

	// Check if the config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// If the file does not exist, create it with default values
		defaultConfig := []byte(`
	token:
	FileName: "token"
	AppName: "gitversity"
	FolderName: ".config"

	services:
	user_grpc:
		host: "grpc.dev.user.api.gitversity.com:443"
	assignment_grpc:
		host: "grpc.dev.assignment.api.gitversity.com:443"
	git_grpc:
		host: "grpc.dev.git.api.gitversity.com:443"
`)

		err = os.WriteFile(configPath, defaultConfig, 0644)
		if err != nil {
			fmt.Println("Error creating default config file:", err)
			return
		}
	}

	viper.AddConfigPath(rootDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

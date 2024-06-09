// config/config.go
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

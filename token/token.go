package token

import (
	"os"
	"path/filepath"

	"github.com/Babatunde50/gitversity-cli/config"
)

func getTokenFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, config.C.Token.FolderName, config.C.Token.AppName)
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(configDir, config.C.Token.FileName), nil
}

func SaveToken(token string) error {
	path, err := getTokenFilePath()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(token), 0600)
}

func LoadToken() (string, error) {
	path, err := getTokenFilePath()
	if err != nil {
		return "", err
	}
	token, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func DeleteToken() error {
	// Get the token file path
	path, err := getTokenFilePath()
	if err != nil {
		return err
	}

	// Remove the token file
	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

package metadata

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

type Metadata struct {
	Timestamp    time.Time         `json:"timestamp"`
	User         string            `json:"user"`
	Machine      string            `json:"machine"`
	RepoId       string            `json:"repo_id"`
	AssignmentId string            `json:"assignment_id"`
	BaseDir      string            `json:"base_dir"`
	CustomData   map[string]string `json:"custom_data"`
}

func (m *Metadata) Write(baseDir string) error {
	identifierPath := filepath.Join(baseDir, ".gitversity")

	// Create the identifier directory
	if err := os.MkdirAll(identifierPath, 0755); err != nil {
		return fmt.Errorf("failed to create identifier directory: %w", err)
	}

	// Store metadata in a JSON file
	metadataFilePath := filepath.Join(identifierPath, "metadata.json")
	if err := writeJSON(metadataFilePath, *m); err != nil {
		return err
	}

	return nil
}

func Read() (*Metadata, error) {

	currentDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	metadataFSPath := path.Join(currentDir, ".gitversity", "metadata.json")
	metadataBytes, err := os.ReadFile(metadataFSPath)

	if err != nil {
		return nil, err
	}

	var md Metadata
	err = json.Unmarshal(metadataBytes, &md)

	if err != nil {
		return nil, err
	}

	return &md, nil
}

func writeJSON(filePath string, data Metadata) error {

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create or open file: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(jsonData); err != nil {
		return fmt.Errorf("failed to write JSON data to file: %w", err)
	}

	return nil
}

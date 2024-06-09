// pkg/hello.go
package pkg

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/services/ggs"
)

// TODO: do not add .gitversity to the git server

type File struct {
	Path     string      `json:"path"`
	Content  string      `json:"content"`
	Mode     fs.FileMode `json:"mode"`
	ModeTime string      `json:"modeTime"`
	CTime    string      `json:"cTime"`
}

func RunStage(cli cli.Cli, repoId string, args []string) error {

	fs, err := getFileDetails(".", args)
	if err != nil {
		return err
	}

	var files []*ggs.AppFile

	for _, f := range fs {
		files = append(files, &ggs.AppFile{
			Path:     f.Path,
			Content:  f.Content,
			Mode:     int32(f.Mode),
			ModeTime: f.ModeTime,
			CTime:    f.CTime,
		})
	}

	return cli.Stage(repoId, files)
}

// GetFileDetails returns the details of the files in the specified directory
func getFileDetails(root string, patterns []string) ([]File, error) {
	var files []File

	// Traverse the directory and handle files
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if path == root {
			return nil
		}

		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		matched := false
		for _, pattern := range patterns {
			if pattern == "" {
				continue
			}
			// Check if the pattern is a directory
			if info.IsDir() {
				if strings.HasPrefix(path, pattern) || pattern == relativePath {
					matched = true
					break
				}
			} else {
				// Match against the file name and relative path
				matchName, err := filepath.Match(pattern, info.Name())
				matchPath, err2 := filepath.Match(pattern, relativePath)
				if err != nil || err2 != nil {
					return err
				}
				if matchName || matchPath || strings.HasPrefix(relativePath, pattern) {
					matched = true
					break
				}
			}
		}

		if !info.IsDir() && matched {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			modeTime := info.ModTime().Format(time.RFC3339)
			ctime := info.ModTime().Format(time.RFC3339) // Unix does not support ctime, using modtime as a substitute

			files = append(files, File{
				Path:     path,
				Content:  string(content),
				Mode:     info.Mode(),
				ModeTime: modeTime,
				CTime:    ctime,
			})
		}

		return nil
	})

	return files, err
}

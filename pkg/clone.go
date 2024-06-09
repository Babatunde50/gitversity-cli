// pkg/hello.go
package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/metadata"
	"github.com/Babatunde50/gitversity-cli/services/ggs"
)

func RunClone(cli cli.Cli, assignmentId string) error {

	// TODO: get userAssignments
	userAssignments, err := cli.GetUserAssignments(assignmentId)

	if err != nil {
		return err
	}

	repoResp, err := cli.GetRepo(userAssignments.RepoId)

	if err != nil {
		return err
	}

	return clone(cli, repoResp.Repo.Name, repoResp.Repo.Id, userAssignments.AssignmentId, repoResp.Repo.Index.Entries)

}

func clone(cli cli.Cli, baseDir, repoId string, assignmentId string, entries []*ggs.IndexEntry) error {

	// Create the base directory if it does not exist
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create base directory %s: %w", baseDir, err)
	}

	for _, entry := range entries {

		// Combine base directory with the entry path
		fullPath := filepath.Join(baseDir, entry.Path)

		// Extract the directory part from the path
		dir := filepath.Dir(fullPath)

		// Create the directory if it does not exist
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory for path %s: %w", dir, err)
		}

		// fetch sha1
		blobResp, err := cli.GetBlob(entry.Sha1)

		if err != nil {
			return err
		}

		fileContent := []byte(blobResp.Content)

		// Write the file content to the path
		if err := os.WriteFile(fullPath, fileContent, 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", entry.Path, err)
		}
	}

	// TODO: update...
	// Generate metadata
	m := metadata.Metadata{
		Timestamp: time.Now(),
		User:      "John Doe <john.doe@example.com>",                          // Replace with actual user information
		Machine:   fmt.Sprintf("%s (%s)", os.Getenv("HOSTNAME"), "127.0.0.1"), // Get hostname and IP
		CustomData: map[string]string{
			"app_version":  "1.2.3",
			"project_name": "MyProject",
		},
		BaseDir:      baseDir,
		RepoId:       repoId,
		AssignmentId: assignmentId,
	}

	if err := m.Write(baseDir); err != nil {
		return err
	}

	return nil
}

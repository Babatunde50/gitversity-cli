// pkg/hello.go
package pkg

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/metadata"
	"github.com/logrusorgru/aurora/v4"
)

type LocalFileInfo struct {
	Path string
	Sha1 string // Content of the file, if available
}

func RunStatus(cli cli.Cli) error {
	md, err := metadata.Read()
	if err != nil {
		return err
	}

	repoResp, err := cli.GetRepo(md.RepoId)
	if err != nil {
		return err
	}

	commitFilesResp, err := cli.GetCommitFiles(repoResp.Repo.Head.CommitId)
	if err != nil {
		return err
	}

	commitFiles := commitFilesResp.GetFiles()
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	localFiles, err := GetAllLocalFiles(currentDir)
	if err != nil {
		return err
	}

	// Create maps to store files by path for efficient lookup
	localFilesMap := make(map[string]LocalFileInfo)          // my local directory
	repoFilesMap := make(map[string]LocalFileInfo)           // files in my repo, with the staged files included
	repoLastCommitFilesMap := make(map[string]LocalFileInfo) // files in the last commit

	for _, file := range commitFiles {
		if !shouldIgnore(file.Path) {
			repoLastCommitFilesMap[file.Path] = LocalFileInfo{
				Path: file.Path,
				Sha1: EncryptToSha1(file.Content),
			}
		}
	}

	for _, entry := range repoResp.Repo.Index.Entries {
		if !shouldIgnore(entry.Path) {
			repoFilesMap[entry.Path] = LocalFileInfo{
				Path: entry.Path,
				Sha1: entry.Sha1,
			}
		}
	}

	for _, entry := range localFiles {
		if !shouldIgnore(entry.Path) {
			localFilesMap[entry.Path] = entry
		}
	}

	var stagedFiles, unstagedModifiedFiles, deletedFiles, untrackFiles []string

	for path, localFile := range localFilesMap {
		// Check if file exists in the Repo
		repoFile, ok := repoFilesMap[path]
		if !ok {
			// File is new
			untrackFiles = append(untrackFiles, path)
		} else {
			// File exists in both, compare SHA-1 hashes
			if localFile.Sha1 != repoFile.Sha1 {
				// File is modified
				unstagedModifiedFiles = append(unstagedModifiedFiles, path)
			}
			lastCommitRepoFile, ok := repoLastCommitFilesMap[path]

			if ok {
				if localFile.Sha1 == repoFile.Sha1 && lastCommitRepoFile.Sha1 != repoFile.Sha1 {
					stagedFiles = append(stagedFiles, path)
				}
			} else {
				if localFile.Sha1 == repoFile.Sha1 {
					stagedFiles = append(stagedFiles, path)
				}
			}
			// Remove file from repoFilesMap to track deleted files
			delete(repoFilesMap, path)
		}
	}

	// Remaining files in repoFilesMap are deleted files
	for path := range repoFilesMap {
		deletedFiles = append(deletedFiles, path)
	}

	printGitStatusMessage(repoResp.Repo.Head.Name, stagedFiles, unstagedModifiedFiles, deletedFiles, untrackFiles)

	return nil
}

func printGitStatusMessage(branchName string, stagedFiles, unstagedModifiedFiles, deletedFiles, untrackedFiles []string) {
	// Print branch name and current commit hash
	fmt.Println(aurora.BrightBlue("On branch " + branchName).String())
	fmt.Println(aurora.BrightCyan("Your branch is up to date with 'origin/" + branchName + "'.").String())

	// fmt.Println(aurora.BrightGreen("nothing to commit, working directory clean").String())

	// Print working directory status
	if len(stagedFiles) == 0 && len(unstagedModifiedFiles) == 0 && len(deletedFiles) == 0 && len(untrackedFiles) == 0 {
		fmt.Println(aurora.BrightGreen("nothing to commit, working directory clean").String())
	}
	// else {
	// 	fmt.Println(aurora.BrightYellow("Changes not staged for commit:").String())
	// }

	// Print staged files
	if len(stagedFiles) > 0 {
		fmt.Println(aurora.BrightYellow("\nChanges to be committed:").String())
		for _, file := range stagedFiles {
			fmt.Println(aurora.BrightGreen("  " + file).String())
		}
	}

	// Print unstaged modified files
	if len(unstagedModifiedFiles) > 0 {
		fmt.Println(aurora.BrightRed("\nChanges not staged for commit:").String())
		for _, file := range unstagedModifiedFiles {
			fmt.Println(aurora.BrightRed("  " + file).String())
		}
	}

	// Print deleted files
	if len(deletedFiles) > 0 {
		fmt.Println(aurora.BrightRed("\nDeleted files:").String())
		for _, file := range deletedFiles {
			fmt.Println(aurora.BrightRed("  " + file).String())
		}
	}

	// Print untracked files
	if len(untrackedFiles) > 0 {
		fmt.Println(aurora.BrightYellow("\nUntracked files:").String())
		for _, file := range untrackedFiles {
			fmt.Println(aurora.BrightYellow("  " + file).String())
		}
	}
}

func GetAllLocalFiles(dirPath string) ([]LocalFileInfo, error) {
	var files []LocalFileInfo

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == ".gitversity" {
			return filepath.SkipDir
		}

		if !d.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			relativePath, err := filepath.Rel(dirPath, path)

			if err != nil {
				return err
			}

			files = append(files, LocalFileInfo{
				Path: relativePath,
				Sha1: EncryptToSha1(string(content)),
			})
		}

		return nil
	})

	return files, err
}

func EncryptToSha1(content string) string {
	hash := sha1.New()

	hash.Write([]byte(content))

	hashSum := hash.Sum(nil)

	return hex.EncodeToString(hashSum)
}

func shouldIgnore(path string) bool {
	// Check if the path starts with .gitversity/
	return strings.HasPrefix(path, ".gitversity/")
}

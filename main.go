package main

import (
	"os"

	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/cmd"
)

func main() {

	gitversityCli := cli.NewGitversityCli()
	rootCmd := cmd.NewRootCommand(gitversityCli)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

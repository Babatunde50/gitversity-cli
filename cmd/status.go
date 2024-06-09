/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/pkg"
	"github.com/spf13/cobra"
)

func NewStatusCmd(cli cli.Cli) *cobra.Command {
	var statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Display the status of the repository in the current directory",
		Long: `The 'status' command shows the status of the repository in reference to the current directory.
	It indicates which changes are staged for the next commit, which changes are not staged, and which files are untracked.
	
	Usage:
	  gitversity status
	
	This command does not take any flags or arguments.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return pkg.RunStatus(cli)
		},
	}

	return statusCmd
}

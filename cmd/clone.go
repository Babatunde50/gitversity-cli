/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/pkg"
	"github.com/spf13/cobra"
)

// TODO: stop user from cloning repo that is not theirs...

func NewCloneCmd(cli cli.Cli) *cobra.Command {
	cloneCmd := &cobra.Command{
		Use:   "clone --assignment-id <assignment ID>",
		Short: "Clone an assignment repository to the current directory",
		Long: `The 'clone' command clones an assignment repository to the current directory. 
	This command requires an assignment ID to identify which repository to clone.
	
	Usage:
	  gitversity clone --assignment-id <assignment ID>
	
	Examples:
	  # Clone an assignment repository with an assignment ID
	  gitversity clone --assignment-id 12345
	
	Flags:
	  --assignment-id    The assignment ID to identify the repository to clone. This flag is required.
	
	This command does not take any additional positional arguments.
	`,

		RunE: func(cmd *cobra.Command, args []string) error {

			assignmentId, _ := cmd.Flags().GetString("assignment_id")

			return pkg.RunClone(cli, assignmentId)

		},
	}

	cloneCmd.Flags().String("assignment_id", "", "gitversity clone --repoId <repoId>")
	cloneCmd.MarkFlagRequired("assignment_id")

	return cloneCmd

}

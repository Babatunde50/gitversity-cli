/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/pkg"
	"github.com/spf13/cobra"
)

func NewPushCmd(cli cli.Cli) *cobra.Command {
	pushCmd := &cobra.Command{
		Use:   "push",
		Short: "Submit the assignment",
		Long: `The 'push' command submits the committed changes to the central repository, effectively submitting your assignment.
	This command should be used after you have added and committed your changes.
	
	Usage:
	  gitversity push
	
	Examples:
	  # Submit the current assignment
	  gitversity push
	
	This command does not take any additional flags or arguments.
	`,
		RunE: func(cmd *cobra.Command, args []string) error {

			pkg.RunPush(cli)

			return nil
		},
	}

	return pushCmd
}

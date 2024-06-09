/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/metadata"
	"github.com/Babatunde50/gitversity-cli/pkg"
	"github.com/spf13/cobra"
)

func NewCommitCmd(cli cli.Cli) *cobra.Command {
	commitCmd := &cobra.Command{
		Use:   "commit --message <commit message>",
		Short: "Add a commit to the repo",
		Long: `The 'commit' command is used to create a new commit in the repository with the changes that have 
	been staged. You must provide a commit message to describe the changes made.
	
	Usage:
	  gitversity commit --message "Your commit message"
	
	Examples:
	  # Commit with a message
	  gitversity commit --message "Initial commit"
	
	  # Commit with a multi-line message
	  gitversity commit --message "Fix bug in user login\n\n- Fixed issue with session handling\n- Added additional logging"
	
	Flags:
	  --message, -m   The commit message to include with the commit. This flag is required.
	
	This command does not take any additional arguments.
	`,
		RunE: func(cmd *cobra.Command, args []string) error {

			message, _ := cmd.Flags().GetString("message")

			md, err := metadata.Read()
			if err != nil {
				return err
			}

			if err = pkg.RunCommit(cli, md.RepoId, message); err != nil {
				return err
			}

			return nil
		},
	}

	commitCmd.Flags().StringP("message", "m", "", "Add a commit to the repo")

	commitCmd.MarkFlagRequired("message")

	return commitCmd
}

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

func NewStageCmd(cli cli.Cli) *cobra.Command {
	return &cobra.Command{
		Use:   "add <file>",
		Short: "Add file to the staging area",
		Long: `The 'add' command stages a specified file, making it ready to be committed to the repository.
	This is typically the first step before committing changes.
	
	Usage:
	  gitversity add <file>
	
	Examples:
	  # Stage a single file
	  gitversity add README.md
	
	  # Stage multiple files
	  gitversity add file1.txt file2.txt
	
	Arguments:
	  <file>   The file or files to stage. You can specify multiple files separated by spaces.
	
	This command does not take any additional flags.
	`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for i, val := range args {
				if val == "." {
					args[i] = "*"
				}
			}

			md, err := metadata.Read()

			if err != nil {
				return err
			}

			return pkg.RunStage(cli, md.RepoId, args)
		},
	}
}

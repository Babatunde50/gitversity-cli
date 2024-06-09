/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/pkg"
	"github.com/spf13/cobra"
)

func NewResetCmd(cli cli.Cli) *cobra.Command {
	var resetCmd = &cobra.Command{
		Use:   "reset [HEAD] [files...]",
		Short: "Remove file(s) from the staging area",
		Args:  cobra.MinimumNArgs(1),
		Long: `The 'reset' command is used to remove file(s) from the staging area, effectively unstaging changes that 
	have been added. This is useful when you want to undo a 'git add' operation.
	
	Usage:
	  gitversity reset [HEAD] [files...]
	
	Examples:
	  # Unstage a specific file
	  gitversity reset HEAD file.txt
	
	  # Unstage multiple files
	  gitversity reset HEAD file1.txt file2.txt
	
	  # Unstage all changes
	  gitversity reset HEAD .
	
	Arguments:
	  [HEAD]     Indicates the commit reference. Typically, this will be 'HEAD' to refer to the latest commit.
	  [files...] One or more files to unstage from the staging area.
	`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) > 0 && args[0] != "HEAD" {
				return errors.New("error")
			}

			// If no file arguments provided, reset all staged files
			if len(args) == 1 {
				args = append(args, ".")
			}

			for i, val := range args {
				if val == "." {
					args[i] = "*"
				}
			}

			args = args[1:]

			return pkg.RunUnstage(cli, args)

		},
	}

	return resetCmd
}

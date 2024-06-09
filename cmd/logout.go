/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/pkg"
	"github.com/spf13/cobra"
)

func NewLogoutCmd(cli cli.Cli) *cobra.Command {
	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout the signed-in user",
		Long: `The 'logout' command logs the currently signed-in user out of the Gitversity system.
	This should be used when you want to end your session securely.
	
	Usage:
	  gitversity logout
	
	Examples:
	  # Log out the current user
	  gitversity logout
	
	This command does not take any additional flags or arguments.
	`,
		RunE: func(cmd *cobra.Command, args []string) error {

			return pkg.RunLogout(cli)
		},
	}

	return logoutCmd
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/pkg"
	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

func NewLoginCmd(cli cli.Cli) *cobra.Command {
	loginCmd := &cobra.Command{
		Use:   "login [email] [password]",
		Short: "Authenticate User",
		Long: `The 'login' command authenticates a user by verifying the provided email and password.
	This should be used to sign in to the Gitversity system before performing any other operations.
	
	Usage:
	  gitversity login [email] [password]
	
	Examples:
	  # Log in using email and password
	  gitversity login user@example.com password123
	
	Arguments:
	  [email]      The email address associated with your Gitversity account.
	  [password]   The password associated with your Gitversity account.
	
	This command does not take any additional flags.
	`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")

			if err := pkg.RunLogin(cli, email, password); err != nil {
				return fmt.Errorf(aurora.Red(formatError(err)).String())
			}

			fmt.Println(aurora.Green("Login successful! Welcome,"), aurora.Bold(email))

			return nil
		},
	}

	loginCmd.Flags().StringP("email", "e", "", "email of the user")
	loginCmd.Flags().StringP("password", "p", "", "password of the user")
	loginCmd.MarkFlagRequired("email")
	loginCmd.MarkFlagRequired("password")

	return loginCmd
}

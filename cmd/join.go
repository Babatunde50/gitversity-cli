/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/pkg"
	"github.com/spf13/cobra"
)

func NewJoinCmd(cli cli.Cli) *cobra.Command {
	var joinCmd = &cobra.Command{
		Use:   "join [invite code]",
		Short: "Join an assignment using an invite code",
		Long: `The 'join' command allows you to join an assignment by providing a valid invite code.
	If you are not registered, you will be prompted to sign up. You can also use the optional 
	'--auto-clone' flag to automatically clone the assignment repository after joining.
	
	Usage:
	  gitversity join [invite code] [flags]
	
	Examples:
	  # Join an assignment with an invite code
	  gitversity join ABC123XYZ
	
	  # Join an assignment with an invite code and auto-clone the repository
	  gitversity join ABC123XYZ --auto-clone
	
	Arguments:
	  [invite code]   The invite code provided to you for joining the assignment.
	
	Flags:
	  --auto-clone    Automatically clone the assignment repository after joining.
	
	This command does not take any additional arguments.
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			inviteCode, _ := cmd.Flags().GetString("invite-code")
			autoClone, _ := cmd.Flags().GetBool("auto-clone")

			return pkg.RunJoin(cli, inviteCode, autoClone)

		},
	}

	joinCmd.Flags().StringP("invite-code", "i", "", "Invite code of the assignment the user is trying to join")
	joinCmd.Flags().BoolP("auto-clone", "a", false, "Ask whether to clone the repository files to the current directory")
	joinCmd.MarkFlagRequired("invite-code")

	return joinCmd
}

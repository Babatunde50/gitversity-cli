package cmd

import (
	"fmt"

	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	cobra.OnInitialize(config.LoadConfig)
}

func NewRootCommand(cli cli.Cli) *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:   "gitversity",
		Short: "CLI tool for submitting assignment on gitversity.com",
		Long: `Gitversity is a command-line application designed to streamline the process of 
		assignment submission by providing Git-like commands. It allows users to clone assignments, 
		add and commit changes, and push updates to a central repository.
		
		Available commands:
		  - login: Log in to the Gitversity system.
		  - logout: Log out of the Gitversity system.
		  - join: Join an assignment repository with an optional auto-clone feature.
		  - clone: Clone an assignment repository using the assignment ID.
		  - add: Add files to the staging area for the next commit.
		  - commit: Commit the staged changes with a commit message.
		  - status: Display the status of the repository, showing changes to be committed and untracked files.
		  - push: Push the committed changes to the central repository.
		  - reset: Reset the current state of the repository to a previous state.
		  - help: Display help information for a specified command.
		
		Use "gitversity [command] --help" for more information about a command.`,
	}

	rootCmd.AddCommand(NewLoginCmd(cli))
	rootCmd.AddCommand(NewCommitCmd(cli))
	rootCmd.AddCommand(NewCloneCmd(cli))
	rootCmd.AddCommand(NewJoinCmd(cli))
	rootCmd.AddCommand(NewResetCmd(cli))
	rootCmd.AddCommand(NewLogoutCmd(cli))
	rootCmd.AddCommand(NewPushCmd(cli))
	rootCmd.AddCommand(NewStageCmd(cli))
	rootCmd.AddCommand(NewStatusCmd(cli))

	return rootCmd
}

func formatError(err error) string {

	fmt.Println(err)

	// Check if the error is a gRPC status error
	st, ok := status.FromError(err)
	if !ok {
		// Not a gRPC error
		return "An unexpected error occurred"
	}

	// Extract the gRPC error code and description
	code := st.Code()
	// description := st.Message()

	// Map the gRPC error code to a user-friendly message
	switch code {
	case codes.Unauthenticated:
		return "Authentication failed: Please check your email and password."
	case codes.PermissionDenied:
		return "Permission denied: You do not have access to this resource."
	case codes.NotFound:
		return "Resource not found: The requested item does not exist."
	default:
		return fmt.Sprintf("Error: %s", err)
	}
}

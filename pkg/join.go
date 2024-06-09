package pkg

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/token"
)

var passwordPrompt = &survey.Password{
	Message: "Input your password?",
}

// the questions to ask
var confirmPasswordPrompt = &survey.Password{
	Message: "Confirm your password?",
}

var fullNamePrompt = &survey.Input{
	Message: "Input your full name",
}

var emailPrompt = &survey.Input{
	Message: "Input your email",
}

var loginOrSignupPrompt = &survey.Confirm{
	Message: "Do you already have an account?",
}

var askOne = survey.AskOne

func RunJoin(cli cli.Cli, inviteCode string, autoClone bool) error {

	assgnmt, err := cli.GetAssignmentByInviteCode(inviteCode)

	if err != nil {
		return err
	}

	_, err = token.LoadToken()

	if err != nil {
		var isLoggingIn bool
		var email string
		var password string
		var fullName string
		var confirmPassword string
		askOne(loginOrSignupPrompt, &isLoggingIn)
		if isLoggingIn {
			askOne(emailPrompt, &email)
			askOne(passwordPrompt, &password)

			err := cli.Login(email, password)

			if err != nil {
				return err
			}
		} else {
			askOne(fullNamePrompt, &fullName)
			askOne(emailPrompt, &email)
			askOne(passwordPrompt, &password)
			askOne(confirmPasswordPrompt, &confirmPassword)

			if err := cli.SignUp(email, fullName, confirmPassword, password); err != nil {
				return err
			}

			return nil

		}
	}

	assgn, err := cli.JoinAssignment(assgnmt.Id)

	if err != nil {
		return err
	}

	if autoClone {
		repoResp, err := cli.GetRepo(assgn.RepoId)

		if err != nil {
			return err
		}

		return clone(cli, repoResp.Repo.Name, repoResp.Repo.Id, assgn.AssignmentId, repoResp.Repo.Index.Entries)

	}

	return nil

}

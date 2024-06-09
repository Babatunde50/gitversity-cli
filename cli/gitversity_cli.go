package cli

import (
	"github.com/Babatunde50/gitversity-cli/api/assignment"
	"github.com/Babatunde50/gitversity-cli/api/git"
	"github.com/Babatunde50/gitversity-cli/api/user"
	"github.com/Babatunde50/gitversity-cli/services/gas"
	"github.com/Babatunde50/gitversity-cli/services/ggs"
)

type GitversityCli struct{}

func NewGitversityCli() *GitversityCli {
	cli := &GitversityCli{}
	return cli
}

func (cli *GitversityCli) Login(email, password string) error {
	return user.Login(email, password)
}

func (cli *GitversityCli) Commit(repoId, message string) error {
	return git.Commit(repoId, message)
}

func (cli *GitversityCli) Stage(repoId string, files []*ggs.AppFile) error {
	return git.Stage(repoId, files)
}

func (cli *GitversityCli) GetRepo(repoId string) (*ggs.GetRepoResponse, error) {
	repo, err := git.GetRepo(repoId)
	if err != nil {
		return repo, err
	}
	return repo, nil
}

func (cli *GitversityCli) GetBlob(sha1 string) (*ggs.GetBlobResponse, error) {
	return git.GetBlob(sha1)
}

func (cli *GitversityCli) GetAssignmentByInviteCode(inviteCode string) (*gas.Assignment, error) {
	return assignment.GetAssignmentByInviteCode(inviteCode)
}

func (cli *GitversityCli) SignUp(email string, fullName string, confirmPassword string, password string) error {
	return user.SignUp(email, fullName, confirmPassword, password)
}

func (cli *GitversityCli) JoinAssignment(id string) (*gas.JoinAssignmentResponse, error) {
	return assignment.JoinAssignment(id)
}

func (cli *GitversityCli) UnStage(repoId string, paths []string) error {
	return git.UnStage(repoId, paths)
}

func (cli *GitversityCli) GetCommitFiles(commitId string) (*ggs.GetCommitFilesResponse, error) {
	return git.GetCommitFiles(commitId)
}

func (cli *GitversityCli) Push(assignmentId string) (*gas.Submission, error) {
	return assignment.SubmitAssignment(assignmentId)
}

func (cli *GitversityCli) GetUserAssignments(assignmentId string) (*gas.UserAssignments, error) {
	return assignment.GetUserAssignments(assignmentId)
}

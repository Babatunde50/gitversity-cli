package cli

import (
	"github.com/Babatunde50/gitversity-cli/services/gas"
	"github.com/Babatunde50/gitversity-cli/services/ggs"
)

type Cli interface {
	Login(email, password string) error
	Commit(repoId, message string) error
	Stage(repoId string, files []*ggs.AppFile) error
	GetRepo(repoId string) (*ggs.GetRepoResponse, error)
	GetBlob(sha1 string) (*ggs.GetBlobResponse, error)
	GetAssignmentByInviteCode(inviteCode string) (*gas.Assignment, error)
	SignUp(email string, fullName string, confirmPassword string, password string) error
	JoinAssignment(id string) (*gas.JoinAssignmentResponse, error)
	UnStage(repoId string, paths []string) error
	GetCommitFiles(commitId string) (*ggs.GetCommitFilesResponse, error)
	Push(assignmentId string) (*gas.Submission, error)
	GetUserAssignments(assignmentId string) (*gas.UserAssignments, error)
}

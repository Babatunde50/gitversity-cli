package git

import (
	"context"
	"fmt"

	"github.com/Babatunde50/gitversity-cli/api"
	"github.com/Babatunde50/gitversity-cli/config"
	"github.com/Babatunde50/gitversity-cli/services/ggs"
	"github.com/Babatunde50/gitversity-cli/token"
	"google.golang.org/grpc/metadata"
)

func GetRepo(repoId string) (*ggs.GetRepoResponse, error) {

	conn, err := api.Setup(config.C.Services["git_grpc"].Host)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	gitService := ggs.NewGitServiceClient(conn)

	jwtToken, err := token.LoadToken()

	if err != nil {
		return nil, err
	}

	authCtx := metadata.AppendToOutgoingContext(context.Background(), "authorization", fmt.Sprintf("Bearer %s", jwtToken))

	repoResp, err := gitService.GetRepo(authCtx, &ggs.GetRepoRequest{
		RepoId: repoId,
	})

	if err != nil {
		return nil, err
	}

	return repoResp, nil

}

func GetCommitFiles(commitId string) (*ggs.GetCommitFilesResponse, error) {

	conn, err := api.Setup(config.C.Services["git_grpc"].Host)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	gitService := ggs.NewGitServiceClient(conn)

	jwtToken, err := token.LoadToken()

	if err != nil {
		return nil, err
	}

	authCtx := metadata.AppendToOutgoingContext(context.Background(), "authorization", fmt.Sprintf("Bearer %s", jwtToken))

	commitFilesResp, err := gitService.GetCommitFiles(authCtx, &ggs.GetCommitFilesRequest{
		CommitId: commitId,
	})

	if err != nil {
		return nil, err
	}

	return commitFilesResp, nil

}

func GetBlob(sha1 string) (*ggs.GetBlobResponse, error) {
	conn, err := api.Setup(config.C.Services["git_grpc"].Host)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	gitService := ggs.NewGitServiceClient(conn)

	jwtToken, err := token.LoadToken()

	if err != nil {
		return nil, err
	}

	authCtx := metadata.AppendToOutgoingContext(context.Background(), "authorization", fmt.Sprintf("Bearer %s", jwtToken))

	blobResp, err := gitService.GetBlob(authCtx, &ggs.GetBlobRequest{
		Sha1: sha1,
	})

	if err != nil {
		return blobResp, err
	}

	return blobResp, nil
}

func Stage(repoId string, files []*ggs.AppFile) error {

	conn, err := api.Setup(config.C.Services["git_grpc"].Host)
	if err != nil {
		return err
	}

	defer conn.Close()

	gitService := ggs.NewGitServiceClient(conn)

	jwtToken, err := token.LoadToken()

	if err != nil {
		return err
	}

	authCtx := metadata.AppendToOutgoingContext(context.Background(), "authorization", fmt.Sprintf("Bearer %s", jwtToken))

	_, err = gitService.StageToRepo(authCtx, &ggs.StageToRepoRequest{
		RepoId: repoId,
		Files:  files,
	})

	if err != nil {
		return err
	}

	return nil
}

func UnStage(repoId string, filePath []string) error {

	conn, err := api.Setup(config.C.Services["git_grpc"].Host)
	if err != nil {
		return err
	}

	defer conn.Close()

	gitService := ggs.NewGitServiceClient(conn)

	jwtToken, err := token.LoadToken()

	if err != nil {
		return err
	}

	authCtx := metadata.AppendToOutgoingContext(context.Background(), "authorization", fmt.Sprintf("Bearer %s", jwtToken))

	_, err = gitService.UnstageFromRepo(authCtx, &ggs.UnstageFromRepoRequest{
		RepoId:   repoId,
		FilePath: filePath,
	})

	if err != nil {
		return err
	}

	return nil
}

func Commit(repoId, message string) error {

	conn, err := api.Setup(config.C.Services["git_grpc"].Host)
	if err != nil {
		return err
	}

	defer conn.Close()

	gitService := ggs.NewGitServiceClient(conn)

	jwtToken, err := token.LoadToken()

	if err != nil {
		return err
	}

	authCtx := metadata.AppendToOutgoingContext(context.Background(), "authorization", fmt.Sprintf("Bearer %s", jwtToken))

	_, err = gitService.CommitToRepo(authCtx, &ggs.CommitToRepoRequest{
		RepoId:        repoId,
		CommitMessage: message,
	})

	if err != nil {
		return err
	}

	return nil

}

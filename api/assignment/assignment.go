package assignment

import (
	"context"
	"errors"
	"fmt"

	"time"

	"github.com/Babatunde50/gitversity-cli/api"
	"github.com/Babatunde50/gitversity-cli/config"
	"github.com/Babatunde50/gitversity-cli/services/gas"
	"github.com/Babatunde50/gitversity-cli/token"
	"google.golang.org/grpc/metadata"
)

var responseTime = 5 * time.Second

func GetAssignmentByInviteCode(inviteCode string) (*gas.Assignment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), responseTime)
	defer cancel()

	conn, err := api.Setup(config.C.Services["assignment_grpc"].Host)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	assignmentService := gas.NewAssignmentServiceClient(conn)

	assignment, err := assignmentService.GetAssignmentByInviteCode(ctx, &gas.GetAssignmentByInviteCodeRequest{
		InviteCode: inviteCode,
	})

	if err != nil {
		return assignment, err
	}

	// TODO: move this to the server..
	if assignment.Status != gas.AssignmentStatus_OPEN {
		return assignment, errors.New("assignment has not been opened")
	}

	return assignment, nil

}

func SubmitAssignment(id string) (*gas.Submission, error) {
	conn, err := api.Setup(config.C.Services["assignment_grpc"].Host)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	assignmentService := gas.NewAssignmentServiceClient(conn)

	jwtToken, err := token.LoadToken()

	if err != nil {
		return nil, err
	}

	authCtx := metadata.AppendToOutgoingContext(context.Background(), "authorization", fmt.Sprintf("Bearer %s", jwtToken))

	return assignmentService.SubmitAssignment(authCtx, &gas.SubmitAssignmentRequest{
		Id: id,
	})
}

func JoinAssignment(id string) (*gas.JoinAssignmentResponse, error) {

	conn, err := api.Setup(config.C.Services["assignment_grpc"].Host)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	assignmentService := gas.NewAssignmentServiceClient(conn)

	jwtToken, err := token.LoadToken()

	if err != nil {
		return nil, err
	}

	authCtx := metadata.AppendToOutgoingContext(context.Background(), "authorization", fmt.Sprintf("Bearer %s", jwtToken))

	return assignmentService.JoinAssignment(authCtx, &gas.JoinAssignmentRequest{
		Id: id,
	})
}

func GetUserAssignments(id string) (*gas.UserAssignments, error) {

	conn, err := api.Setup(config.C.Services["assignment_grpc"].Host)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	assignmentService := gas.NewAssignmentServiceClient(conn)

	jwtToken, err := token.LoadToken()

	if err != nil {
		return nil, err
	}

	authCtx := metadata.AppendToOutgoingContext(context.Background(), "authorization", fmt.Sprintf("Bearer %s", jwtToken))

	return assignmentService.GetUserAssignments(authCtx, &gas.GetAssignmentRequest{
		Id: id,
	})

}

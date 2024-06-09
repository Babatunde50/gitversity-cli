package user

import (
	"context"
	"time"

	"github.com/Babatunde50/gitversity-cli/api"
	"github.com/Babatunde50/gitversity-cli/config"
	"github.com/Babatunde50/gitversity-cli/services/gus"
	"github.com/Babatunde50/gitversity-cli/token"
	_ "github.com/golang/mock/mockgen/model"
)

var responseTime = 5 * time.Second

type UserApi interface {
	Login(email, password string) error
	SignUp(email, fullName, confirmPassword, password string) error
}

func Login(email, password string) error {

	ctx, cancel := context.WithTimeout(context.Background(), responseTime)
	defer cancel()

	conn, err := api.Setup(config.C.Services["user_grpc"].Host)
	if err != nil {
		return err
	}

	defer conn.Close()

	userServer := gus.NewUserServiceClient(conn)

	resp, err := userServer.LoginUser(ctx, &gus.LoginUserRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return err
	}

	if err = token.SaveToken(resp.AccessToken); err != nil {
		return err
	}

	return nil

}

func SignUp(email, fullName, confirmPassword, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), responseTime)
	defer cancel()

	conn, err := api.Setup(config.C.Services["user_grpc"].Host)
	if err != nil {
		return err
	}

	defer conn.Close()

	userServer := gus.NewUserServiceClient(conn)

	_, err = userServer.CreateUser(ctx, &gus.CreateUserRequest{
		FullName:        fullName,
		Email:           email,
		Password:        password,
		PasswordConfirm: confirmPassword,
		Roles:           []string{"USER"},
	})

	if err != nil {
		return err
	}

	return nil
}

// pkg/hello_test.go
package pkg

// import (
// 	"errors"
// 	"testing"

// 	"github.com/Babatunde50/gitversity-cli/cli"
// 	"github.com/golang/mock/gomock"
// )

// func TestRunLogin(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockCli := cli.NewMockCli(ctrl)

// 	tests := []struct {
// 		name     string
// 		email    string
// 		password string
// 		mockErr  error
// 		wantErr  bool
// 	}{
// 		{
// 			name:     "Successful Login",
// 			email:    "test@example.com",
// 			password: "password",
// 			mockErr:  nil,
// 			wantErr:  false,
// 		},
// 		{
// 			name:     "Invalid Credentials",
// 			email:    "invalid@example.com",
// 			password: "wrongpassword",
// 			mockErr:  errors.New("invalid credentials"),
// 			wantErr:  true,
// 		},
// 		{
// 			name:     "Empty Email",
// 			email:    "",
// 			password: "password",
// 			mockErr:  errors.New("email is required"),
// 			wantErr:  true,
// 		},
// 		{
// 			name:     "Empty Password",
// 			email:    "test@example.com",
// 			password: "",
// 			mockErr:  errors.New("password is required"),
// 			wantErr:  true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Set up the expectation
// 			mockCli.EXPECT().Login(tt.email, tt.password).Return(tt.mockErr)

// 			// Call the function under test
// 			err := RunLogin(mockCli, tt.email, tt.password)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RunLogin() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

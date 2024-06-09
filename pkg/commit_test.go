package pkg

// import (
// 	"errors"
// 	"testing"

// 	"github.com/Babatunde50/gitversity-cli/cli"
// 	"github.com/golang/mock/gomock"
// )

// func TestRunCommit(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockCli := cli.NewMockCli(ctrl)

// 	tests := []struct {
// 		name      string
// 		repoId    string
// 		message   string
// 		wantErr   bool
// 		setupMock func()
// 	}{
// 		{
// 			name:    "Successful Commit",
// 			repoId:  "repo123",
// 			message: "Initial commit",
// 			wantErr: false,
// 			setupMock: func() {
// 				mockCli.EXPECT().Commit("repo123", "Initial commit").Return(nil)
// 			},
// 		},
// 		{
// 			name:    "Commit Error",
// 			repoId:  "repo123",
// 			message: "Initial commit",
// 			wantErr: true,
// 			setupMock: func() {
// 				mockCli.EXPECT().Commit("repo123", "Initial commit").Return(errors.New("commit failed"))
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setupMock()

// 			err := RunCommit(mockCli, tt.repoId, tt.message)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RunCommit() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

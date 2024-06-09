// pkg/hello_test.go
package pkg

// import (
// 	"errors"
// 	"os"
// 	"testing"

// 	"github.com/Babatunde50/gitversity-cli/cli"
// 	"github.com/Babatunde50/gitversity-cli/services/ggs"
// 	"github.com/golang/mock/gomock"
// )

// func TestRunClone(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockCli := cli.NewMockCli(ctrl)

// 	tests := []struct {
// 		name       string
// 		repoId     string
// 		repoResp   *ggs.GetRepoResponse
// 		repoErr    error
// 		blobErr    error
// 		wantErr    bool
// 		setupMocks func()
// 	}{
// 		{
// 			name:   "Successful Clone",
// 			repoId: "123",
// 			repoResp: &ggs.GetRepoResponse{
// 				Repo: &ggs.Repo{
// 					Name: "testrepo",
// 					Id:   "123",
// 					Index: &ggs.Index{
// 						Entries: []*ggs.IndexEntry{
// 							{Path: "file1.txt", Sha1: "sha1file1"},
// 						},
// 					},
// 				},
// 			},
// 			repoErr: nil,
// 			blobErr: nil,
// 			wantErr: false,
// 			setupMocks: func() {
// 				mockCli.EXPECT().GetRepo("123").Return(&ggs.GetRepoResponse{
// 					Repo: &ggs.Repo{
// 						Name: "testrepo",
// 						Id:   "123",
// 						Index: &ggs.Index{
// 							Entries: []*ggs.IndexEntry{
// 								{Path: "file1.txt", Sha1: "sha1file1"},
// 							},
// 						},
// 					},
// 				}, nil)
// 				mockCli.EXPECT().GetBlob("sha1file1").Return(&ggs.GetBlobResponse{
// 					Content: "content of file1",
// 				}, nil)
// 			},
// 		},
// 		{
// 			name:    "Repo Fetch Error",
// 			repoId:  "123",
// 			repoErr: errors.New("repo not found"),
// 			wantErr: true,
// 			setupMocks: func() {
// 				mockCli.EXPECT().GetRepo("123").Return(nil, errors.New("repo not found"))
// 			},
// 		},
// 		{
// 			name:   "Blob Fetch Error",
// 			repoId: "123",
// 			repoResp: &ggs.GetRepoResponse{
// 				Repo: &ggs.Repo{
// 					Name: "testrepo",
// 					Id:   "123",
// 					Index: &ggs.Index{
// 						Entries: []*ggs.IndexEntry{
// 							{Path: "file1.txt", Sha1: "sha1file1"},
// 						},
// 					},
// 				},
// 			},
// 			repoErr: nil,
// 			blobErr: errors.New("blob not found"),
// 			wantErr: true,
// 			setupMocks: func() {
// 				mockCli.EXPECT().GetRepo("123").Return(&ggs.GetRepoResponse{
// 					Repo: &ggs.Repo{
// 						Name: "testrepo",
// 						Id:   "123",
// 						Index: &ggs.Index{
// 							Entries: []*ggs.IndexEntry{
// 								{Path: "file1.txt", Sha1: "sha1file1"},
// 							},
// 						},
// 					},
// 				}, nil)
// 				mockCli.EXPECT().GetBlob("sha1file1").Return(nil, errors.New("blob not found"))
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Setup mocks for the current test case
// 			tt.setupMocks()

// 			// Call the function under test
// 			err := RunClone(mockCli, tt.repoId)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RunClone() error = %v, wantErr %v", err, tt.wantErr)
// 			}

// 			// Clean up created files and directories if the test was successful
// 			if !tt.wantErr {
// 				os.RemoveAll(tt.repoResp.Repo.Name)
// 			}
// 		})
// 	}
// }

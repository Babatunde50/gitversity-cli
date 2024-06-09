package pkg

// import (
// 	"errors"
// 	"io/fs"
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"time"

// 	"github.com/Babatunde50/gitversity-cli/cli"
// 	"github.com/Babatunde50/gitversity-cli/services/ggs"
// 	"github.com/golang/mock/gomock"
// )

// // TODO: Make this test pass

// func TestRunStage(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockCli := cli.NewMockCli(ctrl)

// 	// Create a temporary directory to simulate the file system
// 	tempDir := t.TempDir()

// 	// Create some test files
// 	file1Path := filepath.Join(tempDir, "file1.txt")
// 	file2Path := filepath.Join(tempDir, "file2.log")
// 	os.WriteFile(file1Path, []byte("content of file1"), 0644)
// 	os.WriteFile(file2Path, []byte("content of file2"), 0644)

// 	modTime := time.Now().Format(time.RFC3339)

// 	tests := []struct {
// 		name      string
// 		repoId    string
// 		args      []string
// 		wantErr   bool
// 		setupMock func()
// 	}{
// 		{
// 			name:    "Successful Stage",
// 			repoId:  "repo123",
// 			args:    []string{"*.txt"},
// 			wantErr: false,
// 			setupMock: func() {
// 				expectedFiles := []*ggs.AppFile{
// 					{
// 						Path:     "file1.txt",
// 						Content:  "content of file1",
// 						Mode:     int32(fs.FileMode(0644)),
// 						ModeTime: modTime,
// 						CTime:    modTime,
// 					},
// 				}
// 				mockCli.EXPECT().Stage("repo123", expectedFiles).Return(nil)
// 			},
// 		},
// 		{
// 			name:    "Stage Error",
// 			repoId:  "repo123",
// 			args:    []string{"*.txt"},
// 			wantErr: true,
// 			setupMock: func() {
// 				expectedFiles := []*ggs.AppFile{
// 					{
// 						Path:     "file1.txt",
// 						Content:  "content of file1",
// 						Mode:     int32(fs.FileMode(0644)),
// 						ModeTime: modTime,
// 						CTime:    modTime,
// 					},
// 				}
// 				mockCli.EXPECT().Stage("repo123", expectedFiles).Return(errors.New("stage failed"))
// 			},
// 		},
// 		{
// 			name:    "No Matching Files",
// 			repoId:  "repo123",
// 			args:    []string{"*.md"},
// 			wantErr: false,
// 			setupMock: func() {
// 				mockCli.EXPECT().Stage("repo123", []*ggs.AppFile{}).Return(nil)
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setupMock()

// 			err := RunStage(mockCli, tt.repoId, tt.args)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RunStage() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

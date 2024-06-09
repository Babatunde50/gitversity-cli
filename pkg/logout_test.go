package pkg

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestRunLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Define a mock for the deleteTokenFunc
	var mockDeleteTokenFunc func() error

	// Override the deleteTokenFunc with the mock
	deleteTokenFunc = func() error {
		return mockDeleteTokenFunc()
	}

	tests := []struct {
		name      string
		wantErr   bool
		setupMock func()
	}{
		{
			name:    "Successful Logout",
			wantErr: false,
			setupMock: func() {
				mockDeleteTokenFunc = func() error {
					return nil
				}
			},
		},
		{
			name:    "Logout Error",
			wantErr: true,
			setupMock: func() {
				mockDeleteTokenFunc = func() error {
					return errors.New("failed to delete token")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			err := RunLogout(nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunLogout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package pkg

// Mock the survey.AskOne function
// var mockSurveyAskOne = func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
// 	switch p {
// 	case loginOrSignupPrompt:
// 		*response.(*bool) = true // Simulate user has an account
// 	case emailPrompt:
// 		*response.(*string) = "test@example.com"
// 	case passwordPrompt:
// 		*response.(*string) = "password"
// 	case fullNamePrompt:
// 		*response.(*string) = "John Doe"
// 	case confirmPasswordPrompt:
// 		*response.(*string) = "password"
// 	}
// 	return nil
// }

// // TODO: Make this test to pass

// func TestRunJoin(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockCli := cli.NewMockCli(ctrl)

// 	// Override the askOne function with the mock
// 	askOne = mockSurveyAskOne
// 	defer func() { askOne = survey.AskOne }() // Ensure we reset it after the test

// 	tests := []struct {
// 		name              string
// 		inviteCode        string
// 		autoClone         bool
// 		tokenLoadErr      error
// 		loginErr          error
// 		signUpErr         error
// 		getRepoErr        error
// 		getAssignment     *gas.Assignment
// 		getAssignmentErr  error
// 		joinAssignment    *gas.JoinAssignmentResponse
// 		joinAssignmentErr error
// 		wantErr           bool
// 		setupMocks        func()
// 	}{
// 		{
// 			name:              "Successful Join with Auto Clone",
// 			inviteCode:        "invite123",
// 			autoClone:         true,
// 			tokenLoadErr:      errors.New("token not found"),
// 			loginErr:          nil,
// 			signUpErr:         nil,
// 			getRepoErr:        nil,
// 			getAssignment:     &gas.Assignment{Id: "assignment123"},
// 			getAssignmentErr:  nil,
// 			joinAssignment:    &gas.JoinAssignmentResponse{RepoId: "repo123"},
// 			joinAssignmentErr: nil,
// 			wantErr:           false,
// 			setupMocks: func() {
// 				mockCli.EXPECT().GetAssignmentByInviteCode("invite123").Return(&gas.Assignment{Id: "assignment123"}, nil)
// 				mockCli.EXPECT().Login("test@example.com", "password").Return(nil)
// 				mockCli.EXPECT().SignUp("test@example.com", "John Doe", "password", "password").Return(nil)
// 				mockCli.EXPECT().JoinAssignment("assignment123").Return(&gas.JoinAssignmentResponse{RepoId: "repo123"}, nil)
// 				mockCli.EXPECT().GetRepo("repo123").Return(&ggs.GetRepoResponse{Repo: &ggs.Repo{Name: "testrepo", Id: "repo123", Index: &ggs.Index{Entries: []*ggs.IndexEntry{{Path: "file1.txt", Sha1: "sha1file1"}}}}}, nil)
// 				mockCli.EXPECT().GetBlob("sha1file1").Return(&ggs.GetBlobResponse{Content: "content of file1"}, nil)
// 			},
// 		},
// 		{
// 			name:              "Assignment Fetch Error",
// 			inviteCode:        "invalidInvite",
// 			autoClone:         false,
// 			tokenLoadErr:      nil,
// 			loginErr:          nil,
// 			signUpErr:         nil,
// 			getRepoErr:        nil,
// 			getAssignment:     nil,
// 			getAssignmentErr:  errors.New("assignment not found"),
// 			joinAssignment:    nil,
// 			joinAssignmentErr: nil,
// 			wantErr:           true,
// 			setupMocks: func() {
// 				mockCli.EXPECT().GetAssignmentByInviteCode("invalidInvite").Return(nil, errors.New("assignment not found"))
// 			},
// 		},
// 		{
// 			name:              "Join Assignment Error",
// 			inviteCode:        "invite123",
// 			autoClone:         false,
// 			tokenLoadErr:      nil,
// 			loginErr:          nil,
// 			signUpErr:         nil,
// 			getRepoErr:        nil,
// 			getAssignment:     &gas.Assignment{Id: "assignment123"},
// 			getAssignmentErr:  nil,
// 			joinAssignment:    nil,
// 			joinAssignmentErr: errors.New("failed to join assignment"),
// 			wantErr:           true,
// 			setupMocks: func() {
// 				mockCli.EXPECT().GetAssignmentByInviteCode("invite123").Return(&gas.Assignment{Id: "assignment123"}, nil)
// 				mockCli.EXPECT().JoinAssignment("assignment123").Return(nil, errors.New("failed to join assignment"))
// 			},
// 		},
// 		{
// 			name:              "Signup Error",
// 			inviteCode:        "invite123",
// 			autoClone:         false,
// 			tokenLoadErr:      errors.New("token not found"),
// 			loginErr:          nil,
// 			signUpErr:         errors.New("signup failed"),
// 			getRepoErr:        nil,
// 			getAssignment:     nil,
// 			getAssignmentErr:  nil,
// 			joinAssignment:    nil,
// 			joinAssignmentErr: nil,
// 			wantErr:           true,
// 			setupMocks: func() {
// 				mockCli.EXPECT().GetAssignmentByInviteCode("invite123").Return(nil, nil)
// 				mockCli.EXPECT().SignUp("test@example.com", "John Doe", "password", "password").Return(errors.New("signup failed"))
// 			},
// 		},
// 		{
// 			name:              "Load Token Error",
// 			inviteCode:        "invite123",
// 			autoClone:         false,
// 			tokenLoadErr:      errors.New("token not found"),
// 			loginErr:          nil,
// 			signUpErr:         nil,
// 			getRepoErr:        nil,
// 			getAssignment:     nil,
// 			getAssignmentErr:  nil,
// 			joinAssignment:    nil,
// 			joinAssignmentErr: nil,
// 			wantErr:           true,
// 			setupMocks: func() {
// 				mockCli.EXPECT().GetAssignmentByInviteCode("invite123").Return(nil, nil)
// 				mockCli.EXPECT().Login("test@example.com", "password").Return(nil)
// 				mockCli.EXPECT().SignUp("test@example.com", "John Doe", "password", "password").Return(nil)
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Setup mocks for the current test case
// 			tt.setupMocks()

// 			// Call the function under test
// 			err := RunJoin(mockCli, tt.inviteCode, tt.autoClone)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RunJoin() error = %v, wantErr %v", err, tt.wantErr)
// 			}

// 			// Additional checks and clean up if necessary can be added here
// 		})
// 	}
// }

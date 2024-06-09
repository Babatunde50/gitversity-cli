package config

// Token represents token-related configuration.
type Token struct {
	FileName   string
	AppName    string
	FolderName string
}

// Service represents a service configuration.
type Service struct {
	Host string
}

// Config holds the entire configuration structure.
type Config struct {
	Token    Token
	Services map[string]Service
}

// C is the global configuration variable.
var C Config

// LoadConfig initializes the configuration with hardcoded values.
func LoadConfig() {
	C = Config{
		Token: Token{
			FileName:   "token",
			AppName:    "gitversity",
			FolderName: ".config",
		},
		Services: map[string]Service{
			"user_grpc": {
				Host: "grpc.dev.user.api.gitversity.com:443",
			},
			"assignment_grpc": {
				Host: "grpc.dev.assignment.api.gitversity.com:443",
			},
			"git_grpc": {
				Host: "grpc.dev.git.api.gitversity.com:443",
			},
		},
	}
}

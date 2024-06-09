// pkg/hello.go
package pkg

import (
	"github.com/Babatunde50/gitversity-cli/cli"
)

func RunLogin(cli cli.Cli, email, password string) error {
	return cli.Login(email, password)
}

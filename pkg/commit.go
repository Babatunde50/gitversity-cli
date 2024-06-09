// pkg/hello.go
package pkg

import (
	"github.com/Babatunde50/gitversity-cli/cli"
)

func RunCommit(cli cli.Cli, repoId, message string) error {
	return cli.Commit(repoId, message)
}

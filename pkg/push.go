// pkg/hello.go
package pkg

import (
	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/metadata"
	"github.com/Babatunde50/gitversity-cli/services/gas"
)

func RunPush(cli cli.Cli) (*gas.Submission, error) {
	md, err := metadata.Read()
	if err != nil {
		return &gas.Submission{}, err
	}

	return cli.Push(md.AssignmentId)

}

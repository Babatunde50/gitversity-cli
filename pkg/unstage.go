// pkg/hello.go
package pkg

import (
	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/metadata"
)

// TODO: do not consider .gitversity file

func RunUnstage(cli cli.Cli, args []string) error {

	md, err := metadata.Read()

	if err != nil {
		return err
	}

	files, err := getFileDetails(".", args)
	if err != nil {
		return err
	}

	var paths []string
	for _, file := range files {
		paths = append(paths, file.Path)
	}

	if err = cli.UnStage(md.RepoId, paths); err != nil {
		return err
	}

	// TODO (fix) : after reset gitversity status still returning same result..

	return nil
}

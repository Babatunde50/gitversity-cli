package pkg

import (
	"github.com/Babatunde50/gitversity-cli/cli"
	"github.com/Babatunde50/gitversity-cli/token"
)

var deleteTokenFunc = token.DeleteToken

func RunLogout(cli cli.Cli) error {
	return deleteTokenFunc()
}

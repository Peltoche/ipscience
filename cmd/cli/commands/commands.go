package commands

import (
	"github.com/teris-io/cli"
)

func NewApp() cli.App {

	return cli.New("ipscience").
		WithCommand(FetchCmd()).
		WithCommand(MetadatasCmd())
}

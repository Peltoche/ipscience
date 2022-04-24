package commands

import "github.com/teris-io/cli"

func FetchMetadatas() cli.Command {
	return cli.NewCommand("metadatas",
		"Fetch the metadatas for a given DOI").
		WithArg(cli.NewArg("DOI", "A valid DOI")).
		WithAction(fetchMetadatasAction)
}

func fetchMetadatasAction(args []string, options map[string]string) int {
	return 0
}

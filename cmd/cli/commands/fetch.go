package commands

import (
	"fmt"
	"os"

	"github.com/Peltoche/ipscience/pkg/documents"
	"github.com/teris-io/cli"
)

func FetchCmd() cli.Command {
	return cli.NewCommand("fetch", "Fetch a work (paper/article) from a DOI").
		WithArg(cli.NewArg("DOI", "Document Object Identifier").WithType(cli.TypeString)).
		WithArg(cli.NewArg("target", "The path were the file will be saved").WithType(cli.TypeString)).
		WithAction(fetchDocumentAction)
}

func fetchDocumentAction(args []string, options map[string]string) int {
	client := documents.NewScihubClient()

	rawFile, err := client.FetchDocument(args[0])
	if err != nil {
		fmt.Printf("failed to fetch the document: %s\n", err)
		return 1
	}

	err = os.WriteFile(args[1], rawFile, 0666)
	if err != nil {
		fmt.Printf("failed to write the file: %s\n", err)
		return 1
	}

	return 0
}

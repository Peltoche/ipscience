package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Peltoche/ipscience/pkg/metadatas"
	"github.com/teris-io/cli"
)

func WorkCmd() cli.Command {
	return cli.NewCommand("work", "Fetch a work (paper/article) from a DOI").
		WithArg(cli.NewArg("DOI", "Document Object Identifier").WithType(cli.TypeString)).
		WithOption(cli.NewOption("metadatas", "Fetch the metadatas").WithChar('m').WithType(cli.TypeBool)).
		WithAction(workAction)
}

func workAction(args []string, options map[string]string) int {
	ctx := context.Background()

	if _, ok := options["metadatas"]; ok {
		return fetchMetadatas(ctx, args[0])
	}

	return 0
}

func fetchMetadatas(ctx context.Context, doi string) int {
	client := metadatas.NewCrossRefClient()

	work, err := client.FetchWork(ctx, doi)
	if err != nil {
		fmt.Printf("failed to fetch the metadatas: %s\n", err)
		return 1
	}

	err = json.NewEncoder(os.Stdout).Encode(work)
	if err != nil {
		fmt.Printf("failed to print into stdout: %s\n", err)
		return 1
	}

	return 0
}

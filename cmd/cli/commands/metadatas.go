package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Peltoche/ipscience/pkg/metadatas"
	"github.com/teris-io/cli"
)

func MetadatasCmd() cli.Command {
	return cli.NewCommand("metadata", "Fetch the metadatas for a given DOI").
		WithArg(cli.NewArg("DOI", "Document Object Identifier").WithType(cli.TypeString)).
		WithAction(fetchMetadatasAction)
}

func fetchMetadatasAction(args []string, options map[string]string) int {
	ctx := context.Background()

	client := metadatas.NewCrossRefClient()

	work, err := client.FetchWork(ctx, args[0])
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

package main

import (
	"os"

	"github.com/Peltoche/ipscience/cmd/cli/commands"
)

func main() {
	app := commands.NewApp()

	os.Exit(app.Run(os.Args, os.Stdout))
}

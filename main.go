package main

import (
	"os"

	"github.com/alecthomas/kong"

	"github.com/0xmzn/awlist/internal/cli"
)

var version = "0.1.0"

type App struct {
	cli.Globals
	Generate cli.GenerateCmd `cmd:"" help:"generate file from template"`
}

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "--help")
	}

	var app App
	parser := kong.Must(&app,
		kong.Name("awlist"),
		kong.Description("A CLI tool for managing awesome-lists"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Summary: true,
		}),
		kong.Vars{"version": version},
	)

	ctx, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	err = ctx.Run()
	ctx.FatalIfErrorf(err)
}

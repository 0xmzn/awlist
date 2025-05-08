package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"

	"github.com/0xmzn/awlist/internal/cli"
)

var version = "0.1.0"

type App struct {
	Version  kong.VersionFlag `short:"V" help:"Show application version."`
	Generate cli.GenerateCmd `cmd:"" help:"generate file from template"`
}

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "--help")
	}

	var app App
	parser, err := kong.New(&app,
		kong.Name("awlist"),
		kong.Description("A CLI tool for managing awesome-lists"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Summary: true,
		}),
		kong.Vars{"version": version},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating parser:", err)
		os.Exit(1)
	}

	ctx, err := parser.Parse(os.Args[1:])
	if err != nil {
		parser.Printf("%v", err)
		os.Exit(1)
	}

	err = ctx.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package cli

import (
	"fmt"
	htmltemplate "html/template"
	"io"
	"os"
	texttemplate "text/template"

	"github.com/0xmzn/awelist/internal/awesomestore"
)

type GenerateCmd struct {
	Globals    *Globals `kong:"embed"`
	InputFile  string   `arg:"" required:"" help:"Input template file"`
	OutputFile string   `arg:"" optional:"" help:"Output file to write the generated result. Defaults to stdout."`
	Html       bool     `kong:"optional,name='html',help='Generate HTML output.'"`
}

func (cmd *GenerateCmd) Run() error {
	inputFileContent, err := os.ReadFile(cmd.InputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	dataPath, err := GetDataFilePath(cmd.Globals.DataFile)
	if err != nil {
		return fmt.Errorf("failed to get data file path: %w", err)
	}

	store, err := awesomestore.NewStore(dataPath)
	if err != nil {
		return fmt.Errorf("failed to load data from %s: %w", dataPath, err)
	}

	awesomeData := store.Data()

	var outputWriter io.Writer
	if cmd.OutputFile == "" {
		outputWriter = os.Stdout
	} else {
		file, err := os.Create(cmd.OutputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file %s: %w", cmd.OutputFile, err)
		}
		defer file.Close()
		outputWriter = file
	}

	if cmd.Html {
		tmpl := htmltemplate.New("user-template-html").Option("missingkey=error")
		tmpl, err = tmpl.Parse(string(inputFileContent))
		if err != nil {
			return fmt.Errorf("html template parsing error: %w", err)
		}
		err = tmpl.Execute(outputWriter, awesomeData)
		if err != nil {
			return fmt.Errorf("failed to execute html template: %w", err)
		}
	} else {
		tmpl := texttemplate.New("user-template-text").Option("missingkey=error")
		tmpl, err = tmpl.Parse(string(inputFileContent))
		if err != nil {
			return fmt.Errorf("text template parsing error: %w", err)
		}
		err = tmpl.Execute(outputWriter, awesomeData)
		if err != nil {
			return fmt.Errorf("failed to execute text template: %w", err)
		}
	}

	return nil
}

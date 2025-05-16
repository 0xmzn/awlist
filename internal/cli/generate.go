package cli

import (
	"fmt"
	htmltemplate "html/template"
	"io"
	"os"
	texttemplate "text/template"

	"github.com/0xmzn/awelist/internal/awesomestore"
	"github.com/0xmzn/awelist/internal/model"
)

type GenerateCmd struct {
	Globals    *Globals `kong:"embed"`
	InputFile  string   `arg:"" required:"" help:"Input template file"`
	OutputFile string   `arg:"" optional:"" help:"Output file to write the generated result. Defaults to stdout."`
	Html       bool     `kong:"optional,name='html',help='Generate HTML output.'"`
}

func loadAndEnrichData(dataFile string) (*model.AwesomeData, error) {
	dataPath, err := GetDataFilePath(dataFile)
	if err != nil {
		return nil, fmt.Errorf("failed to get data file path: %w", err)
	}

	store, err := awesomestore.NewStore(dataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load data from %s: %w", dataPath, err)
	}

	awesomeData := store.Data()
	model.Enrich(&awesomeData)
	return &awesomeData, nil
}

func getOutputWriter(outputFilePath string) (writer io.Writer, closer func() error, err error) {
	if outputFilePath == "" {
		return os.Stdout, nil, nil
	}
	file, errCreate := os.Create(outputFilePath)
	if errCreate != nil {
		return nil, nil, fmt.Errorf("failed to create output file %s: %w", outputFilePath, errCreate)
	}
	return file, file.Close, nil
}

func executeHTMLTemplate(writer io.Writer, templateContent string, data *model.AwesomeData) error {
	tmpl := htmltemplate.New("user-template-html").Option("missingkey=error")
	parsedTmpl, err := tmpl.Parse(templateContent)
	if err != nil {
		return fmt.Errorf("html template parsing error: %w", err)
	}
	err = parsedTmpl.Execute(writer, data)
	if err != nil {
		return fmt.Errorf("failed to execute html template: %w", err)
	}
	return nil
}

func executeTextTemplate(writer io.Writer, templateContent string, data *model.AwesomeData) error {
	tmpl := texttemplate.New("user-template-text").Option("missingkey=error")
	parsedTmpl, err := tmpl.Parse(templateContent)
	if err != nil {
		return fmt.Errorf("text template parsing error: %w", err)
	}
	err = parsedTmpl.Execute(writer, data)
	if err != nil {
		return fmt.Errorf("failed to execute text template: %w", err)
	}
	return nil
}

func (cmd *GenerateCmd) Run() error {
	inputFileContent, err := os.ReadFile(cmd.InputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	awesomeData, err := loadAndEnrichData(cmd.Globals.DataFile)
	if err != nil {
		return err
	}

	outputWriter, closeFile, err := getOutputWriter(cmd.OutputFile)
	if err != nil {
		return err
	}
	if closeFile != nil {
		defer closeFile()
	}

	templateStr := string(inputFileContent)
	if cmd.Html {
		err = executeHTMLTemplate(outputWriter, templateStr, awesomeData)
	} else {
		err = executeTextTemplate(outputWriter, templateStr, awesomeData)
	}

	if err != nil {
		return err
	}

	return nil
}

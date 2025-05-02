package cli

type GenerateCmd struct {
    InputFile  string `arg:"" required:"" help:"Input template file"`
    OutputFile string `arg:"" required:"" help:"Output file to write the generated result"`
	Html bool `help:"Use html mode"`
}

func (cmd *GenerateCmd) Run() error {
	return nil
}

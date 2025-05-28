package cli

type Globals struct {
	DataFile string `kong:"optional,name='data',short='d',help='Path to the YAML data file. Defaults to awesome.yaml in the current directory if present.'"`
}

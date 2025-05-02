package cli

import "github.com/alecthomas/kong"

type Globals struct {
    Version kong.VersionFlag `short:"V" help:"Show application version."`
}

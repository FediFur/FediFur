package main

import (
	"github.com/urfave/cli/v2"

	"github.com/FediFur/fedifur/cmd"
)

// These are set by build flags
var (
	Version  = ""
	Commit   = ""
	Compiled = ""
)

func main() {
	app := cmd.NewApp(cmd.AppVersion{Version: Version, Commit: Commit, Compiled: Compiled})

	err := cmd.RunApp(app)
	if err != nil {
		cli.OsExiter(1)
	}
}

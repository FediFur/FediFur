package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
)

type AppVersion struct {
	Version  string
	Commit   string
	Compiled string
}

func NewApp(version AppVersion) *cli.App {
	app := cli.NewApp()
	app.Name = "fedifur"
	app.Usage = "usage"
	app.Version = version.Version
	app.Flags = append(app.Flags, cli.VersionFlag)

	cli.VersionPrinter = func(cCtx *cli.Context) {
		template := `
Version:    %s
Go Version: %s
Git Commit: %s
Compiled:   %s
OS/Arch:    %s/%s
`

		text := fmt.Sprintf(
			template,

			app.Version,
			runtime.Version(),
			version.Commit,
			version.Compiled,
			runtime.GOOS,
			runtime.GOARCH,
		)

		fmt.Println(text)
	}

	return app
}

func RunApp(app *cli.App) error {
	err := app.Run(os.Args)
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"github.com/chemiq/notification/cmd/migrate"
	"github.com/chemiq/notification/cmd/server"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "server",
				Aliases:     []string{"s"},
				Description: "Run application to handle endpoints",
				Action:      server.Command,
			},
			{
				Name:        "migrate",
				Aliases:     []string{"m"},
				Description: "Execute database migrations",
				Action:      migrate.Command,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

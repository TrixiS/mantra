package main

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/TrixiS/mantra/pkg/command_context"
	"github.com/TrixiS/mantra/pkg/commands"
	"github.com/asdine/storm/v3"
	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

// TODO: subcommands file which would contain functions for each subcommand (or structs?)
// TODO: add/remove commands to add ssh connections
// TODO: aliases like mantra a / mantra rm
// TODO: mantra list
// TODO: mantra conn [id or name of a connection]
// TODO: boltdb as a data store

const DBFilename = "mantra.db"

func main() {
	if err := clipboard.Init(); err != nil {
		log.Fatal(err)
	}

	exPath, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	db, err := storm.Open(path.Join(filepath.Dir(exPath), DBFilename))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	contextProvider := &command_context.CommandContextProvider{
		DB: db,
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add new connection",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "host",
						Required: true,
					},
					&cli.UintFlag{
						Name:  "port",
						Value: 22,
					},
					&cli.StringFlag{
						Name:  "user",
						Value: "root",
					},
					&cli.StringFlag{
						Name:     "password",
						Required: true,
					},
				},
				Action: contextProvider.Wraps(commands.Add),
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all connections",
				Action:  contextProvider.Wraps(commands.List),
			},
			{
				Name:    "remove",
				Aliases: []string{"rm"},
				Usage:   "Remove a connection",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "id",
						Required: true,
					},
				},
				Action: contextProvider.Wraps(commands.Remove),
			},
			// TODO: store id flag in a var
			{
				Name:    "connect",
				Aliases: []string{"conn", "c"},
				Usage:   "Connect to SSH",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "id",
						Required: true,
					},
				},
				Action: contextProvider.Wraps(commands.Connect),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

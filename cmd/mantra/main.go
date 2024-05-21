package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/TrixiS/mantra/internal/command_context"
	"github.com/TrixiS/mantra/internal/commands"
	"github.com/asdine/storm/v3"
	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

const MantraDirname = ".mantra"
const MantraDBFilename = "mantra.db"

func main() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	mantraDirpath := path.Join(homeDir, MantraDirname)

	if err := os.MkdirAll(mantraDirpath, 0777); err != nil {
		log.Fatal(err)
	}

	if err := clipboard.Init(); err != nil {
		log.Fatal(fmt.Errorf("clipboard init %w", err))
	}

	contextProvider := &command_context.CommandContextProvider{
		DBFactory: func() *storm.DB {
			dbFilepath := path.Join(mantraDirpath, MantraDBFilename)
			db, err := storm.Open(dbFilepath)

			if err != nil {
				log.Fatal(fmt.Errorf("open db %w", err))
			}

			return db
		},
	}

	connectionIDFlag := &cli.IntFlag{
		Name:     "id",
		Required: true,
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
					&cli.StringFlag{
						Name:     "args",
						Required: false,
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
				Flags:   []cli.Flag{connectionIDFlag},
				Action:  contextProvider.Wraps(commands.Remove),
			},
			{
				Name:    "connect",
				Aliases: []string{"conn", "c"},
				Usage:   "Connect to SSH",
				Args:    true,
				Action:  contextProvider.Wraps(commands.Connect),
			},
			{
				Name:    "reveal",
				Aliases: []string{"pwd", "r"},
				Usage:   "Copy connection password",
				Flags:   []cli.Flag{connectionIDFlag},
				Action:  contextProvider.Wraps(commands.Reveal),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

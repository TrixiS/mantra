package command_context

import (
	"github.com/asdine/storm/v3"
	"github.com/urfave/cli/v2"
)

type CommandContextProvider struct {
	DBFactory func() *storm.DB
}

func (provider *CommandContextProvider) Wraps(f CommandFunc) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		commandContext := CommandContext{
			Provider:   provider,
			CLIContext: ctx,
		}

		return f(&commandContext)
	}
}

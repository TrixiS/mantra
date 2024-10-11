package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TrixiS/mantra/internal/mantra/command_context"
	"github.com/TrixiS/mantra/internal/mantra/models"
	"github.com/urfave/cli/v2"
)

// TODO: pretty output

type syncArgs struct {
	address  string
	username string
	password string
}

func syncArgsFromCLIArgs(ctx *cli.Context) syncArgs {
	return syncArgs{
		address:  ctx.String("address"),
		username: ctx.String("username"),
		password: ctx.String("password"),
	}
}

func Push(ctx *command_context.CommandContext) error {
	db := ctx.Provider.DBFactory()
	defer db.Close()

	var connectionModels []models.Connection

	if err := db.All(&connectionModels); err != nil {
		return err
	}

	if len(connectionModels) == 0 {
		fmt.Println("nothing to push")
		return nil
	}

	connectionsJSON, err := json.Marshal(connectionModels)

	if err != nil {
		return err
	}

	args := syncArgsFromCLIArgs(ctx.CLIContext)

	req, err := http.NewRequestWithContext(
		ctx.CLIContext.Context,
		"POST",
		args.address+"/sync/push",
		bytes.NewReader(connectionsJSON),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(args.username, args.password)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to push: %s", res.Status)
	}

	fmt.Println("pushed")
	return nil
}

func Pull(ctx *command_context.CommandContext) error {
	args := syncArgsFromCLIArgs(ctx.CLIContext)

	req, err := http.NewRequestWithContext(
		ctx.CLIContext.Context,
		"GET",
		args.address+"/sync/pull",
		http.NoBody,
	)

	if err != nil {
		return err
	}

	req.SetBasicAuth(args.username, args.password)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to pull: %s", res.Status)
	}

	var connectionModels []models.Connection

	if err := json.NewDecoder(res.Body).Decode(&connectionModels); err != nil {
		return err
	}

	db := ctx.Provider.DBFactory()
	tx, err := db.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := tx.Select().Delete(&models.Connection{}); err != nil {
		return err
	}

	for _, model := range connectionModels {
		if err := tx.Save(&model); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	fmt.Println("pulled")
	return nil
}

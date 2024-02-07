package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/TrixiS/mantra/pkg/command_context"
	"github.com/TrixiS/mantra/pkg/models"
	"github.com/rodaine/table"
	"golang.design/x/clipboard"
)

func Add(ctx *command_context.CommandContext) error {
	connection := models.NewConnectionFromCLIArgs(ctx.CLIContext)

	if err := ctx.Provider.DB.Save(connection); err != nil {
		return err
	}

	fmt.Printf("Connection %s added with ID %d\n", connection.Name, connection.ID)
	return nil
}

// TODO: remove by name as well
func Remove(ctx *command_context.CommandContext) error {
	connectionId := ctx.CLIContext.Int("id")

	if err := ctx.Provider.DB.DeleteStruct(&models.Connection{ID: connectionId}); err != nil {
		return err
	}

	fmt.Printf("Connection with ID %d has been removed\n", connectionId)
	return nil
}

// TODO: an ability to show the table with passwords monkaS
func List(ctx *command_context.CommandContext) error {
	var connections []models.Connection

	if err := ctx.Provider.DB.All(&connections); err != nil {
		return err
	}

	t := table.New("ID", "Name", "Host", "Port", "User")

	for _, conn := range connections {
		t.AddRow(conn.ID, conn.Name, conn.Host, conn.Port, conn.User)
	}

	t.Print()
	return nil
}

func Connect(ctx *command_context.CommandContext) error {
	connectionId := ctx.CLIContext.Int("id")

	var connection models.Connection

	if err := ctx.Provider.DB.One("ID", connectionId, &connection); err != nil {
		return err
	}

	command := exec.Command(
		"ssh",
		"-p", strconv.Itoa(int(connection.Port)),
		fmt.Sprintf("%s@%s", connection.User, connection.Host),
	)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout

	clipboard.Write(clipboard.FmtText, []byte(connection.Password))

	if err := command.Run(); err != nil {
		return err
	}

	return nil
}

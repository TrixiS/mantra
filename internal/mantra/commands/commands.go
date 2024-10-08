package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/TrixiS/mantra/internal/mantra/command_context"
	"github.com/TrixiS/mantra/internal/mantra/models"
	"github.com/asdine/storm/v3"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"golang.design/x/clipboard"
)

func Add(ctx *command_context.CommandContext) error {
	connection := models.NewConnectionFromCLIArgs(ctx.CLIContext)
	db := ctx.Provider.DBFactory()

	defer db.Close()

	if err := db.Save(connection); err != nil {
		return err
	}

	fmt.Printf("connection %s added with ID %d\n", connection.Name, connection.ID)
	return nil
}

func Remove(ctx *command_context.CommandContext) error {
	connectionId := ctx.CLIContext.Int("id")
	db := ctx.Provider.DBFactory()

	defer db.Close()

	if err := db.DeleteStruct(&models.Connection{ID: connectionId}); err != nil {
		return err
	}

	fmt.Printf("connection with ID %d has been removed\n", connectionId)
	return nil
}

func List(ctx *command_context.CommandContext) error {
	db := ctx.Provider.DBFactory()
	defer db.Close()

	var connections []models.Connection

	if err := db.All(&connections); err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Host", "Port", "User"})
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = false
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:   "ID",
			Colors: text.Colors{text.Bold, text.FgHiCyan},
		},
	})

	for _, conn := range connections {
		t.AppendRow(table.Row{conn.ID, conn.Name, conn.Host, conn.Port, conn.User})
	}

	t.Render()
	return nil
}

func Connect(ctx *command_context.CommandContext) error {
	if ctx.CLIContext.NArg() < 1 {
		return fmt.Errorf("specify connection name or ID")
	}

	db := ctx.Provider.DBFactory()
	connection, err := getConnectionByIdentifier(db, ctx.CLIContext.Args().First())
	db.Close()

	if err != nil {
		return fmt.Errorf("connection %w", err)
	}

	cmd := "ssh"
	connectionArgs := strings.Split(connection.Args, " ")
	execArgs := make([]string, len(connectionArgs)+2)

	if ctx.CLIContext.Bool("sftp") {
		cmd = "sftp"
		execArgs[0] = fmt.Sprintf("-P %d", connection.Port)
	} else {
		execArgs[0] = fmt.Sprintf("-p %d", connection.Port)
	}

	execArgs[1] = fmt.Sprintf("%s@%s", connection.User, connection.Host)

	for i, arg := range connectionArgs {
		execArgs[i+2] = arg
	}

	command := exec.Command(cmd, execArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout

	clipboard.Write(clipboard.FmtText, []byte(connection.Password))

	if err := command.Run(); err != nil {
		return fmt.Errorf("ssh %w", err)
	}

	return nil
}

func Reveal(ctx *command_context.CommandContext) error {
	connectionId := ctx.CLIContext.Int("id")

	db := ctx.Provider.DBFactory()
	defer db.Close()

	var connection models.Connection

	if err := db.One("ID", connectionId, &connection); err != nil {
		return fmt.Errorf("connection %w", err)
	}

	clipboard.Write(clipboard.FmtText, []byte(connection.Password))

	fmt.Printf("password for \"%s\" copied to your clipboard\n", connection.Name)
	return nil
}

func Update(ctx *command_context.CommandContext) error {
	connectionId := ctx.CLIContext.Int("id")

	db := ctx.Provider.DBFactory()
	defer db.Close()

	var connection models.Connection

	if err := db.One("ID", connectionId, &connection); err != nil {
		return fmt.Errorf("connection %w", err)
	}

	if ctx.CLIContext.IsSet("name") {
		connection.Name = ctx.CLIContext.String("name")
	}

	if ctx.CLIContext.IsSet("host") {
		connection.Host = ctx.CLIContext.String("host")
	}

	if ctx.CLIContext.IsSet("port") {
		connection.Port = ctx.CLIContext.Uint("port")
	}

	if ctx.CLIContext.IsSet("password") {
		connection.Password = ctx.CLIContext.String("password")
	}

	// TODO: parse args by a parser to an array (or! just use multiple args as -a)
	if ctx.CLIContext.IsSet("args") {
		connection.Args = ctx.CLIContext.String("args")
	}

	if err := db.Update(&connection); err != nil {
		return fmt.Errorf("update connection %w", err)
	}

	fmt.Printf("updated \"%s\"\n", connection.Name)
	return nil
}

func getConnectionByIdentifier(db *storm.DB, ident string) (*models.Connection, error) {
	var connection models.Connection

	id, err := strconv.Atoi(ident)

	if err == nil {
		if err := db.One("ID", id, &connection); err != nil {
			return nil, err
		}
	} else {
		if err := db.One("Name", ident, &connection); err != nil {
			return nil, err
		}
	}

	return &connection, nil
}

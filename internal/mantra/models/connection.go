package models

import "github.com/urfave/cli/v2"

type Connection struct {
	ID       int    `storm:"id,increment"`
	Name     string `storm:"index"`
	Host     string `storm:"index"`
	Port     uint
	User     string
	Password string
	Args     string
}

func NewConnectionFromCLIArgs(ctx *cli.Context) *Connection {
	return &Connection{
		Name:     ctx.String("name"),
		Host:     ctx.String("host"),
		Port:     ctx.Uint("port"),
		User:     ctx.String("user"),
		Password: ctx.String("password"),
		Args:     ctx.String("args"),
	}
}

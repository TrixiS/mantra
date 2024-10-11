package models

import "github.com/urfave/cli/v2"

type Connection struct {
	ID       int    `storm:"id,increment" json:"id"`
	Name     string `storm:"index"        json:"name"`
	Host     string `storm:"index"        json:"host"`
	Port     uint   `                     json:"port"`
	User     string `                     json:"user"`
	Password string `                     json:"password"`
	Args     string `                     json:"args"`
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

package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"github.com/TrixiS/mantra/ent"
	"github.com/TrixiS/mantra/internal/sync_server/controllers"
	"github.com/TrixiS/mantra/internal/sync_server/mw"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	sqlite "modernc.org/sqlite"
)

type config struct {
	Host   string `env:"HOST,required"`
	Port   uint16 `env:"PORT,required"`
	Secret string `env:"SECRET,required"`
}

func main() {
	godotenv.Load()

	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
	db, err := ent.Open(dialect.SQLite, "file:database.sqlite3")

	if err != nil {
		log.Fatalf("db conn: %v", err)
	}

	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("create schema: %v", err)
	}

	e := echo.New()

	userController := &controllers.UserController{
		DB: db,
	}

	userGroup := e.Group("/user", mw.NewSecretAuthMiddleware(cfg.Secret))
	userGroup.POST("", userController.Create)
	userGroup.DELETE("", userController.Delete)

	syncController := &controllers.SyncController{
		DB: db,
	}

	syncGroup := e.Group("/sync", mw.NewSyncAuthMiddleware(db))
	syncGroup.GET("/pull", syncController.Pull)
	syncGroup.GET("/pull", syncController.Pull)
	syncGroup.POST("/push", syncController.Push)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)))
}

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)

	if err != nil {
		return conn, err
	}

	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})

	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to enable fk: %w", err)
	}

	return conn, nil
}

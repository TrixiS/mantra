package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"github.com/TrixiS/mantra/ent"
	"github.com/TrixiS/mantra/internal/sync_server/controllers"
	"github.com/TrixiS/mantra/internal/sync_server/mw"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
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

	db, err := ent.Open(dialect.SQLite, "file:database.sqlite3?_fk=1")

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

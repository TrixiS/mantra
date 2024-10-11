package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/TrixiS/mantra/ent"
	"github.com/TrixiS/mantra/ent/connection"
	"github.com/TrixiS/mantra/ent/user"
	"github.com/TrixiS/mantra/internal/mantra/models"
	"github.com/labstack/echo/v4"
)

var pushTxOptions = &sql.TxOptions{Isolation: sql.LevelSerializable}

type SyncController struct {
	DB *ent.Client
}

func (s *SyncController) Pull(c echo.Context) error {
	usr := c.Get("user").(*ent.User)

	connections, err := s.DB.Connection.Query().
		Where(connection.HasOwnerWith(user.ID(usr.ID))).
		All(c.Request().Context())

	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}

	connectionModels := make([]models.Connection, len(connections))

	for i, conn := range connections {
		connectionModels[i] = models.Connection{
			ID:       conn.LocalID,
			Name:     conn.Name,
			Host:     conn.Host,
			Port:     conn.Port,
			User:     conn.User,
			Password: conn.Password,
			Args:     conn.Args,
		}
	}

	return c.JSON(http.StatusOK, connectionModels)
}

func (s *SyncController) Push(c echo.Context) error {
	var connectionModels []models.Connection

	if err := c.Bind(&connectionModels); err != nil {
		fmt.Println(err)
		return err
	}

	if len(connectionModels) == 0 {
		return echo.ErrForbidden
	}

	ctx := c.Request().Context()
	tx, err := s.DB.BeginTx(ctx, pushTxOptions)

	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}

	defer tx.Rollback()

	usr := c.Get("user").(*ent.User)
	_, err = tx.Connection.Delete().Where(connection.HasOwnerWith(user.ID(usr.ID))).Exec(ctx)

	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}

	connectionBuilders := make([]*ent.ConnectionCreate, len(connectionModels))

	for i, model := range connectionModels {
		connectionBuilders[i] = tx.Connection.Create().
			SetOwner(usr).
			SetLocalID(model.ID).
			SetName(model.Name).
			SetHost(model.Host).
			SetPort(model.Port).
			SetUser(model.User).
			SetPassword(model.Password).
			SetArgs(model.Args)
	}

	err = tx.Connection.CreateBulk(connectionBuilders...).Exec(ctx)

	if err != nil {
		if ent.IsValidationError(err) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}

	if err := tx.Commit(); err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}

	return nil
}

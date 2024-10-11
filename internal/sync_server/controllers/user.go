package controllers

import (
	"net/http"

	"github.com/TrixiS/mantra/ent"
	"github.com/TrixiS/mantra/ent/user"
	"github.com/TrixiS/mantra/internal/sync_server/auth"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	DB *ent.Client
}

type createUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type deleteUserPayload struct {
	Username string `json:"username"`
}

func (u *UserController) Create(c echo.Context) error {
	payload := createUserPayload{}

	if err := c.Bind(&payload); err != nil {
		return err
	}

	passwordHash, err := auth.HashPassword(payload.Password)

	if err != nil {
		return echo.ErrInternalServerError
	}

	usr, err := u.DB.User.Create().
		SetUsername(payload.Username).
		SetPasswordHash(passwordHash).
		Save(c.Request().Context())

	if err != nil {
		if ent.IsConstraintError(err) {
			return echo.ErrConflict
		}

		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

func (u *UserController) Delete(c echo.Context) error {
	payload := deleteUserPayload{}

	if err := c.Bind(&payload); err != nil {
		return err
	}

	rowsCount, err := u.DB.User.Delete().
		Where(user.Username(payload.Username)).
		Exec(c.Request().Context())

	if err != nil {
		return echo.ErrInternalServerError
	}

	if rowsCount == 0 {
		return echo.ErrNotFound
	}

	return nil
}

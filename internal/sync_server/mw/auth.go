package mw

import (
	"github.com/TrixiS/mantra/ent"
	"github.com/TrixiS/mantra/ent/user"
	"github.com/TrixiS/mantra/internal/sync_server/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewSecretAuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")

			if authorization != secret {
				return echo.ErrUnauthorized
			}

			return next(c)
		}
	}
}

func NewPasswordAuthMiddleware(db *ent.Client) echo.MiddlewareFunc {
	return middleware.BasicAuth(
		func(username string, password string, c echo.Context) (bool, error) {
			usr, err := db.User.Query().
				Where(user.Username(username)).
				First(c.Request().Context())

			if err != nil {
				if !ent.IsNotFound(err) {
					c.Logger().Error(err)
				}

				return false, nil
			}

			if !auth.VerifyPassword(password, usr.PasswordHash) {
				return false, nil
			}

			c.Set("user", usr)
			return true, nil
		},
	)
}

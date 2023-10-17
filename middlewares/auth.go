package middlewares

import (
	"chat-hex/business/auth"

	"github.com/labstack/echo/v4"
)

func Auth(next, stop echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return stop(c)
		}

		err := auth.NewService().ValidateToken(tokenString)
		if err != nil {
			return stop(c)
		}

		return next(c)
	}
}
package mw

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

const (
	roleAdmin = "admin"
)

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		role := ctx.Request().Header.Get("user-role")

		if strings.EqualFold(role, roleAdmin) {
			log.Println("Admin user is detected")

		}

		err := next(ctx)

		if err != nil {
			return err
		}

		return nil
	}
}

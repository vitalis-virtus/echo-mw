package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()

	server.Use(middleware.Recover())

	server.Use(Mw)

	server.GET("/status", Handler)

	// Start server
	server.Logger.Fatal(server.Start(":8080"))
}

func Handler(ctx echo.Context) error {
	data := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	duration := time.Until(data)

	answer := fmt.Sprintf("Number of days to New Year: %d", int64(duration.Hours()/24))

	err := ctx.String(http.StatusOK, answer)

	if err != nil {
		return err
	}

	return nil
}

func Mw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		role := ctx.Request().Header.Get("user-role")

		if strings.ToLower(role) == "admin" {
			log.Println("Admin user is detected")

		}

		err := next(ctx)

		if err != nil {
			return err
		}

		return nil
	}
}

package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vitalis-virtus/echo-mw/internal/app/endpoint"
	"github.com/vitalis-virtus/echo-mw/internal/app/mw"
	"github.com/vitalis-virtus/echo-mw/internal/app/service"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(middleware.Recover())

	a.echo.Use(mw.CheckRole)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server is running")
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

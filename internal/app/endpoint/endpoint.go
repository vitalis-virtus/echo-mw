package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	service Service
}

func New(srv Service) *Endpoint {
	return &Endpoint{service: srv}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	days := e.service.DaysLeft()

	answer := fmt.Sprintf("Number of days to New Year: %d", days)

	err := ctx.String(http.StatusOK, answer)

	if err != nil {
		return err
	}

	return nil
}

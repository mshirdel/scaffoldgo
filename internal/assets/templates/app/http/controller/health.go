package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"{{.ModuleName}}/app"
)

type HealthCheckController struct {
	app *app.Application
}

func NewHealthCheckController(app *app.Application) *HealthCheckController {
	return &HealthCheckController{
		app: app,
	}
}

func (h *HealthCheckController) Index(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *HealthCheckController) Liveness(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *HealthCheckController) Readiness(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

package controller

import (
	"fmt"
	"math"
	"strings"

	"github.com/labstack/echo"
	"{{.ModuleName}}/app"
)

type Controller struct {
	app                   *app.Application
	healthCheckController *HealthCheckController
}

func NewController(app *app.Application) *Controller {
	return &Controller{
		app:                   app,
		healthCheckController: NewHealthCheckController(app),
	}
}

func (c *Controller) Routes() *echo.Echo {
	router := c.initEcho()

	// initGeneralMiddlewares(router, &c.app.Cfg.HTTP, &c.app.Cfg.Locale)

	router.GET("/", c.healthCheckController.Index)

	health := router.Group("/healthz")
	{
		health.GET("/liveness", c.healthCheckController.Liveness)
		health.GET("/readiness", c.healthCheckController.Readiness)
	}

	if router.Debug {
		printRoutes(router.Routes())
	}

	return router
}

func (c *Controller) initEcho() *echo.Echo {
	r := echo.New()
	r.Debug = c.app.Cfg.Logging.Level == "debug"
	// r.Validator = validator.New()
	// r.HTTPErrorHandler = HTTPErrorHandler

	return r
}

func printRoutes(routes []*echo.Route) {
	var maxMethodLength, maxPathLength float64
	for _, r := range routes {
		maxMethodLength = math.Max(maxMethodLength, float64(len(r.Method)))
		maxPathLength = math.Max(maxPathLength, float64(len(r.Path)))
	}

	fmt.Printf("\nRegistered http routes:\n")

	for _, r := range routes {
		// do not print middlewares
		if strings.HasPrefix(r.Name, "github.com/labstack/echo") {
			continue
		}

		fmt.Printf("%-*v %-*v --> %v\n", int(maxMethodLength), r.Method, int(maxPathLength), r.Path, r.Name)
	}
}

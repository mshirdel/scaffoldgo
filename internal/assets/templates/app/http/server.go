package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"{{.ModuleName}}/app"
	"{{.ModuleName}}/app/http/controller"
)

type HTTPServer struct {
	app    *app.Application
	server *http.Server
}

func NewHTTPServer(app *app.Application) *HTTPServer {
	controller := controller.NewController(app)

	return &HTTPServer{
		app: app,
		server: &http.Server{
			Addr:              app.Cfg.Server.Address,
			ReadTimeout:       app.Cfg.Server.ReadTimeout,
			ReadHeaderTimeout: app.Cfg.Server.ReadTimeout,
			WriteTimeout:      app.Cfg.Server.WriteTimeout,
			IdleTimeout:       app.Cfg.Server.IdleTimeout,
			Handler:           controller.Routes(),
		},
	}
}

func (h *HTTPServer) Start(_ context.Context) {
	logrus.Infof("starting http server on: %s", h.app.Cfg.Server.Address)

	if err := h.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("failed starting http server: %v", err)
	}
}

func (h *HTTPServer) Shutdown(ctx context.Context) {
	logrus.Info("shutting down http server...")

	// TODO: make it configurable
	deadline, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := h.server.Shutdown(deadline); err != nil {
		logrus.Errorf("failed shutting down http server: %s", err)
	}
}

package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"{{.ModuleName}}/app"
	api "{{.ModuleName}}/app/http"
)

var _serve = &cobra.Command{
	Use:   "serve",
	Short: "serve APIs",
	RunE: func(cmd *cobra.Command, args []string) error {
		app := app.New(_configPath)
		if err := app.InitAll(); err != nil {
			return err
		}
		defer app.Shutdown()

		ctx := context.Background()
		server := api.NewHTTPServer(app)
		defer server.Shutdown(ctx)
		go server.Start(ctx)

		<-handleInterrupts()
		return nil
	},
}

func handleInterrupts() <-chan os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(
		signals,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	return signals
}

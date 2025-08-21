package app

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"{{.ModuleName}}/config"
)

type Application struct {
	configPath string
	Cfg        *config.Config
}

func New(configPath string) *Application {
	return &Application{
		configPath: configPath,
	}
}

func (app *Application) InitAll() error {
	if err := app.InitConfig(); err != nil {
		return err
	}

	// app.InitLogger()
	//
	// if err := app.InitTracer(); err != nil {
	// 	return err
	// }
	//
	// if err := app.InitAllDatabases(); err != nil {
	// 	return err
	// }
	//
	// app.InitAllRepositories()
	// app.InitAllThirdParties()
	// app.InitAllCache()
	//
	// if err := app.InitAllServices(); err != nil {
	// 	return err
	// }

	return nil
}

func (app *Application) InitConfig() (err error) {
	if app.Cfg != nil {
		return
	}

	app.Cfg, err = config.InitViper(app.configPath)
	if err != nil {
		return fmt.Errorf("config initialization failed: %w", err)
	}

	return
}

func (app *Application) Shutdown() {
	logrus.Info("shutting down application...")
}

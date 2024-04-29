package http_server

import (
	"EchoAPI/core/modules"
	"EchoAPI/infrastructure/adapters/repositories"
	"EchoAPI/infrastructure/configs"
	"EchoAPI/infrastructure/controllers/http_server/views"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type ServerApp interface {
	Run() error
}

type serverApp struct {
	httpServer *http.Server
	postgresDB *gorm.DB
}

func New(
	cfg configs.HttpServerConfig,
	postgresDB *gorm.DB,
) (ServerApp, error) {
	router := echo.New()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.EchoConfig.Host, cfg.EchoConfig.Port),
		Handler: router,
	}

	app := &serverApp{
		httpServer: httpServer,
		postgresDB: postgresDB,
	}

	api := router.Group("/api")
	err := views.MountUsers(api, app)
	if err != nil {
		return nil, err
	}
	err = views.MountHealthCheck(api)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func RunServer(cfg configs.HttpServerConfig) (err error) {
	db, err := gorm.Open(postgres.Open(cfg.PostgresConfig.ConnectionString))
	if err != nil {
		return err
	}

	app, err := New(cfg, db)
	if err != nil {
		return err
	}

	err = app.Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *serverApp) Run() error {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *serverApp) Transaction(fc func(modules *modules.Modules) (err error)) (err error) {
	err = s.postgresDB.Transaction(func(tx *gorm.DB) (err error) {
		appModules := modules.New(repositories.NewFactory(s.postgresDB))
		if err != nil {
			return err
		}

		cErr := fc(appModules)
		if cErr != nil {
			return cErr
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *serverApp) PostgresDB() *gorm.DB {
	return s.postgresDB
}

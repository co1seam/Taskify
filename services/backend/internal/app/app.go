package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/co1seam/Taskify/service/backend/internal/config"
	"github.com/co1seam/Taskify/service/backend/internal/repository"
	"github.com/co1seam/Taskify/service/backend/internal/services"
	"github.com/co1seam/Taskify/service/backend/internal/transport/rest"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	return &App{}
}

func (a *App) Initialize() (*rest.Handler, error) {
	cfg := config.NewConfig()
	if err := cfg.LoadConfig(); err != nil {
		return nil, err
	}

	db, err := repository.NewPostrgesDB(repository.Config{
		Host:     cfg.Get("db.hpst"),
  		Port:     cfg.Get("db.port"),
  		Username: cfg.Get("db.username"),
  		Password: cfg.Get("db.password"),
		DBName: cfg.Get("db.dbname"),
		SSLMode: cfg.Get("db.sslmode"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	handler := rest.NewHandler(service)

	return handler, nil
}

func (a *App) Run(port string, handler http.Handler) error {
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	if a.httpServer == nil {
		return fmt.Errorf("http server is not initialized")
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s", err)
		}
		}()

	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	if a.httpServer == nil {
		return fmt.Errorf("http server is not initialized")
	}
	return a.httpServer.Shutdown(ctx)
}
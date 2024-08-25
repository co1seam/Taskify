package main

import (
	"context"

	"github.com/co1seam/Taskify/service/backend/internal/app"
	"github.com/co1seam/Taskify/service/backend/internal/config"
	"github.com/sirupsen/logrus"
)

func main()  {
	a := app.NewApp()
	handler, err := a.Initialize()
	if err != nil {
		logrus.Fatalf("Failed to initialize app: %v", err)
	}
	
	cfg := config.NewConfig()
	if err := cfg.LoadConfig(); err != nil {
		logrus.Fatalf("Failed to load config: %v", err)
	}

	if err := a.Run(cfg.Get("app.port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("Failed to load config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := a.Shutdown(ctx); err != nil {
		logrus.Fatalf("Failed to shutdown app: %v", err)
	}
}
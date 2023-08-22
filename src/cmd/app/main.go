package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ServiceWeaver/weaver"
	"github.com/kiramishima/shining_guardian/src/adapter/database/postgresql/repository"
	"github.com/kiramishima/shining_guardian/src/config"
	"github.com/kiramishima/shining_guardian/src/domain"
	"github.com/kiramishima/shining_guardian/src/handlers"
	"github.com/kiramishima/shining_guardian/src/server"
	"github.com/kiramishima/shining_guardian/src/services"
	"github.com/unrolled/render"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

// serve is called by weaver.Run and contains the body of the application.
func serve(ctx context.Context, app *domain.App) error {
	fmt.Println("Hello")
	// Logger
	logger := config.NewLogger()
	// Load config
	cfg := config.NewConfig()

	// Load Database
	conn, err := config.NewDatabase(cfg, logger)
	if err != nil {
		logger.Fatal("Database error", err)
		return err
	}
	// render
	rendr := render.New()
	// Load Repository
	repo := repository.NewAuthRepository(conn)
	// Load Service
	svc := services.NewAuthService(repo)
	// Load
	h := handlers.NewAuthHandler(logger, svc, rendr)
	//
	srv := server.NewServer(cfg, logger, h, app)

	return srv.Run()
	// return nil
}

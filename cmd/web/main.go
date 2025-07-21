package main

import (
	"context"
	"fmt"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spitch-id/spitch-backend/internal/config"
	"github.com/spitch-id/spitch-backend/internal/connection"
	"github.com/spitch-id/spitch-backend/internal/handler"
	"github.com/spitch-id/spitch-backend/internal/repository"
	"github.com/spitch-id/spitch-backend/internal/usecase"
)

func main() {
	env := config.NewEnv()
	app := config.NewFiber(env)

	config.Validator, config.Translator = config.NewValidator()

	database := connection.NewDatabase(env)
	if database == nil {
		log.Fatal("Failed to connect to the database")
	}

	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUseCase(database, userRepository)
	userHandler := handler.NewUserHandler(config.Validator, userUsecase, *env)

	apiGroup := app.Group("/api")
	config.NewServerConfig(&config.ServerConfig{
		App:         apiGroup,
		Validator:   config.Validator,
		UserHandler: userHandler,
	})

	done := make(chan bool, 1)

	go func() {
		port, _ := strconv.Atoi(env.SERVER_PORT)
		err := app.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			panic(fmt.Sprintf("http server error: %s", err))
		}

	}()

	go gracefulShutdown(app, done)
	<-done
	log.Info("Graceful shutdown complete.")
}

func gracefulShutdown(fiberServer *config.FiberServer, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Info("shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := fiberServer.ShutdownWithContext(ctx); err != nil {
		log.Errorf("Server forced to shutdown with error: %v", err)
	}

	log.Info("Server exiting")

	done <- true
}

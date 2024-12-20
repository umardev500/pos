package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/pos/internal/app/container"
	"github.com/umardev500/pos/pkg"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	v := pkg.NewValidator()
	db := pkg.NewGorm()
	containers := container.RegContainers(db, v)
	pkg.NewRouter(app, containers).Setup()

	ch := make(chan error, 1)
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		log.Info().Msgf("Starting server on port %s", port)

		ch <- app.Listen(":" + port)
	}()

	select {
	case err := <-ch:
		log.Fatal().Err(err).Msg("Failed to start server")
	case <-ctx.Done():
		log.Info().Msg("Shutting down server")
	}

	log.Info().Msg("Server is offline")
}

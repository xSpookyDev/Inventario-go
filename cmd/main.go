package main

import (
	"context"
	"inventario-go/database"
	"inventario-go/internal/repository"
	"inventario-go/internal/service"
	"inventario-go/settings"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),
		fx.Invoke(),
	)
	app.Run()

}

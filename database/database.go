package database

import (
	"context"
	"fmt"
	"inventario-go/settings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {

	log.Info().Msg("Iniciando Connexion a bd")

	if s.DB.User == "" || s.DB.Password == "" || s.DB.Host == "" || s.DB.Port == 0 || s.DB.Name == "" {
		return nil, fmt.Errorf("configuración de base de datos incompleta")
	}

	conneString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", s.DB.User, s.DB.Password, s.DB.Host, s.DB.Port, s.DB.Name)

	// Establecer un tiempo de espera para la conexión
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "mysql", conneString)
	if err != nil {
		log.Err(err).Msg("Error al conectar a la base de datos")
		return nil, fmt.Errorf("error al conectar a la base de datos: %w", err)
	}

	log.Info().Msg("Conexión a la base de datos exitosa")
	return db, nil

}

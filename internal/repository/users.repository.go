package repository

import (
	"context"
	"fmt"
	"inventario-go/internal/entity"

	"github.com/rs/zerolog/log"
)

const (
	queryInsertUser = `
	insert into USERS (email, name, password)
	values(?,?,?);
	`
	queryGetUserByEmail = `
	select
	id,
	email,
	name,
	password
	from USERS
	where email = ?;`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)
	if err != nil {
		log.Err(err).Msg("Error en los campos de la query")
	}
	return nil
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	u := &entity.User{}

	if err := r.db.GetContext(ctx, u, queryGetUserByEmail, email); err != nil {
		log.Err(err).Msgf("Error al ejecutar queryGetUserByEmail con email: %s", email)
		return nil, fmt.Errorf("no se pudo obtener el usuario con email %s: %w", email, err)
	}

	return u, nil
}

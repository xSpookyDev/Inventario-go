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

	qryInsertUserRole = `
		insert into USER_ROLES (user_id, role_id) values (:user_id, :role_id);`

	qryRemoveUserRole = `
		delete from USER_ROLES where user_id = :user_id and role_id = :role_id;`
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
func (r *repo) GetUserRoles(ctx context.Context, userID int) ([]entity.UserRole, error) {
	roles := []entity.UserRole{}

	err := r.db.SelectContext(ctx, &roles, "select user_id, role_id from USER_ROLES where user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryRemoveUserRole, data)

	return err
}
func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertUserRole, data)
	return err
}

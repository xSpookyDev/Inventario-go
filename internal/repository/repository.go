package repository

import (
	"context"
	"inventario-go/internal/entity"

	"github.com/jmoiron/sqlx"
)

//repository is the interface that wraps the Basic CRUD operations

//go:generate mockery --name=Repository -output=repository --inpackage

type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	SaveUserRole(ctx context.Context, userID, roleID int) error
	RemoveUserRole(ctx context.Context, userID, roleID int) error
	GetUserRoles(ctx context.Context, userID int) ([]entity.UserRole, error)

}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

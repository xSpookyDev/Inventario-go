package service

import (
	"context"
	"inventario-go/internal/models"
	"inventario-go/internal/repository"
)

//service is the business logic of the applcation

//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	AddUserRole(ctx context.Context, userID, roleID int) error
	RemoveUserRole(ctx context.Context, userID, roleID int) error
	//GetProducts(ctx context.Context) ([]models.Product, error)
	//GetProduct(ctx context.Context, id int64) (*models.Product, error)
	//AddProdcut(ctx context.Context, product models.Product, userEmail string) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}

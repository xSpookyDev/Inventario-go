package service

import (
	"context"
	"errors"
	"inventario-go/internal/models"
)

var (
	validRolesToAddProduct []int = []int{1, 2}
	ErrInvalidPermission         = errors.New("El usuario no tiene permisos para agregar productos")
)

func (s *serv) GetProducts(ctx context.Context) ([]models.Product, error) {
	pp, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := []models.Product{}

	for _, p := range pp {
		products = append(products, models.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       float32(p.Price),
		})

	}
	return products, nil

}

func (s *serv) GetProduct(ctx context.Context, id int) (*models.Product, error) {
	p, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	product := &models.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
	return product, nil

}

func (s *serv) AddProduct(ctx context.Context, product models.Product, email string) error {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	roles, err := s.repo.GetUserRoles(ctx, u.ID)
	if err != nil {
		return err
	}
	userCanAdd := false

	for _, r := range roles {
		for _, vr := range validRolesToAddProduct {
			if vr == r.RoleID {
				userCanAdd = true
			}
		}
	}

	if !userCanAdd {
		return ErrInvalidPermission
	}

	return s.repo.SaveProduct(ctx, product.Name, product.Description, product.Price, u.ID)
}

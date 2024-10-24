package repository

import (
	"context"

	"inventario-go/internal/entity"
)

const (
	queryInsertProduct = `
	INSERT INTO PRODUCTS (name, description, price, created_by)
	VALUES(?,?,?,?);`

	queryGetAllProducts = ` select id, name, description, price, created_by from PRODUCTS;`

	queryGetProductByID = `select id, name, description, price, created_by from PRODUCTS where id = ?;`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int) error {

	_, err := r.db.ExecContext(ctx, queryInsertProduct, name, description, price, createdBy)
	return err
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {
	pp := []entity.Product{}
	err := r.db.SelectContext(ctx, &pp, queryGetAllProducts)

	if err != nil {
		return nil, err
	}
	return pp, nil
}

func (r *repo) GetProduct(ctx context.Context, id int) (*entity.Product, error) {
	p := &entity.Product{}

	err := r.db.GetContext(ctx, p, queryGetProductByID, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

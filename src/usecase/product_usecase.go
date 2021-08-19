package usecase

import (
	"context"

	"github.com/firmanJS/gin-sqlx/src/database/model"
	"github.com/firmanJS/gin-sqlx/src/repository"
)

type Product struct {
	repo repository.ProductRepositoryInterface
}

func NewProduct(repo repository.ProductRepositoryInterface) ProductUseCaseInterface {
	return &Product{
		repo,
	}
}

func (product *Product) ProductList(ctx context.Context) (*model.Product, error) {

	resp, err := product.repo.GetProduct(ctx)
	if resp == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

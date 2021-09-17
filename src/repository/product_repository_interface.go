package repository

import (
	"context"

	"github.com/firmanJS/gin-sqlx/src/database/model"
)

type ProductRepositoryInterface interface {
	ProductList(ctx context.Context) (*model.Product, error)
}

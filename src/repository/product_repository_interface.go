package repository

import (
	"context"

	"github.com/firmanJS/gin-sqlx/src/database/model"
)

type ProductRepositoryInterface interface {
	GetProduct(ctx context.Context) (*model.Product, error)
}

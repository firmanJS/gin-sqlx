package usecase

import (
	"context"

	"github.com/firmanJS/gin-sqlx/src/database/model"
)

type ProductUseCaseInterface interface {
	ProductList(ctx context.Context) (*model.Product, error)
}

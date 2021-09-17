package usecase

import (
	"github.com/firmanJS/gin-sqlx/src/repository/postgres"
)

type Usecase struct {
	ProductUsecase ProductUseCaseInterface
}

func NewUcase(r *postgres.Repositories) *Usecase {
	return &Usecase{
		ProductUsecase: NewProduct(r.ProductRepos),
	}
}

package postgres

import (
	"github.com/firmanJS/gin-sqlx/src/repository"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ProductRepos repository.ProductRepositoryInterface
}

func NewPostgresRepositories(Conn *sqlx.DB) *Repositories {
	return &Repositories{
		ProductRepos: NewProductRepository(Conn),
	}
}

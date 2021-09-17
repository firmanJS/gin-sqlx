package postgres

import (
	"bytes"
	"context"
	"database/sql"

	"github.com/firmanJS/gin-sqlx/src/database/model"
	"github.com/firmanJS/gin-sqlx/src/repository"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	conn *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) repository.ProductRepositoryInterface {
	return &ProductRepository{
		conn: db,
	}
}

func (r *ProductRepository) ProductList(ctx context.Context) (*model.Product, error) {
	var query bytes.Buffer
	var err error
	var result = &model.Product{}

	query.WriteString(`SELECT * FROM entity_products`)

	if ctx != nil {
		err = r.conn.GetContext(ctx, result, query.String())
	} else {
		err = r.conn.Get(result, query.String())
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

package model

import (
	"time"
)

type Product struct {
	Id         string    `json:"id" db:"id"`
	Id_Product string    `json:"id_product" db:"id_product"`
	Name       string    `json:"name" db:"name"`
	Price      int64     `json:"price" db:"price"`
	Quantity   int64     `json:"quantity" db:"quantity"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

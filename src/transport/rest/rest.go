package rest

import (
	"github.com/firmanJS/gin-sqlx/src/database"
	"github.com/firmanJS/gin-sqlx/src/repository/postgres"
	"github.com/firmanJS/gin-sqlx/src/usecase"
	"github.com/gin-gonic/gin"
)

func NewHandler(r *gin.Engine) {
	db := database.NewConnection()
	product := postgres.NewProductRepository(db)
	productUsecase := usecase.NewProduct(product)
	InitProductRestHttp(r, productUsecase)
}

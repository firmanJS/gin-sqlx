package rest

import (
	"log"
	"net/http"

	"github.com/firmanJS/gin-sqlx/src/usecase"
	"github.com/gin-gonic/gin"
)

type RestProduct struct {
	useCase usecase.ProductUseCaseInterface
}

func InitProductRestHttp(r *gin.Engine, useCase usecase.ProductUseCaseInterface) {
	handler := &RestProduct{useCase}

	api := r.Group("/product")
	{
		api.GET("/", handler.ProductList)
	}
}

func (handler *RestProduct) ProductList(c *gin.Context) {

	data, err := handler.useCase.ProductList(c.Request.Context())
	if err != nil {
		log.Fatal(err)
	}

	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Tracking tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

package main

import (
	"log"
	"strconv"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/firmanJS/gin-sqlx/src/config"
	"github.com/firmanJS/gin-sqlx/src/database"
	"github.com/firmanJS/gin-sqlx/src/repository/postgres"
	"github.com/firmanJS/gin-sqlx/src/route"
	"github.com/firmanJS/gin-sqlx/src/transport/rest"
	"github.com/firmanJS/gin-sqlx/src/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	/**
	  @description Setup Server
	*/
	config, err := config.AppConfig()

	if err != nil {
		log.Fatal(err)
	}

	router := SetupRouter()

	log.Fatal(router.Run(":" + strconv.Itoa(config.AppPort)))
}

func SetupRouter() *gin.Engine {
	/*
	   @description Setup Database Connection
	*/
	db := database.NewConnection()
	/*
	   @description Init Router
	*/
	router := gin.Default()

	/*
	   @description Setup Mode Application
	*/
	config, err := config.AppConfig()
	if err != nil {
		log.Fatal(err)
	}

	if config.AppEnv != "production" && config.AppEnv != "test" {
		gin.SetMode(gin.DebugMode)
	} else if config.AppEnv == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	/*
	   @description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	/**
	  @description Init All Routes
	*/
	product := postgres.NewProductRepository(db)
	productUsecase := usecase.NewProduct(product)
	rest.InitProductRestHttp(router, productUsecase)

	route.InitRoute(router)

	return router
}

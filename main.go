package main

import (
	"fmt"
	"go-api/internal/books/api/v1/checkout"
	"go-api/internal/books/api/v1/create"
	"go-api/internal/books/api/v1/get"
	"go-api/internal/books/api/v1/getall"
	"go-api/internal/books/api/v1/health"
	returnbook "go-api/internal/books/api/v1/return"
	"go-api/internal/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	var cfg config.Config
	err := envconfig.Process("go-api", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	port := fmt.Sprintf(":%s", cfg.Port)

	router := gin.Default()
	router.GET("/books", getall.GetBooks)
	router.GET("/books/:id", get.GetBookById)
	router.GET("/healtz", health.GetHealthcheck)
	router.POST("/books", create.CreateBooks)
	router.PATCH("/checkout", checkout.CheckoutBook)
	router.PATCH("/return", returnbook.ReturnBook)
	router.Run(port)
}

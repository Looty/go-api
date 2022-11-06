package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Looty/go-api/internal"
	"github.com/kelseyhightower/envconfig"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "up")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing id query parameter"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if book.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing id query parameter"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if book.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not available"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func createBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// doesBookExists, err := getBookById(newBook.ID)

	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	return
	// }

	// if newBook.ID <= "0" {
	// 	c.JSON(http.StatusConflict, gin.H{"error": "book ID must be greater than 0"})
	// 	return
	// }

	// if len(doesBookExists.ID) > 0 {
	// 	c.JSON(http.StatusConflict, gin.H{"error": "this book ID already exists"})
	// 	return
	// }

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {

	var cfg internal.Config
	err := envconfig.Process("go-api", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	port := fmt.Sprintf(":%s", cfg.Port)

	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.GET("/healtz", healthCheck)
	router.POST("/books", createBooks)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run(port)

	// e.GET("/healtz", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
}

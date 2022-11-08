package create

import (
	"go-api/internal/config/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBooks(c *gin.Context) {
	var newBook books.Book

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

	books.Bookvar = append(books.Bookvar, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

package get

import (
	"go-api/internal/config/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := books.GetBookInstanceById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

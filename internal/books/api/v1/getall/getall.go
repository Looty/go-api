package getall

import (
	"go-api/internal/config/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books.Bookvar)
}

package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, "up")
}

package healthCheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Healthy is the handler for Healthy endpoint
// @Summary Check the health of the application
// @Router /health [get]
// @Produce  json
// @Success 200 {string} json "Health check passed"
func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

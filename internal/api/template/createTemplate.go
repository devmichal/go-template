package template

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handler(c *gin.Context) {

	c.JSON(http.StatusCreated, "id")
}

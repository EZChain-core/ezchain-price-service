package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Heathcheck(c *gin.Context) {

	c.JSON(
		http.StatusOK, &Response{
			Success: true,
			Message: "healthy",
		},
	)
}


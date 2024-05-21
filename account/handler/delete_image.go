package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteImage",
	})
}

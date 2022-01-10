package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		val := h.Backend.GetAllUsers()
		if val != nil {
			log.Println("GetAllUsers:", val)
		}

		c.JSON(200, val)
	}
}

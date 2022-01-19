package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchAllGroups() gin.HandlerFunc {
	return func(c *gin.Context) {
		val := h.backend.GetAllGroup()
		if val != nil {
			log.Println("GetAllGroup:", val)
		}
		c.JSON(200, val)
	}
}

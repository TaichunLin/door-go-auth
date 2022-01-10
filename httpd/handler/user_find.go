package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FindUserRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		cardId := c.Query("cardId")

		val := h.Backend.GetUser("b2:dm:user:" + cardId)
		if val != nil {
			log.Println("GetUser:", val)
		}
		c.JSON(200, val)
	}
}

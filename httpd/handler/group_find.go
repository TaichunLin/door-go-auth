package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FindGroupRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		groupId := c.Query("groupId")

		val := h.Backend.GetGroup("b2:dm:group:" + groupId)
		if val != nil {
			log.Println("GetGroup:", val)
		}
		c.JSON(200, val)
	}
}

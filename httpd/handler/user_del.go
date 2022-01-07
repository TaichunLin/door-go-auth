package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteUserRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		cardId := c.Query("cardId")

		err := h.backend.Del("b2:dm:user:" + cardId)
		if err != nil {
			log.Println("err:", err)
			c.String(200, "not found this group")
		} else {

			c.JSON(200, gin.H{
				"message": "刪除User成功",
				"cardId":  cardId,
			})
		}
	}
}

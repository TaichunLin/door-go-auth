package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteGroupRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		groupId := c.Query("groupId")

		err := h.Backend.Del("b2:dm:group:" + groupId)
		if err != nil {
			log.Println("err:", err)
			c.String(200, "not found this group")
		} else {

			c.JSON(200, gin.H{
				"message": "刪除Group成功",
				"groupId": groupId,
			})
		}
	}
}

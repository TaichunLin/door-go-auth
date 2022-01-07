package handler

import (
	"log"

	"GO-GIN_REST_API/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddGroupRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		//group=GroupABC&groupId=777&door=B
		groupId := c.Query("groupId")
		door := c.Query("door")
		group := c.Query("group")

		err := h.backend.SetGroup("b2:dm:group:"+groupId, &entity.Group{Group: group, GroupId: groupId, Door: door})
		if err != nil {
			log.Println("SetGroup failed:", err)
		} else {
			c.JSON(200, gin.H{
				"message": "新增Group成功",
				"key":     "b2:dm:group:" + groupId,
				"group":   group,
				"groupId": groupId,
				"door":    door,
			})
		}
	}
}

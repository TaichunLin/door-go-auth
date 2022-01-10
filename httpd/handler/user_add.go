package handler

import (
	"log"

	"GO-GIN_REST_API/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddUserRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		//group=GroupABC&groupId=777&door=B
		username := c.Query("username")
		cardId := c.Query("cardId")
		groupId := c.Query("groupId")
		group := h.Backend.GetGroup("b2:dm:group:" + groupId).Group
		door := h.Backend.GetGroup("b2:dm:group:" + groupId).Door
		log.Println("Get User's Group:", group)

		err := h.Backend.SetUser("b2:dm:user:"+cardId, &entity.User{Username: username, GroupId: &entity.Group{GroupId: groupId, Group: group, Door: door}, CardId: cardId})

		if err != nil {
			log.Println("SetUser failed:", err)
		} else {
			c.JSON(200, gin.H{
				"message":  "新增User成功",
				"key":      "b2:dm:user:" + cardId,
				"username": username,
				"cardId":   cardId,
				"group":    group,
				"groupId":  groupId,
			})
		}
	}
}

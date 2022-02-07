package handler

import (
	"GO-GIN_REST_API/entity"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchAllRoute(x string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch x {
		case "group":
			val := h.backend.GetAllGroup()
			if val != nil {
				log.Println("GetAllGroup:", val)
				c.JSON(200, val)
			}
		case "user":
			val := h.backend.GetAllUsers()
			if val != nil {
				log.Println("GetAllUsers:", val)
			}
			c.JSON(200, val)
		}
	}
}

func (h *Handler) AddRoute(x string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch x {
		case "group":
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

		case "user":

			//group=GroupABC&groupId=777&door=B
			username := c.Query("username")
			cardId := c.Query("cardId")
			groupId := c.Query("groupId")
			group := h.backend.GetGroup("b2:dm:group:" + groupId).Group
			door := h.backend.GetGroup("b2:dm:group:" + groupId).Door
			log.Println("Get User's Group:", group)

			err := h.backend.SetUser("b2:dm:user:"+cardId, &entity.User{Username: username, GroupId: &entity.Group{GroupId: groupId, Group: group, Door: door}, CardId: cardId})

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
}

func (h *Handler) DeleteRoute(x string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch x {
		case "user":
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
		case "group":
			groupId := c.Query("groupId")

			err := h.backend.Del("b2:dm:group:" + groupId)
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
}

func (h *Handler) FindRoute(x string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch x {
		case "user":
			cardId := c.Query("cardId")

			val := h.backend.GetUser("b2:dm:user:" + cardId)
			if val != nil {
				log.Println("GetUser:", val)
			}
			c.JSON(200, val)
		case "group":
			groupId := c.Query("groupId")

			val := h.backend.GetGroup("b2:dm:group:" + groupId)
			if val != nil {
				log.Println("GetGroup:", val)
			}
			c.JSON(200, val)
		}
	}
}

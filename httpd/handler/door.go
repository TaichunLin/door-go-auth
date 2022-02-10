package handler

import (
	"GO-GIN_REST_API/entity"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
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

			group := entity.Group{}
			data, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Println("jsonData:", err)
			}
			err = json.Unmarshal(data, &group)
			if err != nil {
				log.Println(err.Error())
			}
			err = h.backend.Set("b2:dm:group:"+group.GroupId, &entity.Group{Group: group.Group, GroupId: group.GroupId, Door: group.Door})
			if err != nil {
				log.Println("SetGroup failed:", err)
			} else {
				c.JSON(200, gin.H{
					"message":    "新增Group成功",
					"key":        "b2:dm:group:" + group.GroupId,
					"group":      group.Group,
					"groupId":    group.GroupId,
					"door":       group.Door,
					"csrf_token": nosurf.Token(c.Request),
				})
			}

		case "user":

			user := entity.User{}
			data, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Println("jsonData:", err)
			}
			err = json.Unmarshal(data, &user)
			if err != nil {
				log.Println(err.Error())
			}
			log.Println(string(data))
			log.Println(user)

			group := h.backend.GetGroup("b2:dm:group:" + user.GroupId.GroupId).Group
			door := h.backend.GetGroup("b2:dm:group:" + user.GroupId.GroupId).Door

			dberr := h.backend.Set("b2:dm:user:"+user.CardId, &entity.User{Username: user.Username, GroupId: &entity.Group{GroupId: user.GroupId.GroupId, Group: group, Door: door}, CardId: user.CardId})

			if dberr != nil {
				log.Println("SetUser failed:", err)
			} else {
				// c.JSON(200, user)
				c.JSON(200, gin.H{
					"message":    "新增User成功",
					"key":        "b2:dm:user:" + user.CardId,
					"username":   user.Username,
					"cardId":     user.CardId,
					"group":      group,
					"groupId":    user.GroupId.GroupId,
					"csrf_token": nosurf.Token(c.Request),
				})
			}
		}
	}
}

func (h *Handler) DeleteRoute(x string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch x {
		case "user":

			user := entity.User{}
			data, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Println("jsonData:", err)
			}
			err = json.Unmarshal(data, &user)
			if err != nil {
				log.Println(err.Error())
			}
			err = h.backend.Del("b2:dm:user:" + user.CardId)
			if err != nil {
				log.Println("err:", err)
				c.String(200, "not found this group")
			} else {

				c.JSON(200, gin.H{
					"message":    "刪除User成功",
					"cardId":     user.CardId,
					"csrf_token": nosurf.Token(c.Request),
				})
			}
		case "group":
			group := entity.Group{}
			data, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Println("jsonData:", err)
			}
			err = json.Unmarshal(data, &group)
			if err != nil {
				log.Println(err.Error())
			}

			err = h.backend.Del("b2:dm:group:" + group.GroupId)
			if err != nil {
				log.Println("err:", err)
				c.String(200, "not found this group")
			} else {

				c.JSON(200, gin.H{
					"message":    "刪除Group成功",
					"groupId":    group.GroupId,
					"csrf_token": nosurf.Token(c.Request),
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

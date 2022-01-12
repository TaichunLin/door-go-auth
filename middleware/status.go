package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			log.Println("err")
			c.Set("is_logged_in", false)
			return
		}
		log.Println("token:", token)
		log.Println("err:", err)
		if strings.Compare(token, "") == 0 {
			log.Println("false")
			c.Set("is_logged_in", false)
		} else {
			log.Println("true")
			c.Set("is_logged_in", true)
		}
	}
}

package handler

import (
	"github.com/gin-gonic/gin"
)

func ShowIndexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(
			200,
			"index.html",
			gin.H{
				"is_logged_in": c.MustGet("is_logged_in").(bool),
				"Title":        "Home Page",
			},
		)
	}
}

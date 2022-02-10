package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
)

//========nosurf========
func CsrfFailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", nosurf.Reason(r))
	log.Println("==============================================")
	log.Println(nosurf.Reason(r))
}
func SetCSRFTokenResHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-CSRF-Token", nosurf.Token(c.Request))
		c.Next()
	}
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		val, err := ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		deleted, delErr := h.backend.DeleteAuthor(`jwtMetadata:AccessUuid:` + val.AccessUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		c.JSON(http.StatusOK, "Successfully logged out")
	}
}

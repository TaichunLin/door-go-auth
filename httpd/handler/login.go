package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")

		if err := bcrypt.CompareHashAndPassword(h.backend.FetchAuthen("b2:dm:account:"+email).Password, []byte(password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "incorrect password",
			})
			return
		}
		if h.backend.FetchAuthen("b2:dm:account:"+email).Username != username {
			c.JSON(http.StatusUnauthorized, "Please provide valid login username")
			return
		}

		tokenMetadata, err := CreateToken(email)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		saveErr := h.backend.CreateAuthor(email, tokenMetadata)
		if saveErr != nil {
			c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		}

		// tokens := map[string]string{
		// 	"access_token":  tokenMetadata.AccessToken,
		// 	"refresh_token": tokenMetadata.RefreshToken,
		// }
		c.JSON(http.StatusOK, gin.H{
			"access_token":  tokenMetadata.AccessToken,
			"refresh_token": tokenMetadata.RefreshToken,
		})
	}
}

package handler

import (
	"log"
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
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"Title":        "Login",
				"ErrorTitle":   "Login Failed",
				"ErrorMessage": "Please provide valid login username"})
			return
		}

		tokenMetadata, err := CreateToken(email)
		if err != nil {
			c.HTML(http.StatusUnprocessableEntity, "login.html", gin.H{
				"Title":        "Login",
				"ErrorTitle":   "Login Failed",
				"ErrorMessage": err.Error()})
			return
		}
		saveErr := h.backend.CreateAuthor(email, tokenMetadata)
		if saveErr != nil {
			c.HTML(http.StatusUnprocessableEntity, "login.html", gin.H{
				"Title":        "Login",
				"ErrorTitle":   "Login Failed",
				"ErrorMessage": saveErr.Error()})
			return
		}

		tokens := map[string]string{
			"access_token":  tokenMetadata.AccessToken,
			"refresh_token": tokenMetadata.RefreshToken,
		}

		log.Println("Successful Login")
		log.Println(tokens)

		c.SetCookie("token", tokenMetadata.AccessToken, 3600, "", "", false, true)

		c.Redirect(http.StatusFound, "/auth/loginSuccess")
	}
}

func (h *Handler) LoginSuccess() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.HTML(200, "login-successful.html", gin.H{
			"Title": "logged in",
		})

	}
}

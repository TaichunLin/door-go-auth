package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(c *gin.Context) {
	//username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if err := bcrypt.CompareHashAndPassword(h.backend.FetchAuthen("b2:dm:account:"+email).Password, []byte(password)); err != nil {
		ErrorHtml(c, "login.html", "Login", "Login failed", "invalid emailor haven't registered yet?")
		return
	}
	if h.backend.FetchAuthen("b2:dm:account:"+email).Email != email {
		ErrorHtml(c, "login.html", "Login", "Login failed.", "Please provide valid email or haven't registered yet?")
		return
	}
	username := h.backend.FetchAuthen("b2:dm:account:" + email).Username
	tokenMetadata, err := CreateToken(email)

	if err != nil {
		ErrorHtml(c, "login.html", "Login", "Login failed", err.Error())
		return
	}
	saveErr := h.backend.CreateAuthor(email, tokenMetadata)
	if saveErr != nil {
		ErrorHtml(c, "login.html", "Login", "Login failed", saveErr.Error())
		return
	}

	tokens := map[string]string{
		"access_token":  tokenMetadata.AccessToken,
		"refresh_token": tokenMetadata.RefreshToken,
	}

	log.Println("Successful Login")
	log.Println(tokens)

	c.SetCookie("token", tokenMetadata.AccessToken, 900, "", "", false, true)
	c.SetCookie("refresh", tokenMetadata.RefreshToken, 86400, "", "", false, true)
	Html(c, 200, "text.html", "logged in", "You have successfuly logged in.", "Welcome"+username)
}

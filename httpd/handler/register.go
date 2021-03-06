package handler

import (
	"GO-GIN_REST_API/entity"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(c *gin.Context) {

	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	pwd, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	if strings.TrimSpace(password) == "" {
		ErrorHtml(c, "register.html", "Register", "Registration Failed", "the password can't be empty")
	} else if h.SameAccount(email) {
		ErrorHtml(c, "register.html", "Register", "Registration Failed", "the email isn't available or have already exist.")
	} else {
		err := h.backend.Set(`b2:dm:account:`+email, &entity.Accounts{Username: username, Password: pwd, Email: email})
		if err != nil {
			ErrorHtml(c, "register.html", "Register", "Registration Failed", err.Error())
		} else {
			c.HTML(200, "login.html", gin.H{
				"Title":      "Login",
				"Detail":     "registered successfully",
				"Detail2":    "Please log in.",
				"csrf_token": nosurf.Token(c.Request),
			})
		}
	}
}

func (h *Handler) SameAccount(email string) bool {
	val := h.backend.FetchAuthen(`b2:dm:account:` + email)
	if val != nil {
		if val.Email == email {
			return true
		}
	}
	return false
}

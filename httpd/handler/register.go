package handler

import (
	"GO-GIN_REST_API/entity"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {

		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")

		pwd, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

		if strings.TrimSpace(password) == "" {
			c.JSON(http.StatusBadRequest, "the password can't be empty")
			return
		} else if !h.IsAccountAvailable(email) {
			c.JSON(http.StatusBadRequest, "the email isn't available")
			return
		}
		err := h.backend.CreateAuthen(`b2:dm:account:`+email, &entity.Accounts{Username: username, Password: pwd, Email: email})
		if err != nil {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{
				"ErrorTitle":   "Registration Failed",
				"ErrorMessage": err.Error()})
			return
		} else {
			c.HTML(200, "login-successful.html", gin.H{
				"Title":       "registered",
				"DoSomething": "Please log in.",
				"Username":    username,
			})
		}
	}
}

func (h *Handler) IsAccountAvailable(email string) bool {
	val := h.backend.FetchAuthen(`b2:dm:account:` + email)
	if val != nil {
		if val.Email == email {
			return false
		}
	}
	return true
}

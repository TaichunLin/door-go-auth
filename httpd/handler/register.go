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
			c.JSON(http.StatusBadRequest, gin.H{
				"CreateAuthen failede": err,
			})
			return
		} else {
			c.JSON(200, gin.H{
				"message":  "新增anth會員成功",
				"username": username,
				"Email":    email,
				"password": password,
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

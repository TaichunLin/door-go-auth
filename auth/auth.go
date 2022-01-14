package auth

import (
	"GO-GIN_REST_API/cache"
	"GO-GIN_REST_API/entity"
	b64 "encoding/base64"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	redisconf = &cache.RedisConfig{Endpoint: "localhost:6379", Password: "", Database: 0, PoolSize: 0}
	redis, _  = cache.NewRedis(redisconf)
)

func RegisterNewUser(username, password string) (*entity.AuthUser, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	} else if !IsUsernameAvailable(username) {
		return nil, errors.New("the username isn't available")
	}

	u := entity.AuthUser{Username: username, Password: password}

	return &u, nil
}

func IsUsernameAvailable(username string) bool {

	val := redis.GetAuth("b2:dm:session:" + username)
	if val != nil {
		if val.Username == username {
			return false
		}
	}
	return true
}

func ShowRegistrationPage(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{
		"is_logged_in": c.MustGet("is_logged_in").(bool),
		"Title":        "Register",
	})
}

func Register(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := RegisterNewUser(username, password); err == nil {
		token := b64.StdEncoding.EncodeToString([]byte("block2!secret?" + username))
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		err := redis.SetAuth("b2:dm:session:"+username, &entity.AuthUser{Username: username, Password: password, Token: token})
		if err != nil {
			log.Println("SetAuth failed:", err)
		} else {
			session.Set(token, username)
			if err := session.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			}
			c.HTML(200, "login-successful.html", gin.H{
				"is_logged_in": c.MustGet("is_logged_in").(bool),
				"Title":        "Successful registration & Login",
			})
		}

		c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"is_logged_in": c.MustGet("is_logged_in").(bool),
			"Title":        "Register",
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})

	}
}

func IsUserValid(username, password string) bool {
	val := redis.GetAuth("b2:dm:session:" + username)
	if val != nil {
		if val.Username == username && val.Password == password {
			return true
		}
	}
	return false
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"is_logged_in": c.MustGet("is_logged_in").(bool),
		"Title":        "Login",
	})
}

func PerformLogin(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if IsUserValid(username, password) {

		token := b64.StdEncoding.EncodeToString([]byte("block2!secret?"))
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		session.Set(token, username)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		}
		c.HTML(200, "login-successful.html", gin.H{
			"is_logged_in": c.MustGet("is_logged_in").(bool),
			"Title":        "Successful Login & Successfully authenticated user",
		})

	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"is_logged_in": c.MustGet("is_logged_in").(bool),
			"Title":        "Login",
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	token := b64.StdEncoding.EncodeToString([]byte("block2!secret?"))
	session.Delete(token)
	session.Save()
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(b64.StdEncoding.EncodeToString([]byte("block2!secret?")))
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

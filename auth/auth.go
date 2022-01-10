package auth

import (
	"GO-GIN_REST_API/cache"
	"GO-GIN_REST_API/entity"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

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
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := RegisterNewUser(username, password); err == nil {
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		err := redis.SetAuth("b2:dm:session:"+username, &entity.AuthUser{Username: username, Password: password})
		if err != nil {
			log.Println("SetAuth failed:", err)
		} else {
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

func GenerateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
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
	username := c.PostForm("username")
	password := c.PostForm("password")

	if IsUserValid(username, password) {
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		c.HTML(200, "login-successful.html", gin.H{
			"is_logged_in": c.MustGet("is_logged_in").(bool),
			"Title":        "Successful Login",
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
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

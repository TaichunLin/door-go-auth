package auth

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []user{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
	{Username: "user3", Password: "pass3"},
}

func RegisterNewUser(username, password string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	} else if !IsUsernameAvailable(username) {
		return nil, errors.New("the username isn't available")
	}

	u := user{Username: username, Password: password}

	userList = append(userList, u)

	return &u, nil
}

func IsUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
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

		c.HTML(200, "login-successful.html", gin.H{
			"is_logged_in": c.MustGet("is_logged_in").(bool),
			"Title":        "Successful registration & Login",
		})
		c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"is_logged_in": c.MustGet("is_logged_in").(bool),
			"Title":        "Register",
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}

	log.Println(userList)
}

func GenerateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func IsUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
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

	log.Println(userList)
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

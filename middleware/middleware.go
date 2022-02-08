package middleware

import (
	"GO-GIN_REST_API/httpd/handler"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// var (
// 	redisconf = &cache.RedisConfig{Endpoint: "localhost:6379", Password: "", Database: 0, PoolSize: 0}
// 	redis, _  = cache.NewRedis(redisconf)
// )

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		log.Println("tokenString: ", tokenString)
		if err == nil && tokenString != "" {

			val, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("ACCESS_SECRET")), nil
			})
			fmt.Printf("%#v	\n", val)
			log.Println("val: ", val)
			log.Println("err: ", err)
			if _, ok := val.Claims.(jwt.MapClaims); !ok && !val.Valid {
				handler.ErrorHtml(c, "text.html", "Unauthorized", "something wrong 2 ", err.Error())
				return
			}
			c.Next()
		} else {
			refresh, err := c.Cookie("refresh")
			switch refresh != "" {
			case true:
				c.Redirect(http.StatusFound, "/auth/token/refresh")
			case false:
				handler.ErrorHtml(c, "text.html", "Unauthorized", "something wrong 3 ", err.Error())
				c.Abort()
			}
		}

		//get client method
		// switch method {
		// case "addgroup":
		// 	models.addgroup(...)
		// }
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func CustomCors() gin.HandlerFunc {
	var allowedMethods = []string{
		"POST",
		"GET",
		"OPTIONS",
		"PUT",
		"PATCH",
		"DELETE",
	}
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		fmt.Println("origin: ", origin)
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set(
				"Access-Control-Allow-Methods",
				strings.Join(allowedMethods, ", "),
			)
			c.Writer.Header().Set(
				"Access-Control-Allow-Headers",
				c.Request.Header.Get("Access-Control-Request-Headers"),
			)
		}
	}
}

package middleware

import (
	"GO-GIN_REST_API/cache"
	"GO-GIN_REST_API/httpd/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	redisconf = &cache.RedisConfig{Endpoint: "localhost:6379", Password: "", Database: 0, PoolSize: 0}
	redis, _  = cache.NewRedis(redisconf)
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

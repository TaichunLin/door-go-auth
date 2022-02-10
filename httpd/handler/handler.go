package handler

import (
	"GO-GIN_REST_API/cache"

	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
)

type Handler struct {
	backend *cache.RedisClient
}

func NewHandler() *Handler {
	redisconf := &cache.RedisConfig{Endpoint: "localhost:6379", Password: "", Database: 0, PoolSize: 0}
	redis, err := cache.NewRedis(redisconf)
	if err != nil {
		panic(err)
	}
	return &Handler{backend: redis}
}

func Html(c *gin.Context, http int, url string, title string, detail string, detail2 string) {

	c.HTML(http, url, gin.H{
		"Title":      title,
		"Detail":     detail,
		"Detail2":    detail2,
		"csrf_token": nosurf.Token(c.Request),
	})
}

func ErrorHtml(c *gin.Context, url string, title string, errorTitle string, errorMessage string) {
	c.HTML(401, url, gin.H{
		"Title":        title,
		"ErrorTitle":   errorTitle,
		"ErrorMessage": errorMessage,
		"csrf_token":   nosurf.Token(c.Request),
	})
}

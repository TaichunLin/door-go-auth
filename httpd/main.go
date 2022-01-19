package main

import (
	"GO-GIN_REST_API/httpd/handler"
	"GO-GIN_REST_API/middleware"

	"github.com/gin-gonic/gin"
)

// var (
// 	redisconf  = &cache.RedisConfig{Endpoint: "localhost:6379", Password: "", Database: 0, PoolSize: 0}
// 	redis, err = cache.NewRedis(redisconf)
// )

func main() {

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Static("/assets", "./templates/assets")
	r.LoadHTMLGlob("templates/*.html")

	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/user", middleware.TokenAuthMiddleware(), func(c *gin.Context) {
			c.HTML(200, "user.html", gin.H{
				"Title": "會員管理",
			})
		})
		viewRoutes.GET("/group", middleware.TokenAuthMiddleware(), func(c *gin.Context) {
			c.HTML(200, "group.html", gin.H{
				"Title": "部門管理",
			})
		})
	}

	h := handler.NewHandler()
	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/users", middleware.TokenAuthMiddleware(), h.FetchAllUsers())
		apiRoutes.GET("/addUser", middleware.TokenAuthMiddleware(), h.AddUserRoute())
		apiRoutes.GET("/findUser", middleware.TokenAuthMiddleware(), h.FindUserRoute())
		apiRoutes.GET("/deleteUser", middleware.TokenAuthMiddleware(), h.DeleteUserRoute())

		apiRoutes.GET("/groups", middleware.TokenAuthMiddleware(), h.FetchAllGroups())
		apiRoutes.GET("/addGroup", middleware.TokenAuthMiddleware(), h.AddGroupRoute())
		apiRoutes.GET("/findGroup", middleware.TokenAuthMiddleware(), h.FindGroupRoute())
		apiRoutes.GET("/deleteGroup", middleware.TokenAuthMiddleware(), h.DeleteGroupRoute())

		apiRoutes.POST("/login", h.Login())
		apiRoutes.POST("/register", h.Register())
		apiRoutes.POST("/logout", middleware.TokenAuthMiddleware(), h.Logout())
		apiRoutes.POST("/token/refresh", h.Refresh())
	}

	r.Run(":1106")

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

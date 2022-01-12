package main

import (
	"GO-GIN_REST_API/auth"
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
	r.Use(middleware.SetUserStatus())

	r.Static("/assets", "./templates/assets")
	r.LoadHTMLGlob("templates/*.html")

	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/user", middleware.EnsureLoggedIn(), func(c *gin.Context) {
			c.HTML(200, "user.html", gin.H{
				"is_logged_in": c.MustGet("is_logged_in").(bool),
				"Title":        "會員管理",
			})
		})
		viewRoutes.GET("/group", middleware.EnsureLoggedIn(), func(c *gin.Context) {
			c.HTML(200, "group.html", gin.H{
				"is_logged_in": c.MustGet("is_logged_in").(bool),
				"Title":        "部門管理",
			})
		})
	}

	h := handler.NewHandler()
	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/users", middleware.EnsureLoggedIn(), h.FetchAllUsers())
		apiRoutes.GET("/addUser", middleware.EnsureLoggedIn(), h.AddUserRoute())
		apiRoutes.GET("/findUser", middleware.EnsureLoggedIn(), h.FindUserRoute())
		apiRoutes.GET("/deleteUser", middleware.EnsureLoggedIn(), h.DeleteUserRoute())

		apiRoutes.GET("/groups", middleware.EnsureLoggedIn(), h.FetchAllGroups())
		apiRoutes.GET("/addGroup", middleware.EnsureLoggedIn(), h.AddGroupRoute())
		apiRoutes.GET("/findGroup", middleware.EnsureLoggedIn(), h.FindGroupRoute())
		apiRoutes.GET("/deleteGroup", middleware.EnsureLoggedIn(), h.DeleteGroupRoute())
	}

	r.GET("/", handler.ShowIndexPage())

	AuthRoutes := r.Group("/auth")
	{
		AuthRoutes.GET("/register", middleware.EnsureNotLoggedIn(), auth.ShowRegistrationPage)

		AuthRoutes.POST("/register", middleware.EnsureNotLoggedIn(), auth.Register)
		AuthRoutes.GET("/login", middleware.EnsureNotLoggedIn(), auth.ShowLoginPage)
		AuthRoutes.POST("/login", middleware.EnsureNotLoggedIn(), auth.PerformLogin)
		AuthRoutes.GET("/logout", middleware.EnsureLoggedIn(), auth.Logout)
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
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
	r.Use(middleware.CORSMiddleware())
	r.Static("/assets", "./templates/assets")
	r.LoadHTMLGlob("templates/*.html")
	h := handler.NewHandler()
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

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

	AuthRoutes := r.Group("/auth")
	{
		//ShowRegistrationPage
		AuthRoutes.GET("/registerPage", func(c *gin.Context) {
			c.HTML(200, "register.html", gin.H{
				"Title": "Register",
			})
		})
		AuthRoutes.GET("/loginSuccess", h.LoginSuccess())

		AuthRoutes.POST("/register", h.Register())

		//ShowLoginPage
		AuthRoutes.GET("/loginPage", func(c *gin.Context) {
			c.HTML(200, "login.html", gin.H{
				"Title": "Login",
			})
		})
		AuthRoutes.POST("/login", h.Login())
		AuthRoutes.GET("/logout", middleware.TokenAuthMiddleware(), h.Logout())
		AuthRoutes.POST("/token/refresh", h.Refresh())
	}

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
	}

	r.Run(":1106")

}

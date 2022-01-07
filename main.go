package main

import (
	"GO-GIN_REST_API/article"
	"GO-GIN_REST_API/auth"
	"GO-GIN_REST_API/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.Use(middleware.SetUserStatus())

	r.GET("/", article.ShowIndexPage)

	AuthRoutes := r.Group("/auth")
	{
		AuthRoutes.GET("/register", middleware.EnsureNotLoggedIn(), auth.ShowRegistrationPage)

		AuthRoutes.POST("/register", middleware.EnsureNotLoggedIn(), auth.Register)
		AuthRoutes.GET("/login", middleware.EnsureNotLoggedIn(), auth.ShowLoginPage)
		AuthRoutes.POST("/login", middleware.EnsureNotLoggedIn(), auth.PerformLogin)
		AuthRoutes.GET("/logout", middleware.EnsureLoggedIn(), auth.Logout)
	}

	articleRoutes := r.Group("/article")
	{

		articleRoutes.GET("/view/:article_id", article.GetArticle)

		articleRoutes.GET("/create", middleware.EnsureLoggedIn(), article.ShowArticleCreationPage)

		articleRoutes.POST("/create", middleware.EnsureLoggedIn(), article.CreateArticle)
	}

	r.Run(":1106")
}

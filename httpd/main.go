package main

import (
	"GO-GIN_REST_API/httpd/handler"
	"GO-GIN_REST_API/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/nosurf"
)

// var (
// 	redisconf  = &cache.RedisConfig{Endpoint: "localhost:6379", Password: "", Database: 0, PoolSize: 0}
// 	redis, err = cache.NewRedis(redisconf)
// )

func main() {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	//========nosurf========
	csrf := nosurf.New(r)
	csrf.SetFailureHandler(http.HandlerFunc(csrfFailHandler))

	// r.Use(CSRF())

	r.Static("/assets", "./templates/assets")
	r.LoadHTMLGlob("templates/*.html")

	h := handler.NewHandler()

	r.GET("/", func(c *gin.Context) { c.HTML(200, "index.html", nil) })

	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/user", middleware.TokenAuthMiddleware(), func(c *gin.Context) {
			c.Writer.Header().Set("X-CSRF-Token", nosurf.Token(c.Request))

			c.HTML(200, "user.html", gin.H{"Title": "會員管理", "csrf_token": nosurf.Token(c.Request)})
		})
		viewRoutes.GET("/group", middleware.TokenAuthMiddleware(), func(c *gin.Context) {
			c.Writer.Header().Set("X-CSRF-Token", nosurf.Token(c.Request))

			c.HTML(200, "group.html", gin.H{"Title": "部門管理", "csrf_token": nosurf.Token(c.Request)})
		})
	}

	AuthRoutes := r.Group("/auth")
	{
		//ShowRegistrationPage
		AuthRoutes.GET("/registerPage", func(c *gin.Context) {
			c.Writer.Header().Set("X-CSRF-Token", nosurf.Token(c.Request))
			c.HTML(200, "register.html", gin.H{
				"Title":      "Register",
				"csrf_token": nosurf.Token(c.Request),
				// csrf.TemplateTag: csrf.TemplateField(c.Request),
				// "csrfField":      csrf.TemplateField(c.Request),
			})
		})

		AuthRoutes.POST("/register", h.Register)

		//ShowLoginPage
		AuthRoutes.GET("/loginPage", func(c *gin.Context) {
			c.Writer.Header().Set("X-CSRF-Token", nosurf.Token(c.Request))
			c.HTML(200, "login.html", gin.H{
				"Title":      "Login",
				"csrf_token": nosurf.Token(c.Request),
				// csrf.TemplateTag:     csrf.TemplateField(c.Request),
				// "gorilla.csrf.Token": csrf.Token(c.Request),
				// "csrfField":          csrf.Token(c.Request),
			})
		})
		AuthRoutes.POST("/login", h.Login)
		AuthRoutes.GET("/logout", h.Logout)
		AuthRoutes.GET("/token/refresh", h.Refresh)
	}

	apiRoutes := r.Group("/api")
	apiRoutes.Use(middleware.TokenAuthMiddleware())
	{
		apiRoutes.GET("/users", h.FetchAllRoute("user"))
		apiRoutes.GET("/addUser", h.AddRoute("user"))
		apiRoutes.GET("/findUser", h.FindRoute("user"))
		apiRoutes.GET("/deleteUser", h.DeleteRoute("user"))

		apiRoutes.GET("/groups", h.FetchAllRoute("group"))
		apiRoutes.GET("/addGroup", h.AddRoute("group"))
		apiRoutes.GET("/findGroup", h.FindRoute("group"))
		apiRoutes.GET("/deleteGroup", h.DeleteRoute("group"))
	}
	// r.Run(":1106")
	//========nosurf========
	http.ListenAndServe(":1106", csrf)
	// http.ListenAndServe(":1106",
	// 	csrf.Protect([]byte("32-byte-long-auth-key"))(r))
}

// var csrfMd func(http.Handler) http.Handler = csrf.Protect([]byte("32-byte-long-auth-key"),
// 	csrf.MaxAge(0),
// 	csrf.Secure(false),
// 	csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusForbidden)
// 		w.Write([]byte(`{"message": "Forbidden - CSRF token invalid"}`))
// 		w.Write([]byte(csrf.Token(r)))
// 		fmt.Fprintf(w, "%s\n", nosurf.Reason(r))
// 		log.Println("==============================================")
// 		log.Println(fmt.Fprintf(w, "%s\n", nosurf.Reason(r)))
// 		log.Println(csrf.Token(r))
// 	})),
// )

// func CSRF() gin.HandlerFunc {
// 	return adapter.Wrap(csrfMd)
// }

//========nosurf========
func csrfFailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", nosurf.Reason(r))
	log.Println("==============================================")
	log.Println(fmt.Fprintf(w, "%s\n", nosurf.Reason(r)))
}

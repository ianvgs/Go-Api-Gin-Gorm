package routes

import (
	"goagain/controllers"
	"goagain/globals"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"goagain/middleware"
)

func HandleRequests() {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.html")
	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := r.Group("/")
	publicRoutes(public)

	publicPosts := r.Group("/")
	postRoutes(publicPosts)

	publicUsers := r.Group("/")
	userRoutes(publicUsers)

	private := r.Group("/")
	private.Use(middleware.AuthRequired)
	privateRoutes(private)

	r.Run()
}

func publicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.LoginGetHandler())
	g.GET("/register", controllers.RegisterGetHandler())
	g.GET("/about", controllers.AboutGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/", controllers.IndexGetHandler())
}

func privateRoutes(g *gin.RouterGroup) {
	g.GET("/dashboard", controllers.DashboardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())

}

// Post module --- PSQL
func postRoutes(g *gin.RouterGroup) {
	g.GET("/posts", controllers.PostIndex)
	g.POST("/posts", controllers.PostCreate)
	g.GET("/posts/:id", controllers.PostShow)
	g.PUT("/posts/:id", controllers.PostUpdate)
	g.PATCH("/posts/:id", controllers.PostUpdater)
	g.DELETE("/posts/:id", controllers.PostUpdater)

}

// User Module MONGODB
func userRoutes(g *gin.RouterGroup) {
	g.POST("/user", controllers.CreateUser())
	g.GET("/user/:userId", controllers.GetAUser())
	g.PUT("/user/:userId", controllers.EditAUser())
	g.DELETE("/user/:userId", controllers.DeleteAUser())
	g.GET("/users", controllers.GetAllUsers())
}

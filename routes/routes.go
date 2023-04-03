package routes

import (
	"goagain/controllers"
	"goagain/globals"
	"goagain/helpers"
	"goagain/middleware"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	//Middleware debugger
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Logger())

	//Funcoes pra serem usadas nos HTML
	r.SetFuncMap(template.FuncMap{
		"ToUpper": strings.ToUpper,
		"Subt":    helpers.Subt,
	})

	//Middleware 404
	r.Use(func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() == 404 {
			c.HTML(http.StatusBadRequest, "404.html", gin.H{"content": "Page not found."})
		}
	})

	//Configs
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
	g.GET("/login", controllers.RenderLoginPage())
	g.GET("/logout", controllers.LogoutGetHandler())
	g.GET("/register", controllers.RegisterGetHandler())
	g.GET("/about", controllers.AboutGetHandler())
	g.POST("/login", controllers.ValidateLogin())
	g.GET("/", controllers.IndexGetHandler())
	g.GET("/category/:id", controllers.CategoryShow())
	g.GET("/news/:id", controllers.NewsShow())
}

func privateRoutes(g *gin.RouterGroup) {
	g.GET("/dashboard", controllers.DashboardGetHandler())
}

// User Module MONGODB
func userRoutes(g *gin.RouterGroup) {
	g.POST("/user", controllers.CreateUser())
}

// Post noticias, não deixar habilitado
func postRoutes(g *gin.RouterGroup) {
	g.GET("/posts", controllers.PostIndex)
	g.GET("/posts/:id", controllers.PostShow)
}

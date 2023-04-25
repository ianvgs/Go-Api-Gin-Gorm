package routes

import (
	"goagain/controllers"
	"goagain/globals"
	"goagain/helpers"
	"goagain/middleware"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

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

	/* if os.Getenv("GO_ENV") == "production" {
		// Get the absolute path to the executable
		executablePath, err := os.Executable()
		if err != nil {
			log.Fatalf("Error getting executable path: %s", err)
		}
		executableDir := filepath.Dir(executablePath)
		r.Static("/assets", filepath.Join(executableDir, "/assets"))
		r.LoadHTMLGlob(filepath.Join(executableDir, "templates/*.html"))

	} */

	if os.Getenv("GO_ENV") != "production" {
		//Middleware debugger
		gin.SetMode(gin.DebugMode)
		r.Use(gin.Logger())

		//Configs

		r.Static("/assets", "./assets")
		r.LoadHTMLGlob("templates/*.html")
	}
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

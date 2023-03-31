package controllers

import (
	"fmt"
	"goagain/globals"
	"goagain/helpers"
	"goagain/initializers"
	"goagain/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Define a custom function that converts a string to uppercase
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func IndexGetHandler() gin.HandlerFunc {
	colors := []string{"tag is-primary is-medium block", "tag is-link is-medium ", "tag is-light is-danger is-medium", "tag is-dark is-medium ", "tag is-success is-medium ", "tag is-warning is-medium "}
	/*
		initializers.DB.Migrator().DropColumn(&models.Noticia{}, "deleted_at") */

	var noticias []models.Noticia
	initializers.DB.Unscoped().Preload("Colaborador").Limit(5).Find(&noticias)

	for _, noticia := range noticias {
		fmt.Println("Título:", noticia)
	}

	var categorias []models.Categoria
	initializers.DB.Find(&categorias)

	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"content":    "This is an index page...",
			"user":       user,
			"posts":      noticias,
			"categorias": categorias,
			"colors":     colors,
		})

	}
}

func AboutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "about.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "about.html", gin.H{
			"content": "",
			"user":    user,
		})
	}
}

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		log.Printf("%v", user)

		if user != nil {
			c.Redirect(http.StatusBadRequest, "/")
			return

		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"user": user,
		})
	}
}

func RegisterGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "register.html", gin.H{
			"content": "",
			"user":    user,
		})
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Please logout first"})
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserPass(username, password) {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(globals.Userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			return
		}

		session.Delete(globals.Userkey)
		if err := session.Save(); err != nil {
			log.Println("Passei aqui")
			log.Println("Failed to save session:", err)
			return
		}
		log.Println("Passei aqui")

		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		log.Println("%v", user)

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content": "This is a dashboard",
			"user":    user,
		})
	}
}

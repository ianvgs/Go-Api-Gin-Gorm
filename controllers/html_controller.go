package controllers

import (
	"fmt"
	"goagain/globals"
	"goagain/initializers"
	"goagain/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IndexGetHandler() gin.HandlerFunc {
	colors := []string{"tag is-primary is-medium block", "tag is-link is-medium ", "tag is-light is-danger is-medium", "tag is-dark is-medium ", "tag is-success is-medium ", "tag is-warning is-medium "}

	var noticias []models.Noticia
	initializers.DB.Unscoped().Preload("Colaborador").Limit(5).Find(&noticias)
	noticias = append(noticias, models.Noticia{
		Id:            1,
		ColaboradorId: 1,
		Titulo:        "mockTitulo",
		Resumo:        "mockResumo",
		Texto:         "mockTexto",
		IdCategoria:   1,
		Ativo:         "S",
		CreatedAt:     time.Date(2023, 4, 2, 10, 30, 0, 0, time.UTC),
		UpdatedAt:     time.Date(2023, 4, 2, 10, 30, 0, 0, time.UTC),
	})

	fmt.Println("Noticias:", noticias)
	/* for _, noticia := range noticias {
		fmt.Println("Título:", noticia.Colaborador)
	} */

	var categorias []models.Categoria
	initializers.DB.Find(&categorias)
	categorias = append(categorias, models.Categoria{
		Id:        1,
		Nome:      "MockCategNome",
		Descricao: "categMock",
	})

	for _, categoria := range categorias {
		fmt.Println("categoria:", categoria)
	}

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
		errorMsg := session.Get(globals.ErrorMsg)

		user := session.Get(globals.Userkey)

		if user != nil {
			c.Redirect(http.StatusBadRequest, "/")
			return

		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"user":     user,
			"errorMsg": errorMsg,
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
		})
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
		/* if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		} */

		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is a dashboard",
			"user":    "",
		})
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content": "This is a dashboard",
			"user":    user,
		})
	}
}

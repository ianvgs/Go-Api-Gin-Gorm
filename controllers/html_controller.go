package controllers

import (
	"fmt"
	"goagain/globals"
	"goagain/initializers"
	"goagain/models"
	"log"
	"net/http"
	"reflect"

	/* 	"time" */

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IndexGetHandler() gin.HandlerFunc {
	//Declara e inicializa as arrays: colors/noticias/categorias com strings e "modelObjects"
	colors := []string{"tag is-primary is-medium block", "tag is-dark is-medium ", "tag is-success is-medium"}
	var noticias []models.Noticia
	var categorias []models.Categoria

	//busca dados das tabelas noticias e categorias
	initializers.DB.Unscoped().Preload("Colaborador").Limit(3).Find(&noticias)
	initializers.DB.Where("idSite = ?", 2).Find(&categorias)

	/* fmt.Println("Noticias:", noticias) */
	/* for _, noticia := range noticias {
		fmt.Println("Título:", noticia.Colaborador)
	} */

	//Mocks de noticias & categorias
	/* categorias = append(categorias, models.Categoria{
		Id:        1,
		Nome:      "Mocked One",
		Descricao: "categMock",
	}) */

	/* noticias = append(noticias, models.Noticia{
		Id:            1,
		ColaboradorId: 1,
		Titulo:        "MockTitulo na controller + 3 records do DB acima, ou erro",
		Resumo:        "mockResumo",
		Texto:         "mockTexto",
		IdCategoria:   1,
		Ativo:         "S",
		CreatedAt:     time.Date(2023, 4, 2, 10, 30, 0, 0, time.UTC),
		UpdatedAt:     time.Date(2023, 4, 2, 10, 30, 0, 0, time.UTC),
	})
	*/
	// Cria um novo array com a quantidade de elementos de categorias com as cores de colors, pode ser assim: colorArray := make([]string, 4)
	colorArray := make([]string, len(categorias))
	for i := 0; i < len(colorArray); i++ {
		colorIndex := i % len(colors)
		colorArray[i] = colors[colorIndex]
	}

	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.UserKey)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"content":    "This is an index page...",
			"user":       user,
			"posts":      noticias,
			"categorias": categorias,
			"colors":     colorArray,
		})

	}
}

func AboutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.UserKey)
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

func CategoryShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var categoryNews []models.Noticia

		initializers.DB.Unscoped().Preload("Colaborador").Preload("Categoria").Where("idCategoria = ?", id).Limit(3).Find(&categoryNews)

		if reflect.DeepEqual(categoryNews, models.Noticia{}) {
			c.Redirect(http.StatusMovedPermanently, "/")
			return
		}

		c.HTML(http.StatusOK, "categoryNews.html", gin.H{
			"content":      "",
			"param_id":     id,
			"categoryNews": categoryNews,
			"categoryName": categoryNews[0].Categoria.Nome,
		})
	}
}

func NewsShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var news models.Noticia
		initializers.DB.First(&news, id)
		/* 	news.FormatCreatedAt() */
		fmt.Println("Formatted date:", news.CreatedAt)

		formattedDate := news.CreatedAt.Format("Jan 02, 2006")

		log.Println("logging out user:", formattedDate)

		if reflect.DeepEqual(news, models.Noticia{}) {
			c.Redirect(http.StatusMovedPermanently, "/")
			return
		}

		c.HTML(http.StatusOK, "newsShow.html", gin.H{
			"content":      "",
			"param_id":     id,
			"news":         news,
			"formatedDate": formattedDate,
		})
	}
}

func RenderLoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.UserKey)

		if user != nil {
			c.Redirect(http.StatusMovedPermanently, "/")
			return

		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "",
		})
	}
}

func RegisterGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.UserKey)
		if user != nil {
			c.Redirect(http.StatusMovedPermanently, "/")
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
		user := session.Get(globals.UserKey)

		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			return
		}
		session.Delete(globals.UserKey)
		/* session.Set(globals.UserKey, "") */
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/")

	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.UserKey)
		fmt.Println("user:", user)
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content": "This is a dashboard",
			"user":    user,
		})
	}
}

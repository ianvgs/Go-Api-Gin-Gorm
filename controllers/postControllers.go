// nome := c.Params.ByName("nome")
// id := c.Params.ByName("id")
// cpf := c.Param("cpf")
package controllers

import (
	"goagain/initializers"
	"goagain/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	body := models.Noticia{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	res := initializers.DB.Create(&body)

	if res.Error != nil {
		c.Status(400)
		log.Fatal("Erro ao salvar")
	}

	c.JSON(http.StatusOK, gin.H{
		"post": body,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Noticia
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func PostShow(c *gin.Context) {

	/* nome := c.Params.ByName("nome") */

	id := c.Param("id")

	var post models.Noticia
	initializers.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}

func PostUpdate(c *gin.Context) {

	id := c.Param("id")
	println(id)

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post models.Noticia
	initializers.DB.First(&post, id)

	println(body.Title)

	initializers.DB.Model(&post).Updates(models.Noticia{
		Titulo: body.Title, Resumo: body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"post": body,
	})

}

func PostUpdater(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Titulo string
		Resumo string
	}

	c.Bind(&body)

	var post models.Noticia
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Noticia{
		Titulo: body.Titulo, Resumo: body.Resumo,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Updeitado",
		"post":    body,
	})

}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Noticia{}, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleitado",
	})

}

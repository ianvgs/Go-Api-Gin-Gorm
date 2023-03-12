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
	body := models.Post{}
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
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func PostShow(c *gin.Context) {

	/* nome := c.Params.ByName("nome") */

	id := c.Param("id")

	var post models.Post
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

	var post models.Post
	initializers.DB.First(&post, id)

	println(body.Title)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"post": body,
	})

}

func PostUpdater(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Updeitado",
		"post":    body,
	})

}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleitado",
	})

}

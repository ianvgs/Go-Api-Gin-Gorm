package routes

import (
	"goagain/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/posts", controllers.PostIndex)

	r.POST("/posts", controllers.PostCreate)

	r.POST("/user", controllers.CreateUser())
	r.GET("/user/:userId", controllers.GetAUser())
	r.PUT("/user/:userId", controllers.EditAUser())
	r.DELETE("/user/:userId", controllers.DeleteAUser())
	r.GET("/users", controllers.GetAllUsers())

	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.PATCH("/posts/:id", controllers.PostUpdater)
	r.DELETE("/posts/:id", controllers.PostUpdater)
	r.Run()
}

package routes

import (
	"goagain/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/posts", controllers.PostIndex)
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.PATCH("/posts/:id", controllers.PostUpdater)
	r.DELETE("/posts/:id", controllers.PostUpdater)
	r.Run()
}

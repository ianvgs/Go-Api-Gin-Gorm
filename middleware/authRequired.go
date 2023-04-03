package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

/* The middleware for checking if a user has been authenticated is implemented in this file. It simply checks to see if the session exists. If the session is nil, the user is redirected to the /login route. Otherwise, the Gin framework is told to move to the next step. */

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		log.Println("User not logged in")

		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}
	c.Next()
}

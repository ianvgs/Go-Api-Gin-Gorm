package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"goagain/globals"
	"log"
	"net/http"
)

/* The middleware for checking if a user has been authenticated is implemented in this file. It simply checks to see if the session exists. If the session is nil, the user is redirected to the /login route. Otherwise, the Gin framework is told to move to the next step. */

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.UserKey)
	if user == nil {
		log.Println("User not logged in: user:", user)
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}

	if user == "" {
		log.Println("User not logged in: user:", user)
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}
	c.Next()
}

package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"goagain/globals"
	"log"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.UserKey)
	if user == nil || user == "" {
		log.Println("User not logged in: user:", user)
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}
}

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
)


func loginGet(c *gin.Context) {
	session := sessions.Default(c)

	if user := session.Get("user"); user != nil {
		c.Redirect(http.StatusMovedPermanently, "/mypage")
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{})
}
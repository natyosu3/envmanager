package mypage

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
)

func myapageGet(c *gin.Context) {
	session := sessions.Default(c)

	if user := session.Get("user"); user == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	c.HTML(http.StatusOK, "mypage.html", gin.H{})
}
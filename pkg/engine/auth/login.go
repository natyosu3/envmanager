package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func loginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

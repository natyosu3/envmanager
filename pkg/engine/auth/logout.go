package auth

import (
	"envmanager/pkg/session"
	"net/http"
	"github.com/gin-gonic/gin"
)

func logoutGet(c *gin.Context) {
	session.DeleteSession(c, "session")
	c.Redirect(http.StatusFound, "/")
}
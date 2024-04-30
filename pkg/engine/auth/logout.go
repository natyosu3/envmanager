package auth

import (
	"envmanager/pkg/session"
	"net/http"
	"github.com/gin-gonic/gin"
	"envmanager/pkg/model"
)

func logoutGet(c *gin.Context) {
	session.Default(c, "session", &model.Session_model{}).Delete(c)
	c.Redirect(http.StatusFound, "/")
}
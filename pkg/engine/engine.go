package engine

import (
	"github.com/gin-gonic/gin"
	"envmanager/pkg/engine/auth"
)



func Engine(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("web/templates/*/*.html")
	r.Static("/static", "web/static")

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/login", auth.LoginGet())
		authGroup.GET("/signup", auth.SignupGet())
		authGroup.POST("/signup", auth.SignupPost())
	}

	return r
}
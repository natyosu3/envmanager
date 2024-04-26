package engine

import (
	"envmanager/pkg/engine/auth"
	"envmanager/pkg/engine/top"
	"envmanager/pkg/engine/service"

	"github.com/gin-gonic/gin"
)



func Engine(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("web/templates/*/*.html")
	r.Static("/static", "web/static")

	r.GET("/", top.TopGet())
	authGroup := r.Group("/auth")
	{
		authGroup.GET("/login", auth.LoginGet())
		authGroup.POST("/login", auth.LoginPost())
		authGroup.GET("/logout", auth.LogoutGet())
		authGroup.GET("/signup", auth.SignupGet())
		authGroup.POST("/signup", auth.SignupPost())
	}
	serviceGroup := r.Group("/service")
	{	
		serviceGroup.GET("/dashboard", service.DashboardGet())
		serviceGroup.POST("/delete", service.DeleteServicePost())
		serviceGroup.GET("/:id", service.DetailGet())
		serviceGroup.GET("/edit/:id", service.EditServiceGet())
		serviceGroup.POST("/create", service.ServiceCreatePost())
	}

	return r
}
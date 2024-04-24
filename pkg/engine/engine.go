package engine

import (
	"envmanager/pkg/engine/auth"
	"envmanager/pkg/engine/mypage"
	"envmanager/pkg/engine/service"

	"github.com/gin-gonic/gin"
)



func Engine(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("web/templates/*/*.html")
	r.Static("/static", "web/static")

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/login", auth.LoginGet())
		authGroup.POST("/login", auth.LoginPost())
		authGroup.GET("/signup", auth.SignupGet())
		authGroup.POST("/signup", auth.SignupPost())
	}
	mypageGroup := r.Group("/mypage")
	{
		mypageGroup.GET("/", mypage.MypageGet())
	}
	serviceGroup := r.Group("/service")
	{	
		serviceGroup.GET("/dashboard", service.DashboardGet())
		serviceGroup.GET("/:id", service.DetailGet())
		serviceGroup.POST("/create", service.ServiceCreatePost())
	}

	return r
}
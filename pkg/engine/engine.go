package engine

import (
	"envmanager/pkg/engine/auth"
	"envmanager/pkg/engine/mypage"

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

	return r
}
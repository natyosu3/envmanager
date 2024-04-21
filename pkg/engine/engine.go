package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
	"log/slog"
	"os"
	"github.com/joho/godotenv"
	"envmanager/pkg/engine/auth"
)


var (
	REDIS_HOST string
	REDIS_PASSWORD string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}
	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
}


func Engine(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("web/templates/*/*.html")
	r.Static("/static", "web/static")

	store, err := redis.NewStore(10, "tcp", REDIS_HOST, REDIS_PASSWORD, []byte("secret"))
	if err != nil {
		slog.Error(err.Error())
	}
	r.Use(sessions.Sessions("session", store))

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/login", auth.LoginGet())
		authGroup.GET("/signup", auth.SignupGet())
		authGroup.POST("/signup", auth.SignupPost())
	}

	return r
}
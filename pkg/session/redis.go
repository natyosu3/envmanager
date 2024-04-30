package session

import (
	"encoding/json"
	"envmanager/pkg/general/random"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}
	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")

	client = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: REDIS_PASSWORD,
		DB:       0,
	})
}

type Session struct {
	SessionId string
	CookieKey string
	Model     interface{}
}

func Default(c *gin.Context, cookieKey string, model interface{}) *Session {
	SessionId, err := c.Cookie(cookieKey)
	if err != nil {
		SessionId := random.MakeRandomStringId()
		new(c, SessionId, cookieKey, model)
		return &Session{SessionId: SessionId, CookieKey: cookieKey, Model: model}
	}
	return &Session{SessionId: SessionId, CookieKey: cookieKey, Model: model}
}

func new(c *gin.Context, SessionId string, cookieKey string, value interface{}) {
	valueByte, err := json.Marshal(value)
	if err != nil {
		slog.Error("Error setting session: " + err.Error())
		return
	}
	client.Set(c, SessionId, string(valueByte), 24*30*time.Hour)
	c.SetCookie(cookieKey, SessionId, 0, "/", "", false, true)
}

func (s *Session) Set(c *gin.Context, value interface{}) {
	valueByte, err := json.Marshal(value)
	if err != nil {
		slog.Error("Error setting session: " + err.Error())
		return
	}
	client.Set(c, s.SessionId, string(valueByte), 24*30*time.Hour)
}

func (s *Session) Get(c *gin.Context) interface{} {
	SessionId, err := c.Cookie(s.CookieKey)
	if err != nil {
		return nil
	}

	value, err := client.Get(c, SessionId).Bytes()
	if err != nil {
		slog.Error("Error getting session: " + err.Error())
		return nil
	}

	err = json.Unmarshal(value, s.Model)
	if err != nil {
		slog.Error("Error getting session: " + err.Error())
		return nil
	}

	return s.Model
}

func (s *Session) Delete(c *gin.Context) {
	client.Del(c, s.SessionId)
	c.SetCookie(s.CookieKey, "", -1, "/", "", false, true)
}

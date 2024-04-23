package session

import (
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

func NewSession(c *gin.Context, cookieKey string, value []byte) {
  SessionId := random.MakeRandomId()


  client.Set(c, SessionId, string(value), 24*30*time.Hour)
  c.SetCookie(cookieKey, SessionId, 0, "/", "", false, true)
}

func SetSession(c *gin.Context, cookieKey string, value interface{}) {
  SessionId, err := c.Cookie(cookieKey)
  if err != nil {
    return
  }
  client.Set(c, SessionId, value, 24*30*time.Hour)
}

func GetSession(c *gin.Context, cookieKey string) []byte {
	SessionId, err := c.Cookie(cookieKey)
  if err != nil {
    return nil
  }
  value, err := client.Get(c, SessionId).Bytes()
  if err != nil {
    slog.Error("Error getting session: " + err.Error())
    return nil
  }
  return value
}

func DeleteSession(c *gin.Context, cookieKey string) {
  SessionId, err := c.Cookie(cookieKey)
  if err != nil {
    return
  }
  client.Del(c, SessionId)
  c.SetCookie(cookieKey, "", -1, "/", "", false, true)
}
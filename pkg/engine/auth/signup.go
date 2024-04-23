package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"envmanager/pkg/general/encrypt"
	"envmanager/pkg/session"
	"envmanager/pkg/db/create"
	"log/slog"
)


func signupGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}


func signupPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	// パスワードをハッシュ化
	hash, err := encrypt.PasswordEncrypt(password)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Encrypt Error"})
		return
	}

	err = create.CreateUser(username, hash, email)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Create User Error"})
		return
	}

	session.NewSession(c, "session", username)

	c.Redirect(http.StatusMovedPermanently, "/mypage")
}
package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"envmanager/pkg/general/encrypt"
	"envmanager/pkg/db/create"
	"log/slog"
)


func signupGet(c *gin.Context) {
	session := sessions.Default(c)

	if login_st := session.Get("login_st"); login_st != nil {
		c.Redirect(http.StatusMovedPermanently, "/mypage")
		return
	}


	c.HTML(http.StatusOK, "signup.html", gin.H{})
}


func signupPost(c *gin.Context) {
	session := sessions.Default(c)

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

	session.Set("login_st", true)


	c.Redirect(http.StatusMovedPermanently, "/mypage")
}
package auth

import (
	"envmanager/pkg/db/read"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)


func loginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func loginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	slog.Info("Username: " + username)
	slog.Info("Password: " + password)

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password is empty"})
		return
	} else {
		err := read.ReadUser(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{ 
				"error": "User not found", 
				"message": "Please check your username and password",
				"err": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User found"})
		return
	}
}
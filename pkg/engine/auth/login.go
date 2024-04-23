package auth

import (
	"envmanager/pkg/db/read"
	"envmanager/pkg/general/encrypt"
	"envmanager/pkg/session"
	"envmanager/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func loginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func loginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password is empty"})
		return
	} else {
		user, err := read.ReadUser(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{ 
				"error": "User not found", 
				"message": "Please check your username and password",
				"err": err,
			})
			return
		}
		err = encrypt.CompareHashAndPassword(user.Password, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password is incorrect"})
			return
		}
		
		session_info, err := model.Session_model{
			Userid: user.Userid,
			Username: username,
			Logined: true,
		}.Json()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Session error"})
			return
		}
		session.NewSession(c, "session", session_info)
	}
}
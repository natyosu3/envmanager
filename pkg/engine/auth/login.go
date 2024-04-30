package auth

import (
	"envmanager/pkg/db/read"
	"envmanager/pkg/general/encrypt"
	"envmanager/pkg/model"
	"envmanager/pkg/session"
	"net/http"

	"github.com/gin-gonic/gin"
)




func loginGet(c *gin.Context) {
	if data := session.Default(c, "session", &model.Session_model{}).Get(c); data == nil || data.(*model.Session_model).Userid == "" {
		c.HTML(http.StatusOK, "login.html", gin.H{})
		return
	} else {
		data := data.(*model.Session_model)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"session":         data,
			"IsAuthenticated": data.Logined,
		})
		return
	}
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
		
		session_info := model.Session_model{
			Userid: user.Userid,
			Username: username,
			Logined: true,
		}
		session.Default(c, "session", &model.Session_model{}).Set(c, session_info)
		c.Redirect(http.StatusSeeOther, "/service/dashboard")
		return
	}
}
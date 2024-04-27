package auth

import (
	"envmanager/pkg/db/common"
	"envmanager/pkg/db/create"
	"envmanager/pkg/db/read"
	"envmanager/pkg/general/encrypt"
	"envmanager/pkg/model"
	"envmanager/pkg/session"
	"log/slog"
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
)


func signupGet(c *gin.Context) {
	var session_info model.Session_model
	session_data := session.GetSession(c, "session")
	if session_data == nil {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"IsAuthenticated": session_info.Logined,
		})
		return
	} else {
		err := json.Unmarshal(session_data, &session_info)
		if err != nil {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "Session Convert Json Error"})
			return
		}
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"IsAuthenticated": session_info.Logined,
		})
		return
	}
}


func signupPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	// ユーザが既に存在するか確認
	exsist := common.ExsistUser(username)
	if exsist {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exsist"})
		return
	}

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

	user, _ := read.ReadUser(username)
	session_info := model.Session_model{
		Userid: user.Userid,
		Username: username,
		Logined: true,
	}
	jbyte, err := session_info.Json()
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Session Convert Json Error"})
		return
	}
	session.NewSession(c, "session", jbyte)

	c.Redirect(http.StatusMovedPermanently, "/service/dashboard")
}
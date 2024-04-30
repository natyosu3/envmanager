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

	"github.com/gin-gonic/gin"
)

func signupGet(c *gin.Context) {
	if data := session.Default(c, "session", &model.Session_model{}).Get(c); data == nil || data.(*model.Session_model).Userid == "" {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
		return
	} else {
		data := data.(*model.Session_model)
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"session":         data,
			"IsAuthenticated": data.Logined,
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
		Userid:   user.Userid,
		Username: username,
		Logined:  true,
	}
	session.Default(c, "session", &model.Session_model{}).Set(c, session_info)
	c.Redirect(http.StatusMovedPermanently, "/service/dashboard")
}

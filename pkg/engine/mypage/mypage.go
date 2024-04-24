package mypage

import (
	"encoding/json"
	"envmanager/pkg/model"
	"envmanager/pkg/session"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func myapageGet(c *gin.Context) {
	var session_info model.Session_model
	session_data := session.GetSession(c, "session")
	if session_data == nil {
		c.Redirect(http.StatusSeeOther, "/auth/login")
		return
	}

	err := json.Unmarshal(session_data, &session_info)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Session Convert Json Error"})
		return
	}

	c.HTML(http.StatusOK, "mypage.html", gin.H{
		"session": session_info,
		"IsAuthenticated": session_info.Logined,
	})
}
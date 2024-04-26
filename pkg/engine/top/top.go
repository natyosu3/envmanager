package top

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"envmanager/pkg/session"
	"envmanager/pkg/model"
	"encoding/json"
	"log/slog"
)


func indexGet(c *gin.Context) {
	var session_info model.Session_model
	session_data := session.GetSession(c, "session")
	if session_data != nil {
		err := json.Unmarshal(session_data, &session_info)
		if err != nil {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "Session Convert Json Error"})
			return
		}
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"session":         session_info,
		"IsAuthenticated": session_info.Logined,
	})
}
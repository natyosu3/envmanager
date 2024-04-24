package service

import (
	"encoding/json"
	"envmanager/pkg/model"
	"envmanager/pkg/session"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serviceCreatePost(c *gin.Context) {
	service_name := c.PostForm("service_name")
	env_names := c.PostFormArray("env_name")
	env_values := c.PostFormArray("env_value")

	fmt.Println(service_name)
	fmt.Println(env_names)
	fmt.Println(env_values)
}


func serviceGet(c *gin.Context) {
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

	c.HTML(http.StatusOK, "product.html", gin.H{
		"session": session_info,
	})
}
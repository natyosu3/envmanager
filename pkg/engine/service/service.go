package service

import (
	"encoding/json"
	"envmanager/pkg/db/read"
	"envmanager/pkg/model"
	"envmanager/pkg/session"
	"fmt"
	"log/slog"
	"net/http"

	"envmanager/pkg/db/create"

	"github.com/gin-gonic/gin"
)

func dashboardGet(c *gin.Context) {
	var session_info model.Session_model
	session_data := session.GetSession(c, "session")
	if session_data == nil {
		c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	err := json.Unmarshal(session_data, &session_info)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Session Convert Json Error"})
	}

	service_list, err := read.ReadService(session_info.Userid)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Read Service Error"})
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"session": session_info,
		"IsAuthenticated": session_info.Logined,
		"userid": session_info.Userid,
		"env_data": service_list,
	})
}


func detailGet(c *gin.Context) {
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

	service_id := c.Param("id")
	owner, err := read.CheckOwner(session_info.Userid, service_id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Check Owner Error"})
		return
	}
	if !owner {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not Owner"})
		return
	}

	fmt.Println(service_id)
}


func serviceCreatePost(c *gin.Context) {
	type Data struct {
		ServiceId string `json:"service_id"`
		ServiceName string `json:"service_name"`
		Data string `json:"data"`
	}
	var data Data
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session_data := session.GetSession(c, "session")
	if session_data == nil {
		c.Redirect(http.StatusSeeOther, "/auth/login")
		return
	}

	var session_info model.Session_model
	err := json.Unmarshal(session_data, &session_info)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Session Convert Json Error"})
		return
	}

	create.CreateService(session_info.Userid, data.ServiceId, data.ServiceName, data.Data)

	c.Redirect(http.StatusSeeOther, "/service/dashboard")
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
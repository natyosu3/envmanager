package service

import (
	"encoding/json"
	"envmanager/pkg/db/read"
	"envmanager/pkg/db/common"
	"envmanager/pkg/model"
	"envmanager/pkg/session"
	"log/slog"
	"net/http"

	"envmanager/pkg/db/create"
	"envmanager/pkg/db/delete"

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
		"session":         session_info,
		"IsAuthenticated": session_info.Logined,
		"userid":          session_info.Userid,
		"env_data":        service_list,
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
	owner, err := common.CheckOwner(session_info.Userid, service_id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Check Owner Error"})
		return
	}
	if !owner {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not Owner"})
		return
	}

	service_name, envs, err := read.ReadServiceDetail(service_id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Read Service Detail Error"})
		return
	}

	c.HTML(http.StatusOK, "detail.html", gin.H{
		"session": session_info,
		"service_name": service_name,
		"env_data": envs,
		"IsAuthenticated": session_info.Logined,
	})
}

func serviceCreatePost(c *gin.Context) {
	type Data struct {
		ServiceId   string `json:"service_id"`
		ServiceName string `json:"service_name"`
		Data        string `json:"data"`
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


func deleteServicePost(c *gin.Context) {
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

	service_id := c.PostForm("service_id")

	err = delete.DeleteService(service_id, session_info.Userid)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Delete Service Error"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/service/dashboard")
}


func editServiceGet(c *gin.Context) {
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

	service_id := c.Param("id")
	owner, err := common.CheckOwner(session_info.Userid, service_id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Check Owner Error"})
		return
	}
	if !owner {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not Owner"})
		return
	}

	service_name, envs, err := read.ReadServiceDetail(service_id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Read Service Detail Error"})
		return
	}

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"session": session_info,
		"service_name": service_name,
		"service_id": service_id,
		"env_data": envs,
		"IsAuthenticated": session_info.Logined,
	})
}
package service

import (
	"envmanager/pkg/db/common"
	"envmanager/pkg/db/read"
	"envmanager/pkg/model"
	"envmanager/pkg/session"
	"log/slog"
	"net/http"

	"envmanager/pkg/db/create"
	"envmanager/pkg/db/delete"
	"envmanager/pkg/db/update"

	"github.com/gin-gonic/gin"
)

func dashboardGet(c *gin.Context) {
	if data := session.Default(c, "session", &model.Session_model{}).Get(c); data == nil || data.(*model.Session_model).Userid == "" {
		c.Redirect(http.StatusSeeOther, "/auth/login")
		return
	} else {
		data := data.(*model.Session_model)

		service_list, err := read.ReadService(data.Userid)
		if err != nil {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "Read Service Error"})
		}

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"IsAuthenticated": data.Logined,
			"userid":          data.Userid,
			"env_data":        service_list,
		})
	}
}

func detailGet(c *gin.Context) {
	if data := session.Default(c, "session", &model.Session_model{}).Get(c); data == nil || data.(*model.Session_model).Userid == "" {
		c.Redirect(http.StatusSeeOther, "/auth/login")
		return
	} else {
		data := data.(*model.Session_model)

		service_id := c.Param("id")
		owner, err := common.CheckOwner(data.Userid, service_id)

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
			"service_name":    service_name,
			"env_data":        envs,
			"IsAuthenticated": data.Logined,
		})
	}
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

	if session_data := session.Default(c, "session", &model.Session_model{}).Get(c); session_data == nil || session_data.(*model.Session_model).Userid == "" {
		c.Redirect(http.StatusSeeOther, "/auth/login")
		return
	} else {
		session_data := session_data.(*model.Session_model)
		create.CreateService(session_data.Userid, data.ServiceId, data.ServiceName, data.Data)

		c.Redirect(http.StatusSeeOther, "/service/dashboard")
	}
}

func deleteServicePost(c *gin.Context) {
	if data := session.Default(c, "session", &model.Session_model{}).Get(c); data == nil || data.(*model.Session_model).Userid == "" {
		c.Redirect(http.StatusSeeOther, "/auth/login")
		return
	} else {
		data := data.(*model.Session_model)

		service_id := c.PostForm("service_id")
		err := delete.DeleteService(service_id, data.Userid)
		if err != nil {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "Delete Service Error"})
			return
		}
		c.Redirect(http.StatusSeeOther, "/service/dashboard")
	}
}

func editServiceGet(c *gin.Context) {
	if data := session.Default(c, "session", &model.Session_model{}).Get(c); data == nil || data.(*model.Session_model).Userid == "" {
		c.Redirect(http.StatusSeeOther, "/auth/login")
		return
	} else {
		data := data.(*model.Session_model)

		service_id := c.Param("id")
		owner, err := common.CheckOwner(data.Userid, service_id)
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
			"service_name":    service_name,
			"service_id":      service_id,
			"env_data":        envs,
			"IsAuthenticated": data.Logined,
		})
	}
}

func updateServicePost(c *gin.Context) {
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

	if session_data := session.Default(c, "session", &model.Session_model{}).Get(c); session_data == nil || session_data.(*model.Session_model).Userid == "" {
		c.Redirect(http.StatusSeeOther, "/auth/login")
		return
	} else {
		session_data := session_data.(*model.Session_model)

		owner, err := common.CheckOwner(session_data.Userid, data.ServiceId)
		if err != nil {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "Check Owner Error"})
			return
		}
		if !owner {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Not Owner"})
			return
		}

		err = update.UpdateService(data.ServiceId, data.Data, session_data.Userid)
		if err != nil {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "Update Service Error"})
			return
		}
		c.Redirect(http.StatusSeeOther, "/service/dashboard")
	}
}

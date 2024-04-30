package top

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"envmanager/pkg/model"
	"envmanager/pkg/session"
)

func indexGet(c *gin.Context) {
	if data := session.Default(c, "session", &model.Session_model{}).Get(c); data == nil || data.(*model.Session_model).Userid == "" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"IsAuthenticated": false,
		})
		return
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"IsAuthenticated": data.(*model.Session_model).Logined,
			"userid":          data.(*model.Session_model).Userid,
		})
		return
	}
}

package mypage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func myapageGet(c *gin.Context) {
	c.HTML(http.StatusOK, "mypage.html", gin.H{})
}
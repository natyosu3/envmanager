package test

import (
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}
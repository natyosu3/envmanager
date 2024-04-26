package top

import (
	"github.com/gin-gonic/gin"
)

func TopGet() gin.HandlerFunc {
	return indexGet
}
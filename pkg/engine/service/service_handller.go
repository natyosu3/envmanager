package service

import (
	"github.com/gin-gonic/gin"
)

func ServiceCreatePost() gin.HandlerFunc {
	return serviceCreatePost
}


func ServiceGet() gin.HandlerFunc {
	return serviceGet
}


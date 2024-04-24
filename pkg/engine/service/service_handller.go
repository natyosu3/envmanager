package service

import (
	"github.com/gin-gonic/gin"
)


func DashboardGet() gin.HandlerFunc {
	return dashboardGet
}

func DetailGet() gin.HandlerFunc {
	return detailGet
}

func ServiceCreatePost() gin.HandlerFunc {
	return serviceCreatePost
}

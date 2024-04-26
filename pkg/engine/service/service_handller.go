package service

import (
	"github.com/gin-gonic/gin"
)


func DashboardGet() gin.HandlerFunc {
	return dashboardGet
}

func DeleteServicePost() gin.HandlerFunc {
	return deleteServicePost
}

func DetailGet() gin.HandlerFunc {
	return detailGet
}

func ServiceCreatePost() gin.HandlerFunc {
	return serviceCreatePost
}

func EditServiceGet() gin.HandlerFunc {
	return editServiceGet
}

func UpdateServicePost() gin.HandlerFunc {
	return updateServicePost
}
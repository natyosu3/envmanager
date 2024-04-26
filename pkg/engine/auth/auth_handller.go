package auth

import (
	"github.com/gin-gonic/gin"
)

func LoginGet() gin.HandlerFunc {
	return loginGet
}

func LoginPost() gin.HandlerFunc {
	return loginPost
}

func LogoutGet() gin.HandlerFunc {
	return logoutGet
}


func SignupGet() gin.HandlerFunc {
	return signupGet
}

func SignupPost() gin.HandlerFunc {
	return signupPost
}
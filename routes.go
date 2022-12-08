package main

import (
	"github.com/say-bonjour/be-say-bonjour/api"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	r.POST("/api/public/login", api.Login)
	r.POST("/api/public/token/refresh", api.RefreshToken)
	r.GET("/api/public/monitor/ping", api.Ping)

	r.GET("/api/users", api.FindUsers)
	r.POST("/api/public/register", api.Register)
	r.GET("/api/users/me", api.MyUser)
	r.GET("/api/users/:id", api.FindUser)
	r.PUT("/api/users/:id", api.UpdateUser)
	r.DELETE("/api/users/:id", api.DeleteUser)
}

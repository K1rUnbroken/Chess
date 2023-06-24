package main

import (
	"chess/api"
	"chess/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	go service.RunHub()
	go service.RunGameHub()

	user := r.Group("/user")
	{
		user.POST("/login", api.Login)
		user.POST("/register", api.Register)
		// 连接大厅
		user.GET("/connect", api.Connect)
	}

	r.Run()
}

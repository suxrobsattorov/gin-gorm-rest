package main

import (
	"gin-gorm-rest/config"
	"gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	err := router.Run(":8081")
	if err != nil {
		return
	}
}

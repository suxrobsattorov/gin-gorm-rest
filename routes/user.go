package routes

import (
	"gin-gorm-rest/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/", controller.CreateUser)
	router.GET("/", controller.GetUsers)
	router.GET("/:id", controller.GetUserById)
	router.PUT("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}

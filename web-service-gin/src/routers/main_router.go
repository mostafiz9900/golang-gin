package routers

import (
	"mostafiz9900/web-service-gin/src/controllers"

	"github.com/gin-gonic/gin"
)

func MainRouter(router *gin.Engine) {
	userRepo := controllers.New()
	ro := router.Group("/api/v1")
	ro.GET("/main", controllers.GetAllUser())
	ro.GET("/create", userRepo.CreateUser)
	ro.GET("/users", userRepo.GetUsers)
	ro.POST("/insert", userRepo.PostUserInfo)

}

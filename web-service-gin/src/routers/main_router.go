package routers

import "github.com/gin-gonic/gin"

func MainRouter(router *gin.Engine) {
	ro := router.Group("api/v1")
	ro.GET("/main", func(cxt *gin.Context) {
		cxt.JSON(200, gin.H{
			"satus": "success",
			"code":  "Go lang gin mysql api",
		})
	})

}

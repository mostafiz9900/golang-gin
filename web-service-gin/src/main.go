package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default
	r.Get("/", func(cxt *gin.Context) {
		cxt.JSON(200, gin.H{
			"satus": "Success",
			"code":  "Go lang gin mysql api",
		})
	})
}

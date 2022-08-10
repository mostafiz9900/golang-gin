package main

import (
	"mostafiz9900/web-service-gin/src/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	routers.MainRouter(route)
	route.Run()
}
// create a new user

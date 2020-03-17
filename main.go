package main

import (
	"github.com/gin-gonic/gin"
	"helloworld/routes"
)

func main() {
	r := gin.Default()
	routes.InitRoutes(r)
	r.Run()
}

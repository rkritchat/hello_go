package routes

import (
	"github.com/gin-gonic/gin"
	"helloworld/services"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/get", services.GetStudent)
	r.POST("/update", services.UpdateStudent)
}
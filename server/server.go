package server

import (
	"assignment3/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/get", controller.GetStatusData)

	r.Run(":8080")
}

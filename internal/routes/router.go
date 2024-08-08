package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/video-validator/internal/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controllers.Ping)

	return r
}

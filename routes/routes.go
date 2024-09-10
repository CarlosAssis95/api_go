package routes

import (
	"integracaomobilemed/controller"

	"github.com/gin-gonic/gin"
)

func SetupRota() *gin.Engine {
	r := gin.Default()
	r.POST("/dados", controller.AdicionarDados)

	return r
}

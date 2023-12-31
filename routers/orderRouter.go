package routers

import (
	"github.com/gin-gonic/gin"
	orderController "github.com/novitaekasari/assigment2-go/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/orders", orderController.Index)
	v1.GET("/orders/:id", orderController.Show)
	v1.POST("/orders", orderController.Create)
	v1.PUT("/orders/:id", orderController.Update)
	v1.DELETE("/orders/:id", orderController.Delete)

	return router
}



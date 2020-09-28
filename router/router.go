package router

import (

	"github.com/gin-gonic/gin"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/controller"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/controller/get_store_by_id_controller"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controller.HomeController)
	router.GET("/ping", controller.PingController)
	router.GET("/error", controller.ErrorController)

	router.GET("/get-store-by-id/:id", get_store_by_id_controller.GetStoreByIdViaGin)

}
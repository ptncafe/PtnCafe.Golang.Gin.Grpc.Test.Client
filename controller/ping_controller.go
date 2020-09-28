package controller

import (

	"github.com/gin-gonic/gin"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/infrastructure/app_logger"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/infrastructure/common_response"
	"os"
)
func PingController(ctx *gin.Context){
	env := os.Environ()
	app_logger.Logger.Infof("ping %v", env)
	//log.Infof("ping %v", env)
	common_response.JsonOk(ctx, env)
}
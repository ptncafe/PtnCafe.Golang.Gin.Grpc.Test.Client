package controller

import (

	"github.com/gin-gonic/gin"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/infrastructure/common_response"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/model/app_error"
)
func ErrorController(ctx *gin.Context){
	var err = app_error.NewInternalErrorMessage("Đã có lỗi xảy ra, bạn vui lòng quay lại sau.")
	common_response.JsonError(ctx, err)
}
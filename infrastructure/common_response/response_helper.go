package common_response

import (

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/infrastructure/app_logger"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/model"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/model/app_error"
)

func JsonOk(ctx *gin.Context, data interface{} ){
	ctx.JSON(200, model.NewResponseDto(data))
	return
}

func JsonError(ctx *gin.Context, err app_error.AppError) {
	switch  err.TypeError {
		case app_error.BadRequestError:
			JsonBadRequestError(ctx, err)
		case app_error.InternalError:
			JsonInternalError(ctx, err)
		default:
			JsonInternalError(ctx, err)
	}
}

func JsonInternalError(ctx *gin.Context, err app_error.AppError){
	//app_logger.Logger.Errorf("%+v",errors.Wrap(err,""))
	app_logger.Logger.Errorf("%v %+v", ctx.Request.URL, errors.Wrap(err.DetailError,ctx.Request.RequestURI))
	ctx.JSON(500, model.NewErrorResponseDto(err))
	return
}

func JsonBadRequestError(ctx *gin.Context, err app_error.AppError){
	app_logger.Logger.Warnf("%v %+v", ctx.Request.URL, errors.Wrap(err.DetailError,ctx.Request.RequestURI))
	ctx.JSON(400, model.NewErrorResponseDto(err))
	return
}
func JsonBadRequestErrorMessaging(ctx *gin.Context, message string){
	app_logger.Logger.Warnf("%v %v", ctx.Request.URL, message)
	ctx.JSON(400, model.NewErrorResponseDtoMessage(message))
	return
}

func JsonUnauthorizedError(ctx *gin.Context){
	ctx.JSON(401, model.NewErrorResponseDtoMessage("Bạn không có quyền truy cập!"))
	return
}
func JsonNotFoundError(ctx *gin.Context){
	ctx.JSON(404, model.NewErrorResponseDtoMessage("Hệ thống không tìm thấy yêu cầu"))
	return
}


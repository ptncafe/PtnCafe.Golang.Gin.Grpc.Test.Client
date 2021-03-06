package get_store_by_id_controller

import (
	"github.com/gin-gonic/gin"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/infrastructure/app_logger"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/infrastructure/common_response"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/model/app_error"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/module/get_store_by_id_module"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/provider"
	"strconv"
)

func GetStoreByIdViaGin(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, _:= strconv.Atoi(idString)
	//if err != nil {
	//	common_response.JsonBadRequestErrorMessaging(ctx,"Dữ liệu không phù hợp")
	//	return;
	//}
	log:= app_logger.Logger
	shopServicesClient, err := provider.NewShopServicesClient()
	if err != nil {
		common_response.JsonError(ctx,app_error.NewInternalError(err))
		return;
	}
	module := get_store_by_id_module.NewGetStoreByIdMudule(log,shopServicesClient )
	data, errMod := module.GetStoreById(ctx.Request.Context(), int32(id))
	if errMod != nil {
		common_response.JsonError(ctx, *errMod)
		return
	}
	common_response.JsonOk(ctx, data)
	return
}
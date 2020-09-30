package provider

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/infrastructure/app_logger"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/model/app_error"
	store "github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

var grpcClientConn *grpc.ClientConn

func NewShopServicesClient() (store.ShopServiceClient, error) {
	var err error
	if grpcClientConn == nil {
		log := app_logger.Logger

		grpcClientConn, err = grpc.Dial(os.Getenv("GRPC_SHOP_SERVICES_URL"), grpc.WithInsecure())
		log.Infof("NewShopServicesClient %s", os.Getenv("GRPC_SHOP_SERVICES_URL"))

		if err != nil {
			log.Errorf("NewShopServicesClient %+v", errors.Wrap(err, "NewShopServicesClient"))
			return nil, err
		}
	}

	//defer conn.Close()

	client := store.NewShopServiceClient(grpcClientConn)
	return client, nil
}

func GrpcErrorToAppError(err error, log *logrus.Logger) *app_error.AppError{
	resErr, ok := status.FromError(err)
	if ok{
		if resErr.Code() == codes.InvalidArgument {
			log.Warnf("GrpcErrorToAppError %s %+v",resErr.Message(), errors.Wrap(resErr.Err() ,"GrpcErrorToAppError") )
			return app_error.NewInvalidErrorNil(resErr.Message())
		}
		log.Errorf("GrpcErrorToAppError %s %+v",resErr.Message(),errors.Wrap(resErr.Err() ,"GrpcErrorToAppError"))
		return  app_error.NewInternalErrorNil(resErr.Err())
	} else{
		log.Errorf("GrpcErrorToAppError %s %+v",resErr.Message(), errors.Wrap(resErr.Err() ,"GrpcErrorToAppError"))
		return  app_error.NewInternalErrorNil(resErr.Err())
	}
}


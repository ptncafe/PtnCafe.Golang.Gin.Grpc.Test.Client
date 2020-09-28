package provider

import (
	"github.com/pkg/errors"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/infrastructure/app_logger"
	store "github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/proto"
	"google.golang.org/grpc"
)

var grpcClientConn *grpc.ClientConn

func NewShopServicesClient(grpcUrl string) (store.ShopServiceClient, error){
	var err error
	if grpcClientConn == nil {
		grpcClientConn, err = grpc.Dial(grpcUrl, grpc.WithInsecure())
		if err != nil {
			log:= app_logger.Logger
			log.Errorf("NewShopServicesClient %+v", errors.Wrap(err, "NewShopServicesClient"))
			return nil, err
		}
	}

	//defer conn.Close()

	client := store.NewShopServiceClient(grpcClientConn)
	return client , nil
}
package get_store_by_id_module

import (
    "context"
    "github.com/sirupsen/logrus"
    "github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/model/app_error"
    "github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/model/entity"
    store "github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/proto"
    "github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/provider"
)
type module struct{
    log *logrus.Logger
    shopServiceClient store.ShopServiceClient
}

func NewGetStoreByIdMudule(
    log *logrus.Logger,
    shopServiceClient store.ShopServiceClient,
    ) *module{
    return &module{
        log:log,
        shopServiceClient: shopServiceClient,
    }
}


func (m *module) GetStoreById(ctx context.Context, id int32) (*entity.Store, *app_error.AppError){
    data, err := m.shopServiceClient.GetStoreById(ctx, &store.GetStoreByIdReq{
        Id: id,
        IsCache: true,
    })

    if err != nil {
        return nil, provider.GrpcErrorToAppError(err, m.log)
    }
    return &entity.Store{
       Id : id,
       Name: data.Name,
       Code: data.Code,
       Phone: data.Phone,
       ShopStatus: data.ShopStatus,
       ShopType: data.ShopType,
    }, nil
}

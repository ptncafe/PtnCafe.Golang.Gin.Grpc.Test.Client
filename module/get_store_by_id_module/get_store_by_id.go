package get_store_by_id_module

import (
    "github.com/pkg/errors"
    "github.com/sirupsen/logrus"
	"context"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/model/entity"
    store "github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/proto"
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
func (m *module) GetStoreById(ctx context.Context, id int32) (*entity.Store, error){
    data, err := m.shopServiceClient.GetStoreById(ctx, &store.GetStoreByIdReq{
        Id: id,
        IsCache: true,
    })
    if err != nil {
        m.log.Errorf("GetStoreById %d %+v", id, errors.Wrap(err, "m.shopServiceClient.GetStoreById"))
        return nil, err
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

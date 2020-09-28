package get_store_by_id

import (
	"context"
	"github/ptncafe/PtnCafe.Golang.Gin.Grpc.Test/model/entity"
)

type IGetStoreBydIdModule interface{
	GetStoreById(ctx context.Context, id int32, isCache bool ) (*entity.Store, error)
}

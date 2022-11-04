package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGoodsService)

type GoodsService struct {
	v1.UnimplementedGoodsServer

	cac *biz.CategoryUsecase
	log *log.Helper
}

// NewGoodsService new a greeter service.
func NewGoodsService(cac *biz.CategoryUsecase, logger log.Logger) *GoodsService {
	return &GoodsService{cac: cac, log: log.NewHelper(logger)}
}

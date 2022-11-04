package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type GoodsTypeRepo interface {
	CreateGoodsType(context.Context, *domain.GoodsType) (int64, error)
}

type GoodsTypeUsecase struct {
	repo  GoodsTypeRepo
	bRepo BrandRepo
	tx    Transaction
	log   *log.Helper
}

func NewGoodsTypeUsecase(repo GoodsTypeRepo, tx Transaction, bRepo BrandRepo, logger log.Logger) *GoodsTypeUsecase {
	return &GoodsTypeUsecase{
		repo:  repo,
		tx:    tx,
		bRepo: bRepo,
		log:   log.NewHelper(logger),
	}
}

// GoosTypeCreate 创建商品类型
func (gt *GoodsTypeUsecase) GoosTypeCreate(ctx context.Context, r *domain.GoodsType) (int64, error) {
	id, err := gt.repo.CreateGoodsType(ctx, r)
	if err != nil {
		return id, err
	}
	return id, nil
}

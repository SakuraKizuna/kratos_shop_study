// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/conf"
	"goods/internal/data"
	"goods/internal/server"
	"goods/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	brandRepo := data.NewBrandRepo(dataData, logger)
	brandUsecase := biz.NewBrandUsecase(brandRepo, logger)
	categoryRepo := data.NewCategoryRepo(dataData, logger)
	categoryUsecase := biz.NewCategoryUsecase(categoryRepo, logger)
	goodsTypeRepo := data.NewGoodsTypeRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	goodsTypeUsecase := biz.NewGoodsTypeUsecase(goodsTypeRepo, transaction, brandRepo, logger)
	specificationRepo := data.NewSpecificationRepo(dataData, logger)
	specificationUsecase := biz.NewSpecificationUsecase(specificationRepo, goodsTypeRepo, transaction, logger)
	goodsAttrRepo := data.NewGoodsAttrRepo(dataData, logger)
	goodsAttrUsecase := biz.NewGoodsAttrUsecase(goodsAttrRepo, transaction, goodsTypeRepo, logger)
	goodsRepo := data.NewGoodsRepo(dataData, logger)
	goodsSkuRepo := data.NewGoodsSkuRepoRepo(dataData, logger)
	inventoryRepo := data.NewInventoryRepo(dataData, logger)
	esGoodsRepo := data.NewEsGoodsRepo(dataData, logger)
	goodsUsecase := biz.NewGoodsUsecase(goodsRepo, goodsSkuRepo, transaction, goodsTypeRepo, categoryRepo, brandRepo, specificationRepo, goodsAttrRepo, inventoryRepo, esGoodsRepo, logger)
	esGoodsUsecase := biz.NewEsGoodsUsecase(goodsRepo, esGoodsRepo, categoryRepo, logger)
	goodsService := service.NewGoodsService(brandUsecase, categoryUsecase, goodsTypeUsecase, specificationUsecase, goodsAttrUsecase, goodsUsecase, esGoodsUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, goodsService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}

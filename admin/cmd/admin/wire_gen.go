// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"admin/internal/biz"
	"admin/internal/conf"
	"admin/internal/data"
	"admin/internal/server"
	"admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init admin application.
func initApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, confService *conf.Service, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(auth, confService, discovery)
	dataData, err := data.NewData(confData, userClient, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger, auth)
	addressRepo := data.NewAddressRepo(dataData, logger)
	addressUsecase := biz.NewAddressUsecase(userRepo, addressRepo, logger, auth)
	adminService := service.NewAdminService(userUsecase, addressUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, adminService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, httpServer, registrar)
	return app, func() {
	}, nil
}

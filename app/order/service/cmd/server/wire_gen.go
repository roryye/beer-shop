// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/beer-shop/app/order/service/internal/biz"
	"github.com/go-kratos/beer-shop/app/order/service/internal/conf"
	"github.com/go-kratos/beer-shop/app/order/service/internal/data"
	"github.com/go-kratos/beer-shop/app/order/service/internal/server"
	"github.com/go-kratos/beer-shop/app/order/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	orderRepo := data.NewOrderRepo(dataData, logger)
	orderUseCase := biz.NewOrderUseCase(orderRepo, logger)
	orderService := service.NewOrderService(orderUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, orderService)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}

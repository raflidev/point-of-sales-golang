//go:build wireinject
// +build wireinject

package main

import (
	"golang-point-of-sales-system/app"
	"golang-point-of-sales-system/controller"
	"golang-point-of-sales-system/middleware"
	"golang-point-of-sales-system/modules/products/domain/repository"
	"golang-point-of-sales-system/modules/products/domain/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var productSet = wire.NewSet(
	repository.NewProductRepository,
	service.NewProductService,
	controller.NewProductController,
)

func InitServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		productSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}

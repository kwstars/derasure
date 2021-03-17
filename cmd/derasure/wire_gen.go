// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/derasure"
	"github.com/kwstars/derasure/internal/app/derasure/controllers"
	"github.com/kwstars/derasure/internal/app/derasure/repositories"
	"github.com/kwstars/derasure/internal/app/derasure/services/erasure"
	"github.com/kwstars/derasure/pkg/config"
	"github.com/kwstars/derasure/pkg/database"
	"github.com/kwstars/derasure/pkg/log"
	"github.com/kwstars/derasure/pkg/transports/http"
)

// Injectors from wire.go:

func CreateApp(confPath string) (*derasure.App, func(), error) {
	viper, err := config.New(confPath)
	if err != nil {
		return nil, nil, err
	}
	logOptions, err := log.NewLogOptions(viper)
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup, err := log.New(logOptions)
	if err != nil {
		return nil, nil, err
	}
	redisOptions, err := database.NewRedisOptions(viper)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	client, cleanup2, err := database.New(redisOptions)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	repostiory := &repositories.Repostiory{
		Redis: client,
	}
	eliminate := &erasure.Eliminate{
		Repostiory: repostiory,
	}
	banquet := &erasure.Banquet{
		Repostiory: repostiory,
	}
	fishing := &erasure.Fishing{
		Repostiory: repostiory,
	}
	kite := &erasure.Kite{
		Repostiory: repostiory,
	}
	limitedGift := &erasure.LimitedGift{
		Repostiory: repostiory,
	}
	erasureServices := erasure.ErasureServices{
		Logger:      logger,
		Eliminate:   eliminate,
		Banquet:     banquet,
		Fishing:     fishing,
		Kite:        kite,
		LimitedGift: limitedGift,
	}
	controller := &controllers.Controller{
		Service: erasureServices,
	}
	initControllers := controllers.CreateInitControllersFn(controller)
	engine := http.NewRouter(initControllers, logger)
	server, err := http.New(engine)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, err := derasure.NewApp(server, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, database.ProviderSet, repositories.ProviderSet, erasure.ProviderSet, controllers.ProviderSet, http.ProviderSet, derasure.ProviderSet)

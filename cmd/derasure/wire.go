// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app"
	"github.com/kwstars/derasure/internal/app/controllers"
	"github.com/kwstars/derasure/internal/app/model"
	"github.com/kwstars/derasure/internal/app/services"
	globalconfig "github.com/kwstars/derasure/pkg/config"
	"github.com/kwstars/derasure/pkg/db"
	"github.com/kwstars/derasure/pkg/log"
	"github.com/kwstars/derasure/pkg/transports/http"
)

var providerSet = wire.NewSet(
	globalconfig.GlobalProviderSet,
	log.ProviderSet,
	db.ProviderSet,
	model.ProviderSet,
	services.ProviderSet,
	controllers.ProviderSet,
	http.ProviderSet,
	app.ProviderSet,
)

func CreateApp(confPath string) (*app.App, func(), error) {
	panic(wire.Build(providerSet))
}

// +build wireinject

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

var providerSet = wire.NewSet(
	config.ProviderSet,
	log.ProviderSet,
	database.ProviderSet,
	repositories.ProviderSet,
	erasure.ProviderSet,
	controllers.ProviderSet,
	http.ProviderSet,
	derasure.ProviderSet,
)

func CreateApp(confPath string) (*derasure.App, func(), error) {
	panic(wire.Build(providerSet))
}

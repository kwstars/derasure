// +build wireinject

package repositories

import (
	"github.com/google/wire"
	"github.com/kwstars/derasure/pkg/config"
	"github.com/kwstars/derasure/pkg/database"
)

var ProviderSet = wire.NewSet(
	ErasureSet,
)

var testProviderSet = wire.NewSet(
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateErasureRepository(f string) (IErasureRepostiory, func(), error) {
	panic(wire.Build(testProviderSet))
}

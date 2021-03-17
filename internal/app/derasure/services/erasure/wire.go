package erasure

import (
	"github.com/google/wire"
	"github.com/kwstars/derasure/pkg/config"
	"github.com/kwstars/derasure/pkg/database"
	"github.com/kwstars/derasure/pkg/log"
)

var ProviderSet = wire.NewSet(
	DelServiceSet,
	EliminateSet,
	BanquetSet,
	FishingSet,
	KiteSet,
	LimitedGiftSet,
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

//func CreateDerasureService(cf string, sto repositories.ReviewsRepository) (ReviewsService, error) {
//	panic(wire.Build(testProviderSet))
//}

package erasure

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	DelServiceSet,
	EliminateSet,
	BanquetSet,
	FishingSet,
	KiteSet,
	LimitedGiftSet,
)

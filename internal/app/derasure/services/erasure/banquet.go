package erasure

import (
	"context"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/derasure/repositories"
)

var BanquetSet = wire.NewSet(wire.Struct(new(Banquet), "*"))

var _ IDelData = (*Banquet)(nil)

type Banquet struct {
	Repostiory *repositories.ErasureRepostiory
}

func (b *Banquet) Execution(ctx context.Context, uid string) (err error) {
	if err := b.Repostiory.DelKey(ctx, "banquet:"+uid); err != nil {
		return err
	}
	return
}

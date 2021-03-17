package erasure

import (
	"context"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/derasure/repositories"
)

var EliminateSet = wire.NewSet(wire.Struct(new(Eliminate), "*"))

var _ IDelData = (*Eliminate)(nil)

type Eliminate struct {
	Repostiory *repositories.ErasureRepostiory
}

func (b *Eliminate) Execution(ctx context.Context, uid string) (err error) {
	if err := b.Repostiory.DelKey(ctx, "banquet:"+uid); err != nil {
		return err
	}
	return
}

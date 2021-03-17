package erasure

import (
	"context"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/derasure/repositories"
)

var LimitedGiftSet = wire.NewSet(wire.Struct(new(LimitedGift), "*"))

var _ IDelData = (*LimitedGift)(nil)

type LimitedGift struct {
	Repostiory *repositories.ErasureRepostiory
}

func (b *LimitedGift) Execution(ctx context.Context, uid string) (err error) {
	if err := b.Repostiory.DelKey(ctx, "lgift:"+uid); err != nil {
		return err
	}
	return
}

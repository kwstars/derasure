package services

import (
	"context"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/model"
)

var LimitedGiftSet = wire.NewSet(wire.Struct(new(LimitedGift), "*"))

var _ IDel = (*LimitedGift)(nil)

type LimitedGift struct {
	Repostiory *model.Repostiory
}

func (b *LimitedGift) Execution(ctx context.Context, uid string) (err error) {
	if err := b.Repostiory.DelKey(ctx, "lgift:"+uid); err != nil {
		return err
	}
	return
}

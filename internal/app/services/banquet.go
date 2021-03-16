package services

import (
	"context"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/model"
)

var BanquetSet = wire.NewSet(wire.Struct(new(Banquet), "*"))

var _ IDel = (*Banquet)(nil)

type Banquet struct {
	Repostiory *model.Repostiory
}

func (b *Banquet) Execution(ctx context.Context, uid string) (err error) {
	if err := b.Repostiory.DelKey(ctx, "banquet:"+uid); err != nil {
		return err
	}
	return
}

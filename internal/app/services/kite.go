package services

import (
	"context"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/model"
)

var KiteSet = wire.NewSet(wire.Struct(new(Kite), "*"))

var _ IDel = (*Kite)(nil)

type Kite struct {
	Repostiory *model.Repostiory
}

func (b *Kite) Execution(ctx context.Context, uid string) (err error) {
	if err := b.Repostiory.DelKey(ctx, "kite:"+uid); err != nil {
		return err
	}
	return
}

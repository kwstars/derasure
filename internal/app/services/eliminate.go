package services

import (
	"context"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/model"
)

var EliminateSet = wire.NewSet(wire.Struct(new(Eliminate), "*"))

var _ IDel = (*Eliminate)(nil)

type Eliminate struct {
	Repostiory *model.Repostiory
}

func (b *Eliminate) Execution(ctx context.Context, uid string) (err error) {
	if err := b.Repostiory.DelKey(ctx, "banquet:"+uid); err != nil {
		return err
	}
	return
}

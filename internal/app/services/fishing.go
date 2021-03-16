package services

import (
	"context"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/model"
)

var FishingSet = wire.NewSet(wire.Struct(new(Fishing), "*"))

var _ IDel = (*Fishing)(nil)

type Fishing struct {
	Repostiory *model.Repostiory
}

func (b *Fishing) Execution(ctx context.Context, uid string) (err error) {
	if err := b.Repostiory.DelKey(ctx, "fishing:"+uid); err != nil {
		return err
	}
	return
}

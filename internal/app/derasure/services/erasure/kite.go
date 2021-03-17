package erasure

import (
	"context"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/derasure/repositories"
)

var KiteSet = wire.NewSet(wire.Struct(new(Kite), "*"))

var _ IDelData = (*Kite)(nil)

type Kite struct {
	Repostiory *repositories.Repostiory
}

func (b *Kite) Execution(ctx context.Context, uid string) (err error) {
	if err := b.Repostiory.DelKey(ctx, "kite:"+uid); err != nil {
		return err
	}
	return
}

package repositories

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

var ErasureSet = wire.NewSet(
	wire.NewSet(wire.Struct(new(ErasureRepostiory), "*")),
	wire.Bind(new(IErasureRepostiory), new(*ErasureRepostiory)),
)

type IErasureRepostiory interface {
	DelKey(ctx context.Context, key string) (err error)
}

type ErasureRepostiory struct {
	Redis *redis.Client
}

func (db *ErasureRepostiory) DelKey(ctx context.Context, key string) (err error) {
	if _, err = db.Redis.Del(ctx, key).Result(); err != nil {
		return errors.WithStack(err)
	}
	return
}

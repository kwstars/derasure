package repositories

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

var Set = wire.NewSet(wire.Struct(new(Repostiory), "*"))

type Repostiory struct {
	Redis *redis.Client
}

func (db *Repostiory) DelKey(ctx context.Context, key string) (err error) {
	if _, err = db.Redis.Del(ctx, key).Result(); err != nil {
		return errors.WithStack(err)
	}
	return
}

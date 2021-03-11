package model

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

func (db *Repostiory) Del(ctx context.Context, s string) (err error) {
	var cursor uint64
	var n int
	var totalKeys []string
	for {
		var keys []string
		var err error
		keys, cursor, err = db.Redis.Scan(ctx, cursor, s+"*", 5000).Result()
		if err != nil {
			return errors.WithStack(err)
		}
		n += len(keys)
		totalKeys = append(totalKeys, keys...)
		if cursor == 0 {
			break
		}
	}

	if len(totalKeys) > 0 {
		if _, err := db.Redis.Del(ctx, totalKeys...).Result(); err != nil {
			return errors.WithStack(err)
		}
	}

	return
}

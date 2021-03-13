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

func (db *Repostiory) DelKey(ctx context.Context, key string) (err error) {
	if _, err = db.Redis.Del(ctx, key).Result(); err != nil {
		return err
	}
	return
}

func (db *Repostiory) CheckAccountExist(ctx context.Context, uid string) (err error) {
	key := "user:" + uid
	exist, err := db.Redis.Exists(ctx, key).Result()
	if err != nil {
		return err
	}

	if exist == 0 {
		return errors.New("用户不存在")
	}

	return
}

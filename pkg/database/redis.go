package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
)

var ProviderSet = wire.NewSet(New, NewRedisOptions)

type RedisOptions struct {
	Address  string `yaml:"address"`
	DB       int    `yaml:"db"`
	Passwrod string `yaml:"password"`
}

func NewRedisOptions(v *viper.Viper) (o *RedisOptions, err error) {
	o = &RedisOptions{}
	if err = v.UnmarshalKey("redis", o); err != nil {
		return nil, errors.WithStack(err)
	}

	if v, ok := viper.Get("REDIS_ADDR").(string); ok {
		o.Address = v
	}

	if v, ok := viper.Get("REDIS_DB").(int); ok {
		o.DB = v
	}

	if v, ok := viper.Get("REDIS_PASS").(string); ok {
		o.Passwrod = v
	}

	return o, err
}

func New(o *RedisOptions) (db *redis.Client, cf func(), err error) {
	db = redis.NewClient(&redis.Options{
		Addr:     o.Address,
		Password: o.Passwrod, // no password set
		DB:       o.DB,       // use default DB
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			if _, err := cn.Ping(ctx).Result(); err != nil {
				return err
			}
			return nil
		},
	})

	if _, err := db.Ping(context.Background()).Result(); err != nil {
		return nil, nil, err
	}

	cf = func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}

	return
}

package redis

import (
	"context"

	"github.com/artemiyKew/json-rpc-lamoda/config"
	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	db *redis.Client
}

func NewDB(cfg config.Config) *RedisDB {
	return &RedisDB{
		db: redis.NewClient(&redis.Options{
			Addr:     cfg.RedisAddr,
			Password: cfg.RedisPass,
			DB:       cfg.RedisType,
		}),
	}
}

func (r *RedisDB) Ping(ctx context.Context) error {
	return r.db.Ping(ctx).Err()
}

func (r *RedisDB) Close(ctx context.Context) error {
	if err := r.Ping(ctx); err != nil {
		return err
	}
	return r.db.Close()
}

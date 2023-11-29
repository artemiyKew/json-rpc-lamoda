package redis

import (
	"context"
	"fmt"
)

type ProductRepo struct {
	*RedisDB
}

func NewProductRepo(db *RedisDB) *ProductRepo {
	return &ProductRepo{db}
}

func (r *ProductRepo) ReserveProduct(ctx context.Context, uniqueCode string, warehouseID, countProduct int) error {
	key := fmt.Sprintf("reservation:%s", uniqueCode)
	if err := r.db.Set(ctx, key, int64(countProduct), 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *ProductRepo) CancelReservationProduct(ctx context.Context, uniqueCode string, countProduct int) error {
	key := fmt.Sprintf("reservation:%s", uniqueCode)
	newValue, err := r.db.DecrBy(ctx, key, int64(countProduct)).Result()
	if err != nil {
		return err
	}

	if newValue == 0 {
		if err := r.db.Del(ctx, key).Err(); err != nil {
			return err
		}
	}

	return nil
}

func (r *ProductRepo) IsProductReserved(ctx context.Context, uniqueCode string) (bool, error) {
	key := fmt.Sprintf("reservation:%s", uniqueCode)
	exists, err := r.db.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

package encryption

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisKeyCache implements the KeyCache interface using Redis.
type RedisKeyCache struct {
	redis  *redis.Client
	prefix string
}

// NewRedisKeyCache creates a new RedisKeyCache instance.
func NewRedisKeyCache(r *redis.Client, prefix string) *RedisKeyCache {
	return &RedisKeyCache{redis: r, prefix: prefix}
}

// GetKey retrieves the cached key from Redis.
// It returns the key and a bool indicating whether it was found.
func (r *RedisKeyCache) GetKey() (*CachedKey, bool, error) {
	ctx := context.Background()
	fmt.Println("get key called")
	val, err := r.redis.Get(ctx, r.prefix).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, false, nil
		}
		return nil, false, err
	}

	var k CachedKey
	if err := json.Unmarshal([]byte(val), &k); err != nil {
		return nil, false, err
	}

	return &k, true, nil
}

// SetKey caches the provided key in Redis with its TTL (expiry).
// It takes a Key and stores it with the corresponding expiration time.
func (r *RedisKeyCache) SetKey(k *Key) error {
	ctx := context.Background()
	b, err := json.Marshal(k)
	if err != nil {
		return err
	}

	ttl := time.Until(time.Unix(k.ExpiresAt, 0))
	return r.redis.Set(ctx, r.prefix, b, ttl).Err()
}

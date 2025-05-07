package connector

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func ConnectToRedis(ctx context.Context, cfg RedisConfig) (*redis.Client, error) {
	logrus.Print("Connecting to Redis")

	var rdb *redis.Client
	var err error

	for {
		// Attempt to create Redis client
		rdb = redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Password,
			DB:       cfg.DB,
		})

		_, err = rdb.Ping(ctx).Result()
		if err == nil {
			break // Successful connection, exit loop
		}

		// Check if it's a connection error that can be retried
		if !IsRetriableError(err) {
			return nil, errors.Wrap(err, "unable to connect to Redis")
		}

		// Log and retry after a delay
		logrus.Printf("Retrying connecting to Redis: %v\n", err)
		const retryDelay = 5 * time.Second
		time.Sleep(retryDelay)
	}

	logrus.Print("Connected to Redis")
	return rdb, nil
}

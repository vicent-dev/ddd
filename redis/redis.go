package redis

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisConnection struct {
	pool *redis.Pool
}

func NewRedisConnection(host, port string, maxIdle int) *RedisConnection {
	address := host + ":" + port
	return &RedisConnection{&redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", address) },
	}}
}

func (rc RedisConnection) GetConnection(ctx context.Context) (redis.Conn, error) {
	return rc.pool.GetContext(ctx)
}

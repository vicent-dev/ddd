package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisConnection struct {
	pool *redis.Pool
}

func NewRedisConnection() *RedisConnection {
	address := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	maxIdle, err := strconv.Atoi(os.Getenv("REDIS_MAX_IDLE"))
	if err != nil {
		log.Fatal(err)
	}
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

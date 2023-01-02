package redis

import (
	"github.com/go-redis/redis/v9"
	"time"
)

type Config struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

const (
	maxRetries         = 5
	minRetryBackoff    = 300 * time.Millisecond
	maxRetryBackoff    = 500 * time.Millisecond
	dialTimeout        = 5 * time.Second
	readTimeout        = 5 * time.Second
	writeTimeout       = 3 * time.Second
	minIdleConnections = 20
	poolTimeout        = 6 * time.Second
	idleTimeout        = 12 * time.Second
)

func NewUniversalRedisClient(cfg *Config) redis.UniversalClient {

	universalClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:           []string{cfg.Addr},
		Password:        cfg.Password, // no password set
		DB:              cfg.DB,       // use default DB
		MaxRetries:      maxRetries,
		MinRetryBackoff: minRetryBackoff,
		MaxRetryBackoff: maxRetryBackoff,
		DialTimeout:     dialTimeout,
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
		PoolSize:        cfg.PoolSize,
		MinIdleConns:    minIdleConnections,
		PoolTimeout:     poolTimeout,
		ConnMaxIdleTime: idleTimeout,
	})

	return universalClient
}

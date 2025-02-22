package option

import (
	"strings"
	"time"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

// CacheOption holds the options for cache
type CacheOption struct {
	CacheBackend string
	CacheTTL     time.Duration
	RedisOption
}

// RedisOption holds the options for redis cache
type RedisOption struct {
	RedisCACert string
	RedisCert   string
	RedisKey    string
}

// NewCacheOption returns an instance of CacheOption
func NewCacheOption(c *cli.Context) CacheOption {
	return CacheOption{
		CacheBackend: c.String("cache-backend"),
		CacheTTL:     c.Duration("cache-ttl"),
		RedisOption: RedisOption{
			RedisCACert: c.String("redis-ca"),
			RedisCert:   c.String("redis-cert"),
			RedisKey:    c.String("redis-key"),
		},
	}
}

// Init initialize the CacheOption
func (c *CacheOption) Init() error {
	// "redis://" or "fs" are allowed for now
	// An empty value is also allowed for testability
	if !strings.HasPrefix(c.CacheBackend, "redis://") &&
		c.CacheBackend != "fs" && c.CacheBackend != "" {
		return xerrors.Errorf("unsupported cache backend: %s", c.CacheBackend)
	}
	// if one of redis option not nil, make sure CA, cert, and key provided
	if (RedisOption{}) != c.RedisOption {
		if c.RedisCACert == "" || c.RedisCert == "" || c.RedisKey == "" {
			return xerrors.Errorf("you must provide CA, cert and key file path when using tls")
		}
	}
	return nil
}

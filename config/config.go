package config

import (
	"net/url"
	"time"
)

type Config struct {
	timeout time.Duration
	url     *url.URL
}

func (c *Config) Timeout() time.Duration {
	return c.timeout
}

func (c *Config) URL() *url.URL {
	return c.url
}

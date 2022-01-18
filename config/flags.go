package config

import (
	"flag"
	"log"
	"net/url"
	"time"
)

var timeout *time.Duration

var parsedURL *url.URL

func init() {
	timeout = flag.Duration("timeout", time.Second*18000, "The HTTP client timeout")
	flag.Parse()

	rawURL := flag.Arg(0)
	var err error
	parsedURL, err = url.ParseRequestURI(rawURL)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func NewConfigFromFlags() *Config {
	return &Config{
		timeout: *timeout,
		url:     parsedURL,
	}
}

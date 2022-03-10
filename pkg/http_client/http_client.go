package http_client

import (
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	dialContextTimeout        = 5 * time.Second
	clientTLSHandshakeTimeout = 5 * time.Second
	clientRetryWaitTime       = 300 * time.Millisecond

	defaultRetryCount    = 3
	defaultClientTimeout = 5 * time.Second
)

type HttpClientConfig struct {
	ClientTimeout time.Duration
	RetryCount    int
}

type Option func(client *resty.Client)

func ConfigOption(cfg HttpClientConfig) Option {
	return func(client *resty.Client) {
		if cfg.ClientTimeout != 0 {
			client.SetTimeout(cfg.ClientTimeout)
		}
		if cfg.RetryCount != 0 {
			client.SetRetryCount(0)
		}
	}
}

func NewHttpClient(debugMode bool, options ...Option) *resty.Client {
	t := &http.Transport{
		DialContext:         (&net.Dialer{Timeout: dialContextTimeout}).DialContext,
		TLSHandshakeTimeout: clientTLSHandshakeTimeout,
	}

	client := resty.New().
		SetDebug(debugMode).
		SetTimeout(defaultClientTimeout).
		SetRetryCount(defaultRetryCount).
		SetRetryWaitTime(clientRetryWaitTime).
		SetTransport(t)
	for _, option := range options {
		option(client)
	}

	return client
}

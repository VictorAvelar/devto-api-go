package devto

import (
	"context"
	"errors"
	"net/http"
)

//Configuration constants
const (
	BaseURL      string = "https://dev.to/api"
	APIVersion   string = "0.5.1"
	APIKeyHeader string = "api-key"
)

//devto client errors
var (
	ErrMissingConfig = errors.New("missing configuration")
)

type httpClient interface {
	Do(req *http.Request) (res *http.Response, err error)
}

//Client is the main data structure for performing actions
//against dev.to API
type Client struct {
	Context    context.Context
	HTTPClient httpClient
	Config     *Config
}

//NewClient takes a context, a configuration pointer and optionally a
//base http client (bc) to build an Client instance.
func NewClient(ctx context.Context, conf *Config, bc httpClient) (dev *Client, err error) {
	if bc == nil {
		bc = http.DefaultClient
	}

	if ctx == nil {
		ctx = context.Background()
	}

	if conf == nil {
		return nil, ErrMissingConfig
	}

	return &Client{
		Context:    ctx,
		HTTPClient: bc,
		Config:     conf,
	}, nil
}

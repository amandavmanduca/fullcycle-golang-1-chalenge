package interfaces

import (
	"context"
	"net/http"
)

type HttpClientInterface interface {
	Get(ctx context.Context, path string) (*http.Response, error)
}

type AwesomeApiClientInterface interface {
	GetExchangeRate(ctx context.Context) (*float64, error)
}

package interfaces

import (
	"context"
	"net/http"
)

type ClientsContainer struct {
	ServerApiClient ServerApiClientInterface
}

type HttpClientInterface interface {
	Get(ctx context.Context, path string) (*http.Response, error)
}

type ServerApiClientInterface interface {
	GetExchangeRate(ctx context.Context) (*float64, error)
}

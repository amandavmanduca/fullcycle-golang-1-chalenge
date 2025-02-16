package interfaces

import "context"

type HttpClientInterface interface {
	Get(ctx context.Context, path string) ([]byte, error)
}

type AwesomeApiClientInterface interface {
	GetExchangeRate(ctx context.Context) (*float64, error)
}

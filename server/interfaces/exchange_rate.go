package interfaces

import (
	"context"
	"server/structs"
)

type ExchangeRateRepositoryInterface interface {
	Create(ctx context.Context, data structs.ExchangeRate) error
}

type ExchangeRateServiceInterface interface {
	Get(ctx context.Context) (structs.ExchangeRate, error)
}

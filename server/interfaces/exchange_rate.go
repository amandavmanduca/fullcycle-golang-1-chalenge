package interfaces

import (
	"context"

	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/structs"
)

type ExchangeRateRepositoryInterface interface {
	Create(ctx context.Context, data structs.ExchangeRate) error
}

type ExchangeRateServiceInterface interface {
	Get(ctx context.Context) (structs.ExchangeRate, error)
}

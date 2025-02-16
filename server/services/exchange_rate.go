package services

import (
	"context"
	"server/interfaces"
	"server/structs"

	"github.com/google/uuid"
)

type ExchangeRateService struct {
	exchangeRateRepository interfaces.ExchangeRateRepositoryInterface
	awesomeClientApi       interfaces.AwesomeApiClientInterface
}

func NewExchangeRateService(
	exchangeRateRepository interfaces.ExchangeRateRepositoryInterface,
	awesomeClientApi interfaces.AwesomeApiClientInterface,
) interfaces.ExchangeRateServiceInterface {
	return &ExchangeRateService{
		exchangeRateRepository: exchangeRateRepository,
		awesomeClientApi:       awesomeClientApi,
	}
}

func (s *ExchangeRateService) Get(ctx context.Context) (structs.ExchangeRate, error) {
	value, err := s.awesomeClientApi.GetExchangeRate(ctx)
	if err != nil || value == nil {
		return structs.ExchangeRate{}, err
	}
	res := structs.ExchangeRate{
		ID:  uuid.New().String(),
		Bid: *value,
	}
	_, err = s.create(ctx, res)
	if err != nil {
		return structs.ExchangeRate{}, err
	}
	return res, nil
}

func (s *ExchangeRateService) create(ctx context.Context, data structs.ExchangeRate) (structs.ExchangeRate, error) {
	err := s.exchangeRateRepository.Create(ctx, data)
	if err != nil {
		return structs.ExchangeRate{}, err
	}
	return data, err
}

package services

import (
	"context"

	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/interfaces"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/structs"

	"github.com/google/uuid"
)

type exchangeRateService struct {
	exchangeRateRepository interfaces.ExchangeRateRepositoryInterface
	awesomeClientApi       interfaces.AwesomeApiClientInterface
}

func NewExchangeRateService(
	exchangeRateRepository interfaces.ExchangeRateRepositoryInterface,
	awesomeClientApi interfaces.AwesomeApiClientInterface,
) interfaces.ExchangeRateServiceInterface {
	return &exchangeRateService{
		exchangeRateRepository: exchangeRateRepository,
		awesomeClientApi:       awesomeClientApi,
	}
}

func (s *exchangeRateService) Get(ctx context.Context) (structs.ExchangeRate, error) {
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

func (s *exchangeRateService) create(ctx context.Context, data structs.ExchangeRate) (structs.ExchangeRate, error) {
	err := s.exchangeRateRepository.Create(ctx, data)
	if err != nil {
		return structs.ExchangeRate{}, err
	}
	return data, err
}

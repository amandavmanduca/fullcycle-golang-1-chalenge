package services

import (
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/interfaces"
)

func NewServicesContainer(repos interfaces.RepositoriesContainer, clients interfaces.ClientsContainer) *interfaces.ServicesContainer {
	return &interfaces.ServicesContainer{
		ExchangeRateService: NewExchangeRateService(repos.ExchangeRepository, clients.AwesomeApiClient),
	}
}

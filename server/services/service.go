package services

import (
	"server/interfaces"
)

func NewServicesContainer(repos interfaces.RepositoriesContainer, clients interfaces.ClientsContainer) interfaces.ServicesContainer {
	return interfaces.ServicesContainer{
		ExchangeRateService: NewExchangeRateService(repos.ExchangeRepository, clients.AwesomeApiClient),
	}
}

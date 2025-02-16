package interfaces

type ServicesContainer struct {
	ExchangeRateService ExchangeRateServiceInterface
}

type RepositoriesContainer struct {
	ExchangeRepository ExchangeRateRepositoryInterface
}

type ClientsContainer struct {
	AwesomeApiClient AwesomeApiClientInterface
}

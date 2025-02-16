package repositories

import (
	"server/database"
	"server/interfaces"
)

func NewRepositoriesContainer(db database.Database) *interfaces.RepositoriesContainer {
	return &interfaces.RepositoriesContainer{
		ExchangeRepository: NewExchangeRateRepository(db),
	}
}

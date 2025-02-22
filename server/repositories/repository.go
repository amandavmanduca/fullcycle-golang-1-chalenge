package repositories

import (
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/database"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/interfaces"
)

func NewRepositoriesContainer(db database.Database) *interfaces.RepositoriesContainer {
	return &interfaces.RepositoriesContainer{
		ExchangeRepository: NewExchangeRateRepository(db),
	}
}

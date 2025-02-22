package repositories

import (
	"context"
	"time"

	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/database"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/interfaces"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/structs"
)

type exchangeRateRepository struct {
	db            database.Database
	createTimeout *time.Duration
}

func NewExchangeRateRepository(db database.Database) interfaces.ExchangeRateRepositoryInterface {
	timeout := time.Millisecond * 10
	return &exchangeRateRepository{
		db:            db,
		createTimeout: &timeout,
	}
}

func (r *exchangeRateRepository) Create(ctx context.Context, data structs.ExchangeRate) error {
	if r.createTimeout != nil {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, *r.createTimeout)
		defer cancel()
		ctx = ctxWithTimeout
	}
	conn, err := r.db.GetConnection(ctx)
	if err != nil {
		return err
	}
	err = conn.Model(&structs.ExchangeRate{}).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

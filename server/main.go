package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/clients"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/database"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/handlers"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/repositories"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/services"

	"github.com/google/uuid"
)

type ExchangeRate struct {
	ID  string  `json:"id" gorm:"primaryKey"`
	Bid float64 `json:"bid"`
}

func NewExchangeRate(bid float64) *ExchangeRate {
	return &ExchangeRate{
		ID:  uuid.New().String(),
		Bid: bid,
	}
}

func main() {

	sqliteConfig := database.SqliteImpl{
		Dsn: "gorm.db",
	}
	ctx := context.Background()

	dbConn := database.NewSqlite(sqliteConfig)

	_, err := dbConn.Connect(ctx)
	if err != nil {
		fmt.Println("error connecting")
		panic(err)
	}
	err = dbConn.AutoMigrate(ctx, &ExchangeRate{})
	if err != nil {
		fmt.Println("error executing database migrations")
		panic(err)
	}

	clients := clients.NewClientsContainer()
	repos := repositories.NewRepositoriesContainer(dbConn)
	services := services.NewServicesContainer(*repos, clients)
	handlers := handlers.NewHandlersContainer(*services)

	http.HandleFunc("/cotacao", handlers.ExchangeRateHandler.Get)
	http.ListenAndServe(":8080", nil)

}

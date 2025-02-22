package main

import (
	"context"
	"fmt"

	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/client/clients"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/client/file_handler"
)

func main() {
	ctx := context.Background()

	txtFile := file_handler.NewFileHandler("cotacao.txt")
	handler := clients.NewClientsContainer()

	bid, err := handler.ServerApiClient.GetExchangeRate(ctx)
	if err != nil {
		fmt.Println("erro ao obter dolar rate", err.Error())
		return
	}

	err = txtFile.Write(fmt.Sprintf("Dólar: {%f}", *bid))
	if err != nil {
		fmt.Println("erro ao registrar valor", err.Error())
	}
}

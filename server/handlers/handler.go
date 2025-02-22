package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"server/interfaces"
)

type HandlersContainer struct {
	ExchangeRateHandler exchangeRateHandler
}

func NewHandlersContainer(services interfaces.ServicesContainer) HandlersContainer {
	return HandlersContainer{
		ExchangeRateHandler: NewExchangeRateHandler(services.ExchangeRateService),
	}
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HttpResponse(w http.ResponseWriter, data SuccessResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Erro ao gerar JSON", http.StatusInternalServerError)
	}
}

func HttpErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err == context.DeadlineExceeded {
		w.WriteHeader(http.StatusRequestTimeout)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Status:  "error",
		Message: err.Error(),
	}); err != nil {
		http.Error(w, "Erro ao gerar JSON", http.StatusInternalServerError)
	}
}

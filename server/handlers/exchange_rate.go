package handlers

import (
	"context"
	"net/http"
	"server/interfaces"
)

type ExchangeRateHandler struct {
	exchangeRateService interfaces.ExchangeRateServiceInterface
}

func NewExchangeRateHandler(exchangeRateService interfaces.ExchangeRateServiceInterface) ExchangeRateHandler {
	return ExchangeRateHandler{
		exchangeRateService: exchangeRateService,
	}
}

func (h *ExchangeRateHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	res, err := h.exchangeRateService.Get(ctx)
	if err != nil {
		HttpErrorResponse(w, err)
		return
	}

	HttpResponse(w, SuccessResponse{
		Status:  "success",
		Message: "Taxas de câmbio obtidas com sucesso.",
		Data:    res,
	})
}

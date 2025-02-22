package server_api

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/client/interfaces"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/client/structs"
)

type serverApiClient struct {
	httpClient interfaces.HttpClientInterface
}

func NewServerApiClient(httpClient interfaces.HttpClientInterface) interfaces.ServerApiClientInterface {
	return serverApiClient{
		httpClient: httpClient,
	}
}

func (c serverApiClient) GetExchangeRate(ctx context.Context) (*float64, error) {
	resp, err := c.httpClient.Get(ctx, "/cotacao")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response structs.ServerApiResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var exchangeRate structs.ExchangeRateResponse
		err = json.Unmarshal(response.Data, &exchangeRate)
		if err != nil {
			return nil, err
		}

		floatBid, err := exchangeRate.Bid.Float64()
		if err != nil {
			return nil, err
		}
		return &floatBid, nil
	}
	return nil, errors.New(response.Message)
}

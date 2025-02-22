package awesomeapi

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/interfaces"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/structs"
)

type AwesomeApiClient struct {
	httpClient interfaces.HttpClientInterface
}

func NewAwesomeApiClient(httpClient interfaces.HttpClientInterface) interfaces.AwesomeApiClientInterface {
	return AwesomeApiClient{
		httpClient: httpClient,
	}
}

func (c AwesomeApiClient) GetExchangeRate(ctx context.Context) (*float64, error) {
	resp, err := c.httpClient.Get(ctx, "/json/last/USD-BRL")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response structs.AwesomeApiExchangeResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	floatBid, err := strconv.ParseFloat(response.USDBRL.Bid, 64)
	if err != nil {
		return nil, err
	}
	return &floatBid, nil
}

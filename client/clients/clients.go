package clients

import (
	"time"

	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/client/clients/server_api"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/client/interfaces"
)

func NewClientsContainer() interfaces.ClientsContainer {
	apiHttpClientTimeout := time.Millisecond * 300
	apiHttpClient := NewHttpClient("http://localhost:8080", &apiHttpClientTimeout)

	return interfaces.ClientsContainer{
		ServerApiClient: server_api.NewServerApiClient(apiHttpClient),
	}
}

package clients

import (
	"time"

	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/clients/awesomeapi"
	"github.com/amandavmanduca/fullcycle-golang-1-chalenge/server/interfaces"
)

func NewClientsContainer() interfaces.ClientsContainer {
	awesomeApiHttpClientTimeout := time.Millisecond * 200
	awesomeApiHttpClient := NewHttpClient("https://economia.awesomeapi.com.br", &awesomeApiHttpClientTimeout)

	return interfaces.ClientsContainer{
		AwesomeApiClient: awesomeapi.NewAwesomeApiClient(awesomeApiHttpClient),
	}
}

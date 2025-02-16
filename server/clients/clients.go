package clients

import (
	"server/clients/awesomeapi"
	"server/interfaces"
	"time"
)

func NewClientsContainer() interfaces.ClientsContainer {
	awesomeApiHttpClientTimeout := time.Millisecond * 200
	awesomeApiHttpClient := NewHttpClient("https://economia.awesomeapi.com.br", &awesomeApiHttpClientTimeout)

	return interfaces.ClientsContainer{
		AwesomeApiClient: awesomeapi.NewAwesomeApiClient(awesomeApiHttpClient),
	}
}

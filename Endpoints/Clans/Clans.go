package Clans

import (
	"QuaverGo"
)

type Clans struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *Clans {
	return &Clans{APIClient: apiClient, EndpointExtension: "/clan/"}
}

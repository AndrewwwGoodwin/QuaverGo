package Clans

import "quaverGo"

type Clans struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *Clans {
	return &Clans{APIClient: apiClient, EndpointExtension: "/clan/"}
}

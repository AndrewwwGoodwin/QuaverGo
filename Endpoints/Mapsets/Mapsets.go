package Mapsets

import "quaverGo"

type Mapsets struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *Mapsets {
	return &Mapsets{APIClient: apiClient, EndpointExtension: "/mapset/"}
}

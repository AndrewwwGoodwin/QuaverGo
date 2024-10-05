package Mapsets

import (
	"QuaverGo"
)

type Mapsets struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *Mapsets {
	return &Mapsets{APIClient: apiClient, EndpointExtension: "/mapset/"}
}

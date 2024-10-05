package Multiplayer

import (
	"QuaverGo"
)

type Multiplayer struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *Multiplayer {
	return &Multiplayer{APIClient: apiClient, EndpointExtension: "/multiplayer/"}
}

package Multiplayer

import "quaverGo"

type Multiplayer struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *Multiplayer {
	return &Multiplayer{APIClient: apiClient, EndpointExtension: "/multiplayer/"}
}

package Playlists

import (
	"QuaverGo"
)

type Playlists struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *Playlists {
	return &Playlists{APIClient: apiClient, EndpointExtension: "/playlists/"}
}

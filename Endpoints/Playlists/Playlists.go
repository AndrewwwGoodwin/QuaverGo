package Playlists

import "quaverGo"

type Playlists struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *Playlists {
	return &Playlists{APIClient: apiClient, EndpointExtension: "/playlists/"}
}

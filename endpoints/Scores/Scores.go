package Scores

import "quaverGo"

type Scores struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *Scores {
	return &Scores{APIClient: apiClient, EndpointExtension: "/scores/"}
}

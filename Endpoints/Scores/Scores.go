package Scores

import (
	"QuaverGo"
)

type Scores struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *Scores {
	return &Scores{APIClient: apiClient, EndpointExtension: "/scores/"}
}

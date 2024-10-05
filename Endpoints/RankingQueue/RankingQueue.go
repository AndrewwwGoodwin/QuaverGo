package RankingQueue

import (
	"QuaverGo"
)

type RankingQueue struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *RankingQueue {
	return &RankingQueue{APIClient: apiClient, EndpointExtension: "/ranking/"}
}

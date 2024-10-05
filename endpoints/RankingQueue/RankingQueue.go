package RankingQueue

import "quaverGo"

type RankingQueue struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *RankingQueue {
	return &RankingQueue{APIClient: apiClient, EndpointExtension: "/ranking/"}
}

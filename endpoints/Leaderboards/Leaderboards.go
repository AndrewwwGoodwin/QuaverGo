package Leaderboards

import "quaverGo"

type Leaderboards struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *Leaderboards {
	return &Leaderboards{APIClient: apiClient, EndpointExtension: "/leaderboard/"}
}

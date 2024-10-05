package Leaderboards

import (
	"QuaverGo"
)

type Leaderboards struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *Leaderboards {
	return &Leaderboards{APIClient: apiClient, EndpointExtension: "/leaderboard/"}
}

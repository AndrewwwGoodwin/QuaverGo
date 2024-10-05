package ServerStats

import "quaverGo"

type ServerStats struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *ServerStats {
	return &ServerStats{APIClient: apiClient, EndpointExtension: "/server/stats/"}
}

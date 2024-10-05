package ServerStats

import (
	"QuaverGo"
)

type ServerStats struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *ServerStats {
	return &ServerStats{APIClient: apiClient, EndpointExtension: "/server/stats/"}
}

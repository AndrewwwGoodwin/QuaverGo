package Maps

import (
	"QuaverGo"
)

type Maps struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *Maps {
	return &Maps{APIClient: apiClient, EndpointExtension: "/map/"}
}

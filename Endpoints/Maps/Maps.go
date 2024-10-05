package Maps

import "quaverGo"

type Maps struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *Maps {
	return &Maps{APIClient: apiClient, EndpointExtension: "/map/"}
}

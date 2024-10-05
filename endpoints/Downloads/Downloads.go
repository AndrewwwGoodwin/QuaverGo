package Downloads

import "quaverGo"

type Downloads struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *Downloads {
	return &Downloads{APIClient: apiClient, EndpointExtension: "/download/"}
}

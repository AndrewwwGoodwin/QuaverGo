package Downloads

import "QuaverGo"

type Downloads struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *Downloads {
	return &Downloads{APIClient: apiClient, EndpointExtension: "/download/"}
}

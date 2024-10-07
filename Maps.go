package QuaverGo

type Maps struct {
	APIClient         *Client
	EndpointExtension string
}

func initMaps(apiClient *Client) *Maps {
	return &Maps{APIClient: apiClient, EndpointExtension: "/map/"}
}

package QuaverGo

type Mapsets struct {
	APIClient         *Client
	EndpointExtension string
}

func initMapsets(apiClient *Client) *Mapsets {
	return &Mapsets{APIClient: apiClient, EndpointExtension: "/mapset/"}
}

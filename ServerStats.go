package QuaverGo

type ServerStats struct {
	APIClient         *Client
	EndpointExtension string
}

func initServerStats(apiClient *Client) *ServerStats {
	return &ServerStats{APIClient: apiClient, EndpointExtension: "/server/"}
}

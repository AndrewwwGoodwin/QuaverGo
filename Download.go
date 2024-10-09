package QuaverGo

// not sure that I should add these endpoints

type Download struct {
	APIClient         *Client
	EndpointExtension string
}

func initDownload(apiClient *Client) *Download {
	return &Download{APIClient: apiClient, EndpointExtension: "/map/"}
}

package QuaverGo

type Scores struct {
	APIClient         *Client
	EndpointExtension string
}

func initScores(apiClient *Client) *Scores {
	return &Scores{APIClient: apiClient, EndpointExtension: "/scores/"}
}

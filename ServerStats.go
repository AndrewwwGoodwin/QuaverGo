package QuaverGo

import "fmt"

type ServerStats struct {
	APIClient         *Client
	EndpointExtension string
}

func initServerStats(apiClient *Client) *ServerStats {
	return &ServerStats{APIClient: apiClient, EndpointExtension: "/server/"}
}

func (s *ServerStats) GetServerStats() (ServerStatsData, error) {
	var returnData ServerStatsData
	err := fetchData(fmt.Sprintf("%s%sstats", s.APIClient.baseURL, s.EndpointExtension), &returnData)
	if err != nil {
		return ServerStatsData{}, err
	}
	return returnData, nil
}

type ServerStatsData struct {
	OnlineUsers  int `json:"online_users"`
	TotalMapsets int `json:"total_mapsets"`
	TotalScores  int `json:"total_scores"`
	TotalUsers   int `json:"total_users"`
}

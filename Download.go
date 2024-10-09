package QuaverGo

import (
	"fmt"
)

// not sure that I should add these endpoints

type Download struct {
	APIClient         *Client
	EndpointExtension string
}

func initDownload(apiClient *Client) *Download {
	return &Download{APIClient: apiClient, EndpointExtension: "/map/"}
}

func (d *Download) Map(id int) (DownloadData, error) {
	endpoint := fmt.Sprintf("%s%smap/%d", d.APIClient.baseURL, d.EndpointExtension, id)
	var returnData DownloadData
	err := d.APIClient.fetchData(endpoint, &returnData)
	if err != nil {
		return DownloadData{}, err
	}
	return returnData, nil
}

type DownloadData struct {
}

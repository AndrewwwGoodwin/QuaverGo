package QuaverGo

import (
	"fmt"
	"time"
)

type Mapsets struct {
	APIClient         *Client
	EndpointExtension string
}

func initMapsets(apiClient *Client) *Mapsets {
	return &Mapsets{APIClient: apiClient, EndpointExtension: "/mapset/"}
}

func (m Mapsets) ByID(id int) (Mapset, error) {
	var returnedMapset MapsetJson
	err := fetchData(fmt.Sprintf("%s%s%d", m.APIClient.baseURL, m.EndpointExtension, id), &returnedMapset)
	if err != nil {
		return Mapset{}, err
	}
	return returnedMapset.Mapset, nil
}

type Mapset struct {
	Id              int       `json:"id"`
	PackageMd5      string    `json:"package_md5"`
	CreatorId       int       `json:"creator_id"`
	CreatorUsername string    `json:"creator_username"`
	Artist          string    `json:"artist"`
	Title           string    `json:"title"`
	Source          string    `json:"source"`
	Tags            string    `json:"tags"`
	Description     string    `json:"description"`
	DateSubmitted   time.Time `json:"date_submitted"`
	DateLastUpdated time.Time `json:"date_last_updated"`
	IsVisible       bool      `json:"is_visible"`
	IsExplicit      bool      `json:"is_explicit"`
	Maps            []Map     `json:"maps"`
}

type MapsetJson struct {
	Mapset Mapset `json:"mapset"`
}

// RankedList returns an []int of all the ranked maps in quaver
func (m Mapsets) RankedList() ([]int, error) {
	var returnedMapset RankedList
	err := fetchData(fmt.Sprintf("%s%sranked", m.APIClient.baseURL, m.EndpointExtension), &returnedMapset)
	if err != nil {
		return nil, err
	}
	return returnedMapset.RankedMapsets, nil
}

type RankedList struct {
	RankedMapsets []int `json:"ranked_mapsets"`
}

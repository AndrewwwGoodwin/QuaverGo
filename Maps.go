package QuaverGo

import "fmt"

type Maps struct {
	APIClient         *Client
	EndpointExtension string
}

func initMaps(apiClient *Client) *Maps {
	return &Maps{APIClient: apiClient, EndpointExtension: "/map/"}
}

func (m Maps) ByID(id int) (Map, error) {
	var returnedMap MapJson
	err := fetchData(fmt.Sprintf("%s%s%d", m.APIClient.baseURL, m.EndpointExtension, id), &returnedMap)
	if err != nil {
		return Map{}, err
	}
	return returnedMap.Map, nil
}

type MapJson struct {
	Map Map `json:"map"`
}

type Map struct {
	Id                   int     `json:"id"`
	MapsetId             int     `json:"mapset_id"`
	Md5                  string  `json:"md5"`
	AlternativeMd5       string  `json:"alternative_md5"`
	CreatorId            int     `json:"creator_id"`
	CreatorUsername      string  `json:"creator_username"`
	GameMode             int     `json:"game_mode"`
	RankedStatus         int     `json:"ranked_status"`
	Artist               string  `json:"artist"`
	Title                string  `json:"title"`
	Source               string  `json:"source"`
	Tags                 string  `json:"tags"`
	Description          string  `json:"description"`
	DifficultyName       string  `json:"difficulty_name"`
	Length               int     `json:"length"`
	Bpm                  float64 `json:"bpm"`
	DifficultyRating     float64 `json:"difficulty_rating"`
	CountHitobjectNormal int     `json:"count_hitobject_normal"`
	CountHitobjectLong   int     `json:"count_hitobject_long"`
	LongNotePercentage   float64 `json:"long_note_percentage"`
	MaxCombo             int     `json:"max_combo"`
	PlayCount            int     `json:"play_count"`
	FailCount            int     `json:"fail_count"`
	PlayAttempts         int     `json:"play_attempts"`
	ModsPending          int     `json:"mods_pending"`
	ModsAccepted         int     `json:"mods_accepted"`
	ModsDenied           int     `json:"mods_denied"`
	ModsIgnored          int     `json:"mods_ignored"`
	OnlineOffset         int     `json:"online_offset"`
	IsClanRanked         bool    `json:"is_clan_ranked"`
}

package QuaverGo

import (
	"fmt"
	"time"
)

type Maps struct {
	APIClient         *Client
	EndpointExtension string
}

func initMaps(apiClient *Client) *Maps {
	return &Maps{APIClient: apiClient, EndpointExtension: "/map/"}
}

// ByID Retrieves info about a given map
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

// GetMods Retrieves a list of mods on a given map.
func (m Maps) GetMods(id int) (MapModsData, error) {
	var returnedMapMods MapModsData
	err := fetchData(fmt.Sprintf("%s%s%d/mods", m.APIClient.baseURL, m.EndpointExtension, id), &returnedMapMods)
	if err != nil {
		return MapModsData{}, err
	}
	return returnedMapMods, nil
}

type MapModsData struct {
	Mods []struct {
		Id           int       `json:"id"`
		MapId        int       `json:"map_id"`
		AuthorId     int       `json:"author_id"`
		Timestamp    time.Time `json:"timestamp"`
		MapTimestamp *string   `json:"map_timestamp"`
		Comment      string    `json:"comment"`
		Status       string    `json:"status"`
		Type         string    `json:"type"`
		Author       struct {
			Id              int         `json:"id"`
			SteamId         string      `json:"steam_id"`
			Username        string      `json:"username"`
			TimeRegistered  time.Time   `json:"time_registered"`
			Allowed         bool        `json:"allowed"`
			Privileges      int         `json:"privileges"`
			Usergroups      int         `json:"usergroups"`
			MuteEndTime     time.Time   `json:"mute_end_time"`
			LatestActivity  time.Time   `json:"latest_activity"`
			Country         string      `json:"country"`
			AvatarUrl       string      `json:"avatar_url"`
			Twitter         *string     `json:"twitter"`
			Title           string      `json:"title"`
			TwitchUsername  *string     `json:"twitch_username"`
			DonatorEndTime  time.Time   `json:"donator_end_time"`
			DiscordId       *string     `json:"discord_id"`
			MiscInformation interface{} `json:"misc_information"`
			ClanId          *int        `json:"clan_id"`
			ClanLeaveTime   time.Time   `json:"clan_leave_time"`
			ClientStatus    interface{} `json:"client_status"`
		} `json:"author"`
		Replies []struct {
			Id        int       `json:"id"`
			MapModId  int       `json:"map_mod_id"`
			AuthorId  int       `json:"author_id"`
			Timestamp time.Time `json:"timestamp"`
			Comments  string    `json:"comments"`
			Spam      bool      `json:"spam"`
			Author    struct {
				Id              int       `json:"id"`
				SteamId         string    `json:"steam_id"`
				Username        string    `json:"username"`
				TimeRegistered  time.Time `json:"time_registered"`
				Allowed         bool      `json:"allowed"`
				Privileges      int       `json:"privileges"`
				Usergroups      int       `json:"usergroups"`
				MuteEndTime     time.Time `json:"mute_end_time"`
				LatestActivity  time.Time `json:"latest_activity"`
				Country         string    `json:"country"`
				AvatarUrl       string    `json:"avatar_url"`
				Twitter         *string   `json:"twitter"`
				Title           string    `json:"title"`
				TwitchUsername  *string   `json:"twitch_username"`
				DonatorEndTime  time.Time `json:"donator_end_time"`
				DiscordId       *string   `json:"discord_id"`
				MiscInformation *struct {
					DefaultMode       int    `json:"default_mode,omitempty"`
					Twitch            string `json:"twitch,omitempty"`
					NotifActionMapset bool   `json:"notif_action_mapset,omitempty"`
					Discord           string `json:"discord,omitempty"`
					Twitter           string `json:"twitter,omitempty"`
					Youtube           string `json:"youtube,omitempty"`
				} `json:"misc_information"`
				ClanId        *int        `json:"clan_id"`
				ClanLeaveTime time.Time   `json:"clan_leave_time"`
				ClientStatus  interface{} `json:"client_status"`
			} `json:"author"`
		} `json:"replies"`
	} `json:"mods"`
}

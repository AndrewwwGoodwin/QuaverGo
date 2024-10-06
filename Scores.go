package QuaverGo

import (
	"fmt"
	"time"
)

type Scores struct {
	APIClient         *Client
	EndpointExtension string
}

func initScores(apiClient *Client) *Scores {
	return &Scores{APIClient: apiClient, EndpointExtension: "/scores/"}
}

// GetMapLeaderboard gets the top 50 scores global leaderboard scores from a map's given md5 id
func (s Scores) GetMapLeaderboard(mapMD5 string) ([]MapScores, error) {
	var returnedScores MapScoreJson
	err := fetchData(fmt.Sprintf("%s%s%s/global", s.APIClient.baseURL, s.EndpointExtension, mapMD5), &returnedScores)
	if err != nil {
		return nil, err
	}
	return returnedScores.Scores, nil
}

type MapScoreJson struct {
	Scores []MapScores `json:"scores"`
}

type MapScores struct {
	Id                int         `json:"id"`
	UserId            int         `json:"user_id"`
	MapMd5            string      `json:"map_md5"`
	ReplayMd5         string      `json:"replay_md5"`
	Mode              int         `json:"mode"`
	Timestamp         time.Time   `json:"timestamp"`
	IsPersonalBest    bool        `json:"is_personal_best"`
	PerformanceRating float64     `json:"performance_rating"`
	Modifiers         int64       `json:"modifiers"`
	Failed            bool        `json:"failed"`
	TotalScore        int         `json:"total_score"`
	Accuracy          float64     `json:"accuracy"`
	MaxCombo          int         `json:"max_combo"`
	CountMarvelous    int         `json:"count_marvelous"`
	CountPerfect      int         `json:"count_perfect"`
	CountGreat        int         `json:"count_great"`
	CountGood         int         `json:"count_good"`
	CountOkay         int         `json:"count_okay"`
	CountMiss         int         `json:"count_miss"`
	Grade             string      `json:"grade"`
	ScrollSpeed       int         `json:"scroll_speed"`
	IsDonatorScore    bool        `json:"is_donator_score"`
	TournamentGameId  int         `json:"tournament_game_id"`
	ClanId            interface{} `json:"clan_id"`
	User              UserCompact `json:"user"`
}

// UserCompact is literally just JsonUser without the 4/7k stats
type UserCompact struct {
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
	Twitter         interface{} `json:"twitter"`
	Title           string      `json:"title"`
	TwitchUsername  interface{} `json:"twitch_username"`
	DonatorEndTime  time.Time   `json:"donator_end_time"`
	DiscordId       string      `json:"discord_id"`
	MiscInformation interface{} `json:"misc_information"`
	ClanId          interface{} `json:"clan_id"`
	ClanLeaveTime   time.Time   `json:"clan_leave_time"`
	ClientStatus    interface{} `json:"client_status"`
}

// GetMapModLeaderboard takes in a map MD5 and a mod id, and returns the leaderboard for the given mod.
func (s Scores) GetMapModLeaderboard(mapMD5 string, mods int) ([]MapScores, error) {
	var returnedScores MapScoreJson
	err := fetchData(fmt.Sprintf("%s%s%s/mods/%d", s.APIClient.baseURL, s.EndpointExtension, mapMD5, mods), &returnedScores)
	if err != nil {
		return nil, err
	}
	return returnedScores.Scores, nil
}

// GetMapRateLeaderboard takes in a mapMD5 and a rate mod id, and returns the leaderboard for the given rate.
func (s Scores) GetMapRateLeaderboard(mapMD5 string, mods int) ([]MapScores, error) {
	var returnedScores MapScoreJson
	err := fetchData(fmt.Sprintf("%s%s%s/rate/%d", s.APIClient.baseURL, s.EndpointExtension, mapMD5, mods), &returnedScores)
	if err != nil {
		return nil, err
	}
	return returnedScores.Scores, nil
}

// GetUserBestScore returns a user's single best global score on a provided beatmap. Takes in mapMD5 and userID
func (s Scores) GetUserBestScore(mapMD5 string, userId int) (MapScore, error) {
	var returnedScore MapScore
	err := fetchData(fmt.Sprintf("%s%s%s/%d/global", s.APIClient.baseURL, s.EndpointExtension, mapMD5, userId), &returnedScore)
	if err != nil {
		return MapScore{}, err
	}
	return returnedScore, nil
}

type MapScore struct {
	Score MapScores `json:"score"`
}

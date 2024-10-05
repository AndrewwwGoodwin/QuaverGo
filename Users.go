package QuaverGo

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Users struct {
	APIClient         *Client
	EndpointExtension string
}

func initUsers(apiClient *Client) *Users {
	return &Users{APIClient: apiClient, EndpointExtension: "/user/"}
}

func (u Users) ID(id int) (*JsonUser, error) {

	responseData, err := u.APIClient.AttemptRequest(fmt.Sprintf("%s%s%d", u.APIClient.baseURL, u.EndpointExtension, id))
	if err != nil {
		return nil, err
	}
	defer responseData.Body.Close()
	dataStream, err := io.ReadAll(responseData.Body)
	if err != nil {
		return nil, err
	}
	var UserData QuaverUser
	err = json.Unmarshal(dataStream, &UserData)
	if err != nil {
		return nil, err
	}

	return &UserData.User, nil
}

func (u Users) Achievements(id int) (*[]Achievement, error) {
	responseData, err := u.APIClient.AttemptRequest(fmt.Sprintf("%s%s%d/achievements", u.APIClient.baseURL, u.EndpointExtension, id))
	if err != nil {
		return nil, err
	}
	defer responseData.Body.Close()
	dataStream, err := io.ReadAll(responseData.Body)
	if err != nil {
		return nil, err
	}
	var ReturnedAchievements AchievementJson
	err = json.Unmarshal(dataStream, &ReturnedAchievements)
	if err != nil {
		return nil, err
	}

	return &ReturnedAchievements.Achievements, nil
}

func (u Users) Activity(id int) (*[]Activity, error) {
	responseData, err := u.APIClient.AttemptRequest(fmt.Sprintf("%s%s%d/activity", u.APIClient.baseURL, u.EndpointExtension, id))
	if err != nil {
		return nil, err
	}
	defer responseData.Body.Close()
	dataStream, err := io.ReadAll(responseData.Body)
	if err != nil {
		return nil, err
	}
	var ReturnedActivities ActivityJSON
	err = json.Unmarshal(dataStream, &ReturnedActivities)
	if err != nil {
		return nil, err
	}

	return &ReturnedActivities.Activities, nil
}

type ActivityJSON struct {
	Activities []Activity `json:"activities"`
}

type Activity struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Type      int       `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Value     string    `json:"value"`
	MapsetID  int       `json:"mapset_id"`
}

type AchievementJson struct {
	Achievements []Achievement `json:"achievements"`
}

type Achievement struct {
	ID           int    `json:"id"`
	Difficulty   string `json:"difficulty"`
	SteamAPIName string `json:"steam_api_name"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsUnlocked   bool   `json:"is_unlocked"`
}

type QuaverUser struct {
	User JsonUser `json:"user"`
}

type JsonUser struct {
	Id              int             `json:"id"`
	SteamId         string          `json:"steam_id"`
	Username        string          `json:"username"`
	TimeRegistered  time.Time       `json:"time_registered"`
	Allowed         bool            `json:"allowed"`
	Privileges      int             `json:"privileges"`
	Usergroups      int             `json:"usergroups"`
	MuteEndTime     time.Time       `json:"mute_end_time"`
	LatestActivity  time.Time       `json:"latest_activity"`
	Country         string          `json:"country"`
	AvatarUrl       string          `json:"avatar_url"`
	Twitter         interface{}     `json:"twitter"`
	Title           interface{}     `json:"title"`
	TwitchUsername  interface{}     `json:"twitch_username"`
	DonatorEndTime  time.Time       `json:"donator_end_time"`
	DiscordId       interface{}     `json:"discord_id"`
	MiscInformation interface{}     `json:"misc_information"`
	ClanId          interface{}     `json:"clan_id"`
	ClanLeaveTime   time.Time       `json:"clan_leave_time"`
	ClientStatus    interface{}     `json:"client_status"`
	StatsKeys4      QuaverUserStats `json:"stats_keys4"`
	StatsKeys7      QuaverUserStats `json:"stats_key7"`
}

type QuaverUserStatsRanks struct {
	Global    int `json:"global"`
	Country   int `json:"country"`
	TotalHits int `json:"total_hits"`
}

type QuaverUserStats struct {
	Ranks                    QuaverUserStatsRanks `json:"ranks"`
	TotalScore               int                  `json:"total_score"`
	RankedScore              int                  `json:"ranked_score"`
	OverallAccuracy          float64              `json:"overall_accuracy"`
	OverallPerformanceRating float64              `json:"overall_performance_rating"`
	PlayCount                int                  `json:"play_count"`
	FailCount                int                  `json:"fail_count"`
	MaxCombo                 int                  `json:"max_combo"`
	TotalMarvelous           int                  `json:"total_marvelous"`
	TotalPerfect             int                  `json:"total_perfect"`
	TotalGreat               int                  `json:"total_great"`
	TotalGood                int                  `json:"total_good"`
	TotalOkay                int                  `json:"total_okay"`
	TotalMiss                int                  `json:"total_miss"`
	CountGradeX              int                  `json:"count_grade_x"`
	CountGradeSs             int                  `json:"count_grade_ss"`
	CountGradeS              int                  `json:"count_grade_s"`
	CountGradeA              int                  `json:"count_grade_a"`
	CountGradeB              int                  `json:"count_grade_b"`
	CountGradeC              int                  `json:"count_grade_c"`
	CountGradeD              int                  `json:"count_grade_d"`
}

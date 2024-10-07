package QuaverGo

import (
	"fmt"
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
	var UserData QuaverUser
	err := fetchData(fmt.Sprintf("%s%s%d", u.APIClient.baseURL, u.EndpointExtension, id), &UserData)
	if err != nil {
		return nil, err
	}

	return &UserData.User, nil
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
	DiscordId       string          `json:"discord_id"`
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

func (u Users) Achievements(id int) (*[]Achievement, error) {
	var ReturnedAchievements AchievementJson
	err := fetchData(fmt.Sprintf("%s%s%d/achievements", u.APIClient.baseURL, u.EndpointExtension, id), &ReturnedAchievements)
	if err != nil {
		return nil, err
	}

	return &ReturnedAchievements.Achievements, nil
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

func (u Users) Activity(id int) (*[]Activity, error) {
	var ReturnedActivities ActivityJSON
	err := fetchData(fmt.Sprintf("%s%s%d/activity", u.APIClient.baseURL, u.EndpointExtension, id), &ReturnedActivities)
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

func (u Users) Badges(id int) (*[]Badge, error) {
	var returnedBadges BadgesJSON
	err := fetchData(fmt.Sprintf("%s%s%d/badges", u.APIClient.baseURL, u.EndpointExtension, id), &returnedBadges)
	if err != nil {
		return nil, err
	}
	return &returnedBadges.Badges, nil
}

type BadgesJSON struct {
	Badges []Badge `json:"badges"`
}

type Badge struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (u Users) Mapsets(id int) (*[]Mapset, error) {
	var returnedMapsets MapsetJSON
	err := fetchData(fmt.Sprintf("%s%s%d/mapsets", u.APIClient.baseURL, u.EndpointExtension, id), &returnedMapsets)
	if err != nil {
		return nil, err
	}
	return &returnedMapsets.Mapsets, nil
}

type MapsetJSON struct {
	Mapsets []Mapset `json:"mapsets"`
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

func (u Users) Playlists(id int) (*[]Playlists, error) {
	var returnedPlaylists PlaylistsJson
	err := fetchData(fmt.Sprintf("%s%s%d/playlists", u.APIClient.baseURL, u.EndpointExtension, id), &returnedPlaylists)
	if err != nil {
		return nil, err
	}
	return &returnedPlaylists.Playlists, nil
}

type PlaylistsJson struct {
	Playlists []Playlists `json:"playlists"`
}

type Playlists struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	MapCount        int       `json:"map_count"`
	Timestamp       time.Time `json:"timestamp"`
	TimeLastUpdated time.Time `json:"time_last_updated"`
}

// ScoresBest valid modes are 1 (4k) and 2 (7k)
func (u Users) ScoresBest(id int, mode int) (*[]ScoresData, error) {
	var returnedScores ScoresJSON
	err := fetchData(fmt.Sprintf("%s%s%d/scores/%d/best", u.APIClient.baseURL, u.EndpointExtension, id, mode), &returnedScores)
	if err != nil {
		return nil, err
	}
	return &returnedScores.Scores, nil
}

func (u Users) ScoresRecent(id int, mode int) (*[]ScoresData, error) {
	var returnedScores ScoresJSON
	err := fetchData(fmt.Sprintf("%s%s%d/scores/%d/recent", u.APIClient.baseURL, u.EndpointExtension, id, mode), &returnedScores)
	if err != nil {
		return nil, err
	}
	return &returnedScores.Scores, nil
}

func (u Users) ScoresFirstPlace(id int, mode int) (*[]ScoresData, error) {
	var returnedScores ScoresJSON
	err := fetchData(fmt.Sprintf("%s%s%d/scores/%d/firstplace", u.APIClient.baseURL, u.EndpointExtension, id, mode), &returnedScores)
	if err != nil {
		return nil, err
	}
	return &returnedScores.Scores, nil
}

// ScoresByGrade valid grades are X, SS, S, A, B, C, D
// valid modes are (1, 4k), or (2, 7k)
func (u Users) ScoresByGrade(id int, mode int, grade string) (*[]ScoresData, error) {
	var returnedScores ScoresJSON
	err := fetchData(fmt.Sprintf("%s%s%d/scores/%d/grades/%s", u.APIClient.baseURL, u.EndpointExtension, id, mode, grade), &returnedScores)
	if err != nil {
		return nil, err
	}
	return &returnedScores.Scores, nil
}

type ScoresJSON struct {
	Scores []ScoresData `json:"scores"`
}
type ScoresData struct {
	Id                int         `json:"id"`
	UserId            int         `json:"user_id"`
	MapMd5            string      `json:"map_md5"`
	ReplayMd5         string      `json:"replay_md5"`
	Timestamp         time.Time   `json:"timestamp"`
	IsPersonalBest    bool        `json:"is_personal_best"`
	PerformanceRating float64     `json:"performance_rating"`
	Modifiers         int         `json:"modifiers"`
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
	TournamentGameId  interface{} `json:"tournament_game_id"`
	ClanId            interface{} `json:"clan_id"`
	Map               Map         `json:"map"`
}

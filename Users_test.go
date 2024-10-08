package QuaverGo

import (
	"testing"
)

// attempts to get a user's profile from their user id
func TestGetUserByID(t *testing.T) {
	client := Init()
	returnedUser, err := client.Users.ID(371737)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUser)
	//fmt.Println(id)
}

// attempts to get a user's achievements from their user id
func TestGetUserAchievementsByID(t *testing.T) {
	client := Init()
	returnedUser, err := client.Users.Achievements(371737)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUser)
}

func TestGetUserActivityByID(t *testing.T) {
	client := Init()
	returnedUserActivity, err := client.Users.Activity(371737)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUserActivity)
}

func TestGetUserBadgesByID(t *testing.T) {
	client := Init()
	returnedBadges, err := client.Users.Badges(608)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedBadges)
}

func TestGetUserMapsetsByID(t *testing.T) {
	client := Init()
	returnedUserMapsets, err := client.Users.Mapsets(608)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUserMapsets)
}

func TestGetUserPlaylistsByID(t *testing.T) {
	client := Init()
	returnedUserPlaylists, err := client.Users.Playlists(398945)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUserPlaylists)
}

func TestGetUserScoresBestByID(t *testing.T) {
	client := Init()
	returnedUserScores, err := client.Users.ScoresBest(371737, 1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUserScores)
}

func TestGetUserScoresRecentByID(t *testing.T) {
	client := Init()
	returnedUserScores, err := client.Users.ScoresRecent(371737, 1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUserScores)
}

func TestGetUserScoresFirstPlaceByID(t *testing.T) {
	client := Init()
	returnedUserScores, err := client.Users.ScoresFirstPlace(371737, 1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUserScores)
}

func TestGetUserScoresByGradeByID(t *testing.T) {
	client := Init()
	returnedUserScores, err := client.Users.ScoresByGrade(371737, 1, "s")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUserScores)
}

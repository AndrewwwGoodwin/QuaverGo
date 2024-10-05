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
	returnedUser, err := client.Users.Activity(371737)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUser)
}

func TestGetUserBadgesByID(t *testing.T) {
	client := Init()
	returnedUser, err := client.Users.Badges(608)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(*returnedUser)
}

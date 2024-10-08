package QuaverGo

import "testing"

// attempts to get a user's profile from their user id
func TestScores_GetMapLeaderboard(t *testing.T) {
	client := Init()
	returnedScores, err := client.Scores.GetMapLeaderboard("0ec514e829ae12273cf166f65ad53d25")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(returnedScores[0])
	//fmt.Println(id)
}

func TestScores_GetMapModLeaderboard(t *testing.T) {
	client := Init()
	returnedScores, err := client.Scores.GetMapModLeaderboard("0ec514e829ae12273cf166f65ad53d25", 4)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(returnedScores[0])
}

func TestScores_GetMapRateLeaderboard(t *testing.T) {
	client := Init()
	returnedScores, err := client.Scores.GetMapRateLeaderboard("0ec514e829ae12273cf166f65ad53d25", 128)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(returnedScores[0])
}

func TestScores_GetBeatmapUserBestScore(t *testing.T) {
	client := Init()
	returnedScore, err := client.Scores.GetBeatmapUserBestScore("0ec514e829ae12273cf166f65ad53d25", 371737)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(returnedScore)
}

func TestScores_GetUserBestAllScoreboards(t *testing.T) {
	client := Init()
	returnedScore, err := client.Scores.GetUserBestAllScoreboards("0ec514e829ae12273cf166f65ad53d25", 371737)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(returnedScore)
}

// might not work, not sure.
//
// Alright I found where they hide all the mod identifiers https://github.com/Quaver/Quaver.API/blob/master/Quaver.API/Enums/ModIdentifier.cs
func TestScores_GetBeatmapUserBestScoreMods(t *testing.T) {
	client := Init()
	returnedScore, err := client.Scores.GetBeatmapUserBestScoreMods("0ec514e829ae12273cf166f65ad53d25", 371737, -1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(returnedScore)
}

// might not work, not sure. todo figure out how rates work since i have no clue
func TestScores_GetBeatmapUserBestScoreRate(t *testing.T) {
	client := Init()
	returnedScore, err := client.Scores.GetBeatmapUserBestScoreRate("0ec514e829ae12273cf166f65ad53d25", 371737, 8)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(returnedScore)
}

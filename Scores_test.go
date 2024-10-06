package QuaverGo

import "testing"

// attempts to get a user's profile from their user id
func TestScores_GetMapLeaderboardByMD5(t *testing.T) {
	client := Init()
	returnedScores, err := client.Scores.GetMapLeaderboardByMD5("0ec514e829ae12273cf166f65ad53d25")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(returnedScores[0])
	//fmt.Println(id)
}

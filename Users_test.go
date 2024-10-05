package QuaverGo

import (
	"testing"
)

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

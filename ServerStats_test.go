package QuaverGo

import "testing"

func TestServerStats_GetServerStats(t *testing.T) {
	client := Init()
	stats, err := client.ServerStats.GetServerStats()
	if err != nil {
		t.Error(err)
	}
	t.Log(stats)
}

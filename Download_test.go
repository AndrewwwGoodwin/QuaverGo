package QuaverGo

import "testing"

func TestDownload_Map(t *testing.T) {
	client := Init()
	data, err := client.Download.Map(123)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

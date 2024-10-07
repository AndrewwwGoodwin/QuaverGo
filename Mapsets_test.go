package QuaverGo

import "testing"

func TestMapsets_ByID(t *testing.T) {
	client := Init()
	returnedMapset, err := client.Mapsets.ByID(35323)
	if err != nil {
		t.Errorf("Mapsets.ByID() returned error: %v", err)
	}
	t.Log(returnedMapset)
}

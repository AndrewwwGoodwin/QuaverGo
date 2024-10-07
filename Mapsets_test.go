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

func TestMapsets_RankedList(t *testing.T) {
	client := Init()
	returnedMapset, err := client.Mapsets.RankedList()
	if err != nil {
		t.Errorf("Mapsets.RankedList() returned error: %v", err)
	}
	t.Log(returnedMapset)
}

func TestMapsets_OffsetList(t *testing.T) {
	client := Init()
	returnedOffsets, err := client.Mapsets.OffsetList()
	if err != nil {
		t.Errorf("Mapsets.OffsetList() returned error: %v", err)
	}
	t.Log(returnedOffsets)
}

package QuaverGo

import "testing"

func TestMaps_ByID(t *testing.T) {
	client := Init()
	returnedMap, err := client.Maps.ByID(149325)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(returnedMap)
}

func TestMaps_GetMods(t *testing.T) {
	client := Init()
	returnedMapMods, err := client.Maps.GetMods(380)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(returnedMapMods)
}

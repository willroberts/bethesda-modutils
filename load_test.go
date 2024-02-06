package modutils

import "testing"

func TestLoad(t *testing.T) {
	modFile, err := Load("testdata/fallout4-perks.esp")
	if err != nil {
		t.Error(err)
	}

	modFile.Metadata.Print()
}

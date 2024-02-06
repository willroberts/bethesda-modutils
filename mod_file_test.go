package modutils

import "testing"

func TestLoadModFile(t *testing.T) {
	modFile, err := LoadModFile("testdata/fallout4-perks.esp")
	if err != nil {
		t.Error(err)
	}

	modFile.Metadata.Print()
}

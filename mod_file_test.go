package modutils

import "testing"

func TestLoadModFile(t *testing.T) {
	modFile, err := LoadModFile("testdata/fallout4-perks.esp")
	if err != nil {
		t.Error(err)
	}

	modFile.Metadata.Print()
	for _, f := range modFile.Metadata.Fields {
		f.Print()
	}
	for _, g := range modFile.Groups {
		g.Print()
	}
}

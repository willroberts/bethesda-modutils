package modutils

import (
	"testing"
)

func TestLoadModFile(t *testing.T) {
	_, err := LoadModFile("testdata/fallout4-perks.esp")
	if err != nil {
		t.Fatal(err)
	}
}

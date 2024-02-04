package modutils

import "testing"

func TestLoad(t *testing.T) {
	_, err := Load("testdata/fallout4-perks.esp")
	if err != nil {
		t.Error(err)
	}
}

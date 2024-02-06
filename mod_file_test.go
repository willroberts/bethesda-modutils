package modutils

import (
	"fmt"
	"testing"
)

func TestLoadModFile(t *testing.T) {
	modFile, err := LoadModFile("testdata/fallout4-perks.esp")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("== Mod Metadata ==")
	modFile.Metadata.Print()
	fmt.Println("== Metadata Fields ==")
	for _, f := range modFile.Metadata.Fields {
		f.Print()
	}
	fmt.Println("== Mod Groups ==")
	for _, g := range modFile.Groups {
		g.Print()
		fmt.Println("== Group Records ==")
		for _, r := range g.Records {
			r.Print()
			fmt.Println("== Record Fields ==")
			for _, f := range r.Fields {
				f.Print()
			}
		}
	}
}

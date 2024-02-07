package main

import (
	"flag"
	"fmt"
	"log"

	modutils "github.com/willroberts/bethesda-modutils"
)

var espPath string

func init() {
	flag.StringVar(&espPath, "esp-path", "", "Path to ESM, ESP, or ESL file")
	flag.Parse()
}

func main() {
	if espPath == "" {
		log.Fatal("esp-path may not be blank")
	}

	mod, err := modutils.LoadModFile(espPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mod File:", espPath)
	printRecord(mod.Metadata)
	for _, f := range mod.Metadata.Fields {
		f.Print()
	}
	for _, g := range mod.Groups {
		if string(g.Label) == "PERK" {
			continue // Skip perks for a sec.
		}
		fmt.Println("Group:", string(g.Label))
		for _, r := range g.Records {
			printRecord(r)
			for _, f := range r.Fields {
				f.Print()
			}
		}
	}
}

func printRecord(r *modutils.Record) {
	fmt.Printf(
		"- Record %s has FormID %d, Flags %d, and Version %d\n",
		string(r.Type),
		r.FormID,
		r.Flags,
		r.Version,
	)
}

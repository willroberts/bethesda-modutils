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
		printField(f)
	}
	for _, g := range mod.Groups {
		if string(g.Label) == "PERK" {
			continue // Skip perks for a sec.
		}
		fmt.Println("Group:", string(g.Label))
		for _, r := range g.Records {
			printRecord(r)
			for _, f := range r.Fields {
				printField(f)
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

// TODO: Move to espcat after removing private field/method access.
func printField(f *modutils.Field) {
	fType := string(f.Type)
	switch fType {
	case "ANAM":
		//fmt.Printf("  - %s field has value: %s", fType, f.StringValue)
	case "EDID":
		fmt.Printf("  - %s field has value: %s\n", fType, f.StringValue)
	case "FULL":
		fmt.Printf("  - %s field has value: %s\n", fType, f.StringValue)
	case "DESC":
		fmt.Printf("  - %s field has value: %s\n", fType, f.StringValue)
	case "CNAM":
		fmt.Printf("  - %s field has value: %s\n", fType, f.StringValue)
	case "MAST":
		fmt.Printf("  - %s field has value: %s\n", fType, f.StringValue)
	case "INTV":
		fmt.Printf("  - %s field has value %d\n", fType, f.Uint16Value)
	default:
		//fmt.Printf("  - %s field has binary value %v\n", fType, f.BinaryValue)
	}
}

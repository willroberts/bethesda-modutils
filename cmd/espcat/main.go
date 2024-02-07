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

	fmt.Println("=== Mod Metadata ===")
	mod.Metadata.Print()
	fmt.Println("=== Metadata Fields ===")
	for _, f := range mod.Metadata.Fields {
		f.Print()
	}
	fmt.Println("=== Mod Groups ===")
	for _, g := range mod.Groups {
		g.Print()
		fmt.Println("=== Group Records ===")
		for _, r := range g.Records {
			r.Print()
			fmt.Println("=== Record Fields ===")
			for _, f := range r.Fields {
				f.Print()
			}
		}
	}
}

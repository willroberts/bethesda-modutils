package modutils

import (
	"bytes"
	"fmt"
	"io"
)

type Field struct {
	Type []byte
	Size uint16

	rawData []byte
}

func ReadField(r io.Reader) (*Field, error) {
	field := &Field{}
	var err error

	field.Type, err = readBytes(4, r)
	if err != nil {
		return nil, err
	}

	field.Size, err = readUint16(r)
	if err != nil {
		return nil, err
	}

	field.rawData, err = readBytes(uint(field.Size), r)
	if err != nil {
		return nil, err
	}

	return field, nil
}

// TODO: Move to espcat after removing private field/method access.
func (f *Field) Print() error {
	fType := string(f.Type)
	r := bytes.NewReader(f.rawData)

	switch fType {
	case "HEDR": // Header.
		fmt.Printf("  - %s field has value: ", fType)
		fmt.Println(f.rawData) // FIXME: Parsing. Are these flags? 12 bytes.
	case "ANAM": // Abbreviated Name.
		//fmt.Printf("  - %s field has value: ", fType)
		//fmt.Println(string(f.rawData))
	case "EDID": // Specific ID.
		fmt.Printf("  - %s field has value: ", fType)
		fmt.Println(string(f.rawData))
	case "FULL": // Full Name.
		fmt.Printf("  - %s field has value: ", fType)
		fmt.Println(string(f.rawData))
	case "DESC":
		fmt.Printf("  - %s field has value: ", fType)
		fmt.Println(string(f.rawData))
	case "CNAM":
		fmt.Printf("  - %s field has value: ", fType)
		fmt.Println(string(f.rawData))
	case "MAST": // Master File.
		fmt.Printf("  - %s field has value: ", fType)
		fmt.Println(string(f.rawData))
	case "INTV": // Internal Version.
		fmt.Printf("  - %s field has value: ", fType)
		v, err := readUint16(r)
		if err != nil {
			return err
		}
		fmt.Println(v)
	default:
		//fmt.Printf("  - %s field has value: ", fType)
		//fmt.Println(f.rawData)
	}

	return nil
}

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

func (f *Field) Print() error {
	fType := string(f.Type)
	fmt.Printf("%s field has value: ", fType)
	r := bytes.NewReader(f.rawData)

	switch fType {
	case "HEDR": // Header.
		fmt.Println(f.rawData) // FIXME: Parsing. Are these flags? 12 bytes.
	case "CNAM":
		fmt.Println(string(f.rawData))
	case "MAST": // Master File.
		fmt.Println(string(f.rawData))
	case "DATA": // Unused?
		v, err := readUint32(r)
		if err != nil {
			return err
		}
		fmt.Println(v)
	case "INTV": // Internal Version?
		v, err := readUint16(r)
		if err != nil {
			return err
		}
		fmt.Println(v)
	default:
		fmt.Println("Unknown field type", fType)
	}

	return nil
}

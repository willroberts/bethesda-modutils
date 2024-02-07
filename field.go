package modutils

import (
	"bytes"
	"io"
)

type Field struct {
	Type []byte
	Size uint16

	StringValue string
	Uint16Value uint16
	BinaryValue []byte

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

	if err := field.ParseValue(); err != nil {
		return nil, err
	}

	return field, nil
}

func (f *Field) ParseValue() error {
	r := bytes.NewReader(f.rawData)

	switch string(f.Type) {
	case "ANAM": // Abbreviated Name.
		f.StringValue = string(f.rawData)
	case "EDID": // Specific ID.
		f.StringValue = string(f.rawData)
	case "FULL": // Full Name.
		f.StringValue = string(f.rawData)
	case "DESC": // Description.
		f.StringValue = string(f.rawData)
	case "CNAM":
		f.StringValue = string(f.rawData)
	case "MAST": // Master File.
		f.StringValue = string(f.rawData)
	case "INTV": // Internal Version.
		v, err := readUint16(r)
		if err != nil {
			return err
		}
		f.Uint16Value = v
	default:
		// Expose the raw bytes for now, until we determine how to parse.
		f.BinaryValue = f.rawData
	}

	return nil
}

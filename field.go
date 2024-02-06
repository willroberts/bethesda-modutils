package modutils

import (
	"fmt"
	"io"
)

type Field struct {
	Type []byte
	Size uint16
	Data []byte
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

	field.Data, err = readBytes(uint(field.Size), r)
	if err != nil {
		return nil, err
	}

	return field, nil
}

func (f *Field) Print() {
	fmt.Println("==========")
	fmt.Println("Field Type:", string(f.Type))
	fmt.Println("Field Size:", f.Size)
	//fmt.Println("Field Data:", string(f.Data))
}

package modutils

import (
	"bytes"
	"fmt"
	"io"
)

type Record struct {
	Type      []byte
	Size      uint32
	Flags     uint32
	FormID    uint32
	Timestamp uint16
	VCSInfo   uint16
	Version   uint16
	Unknown   uint16
	Data      []byte
}

func ReadRecord(r io.Reader) (*Record, error) {
	record := &Record{}
	var err error

	record.Type, err = readBytes(4, r)
	if err != nil {
		return nil, err
	}

	record.Size, err = readUint32(r)
	if err != nil {
		return nil, err
	}

	record.Flags, err = readUint32(r)
	if err != nil {
		return nil, err
	}

	record.FormID, err = readUint32(r)
	if err != nil {
		return nil, err
	}

	record.Timestamp, err = readUint16(r)
	if err != nil {
		return nil, err
	}

	record.VCSInfo, err = readUint16(r)
	if err != nil {
		return nil, err
	}

	record.Version, err = readUint16(r)
	if err != nil {
		return nil, err
	}

	record.Unknown, err = readUint16(r)
	if err != nil {
		return nil, err
	}

	record.Data, err = readBytes(uint(record.Size), r)
	if err != nil {
		return nil, err
	}

	// Testing field parser.
	record.Print()
	field, err := ReadField(bytes.NewReader(record.Data))
	if err != nil {
		return nil, err
	}
	field.Print()

	return record, nil
}

func (r *Record) Print() {
	fmt.Println("Record Type:", string(r.Type))
	fmt.Println("Record Size:", r.Size)
	fmt.Println("Record Flags:", r.Flags)
	fmt.Println("Record FormID:", r.FormID)
	fmt.Println("Record Timestamp:", r.Timestamp)
	fmt.Println("Record VCSInfo:", r.VCSInfo)
	fmt.Println("Record Version:", r.Version)
	fmt.Println("Record Data:", string(r.Data))
}

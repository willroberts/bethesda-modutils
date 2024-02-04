package modutils

import (
	"encoding/binary"
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

	record.Data, err = readBytes(record.Size, r)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func PrintRecord(r *Record) {
	fmt.Println("Record Type:", string(r.Type))
	fmt.Println("Record Size:", r.Size)
	fmt.Println("Record Flags:", r.Flags)
	fmt.Println("Record FormID:", r.FormID)
	fmt.Println("Record Timestamp:", r.Timestamp)
	fmt.Println("Record VCSInfo:", r.VCSInfo)
	fmt.Println("Record Version:", r.Version)
	//fmt.Println("Record Data:", r.Data)
}

func readBytes(n uint32, r io.Reader) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(r, buf)
	return buf, err
}

func readUint32(r io.Reader) (uint32, error) {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(buf), nil
}

func readUint16(r io.Reader) (uint16, error) {
	buf := make([]byte, 2)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(buf), nil
}

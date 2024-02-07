package modutils

import (
	"bytes"
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
	Fields    []*Field

	rawData []byte
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

	record.rawData, err = readBytes(uint(record.Size), r)
	if err != nil {
		return nil, err
	}

	if err := record.readAllFields(); err != nil {
		return nil, err
	}

	return record, nil
}

func (r *Record) readAllFields() error {
	fields := make([]*Field, 0)
	reader := bytes.NewReader(r.rawData)
	var bytesRead uint = 0

	for bytesRead < uint(r.Size) {
		f, err := ReadField(reader)
		if err != nil {
			return err
		}
		bytesRead += (6 + uint(f.Size)) // FieldHeaderSize
		fields = append(fields, f)
	}

	r.Fields = fields
	return nil
}

package modutils

import (
	"bytes"
	"fmt"
	"io"
)

type Group struct {
	RecordType []byte
	Size       uint32
	Label      []byte
	GroupType  []byte
	Timestamp  uint16
	VCSInfo    uint16
	Unknown    uint32
	Records    []*Record

	rawData []byte
}

func ReadGroup(r io.Reader) (*Group, error) {
	g := &Group{}
	var err error

	g.RecordType, err = readBytes(4, r)
	if err != nil {
		return nil, err
	}

	if string(g.RecordType) != "GRUP" {
		return nil, fmt.Errorf("invalid group record type: expected GRUP, got %s", string(g.RecordType))
	}

	s := string(g.RecordType)
	if s != "GRUP" {
		return nil, fmt.Errorf("expected record type GRUP; got %s", s)
	}

	g.Size, err = readUint32(r)
	if err != nil {
		return nil, err
	}

	g.Label, err = readBytes(4, r)
	if err != nil {
		return nil, err
	}

	g.GroupType, err = readBytes(4, r)
	if err != nil {
		return nil, err
	}

	g.Timestamp, err = readUint16(r)
	if err != nil {
		return nil, err
	}

	g.VCSInfo, err = readUint16(r)
	if err != nil {
		return nil, err
	}

	g.Unknown, err = readUint32(r)
	if err != nil {
		return nil, err
	}

	g.rawData, err = readBytes(uint(g.Size)-24, r)
	if err != nil {
		return nil, err
	}

	if err := g.readAllRecords(); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Group) readAllRecords() error {
	records := make([]*Record, 0)
	reader := bytes.NewReader(g.rawData)
	var bytesRead uint = 0

	for bytesRead < uint(g.Size)-24 {
		r, err := ReadRecord(reader)
		if err != nil {
			return err
		}
		bytesRead += (24 + uint(r.Size)) // RecordHeaderSize
		records = append(records, r)
	}

	g.Records = records
	return nil
}
